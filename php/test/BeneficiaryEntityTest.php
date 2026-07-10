<?php
declare(strict_types=1);

// Beneficiary entity test

require_once __DIR__ . '/../nofrixion_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class BeneficiaryEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = NofrixionSDK::test(null, null);
        $ent = $testsdk->Beneficiary(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = beneficiary_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["create", "list", "update", "load", "remove"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "beneficiary." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_BENEFICIARY_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // CREATE
        $beneficiary_ref01_ent = $client->Beneficiary(null);
        $beneficiary_ref01_data = Helpers::to_map(Vs::getprop(
            Vs::getpath($setup["data"], "new.beneficiary"), "beneficiary_ref01"));
        $beneficiary_ref01_data["merchant_id"] = $setup["idmap"]["merchant01"];

        $beneficiary_ref01_data_result = $beneficiary_ref01_ent->create($beneficiary_ref01_data, null);
        $beneficiary_ref01_data = Helpers::to_map($beneficiary_ref01_data_result);
        $this->assertNotNull($beneficiary_ref01_data);
        $this->assertNotNull($beneficiary_ref01_data["id"]);

        // LIST
        $beneficiary_ref01_match = [
            "merchant_id" => $setup["idmap"]["merchant01"],
        ];

        $beneficiary_ref01_list_result = $beneficiary_ref01_ent->list($beneficiary_ref01_match, null);
        $this->assertIsArray($beneficiary_ref01_list_result);

        $found_item = sdk_select(
            Runner::entity_list_to_data($beneficiary_ref01_list_result),
            ["id" => $beneficiary_ref01_data["id"]]);
        $this->assertNotEmpty($found_item);

        // UPDATE
        $beneficiary_ref01_data_up0_up = [
            "id" => $beneficiary_ref01_data["id"],
        ];

        $beneficiary_ref01_markdef_up0_name = "approval_callback_url";
        $beneficiary_ref01_markdef_up0_value = "Mark01-beneficiary_ref01_" . $setup["now"];
        $beneficiary_ref01_data_up0_up[$beneficiary_ref01_markdef_up0_name] = $beneficiary_ref01_markdef_up0_value;

        $beneficiary_ref01_resdata_up0_result = $beneficiary_ref01_ent->update($beneficiary_ref01_data_up0_up, null);
        $beneficiary_ref01_resdata_up0 = Helpers::to_map($beneficiary_ref01_resdata_up0_result);
        $this->assertNotNull($beneficiary_ref01_resdata_up0);
        $this->assertEquals($beneficiary_ref01_resdata_up0["id"], $beneficiary_ref01_data_up0_up["id"]);
        $this->assertEquals($beneficiary_ref01_resdata_up0[$beneficiary_ref01_markdef_up0_name], $beneficiary_ref01_markdef_up0_value);

        // LOAD
        $beneficiary_ref01_match_dt0 = [
            "id" => $beneficiary_ref01_data["id"],
        ];
        $beneficiary_ref01_data_dt0_loaded = $beneficiary_ref01_ent->load($beneficiary_ref01_match_dt0, null);
        $beneficiary_ref01_data_dt0_load_result = Helpers::to_map($beneficiary_ref01_data_dt0_loaded);
        $this->assertNotNull($beneficiary_ref01_data_dt0_load_result);
        $this->assertEquals($beneficiary_ref01_data_dt0_load_result["id"], $beneficiary_ref01_data["id"]);

        // REMOVE
        $beneficiary_ref01_match_rm0 = [
            "id" => $beneficiary_ref01_data["id"],
        ];
        $beneficiary_ref01_ent->remove($beneficiary_ref01_match_rm0, null);

        // LIST
        $beneficiary_ref01_match_rt0 = [
            "merchant_id" => $setup["idmap"]["merchant01"],
        ];

        $beneficiary_ref01_list_rt0_result = $beneficiary_ref01_ent->list($beneficiary_ref01_match_rt0, null);
        $this->assertIsArray($beneficiary_ref01_list_rt0_result);

        $not_found_item = sdk_select(
            Runner::entity_list_to_data($beneficiary_ref01_list_rt0_result),
            ["id" => $beneficiary_ref01_data["id"]]);
        $this->assertEmpty($not_found_item);

    }
}

function beneficiary_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/beneficiary/BeneficiaryTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = NofrixionSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["beneficiary01", "beneficiary02", "beneficiary03", "merchant01", "merchant02", "merchant03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("NOFRIXION_TEST_BENEFICIARY_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "NOFRIXION_TEST_BENEFICIARY_ENTID" => $idmap,
        "NOFRIXION_TEST_LIVE" => "FALSE",
        "NOFRIXION_TEST_EXPLAIN" => "FALSE",
        "NOFRIXION_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["NOFRIXION_TEST_BENEFICIARY_ENTID"]);
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
