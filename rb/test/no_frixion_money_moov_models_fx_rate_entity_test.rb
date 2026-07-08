# NoFrixionMoneyMoovModelsFxRate entity test

require "minitest/autorun"
require "json"
require_relative "../Nofrixion_sdk"
require_relative "runner"

class NoFrixionMoneyMoovModelsFxRateEntityTest < Minitest::Test
  def test_create_instance
    testsdk = NofrixionSDK.test(nil, nil)
    ent = testsdk.NoFrixionMoneyMoovModelsFxRate(nil)
    assert !ent.nil?
  end

  def test_basic_flow
    setup = no_frixion_money_moov_models_fx_rate_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["list", "load"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "no_frixion_money_moov_models_fx_rate." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_FX_RATE_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # Bootstrap entity data from existing test data.
    no_frixion_money_moov_models_fx_rate_ref01_data_raw = Vs.items(Helpers.to_map(
      Vs.getpath(setup[:data], "existing.no_frixion_money_moov_models_fx_rate")))
    no_frixion_money_moov_models_fx_rate_ref01_data = nil
    if no_frixion_money_moov_models_fx_rate_ref01_data_raw.length > 0
      no_frixion_money_moov_models_fx_rate_ref01_data = Helpers.to_map(no_frixion_money_moov_models_fx_rate_ref01_data_raw[0][1])
    end

    # LIST
    no_frixion_money_moov_models_fx_rate_ref01_ent = client.NoFrixionMoneyMoovModelsFxRate(nil)
    no_frixion_money_moov_models_fx_rate_ref01_match = {
      "destination" => setup[:idmap]["destination01"],
      "source" => setup[:idmap]["source01"],
    }

    no_frixion_money_moov_models_fx_rate_ref01_list_result = no_frixion_money_moov_models_fx_rate_ref01_ent.list(no_frixion_money_moov_models_fx_rate_ref01_match, nil)
    assert no_frixion_money_moov_models_fx_rate_ref01_list_result.is_a?(Array)

    # LOAD
    no_frixion_money_moov_models_fx_rate_ref01_match_dt0 = {}
    no_frixion_money_moov_models_fx_rate_ref01_data_dt0_loaded = no_frixion_money_moov_models_fx_rate_ref01_ent.load(no_frixion_money_moov_models_fx_rate_ref01_match_dt0, nil)
    assert !no_frixion_money_moov_models_fx_rate_ref01_data_dt0_loaded.nil?

  end
end

def no_frixion_money_moov_models_fx_rate_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "no_frixion_money_moov_models_fx_rate", "NoFrixionMoneyMoovModelsFxRateTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = NofrixionSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["no_frixion_money_moov_models_fx_rate01", "no_frixion_money_moov_models_fx_rate02", "no_frixion_money_moov_models_fx_rate03", "fxallheldrate01", "fxallheldrate02", "fxallheldrate03", "fxheldrate01", "fxheldrate02", "fxheldrate03", "destination01", "source01"],
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
  entid_env_raw = ENV["NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_FX_RATE_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_FX_RATE_ENTID" => idmap,
    "NOFRIXION_TEST_LIVE" => "FALSE",
    "NOFRIXION_TEST_EXPLAIN" => "FALSE",
    "NOFRIXION_APIKEY" => "NONE",
  })

  idmap_resolved = Helpers.to_map(
    env["NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_FX_RATE_ENTID"])
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
