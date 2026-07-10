<?php
declare(strict_types=1);

// Payout entity test

require_once __DIR__ . '/../nofrixion_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class PayoutEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = NofrixionSDK::test(null, null);
        $ent = $testsdk->Payout(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = payout_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["create", "list", "update", "load", "remove"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "payout." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_PAYOUT_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // CREATE
        $payout_ref01_ent = $client->Payout(null);
        $payout_ref01_data = Helpers::to_map(Vs::getprop(
            Vs::getpath($setup["data"], "new.payout"), "payout_ref01"));
        $payout_ref01_data["account_id"] = $setup["idmap"]["account01"];
        $payout_ref01_data["destination"] = $setup["idmap"]["destination01"];
        $payout_ref01_data["merchant_id"] = $setup["idmap"]["merchant01"];
        $payout_ref01_data["source"] = $setup["idmap"]["source01"];

        $payout_ref01_data_result = $payout_ref01_ent->create($payout_ref01_data, null);
        $payout_ref01_data = Helpers::to_map($payout_ref01_data_result);
        $this->assertNotNull($payout_ref01_data);
        $this->assertNotNull($payout_ref01_data["id"]);

        // LIST
        $payout_ref01_match = [
            "merchant_id" => $setup["idmap"]["merchant01"],
        ];

        $payout_ref01_list_result = $payout_ref01_ent->list($payout_ref01_match, null);
        $this->assertIsArray($payout_ref01_list_result);

        $found_item = sdk_select(
            Runner::entity_list_to_data($payout_ref01_list_result),
            ["id" => $payout_ref01_data["id"]]);
        $this->assertNotEmpty($found_item);

        // UPDATE
        $payout_ref01_data_up0_up = [
            "id" => $payout_ref01_data["id"],
        ];

        $payout_ref01_markdef_up0_name = "account_id";
        $payout_ref01_markdef_up0_value = "Mark01-payout_ref01_" . $setup["now"];
        $payout_ref01_data_up0_up[$payout_ref01_markdef_up0_name] = $payout_ref01_markdef_up0_value;

        $payout_ref01_resdata_up0_result = $payout_ref01_ent->update($payout_ref01_data_up0_up, null);
        $payout_ref01_resdata_up0 = Helpers::to_map($payout_ref01_resdata_up0_result);
        $this->assertNotNull($payout_ref01_resdata_up0);
        $this->assertEquals($payout_ref01_resdata_up0["id"], $payout_ref01_data_up0_up["id"]);
        $this->assertEquals($payout_ref01_resdata_up0[$payout_ref01_markdef_up0_name], $payout_ref01_markdef_up0_value);

        // LOAD
        $payout_ref01_match_dt0 = [
            "id" => $payout_ref01_data["id"],
        ];
        $payout_ref01_data_dt0_loaded = $payout_ref01_ent->load($payout_ref01_match_dt0, null);
        $payout_ref01_data_dt0_load_result = Helpers::to_map($payout_ref01_data_dt0_loaded);
        $this->assertNotNull($payout_ref01_data_dt0_load_result);
        $this->assertEquals($payout_ref01_data_dt0_load_result["id"], $payout_ref01_data["id"]);

        // REMOVE
        $payout_ref01_match_rm0 = [
            "id" => $payout_ref01_data["id"],
        ];
        $payout_ref01_ent->remove($payout_ref01_match_rm0, null);

        // LIST
        $payout_ref01_match_rt0 = [
            "merchant_id" => $setup["idmap"]["merchant01"],
        ];

        $payout_ref01_list_rt0_result = $payout_ref01_ent->list($payout_ref01_match_rt0, null);
        $this->assertIsArray($payout_ref01_list_rt0_result);

        $not_found_item = sdk_select(
            Runner::entity_list_to_data($payout_ref01_list_rt0_result),
            ["id" => $payout_ref01_data["id"]]);
        $this->assertEmpty($not_found_item);

    }
}

function payout_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/payout/PayoutTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = NofrixionSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["payout01", "payout02", "payout03", "account01", "account02", "account03", "merchant01", "merchant02", "merchant03", "fxquote01", "fxquote02", "fxquote03", "destination01", "source01"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("NOFRIXION_TEST_PAYOUT_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "NOFRIXION_TEST_PAYOUT_ENTID" => $idmap,
        "NOFRIXION_TEST_LIVE" => "FALSE",
        "NOFRIXION_TEST_EXPLAIN" => "FALSE",
        "NOFRIXION_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["NOFRIXION_TEST_PAYOUT_ENTID"]);
    if ($idmap_resolved === null) {
        $idmap_resolved = Helpers::to_map($idmap);
    }

    if ($env["NOFRIXION_TEST_LIVE"] === "TRUE") {
        $merged_opts = Vs::merge([
            [
                "apikey" => $env["NOFRIXION_APIKEY"],
            ],
            $extra ?? [],
        ]);
        $client = new NofrixionSDK(Helpers::to_map($merged_opts));
    }

    $live = $env["NOFRIXION_TEST_LIVE"] === "TRUE";
    return [
        "client" => $client,
        "data" => $entity_data,
        "idmap" => $idmap_resolved,
        "env" => $env,
        "explain" => $env["NOFRIXION_TEST_EXPLAIN"] === "TRUE",
        "live" => $live,
        "synthetic_only" => $live && !$idmap_overridden,
        "now" => (int)(microtime(true) * 1000),
    ];
}
