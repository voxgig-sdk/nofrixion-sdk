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

func TestNoFrixionMoneyMoovModelsCardCustomerTokenEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.NoFrixionMoneyMoovModelsCardCustomerToken(nil)
		if ent == nil {
			t.Fatal("expected non-nil NoFrixionMoneyMoovModelsCardCustomerTokenEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := no_frixion_money_moov_models_card_customer_tokenBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"list", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "no_frixion_money_moov_models_card_customer_token." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_CARD_CUSTOMER_TOKEN_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		noFrixionMoneyMoovModelsCardCustomerTokenRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.no_frixion_money_moov_models_card_customer_token", setup.data)))
		var noFrixionMoneyMoovModelsCardCustomerTokenRef01Data map[string]any
		if len(noFrixionMoneyMoovModelsCardCustomerTokenRef01DataRaw) > 0 {
			noFrixionMoneyMoovModelsCardCustomerTokenRef01Data = core.ToMapAny(noFrixionMoneyMoovModelsCardCustomerTokenRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = noFrixionMoneyMoovModelsCardCustomerTokenRef01Data

		// LIST
		noFrixionMoneyMoovModelsCardCustomerTokenRef01Ent := client.NoFrixionMoneyMoovModelsCardCustomerToken(nil)
		noFrixionMoneyMoovModelsCardCustomerTokenRef01Match := map[string]any{
			"customer_email_address": setup.idmap["customer_email_address01"],
			"merchant_id": setup.idmap["merchant01"],
		}

		noFrixionMoneyMoovModelsCardCustomerTokenRef01ListResult, err := noFrixionMoneyMoovModelsCardCustomerTokenRef01Ent.List(noFrixionMoneyMoovModelsCardCustomerTokenRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, noFrixionMoneyMoovModelsCardCustomerTokenRef01ListOk := noFrixionMoneyMoovModelsCardCustomerTokenRef01ListResult.([]any)
		if !noFrixionMoneyMoovModelsCardCustomerTokenRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", noFrixionMoneyMoovModelsCardCustomerTokenRef01ListResult)
		}

		// LOAD
		noFrixionMoneyMoovModelsCardCustomerTokenRef01MatchDt0 := map[string]any{
			"id": noFrixionMoneyMoovModelsCardCustomerTokenRef01Data["id"],
		}
		noFrixionMoneyMoovModelsCardCustomerTokenRef01DataDt0Loaded, err := noFrixionMoneyMoovModelsCardCustomerTokenRef01Ent.Load(noFrixionMoneyMoovModelsCardCustomerTokenRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		noFrixionMoneyMoovModelsCardCustomerTokenRef01DataDt0LoadResult := core.ToMapAny(noFrixionMoneyMoovModelsCardCustomerTokenRef01DataDt0Loaded)
		if noFrixionMoneyMoovModelsCardCustomerTokenRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if noFrixionMoneyMoovModelsCardCustomerTokenRef01DataDt0LoadResult["id"] != noFrixionMoneyMoovModelsCardCustomerTokenRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

		// REMOVE
		noFrixionMoneyMoovModelsCardCustomerTokenRef01MatchRm0 := map[string]any{
			"id": noFrixionMoneyMoovModelsCardCustomerTokenRef01Data["id"],
		}
		_, err = noFrixionMoneyMoovModelsCardCustomerTokenRef01Ent.Remove(noFrixionMoneyMoovModelsCardCustomerTokenRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		noFrixionMoneyMoovModelsCardCustomerTokenRef01MatchRt0 := map[string]any{
			"customer_email_address": setup.idmap["customer_email_address01"],
			"merchant_id": setup.idmap["merchant01"],
		}

		noFrixionMoneyMoovModelsCardCustomerTokenRef01ListRt0Result, err := noFrixionMoneyMoovModelsCardCustomerTokenRef01Ent.List(noFrixionMoneyMoovModelsCardCustomerTokenRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, noFrixionMoneyMoovModelsCardCustomerTokenRef01ListRt0Ok := noFrixionMoneyMoovModelsCardCustomerTokenRef01ListRt0Result.([]any)
		if !noFrixionMoneyMoovModelsCardCustomerTokenRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", noFrixionMoneyMoovModelsCardCustomerTokenRef01ListRt0Result)
		}

	})
}

func no_frixion_money_moov_models_card_customer_tokenBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "no_frixion_money_moov_models_card_customer_token", "NoFrixionMoneyMoovModelsCardCustomerTokenTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read no_frixion_money_moov_models_card_customer_token test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse no_frixion_money_moov_models_card_customer_token test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"no_frixion_money_moov_models_card_customer_token01", "no_frixion_money_moov_models_card_customer_token02", "no_frixion_money_moov_models_card_customer_token03", "removeall01", "removeall02", "removeall03", "customertoken01", "customertoken02", "customertoken03", "customer_email_address01", "merchant01"},
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
	entidEnvRaw := os.Getenv("NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_CARD_CUSTOMER_TOKEN_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_CARD_CUSTOMER_TOKEN_ENTID": idmap,
		"NOFRIXION_TEST_LIVE":      "FALSE",
		"NOFRIXION_TEST_EXPLAIN":   "FALSE",
		"NOFRIXION_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_CARD_CUSTOMER_TOKEN_ENTID"])
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
