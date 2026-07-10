<?php
declare(strict_types=1);

// UserInvite entity test

require_once __DIR__ . '/../nofrixion_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class UserInviteEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = NofrixionSDK::test(null, null);
        $ent = $testsdk->UserInvite(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = user_invite_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["create", "list", "update", "load", "remove"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "user_invite." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_USER_INVITE_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // CREATE
        $user_invite_ref01_ent = $client->UserInvite(null);
        $user_invite_ref01_data = Helpers::to_map(Vs::getprop(
            Vs::getpath($setup["data"], "new.user_invite"), "user_invite_ref01"));
        $user_invite_ref01_data["merchant_id"] = $setup["idmap"]["merchant01"];

        $user_invite_ref01_data_result = $user_invite_ref01_ent->create($user_invite_ref01_data, null);
        $user_invite_ref01_data = Helpers::to_map($user_invite_ref01_data_result);
        $this->assertNotNull($user_invite_ref01_data);
        $this->assertNotNull($user_invite_ref01_data["id"]);

        // LIST
        $user_invite_ref01_match = [
            "merchant_id" => $setup["idmap"]["merchant01"],
        ];

        $user_invite_ref01_list_result = $user_invite_ref01_ent->list($user_invite_ref01_match, null);
        $this->assertIsArray($user_invite_ref01_list_result);

        $found_item = sdk_select(
            Runner::entity_list_to_data($user_invite_ref01_list_result),
            ["id" => $user_invite_ref01_data["id"]]);
        $this->assertNotEmpty($found_item);

        // UPDATE
        $user_invite_ref01_data_up0_up = [
            "id" => $user_invite_ref01_data["id"],
        ];

        $user_invite_ref01_markdef_up0_name = "initial_role_id";
        $user_invite_ref01_markdef_up0_value = "Mark01-user_invite_ref01_" . $setup["now"];
        $user_invite_ref01_data_up0_up[$user_invite_ref01_markdef_up0_name] = $user_invite_ref01_markdef_up0_value;

        $user_invite_ref01_resdata_up0_result = $user_invite_ref01_ent->update($user_invite_ref01_data_up0_up, null);
        $user_invite_ref01_resdata_up0 = Helpers::to_map($user_invite_ref01_resdata_up0_result);
        $this->assertNotNull($user_invite_ref01_resdata_up0);
        $this->assertEquals($user_invite_ref01_resdata_up0["id"], $user_invite_ref01_data_up0_up["id"]);
        $this->assertEquals($user_invite_ref01_resdata_up0[$user_invite_ref01_markdef_up0_name], $user_invite_ref01_markdef_up0_value);

        // LOAD
        $user_invite_ref01_match_dt0 = [
            "id" => $user_invite_ref01_data["id"],
        ];
        $user_invite_ref01_data_dt0_loaded = $user_invite_ref01_ent->load($user_invite_ref01_match_dt0, null);
        $user_invite_ref01_data_dt0_load_result = Helpers::to_map($user_invite_ref01_data_dt0_loaded);
        $this->assertNotNull($user_invite_ref01_data_dt0_load_result);
        $this->assertEquals($user_invite_ref01_data_dt0_load_result["id"], $user_invite_ref01_data["id"]);

        // REMOVE
        $user_invite_ref01_match_rm0 = [
            "id" => $user_invite_ref01_data["id"],
        ];
        $user_invite_ref01_ent->remove($user_invite_ref01_match_rm0, null);

        // LIST
        $user_invite_ref01_match_rt0 = [
            "merchant_id" => $setup["idmap"]["merchant01"],
        ];

        $user_invite_ref01_list_rt0_result = $user_invite_ref01_ent->list($user_invite_ref01_match_rt0, null);
        $this->assertIsArray($user_invite_ref01_list_rt0_result);

        $not_found_item = sdk_select(
            Runner::entity_list_to_data($user_invite_ref01_list_rt0_result),
            ["id" => $user_invite_ref01_data["id"]]);
        $this->assertEmpty($not_found_item);

    }
}

function user_invite_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/user_invite/UserInviteTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = NofrixionSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["user_invite01", "user_invite02", "user_invite03", "merchant01", "merchant02", "merchant03", "userinvite01", "userinvite02", "userinvite03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("NOFRIXION_TEST_USER_INVITE_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "NOFRIXION_TEST_USER_INVITE_ENTID" => $idmap,
        "NOFRIXION_TEST_LIVE" => "FALSE",
        "NOFRIXION_TEST_EXPLAIN" => "FALSE",
        "NOFRIXION_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["NOFRIXION_TEST_USER_INVITE_ENTID"]);
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
