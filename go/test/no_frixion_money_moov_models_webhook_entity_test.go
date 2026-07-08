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

func TestNoFrixionMoneyMoovModelsWebhookEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.NoFrixionMoneyMoovModelsWebhook(nil)
		if ent == nil {
			t.Fatal("expected non-nil NoFrixionMoneyMoovModelsWebhookEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := no_frixion_money_moov_models_webhookBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "update", "load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "no_frixion_money_moov_models_webhook." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_WEBHOOK_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		noFrixionMoneyMoovModelsWebhookRef01Ent := client.NoFrixionMoneyMoovModelsWebhook(nil)
		noFrixionMoneyMoovModelsWebhookRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "no_frixion_money_moov_models_webhook"}, setup.data), "no_frixion_money_moov_models_webhook_ref01"))
		noFrixionMoneyMoovModelsWebhookRef01Data["merchant_id"] = setup.idmap["merchant01"]

		noFrixionMoneyMoovModelsWebhookRef01DataResult, err := noFrixionMoneyMoovModelsWebhookRef01Ent.Create(noFrixionMoneyMoovModelsWebhookRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		noFrixionMoneyMoovModelsWebhookRef01Data = core.ToMapAny(noFrixionMoneyMoovModelsWebhookRef01DataResult)
		if noFrixionMoneyMoovModelsWebhookRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if noFrixionMoneyMoovModelsWebhookRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		noFrixionMoneyMoovModelsWebhookRef01Match := map[string]any{
			"merchant_id": setup.idmap["merchant01"],
		}

		noFrixionMoneyMoovModelsWebhookRef01ListResult, err := noFrixionMoneyMoovModelsWebhookRef01Ent.List(noFrixionMoneyMoovModelsWebhookRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		noFrixionMoneyMoovModelsWebhookRef01List, noFrixionMoneyMoovModelsWebhookRef01ListOk := noFrixionMoneyMoovModelsWebhookRef01ListResult.([]any)
		if !noFrixionMoneyMoovModelsWebhookRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", noFrixionMoneyMoovModelsWebhookRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(noFrixionMoneyMoovModelsWebhookRef01List), map[string]any{"id": noFrixionMoneyMoovModelsWebhookRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// UPDATE
		noFrixionMoneyMoovModelsWebhookRef01DataUp0Up := map[string]any{
			"id": noFrixionMoneyMoovModelsWebhookRef01Data["id"],
		}

		noFrixionMoneyMoovModelsWebhookRef01MarkdefUp0Name := "destination_url"
		noFrixionMoneyMoovModelsWebhookRef01MarkdefUp0Value := fmt.Sprintf("Mark01-no_frixion_money_moov_models_webhook_ref01_%d", setup.now)
		noFrixionMoneyMoovModelsWebhookRef01DataUp0Up[noFrixionMoneyMoovModelsWebhookRef01MarkdefUp0Name] = noFrixionMoneyMoovModelsWebhookRef01MarkdefUp0Value

		noFrixionMoneyMoovModelsWebhookRef01ResdataUp0Result, err := noFrixionMoneyMoovModelsWebhookRef01Ent.Update(noFrixionMoneyMoovModelsWebhookRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		noFrixionMoneyMoovModelsWebhookRef01ResdataUp0 := core.ToMapAny(noFrixionMoneyMoovModelsWebhookRef01ResdataUp0Result)
		if noFrixionMoneyMoovModelsWebhookRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if noFrixionMoneyMoovModelsWebhookRef01ResdataUp0["id"] != noFrixionMoneyMoovModelsWebhookRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if noFrixionMoneyMoovModelsWebhookRef01ResdataUp0[noFrixionMoneyMoovModelsWebhookRef01MarkdefUp0Name] != noFrixionMoneyMoovModelsWebhookRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", noFrixionMoneyMoovModelsWebhookRef01MarkdefUp0Name, noFrixionMoneyMoovModelsWebhookRef01ResdataUp0[noFrixionMoneyMoovModelsWebhookRef01MarkdefUp0Name])
		}

		// LOAD
		noFrixionMoneyMoovModelsWebhookRef01MatchDt0 := map[string]any{
			"id": noFrixionMoneyMoovModelsWebhookRef01Data["id"],
		}
		noFrixionMoneyMoovModelsWebhookRef01DataDt0Loaded, err := noFrixionMoneyMoovModelsWebhookRef01Ent.Load(noFrixionMoneyMoovModelsWebhookRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		noFrixionMoneyMoovModelsWebhookRef01DataDt0LoadResult := core.ToMapAny(noFrixionMoneyMoovModelsWebhookRef01DataDt0Loaded)
		if noFrixionMoneyMoovModelsWebhookRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if noFrixionMoneyMoovModelsWebhookRef01DataDt0LoadResult["id"] != noFrixionMoneyMoovModelsWebhookRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

	})
}

func no_frixion_money_moov_models_webhookBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "no_frixion_money_moov_models_webhook", "NoFrixionMoneyMoovModelsWebhookTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read no_frixion_money_moov_models_webhook test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse no_frixion_money_moov_models_webhook test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"no_frixion_money_moov_models_webhook01", "no_frixion_money_moov_models_webhook02", "no_frixion_money_moov_models_webhook03", "merchant01", "merchant02", "merchant03"},
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
	entidEnvRaw := os.Getenv("NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_WEBHOOK_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_WEBHOOK_ENTID": idmap,
		"NOFRIXION_TEST_LIVE":      "FALSE",
		"NOFRIXION_TEST_EXPLAIN":   "FALSE",
		"NOFRIXION_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_WEBHOOK_ENTID"])
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
