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

	t.Run("basic", func(t *testing.T) {
		setup := payrunBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "update", "remove"} {
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

		// UPDATE
		payrunRef01DataUp0Up := map[string]any{
			"id": payrunRef01Data["id"],
		}

		payrunRef01MarkdefUp0Name := "note"
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

		// REMOVE
		payrunRef01MatchRm0 := map[string]any{
			"id": payrunRef01Data["id"],
		}
		_, err = payrunRef01Ent.Remove(payrunRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
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
		[]any{"payrun01", "payrun02", "payrun03"},
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
