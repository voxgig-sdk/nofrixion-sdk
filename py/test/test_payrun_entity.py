# Payrun entity test

import json
import os
import time

import pytest

from utility.voxgig_struct import voxgig_struct as vs
from nofrixion_sdk import NofrixionSDK
from core import helpers

_TEST_DIR = os.path.dirname(os.path.abspath(__file__))
from test import runner


class TestPayrunEntity:

    def test_should_create_instance(self):
        testsdk = NofrixionSDK.test(None, None)
        ent = testsdk.Payrun(None)
        assert ent is not None

    def test_should_run_basic_flow(self):
        setup = _payrun_basic_setup(None)
        # Per-op sdk-test-control.json skip — basic test exercises a flow with
        # multiple ops; skipping any one skips the whole flow (steps depend
        # on each other).
        _live = setup.get("live", False)
        for _op in ["create", "update", "remove"]:
            _skip, _reason = runner.is_control_skipped("entityOp", "payrun." + _op, "live" if _live else "unit")
            if _skip:
                pytest.skip(_reason or "skipped via sdk-test-control.json")
                return
        # The basic flow consumes synthetic IDs from the fixture. In live mode
        # without an *_ENTID env override, those IDs hit the live API and 4xx.
        if setup.get("synthetic_only"):
            pytest.skip("live entity test uses synthetic IDs from fixture — "
                        "set NOFRIXION_TEST_PAYRUN_ENTID JSON to run live")
        client = setup["client"]

        # CREATE
        payrun_ref01_ent = client.Payrun(None)
        payrun_ref01_data = helpers.to_map(vs.getprop(
            vs.getpath(setup["data"], "new.payrun"), "payrun_ref01"))

        payrun_ref01_data = helpers.to_map(payrun_ref01_ent.create(payrun_ref01_data, None))
        assert payrun_ref01_data is not None
        assert payrun_ref01_data["id"] is not None

        # UPDATE
        payrun_ref01_data_up0_up = {
            "id": payrun_ref01_data["id"],
        }

        payrun_ref01_markdef_up0_name = "note"
        payrun_ref01_markdef_up0_value = "Mark01-payrun_ref01_" + str(setup["now"])
        payrun_ref01_data_up0_up[payrun_ref01_markdef_up0_name] = payrun_ref01_markdef_up0_value

        payrun_ref01_resdata_up0 = helpers.to_map(payrun_ref01_ent.update(payrun_ref01_data_up0_up, None))
        assert payrun_ref01_resdata_up0 is not None
        assert payrun_ref01_resdata_up0["id"] == payrun_ref01_data_up0_up["id"]
        assert payrun_ref01_resdata_up0[payrun_ref01_markdef_up0_name] == payrun_ref01_markdef_up0_value

        # REMOVE
        payrun_ref01_match_rm0 = {
            "id": payrun_ref01_data["id"],
        }
        payrun_ref01_ent.remove(payrun_ref01_match_rm0, None)



def _payrun_basic_setup(extra):
    runner.load_env_local()

    entity_data_file = os.path.join(_TEST_DIR, "../../.sdk/test/entity/payrun/PayrunTestData.json")
    with open(entity_data_file, "r") as f:
        entity_data_source = f.read()

    entity_data = json.loads(entity_data_source)

    options = {}
    options["entity"] = entity_data.get("existing")

    client = NofrixionSDK.test(options, extra)

    # Generate idmap via transform.
    idmap = vs.transform(
        ["payrun01", "payrun02", "payrun03"],
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
        "NOFRIXION_TEST_PAYRUN_ENTID")
    _idmap_overridden = _entid_env_raw is not None and _entid_env_raw.strip().startswith("{")

    env = runner.env_override({
        "NOFRIXION_TEST_PAYRUN_ENTID": idmap,
        "NOFRIXION_TEST_LIVE": "FALSE",
        "NOFRIXION_TEST_EXPLAIN": "FALSE",
        "NOFRIXION_APIKEY": "NONE",
    })

    idmap_resolved = helpers.to_map(
        env.get("NOFRIXION_TEST_PAYRUN_ENTID"))
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
