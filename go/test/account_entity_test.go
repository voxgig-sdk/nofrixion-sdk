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

func TestAccountEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.Account(nil)
		if ent == nil {
			t.Fatal("expected non-nil AccountEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := accountBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "update", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "account." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_ACCOUNT_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		accountRef01Ent := client.Account(nil)
		accountRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "account"}, setup.data), "account_ref01"))
		accountRef01Data["account_id"] = setup.idmap["account01"]
		accountRef01Data["merchant_id"] = setup.idmap["merchant01"]

		accountRef01DataResult, err := accountRef01Ent.Create(accountRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		accountRef01Data = core.ToMapAny(accountRef01DataResult)
		if accountRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if accountRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		accountRef01Match := map[string]any{
			"merchant_id": setup.idmap["merchant01"],
		}

		accountRef01ListResult, err := accountRef01Ent.List(accountRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		accountRef01List, accountRef01ListOk := accountRef01ListResult.([]any)
		if !accountRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", accountRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(accountRef01List), map[string]any{"id": accountRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// UPDATE
		accountRef01DataUp0Up := map[string]any{
			"id": accountRef01Data["id"],
		}

		accountRef01MarkdefUp0Name := "account_id"
		accountRef01MarkdefUp0Value := fmt.Sprintf("Mark01-account_ref01_%d", setup.now)
		accountRef01DataUp0Up[accountRef01MarkdefUp0Name] = accountRef01MarkdefUp0Value

		accountRef01ResdataUp0Result, err := accountRef01Ent.Update(accountRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		accountRef01ResdataUp0 := core.ToMapAny(accountRef01ResdataUp0Result)
		if accountRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if accountRef01ResdataUp0["id"] != accountRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if accountRef01ResdataUp0[accountRef01MarkdefUp0Name] != accountRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", accountRef01MarkdefUp0Name, accountRef01ResdataUp0[accountRef01MarkdefUp0Name])
		}

		// LOAD
		accountRef01MatchDt0 := map[string]any{
			"id": accountRef01Data["id"],
		}
		accountRef01DataDt0Loaded, err := accountRef01Ent.Load(accountRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		accountRef01DataDt0LoadResult := core.ToMapAny(accountRef01DataDt0Loaded)
		if accountRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if accountRef01DataDt0LoadResult["id"] != accountRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

		// REMOVE
		accountRef01MatchRm0 := map[string]any{
			"id": accountRef01Data["id"],
		}
		_, err = accountRef01Ent.Remove(accountRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		accountRef01MatchRt0 := map[string]any{
			"merchant_id": setup.idmap["merchant01"],
		}

		accountRef01ListRt0Result, err := accountRef01Ent.List(accountRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		accountRef01ListRt0, accountRef01ListRt0Ok := accountRef01ListRt0Result.([]any)
		if !accountRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", accountRef01ListRt0Result)
		}

		notFoundItem := vs.Select(entityListToData(accountRef01ListRt0), map[string]any{"id": accountRef01Data["id"]})
		if !vs.IsEmpty(notFoundItem) {
			t.Fatal("expected removed entity to not be in list")
		}

	})
}

func accountBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "account", "AccountTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read account test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse account test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"account01", "account02", "account03", "merchant01", "merchant02", "merchant03", "topup01", "topup02", "topup03"},
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
	entidEnvRaw := os.Getenv("NOFRIXION_TEST_ACCOUNT_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"NOFRIXION_TEST_ACCOUNT_ENTID": idmap,
		"NOFRIXION_TEST_LIVE":      "FALSE",
		"NOFRIXION_TEST_EXPLAIN":   "FALSE",
		"NOFRIXION_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["NOFRIXION_TEST_ACCOUNT_ENTID"])
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
