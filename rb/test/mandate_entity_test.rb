# Mandate entity test

require "minitest/autorun"
require "json"
require_relative "../Nofrixion_sdk"
require_relative "runner"

class MandateEntityTest < Minitest::Test
  def test_create_instance
    testsdk = NofrixionSDK.test(nil, nil)
    ent = testsdk.Mandate(nil)
    assert !ent.nil?
  end

  def test_basic_flow
    setup = mandate_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["create", "load"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "mandate." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_MANDATE_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # CREATE
    mandate_ref01_ent = client.Mandate(nil)
    mandate_ref01_data = Helpers.to_map(Vs.getprop(
      Vs.getpath(setup[:data], "new.mandate"), "mandate_ref01"))

    mandate_ref01_data_result = mandate_ref01_ent.create(mandate_ref01_data, nil)
    mandate_ref01_data = Helpers.to_map(mandate_ref01_data_result)
    assert !mandate_ref01_data.nil?
    assert !mandate_ref01_data["id"].nil?

    # LOAD
    mandate_ref01_match_dt0 = {
      "id" => mandate_ref01_data["id"],
    }
    mandate_ref01_data_dt0_loaded = mandate_ref01_ent.load(mandate_ref01_match_dt0, nil)
    mandate_ref01_data_dt0_load_result = Helpers.to_map(mandate_ref01_data_dt0_loaded)
    assert !mandate_ref01_data_dt0_load_result.nil?
    assert_equal mandate_ref01_data_dt0_load_result["id"], mandate_ref01_data["id"]

  end
end

def mandate_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "mandate", "MandateTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = NofrixionSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["mandate01", "mandate02", "mandate03"],
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
  entid_env_raw = ENV["NOFRIXION_TEST_MANDATE_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "NOFRIXION_TEST_MANDATE_ENTID" => idmap,
    "NOFRIXION_TEST_LIVE" => "FALSE",
    "NOFRIXION_TEST_EXPLAIN" => "FALSE",
    "NOFRIXION_APIKEY" => "NONE",
  })

  idmap_resolved = Helpers.to_map(
    env["NOFRIXION_TEST_MANDATE_ENTID"])
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
