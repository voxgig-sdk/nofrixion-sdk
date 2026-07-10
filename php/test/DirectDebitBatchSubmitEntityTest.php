<?php
declare(strict_types=1);

// DirectDebitBatchSubmit entity test

require_once __DIR__ . '/../nofrixion_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class DirectDebitBatchSubmitEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = NofrixionSDK::test(null, null);
        $ent = $testsdk->DirectDebitBatchSubmit(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = direct_debit_batch_submit_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["create"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "direct_debit_batch_submit." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_DIRECT_DEBIT_BATCH_SUBMIT_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // CREATE
        $direct_debit_batch_submit_ref01_ent = $client->DirectDebitBatchSubmit(null);
        $direct_debit_batch_submit_ref01_data = Helpers::to_map(Vs::getprop(
            Vs::getpath($setup["data"], "new.direct_debit_batch_submit"), "direct_debit_batch_submit_ref01"));

        $direct_debit_batch_submit_ref01_data_result = $direct_debit_batch_submit_ref01_ent->create($direct_debit_batch_submit_ref01_data, null);
        $direct_debit_batch_submit_ref01_data = Helpers::to_map($direct_debit_batch_submit_ref01_data_result);
        $this->assertNotNull($direct_debit_batch_submit_ref01_data);

    }
}

function direct_debit_batch_submit_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/direct_debit_batch_submit/DirectDebitBatchSubmitTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = NofrixionSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["direct_debit_batch_submit01", "direct_debit_batch_submit02", "direct_debit_batch_submit03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("NOFRIXION_TEST_DIRECT_DEBIT_BATCH_SUBMIT_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "NOFRIXION_TEST_DIRECT_DEBIT_BATCH_SUBMIT_ENTID" => $idmap,
        "NOFRIXION_TEST_LIVE" => "FALSE",
        "NOFRIXION_TEST_EXPLAIN" => "FALSE",
        "NOFRIXION_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["NOFRIXION_TEST_DIRECT_DEBIT_BATCH_SUBMIT_ENTID"]);
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
