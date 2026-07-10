<?php
declare(strict_types=1);

// NoFrixionMoneyMoovModelsRule entity test

require_once __DIR__ . '/../nofrixion_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class NoFrixionMoneyMoovModelsRuleEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = NofrixionSDK::test(null, null);
        $ent = $testsdk->NoFrixionMoneyMoovModelsRule(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = no_frixion_money_moov_models_rule_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["create", "update", "load"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "no_frixion_money_moov_models_rule." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_RULE_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // CREATE
        $no_frixion_money_moov_models_rule_ref01_ent = $client->NoFrixionMoneyMoovModelsRule(null);
        $no_frixion_money_moov_models_rule_ref01_data = Helpers::to_map(Vs::getprop(
            Vs::getpath($setup["data"], "new.no_frixion_money_moov_models_rule"), "no_frixion_money_moov_models_rule_ref01"));

        $no_frixion_money_moov_models_rule_ref01_data_result = $no_frixion_money_moov_models_rule_ref01_ent->create($no_frixion_money_moov_models_rule_ref01_data, null);
        $no_frixion_money_moov_models_rule_ref01_data = Helpers::to_map($no_frixion_money_moov_models_rule_ref01_data_result);
        $this->assertNotNull($no_frixion_money_moov_models_rule_ref01_data);
        $this->assertNotNull($no_frixion_money_moov_models_rule_ref01_data["id"]);

        // UPDATE
        $no_frixion_money_moov_models_rule_ref01_data_up0_up = [
            "id" => $no_frixion_money_moov_models_rule_ref01_data["id"],
        ];

        $no_frixion_money_moov_models_rule_ref01_markdef_up0_name = "account_id";
        $no_frixion_money_moov_models_rule_ref01_markdef_up0_value = "Mark01-no_frixion_money_moov_models_rule_ref01_" . $setup["now"];
        $no_frixion_money_moov_models_rule_ref01_data_up0_up[$no_frixion_money_moov_models_rule_ref01_markdef_up0_name] = $no_frixion_money_moov_models_rule_ref01_markdef_up0_value;

        $no_frixion_money_moov_models_rule_ref01_resdata_up0_result = $no_frixion_money_moov_models_rule_ref01_ent->update($no_frixion_money_moov_models_rule_ref01_data_up0_up, null);
        $no_frixion_money_moov_models_rule_ref01_resdata_up0 = Helpers::to_map($no_frixion_money_moov_models_rule_ref01_resdata_up0_result);
        $this->assertNotNull($no_frixion_money_moov_models_rule_ref01_resdata_up0);
        $this->assertEquals($no_frixion_money_moov_models_rule_ref01_resdata_up0["id"], $no_frixion_money_moov_models_rule_ref01_data_up0_up["id"]);
        $this->assertEquals($no_frixion_money_moov_models_rule_ref01_resdata_up0[$no_frixion_money_moov_models_rule_ref01_markdef_up0_name], $no_frixion_money_moov_models_rule_ref01_markdef_up0_value);

        // LOAD
        $no_frixion_money_moov_models_rule_ref01_match_dt0 = [
            "id" => $no_frixion_money_moov_models_rule_ref01_data["id"],
        ];
        $no_frixion_money_moov_models_rule_ref01_data_dt0_loaded = $no_frixion_money_moov_models_rule_ref01_ent->load($no_frixion_money_moov_models_rule_ref01_match_dt0, null);
        $no_frixion_money_moov_models_rule_ref01_data_dt0_load_result = Helpers::to_map($no_frixion_money_moov_models_rule_ref01_data_dt0_loaded);
        $this->assertNotNull($no_frixion_money_moov_models_rule_ref01_data_dt0_load_result);
        $this->assertEquals($no_frixion_money_moov_models_rule_ref01_data_dt0_load_result["id"], $no_frixion_money_moov_models_rule_ref01_data["id"]);

    }
}

function no_frixion_money_moov_models_rule_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/no_frixion_money_moov_models_rule/NoFrixionMoneyMoovModelsRuleTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = NofrixionSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["no_frixion_money_moov_models_rule01", "no_frixion_money_moov_models_rule02", "no_frixion_money_moov_models_rule03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_RULE_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_RULE_ENTID" => $idmap,
        "NOFRIXION_TEST_LIVE" => "FALSE",
        "NOFRIXION_TEST_EXPLAIN" => "FALSE",
        "NOFRIXION_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_RULE_ENTID"]);
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
