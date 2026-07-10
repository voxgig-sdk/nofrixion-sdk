# CardCustomerToken entity test

require "minitest/autorun"
require "json"
require_relative "../Nofrixion_sdk"
require_relative "runner"

class CardCustomerTokenEntityTest < Minitest::Test
  def test_create_instance
    testsdk = NofrixionSDK.test(nil, nil)
    ent = testsdk.CardCustomerToken(nil)
    assert !ent.nil?
  end

  def test_basic_flow
    setup = card_customer_token_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["list", "load", "remove"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "card_customer_token." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_CARD_CUSTOMER_TOKEN_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # Bootstrap entity data from existing test data.
    card_customer_token_ref01_data_raw = Vs.items(Helpers.to_map(
      Vs.getpath(setup[:data], "existing.card_customer_token")))
    card_customer_token_ref01_data = nil
    if card_customer_token_ref01_data_raw.length > 0
      card_customer_token_ref01_data = Helpers.to_map(card_customer_token_ref01_data_raw[0][1])
    end

    # LIST
    card_customer_token_ref01_ent = client.CardCustomerToken(nil)
    card_customer_token_ref01_match = {
      "customer_email_address" => setup[:idmap]["customer_email_address01"],
      "merchant_id" => setup[:idmap]["merchant01"],
    }

    card_customer_token_ref01_list_result = card_customer_token_ref01_ent.list(card_customer_token_ref01_match, nil)
    assert card_customer_token_ref01_list_result.is_a?(Array)

    # LOAD
    card_customer_token_ref01_match_dt0 = {
      "id" => card_customer_token_ref01_data["id"],
    }
    card_customer_token_ref01_data_dt0_loaded = card_customer_token_ref01_ent.load(card_customer_token_ref01_match_dt0, nil)
    card_customer_token_ref01_data_dt0_load_result = Helpers.to_map(card_customer_token_ref01_data_dt0_loaded)
    assert !card_customer_token_ref01_data_dt0_load_result.nil?
    assert_equal card_customer_token_ref01_data_dt0_load_result["id"], card_customer_token_ref01_data["id"]

    # REMOVE
    card_customer_token_ref01_match_rm0 = {
      "id" => card_customer_token_ref01_data["id"],
    }
    card_customer_token_ref01_ent.remove(card_customer_token_ref01_match_rm0, nil)

    # LIST
    card_customer_token_ref01_match_rt0 = {
      "customer_email_address" => setup[:idmap]["customer_email_address01"],
      "merchant_id" => setup[:idmap]["merchant01"],
    }

    card_customer_token_ref01_list_rt0_result = card_customer_token_ref01_ent.list(card_customer_token_ref01_match_rt0, nil)
    assert card_customer_token_ref01_list_rt0_result.is_a?(Array)

  end
end

def card_customer_token_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "card_customer_token", "CardCustomerTokenTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = NofrixionSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["card_customer_token01", "card_customer_token02", "card_customer_token03", "removeall01", "removeall02", "removeall03", "customertoken01", "customertoken02", "customertoken03", "customer_email_address01", "merchant01"],
    {
      "`$PACK`" => ["", {
        "`$KEY`" => "`$COPY`",
        "`$VAL`" => ["`$FORMAT`", "upper", "`$COPY`"],
      }],
    }
  )

  # Detect ENTID env override before envOverride consumes it. When live
  # mode is on without a real override, the basic test runs against synthetic
  # IDs from the fixture and 4xx's. Surface this so the test can skip.
  entid_env_raw = ENV["NOFRIXION_TEST_CARD_CUSTOMER_TOKEN_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "NOFRIXION_TEST_CARD_CUSTOMER_TOKEN_ENTID" => idmap,
    "NOFRIXION_TEST_LIVE" => "FALSE",
    "NOFRIXION_TEST_EXPLAIN" => "FALSE",
    "NOFRIXION_APIKEY" => "NONE",
  })

  idmap_resolved = Helpers.to_map(
    env["NOFRIXION_TEST_CARD_CUSTOMER_TOKEN_ENTID"])
  if idmap_resolved.nil?
    idmap_resolved = Helpers.to_map(idmap)
  end

  if env["NOFRIXION_TEST_LIVE"] == "TRUE"
    merged_opts = Vs.merge([
      {
        "apikey" => env["NOFRIXION_APIKEY"],
      },
      extra || {},
    ])
    client = NofrixionSDK.new(Helpers.to_map(merged_opts))
  end

  live = env["NOFRIXION_TEST_LIVE"] == "TRUE"
  {
    client: client,
    data: entity_data,
    idmap: idmap_resolved,
    env: env,
    explain: env["NOFRIXION_TEST_EXPLAIN"] == "TRUE",
    live: live,
    synthetic_only: live && !idmap_overridden,
    now: (Time.now.to_f * 1000).to_i,
  }
end
