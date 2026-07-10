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

func TestUserInviteEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.UserInvite(nil)
		if ent == nil {
			t.Fatal("expected non-nil UserInviteEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := user_inviteBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "update", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "user_invite." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_USER_INVITE_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		userInviteRef01Ent := client.UserInvite(nil)
		userInviteRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "user_invite"}, setup.data), "user_invite_ref01"))
		userInviteRef01Data["merchant_id"] = setup.idmap["merchant01"]

		userInviteRef01DataResult, err := userInviteRef01Ent.Create(userInviteRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		userInviteRef01Data = core.ToMapAny(userInviteRef01DataResult)
		if userInviteRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if userInviteRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		userInviteRef01Match := map[string]any{
			"merchant_id": setup.idmap["merchant01"],
		}

		userInviteRef01ListResult, err := userInviteRef01Ent.List(userInviteRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		userInviteRef01List, userInviteRef01ListOk := userInviteRef01ListResult.([]any)
		if !userInviteRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", userInviteRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(userInviteRef01List), map[string]any{"id": userInviteRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// UPDATE
		userInviteRef01DataUp0Up := map[string]any{
			"id": userInviteRef01Data["id"],
		}

		userInviteRef01MarkdefUp0Name := "initial_role_id"
		userInviteRef01MarkdefUp0Value := fmt.Sprintf("Mark01-user_invite_ref01_%d", setup.now)
		userInviteRef01DataUp0Up[userInviteRef01MarkdefUp0Name] = userInviteRef01MarkdefUp0Value

		userInviteRef01ResdataUp0Result, err := userInviteRef01Ent.Update(userInviteRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		userInviteRef01ResdataUp0 := core.ToMapAny(userInviteRef01ResdataUp0Result)
		if userInviteRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if userInviteRef01ResdataUp0["id"] != userInviteRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if userInviteRef01ResdataUp0[userInviteRef01MarkdefUp0Name] != userInviteRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", userInviteRef01MarkdefUp0Name, userInviteRef01ResdataUp0[userInviteRef01MarkdefUp0Name])
		}

		// LOAD
		userInviteRef01MatchDt0 := map[string]any{
			"id": userInviteRef01Data["id"],
		}
		userInviteRef01DataDt0Loaded, err := userInviteRef01Ent.Load(userInviteRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		userInviteRef01DataDt0LoadResult := core.ToMapAny(userInviteRef01DataDt0Loaded)
		if userInviteRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if userInviteRef01DataDt0LoadResult["id"] != userInviteRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

		// REMOVE
		userInviteRef01MatchRm0 := map[string]any{
			"id": userInviteRef01Data["id"],
		}
		_, err = userInviteRef01Ent.Remove(userInviteRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		userInviteRef01MatchRt0 := map[string]any{
			"merchant_id": setup.idmap["merchant01"],
		}

		userInviteRef01ListRt0Result, err := userInviteRef01Ent.List(userInviteRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		userInviteRef01ListRt0, userInviteRef01ListRt0Ok := userInviteRef01ListRt0Result.([]any)
		if !userInviteRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", userInviteRef01ListRt0Result)
		}

		notFoundItem := vs.Select(entityListToData(userInviteRef01ListRt0), map[string]any{"id": userInviteRef01Data["id"]})
		if !vs.IsEmpty(notFoundItem) {
			t.Fatal("expected removed entity to not be in list")
		}

	})
}

func user_inviteBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "user_invite", "UserInviteTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read user_invite test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse user_invite test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"user_invite01", "user_invite02", "user_invite03", "merchant01", "merchant02", "merchant03", "userinvite01", "userinvite02", "userinvite03"},
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
	entidEnvRaw := os.Getenv("NOFRIXION_TEST_USER_INVITE_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"NOFRIXION_TEST_USER_INVITE_ENTID": idmap,
		"NOFRIXION_TEST_LIVE":      "FALSE",
		"NOFRIXION_TEST_EXPLAIN":   "FALSE",
		"NOFRIXION_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["NOFRIXION_TEST_USER_INVITE_ENTID"])
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
