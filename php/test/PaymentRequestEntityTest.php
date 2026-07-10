<?php
declare(strict_types=1);

// PaymentRequest entity test

require_once __DIR__ . '/../nofrixion_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class PaymentRequestEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = NofrixionSDK::test(null, null);
        $ent = $testsdk->PaymentRequest(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = payment_request_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["create", "list", "update", "load", "remove"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "payment_request." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_PAYMENT_REQUEST_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // CREATE
        $payment_request_ref01_ent = $client->PaymentRequest(null);
        $payment_request_ref01_data = Helpers::to_map(Vs::getprop(
            Vs::getpath($setup["data"], "new.payment_request"), "payment_request_ref01"));
        $payment_request_ref01_data["paymentrequest_id"] = $setup["idmap"]["paymentrequest01"];

        $payment_request_ref01_data_result = $payment_request_ref01_ent->create($payment_request_ref01_data, null);
        $payment_request_ref01_data = Helpers::to_map($payment_request_ref01_data_result);
        $this->assertNotNull($payment_request_ref01_data);
        $this->assertNotNull($payment_request_ref01_data["id"]);

        // LIST
        $payment_request_ref01_match = [];

        $payment_request_ref01_list_result = $payment_request_ref01_ent->list($payment_request_ref01_match, null);
        $this->assertIsArray($payment_request_ref01_list_result);

        $found_item = sdk_select(
            Runner::entity_list_to_data($payment_request_ref01_list_result),
            ["id" => $payment_request_ref01_data["id"]]);
        $this->assertNotEmpty($found_item);

        // UPDATE
        $payment_request_ref01_data_up0_up = [
            "id" => $payment_request_ref01_data["id"],
        ];

        $payment_request_ref01_markdef_up0_name = "base_origin_url";
        $payment_request_ref01_markdef_up0_value = "Mark01-payment_request_ref01_" . $setup["now"];
        $payment_request_ref01_data_up0_up[$payment_request_ref01_markdef_up0_name] = $payment_request_ref01_markdef_up0_value;

        $payment_request_ref01_resdata_up0_result = $payment_request_ref01_ent->update($payment_request_ref01_data_up0_up, null);
        $payment_request_ref01_resdata_up0 = Helpers::to_map($payment_request_ref01_resdata_up0_result);
        $this->assertNotNull($payment_request_ref01_resdata_up0);
        $this->assertEquals($payment_request_ref01_resdata_up0["id"], $payment_request_ref01_data_up0_up["id"]);
        $this->assertEquals($payment_request_ref01_resdata_up0[$payment_request_ref01_markdef_up0_name], $payment_request_ref01_markdef_up0_value);

        // LOAD
        $payment_request_ref01_match_dt0 = [
            "id" => $payment_request_ref01_data["id"],
        ];
        $payment_request_ref01_data_dt0_loaded = $payment_request_ref01_ent->load($payment_request_ref01_match_dt0, null);
        $payment_request_ref01_data_dt0_load_result = Helpers::to_map($payment_request_ref01_data_dt0_loaded);
        $this->assertNotNull($payment_request_ref01_data_dt0_load_result);
        $this->assertEquals($payment_request_ref01_data_dt0_load_result["id"], $payment_request_ref01_data["id"]);

        // REMOVE
        $payment_request_ref01_match_rm0 = [
            "id" => $payment_request_ref01_data["id"],
        ];
        $payment_request_ref01_ent->remove($payment_request_ref01_match_rm0, null);

        // LIST
        $payment_request_ref01_match_rt0 = [];

        $payment_request_ref01_list_rt0_result = $payment_request_ref01_ent->list($payment_request_ref01_match_rt0, null);
        $this->assertIsArray($payment_request_ref01_list_rt0_result);

        $not_found_item = sdk_select(
            Runner::entity_list_to_data($payment_request_ref01_list_rt0_result),
            ["id" => $payment_request_ref01_data["id"]]);
        $this->assertEmpty($not_found_item);

    }
}

function payment_request_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/payment_request/PaymentRequestTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = NofrixionSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["payment_request01", "payment_request02", "payment_request03", "paymentrequest01", "paymentrequest02", "paymentrequest03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("NOFRIXION_TEST_PAYMENT_REQUEST_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "NOFRIXION_TEST_PAYMENT_REQUEST_ENTID" => $idmap,
        "NOFRIXION_TEST_LIVE" => "FALSE",
        "NOFRIXION_TEST_EXPLAIN" => "FALSE",
        "NOFRIXION_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["NOFRIXION_TEST_PAYMENT_REQUEST_ENTID"]);
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
