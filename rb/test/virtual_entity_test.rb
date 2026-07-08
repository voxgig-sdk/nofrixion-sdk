# Virtual entity test

require "minitest/autorun"
require "json"
require_relative "../Nofrixion_sdk"
require_relative "runner"

class VirtualEntityTest < Minitest::Test
  def test_create_instance
    testsdk = NofrixionSDK.test(nil, nil)
    ent = testsdk.Virtual(nil)
    assert !ent.nil?
  end

  def test_basic_flow
    setup = virtual_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["create", "update"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "virtual." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_VIRTUAL_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # CREATE
    virtual_ref01_ent = client.Virtual(nil)
    virtual_ref01_data = Helpers.to_map(Vs.getprop(
      Vs.getpath(setup[:data], "new.virtual"), "virtual_ref01"))
    virtual_ref01_data["account_id"] = setup[:idmap]["account01"]

    virtual_ref01_data_result = virtual_ref01_ent.create(virtual_ref01_data, nil)
    virtual_ref01_data = Helpers.to_map(virtual_ref01_data_result)
    assert !virtual_ref01_data.nil?
    assert !virtual_ref01_data["id"].nil?

    # UPDATE
    virtual_ref01_data_up0_up = {
      "id" => virtual_ref01_data["id"],
      "account_id" => setup[:idmap]["account_id"],
    }

    virtual_ref01_markdef_up0_name = "account_name"
    virtual_ref01_markdef_up0_value = "Mark01-virtual_ref01_#{setup[:now]}"
    virtual_ref01_data_up0_up[virtual_ref01_markdef_up0_name] = virtual_ref01_markdef_up0_value

    virtual_ref01_resdata_up0_result = virtual_ref01_ent.update(virtual_ref01_data_up0_up, nil)
    virtual_ref01_resdata_up0 = Helpers.to_map(virtual_ref01_resdata_up0_result)
    assert !virtual_ref01_resdata_up0.nil?
    assert_equal virtual_ref01_resdata_up0["id"], virtual_ref01_data_up0_up["id"]
    assert_equal virtual_ref01_resdata_up0[virtual_ref01_markdef_up0_name], virtual_ref01_markdef_up0_value

  end
end

def virtual_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "virtual", "VirtualTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = NofrixionSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["virtual01", "virtual02", "virtual03", "account01", "account02", "account03"],
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
  entid_env_raw = ENV["NOFRIXION_TEST_VIRTUAL_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "NOFRIXION_TEST_VIRTUAL_ENTID" => idmap,
    "NOFRIXION_TEST_LIVE" => "FALSE",
    "NOFRIXION_TEST_EXPLAIN" => "FALSE",
    "NOFRIXION_APIKEY" => "NONE",
  })

  idmap_resolved = Helpers.to_map(
    env["NOFRIXION_TEST_VIRTUAL_ENTID"])
  if idmap_resolved.nil?
    idmap_resolved = Helpers.to_map(idmap)
  end
  if idmap_resolved["account_id"].nil?
    idmap_resolved["account_id"] = idmap_resolved["account01"]
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
