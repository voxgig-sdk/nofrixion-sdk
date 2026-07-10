# Payrun entity test

require "minitest/autorun"
require "json"
require_relative "../Nofrixion_sdk"
require_relative "runner"

class PayrunEntityTest < Minitest::Test
  def test_create_instance
    testsdk = NofrixionSDK.test(nil, nil)
    ent = testsdk.Payrun(nil)
    assert !ent.nil?
  end

  def test_basic_flow
    setup = payrun_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["create", "list", "update", "load", "remove"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "payrun." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_PAYRUN_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # CREATE
    payrun_ref01_ent = client.Payrun(nil)
    payrun_ref01_data = Helpers.to_map(Vs.getprop(
      Vs.getpath(setup[:data], "new.payrun"), "payrun_ref01"))
    payrun_ref01_data["merchant_i_d"] = setup[:idmap]["merchant_i_d01"]

    payrun_ref01_data_result = payrun_ref01_ent.create(payrun_ref01_data, nil)
    payrun_ref01_data = Helpers.to_map(payrun_ref01_data_result)
    assert !payrun_ref01_data.nil?
    assert !payrun_ref01_data["id"].nil?

    # LIST
    payrun_ref01_match = {}

    payrun_ref01_list_result = payrun_ref01_ent.list(payrun_ref01_match, nil)
    assert payrun_ref01_list_result.is_a?(Array)

    found_item = Vs.select(
      Runner.entity_list_to_data(payrun_ref01_list_result),
      { "id" => payrun_ref01_data["id"] })
    assert !Vs.isempty(found_item)

    # UPDATE
    payrun_ref01_data_up0_up = {
      "id" => payrun_ref01_data["id"],
    }

    payrun_ref01_markdef_up0_name = "authorisation_date"
    payrun_ref01_markdef_up0_value = "Mark01-payrun_ref01_#{setup[:now]}"
    payrun_ref01_data_up0_up[payrun_ref01_markdef_up0_name] = payrun_ref01_markdef_up0_value

    payrun_ref01_resdata_up0_result = payrun_ref01_ent.update(payrun_ref01_data_up0_up, nil)
    payrun_ref01_resdata_up0 = Helpers.to_map(payrun_ref01_resdata_up0_result)
    assert !payrun_ref01_resdata_up0.nil?
    assert_equal payrun_ref01_resdata_up0["id"], payrun_ref01_data_up0_up["id"]
    assert_equal payrun_ref01_resdata_up0[payrun_ref01_markdef_up0_name], payrun_ref01_markdef_up0_value

    # LOAD
    payrun_ref01_match_dt0 = {
      "id" => payrun_ref01_data["id"],
    }
    payrun_ref01_data_dt0_loaded = payrun_ref01_ent.load(payrun_ref01_match_dt0, nil)
    payrun_ref01_data_dt0_load_result = Helpers.to_map(payrun_ref01_data_dt0_loaded)
    assert !payrun_ref01_data_dt0_load_result.nil?
    assert_equal payrun_ref01_data_dt0_load_result["id"], payrun_ref01_data["id"]

    # REMOVE
    payrun_ref01_match_rm0 = {
      "id" => payrun_ref01_data["id"],
    }
    payrun_ref01_ent.remove(payrun_ref01_match_rm0, nil)

    # LIST
    payrun_ref01_match_rt0 = {}

    payrun_ref01_list_rt0_result = payrun_ref01_ent.list(payrun_ref01_match_rt0, nil)
    assert payrun_ref01_list_rt0_result.is_a?(Array)

    not_found_item = Vs.select(
      Runner.entity_list_to_data(payrun_ref01_list_rt0_result),
      { "id" => payrun_ref01_data["id"] })
    assert Vs.isempty(not_found_item)

  end
end

def payrun_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "payrun", "PayrunTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = NofrixionSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["payrun01", "payrun02", "payrun03", "merchant_i_d01"],
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
  entid_env_raw = ENV["NOFRIXION_TEST_PAYRUN_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "NOFRIXION_TEST_PAYRUN_ENTID" => idmap,
    "NOFRIXION_TEST_LIVE" => "FALSE",
    "NOFRIXION_TEST_EXPLAIN" => "FALSE",
    "NOFRIXION_APIKEY" => "NONE",
  })

  idmap_resolved = Helpers.to_map(
    env["NOFRIXION_TEST_PAYRUN_ENTID"])
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
