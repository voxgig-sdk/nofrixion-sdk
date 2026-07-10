<?php
declare(strict_types=1);

// Webhook entity test

require_once __DIR__ . '/../nofrixion_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class WebhookEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = NofrixionSDK::test(null, null);
        $ent = $testsdk->Webhook(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = webhook_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["create", "list", "update", "load", "remove"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "webhook." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_WEBHOOK_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // CREATE
        $webhook_ref01_ent = $client->Webhook(null);
        $webhook_ref01_data = Helpers::to_map(Vs::getprop(
            Vs::getpath($setup["data"], "new.webhook"), "webhook_ref01"));
        $webhook_ref01_data["merchant_id"] = $setup["idmap"]["merchant01"];

        $webhook_ref01_data_result = $webhook_ref01_ent->create($webhook_ref01_data, null);
        $webhook_ref01_data = Helpers::to_map($webhook_ref01_data_result);
        $this->assertNotNull($webhook_ref01_data);
        $this->assertNotNull($webhook_ref01_data["id"]);

        // LIST
        $webhook_ref01_match = [
            "merchant_id" => $setup["idmap"]["merchant01"],
        ];

        $webhook_ref01_list_result = $webhook_ref01_ent->list($webhook_ref01_match, null);
        $this->assertIsArray($webhook_ref01_list_result);

        $found_item = sdk_select(
            Runner::entity_list_to_data($webhook_ref01_list_result),
            ["id" => $webhook_ref01_data["id"]]);
        $this->assertNotEmpty($found_item);

        // UPDATE
        $webhook_ref01_data_up0_up = [
            "id" => $webhook_ref01_data["id"],
        ];

        $webhook_ref01_markdef_up0_name = "destination_url";
        $webhook_ref01_markdef_up0_value = "Mark01-webhook_ref01_" . $setup["now"];
        $webhook_ref01_data_up0_up[$webhook_ref01_markdef_up0_name] = $webhook_ref01_markdef_up0_value;

        $webhook_ref01_resdata_up0_result = $webhook_ref01_ent->update($webhook_ref01_data_up0_up, null);
        $webhook_ref01_resdata_up0 = Helpers::to_map($webhook_ref01_resdata_up0_result);
        $this->assertNotNull($webhook_ref01_resdata_up0);
        $this->assertEquals($webhook_ref01_resdata_up0["id"], $webhook_ref01_data_up0_up["id"]);
        $this->assertEquals($webhook_ref01_resdata_up0[$webhook_ref01_markdef_up0_name], $webhook_ref01_markdef_up0_value);

        // LOAD
        $webhook_ref01_match_dt0 = [
            "id" => $webhook_ref01_data["id"],
        ];
        $webhook_ref01_data_dt0_loaded = $webhook_ref01_ent->load($webhook_ref01_match_dt0, null);
        $webhook_ref01_data_dt0_load_result = Helpers::to_map($webhook_ref01_data_dt0_loaded);
        $this->assertNotNull($webhook_ref01_data_dt0_load_result);
        $this->assertEquals($webhook_ref01_data_dt0_load_result["id"], $webhook_ref01_data["id"]);

        // REMOVE
        $webhook_ref01_match_rm0 = [
            "id" => $webhook_ref01_data["id"],
        ];
        $webhook_ref01_ent->remove($webhook_ref01_match_rm0, null);

        // LIST
        $webhook_ref01_match_rt0 = [
            "merchant_id" => $setup["idmap"]["merchant01"],
        ];

        $webhook_ref01_list_rt0_result = $webhook_ref01_ent->list($webhook_ref01_match_rt0, null);
        $this->assertIsArray($webhook_ref01_list_rt0_result);

        $not_found_item = sdk_select(
            Runner::entity_list_to_data($webhook_ref01_list_rt0_result),
            ["id" => $webhook_ref01_data["id"]]);
        $this->assertEmpty($not_found_item);

    }
}

function webhook_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/webhook/WebhookTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = NofrixionSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["webhook01", "webhook02", "webhook03", "merchant01", "merchant02", "merchant03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("NOFRIXION_TEST_WEBHOOK_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "NOFRIXION_TEST_WEBHOOK_ENTID" => $idmap,
        "NOFRIXION_TEST_LIVE" => "FALSE",
        "NOFRIXION_TEST_EXPLAIN" => "FALSE",
        "NOFRIXION_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["NOFRIXION_TEST_WEBHOOK_ENTID"]);
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
