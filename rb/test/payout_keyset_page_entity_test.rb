# PayoutKeysetPage entity test

require "minitest/autorun"
require "json"
require_relative "../Nofrixion_sdk"
require_relative "runner"

class PayoutKeysetPageEntityTest < Minitest::Test
  def test_create_instance
    testsdk = NofrixionSDK.test(nil, nil)
    ent = testsdk.PayoutKeysetPage(nil)
    assert !ent.nil?
  end

  # Feature #4: the entity stream(action, ...) method runs the op pipeline and
  # returns an Enumerator over result items. With the streaming feature active
  # it yields the feature's incremental output; otherwise it falls back to the
  # materialised list so stream always yields.
  def test_stream
    seed = {
      "entity" => {
        "payout_keyset_page" => {
          "s1" => { "id" => "s1" },
          "s2" => { "id" => "s2" },
          "s3" => { "id" => "s3" },
        },
      },
    }

    # Fallback: streaming inactive -> yields the materialised list items.
    base = NofrixionSDK.test(seed, nil)
    seen = base.PayoutKeysetPage(nil).stream("list", nil, nil).to_a
    assert_equal 3, seen.length

    # Inbound: streaming active -> yields each item from the feature.
    cfg = NofrixionConfig.make_config
    if cfg["feature"].is_a?(Hash) && cfg["feature"].key?("streaming")
      sdk = NofrixionSDK.test(seed, { "feature" => { "streaming" => { "active" => true } } })
      got = []
      sdk.PayoutKeysetPage(nil).stream("list", nil, nil).each do |item|
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
    setup = payout_keyset_page_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["list"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "payout_keyset_page." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_PAYOUT_KEYSET_PAGE_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # Bootstrap entity data from existing test data.
    payout_keyset_page_ref01_data_raw = Vs.items(Helpers.to_map(
      Vs.getpath(setup[:data], "existing.payout_keyset_page")))
    payout_keyset_page_ref01_data = nil
    if payout_keyset_page_ref01_data_raw.length > 0
      payout_keyset_page_ref01_data = Helpers.to_map(payout_keyset_page_ref01_data_raw[0][1])
    end

    # LIST
    payout_keyset_page_ref01_ent = client.PayoutKeysetPage(nil)
    payout_keyset_page_ref01_match = {
      "merchant_id" => setup[:idmap]["merchant01"],
    }

    payout_keyset_page_ref01_list_result = payout_keyset_page_ref01_ent.list(payout_keyset_page_ref01_match, nil)
    assert payout_keyset_page_ref01_list_result.is_a?(Array)

  end
end

def payout_keyset_page_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "payout_keyset_page", "PayoutKeysetPageTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = NofrixionSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["payout_keyset_page01", "payout_keyset_page02", "payout_keyset_page03", "account01", "account02", "account03", "merchant01", "merchant02", "merchant03", "payout01", "payout02", "payout03"],
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
  entid_env_raw = ENV["NOFRIXION_TEST_PAYOUT_KEYSET_PAGE_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "NOFRIXION_TEST_PAYOUT_KEYSET_PAGE_ENTID" => idmap,
    "NOFRIXION_TEST_LIVE" => "FALSE",
    "NOFRIXION_TEST_EXPLAIN" => "FALSE",
    "NOFRIXION_APIKEY" => "NONE",
  })

  idmap_resolved = Helpers.to_map(
    env["NOFRIXION_TEST_PAYOUT_KEYSET_PAGE_ENTID"])
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
