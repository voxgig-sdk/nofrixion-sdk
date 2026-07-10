<?php
declare(strict_types=1);

// NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4 entity test

require_once __DIR__ . '/../nofrixion_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4EntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = NofrixionSDK::test(null, null);
        $ent = $testsdk->NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = no_frixion_money_moov_models_payment_requests_merchant_payment4_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["remove"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "no_frixion_money_moov_models_payment_requests_merchant_payment4." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_PAYMENT_REQUESTS_MERCHANT_PAYMENT__ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // Bootstrap entity data from existing test data.
        $no_frixion_money_moov_models_payment_requests_merchant_payment4_ref01_data_raw = Vs::items(Helpers::to_map(
            Vs::getpath($setup["data"], "existing.no_frixion_money_moov_models_payment_requests_merchant_payment4")));
        $no_frixion_money_moov_models_payment_requests_merchant_payment4_ref01_data = null;
        if (count($no_frixion_money_moov_models_payment_requests_merchant_payment4_ref01_data_raw) > 0) {
            $no_frixion_money_moov_models_payment_requests_merchant_payment4_ref01_data = Helpers::to_map($no_frixion_money_moov_models_payment_requests_merchant_payment4_ref01_data_raw[0][1]);
        }

        // REMOVE
        $no_frixion_money_moov_models_payment_requests_merchant_payment4_ref01_ent = $client->NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4(null);
        $no_frixion_money_moov_models_payment_requests_merchant_payment4_ref01_match_rm0 = [
            "id" => $no_frixion_money_moov_models_payment_requests_merchant_payment4_ref01_data["id"],
        ];
        $no_frixion_money_moov_models_payment_requests_merchant_payment4_ref01_ent->remove($no_frixion_money_moov_models_payment_requests_merchant_payment4_ref01_match_rm0, null);

    }
}

function no_frixion_money_moov_models_payment_requests_merchant_payment4_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/no_frixion_money_moov_models_payment_requests_merchant_payment4/NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4TestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = NofrixionSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["no_frixion_money_moov_models_payment_requests_merchant_payment401", "no_frixion_money_moov_models_payment_requests_merchant_payment402", "no_frixion_money_moov_models_payment_requests_merchant_payment403", "paymentrequest01", "paymentrequest02", "paymentrequest03", "template01", "template02", "template03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_PAYMENT_REQUESTS_MERCHANT_PAYMENT__ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_PAYMENT_REQUESTS_MERCHANT_PAYMENT__ENTID" => $idmap,
        "NOFRIXION_TEST_LIVE" => "FALSE",
        "NOFRIXION_TEST_EXPLAIN" => "FALSE",
        "NOFRIXION_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_PAYMENT_REQUESTS_MERCHANT_PAYMENT__ENTID"]);
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
