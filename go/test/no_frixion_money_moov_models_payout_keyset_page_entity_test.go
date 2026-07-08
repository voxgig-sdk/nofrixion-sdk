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

func TestNoFrixionMoneyMoovModelsPayoutKeysetPageEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.NoFrixionMoneyMoovModelsPayoutKeysetPage(nil)
		if ent == nil {
			t.Fatal("expected non-nil NoFrixionMoneyMoovModelsPayoutKeysetPageEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := no_frixion_money_moov_models_payout_keyset_pageBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"list"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "no_frixion_money_moov_models_payout_keyset_page." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_PAYOUT_KEYSET_PAGE_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		noFrixionMoneyMoovModelsPayoutKeysetPageRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.no_frixion_money_moov_models_payout_keyset_page", setup.data)))
		var noFrixionMoneyMoovModelsPayoutKeysetPageRef01Data map[string]any
		if len(noFrixionMoneyMoovModelsPayoutKeysetPageRef01DataRaw) > 0 {
			noFrixionMoneyMoovModelsPayoutKeysetPageRef01Data = core.ToMapAny(noFrixionMoneyMoovModelsPayoutKeysetPageRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = noFrixionMoneyMoovModelsPayoutKeysetPageRef01Data

		// LIST
		noFrixionMoneyMoovModelsPayoutKeysetPageRef01Ent := client.NoFrixionMoneyMoovModelsPayoutKeysetPage(nil)
		noFrixionMoneyMoovModelsPayoutKeysetPageRef01Match := map[string]any{
			"merchant_id": setup.idmap["merchant01"],
		}

		noFrixionMoneyMoovModelsPayoutKeysetPageRef01ListResult, err := noFrixionMoneyMoovModelsPayoutKeysetPageRef01Ent.List(noFrixionMoneyMoovModelsPayoutKeysetPageRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, noFrixionMoneyMoovModelsPayoutKeysetPageRef01ListOk := noFrixionMoneyMoovModelsPayoutKeysetPageRef01ListResult.([]any)
		if !noFrixionMoneyMoovModelsPayoutKeysetPageRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", noFrixionMoneyMoovModelsPayoutKeysetPageRef01ListResult)
		}

	})
}

func no_frixion_money_moov_models_payout_keyset_pageBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "no_frixion_money_moov_models_payout_keyset_page", "NoFrixionMoneyMoovModelsPayoutKeysetPageTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read no_frixion_money_moov_models_payout_keyset_page test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse no_frixion_money_moov_models_payout_keyset_page test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"no_frixion_money_moov_models_payout_keyset_page01", "no_frixion_money_moov_models_payout_keyset_page02", "no_frixion_money_moov_models_payout_keyset_page03", "account01", "account02", "account03", "merchant01", "merchant02", "merchant03", "payout01", "payout02", "payout03"},
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
	entidEnvRaw := os.Getenv("NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_PAYOUT_KEYSET_PAGE_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_PAYOUT_KEYSET_PAGE_ENTID": idmap,
		"NOFRIXION_TEST_LIVE":      "FALSE",
		"NOFRIXION_TEST_EXPLAIN":   "FALSE",
		"NOFRIXION_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_PAYOUT_KEYSET_PAGE_ENTID"])
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
