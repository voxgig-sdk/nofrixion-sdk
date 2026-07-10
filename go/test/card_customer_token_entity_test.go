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

func TestCardCustomerTokenEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.CardCustomerToken(nil)
		if ent == nil {
			t.Fatal("expected non-nil CardCustomerTokenEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := card_customer_tokenBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"list", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "card_customer_token." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_CARD_CUSTOMER_TOKEN_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		cardCustomerTokenRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.card_customer_token", setup.data)))
		var cardCustomerTokenRef01Data map[string]any
		if len(cardCustomerTokenRef01DataRaw) > 0 {
			cardCustomerTokenRef01Data = core.ToMapAny(cardCustomerTokenRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = cardCustomerTokenRef01Data

		// LIST
		cardCustomerTokenRef01Ent := client.CardCustomerToken(nil)
		cardCustomerTokenRef01Match := map[string]any{
			"customer_email_address": setup.idmap["customer_email_address01"],
			"merchant_id": setup.idmap["merchant01"],
		}

		cardCustomerTokenRef01ListResult, err := cardCustomerTokenRef01Ent.List(cardCustomerTokenRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, cardCustomerTokenRef01ListOk := cardCustomerTokenRef01ListResult.([]any)
		if !cardCustomerTokenRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", cardCustomerTokenRef01ListResult)
		}

		// LOAD
		cardCustomerTokenRef01MatchDt0 := map[string]any{
			"id": cardCustomerTokenRef01Data["id"],
		}
		cardCustomerTokenRef01DataDt0Loaded, err := cardCustomerTokenRef01Ent.Load(cardCustomerTokenRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		cardCustomerTokenRef01DataDt0LoadResult := core.ToMapAny(cardCustomerTokenRef01DataDt0Loaded)
		if cardCustomerTokenRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if cardCustomerTokenRef01DataDt0LoadResult["id"] != cardCustomerTokenRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

		// REMOVE
		cardCustomerTokenRef01MatchRm0 := map[string]any{
			"id": cardCustomerTokenRef01Data["id"],
		}
		_, err = cardCustomerTokenRef01Ent.Remove(cardCustomerTokenRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		cardCustomerTokenRef01MatchRt0 := map[string]any{
			"customer_email_address": setup.idmap["customer_email_address01"],
			"merchant_id": setup.idmap["merchant01"],
		}

		cardCustomerTokenRef01ListRt0Result, err := cardCustomerTokenRef01Ent.List(cardCustomerTokenRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, cardCustomerTokenRef01ListRt0Ok := cardCustomerTokenRef01ListRt0Result.([]any)
		if !cardCustomerTokenRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", cardCustomerTokenRef01ListRt0Result)
		}

	})
}

func card_customer_tokenBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "card_customer_token", "CardCustomerTokenTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read card_customer_token test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse card_customer_token test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"card_customer_token01", "card_customer_token02", "card_customer_token03", "removeall01", "removeall02", "removeall03", "customertoken01", "customertoken02", "customertoken03", "customer_email_address01", "merchant01"},
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
	entidEnvRaw := os.Getenv("NOFRIXION_TEST_CARD_CUSTOMER_TOKEN_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"NOFRIXION_TEST_CARD_CUSTOMER_TOKEN_ENTID": idmap,
		"NOFRIXION_TEST_LIVE":      "FALSE",
		"NOFRIXION_TEST_EXPLAIN":   "FALSE",
		"NOFRIXION_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["NOFRIXION_TEST_CARD_CUSTOMER_TOKEN_ENTID"])
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
