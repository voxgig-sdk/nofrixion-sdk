# NoFrixionMoneyMoovModelsWebhook entity test

require "minitest/autorun"
require "json"
require_relative "../Nofrixion_sdk"
require_relative "runner"

class NoFrixionMoneyMoovModelsWebhookEntityTest < Minitest::Test
  def test_create_instance
    testsdk = NofrixionSDK.test(nil, nil)
    ent = testsdk.NoFrixionMoneyMoovModelsWebhook(nil)
    assert !ent.nil?
  end

  def test_basic_flow
    setup = no_frixion_money_moov_models_webhook_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["create", "list", "update", "load"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "no_frixion_money_moov_models_webhook." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_WEBHOOK_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # CREATE
    no_frixion_money_moov_models_webhook_ref01_ent = client.NoFrixionMoneyMoovModelsWebhook(nil)
    no_frixion_money_moov_models_webhook_ref01_data = Helpers.to_map(Vs.getprop(
      Vs.getpath(setup[:data], "new.no_frixion_money_moov_models_webhook"), "no_frixion_money_moov_models_webhook_ref01"))
    no_frixion_money_moov_models_webhook_ref01_data["merchant_id"] = setup[:idmap]["merchant01"]

    no_frixion_money_moov_models_webhook_ref01_data_result = no_frixion_money_moov_models_webhook_ref01_ent.create(no_frixion_money_moov_models_webhook_ref01_data, nil)
    no_frixion_money_moov_models_webhook_ref01_data = Helpers.to_map(no_frixion_money_moov_models_webhook_ref01_data_result)
    assert !no_frixion_money_moov_models_webhook_ref01_data.nil?
    assert !no_frixion_money_moov_models_webhook_ref01_data["id"].nil?

    # LIST
    no_frixion_money_moov_models_webhook_ref01_match = {
      "merchant_id" => setup[:idmap]["merchant01"],
    }

    no_frixion_money_moov_models_webhook_ref01_list_result = no_frixion_money_moov_models_webhook_ref01_ent.list(no_frixion_money_moov_models_webhook_ref01_match, nil)
    assert no_frixion_money_moov_models_webhook_ref01_list_result.is_a?(Array)

    found_item = Vs.select(
      Runner.entity_list_to_data(no_frixion_money_moov_models_webhook_ref01_list_result),
      { "id" => no_frixion_money_moov_models_webhook_ref01_data["id"] })
    assert !Vs.isempty(found_item)

    # UPDATE
    no_frixion_money_moov_models_webhook_ref01_data_up0_up = {
      "id" => no_frixion_money_moov_models_webhook_ref01_data["id"],
    }

    no_frixion_money_moov_models_webhook_ref01_markdef_up0_name = "destination_url"
    no_frixion_money_moov_models_webhook_ref01_markdef_up0_value = "Mark01-no_frixion_money_moov_models_webhook_ref01_#{setup[:now]}"
    no_frixion_money_moov_models_webhook_ref01_data_up0_up[no_frixion_money_moov_models_webhook_ref01_markdef_up0_name] = no_frixion_money_moov_models_webhook_ref01_markdef_up0_value

    no_frixion_money_moov_models_webhook_ref01_resdata_up0_result = no_frixion_money_moov_models_webhook_ref01_ent.update(no_frixion_money_moov_models_webhook_ref01_data_up0_up, nil)
    no_frixion_money_moov_models_webhook_ref01_resdata_up0 = Helpers.to_map(no_frixion_money_moov_models_webhook_ref01_resdata_up0_result)
    assert !no_frixion_money_moov_models_webhook_ref01_resdata_up0.nil?
    assert_equal no_frixion_money_moov_models_webhook_ref01_resdata_up0["id"], no_frixion_money_moov_models_webhook_ref01_data_up0_up["id"]
    assert_equal no_frixion_money_moov_models_webhook_ref01_resdata_up0[no_frixion_money_moov_models_webhook_ref01_markdef_up0_name], no_frixion_money_moov_models_webhook_ref01_markdef_up0_value

    # LOAD
    no_frixion_money_moov_models_webhook_ref01_match_dt0 = {
      "id" => no_frixion_money_moov_models_webhook_ref01_data["id"],
    }
    no_frixion_money_moov_models_webhook_ref01_data_dt0_loaded = no_frixion_money_moov_models_webhook_ref01_ent.load(no_frixion_money_moov_models_webhook_ref01_match_dt0, nil)
    no_frixion_money_moov_models_webhook_ref01_data_dt0_load_result = Helpers.to_map(no_frixion_money_moov_models_webhook_ref01_data_dt0_loaded)
    assert !no_frixion_money_moov_models_webhook_ref01_data_dt0_load_result.nil?
    assert_equal no_frixion_money_moov_models_webhook_ref01_data_dt0_load_result["id"], no_frixion_money_moov_models_webhook_ref01_data["id"]

  end
end

def no_frixion_money_moov_models_webhook_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "no_frixion_money_moov_models_webhook", "NoFrixionMoneyMoovModelsWebhookTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = NofrixionSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["no_frixion_money_moov_models_webhook01", "no_frixion_money_moov_models_webhook02", "no_frixion_money_moov_models_webhook03", "merchant01", "merchant02", "merchant03"],
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
  entid_env_raw = ENV["NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_WEBHOOK_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_WEBHOOK_ENTID" => idmap,
    "NOFRIXION_TEST_LIVE" => "FALSE",
    "NOFRIXION_TEST_EXPLAIN" => "FALSE",
    "NOFRIXION_APIKEY" => "NONE",
  })

  idmap_resolved = Helpers.to_map(
    env["NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_WEBHOOK_ENTID"])
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
