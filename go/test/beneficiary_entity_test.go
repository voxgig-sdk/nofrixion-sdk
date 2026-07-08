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

func TestBeneficiaryEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.Beneficiary(nil)
		if ent == nil {
			t.Fatal("expected non-nil BeneficiaryEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := beneficiaryBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "update", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "beneficiary." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_BENEFICIARY_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		beneficiaryRef01Ent := client.Beneficiary(nil)
		beneficiaryRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "beneficiary"}, setup.data), "beneficiary_ref01"))
		beneficiaryRef01Data["merchant_id"] = setup.idmap["merchant01"]

		beneficiaryRef01DataResult, err := beneficiaryRef01Ent.Create(beneficiaryRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		beneficiaryRef01Data = core.ToMapAny(beneficiaryRef01DataResult)
		if beneficiaryRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if beneficiaryRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// UPDATE
		beneficiaryRef01DataUp0Up := map[string]any{
			"id": beneficiaryRef01Data["id"],
		}

		beneficiaryRef01MarkdefUp0Name := "approval_callback_url"
		beneficiaryRef01MarkdefUp0Value := fmt.Sprintf("Mark01-beneficiary_ref01_%d", setup.now)
		beneficiaryRef01DataUp0Up[beneficiaryRef01MarkdefUp0Name] = beneficiaryRef01MarkdefUp0Value

		beneficiaryRef01ResdataUp0Result, err := beneficiaryRef01Ent.Update(beneficiaryRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		beneficiaryRef01ResdataUp0 := core.ToMapAny(beneficiaryRef01ResdataUp0Result)
		if beneficiaryRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if beneficiaryRef01ResdataUp0["id"] != beneficiaryRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if beneficiaryRef01ResdataUp0[beneficiaryRef01MarkdefUp0Name] != beneficiaryRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", beneficiaryRef01MarkdefUp0Name, beneficiaryRef01ResdataUp0[beneficiaryRef01MarkdefUp0Name])
		}

		// LOAD
		beneficiaryRef01MatchDt0 := map[string]any{
			"id": beneficiaryRef01Data["id"],
		}
		beneficiaryRef01DataDt0Loaded, err := beneficiaryRef01Ent.Load(beneficiaryRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		beneficiaryRef01DataDt0LoadResult := core.ToMapAny(beneficiaryRef01DataDt0Loaded)
		if beneficiaryRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if beneficiaryRef01DataDt0LoadResult["id"] != beneficiaryRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

		// REMOVE
		beneficiaryRef01MatchRm0 := map[string]any{
			"id": beneficiaryRef01Data["id"],
		}
		_, err = beneficiaryRef01Ent.Remove(beneficiaryRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

	})
}

func beneficiaryBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "beneficiary", "BeneficiaryTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read beneficiary test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse beneficiary test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"beneficiary01", "beneficiary02", "beneficiary03", "merchant01", "merchant02", "merchant03"},
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
	entidEnvRaw := os.Getenv("NOFRIXION_TEST_BENEFICIARY_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"NOFRIXION_TEST_BENEFICIARY_ENTID": idmap,
		"NOFRIXION_TEST_LIVE":      "FALSE",
		"NOFRIXION_TEST_EXPLAIN":   "FALSE",
		"NOFRIXION_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["NOFRIXION_TEST_BENEFICIARY_ENTID"])
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
