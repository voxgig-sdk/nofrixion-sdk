# Payment entity test

import json
import os
import time

import pytest

from utility.voxgig_struct import voxgig_struct as vs
from nofrixion_sdk import NofrixionSDK
from core import helpers

_TEST_DIR = os.path.dirname(os.path.abspath(__file__))
from test import runner


class TestPaymentEntity:

    def test_should_create_instance(self):
        testsdk = NofrixionSDK.test(None, None)
        ent = testsdk.Payment(None)
        assert ent is not None

    def test_should_run_basic_flow(self):
        setup = _payment_basic_setup(None)
        # Per-op sdk-test-control.json skip — basic test exercises a flow with
        # multiple ops; skipping any one skips the whole flow (steps depend
        # on each other).
        _live = setup.get("live", False)
        for _op in ["create", "update", "load"]:
            _skip, _reason = runner.is_control_skipped("entityOp", "payment." + _op, "live" if _live else "unit")
            if _skip:
                pytest.skip(_reason or "skipped via sdk-test-control.json")
                return
        # The basic flow consumes synthetic IDs from the fixture. In live mode
        # without an *_ENTID env override, those IDs hit the live API and 4xx.
        if setup.get("synthetic_only"):
            pytest.skip("live entity test uses synthetic IDs from fixture — "
                        "set NOFRIXION_TEST_PAYMENT_ENTID JSON to run live")
        client = setup["client"]

        # CREATE
        payment_ref01_ent = client.Payment(None)
        payment_ref01_data = helpers.to_map(vs.getprop(
            vs.getpath(setup["data"], "new.payment"), "payment_ref01"))

        payment_ref01_data = helpers.to_map(payment_ref01_ent.create(payment_ref01_data, None))
        assert payment_ref01_data is not None
        assert payment_ref01_data["id"] is not None

        # UPDATE
        payment_ref01_data_up0_up = {
            "id": payment_ref01_data["id"],
        }

        payment_ref01_markdef_up0_name = "base_origin_url"
        payment_ref01_markdef_up0_value = "Mark01-payment_ref01_" + str(setup["now"])
        payment_ref01_data_up0_up[payment_ref01_markdef_up0_name] = payment_ref01_markdef_up0_value

        payment_ref01_resdata_up0 = helpers.to_map(payment_ref01_ent.update(payment_ref01_data_up0_up, None))
        assert payment_ref01_resdata_up0 is not None
        assert payment_ref01_resdata_up0["id"] == payment_ref01_data_up0_up["id"]
        assert payment_ref01_resdata_up0[payment_ref01_markdef_up0_name] == payment_ref01_markdef_up0_value

        # LOAD
        payment_ref01_match_dt0 = {
            "id": payment_ref01_data["id"],
        }
        payment_ref01_data_dt0_loaded = payment_ref01_ent.load(payment_ref01_match_dt0, None)
        payment_ref01_data_dt0_load_result = helpers.to_map(payment_ref01_data_dt0_loaded)
        assert payment_ref01_data_dt0_load_result is not None
        assert payment_ref01_data_dt0_load_result["id"] == payment_ref01_data["id"]



def _payment_basic_setup(extra):
    runner.load_env_local()

    entity_data_file = os.path.join(_TEST_DIR, "../../.sdk/test/entity/payment/PaymentTestData.json")
    with open(entity_data_file, "r") as f:
        entity_data_source = f.read()

    entity_data = json.loads(entity_data_source)

    options = {}
    options["entity"] = entity_data.get("existing")

    client = NofrixionSDK.test(options, extra)

    # Generate idmap via transform.
    idmap = vs.transform(
        ["payment01", "payment02", "payment03", "getbyorderid01", "getbyorderid02", "getbyorderid03"],
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
        "NOFRIXION_TEST_PAYMENT_ENTID")
    _idmap_overridden = _entid_env_raw is not None and _entid_env_raw.strip().startswith("{")

    env = runner.env_override({
        "NOFRIXION_TEST_PAYMENT_ENTID": idmap,
        "NOFRIXION_TEST_LIVE": "FALSE",
        "NOFRIXION_TEST_EXPLAIN": "FALSE",
        "NOFRIXION_APIKEY": "NONE",
    })

    idmap_resolved = helpers.to_map(
        env.get("NOFRIXION_TEST_PAYMENT_ENTID"))
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
