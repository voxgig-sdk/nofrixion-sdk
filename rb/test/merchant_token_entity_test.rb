# MerchantToken entity test

require "minitest/autorun"
require "json"
require_relative "../Nofrixion_sdk"
require_relative "runner"

class MerchantTokenEntityTest < Minitest::Test
  def test_create_instance
    testsdk = NofrixionSDK.test(nil, nil)
    ent = testsdk.MerchantToken(nil)
    assert !ent.nil?
  end

  # Feature #4: the entity stream(action, ...) method runs the op pipeline and
  # returns an Enumerator over result items. With the streaming feature active
  # it yields the feature's incremental output; otherwise it falls back to the
  # materialised list so stream always yields.
  def test_stream
    seed = {
      "entity" => {
        "merchant_token" => {
          "s1" => { "id" => "s1" },
          "s2" => { "id" => "s2" },
          "s3" => { "id" => "s3" },
        },
      },
    }

    # Fallback: streaming inactive -> yields the materialised list items.
    base = NofrixionSDK.test(seed, nil)
    seen = base.MerchantToken(nil).stream("list", nil, nil).to_a
    assert_equal 3, seen.length

    # Inbound: streaming active -> yields each item from the feature.
    cfg = NofrixionConfig.make_config
    if cfg["feature"].is_a?(Hash) && cfg["feature"].key?("streaming")
      sdk = NofrixionSDK.test(seed, { "feature" => { "streaming" => { "active" => true } } })
      got = []
      sdk.MerchantToken(nil).stream("list", nil, nil).each do |item|
        if item.is_a?(Array)
          got.concat(item)
        else
          got << item
        end
      end
      assert_equal 3, got.length
    end
  end

  def test_basic_flow
    setup = merchant_token_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["create", "list", "update", "load"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "merchant_token." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_MERCHANT_TOKEN_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # CREATE
    merchant_token_ref01_ent = client.MerchantToken(nil)
    merchant_token_ref01_data = Helpers.to_map(Vs.getprop(
      Vs.getpath(setup[:data], "new.merchant_token"), "merchant_token_ref01"))
    merchant_token_ref01_data["merchant_id"] = setup[:idmap]["merchant01"]

    merchant_token_ref01_data_result = merchant_token_ref01_ent.create(merchant_token_ref01_data, nil)
    merchant_token_ref01_data = Helpers.to_map(merchant_token_ref01_data_result)
    assert !merchant_token_ref01_data.nil?
    assert !merchant_token_ref01_data["id"].nil?

    # LIST
    merchant_token_ref01_match = {
      "merchant_id" => setup[:idmap]["merchant01"],
    }

    merchant_token_ref01_list_result = merchant_token_ref01_ent.list(merchant_token_ref01_match, nil)
    assert merchant_token_ref01_list_result.is_a?(Array)

    found_item = Vs.select(
      Runner.entity_list_to_data(merchant_token_ref01_list_result),
      { "id" => merchant_token_ref01_data["id"] })
    assert !Vs.isempty(found_item)

    # UPDATE
    merchant_token_ref01_data_up0_up = {
      "id" => merchant_token_ref01_data["id"],
    }

    merchant_token_ref01_markdef_up0_name = "description"
    merchant_token_ref01_markdef_up0_value = "Mark01-merchant_token_ref01_#{setup[:now]}"
    merchant_token_ref01_data_up0_up[merchant_token_ref01_markdef_up0_name] = merchant_token_ref01_markdef_up0_value

    merchant_token_ref01_resdata_up0_result = merchant_token_ref01_ent.update(merchant_token_ref01_data_up0_up, nil)
    merchant_token_ref01_resdata_up0 = Helpers.to_map(merchant_token_ref01_resdata_up0_result)
    assert !merchant_token_ref01_resdata_up0.nil?
    assert_equal merchant_token_ref01_resdata_up0["id"], merchant_token_ref01_data_up0_up["id"]
    assert_equal merchant_token_ref01_resdata_up0[merchant_token_ref01_markdef_up0_name], merchant_token_ref01_markdef_up0_value

    # LOAD
    merchant_token_ref01_match_dt0 = {
      "id" => merchant_token_ref01_data["id"],
    }
    merchant_token_ref01_data_dt0_loaded = merchant_token_ref01_ent.load(merchant_token_ref01_match_dt0, nil)
    merchant_token_ref01_data_dt0_load_result = Helpers.to_map(merchant_token_ref01_data_dt0_loaded)
    assert !merchant_token_ref01_data_dt0_load_result.nil?
    assert_equal merchant_token_ref01_data_dt0_load_result["id"], merchant_token_ref01_data["id"]

  end
end

def merchant_token_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "merchant_token", "MerchantTokenTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = NofrixionSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["merchant_token01", "merchant_token02", "merchant_token03", "merchant01", "merchant02", "merchant03"],
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
  entid_env_raw = ENV["NOFRIXION_TEST_MERCHANT_TOKEN_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "NOFRIXION_TEST_MERCHANT_TOKEN_ENTID" => idmap,
    "NOFRIXION_TEST_LIVE" => "FALSE",
    "NOFRIXION_TEST_EXPLAIN" => "FALSE",
    "NOFRIXION_APIKEY" => "NONE",
  })

  idmap_resolved = Helpers.to_map(
    env["NOFRIXION_TEST_MERCHANT_TOKEN_ENTID"])
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
