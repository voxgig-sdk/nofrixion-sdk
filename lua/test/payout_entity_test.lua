-- Payout entity test

local json = require("dkjson")
local vs = require("utility.struct.struct")
local sdk = require("nofrixion_sdk")
local helpers = require("core.helpers")
local runner = require("test.runner")

local _test_dir = debug.getinfo(1, "S").source:match("^@(.+/)")  or "./"

describe("PayoutEntity", function()
  it("should create instance", function()
    local testsdk = sdk.test(nil, nil)
    local ent = testsdk:Payout(nil)
    assert.is_not_nil(ent)
  end)

  -- Feature #4: the entity stream(action, ...) method runs the op pipeline and
  -- returns an iterator over result items. With the streaming feature active it
  -- yields the feature's incremental output; otherwise it falls back to the
  -- materialised list so stream always yields.
  it("should stream", function()
    local seed = {
      entity = {
        ["payout"] = {
          s1 = { id = "s1" },
          s2 = { id = "s2" },
          s3 = { id = "s3" },
        },
      },
    }

    -- Fallback: streaming inactive -> yields the materialised list items.
    local base = sdk.test(seed, nil)
    local seen = {}
    for item in base:Payout(nil):stream("list", nil, nil) do
      table.insert(seen, item)
    end
    assert.are.equal(3, #seen)

    -- Inbound: streaming active -> yields each item from the feature.
    local config = require("config")()
    if type(config.feature) == "table" and config.feature.streaming ~= nil then
      local streamsdk = sdk.test(seed, { feature = { streaming = { active = true } } })
      local got = {}
      for item in streamsdk:Payout(nil):stream("list", nil, nil) do
        if vs.islist(item) then
          for _, sub in ipairs(item) do
            table.insert(got, sub)
          end
        else
          table.insert(got, item)
        end
      end
      assert.are.equal(3, #got)
    end
  end)

  it("should run basic flow", function()
    local setup = payout_basic_setup(nil)
    -- Per-op sdk-test-control.json skip.
    local _live = setup.live or false
    for _, _op in ipairs({"create", "list", "update", "load", "remove"}) do
      local _should_skip, _reason = runner.is_control_skipped("entityOp", "payout." .. _op, _live and "live" or "unit")
      if _should_skip then
        pending(_reason or "skipped via sdk-test-control.json")
        return
      end
    end
    -- The basic flow consumes synthetic IDs from the fixture. In live mode
    -- without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup.synthetic_only then
      pending("live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_PAYOUT_ENTID JSON to run live")
      return
    end
    local client = setup.client

    -- CREATE
    local payout_ref01_ent = client:Payout(nil)
    local payout_ref01_data = helpers.to_map(vs.getprop(
      vs.getpath(setup.data, "new.payout"), "payout_ref01"))
    payout_ref01_data["account_id"] = setup.idmap["account01"]
    payout_ref01_data["destination"] = setup.idmap["destination01"]
    payout_ref01_data["merchant_id"] = setup.idmap["merchant01"]
    payout_ref01_data["source"] = setup.idmap["source01"]

    local payout_ref01_data_result, err = payout_ref01_ent:create(payout_ref01_data, nil)
    assert.is_nil(err)
    payout_ref01_data = helpers.to_map(payout_ref01_data_result)
    assert.is_not_nil(payout_ref01_data)
    assert.is_not_nil(payout_ref01_data["id"])

    -- LIST
    local payout_ref01_match = {
      ["merchant_id"] = setup.idmap["merchant01"],
    }

    local payout_ref01_list_result, err = payout_ref01_ent:list(payout_ref01_match, nil)
    assert.is_nil(err)
    assert.is_table(payout_ref01_list_result)

    local found_item = vs.select(
      runner.entity_list_to_data(payout_ref01_list_result),
      { id = payout_ref01_data["id"] })
    assert.is_false(vs.isempty(found_item))

    -- UPDATE
    local payout_ref01_data_up0_up = {
      id = payout_ref01_data["id"],
    }

    local payout_ref01_markdef_up0_name = "account_id"
    local payout_ref01_markdef_up0_value = "Mark01-payout_ref01_" .. tostring(setup.now)
    payout_ref01_data_up0_up[payout_ref01_markdef_up0_name] = payout_ref01_markdef_up0_value

    local payout_ref01_resdata_up0_result, err = payout_ref01_ent:update(payout_ref01_data_up0_up, nil)
    assert.is_nil(err)
    local payout_ref01_resdata_up0 = helpers.to_map(payout_ref01_resdata_up0_result)
    assert.is_not_nil(payout_ref01_resdata_up0)
    assert.are.equal(payout_ref01_resdata_up0["id"], payout_ref01_data_up0_up["id"])
    assert.are.equal(payout_ref01_resdata_up0[payout_ref01_markdef_up0_name], payout_ref01_markdef_up0_value)

    -- LOAD
    local payout_ref01_match_dt0 = {
      id = payout_ref01_data["id"],
    }
    local payout_ref01_data_dt0_loaded, err = payout_ref01_ent:load(payout_ref01_match_dt0, nil)
    assert.is_nil(err)
    local payout_ref01_data_dt0_load_result = helpers.to_map(payout_ref01_data_dt0_loaded)
    assert.is_not_nil(payout_ref01_data_dt0_load_result)
    assert.are.equal(payout_ref01_data_dt0_load_result["id"], payout_ref01_data["id"])

    -- REMOVE
    local payout_ref01_match_rm0 = {
      id = payout_ref01_data["id"],
    }
    local _, err = payout_ref01_ent:remove(payout_ref01_match_rm0, nil)
    assert.is_nil(err)

    -- LIST
    local payout_ref01_match_rt0 = {
      ["merchant_id"] = setup.idmap["merchant01"],
    }

    local payout_ref01_list_rt0_result, err = payout_ref01_ent:list(payout_ref01_match_rt0, nil)
    assert.is_nil(err)
    assert.is_table(payout_ref01_list_rt0_result)

    local not_found_item = vs.select(
      runner.entity_list_to_data(payout_ref01_list_rt0_result),
      { id = payout_ref01_data["id"] })
    assert.is_true(vs.isempty(not_found_item))

  end)
end)

function payout_basic_setup(extra)
  runner.load_env_local()

  local entity_data_file = _test_dir .. "../../.sdk/test/entity/payout/PayoutTestData.json"
  local f = io.open(entity_data_file, "r")
  if f == nil then
    error("failed to read payout test data: " .. entity_data_file)
  end
  local entity_data_source = f:read("*a")
  f:close()

  local entity_data = json.decode(entity_data_source)

  local options = {}
  options["entity"] = entity_data["existing"]

  local client = sdk.test(options, extra)

  -- Generate idmap via transform.
  local idmap = vs.transform(
    { "payout01", "payout02", "payout03", "account01", "account02", "account03", "merchant01", "merchant02", "merchant03", "fxquote01", "fxquote02", "fxquote03", "destination01", "source01" },
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
  local entid_env_raw = os.getenv("NOFRIXION_TEST_PAYOUT_ENTID")
  local idmap_overridden = entid_env_raw ~= nil and entid_env_raw:match("^%s*{") ~= nil

  local env = runner.env_override({
    ["NOFRIXION_TEST_PAYOUT_ENTID"] = idmap,
    ["NOFRIXION_TEST_LIVE"] = "FALSE",
    ["NOFRIXION_TEST_EXPLAIN"] = "FALSE",
    ["NOFRIXION_APIKEY"] = "NONE",
  })

  local idmap_resolved = helpers.to_map(
    env["NOFRIXION_TEST_PAYOUT_ENTID"])
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
