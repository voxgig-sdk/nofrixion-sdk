# Role entity test

require "minitest/autorun"
require "json"
require_relative "../Nofrixion_sdk"
require_relative "runner"

class RoleEntityTest < Minitest::Test
  def test_create_instance
    testsdk = NofrixionSDK.test(nil, nil)
    ent = testsdk.Role(nil)
    assert !ent.nil?
  end

  def test_basic_flow
    setup = role_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["create"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "role." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_ROLE_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # CREATE
    role_ref01_ent = client.Role(nil)
    role_ref01_data = Helpers.to_map(Vs.getprop(
      Vs.getpath(setup[:data], "new.role"), "role_ref01"))
    role_ref01_data["merchant_id"] = setup[:idmap]["merchant01"]

    role_ref01_data_result = role_ref01_ent.create(role_ref01_data, nil)
    role_ref01_data = Helpers.to_map(role_ref01_data_result)
    assert !role_ref01_data.nil?

  end
end

def role_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "role", "RoleTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = NofrixionSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["role01", "role02", "role03", "merchant01", "merchant02", "merchant03"],
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
  entid_env_raw = ENV["NOFRIXION_TEST_ROLE_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "NOFRIXION_TEST_ROLE_ENTID" => idmap,
    "NOFRIXION_TEST_LIVE" => "FALSE",
    "NOFRIXION_TEST_EXPLAIN" => "FALSE",
    "NOFRIXION_APIKEY" => "NONE",
  })

  idmap_resolved = Helpers.to_map(
    env["NOFRIXION_TEST_ROLE_ENTID"])
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
