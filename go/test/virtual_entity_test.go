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

func TestVirtualEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.Virtual(nil)
		if ent == nil {
			t.Fatal("expected non-nil VirtualEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := virtualBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "update"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "virtual." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_VIRTUAL_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		virtualRef01Ent := client.Virtual(nil)
		virtualRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "virtual"}, setup.data), "virtual_ref01"))
		virtualRef01Data["account_id"] = setup.idmap["account01"]

		virtualRef01DataResult, err := virtualRef01Ent.Create(virtualRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		virtualRef01Data = core.ToMapAny(virtualRef01DataResult)
		if virtualRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if virtualRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// UPDATE
		virtualRef01DataUp0Up := map[string]any{
			"id": virtualRef01Data["id"],
			"account_id": setup.idmap["account_id"],
		}

		virtualRef01MarkdefUp0Name := "account_name"
		virtualRef01MarkdefUp0Value := fmt.Sprintf("Mark01-virtual_ref01_%d", setup.now)
		virtualRef01DataUp0Up[virtualRef01MarkdefUp0Name] = virtualRef01MarkdefUp0Value

		virtualRef01ResdataUp0Result, err := virtualRef01Ent.Update(virtualRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		virtualRef01ResdataUp0 := core.ToMapAny(virtualRef01ResdataUp0Result)
		if virtualRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if virtualRef01ResdataUp0["id"] != virtualRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if virtualRef01ResdataUp0[virtualRef01MarkdefUp0Name] != virtualRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", virtualRef01MarkdefUp0Name, virtualRef01ResdataUp0[virtualRef01MarkdefUp0Name])
		}

	})
}

func virtualBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "virtual", "VirtualTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read virtual test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse virtual test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"virtual01", "virtual02", "virtual03", "account01", "account02", "account03"},
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
	entidEnvRaw := os.Getenv("NOFRIXION_TEST_VIRTUAL_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"NOFRIXION_TEST_VIRTUAL_ENTID": idmap,
		"NOFRIXION_TEST_LIVE":      "FALSE",
		"NOFRIXION_TEST_EXPLAIN":   "FALSE",
		"NOFRIXION_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["NOFRIXION_TEST_VIRTUAL_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}
	// Add account_id alias for update test.
	if idmapResolved["account_id"] == nil {
		idmapResolved["account_id"] = idmapResolved["account01"]
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
