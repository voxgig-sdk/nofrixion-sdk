# Payout entity test

require "minitest/autorun"
require "json"
require_relative "../Nofrixion_sdk"
require_relative "runner"

class PayoutEntityTest < Minitest::Test
  def test_create_instance
    testsdk = NofrixionSDK.test(nil, nil)
    ent = testsdk.Payout(nil)
    assert !ent.nil?
  end

  def test_basic_flow
    setup = payout_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["create", "update", "load", "remove"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "payout." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_PAYOUT_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # CREATE
    payout_ref01_ent = client.Payout(nil)
    payout_ref01_data = Helpers.to_map(Vs.getprop(
      Vs.getpath(setup[:data], "new.payout"), "payout_ref01"))
    payout_ref01_data["destination"] = setup[:idmap]["destination01"]
    payout_ref01_data["source"] = setup[:idmap]["source01"]

    payout_ref01_data_result = payout_ref01_ent.create(payout_ref01_data, nil)
    payout_ref01_data = Helpers.to_map(payout_ref01_data_result)
    assert !payout_ref01_data.nil?
    assert !payout_ref01_data["id"].nil?

    # UPDATE
    payout_ref01_data_up0_up = {
      "id" => payout_ref01_data["id"],
    }

    payout_ref01_markdef_up0_name = "account_id"
    payout_ref01_markdef_up0_value = "Mark01-payout_ref01_#{setup[:now]}"
    payout_ref01_data_up0_up[payout_ref01_markdef_up0_name] = payout_ref01_markdef_up0_value

    payout_ref01_resdata_up0_result = payout_ref01_ent.update(payout_ref01_data_up0_up, nil)
    payout_ref01_resdata_up0 = Helpers.to_map(payout_ref01_resdata_up0_result)
    assert !payout_ref01_resdata_up0.nil?
    assert_equal payout_ref01_resdata_up0["id"], payout_ref01_data_up0_up["id"]
    assert_equal payout_ref01_resdata_up0[payout_ref01_markdef_up0_name], payout_ref01_markdef_up0_value

    # LOAD
    payout_ref01_match_dt0 = {
      "id" => payout_ref01_data["id"],
    }
    payout_ref01_data_dt0_loaded = payout_ref01_ent.load(payout_ref01_match_dt0, nil)
    payout_ref01_data_dt0_load_result = Helpers.to_map(payout_ref01_data_dt0_loaded)
    assert !payout_ref01_data_dt0_load_result.nil?
    assert_equal payout_ref01_data_dt0_load_result["id"], payout_ref01_data["id"]

    # REMOVE
    payout_ref01_match_rm0 = {
      "id" => payout_ref01_data["id"],
    }
    payout_ref01_ent.remove(payout_ref01_match_rm0, nil)

  end
end

def payout_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "payout", "PayoutTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = NofrixionSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["payout01", "payout02", "payout03", "fxquote01", "fxquote02", "fxquote03", "destination01", "source01"],
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
  entid_env_raw = ENV["NOFRIXION_TEST_PAYOUT_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "NOFRIXION_TEST_PAYOUT_ENTID" => idmap,
    "NOFRIXION_TEST_LIVE" => "FALSE",
    "NOFRIXION_TEST_EXPLAIN" => "FALSE",
    "NOFRIXION_APIKEY" => "NONE",
  })

  idmap_resolved = Helpers.to_map(
    env["NOFRIXION_TEST_PAYOUT_ENTID"])
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
