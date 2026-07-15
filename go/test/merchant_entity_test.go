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

func TestMerchantEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.Merchant(nil)
		if ent == nil {
			t.Fatal("expected non-nil MerchantEntity")
		}
	})

	// Feature #4: the entity Stream(action, ...) method runs the op pipeline and
	// returns a channel over result items. With the streaming feature active it
	// yields the feature's incremental output; otherwise it falls back to the
	// materialised list so Stream always yields.
	t.Run("stream", func(t *testing.T) {
		seed := map[string]any{
			"entity": map[string]any{
				"merchant": map[string]any{
					"s1": map[string]any{"id": "s1"},
					"s2": map[string]any{"id": "s2"},
					"s3": map[string]any{"id": "s3"},
				},
			},
		}

		// Fallback: streaming inactive -> yields the materialised list items.
		base := sdk.TestSDK(seed, nil)
		var seen []any
		for item := range base.Merchant(nil).Stream("list", nil, nil) {
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
			for item := range streamSdk.Merchant(nil).Stream("list", nil, nil) {
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
		setup := merchantBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"list", "update", "load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "merchant." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_MERCHANT_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		merchantRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.merchant", setup.data)))
		var merchantRef01Data map[string]any
		if len(merchantRef01DataRaw) > 0 {
			merchantRef01Data = core.ToMapAny(merchantRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = merchantRef01Data

		// LIST
		merchantRef01Ent := client.Merchant(nil)
		merchantRef01Match := map[string]any{}

		merchantRef01ListResult, err := merchantRef01Ent.List(merchantRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, merchantRef01ListOk := merchantRef01ListResult.([]any)
		if !merchantRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", merchantRef01ListResult)
		}

		// UPDATE
		merchantRef01DataUp0Up := map[string]any{
			"id": merchantRef01Data["id"],
		}

		merchantRef01MarkdefUp0Name := "card_payment_processor"
		merchantRef01MarkdefUp0Value := fmt.Sprintf("Mark01-merchant_ref01_%d", setup.now)
		merchantRef01DataUp0Up[merchantRef01MarkdefUp0Name] = merchantRef01MarkdefUp0Value

		merchantRef01ResdataUp0Result, err := merchantRef01Ent.Update(merchantRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		merchantRef01ResdataUp0 := core.ToMapAny(merchantRef01ResdataUp0Result)
		if merchantRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if merchantRef01ResdataUp0["id"] != merchantRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if merchantRef01ResdataUp0[merchantRef01MarkdefUp0Name] != merchantRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", merchantRef01MarkdefUp0Name, merchantRef01ResdataUp0[merchantRef01MarkdefUp0Name])
		}

		// LOAD
		merchantRef01MatchDt0 := map[string]any{
			"id": merchantRef01Data["id"],
		}
		merchantRef01DataDt0Loaded, err := merchantRef01Ent.Load(merchantRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		merchantRef01DataDt0LoadResult := core.ToMapAny(merchantRef01DataDt0Loaded)
		if merchantRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if merchantRef01DataDt0LoadResult["id"] != merchantRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

	})
}

func merchantBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "merchant", "MerchantTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read merchant test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse merchant test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"merchant01", "merchant02", "merchant03", "user01", "user02", "user03", "tag01", "tag02", "tag03"},
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
	entidEnvRaw := os.Getenv("NOFRIXION_TEST_MERCHANT_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"NOFRIXION_TEST_MERCHANT_ENTID": idmap,
		"NOFRIXION_TEST_LIVE":      "FALSE",
		"NOFRIXION_TEST_EXPLAIN":   "FALSE",
		"NOFRIXION_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["NOFRIXION_TEST_MERCHANT_ENTID"])
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
