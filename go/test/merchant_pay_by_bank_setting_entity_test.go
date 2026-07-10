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

func TestMerchantPayByBankSettingEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.MerchantPayByBankSetting(nil)
		if ent == nil {
			t.Fatal("expected non-nil MerchantPayByBankSettingEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := merchant_pay_by_bank_settingBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"list"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "merchant_pay_by_bank_setting." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_MERCHANT_PAY_BY_BANK_SETTING_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		merchantPayByBankSettingRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.merchant_pay_by_bank_setting", setup.data)))
		var merchantPayByBankSettingRef01Data map[string]any
		if len(merchantPayByBankSettingRef01DataRaw) > 0 {
			merchantPayByBankSettingRef01Data = core.ToMapAny(merchantPayByBankSettingRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = merchantPayByBankSettingRef01Data

		// LIST
		merchantPayByBankSettingRef01Ent := client.MerchantPayByBankSetting(nil)
		merchantPayByBankSettingRef01Match := map[string]any{
			"merchant_id": setup.idmap["merchant01"],
		}

		merchantPayByBankSettingRef01ListResult, err := merchantPayByBankSettingRef01Ent.List(merchantPayByBankSettingRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, merchantPayByBankSettingRef01ListOk := merchantPayByBankSettingRef01ListResult.([]any)
		if !merchantPayByBankSettingRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", merchantPayByBankSettingRef01ListResult)
		}

	})
}

func merchant_pay_by_bank_settingBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "merchant_pay_by_bank_setting", "MerchantPayByBankSettingTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read merchant_pay_by_bank_setting test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse merchant_pay_by_bank_setting test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"merchant_pay_by_bank_setting01", "merchant_pay_by_bank_setting02", "merchant_pay_by_bank_setting03", "merchant01", "merchant02", "merchant03"},
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
	entidEnvRaw := os.Getenv("NOFRIXION_TEST_MERCHANT_PAY_BY_BANK_SETTING_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"NOFRIXION_TEST_MERCHANT_PAY_BY_BANK_SETTING_ENTID": idmap,
		"NOFRIXION_TEST_LIVE":      "FALSE",
		"NOFRIXION_TEST_EXPLAIN":   "FALSE",
		"NOFRIXION_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["NOFRIXION_TEST_MERCHANT_PAY_BY_BANK_SETTING_ENTID"])
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
