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

func TestNoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2(nil)
		if ent == nil {
			t.Fatal("expected non-nil NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := no_frixion_money_moov_models_payment_requests_merchant_payment2BasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "no_frixion_money_moov_models_payment_requests_merchant_payment2." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_PAYMENT_REQUESTS_MERCHANT_PAYMENT__ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Ref01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.no_frixion_money_moov_models_payment_requests_merchant_payment2", setup.data)))
		var noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Ref01Data map[string]any
		if len(noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Ref01DataRaw) > 0 {
			noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Ref01Data = core.ToMapAny(noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Ref01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Ref01Data

		// LOAD
		noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Ref01Ent := client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2(nil)
		noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Ref01MatchDt0 := map[string]any{
			"id": noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Ref01Data["id"],
		}
		noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Ref01DataDt0Loaded, err := noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Ref01Ent.Load(noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Ref01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Ref01DataDt0LoadResult := core.ToMapAny(noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Ref01DataDt0Loaded)
		if noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Ref01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Ref01DataDt0LoadResult["id"] != noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Ref01Data["id"] {
			t.Fatal("expected load result id to match")
		}

	})
}

func no_frixion_money_moov_models_payment_requests_merchant_payment2BasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "no_frixion_money_moov_models_payment_requests_merchant_payment2", "NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2TestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read no_frixion_money_moov_models_payment_requests_merchant_payment2 test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse no_frixion_money_moov_models_payment_requests_merchant_payment2 test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"no_frixion_money_moov_models_payment_requests_merchant_payment201", "no_frixion_money_moov_models_payment_requests_merchant_payment202", "no_frixion_money_moov_models_payment_requests_merchant_payment203", "paymentrequest01", "paymentrequest02", "paymentrequest03", "template01", "template02", "template03"},
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
	entidEnvRaw := os.Getenv("NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_PAYMENT_REQUESTS_MERCHANT_PAYMENT__ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_PAYMENT_REQUESTS_MERCHANT_PAYMENT__ENTID": idmap,
		"NOFRIXION_TEST_LIVE":      "FALSE",
		"NOFRIXION_TEST_EXPLAIN":   "FALSE",
		"NOFRIXION_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_PAYMENT_REQUESTS_MERCHANT_PAYMENT__ENTID"])
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
