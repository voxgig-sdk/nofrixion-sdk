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

func TestMerchantTokenEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.MerchantToken(nil)
		if ent == nil {
			t.Fatal("expected non-nil MerchantTokenEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := merchant_tokenBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "update", "load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "merchant_token." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_MERCHANT_TOKEN_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		merchantTokenRef01Ent := client.MerchantToken(nil)
		merchantTokenRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "merchant_token"}, setup.data), "merchant_token_ref01"))
		merchantTokenRef01Data["merchant_id"] = setup.idmap["merchant01"]

		merchantTokenRef01DataResult, err := merchantTokenRef01Ent.Create(merchantTokenRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		merchantTokenRef01Data = core.ToMapAny(merchantTokenRef01DataResult)
		if merchantTokenRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if merchantTokenRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		merchantTokenRef01Match := map[string]any{
			"merchant_id": setup.idmap["merchant01"],
		}

		merchantTokenRef01ListResult, err := merchantTokenRef01Ent.List(merchantTokenRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		merchantTokenRef01List, merchantTokenRef01ListOk := merchantTokenRef01ListResult.([]any)
		if !merchantTokenRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", merchantTokenRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(merchantTokenRef01List), map[string]any{"id": merchantTokenRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// UPDATE
		merchantTokenRef01DataUp0Up := map[string]any{
			"id": merchantTokenRef01Data["id"],
		}

		merchantTokenRef01MarkdefUp0Name := "description"
		merchantTokenRef01MarkdefUp0Value := fmt.Sprintf("Mark01-merchant_token_ref01_%d", setup.now)
		merchantTokenRef01DataUp0Up[merchantTokenRef01MarkdefUp0Name] = merchantTokenRef01MarkdefUp0Value

		merchantTokenRef01ResdataUp0Result, err := merchantTokenRef01Ent.Update(merchantTokenRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		merchantTokenRef01ResdataUp0 := core.ToMapAny(merchantTokenRef01ResdataUp0Result)
		if merchantTokenRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if merchantTokenRef01ResdataUp0["id"] != merchantTokenRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if merchantTokenRef01ResdataUp0[merchantTokenRef01MarkdefUp0Name] != merchantTokenRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", merchantTokenRef01MarkdefUp0Name, merchantTokenRef01ResdataUp0[merchantTokenRef01MarkdefUp0Name])
		}

		// LOAD
		merchantTokenRef01MatchDt0 := map[string]any{
			"id": merchantTokenRef01Data["id"],
		}
		merchantTokenRef01DataDt0Loaded, err := merchantTokenRef01Ent.Load(merchantTokenRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		merchantTokenRef01DataDt0LoadResult := core.ToMapAny(merchantTokenRef01DataDt0Loaded)
		if merchantTokenRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if merchantTokenRef01DataDt0LoadResult["id"] != merchantTokenRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

	})
}

func merchant_tokenBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "merchant_token", "MerchantTokenTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read merchant_token test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse merchant_token test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"merchant_token01", "merchant_token02", "merchant_token03", "merchant01", "merchant02", "merchant03"},
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
	entidEnvRaw := os.Getenv("NOFRIXION_TEST_MERCHANT_TOKEN_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"NOFRIXION_TEST_MERCHANT_TOKEN_ENTID": idmap,
		"NOFRIXION_TEST_LIVE":      "FALSE",
		"NOFRIXION_TEST_EXPLAIN":   "FALSE",
		"NOFRIXION_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["NOFRIXION_TEST_MERCHANT_TOKEN_ENTID"])
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
