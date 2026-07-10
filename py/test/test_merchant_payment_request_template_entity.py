# MerchantPaymentRequestTemplate entity test

import json
import os
import time

import pytest

from utility.voxgig_struct import voxgig_struct as vs
from nofrixion_sdk import NofrixionSDK
from core import helpers

_TEST_DIR = os.path.dirname(os.path.abspath(__file__))
from test import runner


class TestMerchantPaymentRequestTemplateEntity:

    def test_should_create_instance(self):
        testsdk = NofrixionSDK.test(None, None)
        ent = testsdk.MerchantPaymentRequestTemplate(None)
        assert ent is not None

    def test_should_run_basic_flow(self):
        setup = _merchant_payment_request_template_basic_setup(None)
        # Per-op sdk-test-control.json skip — basic test exercises a flow with
        # multiple ops; skipping any one skips the whole flow (steps depend
        # on each other).
        _live = setup.get("live", False)
        for _op in ["list", "update", "load", "remove"]:
            _skip, _reason = runner.is_control_skipped("entityOp", "merchant_payment_request_template." + _op, "live" if _live else "unit")
            if _skip:
                pytest.skip(_reason or "skipped via sdk-test-control.json")
                return
        # The basic flow consumes synthetic IDs from the fixture. In live mode
        # without an *_ENTID env override, those IDs hit the live API and 4xx.
        if setup.get("synthetic_only"):
            pytest.skip("live entity test uses synthetic IDs from fixture — "
                        "set NOFRIXION_TEST_MERCHANT_PAYMENT_REQUEST_TEMPLATE_ENTID JSON to run live")
        client = setup["client"]

        # Bootstrap entity data from existing test data.
        merchant_payment_request_template_ref01_data_raw = vs.items(helpers.to_map(
            vs.getpath(setup["data"], "existing.merchant_payment_request_template")))
        merchant_payment_request_template_ref01_data = None
        if len(merchant_payment_request_template_ref01_data_raw) > 0:
            merchant_payment_request_template_ref01_data = helpers.to_map(merchant_payment_request_template_ref01_data_raw[0][1])

        # LIST
        merchant_payment_request_template_ref01_ent = client.MerchantPaymentRequestTemplate(None)
        merchant_payment_request_template_ref01_match = {
            "merchant_id": setup["idmap"]["merchant01"],
        }

        merchant_payment_request_template_ref01_list_result = merchant_payment_request_template_ref01_ent.list(merchant_payment_request_template_ref01_match, None)
        assert isinstance(merchant_payment_request_template_ref01_list_result, list)

        # UPDATE
        merchant_payment_request_template_ref01_data_up0_up = {
            "id": merchant_payment_request_template_ref01_data["id"],
            "paymentrequest_id": setup["idmap"]["paymentrequest_id"],
        }

        merchant_payment_request_template_ref01_markdef_up0_name = "description"
        merchant_payment_request_template_ref01_markdef_up0_value = "Mark01-merchant_payment_request_template_ref01_" + str(setup["now"])
        merchant_payment_request_template_ref01_data_up0_up[merchant_payment_request_template_ref01_markdef_up0_name] = merchant_payment_request_template_ref01_markdef_up0_value

        merchant_payment_request_template_ref01_resdata_up0 = helpers.to_map(merchant_payment_request_template_ref01_ent.update(merchant_payment_request_template_ref01_data_up0_up, None))
        assert merchant_payment_request_template_ref01_resdata_up0 is not None
        assert merchant_payment_request_template_ref01_resdata_up0["id"] == merchant_payment_request_template_ref01_data_up0_up["id"]
        assert merchant_payment_request_template_ref01_resdata_up0[merchant_payment_request_template_ref01_markdef_up0_name] == merchant_payment_request_template_ref01_markdef_up0_value

        # LOAD
        merchant_payment_request_template_ref01_match_dt0 = {
            "id": merchant_payment_request_template_ref01_data["id"],
        }
        merchant_payment_request_template_ref01_data_dt0_loaded = merchant_payment_request_template_ref01_ent.load(merchant_payment_request_template_ref01_match_dt0, None)
        merchant_payment_request_template_ref01_data_dt0_load_result = helpers.to_map(merchant_payment_request_template_ref01_data_dt0_loaded)
        assert merchant_payment_request_template_ref01_data_dt0_load_result is not None
        assert merchant_payment_request_template_ref01_data_dt0_load_result["id"] == merchant_payment_request_template_ref01_data["id"]

        # REMOVE
        merchant_payment_request_template_ref01_match_rm0 = {
            "id": merchant_payment_request_template_ref01_data["id"],
        }
        merchant_payment_request_template_ref01_ent.remove(merchant_payment_request_template_ref01_match_rm0, None)

        # LIST
        merchant_payment_request_template_ref01_match_rt0 = {
            "merchant_id": setup["idmap"]["merchant01"],
        }

        merchant_payment_request_template_ref01_list_rt0_result = merchant_payment_request_template_ref01_ent.list(merchant_payment_request_template_ref01_match_rt0, None)
        assert isinstance(merchant_payment_request_template_ref01_list_rt0_result, list)



def _merchant_payment_request_template_basic_setup(extra):
    runner.load_env_local()

    entity_data_file = os.path.join(_TEST_DIR, "../../.sdk/test/entity/merchant_payment_request_template/MerchantPaymentRequestTemplateTestData.json")
    with open(entity_data_file, "r") as f:
        entity_data_source = f.read()

    entity_data = json.loads(entity_data_source)

    options = {}
    options["entity"] = entity_data.get("existing")

    client = NofrixionSDK.test(options, extra)

    # Generate idmap via transform.
    idmap = vs.transform(
        ["merchant_payment_request_template01", "merchant_payment_request_template02", "merchant_payment_request_template03", "paymentrequest01", "paymentrequest02", "paymentrequest03", "merchant01"],
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
        "NOFRIXION_TEST_MERCHANT_PAYMENT_REQUEST_TEMPLATE_ENTID")
    _idmap_overridden = _entid_env_raw is not None and _entid_env_raw.strip().startswith("{")

    env = runner.env_override({
        "NOFRIXION_TEST_MERCHANT_PAYMENT_REQUEST_TEMPLATE_ENTID": idmap,
        "NOFRIXION_TEST_LIVE": "FALSE",
        "NOFRIXION_TEST_EXPLAIN": "FALSE",
        "NOFRIXION_APIKEY": "NONE",
    })

    idmap_resolved = helpers.to_map(
        env.get("NOFRIXION_TEST_MERCHANT_PAYMENT_REQUEST_TEMPLATE_ENTID"))
    if idmap_resolved is None:
        idmap_resolved = helpers.to_map(idmap)
    if idmap_resolved.get("paymentrequest_id") is None:
        idmap_resolved["paymentrequest_id"] = idmap_resolved.get("paymentrequest01")

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
