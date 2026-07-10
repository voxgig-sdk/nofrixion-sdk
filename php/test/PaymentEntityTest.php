<?php
declare(strict_types=1);

// Payment entity test

require_once __DIR__ . '/../nofrixion_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class PaymentEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = NofrixionSDK::test(null, null);
        $ent = $testsdk->Payment(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = payment_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["create", "update", "load"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "payment." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_PAYMENT_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // CREATE
        $payment_ref01_ent = $client->Payment(null);
        $payment_ref01_data = Helpers::to_map(Vs::getprop(
            Vs::getpath($setup["data"], "new.payment"), "payment_ref01"));

        $payment_ref01_data_result = $payment_ref01_ent->create($payment_ref01_data, null);
        $payment_ref01_data = Helpers::to_map($payment_ref01_data_result);
        $this->assertNotNull($payment_ref01_data);
        $this->assertNotNull($payment_ref01_data["id"]);

        // UPDATE
        $payment_ref01_data_up0_up = [
            "id" => $payment_ref01_data["id"],
        ];

        $payment_ref01_markdef_up0_name = "base_origin_url";
        $payment_ref01_markdef_up0_value = "Mark01-payment_ref01_" . $setup["now"];
        $payment_ref01_data_up0_up[$payment_ref01_markdef_up0_name] = $payment_ref01_markdef_up0_value;

        $payment_ref01_resdata_up0_result = $payment_ref01_ent->update($payment_ref01_data_up0_up, null);
        $payment_ref01_resdata_up0 = Helpers::to_map($payment_ref01_resdata_up0_result);
        $this->assertNotNull($payment_ref01_resdata_up0);
        $this->assertEquals($payment_ref01_resdata_up0["id"], $payment_ref01_data_up0_up["id"]);
        $this->assertEquals($payment_ref01_resdata_up0[$payment_ref01_markdef_up0_name], $payment_ref01_markdef_up0_value);

        // LOAD
        $payment_ref01_match_dt0 = [
            "id" => $payment_ref01_data["id"],
        ];
        $payment_ref01_data_dt0_loaded = $payment_ref01_ent->load($payment_ref01_match_dt0, null);
        $payment_ref01_data_dt0_load_result = Helpers::to_map($payment_ref01_data_dt0_loaded);
        $this->assertNotNull($payment_ref01_data_dt0_load_result);
        $this->assertEquals($payment_ref01_data_dt0_load_result["id"], $payment_ref01_data["id"]);

    }
}

function payment_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/payment/PaymentTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = NofrixionSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["payment01", "payment02", "payment03", "getbyorderid01", "getbyorderid02", "getbyorderid03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("NOFRIXION_TEST_PAYMENT_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "NOFRIXION_TEST_PAYMENT_ENTID" => $idmap,
        "NOFRIXION_TEST_LIVE" => "FALSE",
        "NOFRIXION_TEST_EXPLAIN" => "FALSE",
        "NOFRIXION_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["NOFRIXION_TEST_PAYMENT_ENTID"]);
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
