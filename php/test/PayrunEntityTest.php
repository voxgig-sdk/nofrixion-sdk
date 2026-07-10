<?php
declare(strict_types=1);

// Payrun entity test

require_once __DIR__ . '/../nofrixion_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class PayrunEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = NofrixionSDK::test(null, null);
        $ent = $testsdk->Payrun(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = payrun_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["create", "update", "remove"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "payrun." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_PAYRUN_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // CREATE
        $payrun_ref01_ent = $client->Payrun(null);
        $payrun_ref01_data = Helpers::to_map(Vs::getprop(
            Vs::getpath($setup["data"], "new.payrun"), "payrun_ref01"));

        $payrun_ref01_data_result = $payrun_ref01_ent->create($payrun_ref01_data, null);
        $payrun_ref01_data = Helpers::to_map($payrun_ref01_data_result);
        $this->assertNotNull($payrun_ref01_data);
        $this->assertNotNull($payrun_ref01_data["id"]);

        // UPDATE
        $payrun_ref01_data_up0_up = [
            "id" => $payrun_ref01_data["id"],
        ];

        $payrun_ref01_markdef_up0_name = "note";
        $payrun_ref01_markdef_up0_value = "Mark01-payrun_ref01_" . $setup["now"];
        $payrun_ref01_data_up0_up[$payrun_ref01_markdef_up0_name] = $payrun_ref01_markdef_up0_value;

        $payrun_ref01_resdata_up0_result = $payrun_ref01_ent->update($payrun_ref01_data_up0_up, null);
        $payrun_ref01_resdata_up0 = Helpers::to_map($payrun_ref01_resdata_up0_result);
        $this->assertNotNull($payrun_ref01_resdata_up0);
        $this->assertEquals($payrun_ref01_resdata_up0["id"], $payrun_ref01_data_up0_up["id"]);
        $this->assertEquals($payrun_ref01_resdata_up0[$payrun_ref01_markdef_up0_name], $payrun_ref01_markdef_up0_value);

        // REMOVE
        $payrun_ref01_match_rm0 = [
            "id" => $payrun_ref01_data["id"],
        ];
        $payrun_ref01_ent->remove($payrun_ref01_match_rm0, null);

    }
}

function payrun_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/payrun/PayrunTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = NofrixionSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["payrun01", "payrun02", "payrun03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("NOFRIXION_TEST_PAYRUN_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "NOFRIXION_TEST_PAYRUN_ENTID" => $idmap,
        "NOFRIXION_TEST_LIVE" => "FALSE",
        "NOFRIXION_TEST_EXPLAIN" => "FALSE",
        "NOFRIXION_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["NOFRIXION_TEST_PAYRUN_ENTID"]);
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
