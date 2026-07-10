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

func TestPayoutEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.Payout(nil)
		if ent == nil {
			t.Fatal("expected non-nil PayoutEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := payoutBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "update", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "payout." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_PAYOUT_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		payoutRef01Ent := client.Payout(nil)
		payoutRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "payout"}, setup.data), "payout_ref01"))
		payoutRef01Data["account_id"] = setup.idmap["account01"]
		payoutRef01Data["destination"] = setup.idmap["destination01"]
		payoutRef01Data["merchant_id"] = setup.idmap["merchant01"]
		payoutRef01Data["source"] = setup.idmap["source01"]

		payoutRef01DataResult, err := payoutRef01Ent.Create(payoutRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		payoutRef01Data = core.ToMapAny(payoutRef01DataResult)
		if payoutRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if payoutRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		payoutRef01Match := map[string]any{
			"merchant_id": setup.idmap["merchant01"],
		}

		payoutRef01ListResult, err := payoutRef01Ent.List(payoutRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		payoutRef01List, payoutRef01ListOk := payoutRef01ListResult.([]any)
		if !payoutRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", payoutRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(payoutRef01List), map[string]any{"id": payoutRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// UPDATE
		payoutRef01DataUp0Up := map[string]any{
			"id": payoutRef01Data["id"],
		}

		payoutRef01MarkdefUp0Name := "account_id"
		payoutRef01MarkdefUp0Value := fmt.Sprintf("Mark01-payout_ref01_%d", setup.now)
		payoutRef01DataUp0Up[payoutRef01MarkdefUp0Name] = payoutRef01MarkdefUp0Value

		payoutRef01ResdataUp0Result, err := payoutRef01Ent.Update(payoutRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		payoutRef01ResdataUp0 := core.ToMapAny(payoutRef01ResdataUp0Result)
		if payoutRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if payoutRef01ResdataUp0["id"] != payoutRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if payoutRef01ResdataUp0[payoutRef01MarkdefUp0Name] != payoutRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", payoutRef01MarkdefUp0Name, payoutRef01ResdataUp0[payoutRef01MarkdefUp0Name])
		}

		// LOAD
		payoutRef01MatchDt0 := map[string]any{
			"id": payoutRef01Data["id"],
		}
		payoutRef01DataDt0Loaded, err := payoutRef01Ent.Load(payoutRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		payoutRef01DataDt0LoadResult := core.ToMapAny(payoutRef01DataDt0Loaded)
		if payoutRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if payoutRef01DataDt0LoadResult["id"] != payoutRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

		// REMOVE
		payoutRef01MatchRm0 := map[string]any{
			"id": payoutRef01Data["id"],
		}
		_, err = payoutRef01Ent.Remove(payoutRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		payoutRef01MatchRt0 := map[string]any{
			"merchant_id": setup.idmap["merchant01"],
		}

		payoutRef01ListRt0Result, err := payoutRef01Ent.List(payoutRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		payoutRef01ListRt0, payoutRef01ListRt0Ok := payoutRef01ListRt0Result.([]any)
		if !payoutRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", payoutRef01ListRt0Result)
		}

		notFoundItem := vs.Select(entityListToData(payoutRef01ListRt0), map[string]any{"id": payoutRef01Data["id"]})
		if !vs.IsEmpty(notFoundItem) {
			t.Fatal("expected removed entity to not be in list")
		}

	})
}

func payoutBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "payout", "PayoutTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read payout test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse payout test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"payout01", "payout02", "payout03", "account01", "account02", "account03", "merchant01", "merchant02", "merchant03", "fxquote01", "fxquote02", "fxquote03", "destination01", "source01"},
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
	entidEnvRaw := os.Getenv("NOFRIXION_TEST_PAYOUT_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"NOFRIXION_TEST_PAYOUT_ENTID": idmap,
		"NOFRIXION_TEST_LIVE":      "FALSE",
		"NOFRIXION_TEST_EXPLAIN":   "FALSE",
		"NOFRIXION_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["NOFRIXION_TEST_PAYOUT_ENTID"])
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
