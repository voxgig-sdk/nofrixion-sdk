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

func TestNoFrixionMoneyMoovModelsOpenBankingConsentEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.NoFrixionMoneyMoovModelsOpenBankingConsent(nil)
		if ent == nil {
			t.Fatal("expected non-nil NoFrixionMoneyMoovModelsOpenBankingConsentEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := no_frixion_money_moov_models_open_banking_consentBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "update", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "no_frixion_money_moov_models_open_banking_consent." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_OPEN_BANKING_CONSENT_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		noFrixionMoneyMoovModelsOpenBankingConsentRef01Ent := client.NoFrixionMoneyMoovModelsOpenBankingConsent(nil)
		noFrixionMoneyMoovModelsOpenBankingConsentRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "no_frixion_money_moov_models_open_banking_consent"}, setup.data), "no_frixion_money_moov_models_open_banking_consent_ref01"))
		noFrixionMoneyMoovModelsOpenBankingConsentRef01Data["email"] = setup.idmap["email01"]
		noFrixionMoneyMoovModelsOpenBankingConsentRef01Data["merchant_id"] = setup.idmap["merchant01"]

		noFrixionMoneyMoovModelsOpenBankingConsentRef01DataResult, err := noFrixionMoneyMoovModelsOpenBankingConsentRef01Ent.Create(noFrixionMoneyMoovModelsOpenBankingConsentRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		noFrixionMoneyMoovModelsOpenBankingConsentRef01Data = core.ToMapAny(noFrixionMoneyMoovModelsOpenBankingConsentRef01DataResult)
		if noFrixionMoneyMoovModelsOpenBankingConsentRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if noFrixionMoneyMoovModelsOpenBankingConsentRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		noFrixionMoneyMoovModelsOpenBankingConsentRef01Match := map[string]any{
			"email": setup.idmap["email01"],
			"merchant_id": setup.idmap["merchant01"],
		}

		noFrixionMoneyMoovModelsOpenBankingConsentRef01ListResult, err := noFrixionMoneyMoovModelsOpenBankingConsentRef01Ent.List(noFrixionMoneyMoovModelsOpenBankingConsentRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		noFrixionMoneyMoovModelsOpenBankingConsentRef01List, noFrixionMoneyMoovModelsOpenBankingConsentRef01ListOk := noFrixionMoneyMoovModelsOpenBankingConsentRef01ListResult.([]any)
		if !noFrixionMoneyMoovModelsOpenBankingConsentRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", noFrixionMoneyMoovModelsOpenBankingConsentRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(noFrixionMoneyMoovModelsOpenBankingConsentRef01List), map[string]any{"id": noFrixionMoneyMoovModelsOpenBankingConsentRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// UPDATE
		noFrixionMoneyMoovModelsOpenBankingConsentRef01DataUp0Up := map[string]any{
			"id": noFrixionMoneyMoovModelsOpenBankingConsentRef01Data["id"],
		}

		noFrixionMoneyMoovModelsOpenBankingConsentRef01MarkdefUp0Name := "authorisation_url"
		noFrixionMoneyMoovModelsOpenBankingConsentRef01MarkdefUp0Value := fmt.Sprintf("Mark01-no_frixion_money_moov_models_open_banking_consent_ref01_%d", setup.now)
		noFrixionMoneyMoovModelsOpenBankingConsentRef01DataUp0Up[noFrixionMoneyMoovModelsOpenBankingConsentRef01MarkdefUp0Name] = noFrixionMoneyMoovModelsOpenBankingConsentRef01MarkdefUp0Value

		noFrixionMoneyMoovModelsOpenBankingConsentRef01ResdataUp0Result, err := noFrixionMoneyMoovModelsOpenBankingConsentRef01Ent.Update(noFrixionMoneyMoovModelsOpenBankingConsentRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		noFrixionMoneyMoovModelsOpenBankingConsentRef01ResdataUp0 := core.ToMapAny(noFrixionMoneyMoovModelsOpenBankingConsentRef01ResdataUp0Result)
		if noFrixionMoneyMoovModelsOpenBankingConsentRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if noFrixionMoneyMoovModelsOpenBankingConsentRef01ResdataUp0["id"] != noFrixionMoneyMoovModelsOpenBankingConsentRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if noFrixionMoneyMoovModelsOpenBankingConsentRef01ResdataUp0[noFrixionMoneyMoovModelsOpenBankingConsentRef01MarkdefUp0Name] != noFrixionMoneyMoovModelsOpenBankingConsentRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", noFrixionMoneyMoovModelsOpenBankingConsentRef01MarkdefUp0Name, noFrixionMoneyMoovModelsOpenBankingConsentRef01ResdataUp0[noFrixionMoneyMoovModelsOpenBankingConsentRef01MarkdefUp0Name])
		}

		// LOAD
		noFrixionMoneyMoovModelsOpenBankingConsentRef01MatchDt0 := map[string]any{
			"id": noFrixionMoneyMoovModelsOpenBankingConsentRef01Data["id"],
		}
		noFrixionMoneyMoovModelsOpenBankingConsentRef01DataDt0Loaded, err := noFrixionMoneyMoovModelsOpenBankingConsentRef01Ent.Load(noFrixionMoneyMoovModelsOpenBankingConsentRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		noFrixionMoneyMoovModelsOpenBankingConsentRef01DataDt0LoadResult := core.ToMapAny(noFrixionMoneyMoovModelsOpenBankingConsentRef01DataDt0Loaded)
		if noFrixionMoneyMoovModelsOpenBankingConsentRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if noFrixionMoneyMoovModelsOpenBankingConsentRef01DataDt0LoadResult["id"] != noFrixionMoneyMoovModelsOpenBankingConsentRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

		// REMOVE
		noFrixionMoneyMoovModelsOpenBankingConsentRef01MatchRm0 := map[string]any{
			"id": noFrixionMoneyMoovModelsOpenBankingConsentRef01Data["id"],
		}
		_, err = noFrixionMoneyMoovModelsOpenBankingConsentRef01Ent.Remove(noFrixionMoneyMoovModelsOpenBankingConsentRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		noFrixionMoneyMoovModelsOpenBankingConsentRef01MatchRt0 := map[string]any{
			"email": setup.idmap["email01"],
			"merchant_id": setup.idmap["merchant01"],
		}

		noFrixionMoneyMoovModelsOpenBankingConsentRef01ListRt0Result, err := noFrixionMoneyMoovModelsOpenBankingConsentRef01Ent.List(noFrixionMoneyMoovModelsOpenBankingConsentRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		noFrixionMoneyMoovModelsOpenBankingConsentRef01ListRt0, noFrixionMoneyMoovModelsOpenBankingConsentRef01ListRt0Ok := noFrixionMoneyMoovModelsOpenBankingConsentRef01ListRt0Result.([]any)
		if !noFrixionMoneyMoovModelsOpenBankingConsentRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", noFrixionMoneyMoovModelsOpenBankingConsentRef01ListRt0Result)
		}

		notFoundItem := vs.Select(entityListToData(noFrixionMoneyMoovModelsOpenBankingConsentRef01ListRt0), map[string]any{"id": noFrixionMoneyMoovModelsOpenBankingConsentRef01Data["id"]})
		if !vs.IsEmpty(notFoundItem) {
			t.Fatal("expected removed entity to not be in list")
		}

	})
}

func no_frixion_money_moov_models_open_banking_consentBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "no_frixion_money_moov_models_open_banking_consent", "NoFrixionMoneyMoovModelsOpenBankingConsentTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read no_frixion_money_moov_models_open_banking_consent test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse no_frixion_money_moov_models_open_banking_consent test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"no_frixion_money_moov_models_open_banking_consent01", "no_frixion_money_moov_models_open_banking_consent02", "no_frixion_money_moov_models_open_banking_consent03", "consent01", "consent02", "consent03", "email01", "merchant01"},
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
	entidEnvRaw := os.Getenv("NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_OPEN_BANKING_CONSENT_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_OPEN_BANKING_CONSENT_ENTID": idmap,
		"NOFRIXION_TEST_LIVE":      "FALSE",
		"NOFRIXION_TEST_EXPLAIN":   "FALSE",
		"NOFRIXION_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_OPEN_BANKING_CONSENT_ENTID"])
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
