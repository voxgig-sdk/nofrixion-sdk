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

func TestPaymentEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.Payment(nil)
		if ent == nil {
			t.Fatal("expected non-nil PaymentEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := paymentBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "update", "load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "payment." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_PAYMENT_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		paymentRef01Ent := client.Payment(nil)
		paymentRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "payment"}, setup.data), "payment_ref01"))

		paymentRef01DataResult, err := paymentRef01Ent.Create(paymentRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		paymentRef01Data = core.ToMapAny(paymentRef01DataResult)
		if paymentRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if paymentRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// UPDATE
		paymentRef01DataUp0Up := map[string]any{
			"id": paymentRef01Data["id"],
		}

		paymentRef01MarkdefUp0Name := "base_origin_url"
		paymentRef01MarkdefUp0Value := fmt.Sprintf("Mark01-payment_ref01_%d", setup.now)
		paymentRef01DataUp0Up[paymentRef01MarkdefUp0Name] = paymentRef01MarkdefUp0Value

		paymentRef01ResdataUp0Result, err := paymentRef01Ent.Update(paymentRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		paymentRef01ResdataUp0 := core.ToMapAny(paymentRef01ResdataUp0Result)
		if paymentRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if paymentRef01ResdataUp0["id"] != paymentRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if paymentRef01ResdataUp0[paymentRef01MarkdefUp0Name] != paymentRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", paymentRef01MarkdefUp0Name, paymentRef01ResdataUp0[paymentRef01MarkdefUp0Name])
		}

		// LOAD
		paymentRef01MatchDt0 := map[string]any{
			"id": paymentRef01Data["id"],
		}
		paymentRef01DataDt0Loaded, err := paymentRef01Ent.Load(paymentRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		paymentRef01DataDt0LoadResult := core.ToMapAny(paymentRef01DataDt0Loaded)
		if paymentRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if paymentRef01DataDt0LoadResult["id"] != paymentRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

	})
}

func paymentBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "payment", "PaymentTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read payment test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse payment test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"payment01", "payment02", "payment03", "getbyorderid01", "getbyorderid02", "getbyorderid03"},
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
	entidEnvRaw := os.Getenv("NOFRIXION_TEST_PAYMENT_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"NOFRIXION_TEST_PAYMENT_ENTID": idmap,
		"NOFRIXION_TEST_LIVE":      "FALSE",
		"NOFRIXION_TEST_EXPLAIN":   "FALSE",
		"NOFRIXION_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["NOFRIXION_TEST_PAYMENT_ENTID"])
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
