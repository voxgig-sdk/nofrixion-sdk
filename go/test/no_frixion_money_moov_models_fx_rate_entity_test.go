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

func TestNoFrixionMoneyMoovModelsFxRateEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.NoFrixionMoneyMoovModelsFxRate(nil)
		if ent == nil {
			t.Fatal("expected non-nil NoFrixionMoneyMoovModelsFxRateEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := no_frixion_money_moov_models_fx_rateBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"list", "load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "no_frixion_money_moov_models_fx_rate." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_FX_RATE_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		noFrixionMoneyMoovModelsFxRateRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.no_frixion_money_moov_models_fx_rate", setup.data)))
		var noFrixionMoneyMoovModelsFxRateRef01Data map[string]any
		if len(noFrixionMoneyMoovModelsFxRateRef01DataRaw) > 0 {
			noFrixionMoneyMoovModelsFxRateRef01Data = core.ToMapAny(noFrixionMoneyMoovModelsFxRateRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = noFrixionMoneyMoovModelsFxRateRef01Data

		// LIST
		noFrixionMoneyMoovModelsFxRateRef01Ent := client.NoFrixionMoneyMoovModelsFxRate(nil)
		noFrixionMoneyMoovModelsFxRateRef01Match := map[string]any{
			"destination": setup.idmap["destination01"],
			"source": setup.idmap["source01"],
		}

		noFrixionMoneyMoovModelsFxRateRef01ListResult, err := noFrixionMoneyMoovModelsFxRateRef01Ent.List(noFrixionMoneyMoovModelsFxRateRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, noFrixionMoneyMoovModelsFxRateRef01ListOk := noFrixionMoneyMoovModelsFxRateRef01ListResult.([]any)
		if !noFrixionMoneyMoovModelsFxRateRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", noFrixionMoneyMoovModelsFxRateRef01ListResult)
		}

		// LOAD
		noFrixionMoneyMoovModelsFxRateRef01MatchDt0 := map[string]any{}
		noFrixionMoneyMoovModelsFxRateRef01DataDt0Loaded, err := noFrixionMoneyMoovModelsFxRateRef01Ent.Load(noFrixionMoneyMoovModelsFxRateRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		if noFrixionMoneyMoovModelsFxRateRef01DataDt0Loaded == nil {
			t.Fatal("expected load result to be non-nil")
		}

	})
}

func no_frixion_money_moov_models_fx_rateBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "no_frixion_money_moov_models_fx_rate", "NoFrixionMoneyMoovModelsFxRateTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read no_frixion_money_moov_models_fx_rate test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse no_frixion_money_moov_models_fx_rate test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"no_frixion_money_moov_models_fx_rate01", "no_frixion_money_moov_models_fx_rate02", "no_frixion_money_moov_models_fx_rate03", "fxallheldrate01", "fxallheldrate02", "fxallheldrate03", "fxheldrate01", "fxheldrate02", "fxheldrate03", "destination01", "source01"},
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
	entidEnvRaw := os.Getenv("NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_FX_RATE_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_FX_RATE_ENTID": idmap,
		"NOFRIXION_TEST_LIVE":      "FALSE",
		"NOFRIXION_TEST_EXPLAIN":   "FALSE",
		"NOFRIXION_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_FX_RATE_ENTID"])
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
