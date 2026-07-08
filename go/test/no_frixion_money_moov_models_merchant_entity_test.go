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

func TestNoFrixionMoneyMoovModelsMerchantEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.NoFrixionMoneyMoovModelsMerchant(nil)
		if ent == nil {
			t.Fatal("expected non-nil NoFrixionMoneyMoovModelsMerchantEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := no_frixion_money_moov_models_merchantBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"list", "update", "load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "no_frixion_money_moov_models_merchant." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_MERCHANT_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		noFrixionMoneyMoovModelsMerchantRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.no_frixion_money_moov_models_merchant", setup.data)))
		var noFrixionMoneyMoovModelsMerchantRef01Data map[string]any
		if len(noFrixionMoneyMoovModelsMerchantRef01DataRaw) > 0 {
			noFrixionMoneyMoovModelsMerchantRef01Data = core.ToMapAny(noFrixionMoneyMoovModelsMerchantRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = noFrixionMoneyMoovModelsMerchantRef01Data

		// LIST
		noFrixionMoneyMoovModelsMerchantRef01Ent := client.NoFrixionMoneyMoovModelsMerchant(nil)
		noFrixionMoneyMoovModelsMerchantRef01Match := map[string]any{}

		noFrixionMoneyMoovModelsMerchantRef01ListResult, err := noFrixionMoneyMoovModelsMerchantRef01Ent.List(noFrixionMoneyMoovModelsMerchantRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, noFrixionMoneyMoovModelsMerchantRef01ListOk := noFrixionMoneyMoovModelsMerchantRef01ListResult.([]any)
		if !noFrixionMoneyMoovModelsMerchantRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", noFrixionMoneyMoovModelsMerchantRef01ListResult)
		}

		// UPDATE
		noFrixionMoneyMoovModelsMerchantRef01DataUp0Up := map[string]any{
			"id": noFrixionMoneyMoovModelsMerchantRef01Data["id"],
		}

		noFrixionMoneyMoovModelsMerchantRef01MarkdefUp0Name := "card_payment_processor"
		noFrixionMoneyMoovModelsMerchantRef01MarkdefUp0Value := fmt.Sprintf("Mark01-no_frixion_money_moov_models_merchant_ref01_%d", setup.now)
		noFrixionMoneyMoovModelsMerchantRef01DataUp0Up[noFrixionMoneyMoovModelsMerchantRef01MarkdefUp0Name] = noFrixionMoneyMoovModelsMerchantRef01MarkdefUp0Value

		noFrixionMoneyMoovModelsMerchantRef01ResdataUp0Result, err := noFrixionMoneyMoovModelsMerchantRef01Ent.Update(noFrixionMoneyMoovModelsMerchantRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		noFrixionMoneyMoovModelsMerchantRef01ResdataUp0 := core.ToMapAny(noFrixionMoneyMoovModelsMerchantRef01ResdataUp0Result)
		if noFrixionMoneyMoovModelsMerchantRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if noFrixionMoneyMoovModelsMerchantRef01ResdataUp0["id"] != noFrixionMoneyMoovModelsMerchantRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if noFrixionMoneyMoovModelsMerchantRef01ResdataUp0[noFrixionMoneyMoovModelsMerchantRef01MarkdefUp0Name] != noFrixionMoneyMoovModelsMerchantRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", noFrixionMoneyMoovModelsMerchantRef01MarkdefUp0Name, noFrixionMoneyMoovModelsMerchantRef01ResdataUp0[noFrixionMoneyMoovModelsMerchantRef01MarkdefUp0Name])
		}

		// LOAD
		noFrixionMoneyMoovModelsMerchantRef01MatchDt0 := map[string]any{
			"id": noFrixionMoneyMoovModelsMerchantRef01Data["id"],
		}
		noFrixionMoneyMoovModelsMerchantRef01DataDt0Loaded, err := noFrixionMoneyMoovModelsMerchantRef01Ent.Load(noFrixionMoneyMoovModelsMerchantRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		noFrixionMoneyMoovModelsMerchantRef01DataDt0LoadResult := core.ToMapAny(noFrixionMoneyMoovModelsMerchantRef01DataDt0Loaded)
		if noFrixionMoneyMoovModelsMerchantRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if noFrixionMoneyMoovModelsMerchantRef01DataDt0LoadResult["id"] != noFrixionMoneyMoovModelsMerchantRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

	})
}

func no_frixion_money_moov_models_merchantBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "no_frixion_money_moov_models_merchant", "NoFrixionMoneyMoovModelsMerchantTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read no_frixion_money_moov_models_merchant test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse no_frixion_money_moov_models_merchant test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"no_frixion_money_moov_models_merchant01", "no_frixion_money_moov_models_merchant02", "no_frixion_money_moov_models_merchant03"},
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
	entidEnvRaw := os.Getenv("NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_MERCHANT_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_MERCHANT_ENTID": idmap,
		"NOFRIXION_TEST_LIVE":      "FALSE",
		"NOFRIXION_TEST_EXPLAIN":   "FALSE",
		"NOFRIXION_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_MERCHANT_ENTID"])
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
