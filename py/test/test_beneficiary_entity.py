# Beneficiary entity test

import json
import os
import time

import pytest

from utility.voxgig_struct import voxgig_struct as vs
from nofrixion_sdk import NofrixionSDK
from core import helpers

_TEST_DIR = os.path.dirname(os.path.abspath(__file__))
from test import runner


class TestBeneficiaryEntity:

    def test_should_create_instance(self):
        testsdk = NofrixionSDK.test(None, None)
        ent = testsdk.Beneficiary(None)
        assert ent is not None

    def test_should_stream(self):
        # Feature #4: the entity stream(action, ...) method runs the op
        # pipeline and yields result items. With the streaming feature active
        # it yields the feature's incremental output; otherwise it falls back
        # to the materialised list so stream always yields.
        seed = {
            "entity": {
                "beneficiary": {
                    "s1": {"id": "s1"},
                    "s2": {"id": "s2"},
                    "s3": {"id": "s3"},
                }
            }
        }

        # Fallback: streaming inactive -> yields the materialised list items.
        base = NofrixionSDK.test(seed, None)
        seen = list(base.Beneficiary(None).stream("list", None, None))
        assert len(seen) == 3

        # Inbound: streaming active -> yields each item from the feature.
        from config import make_config
        cfg = make_config()
        if isinstance(cfg.get("feature"), dict) and "streaming" in cfg["feature"]:
            sdk = NofrixionSDK.test(
                seed, {"feature": {"streaming": {"active": True}}})
            got = []
            for item in sdk.Beneficiary(None).stream("list", None, None):
                if isinstance(item, list):
                    got.extend(item)
                else:
                    got.append(item)
            assert len(got) == 3

    def test_should_run_basic_flow(self):
        setup = _beneficiary_basic_setup(None)
        # Per-op sdk-test-control.json skip — basic test exercises a flow with
        # multiple ops; skipping any one skips the whole flow (steps depend
        # on each other).
        _live = setup.get("live", False)
        for _op in ["create", "list", "update", "load", "remove"]:
            _skip, _reason = runner.is_control_skipped("entityOp", "beneficiary." + _op, "live" if _live else "unit")
            if _skip:
                pytest.skip(_reason or "skipped via sdk-test-control.json")
                return
        # The basic flow consumes synthetic IDs from the fixture. In live mode
        # without an *_ENTID env override, those IDs hit the live API and 4xx.
        if setup.get("synthetic_only"):
            pytest.skip("live entity test uses synthetic IDs from fixture — "
                        "set NOFRIXION_TEST_BENEFICIARY_ENTID JSON to run live")
        client = setup["client"]

        # CREATE
        beneficiary_ref01_ent = client.Beneficiary(None)
        beneficiary_ref01_data = helpers.to_map(vs.getprop(
            vs.getpath(setup["data"], "new.beneficiary"), "beneficiary_ref01"))
        beneficiary_ref01_data["merchant_id"] = setup["idmap"]["merchant01"]

        beneficiary_ref01_data = helpers.to_map(beneficiary_ref01_ent.create(beneficiary_ref01_data, None))
        assert beneficiary_ref01_data is not None
        assert beneficiary_ref01_data["id"] is not None

        # LIST
        beneficiary_ref01_match = {
            "merchant_id": setup["idmap"]["merchant01"],
        }

        beneficiary_ref01_list_result = beneficiary_ref01_ent.list(beneficiary_ref01_match, None)
        assert isinstance(beneficiary_ref01_list_result, list)

        found_item = vs.select(
            runner.entity_list_to_data(beneficiary_ref01_list_result),
            {"id": beneficiary_ref01_data["id"]})
        assert not vs.isempty(found_item)

        # UPDATE
        beneficiary_ref01_data_up0_up = {
            "id": beneficiary_ref01_data["id"],
        }

        beneficiary_ref01_markdef_up0_name = "approval_callback_url"
        beneficiary_ref01_markdef_up0_value = "Mark01-beneficiary_ref01_" + str(setup["now"])
        beneficiary_ref01_data_up0_up[beneficiary_ref01_markdef_up0_name] = beneficiary_ref01_markdef_up0_value

        beneficiary_ref01_resdata_up0 = helpers.to_map(beneficiary_ref01_ent.update(beneficiary_ref01_data_up0_up, None))
        assert beneficiary_ref01_resdata_up0 is not None
        assert beneficiary_ref01_resdata_up0["id"] == beneficiary_ref01_data_up0_up["id"]
        assert beneficiary_ref01_resdata_up0[beneficiary_ref01_markdef_up0_name] == beneficiary_ref01_markdef_up0_value

        # LOAD
        beneficiary_ref01_match_dt0 = {
            "id": beneficiary_ref01_data["id"],
        }
        beneficiary_ref01_data_dt0_loaded = beneficiary_ref01_ent.load(beneficiary_ref01_match_dt0, None)
        beneficiary_ref01_data_dt0_load_result = helpers.to_map(beneficiary_ref01_data_dt0_loaded)
        assert beneficiary_ref01_data_dt0_load_result is not None
        assert beneficiary_ref01_data_dt0_load_result["id"] == beneficiary_ref01_data["id"]

        # REMOVE
        beneficiary_ref01_match_rm0 = {
            "id": beneficiary_ref01_data["id"],
        }
        beneficiary_ref01_ent.remove(beneficiary_ref01_match_rm0, None)

        # LIST
        beneficiary_ref01_match_rt0 = {
            "merchant_id": setup["idmap"]["merchant01"],
        }

        beneficiary_ref01_list_rt0_result = beneficiary_ref01_ent.list(beneficiary_ref01_match_rt0, None)
        assert isinstance(beneficiary_ref01_list_rt0_result, list)

        not_found_item = vs.select(
            runner.entity_list_to_data(beneficiary_ref01_list_rt0_result),
            {"id": beneficiary_ref01_data["id"]})
        assert vs.isempty(not_found_item)



def _beneficiary_basic_setup(extra):
    runner.load_env_local()

    entity_data_file = os.path.join(_TEST_DIR, "../../.sdk/test/entity/beneficiary/BeneficiaryTestData.json")
    with open(entity_data_file, "r") as f:
        entity_data_source = f.read()

    entity_data = json.loads(entity_data_source)

    options = {}
    options["entity"] = entity_data.get("existing")

    client = NofrixionSDK.test(options, extra)

    # Generate idmap via transform.
    idmap = vs.transform(
        ["beneficiary01", "beneficiary02", "beneficiary03", "merchant01", "merchant02", "merchant03"],
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
        "NOFRIXION_TEST_BENEFICIARY_ENTID")
    _idmap_overridden = _entid_env_raw is not None and _entid_env_raw.strip().startswith("{")

    env = runner.env_override({
        "NOFRIXION_TEST_BENEFICIARY_ENTID": idmap,
        "NOFRIXION_TEST_LIVE": "FALSE",
        "NOFRIXION_TEST_EXPLAIN": "FALSE",
        "NOFRIXION_APIKEY": "NONE",
    })

    idmap_resolved = helpers.to_map(
        env.get("NOFRIXION_TEST_BENEFICIARY_ENTID"))
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
