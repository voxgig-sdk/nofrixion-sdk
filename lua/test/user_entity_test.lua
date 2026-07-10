-- User entity test

local json = require("dkjson")
local vs = require("utility.struct.struct")
local sdk = require("nofrixion_sdk")
local helpers = require("core.helpers")
local runner = require("test.runner")

local _test_dir = debug.getinfo(1, "S").source:match("^@(.+/)")  or "./"

describe("UserEntity", function()
  it("should create instance", function()
    local testsdk = sdk.test(nil, nil)
    local ent = testsdk:User(nil)
    assert.is_not_nil(ent)
  end)

  it("should run basic flow", function()
    local setup = user_basic_setup(nil)
    -- Per-op sdk-test-control.json skip.
    local _live = setup.live or false
    for _, _op in ipairs({"list", "update"}) do
      local _should_skip, _reason = runner.is_control_skipped("entityOp", "user." .. _op, _live and "live" or "unit")
      if _should_skip then
        pending(_reason or "skipped via sdk-test-control.json")
        return
      end
    end
    -- The basic flow consumes synthetic IDs from the fixture. In live mode
    -- without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup.synthetic_only then
      pending("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_USER_ENTID JSON to run live")
      return
    end
    local client = setup.client

    -- Bootstrap entity data from existing test data.
    local user_ref01_data_raw = vs.items(helpers.to_map(
      vs.getpath(setup.data, "existing.user")))
    local user_ref01_data = nil
    if #user_ref01_data_raw > 0 then
      user_ref01_data = helpers.to_map(user_ref01_data_raw[1][2])
    end

    -- LIST
    local user_ref01_ent = client:User(nil)
    local user_ref01_match = {}

    local user_ref01_list_result, err = user_ref01_ent:list(user_ref01_match, nil)
    assert.is_nil(err)
    assert.is_table(user_ref01_list_result)

    -- UPDATE
    local user_ref01_data_up0_up = {
      id = user_ref01_data["id"],
    }

    local user_ref01_markdef_up0_name = "email_address"
    local user_ref01_markdef_up0_value = "Mark01-user_ref01_" .. tostring(setup.now)
    user_ref01_data_up0_up[user_ref01_markdef_up0_name] = user_ref01_markdef_up0_value

    local user_ref01_resdata_up0_result, err = user_ref01_ent:update(user_ref01_data_up0_up, nil)
    assert.is_nil(err)
    local user_ref01_resdata_up0 = helpers.to_map(user_ref01_resdata_up0_result)
    assert.is_not_nil(user_ref01_resdata_up0)
    assert.are.equal(user_ref01_resdata_up0["id"], user_ref01_data_up0_up["id"])
    assert.are.equal(user_ref01_resdata_up0[user_ref01_markdef_up0_name], user_ref01_markdef_up0_value)

  end)
end)

function user_basic_setup(extra)
  runner.load_env_local()

  local entity_data_file = _test_dir .. "../../.sdk/test/entity/user/UserTestData.json"
  local f = io.open(entity_data_file, "r")
  if f == nil then
    error("failed to read user test data: " .. entity_data_file)
  end
  local entity_data_source = f:read("*a")
  f:close()

  local entity_data = json.decode(entity_data_source)

  local options = {}
  options["entity"] = entity_data["existing"]

  local client = sdk.test(options, extra)

  -- Generate idmap via transform.
  local idmap = vs.transform(
    { "user01", "user02", "user03", "merchant01", "merchant02", "merchant03" },
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
  local entid_env_raw = os.getenv("NOFRIXION_TEST_USER_ENTID")
  local idmap_overridden = entid_env_raw ~= nil and entid_env_raw:match("^%s*{") ~= nil

  local env = runner.env_override({
    ["NOFRIXION_TEST_USER_ENTID"] = idmap,
    ["NOFRIXION_TEST_LIVE"] = "FALSE",
    ["NOFRIXION_TEST_EXPLAIN"] = "FALSE",
    ["NOFRIXION_APIKEY"] = "NONE",
  })

  local idmap_resolved = helpers.to_map(
    env["NOFRIXION_TEST_USER_ENTID"])
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
