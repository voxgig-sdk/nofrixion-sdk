package sdktest

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	sdk "github.com/voxgig-sdk/nofrixion-sdk/go"
	"github.com/voxgig-sdk/nofrixion-sdk/go/core"

	vs "github.com/voxgig-sdk/nofrixion-sdk/go/utility/struct"
)

func TestPayrunEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.Payrun(nil)
		if ent == nil {
			t.Fatal("expected non-nil PayrunEntity")
		}
	})

	// Feature #4: the entity Stream(action, ...) method runs the op pipeline and
	// returns a channel over result items. With the streaming feature active it
	// yields the feature's incremental output; otherwise it falls back to the
	// materialised list so Stream always yields.
	t.Run("stream", func(t *testing.T) {
		seed := map[string]any{
			"entity": map[string]any{
				"payrun": map[string]any{
					"s1": map[string]any{"id": "s1"},
					"s2": map[string]any{"id": "s2"},
					"s3": map[string]any{"id": "s3"},
				},
			},
		}

		// Fallback: streaming inactive -> yields the materialised list items.
		base := sdk.TestSDK(seed, nil)
		var seen []any
		for item := range base.Payrun(nil).Stream("list", nil, nil) {
			seen = append(seen, item)
		}
		if len(seen) != 3 {
			t.Fatalf("expected 3 streamed items, got %d", len(seen))
		}

		// Inbound: streaming active -> yields each item from the feature iterator.
		hasStreaming := false
		if fm, ok := core.MakeConfig()["feature"].(map[string]any); ok {
			_, hasStreaming = fm["streaming"]
		}
		if hasStreaming {
			streamSdk := sdk.TestSDK(seed, map[string]any{
				"feature": map[string]any{"streaming": map[string]any{"active": true}},
			})
			var got []any
			for item := range streamSdk.Payrun(nil).Stream("list", nil, nil) {
				if sub, ok := item.([]any); ok {
					got = append(got, sub...)
				} else {
					got = append(got, item)
				}
			}
			if len(got) != 3 {
				t.Fatalf("expected 3 items via streaming feature, got %d", len(got))
			}
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := payrunBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "update", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "payrun." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		// The basic flow consumes synthetic IDs from the fixture. In live mode
		// without an *_ENTID env override, those IDs hit the live API and 4xx.
		if setup.syntheticOnly {
			t.Skip("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_PAYRUN_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		payrunRef01Ent := client.Payrun(nil)
		payrunRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "payrun"}, setup.data), "payrun_ref01"))
		payrunRef01Data["merchant_i_d"] = setup.idmap["merchant_i_d01"]

		payrunRef01DataResult, err := payrunRef01Ent.Create(payrunRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		payrunRef01Data = core.ToMapAny(payrunRef01DataResult)
		if payrunRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if payrunRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		payrunRef01Match := map[string]any{}

		payrunRef01ListResult, err := payrunRef01Ent.List(payrunRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		payrunRef01List, payrunRef01ListOk := payrunRef01ListResult.([]any)
		if !payrunRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", payrunRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(payrunRef01List), map[string]any{"id": payrunRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// UPDATE
		payrunRef01DataUp0Up := map[string]any{
			"id": payrunRef01Data["id"],
		}

		payrunRef01MarkdefUp0Name := "authorisation_date"
		payrunRef01MarkdefUp0Value := fmt.Sprintf("Mark01-payrun_ref01_%d", setup.now)
		payrunRef01DataUp0Up[payrunRef01MarkdefUp0Name] = payrunRef01MarkdefUp0Value

		payrunRef01ResdataUp0Result, err := payrunRef01Ent.Update(payrunRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		payrunRef01ResdataUp0 := core.ToMapAny(payrunRef01ResdataUp0Result)
		if payrunRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if payrunRef01ResdataUp0["id"] != payrunRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if payrunRef01ResdataUp0[payrunRef01MarkdefUp0Name] != payrunRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", payrunRef01MarkdefUp0Name, payrunRef01ResdataUp0[payrunRef01MarkdefUp0Name])
		}

		// LOAD
		payrunRef01MatchDt0 := map[string]any{
			"id": payrunRef01Data["id"],
		}
		payrunRef01DataDt0Loaded, err := payrunRef01Ent.Load(payrunRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		payrunRef01DataDt0LoadResult := core.ToMapAny(payrunRef01DataDt0Loaded)
		if payrunRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if payrunRef01DataDt0LoadResult["id"] != payrunRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

		// REMOVE
		payrunRef01MatchRm0 := map[string]any{
			"id": payrunRef01Data["id"],
		}
		_, err = payrunRef01Ent.Remove(payrunRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		payrunRef01MatchRt0 := map[string]any{}

		payrunRef01ListRt0Result, err := payrunRef01Ent.List(payrunRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		payrunRef01ListRt0, payrunRef01ListRt0Ok := payrunRef01ListRt0Result.([]any)
		if !payrunRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", payrunRef01ListRt0Result)
		}

		notFoundItem := vs.Select(entityListToData(payrunRef01ListRt0), map[string]any{"id": payrunRef01Data["id"]})
		if !vs.IsEmpty(notFoundItem) {
			t.Fatal("expected removed entity to not be in list")
		}

	})
}

func payrunBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "payrun", "PayrunTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read payrun test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse payrun test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"payrun01", "payrun02", "payrun03", "merchant_i_d01"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Detect ENTID env override before envOverride consumes it. When live
	// mode is on without a real override, the basic test runs against synthetic
	// IDs from the fixture and 4xx's. Surface this so the test can skip.
	entidEnvRaw := os.Getenv("NOFRIXION_TEST_PAYRUN_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"NOFRIXION_TEST_PAYRUN_ENTID": idmap,
		"NOFRIXION_TEST_LIVE":      "FALSE",
		"NOFRIXION_TEST_EXPLAIN":   "FALSE",
		"NOFRIXION_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["NOFRIXION_TEST_PAYRUN_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["NOFRIXION_TEST_LIVE"] == "TRUE" {
		mergedOpts := vs.Merge([]any{
			map[string]any{
				"apikey": env["NOFRIXION_APIKEY"],
			},
			extra,
		})
		client = sdk.NewNofrixionSDK(core.ToMapAny(mergedOpts))
	}

	live := env["NOFRIXION_TEST_LIVE"] == "TRUE"
	return &entityTestSetup{
		client:        client,
		data:          entityData,
		idmap:         idmapResolved,
		env:           env,
		explain:       env["NOFRIXION_TEST_EXPLAIN"] == "TRUE",
		live:          live,
		syntheticOnly: live && !idmapOverridden,
		now:           time.Now().UnixMilli(),
	}
}
