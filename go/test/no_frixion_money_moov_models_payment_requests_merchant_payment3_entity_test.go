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

func TestNoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3(nil)
		if ent == nil {
			t.Fatal("expected non-nil NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := no_frixion_money_moov_models_payment_requests_merchant_payment3BasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"update"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "no_frixion_money_moov_models_payment_requests_merchant_payment3." + _op, _mode); _shouldSkip {
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
		noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Ref01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.no_frixion_money_moov_models_payment_requests_merchant_payment3", setup.data)))
		var noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Ref01Data map[string]any
		if len(noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Ref01DataRaw) > 0 {
			noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Ref01Data = core.ToMapAny(noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Ref01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Ref01Data

		// UPDATE
		noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Ref01Ent := client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3(nil)
		noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Ref01DataUp0Up := map[string]any{
			"id": noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Ref01Data["id"],
			"paymentrequest_id": setup.idmap["paymentrequest_id"],
		}

		noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Ref01MarkdefUp0Name := "description"
		noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Ref01MarkdefUp0Value := fmt.Sprintf("Mark01-no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_%d", setup.now)
		noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Ref01DataUp0Up[noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Ref01MarkdefUp0Name] = noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Ref01MarkdefUp0Value

		noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Ref01ResdataUp0Result, err := noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Ref01Ent.Update(noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Ref01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Ref01ResdataUp0 := core.ToMapAny(noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Ref01ResdataUp0Result)
		if noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Ref01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Ref01ResdataUp0["id"] != noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Ref01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Ref01ResdataUp0[noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Ref01MarkdefUp0Name] != noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Ref01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Ref01MarkdefUp0Name, noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Ref01ResdataUp0[noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Ref01MarkdefUp0Name])
		}

	})
}

func no_frixion_money_moov_models_payment_requests_merchant_payment3BasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "no_frixion_money_moov_models_payment_requests_merchant_payment3", "NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3TestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read no_frixion_money_moov_models_payment_requests_merchant_payment3 test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse no_frixion_money_moov_models_payment_requests_merchant_payment3 test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"no_frixion_money_moov_models_payment_requests_merchant_payment301", "no_frixion_money_moov_models_payment_requests_merchant_payment302", "no_frixion_money_moov_models_payment_requests_merchant_payment303", "paymentrequest01", "paymentrequest02", "paymentrequest03", "template01", "template02", "template03"},
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
	// Add paymentrequest_id alias for update test.
	if idmapResolved["paymentrequest_id"] == nil {
		idmapResolved["paymentrequest_id"] = idmapResolved["paymentrequest01"]
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
