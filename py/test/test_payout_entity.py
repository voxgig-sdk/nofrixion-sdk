# Payout entity test

import json
import os
import time

import pytest

from utility.voxgig_struct import voxgig_struct as vs
from nofrixion_sdk import NofrixionSDK
from core import helpers

_TEST_DIR = os.path.dirname(os.path.abspath(__file__))
from test import runner


class TestPayoutEntity:

    def test_should_create_instance(self):
        testsdk = NofrixionSDK.test(None, None)
        ent = testsdk.Payout(None)
        assert ent is not None

    def test_should_run_basic_flow(self):
        setup = _payout_basic_setup(None)
        # Per-op sdk-test-control.json skip — basic test exercises a flow with
        # multiple ops; skipping any one skips the whole flow (steps depend
        # on each other).
        _live = setup.get("live", False)
        for _op in ["create", "update", "load", "remove"]:
            _skip, _reason = runner.is_control_skipped("entityOp", "payout." + _op, "live" if _live else "unit")
            if _skip:
                pytest.skip(_reason or "skipped via sdk-test-control.json")
                return
        # The basic flow consumes synthetic IDs from the fixture. In live mode
        # without an *_ENTID env override, those IDs hit the live API and 4xx.
        if setup.get("synthetic_only"):
            pytest.skip("live entity test uses synthetic IDs from fixture — "
                        "set NOFRIXION_TEST_PAYOUT_ENTID JSON to run live")
        client = setup["client"]

        # CREATE
        payout_ref01_ent = client.Payout(None)
        payout_ref01_data = helpers.to_map(vs.getprop(
            vs.getpath(setup["data"], "new.payout"), "payout_ref01"))
        payout_ref01_data["destination"] = setup["idmap"]["destination01"]
        payout_ref01_data["source"] = setup["idmap"]["source01"]

        payout_ref01_data = helpers.to_map(payout_ref01_ent.create(payout_ref01_data, None))
        assert payout_ref01_data is not None
        assert payout_ref01_data["id"] is not None

        # UPDATE
        payout_ref01_data_up0_up = {
            "id": payout_ref01_data["id"],
        }

        payout_ref01_markdef_up0_name = "account_id"
        payout_ref01_markdef_up0_value = "Mark01-payout_ref01_" + str(setup["now"])
        payout_ref01_data_up0_up[payout_ref01_markdef_up0_name] = payout_ref01_markdef_up0_value

        payout_ref01_resdata_up0 = helpers.to_map(payout_ref01_ent.update(payout_ref01_data_up0_up, None))
        assert payout_ref01_resdata_up0 is not None
        assert payout_ref01_resdata_up0["id"] == payout_ref01_data_up0_up["id"]
        assert payout_ref01_resdata_up0[payout_ref01_markdef_up0_name] == payout_ref01_markdef_up0_value

        # LOAD
        payout_ref01_match_dt0 = {
            "id": payout_ref01_data["id"],
        }
        payout_ref01_data_dt0_loaded = payout_ref01_ent.load(payout_ref01_match_dt0, None)
        payout_ref01_data_dt0_load_result = helpers.to_map(payout_ref01_data_dt0_loaded)
        assert payout_ref01_data_dt0_load_result is not None
        assert payout_ref01_data_dt0_load_result["id"] == payout_ref01_data["id"]

        # REMOVE
        payout_ref01_match_rm0 = {
            "id": payout_ref01_data["id"],
        }
        payout_ref01_ent.remove(payout_ref01_match_rm0, None)



def _payout_basic_setup(extra):
    runner.load_env_local()

    entity_data_file = os.path.join(_TEST_DIR, "../../.sdk/test/entity/payout/PayoutTestData.json")
    with open(entity_data_file, "r") as f:
        entity_data_source = f.read()

    entity_data = json.loads(entity_data_source)

    options = {}
    options["entity"] = entity_data.get("existing")

    client = NofrixionSDK.test(options, extra)

    # Generate idmap via transform.
    idmap = vs.transform(
        ["payout01", "payout02", "payout03", "fxquote01", "fxquote02", "fxquote03", "destination01", "source01"],
        {
            "`$PACK`": ["", {
                "`$KEY`": "`$COPY`",
                "`$VAL`": ["`$FORMAT`", "upper", "`$COPY`"],
            }],
        }
    )

    # Detect ENTID env override before envOverride consumes it. When live
    # mode is on without a real override, the basic test runs against synthetic
    # IDs from the fixture and 4xx's. We surface this so the test can skip.
    _entid_env_raw = os.environ.get(
        "NOFRIXION_TEST_PAYOUT_ENTID")
    _idmap_overridden = _entid_env_raw is not None and _entid_env_raw.strip().startswith("{")

    env = runner.env_override({
        "NOFRIXION_TEST_PAYOUT_ENTID": idmap,
        "NOFRIXION_TEST_LIVE": "FALSE",
        "NOFRIXION_TEST_EXPLAIN": "FALSE",
        "NOFRIXION_APIKEY": "NONE",
    })

    idmap_resolved = helpers.to_map(
        env.get("NOFRIXION_TEST_PAYOUT_ENTID"))
    if idmap_resolved is None:
        idmap_resolved = helpers.to_map(idmap)

    if env.get("NOFRIXION_TEST_LIVE") == "TRUE":
        merged_opts = vs.merge([
            {
                "apikey": env.get("NOFRIXION_APIKEY"),
            },
            extra or {},
        ])
        client = NofrixionSDK(helpers.to_map(merged_opts))

    _live = env.get("NOFRIXION_TEST_LIVE") == "TRUE"
    return {
        "client": client,
        "data": entity_data,
        "idmap": idmap_resolved,
        "env": env,
        "explain": env.get("NOFRIXION_TEST_EXPLAIN") == "TRUE",
        "live": _live,
        "synthetic_only": _live and not _idmap_overridden,
        "now": int(time.time() * 1000),
    }
