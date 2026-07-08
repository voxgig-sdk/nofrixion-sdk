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

func TestNoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment(nil)
		if ent == nil {
			t.Fatal("expected non-nil NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := no_frixion_money_moov_api_features_payment_requests_paymentBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "no_frixion_money_moov_api_features_payment_requests_payment." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_API_FEATURES_PAYMENT_REQUESTS_PAYMENT_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		noFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentRef01Ent := client.NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment(nil)
		noFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "no_frixion_money_moov_api_features_payment_requests_payment"}, setup.data), "no_frixion_money_moov_api_features_payment_requests_payment_ref01"))

		noFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentRef01DataResult, err := noFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentRef01Ent.Create(noFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		noFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentRef01Data = core.ToMapAny(noFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentRef01DataResult)
		if noFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}

	})
}

func no_frixion_money_moov_api_features_payment_requests_paymentBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "no_frixion_money_moov_api_features_payment_requests_payment", "NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read no_frixion_money_moov_api_features_payment_requests_payment test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse no_frixion_money_moov_api_features_payment_requests_payment test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"no_frixion_money_moov_api_features_payment_requests_payment01", "no_frixion_money_moov_api_features_payment_requests_payment02", "no_frixion_money_moov_api_features_payment_requests_payment03"},
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
	entidEnvRaw := os.Getenv("NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_API_FEATURES_PAYMENT_REQUESTS_PAYMENT_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_API_FEATURES_PAYMENT_REQUESTS_PAYMENT_ENTID": idmap,
		"NOFRIXION_TEST_LIVE":      "FALSE",
		"NOFRIXION_TEST_EXPLAIN":   "FALSE",
		"NOFRIXION_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_API_FEATURES_PAYMENT_REQUESTS_PAYMENT_ENTID"])
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
