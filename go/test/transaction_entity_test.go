package sdktest

import (
	"encoding/json"
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

func TestTransactionEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.Transaction(nil)
		if ent == nil {
			t.Fatal("expected non-nil TransactionEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := transactionBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "transaction." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_TRANSACTION_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		transactionRef01Ent := client.Transaction(nil)
		transactionRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "transaction"}, setup.data), "transaction_ref01"))
		transactionRef01Data["account_id"] = setup.idmap["account01"]
		transactionRef01Data["merchant_id"] = setup.idmap["merchant01"]
		transactionRef01Data["transaction_id"] = setup.idmap["transaction01"]

		transactionRef01DataResult, err := transactionRef01Ent.Create(transactionRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		transactionRef01Data = core.ToMapAny(transactionRef01DataResult)
		if transactionRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if transactionRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		transactionRef01Match := map[string]any{}

		transactionRef01ListResult, err := transactionRef01Ent.List(transactionRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		transactionRef01List, transactionRef01ListOk := transactionRef01ListResult.([]any)
		if !transactionRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", transactionRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(transactionRef01List), map[string]any{"id": transactionRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// LOAD
		transactionRef01MatchDt0 := map[string]any{
			"id": transactionRef01Data["id"],
		}
		transactionRef01DataDt0Loaded, err := transactionRef01Ent.Load(transactionRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		transactionRef01DataDt0LoadResult := core.ToMapAny(transactionRef01DataDt0Loaded)
		if transactionRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if transactionRef01DataDt0LoadResult["id"] != transactionRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

		// REMOVE
		transactionRef01MatchRm0 := map[string]any{
			"id": transactionRef01Data["id"],
		}
		_, err = transactionRef01Ent.Remove(transactionRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		transactionRef01MatchRt0 := map[string]any{}

		transactionRef01ListRt0Result, err := transactionRef01Ent.List(transactionRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		transactionRef01ListRt0, transactionRef01ListRt0Ok := transactionRef01ListRt0Result.([]any)
		if !transactionRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", transactionRef01ListRt0Result)
		}

		notFoundItem := vs.Select(entityListToData(transactionRef01ListRt0), map[string]any{"id": transactionRef01Data["id"]})
		if !vs.IsEmpty(notFoundItem) {
			t.Fatal("expected removed entity to not be in list")
		}

	})
}

func transactionBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "transaction", "TransactionTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read transaction test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse transaction test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"transaction01", "transaction02", "transaction03", "account01", "account02", "account03", "merchant01", "merchant02", "merchant03", "from01", "from02", "from03"},
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
	entidEnvRaw := os.Getenv("NOFRIXION_TEST_TRANSACTION_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"NOFRIXION_TEST_TRANSACTION_ENTID": idmap,
		"NOFRIXION_TEST_LIVE":      "FALSE",
		"NOFRIXION_TEST_EXPLAIN":   "FALSE",
		"NOFRIXION_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["NOFRIXION_TEST_TRANSACTION_ENTID"])
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
