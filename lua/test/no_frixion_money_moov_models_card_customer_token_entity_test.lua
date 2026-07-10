-- NoFrixionMoneyMoovModelsCardCustomerToken entity test

local json = require("dkjson")
local vs = require("utility.struct.struct")
local sdk = require("nofrixion_sdk")
local helpers = require("core.helpers")
local runner = require("test.runner")

local _test_dir = debug.getinfo(1, "S").source:match("^@(.+/)")  or "./"

describe("NoFrixionMoneyMoovModelsCardCustomerTokenEntity", function()
  it("should create instance", function()
    local testsdk = sdk.test(nil, nil)
    local ent = testsdk:NoFrixionMoneyMoovModelsCardCustomerToken(nil)
    assert.is_not_nil(ent)
  end)

  it("should run basic flow", function()
    local setup = no_frixion_money_moov_models_card_customer_token_basic_setup(nil)
    -- Per-op sdk-test-control.json skip.
    local _live = setup.live or false
    for _, _op in ipairs({"list", "load", "remove"}) do
      local _should_skip, _reason = runner.is_control_skipped("entityOp", "no_frixion_money_moov_models_card_customer_token." .. _op, _live and "live" or "unit")
      if _should_skip then
        pending(_reason or "skipped via sdk-test-control.json")
        return
      end
    end
    -- The basic flow consumes synthetic IDs from the fixture. In live mode
    -- without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup.synthetic_only then
      pending("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_CARD_CUSTOMER_TOKEN_ENTID JSON to run live")
      return
    end
    local client = setup.client

    -- Bootstrap entity data from existing test data.
    local no_frixion_money_moov_models_card_customer_token_ref01_data_raw = vs.items(helpers.to_map(
      vs.getpath(setup.data, "existing.no_frixion_money_moov_models_card_customer_token")))
    local no_frixion_money_moov_models_card_customer_token_ref01_data = nil
    if #no_frixion_money_moov_models_card_customer_token_ref01_data_raw > 0 then
      no_frixion_money_moov_models_card_customer_token_ref01_data = helpers.to_map(no_frixion_money_moov_models_card_customer_token_ref01_data_raw[1][2])
    end

    -- LIST
    local no_frixion_money_moov_models_card_customer_token_ref01_ent = client:NoFrixionMoneyMoovModelsCardCustomerToken(nil)
    local no_frixion_money_moov_models_card_customer_token_ref01_match = {
      ["customer_email_address"] = setup.idmap["customer_email_address01"],
      ["merchant_id"] = setup.idmap["merchant01"],
    }

    local no_frixion_money_moov_models_card_customer_token_ref01_list_result, err = no_frixion_money_moov_models_card_customer_token_ref01_ent:list(no_frixion_money_moov_models_card_customer_token_ref01_match, nil)
    assert.is_nil(err)
    assert.is_table(no_frixion_money_moov_models_card_customer_token_ref01_list_result)

    -- LOAD
    local no_frixion_money_moov_models_card_customer_token_ref01_match_dt0 = {
      id = no_frixion_money_moov_models_card_customer_token_ref01_data["id"],
    }
    local no_frixion_money_moov_models_card_customer_token_ref01_data_dt0_loaded, err = no_frixion_money_moov_models_card_customer_token_ref01_ent:load(no_frixion_money_moov_models_card_customer_token_ref01_match_dt0, nil)
    assert.is_nil(err)
    local no_frixion_money_moov_models_card_customer_token_ref01_data_dt0_load_result = helpers.to_map(no_frixion_money_moov_models_card_customer_token_ref01_data_dt0_loaded)
    assert.is_not_nil(no_frixion_money_moov_models_card_customer_token_ref01_data_dt0_load_result)
    assert.are.equal(no_frixion_money_moov_models_card_customer_token_ref01_data_dt0_load_result["id"], no_frixion_money_moov_models_card_customer_token_ref01_data["id"])

    -- REMOVE
    local no_frixion_money_moov_models_card_customer_token_ref01_match_rm0 = {
      id = no_frixion_money_moov_models_card_customer_token_ref01_data["id"],
    }
    local _, err = no_frixion_money_moov_models_card_customer_token_ref01_ent:remove(no_frixion_money_moov_models_card_customer_token_ref01_match_rm0, nil)
    assert.is_nil(err)

    -- LIST
    local no_frixion_money_moov_models_card_customer_token_ref01_match_rt0 = {
      ["customer_email_address"] = setup.idmap["customer_email_address01"],
      ["merchant_id"] = setup.idmap["merchant01"],
    }

    local no_frixion_money_moov_models_card_customer_token_ref01_list_rt0_result, err = no_frixion_money_moov_models_card_customer_token_ref01_ent:list(no_frixion_money_moov_models_card_customer_token_ref01_match_rt0, nil)
    assert.is_nil(err)
    assert.is_table(no_frixion_money_moov_models_card_customer_token_ref01_list_rt0_result)

  end)
end)

function no_frixion_money_moov_models_card_customer_token_basic_setup(extra)
  runner.load_env_local()

  local entity_data_file = _test_dir .. "../../.sdk/test/entity/no_frixion_money_moov_models_card_customer_token/NoFrixionMoneyMoovModelsCardCustomerTokenTestData.json"
  local f = io.open(entity_data_file, "r")
  if f == nil then
    error("failed to read no_frixion_money_moov_models_card_customer_token test data: " .. entity_data_file)
  end
  local entity_data_source = f:read("*a")
  f:close()

  local entity_data = json.decode(entity_data_source)

  local options = {}
  options["entity"] = entity_data["existing"]

  local client = sdk.test(options, extra)

  -- Generate idmap via transform.
  local idmap = vs.transform(
    { "no_frixion_money_moov_models_card_customer_token01", "no_frixion_money_moov_models_card_customer_token02", "no_frixion_money_moov_models_card_customer_token03", "removeall01", "removeall02", "removeall03", "customertoken01", "customertoken02", "customertoken03", "customer_email_address01", "merchant01" },
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
  local entid_env_raw = os.getenv("NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_CARD_CUSTOMER_TOKEN_ENTID")
  local idmap_overridden = entid_env_raw ~= nil and entid_env_raw:match("^%s*{") ~= nil

  local env = runner.env_override({
    ["NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_CARD_CUSTOMER_TOKEN_ENTID"] = idmap,
    ["NOFRIXION_TEST_LIVE"] = "FALSE",
    ["NOFRIXION_TEST_EXPLAIN"] = "FALSE",
    ["NOFRIXION_APIKEY"] = "NONE",
  })

  local idmap_resolved = helpers.to_map(
    env["NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_CARD_CUSTOMER_TOKEN_ENTID"])
  if idmap_resolved == nil then
    idmap_resolved = helpers.to_map(idmap)
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
