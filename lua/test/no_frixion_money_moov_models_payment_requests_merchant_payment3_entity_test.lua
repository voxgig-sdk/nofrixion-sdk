-- NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3 entity test

local json = require("dkjson")
local vs = require("utility.struct.struct")
local sdk = require("nofrixion_sdk")
local helpers = require("core.helpers")
local runner = require("test.runner")

local _test_dir = debug.getinfo(1, "S").source:match("^@(.+/)")  or "./"

describe("NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity", function()
  it("should create instance", function()
    local testsdk = sdk.test(nil, nil)
    local ent = testsdk:NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3(nil)
    assert.is_not_nil(ent)
  end)

  it("should run basic flow", function()
    local setup = no_frixion_money_moov_models_payment_requests_merchant_payment3_basic_setup(nil)
    -- Per-op sdk-test-control.json skip.
    local _live = setup.live or false
    for _, _op in ipairs({"update"}) do
      local _should_skip, _reason = runner.is_control_skipped("entityOp", "no_frixion_money_moov_models_payment_requests_merchant_payment3." .. _op, _live and "live" or "unit")
      if _should_skip then
        pending(_reason or "skipped via sdk-test-control.json")
        return
      end
    end
    -- The basic flow consumes synthetic IDs from the fixture. In live mode
    -- without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup.synthetic_only then
      pending("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_PAYMENT_REQUESTS_MERCHANT_PAYMENT__ENTID JSON to run live")
      return
    end
    local client = setup.client

    -- Bootstrap entity data from existing test data.
    local no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_data_raw = vs.items(helpers.to_map(
      vs.getpath(setup.data, "existing.no_frixion_money_moov_models_payment_requests_merchant_payment3")))
    local no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_data = nil
    if #no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_data_raw > 0 then
      no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_data = helpers.to_map(no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_data_raw[1][2])
    end

    -- UPDATE
    local no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_ent = client:NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3(nil)
    local no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_data_up0_up = {
      id = no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_data["id"],
      ["paymentrequest_id"] = setup.idmap["paymentrequest_id"],
    }

    local no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_markdef_up0_name = "description"
    local no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_markdef_up0_value = "Mark01-no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_" .. tostring(setup.now)
    no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_data_up0_up[no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_markdef_up0_name] = no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_markdef_up0_value

    local no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_resdata_up0_result, err = no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_ent:update(no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_data_up0_up, nil)
    assert.is_nil(err)
    local no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_resdata_up0 = helpers.to_map(no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_resdata_up0_result)
    assert.is_not_nil(no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_resdata_up0)
    assert.are.equal(no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_resdata_up0["id"], no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_data_up0_up["id"])
    assert.are.equal(no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_resdata_up0[no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_markdef_up0_name], no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_markdef_up0_value)

  end)
end)

function no_frixion_money_moov_models_payment_requests_merchant_payment3_basic_setup(extra)
  runner.load_env_local()

  local entity_data_file = _test_dir .. "../../.sdk/test/entity/no_frixion_money_moov_models_payment_requests_merchant_payment3/NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3TestData.json"
  local f = io.open(entity_data_file, "r")
  if f == nil then
    error("failed to read no_frixion_money_moov_models_payment_requests_merchant_payment3 test data: " .. entity_data_file)
  end
  local entity_data_source = f:read("*a")
  f:close()

  local entity_data = json.decode(entity_data_source)

  local options = {}
  options["entity"] = entity_data["existing"]

  local client = sdk.test(options, extra)

  -- Generate idmap via transform.
  local idmap = vs.transform(
    { "no_frixion_money_moov_models_payment_requests_merchant_payment301", "no_frixion_money_moov_models_payment_requests_merchant_payment302", "no_frixion_money_moov_models_payment_requests_merchant_payment303", "paymentrequest01", "paymentrequest02", "paymentrequest03", "template01", "template02", "template03" },
    {
      ["`$PACK`"] = { "", {
        ["`$KEY`"] = "`$COPY`",
        ["`$VAL`"] = { "`$FORMAT`", "upper", "`$COPY`" },
      }},
    }
  )

  -- Detect ENTID env override before envOverride consumes it. When live
  -- mode is on without a real override, the basic test runs against synthetic
  -- IDs from the fixture and 4xx's. Surface this so the test can skip.
  local entid_env_raw = os.getenv("NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_PAYMENT_REQUESTS_MERCHANT_PAYMENT__ENTID")
  local idmap_overridden = entid_env_raw ~= nil and entid_env_raw:match("^%s*{") ~= nil

  local env = runner.env_override({
    ["NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_PAYMENT_REQUESTS_MERCHANT_PAYMENT__ENTID"] = idmap,
    ["NOFRIXION_TEST_LIVE"] = "FALSE",
    ["NOFRIXION_TEST_EXPLAIN"] = "FALSE",
    ["NOFRIXION_APIKEY"] = "NONE",
  })

  local idmap_resolved = helpers.to_map(
    env["NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_PAYMENT_REQUESTS_MERCHANT_PAYMENT__ENTID"])
  if idmap_resolved == nil then
    idmap_resolved = helpers.to_map(idmap)
  end
  if idmap_resolved["paymentrequest_id"] == nil then
    idmap_resolved["paymentrequest_id"] = idmap_resolved["paymentrequest01"]
  end

  if env["NOFRIXION_TEST_LIVE"] == "TRUE" then
    local merged_opts = vs.merge({
      {
        apikey = env["NOFRIXION_APIKEY"],
      },
      extra or {},
    })
    client = sdk.new(helpers.to_map(merged_opts))
  end

  local live = env["NOFRIXION_TEST_LIVE"] == "TRUE"
  return {
    client = client,
    data = entity_data,
    idmap = idmap_resolved,
    env = env,
    explain = env["NOFRIXION_TEST_EXPLAIN"] == "TRUE",
    live = live,
    synthetic_only = live and not idmap_overridden,
    now = os.time() * 1000,
  }
end
