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

func TestMerchantPaymentRequestTemplateEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.MerchantPaymentRequestTemplate(nil)
		if ent == nil {
			t.Fatal("expected non-nil MerchantPaymentRequestTemplateEntity")
		}
	})

	// Feature #4: the entity Stream(action, ...) method runs the op pipeline and
	// returns a channel over result items. With the streaming feature active it
	// yields the feature's incremental output; otherwise it falls back to the
	// materialised list so Stream always yields.
	t.Run("stream", func(t *testing.T) {
		seed := map[string]any{
			"entity": map[string]any{
				"merchant_payment_request_template": map[string]any{
					"s1": map[string]any{"id": "s1"},
					"s2": map[string]any{"id": "s2"},
					"s3": map[string]any{"id": "s3"},
				},
			},
		}

		// Fallback: streaming inactive -> yields the materialised list items.
		base := sdk.TestSDK(seed, nil)
		var seen []any
		for item := range base.MerchantPaymentRequestTemplate(nil).Stream("list", nil, nil) {
			seen = append(seen, item)
		}
		if len(seen) != 3 {
			t.Fatalf("expected 3 streamed items, got %d", len(seen))
		}

		// Inbound: streaming active -> yields each item from the feature iterator.
		hasStreaming := false
		if fm, ok := core.MakeConfig()["feature"].(map[string]any); ok {
			_, hasStreaming = fm["streaming"]
		}
		if hasStreaming {
			streamSdk := sdk.TestSDK(seed, map[string]any{
				"feature": map[string]any{"streaming": map[string]any{"active": true}},
			})
			var got []any
			for item := range streamSdk.MerchantPaymentRequestTemplate(nil).Stream("list", nil, nil) {
				if sub, ok := item.([]any); ok {
					got = append(got, sub...)
				} else {
					got = append(got, item)
				}
			}
			if len(got) != 3 {
				t.Fatalf("expected 3 items via streaming feature, got %d", len(got))
			}
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := merchant_payment_request_templateBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"list", "update", "load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "merchant_payment_request_template." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_MERCHANT_PAYMENT_REQUEST_TEMPLATE_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		merchantPaymentRequestTemplateRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.merchant_payment_request_template", setup.data)))
		var merchantPaymentRequestTemplateRef01Data map[string]any
		if len(merchantPaymentRequestTemplateRef01DataRaw) > 0 {
			merchantPaymentRequestTemplateRef01Data = core.ToMapAny(merchantPaymentRequestTemplateRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = merchantPaymentRequestTemplateRef01Data

		// LIST
		merchantPaymentRequestTemplateRef01Ent := client.MerchantPaymentRequestTemplate(nil)
		merchantPaymentRequestTemplateRef01Match := map[string]any{
			"merchant_id": setup.idmap["merchant01"],
		}

		merchantPaymentRequestTemplateRef01ListResult, err := merchantPaymentRequestTemplateRef01Ent.List(merchantPaymentRequestTemplateRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, merchantPaymentRequestTemplateRef01ListOk := merchantPaymentRequestTemplateRef01ListResult.([]any)
		if !merchantPaymentRequestTemplateRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", merchantPaymentRequestTemplateRef01ListResult)
		}

		// UPDATE
		merchantPaymentRequestTemplateRef01DataUp0Up := map[string]any{
			"id": merchantPaymentRequestTemplateRef01Data["id"],
			"paymentrequest_id": setup.idmap["paymentrequest_id"],
		}

		merchantPaymentRequestTemplateRef01MarkdefUp0Name := "description"
		merchantPaymentRequestTemplateRef01MarkdefUp0Value := fmt.Sprintf("Mark01-merchant_payment_request_template_ref01_%d", setup.now)
		merchantPaymentRequestTemplateRef01DataUp0Up[merchantPaymentRequestTemplateRef01MarkdefUp0Name] = merchantPaymentRequestTemplateRef01MarkdefUp0Value

		merchantPaymentRequestTemplateRef01ResdataUp0Result, err := merchantPaymentRequestTemplateRef01Ent.Update(merchantPaymentRequestTemplateRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		merchantPaymentRequestTemplateRef01ResdataUp0 := core.ToMapAny(merchantPaymentRequestTemplateRef01ResdataUp0Result)
		if merchantPaymentRequestTemplateRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if merchantPaymentRequestTemplateRef01ResdataUp0["id"] != merchantPaymentRequestTemplateRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if merchantPaymentRequestTemplateRef01ResdataUp0[merchantPaymentRequestTemplateRef01MarkdefUp0Name] != merchantPaymentRequestTemplateRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", merchantPaymentRequestTemplateRef01MarkdefUp0Name, merchantPaymentRequestTemplateRef01ResdataUp0[merchantPaymentRequestTemplateRef01MarkdefUp0Name])
		}

		// LOAD
		merchantPaymentRequestTemplateRef01MatchDt0 := map[string]any{
			"id": merchantPaymentRequestTemplateRef01Data["id"],
		}
		merchantPaymentRequestTemplateRef01DataDt0Loaded, err := merchantPaymentRequestTemplateRef01Ent.Load(merchantPaymentRequestTemplateRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		merchantPaymentRequestTemplateRef01DataDt0LoadResult := core.ToMapAny(merchantPaymentRequestTemplateRef01DataDt0Loaded)
		if merchantPaymentRequestTemplateRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if merchantPaymentRequestTemplateRef01DataDt0LoadResult["id"] != merchantPaymentRequestTemplateRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

	})
}

func merchant_payment_request_templateBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "merchant_payment_request_template", "MerchantPaymentRequestTemplateTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read merchant_payment_request_template test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse merchant_payment_request_template test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"merchant_payment_request_template01", "merchant_payment_request_template02", "merchant_payment_request_template03", "paymentrequest01", "paymentrequest02", "paymentrequest03", "merchant01"},
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
	entidEnvRaw := os.Getenv("NOFRIXION_TEST_MERCHANT_PAYMENT_REQUEST_TEMPLATE_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"NOFRIXION_TEST_MERCHANT_PAYMENT_REQUEST_TEMPLATE_ENTID": idmap,
		"NOFRIXION_TEST_LIVE":      "FALSE",
		"NOFRIXION_TEST_EXPLAIN":   "FALSE",
		"NOFRIXION_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["NOFRIXION_TEST_MERCHANT_PAYMENT_REQUEST_TEMPLATE_ENTID"])
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
