-- Payment entity test

local json = require("dkjson")
local vs = require("utility.struct.struct")
local sdk = require("nofrixion_sdk")
local helpers = require("core.helpers")
local runner = require("test.runner")

local _test_dir = debug.getinfo(1, "S").source:match("^@(.+/)")  or "./"

describe("PaymentEntity", function()
  it("should create instance", function()
    local testsdk = sdk.test(nil, nil)
    local ent = testsdk:Payment(nil)
    assert.is_not_nil(ent)
  end)

  it("should run basic flow", function()
    local setup = payment_basic_setup(nil)
    -- Per-op sdk-test-control.json skip.
    local _live = setup.live or false
    for _, _op in ipairs({"create", "update", "load"}) do
      local _should_skip, _reason = runner.is_control_skipped("entityOp", "payment." .. _op, _live and "live" or "unit")
      if _should_skip then
        pending(_reason or "skipped via sdk-test-control.json")
        return
      end
    end
    -- The basic flow consumes synthetic IDs from the fixture. In live mode
    -- without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup.synthetic_only then
      pending("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_PAYMENT_ENTID JSON to run live")
      return
    end
    local client = setup.client

    -- CREATE
    local payment_ref01_ent = client:Payment(nil)
    local payment_ref01_data = helpers.to_map(vs.getprop(
      vs.getpath(setup.data, "new.payment"), "payment_ref01"))

    local payment_ref01_data_result, err = payment_ref01_ent:create(payment_ref01_data, nil)
    assert.is_nil(err)
    payment_ref01_data = helpers.to_map(payment_ref01_data_result)
    assert.is_not_nil(payment_ref01_data)
    assert.is_not_nil(payment_ref01_data["id"])

    -- UPDATE
    local payment_ref01_data_up0_up = {
      id = payment_ref01_data["id"],
    }

    local payment_ref01_markdef_up0_name = "base_origin_url"
    local payment_ref01_markdef_up0_value = "Mark01-payment_ref01_" .. tostring(setup.now)
    payment_ref01_data_up0_up[payment_ref01_markdef_up0_name] = payment_ref01_markdef_up0_value

    local payment_ref01_resdata_up0_result, err = payment_ref01_ent:update(payment_ref01_data_up0_up, nil)
    assert.is_nil(err)
    local payment_ref01_resdata_up0 = helpers.to_map(payment_ref01_resdata_up0_result)
    assert.is_not_nil(payment_ref01_resdata_up0)
    assert.are.equal(payment_ref01_resdata_up0["id"], payment_ref01_data_up0_up["id"])
    assert.are.equal(payment_ref01_resdata_up0[payment_ref01_markdef_up0_name], payment_ref01_markdef_up0_value)

    -- LOAD
    local payment_ref01_match_dt0 = {
      id = payment_ref01_data["id"],
    }
    local payment_ref01_data_dt0_loaded, err = payment_ref01_ent:load(payment_ref01_match_dt0, nil)
    assert.is_nil(err)
    local payment_ref01_data_dt0_load_result = helpers.to_map(payment_ref01_data_dt0_loaded)
    assert.is_not_nil(payment_ref01_data_dt0_load_result)
    assert.are.equal(payment_ref01_data_dt0_load_result["id"], payment_ref01_data["id"])

  end)
end)

function payment_basic_setup(extra)
  runner.load_env_local()

  local entity_data_file = _test_dir .. "../../.sdk/test/entity/payment/PaymentTestData.json"
  local f = io.open(entity_data_file, "r")
  if f == nil then
    error("failed to read payment test data: " .. entity_data_file)
  end
  local entity_data_source = f:read("*a")
  f:close()

  local entity_data = json.decode(entity_data_source)

  local options = {}
  options["entity"] = entity_data["existing"]

  local client = sdk.test(options, extra)

  -- Generate idmap via transform.
  local idmap = vs.transform(
    { "payment01", "payment02", "payment03", "getbyorderid01", "getbyorderid02", "getbyorderid03" },
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
  local entid_env_raw = os.getenv("NOFRIXION_TEST_PAYMENT_ENTID")
  local idmap_overridden = entid_env_raw ~= nil and entid_env_raw:match("^%s*{") ~= nil

  local env = runner.env_override({
    ["NOFRIXION_TEST_PAYMENT_ENTID"] = idmap,
    ["NOFRIXION_TEST_LIVE"] = "FALSE",
    ["NOFRIXION_TEST_EXPLAIN"] = "FALSE",
    ["NOFRIXION_APIKEY"] = "NONE",
  })

  local idmap_resolved = helpers.to_map(
    env["NOFRIXION_TEST_PAYMENT_ENTID"])
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
