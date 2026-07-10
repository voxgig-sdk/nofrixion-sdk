# MerchantPaymentRequestTemplate entity test

require "minitest/autorun"
require "json"
require_relative "../Nofrixion_sdk"
require_relative "runner"

class MerchantPaymentRequestTemplateEntityTest < Minitest::Test
  def test_create_instance
    testsdk = NofrixionSDK.test(nil, nil)
    ent = testsdk.MerchantPaymentRequestTemplate(nil)
    assert !ent.nil?
  end

  def test_basic_flow
    setup = merchant_payment_request_template_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["list", "update", "load", "remove"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "merchant_payment_request_template." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_MERCHANT_PAYMENT_REQUEST_TEMPLATE_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # Bootstrap entity data from existing test data.
    merchant_payment_request_template_ref01_data_raw = Vs.items(Helpers.to_map(
      Vs.getpath(setup[:data], "existing.merchant_payment_request_template")))
    merchant_payment_request_template_ref01_data = nil
    if merchant_payment_request_template_ref01_data_raw.length > 0
      merchant_payment_request_template_ref01_data = Helpers.to_map(merchant_payment_request_template_ref01_data_raw[0][1])
    end

    # LIST
    merchant_payment_request_template_ref01_ent = client.MerchantPaymentRequestTemplate(nil)
    merchant_payment_request_template_ref01_match = {
      "merchant_id" => setup[:idmap]["merchant01"],
    }

    merchant_payment_request_template_ref01_list_result = merchant_payment_request_template_ref01_ent.list(merchant_payment_request_template_ref01_match, nil)
    assert merchant_payment_request_template_ref01_list_result.is_a?(Array)

    # UPDATE
    merchant_payment_request_template_ref01_data_up0_up = {
      "id" => merchant_payment_request_template_ref01_data["id"],
      "paymentrequest_id" => setup[:idmap]["paymentrequest_id"],
    }

    merchant_payment_request_template_ref01_markdef_up0_name = "description"
    merchant_payment_request_template_ref01_markdef_up0_value = "Mark01-merchant_payment_request_template_ref01_#{setup[:now]}"
    merchant_payment_request_template_ref01_data_up0_up[merchant_payment_request_template_ref01_markdef_up0_name] = merchant_payment_request_template_ref01_markdef_up0_value

    merchant_payment_request_template_ref01_resdata_up0_result = merchant_payment_request_template_ref01_ent.update(merchant_payment_request_template_ref01_data_up0_up, nil)
    merchant_payment_request_template_ref01_resdata_up0 = Helpers.to_map(merchant_payment_request_template_ref01_resdata_up0_result)
    assert !merchant_payment_request_template_ref01_resdata_up0.nil?
    assert_equal merchant_payment_request_template_ref01_resdata_up0["id"], merchant_payment_request_template_ref01_data_up0_up["id"]
    assert_equal merchant_payment_request_template_ref01_resdata_up0[merchant_payment_request_template_ref01_markdef_up0_name], merchant_payment_request_template_ref01_markdef_up0_value

    # LOAD
    merchant_payment_request_template_ref01_match_dt0 = {
      "id" => merchant_payment_request_template_ref01_data["id"],
    }
    merchant_payment_request_template_ref01_data_dt0_loaded = merchant_payment_request_template_ref01_ent.load(merchant_payment_request_template_ref01_match_dt0, nil)
    merchant_payment_request_template_ref01_data_dt0_load_result = Helpers.to_map(merchant_payment_request_template_ref01_data_dt0_loaded)
    assert !merchant_payment_request_template_ref01_data_dt0_load_result.nil?
    assert_equal merchant_payment_request_template_ref01_data_dt0_load_result["id"], merchant_payment_request_template_ref01_data["id"]

    # REMOVE
    merchant_payment_request_template_ref01_match_rm0 = {
      "id" => merchant_payment_request_template_ref01_data["id"],
    }
    merchant_payment_request_template_ref01_ent.remove(merchant_payment_request_template_ref01_match_rm0, nil)

    # LIST
    merchant_payment_request_template_ref01_match_rt0 = {
      "merchant_id" => setup[:idmap]["merchant01"],
    }

    merchant_payment_request_template_ref01_list_rt0_result = merchant_payment_request_template_ref01_ent.list(merchant_payment_request_template_ref01_match_rt0, nil)
    assert merchant_payment_request_template_ref01_list_rt0_result.is_a?(Array)

  end
end

def merchant_payment_request_template_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "merchant_payment_request_template", "MerchantPaymentRequestTemplateTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = NofrixionSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["merchant_payment_request_template01", "merchant_payment_request_template02", "merchant_payment_request_template03", "paymentrequest01", "paymentrequest02", "paymentrequest03", "merchant01"],
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
  entid_env_raw = ENV["NOFRIXION_TEST_MERCHANT_PAYMENT_REQUEST_TEMPLATE_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "NOFRIXION_TEST_MERCHANT_PAYMENT_REQUEST_TEMPLATE_ENTID" => idmap,
    "NOFRIXION_TEST_LIVE" => "FALSE",
    "NOFRIXION_TEST_EXPLAIN" => "FALSE",
    "NOFRIXION_APIKEY" => "NONE",
  })

  idmap_resolved = Helpers.to_map(
    env["NOFRIXION_TEST_MERCHANT_PAYMENT_REQUEST_TEMPLATE_ENTID"])
  if idmap_resolved.nil?
    idmap_resolved = Helpers.to_map(idmap)
  end
  if idmap_resolved["paymentrequest_id"].nil?
    idmap_resolved["paymentrequest_id"] = idmap_resolved["paymentrequest01"]
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
