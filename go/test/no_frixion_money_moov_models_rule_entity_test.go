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

func TestNoFrixionMoneyMoovModelsRuleEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.NoFrixionMoneyMoovModelsRule(nil)
		if ent == nil {
			t.Fatal("expected non-nil NoFrixionMoneyMoovModelsRuleEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := no_frixion_money_moov_models_ruleBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "update", "load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "no_frixion_money_moov_models_rule." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_RULE_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		noFrixionMoneyMoovModelsRuleRef01Ent := client.NoFrixionMoneyMoovModelsRule(nil)
		noFrixionMoneyMoovModelsRuleRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "no_frixion_money_moov_models_rule"}, setup.data), "no_frixion_money_moov_models_rule_ref01"))

		noFrixionMoneyMoovModelsRuleRef01DataResult, err := noFrixionMoneyMoovModelsRuleRef01Ent.Create(noFrixionMoneyMoovModelsRuleRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		noFrixionMoneyMoovModelsRuleRef01Data = core.ToMapAny(noFrixionMoneyMoovModelsRuleRef01DataResult)
		if noFrixionMoneyMoovModelsRuleRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if noFrixionMoneyMoovModelsRuleRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// UPDATE
		noFrixionMoneyMoovModelsRuleRef01DataUp0Up := map[string]any{
			"id": noFrixionMoneyMoovModelsRuleRef01Data["id"],
		}

		noFrixionMoneyMoovModelsRuleRef01MarkdefUp0Name := "account_id"
		noFrixionMoneyMoovModelsRuleRef01MarkdefUp0Value := fmt.Sprintf("Mark01-no_frixion_money_moov_models_rule_ref01_%d", setup.now)
		noFrixionMoneyMoovModelsRuleRef01DataUp0Up[noFrixionMoneyMoovModelsRuleRef01MarkdefUp0Name] = noFrixionMoneyMoovModelsRuleRef01MarkdefUp0Value

		noFrixionMoneyMoovModelsRuleRef01ResdataUp0Result, err := noFrixionMoneyMoovModelsRuleRef01Ent.Update(noFrixionMoneyMoovModelsRuleRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		noFrixionMoneyMoovModelsRuleRef01ResdataUp0 := core.ToMapAny(noFrixionMoneyMoovModelsRuleRef01ResdataUp0Result)
		if noFrixionMoneyMoovModelsRuleRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if noFrixionMoneyMoovModelsRuleRef01ResdataUp0["id"] != noFrixionMoneyMoovModelsRuleRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if noFrixionMoneyMoovModelsRuleRef01ResdataUp0[noFrixionMoneyMoovModelsRuleRef01MarkdefUp0Name] != noFrixionMoneyMoovModelsRuleRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", noFrixionMoneyMoovModelsRuleRef01MarkdefUp0Name, noFrixionMoneyMoovModelsRuleRef01ResdataUp0[noFrixionMoneyMoovModelsRuleRef01MarkdefUp0Name])
		}

		// LOAD
		noFrixionMoneyMoovModelsRuleRef01MatchDt0 := map[string]any{
			"id": noFrixionMoneyMoovModelsRuleRef01Data["id"],
		}
		noFrixionMoneyMoovModelsRuleRef01DataDt0Loaded, err := noFrixionMoneyMoovModelsRuleRef01Ent.Load(noFrixionMoneyMoovModelsRuleRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		noFrixionMoneyMoovModelsRuleRef01DataDt0LoadResult := core.ToMapAny(noFrixionMoneyMoovModelsRuleRef01DataDt0Loaded)
		if noFrixionMoneyMoovModelsRuleRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if noFrixionMoneyMoovModelsRuleRef01DataDt0LoadResult["id"] != noFrixionMoneyMoovModelsRuleRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

	})
}

func no_frixion_money_moov_models_ruleBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "no_frixion_money_moov_models_rule", "NoFrixionMoneyMoovModelsRuleTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read no_frixion_money_moov_models_rule test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse no_frixion_money_moov_models_rule test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"no_frixion_money_moov_models_rule01", "no_frixion_money_moov_models_rule02", "no_frixion_money_moov_models_rule03"},
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
	entidEnvRaw := os.Getenv("NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_RULE_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_RULE_ENTID": idmap,
		"NOFRIXION_TEST_LIVE":      "FALSE",
		"NOFRIXION_TEST_EXPLAIN":   "FALSE",
		"NOFRIXION_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_RULE_ENTID"])
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
