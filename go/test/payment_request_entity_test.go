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

func TestPaymentRequestEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.PaymentRequest(nil)
		if ent == nil {
			t.Fatal("expected non-nil PaymentRequestEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := payment_requestBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "update", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "payment_request." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_PAYMENT_REQUEST_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		paymentRequestRef01Ent := client.PaymentRequest(nil)
		paymentRequestRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "payment_request"}, setup.data), "payment_request_ref01"))

		paymentRequestRef01DataResult, err := paymentRequestRef01Ent.Create(paymentRequestRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		paymentRequestRef01Data = core.ToMapAny(paymentRequestRef01DataResult)
		if paymentRequestRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if paymentRequestRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		paymentRequestRef01Match := map[string]any{}

		paymentRequestRef01ListResult, err := paymentRequestRef01Ent.List(paymentRequestRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		paymentRequestRef01List, paymentRequestRef01ListOk := paymentRequestRef01ListResult.([]any)
		if !paymentRequestRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", paymentRequestRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(paymentRequestRef01List), map[string]any{"id": paymentRequestRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// UPDATE
		paymentRequestRef01DataUp0Up := map[string]any{
			"id": paymentRequestRef01Data["id"],
		}

		paymentRequestRef01MarkdefUp0Name := "base_origin_url"
		paymentRequestRef01MarkdefUp0Value := fmt.Sprintf("Mark01-payment_request_ref01_%d", setup.now)
		paymentRequestRef01DataUp0Up[paymentRequestRef01MarkdefUp0Name] = paymentRequestRef01MarkdefUp0Value

		paymentRequestRef01ResdataUp0Result, err := paymentRequestRef01Ent.Update(paymentRequestRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		paymentRequestRef01ResdataUp0 := core.ToMapAny(paymentRequestRef01ResdataUp0Result)
		if paymentRequestRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if paymentRequestRef01ResdataUp0["id"] != paymentRequestRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if paymentRequestRef01ResdataUp0[paymentRequestRef01MarkdefUp0Name] != paymentRequestRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", paymentRequestRef01MarkdefUp0Name, paymentRequestRef01ResdataUp0[paymentRequestRef01MarkdefUp0Name])
		}

		// LOAD
		paymentRequestRef01MatchDt0 := map[string]any{
			"id": paymentRequestRef01Data["id"],
		}
		paymentRequestRef01DataDt0Loaded, err := paymentRequestRef01Ent.Load(paymentRequestRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		paymentRequestRef01DataDt0LoadResult := core.ToMapAny(paymentRequestRef01DataDt0Loaded)
		if paymentRequestRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if paymentRequestRef01DataDt0LoadResult["id"] != paymentRequestRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

		// REMOVE
		paymentRequestRef01MatchRm0 := map[string]any{
			"id": paymentRequestRef01Data["id"],
		}
		_, err = paymentRequestRef01Ent.Remove(paymentRequestRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		paymentRequestRef01MatchRt0 := map[string]any{}

		paymentRequestRef01ListRt0Result, err := paymentRequestRef01Ent.List(paymentRequestRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		paymentRequestRef01ListRt0, paymentRequestRef01ListRt0Ok := paymentRequestRef01ListRt0Result.([]any)
		if !paymentRequestRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", paymentRequestRef01ListRt0Result)
		}

		notFoundItem := vs.Select(entityListToData(paymentRequestRef01ListRt0), map[string]any{"id": paymentRequestRef01Data["id"]})
		if !vs.IsEmpty(notFoundItem) {
			t.Fatal("expected removed entity to not be in list")
		}

	})
}

func payment_requestBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "payment_request", "PaymentRequestTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read payment_request test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse payment_request test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"payment_request01", "payment_request02", "payment_request03", "paymentrequest01", "paymentrequest02", "paymentrequest03"},
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
	entidEnvRaw := os.Getenv("NOFRIXION_TEST_PAYMENT_REQUEST_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"NOFRIXION_TEST_PAYMENT_REQUEST_ENTID": idmap,
		"NOFRIXION_TEST_LIVE":      "FALSE",
		"NOFRIXION_TEST_EXPLAIN":   "FALSE",
		"NOFRIXION_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["NOFRIXION_TEST_PAYMENT_REQUEST_ENTID"])
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
