<?php
declare(strict_types=1);

// Transaction entity test

require_once __DIR__ . '/../nofrixion_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class TransactionEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = NofrixionSDK::test(null, null);
        $ent = $testsdk->Transaction(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = transaction_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["create", "list", "load", "remove"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "transaction." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_TRANSACTION_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // CREATE
        $transaction_ref01_ent = $client->Transaction(null);
        $transaction_ref01_data = Helpers::to_map(Vs::getprop(
            Vs::getpath($setup["data"], "new.transaction"), "transaction_ref01"));
        $transaction_ref01_data["account_id"] = $setup["idmap"]["account01"];
        $transaction_ref01_data["merchant_id"] = $setup["idmap"]["merchant01"];
        $transaction_ref01_data["transaction_id"] = $setup["idmap"]["transaction01"];

        $transaction_ref01_data_result = $transaction_ref01_ent->create($transaction_ref01_data, null);
        $transaction_ref01_data = Helpers::to_map($transaction_ref01_data_result);
        $this->assertNotNull($transaction_ref01_data);
        $this->assertNotNull($transaction_ref01_data["id"]);

        // LIST
        $transaction_ref01_match = [];

        $transaction_ref01_list_result = $transaction_ref01_ent->list($transaction_ref01_match, null);
        $this->assertIsArray($transaction_ref01_list_result);

        $found_item = sdk_select(
            Runner::entity_list_to_data($transaction_ref01_list_result),
            ["id" => $transaction_ref01_data["id"]]);
        $this->assertNotEmpty($found_item);

        // LOAD
        $transaction_ref01_match_dt0 = [
            "id" => $transaction_ref01_data["id"],
        ];
        $transaction_ref01_data_dt0_loaded = $transaction_ref01_ent->load($transaction_ref01_match_dt0, null);
        $transaction_ref01_data_dt0_load_result = Helpers::to_map($transaction_ref01_data_dt0_loaded);
        $this->assertNotNull($transaction_ref01_data_dt0_load_result);
        $this->assertEquals($transaction_ref01_data_dt0_load_result["id"], $transaction_ref01_data["id"]);

        // REMOVE
        $transaction_ref01_match_rm0 = [
            "id" => $transaction_ref01_data["id"],
        ];
        $transaction_ref01_ent->remove($transaction_ref01_match_rm0, null);

        // LIST
        $transaction_ref01_match_rt0 = [];

        $transaction_ref01_list_rt0_result = $transaction_ref01_ent->list($transaction_ref01_match_rt0, null);
        $this->assertIsArray($transaction_ref01_list_rt0_result);

        $not_found_item = sdk_select(
            Runner::entity_list_to_data($transaction_ref01_list_rt0_result),
            ["id" => $transaction_ref01_data["id"]]);
        $this->assertEmpty($not_found_item);

    }
}

function transaction_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/transaction/TransactionTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = NofrixionSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["transaction01", "transaction02", "transaction03", "account01", "account02", "account03", "merchant01", "merchant02", "merchant03", "from01", "from02", "from03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("NOFRIXION_TEST_TRANSACTION_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "NOFRIXION_TEST_TRANSACTION_ENTID" => $idmap,
        "NOFRIXION_TEST_LIVE" => "FALSE",
        "NOFRIXION_TEST_EXPLAIN" => "FALSE",
        "NOFRIXION_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["NOFRIXION_TEST_TRANSACTION_ENTID"]);
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
