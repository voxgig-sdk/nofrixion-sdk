# NoFrixionMoneyMoovModelsOpenBankingTransaction direct test

import json
import pytest

from utility.voxgig_struct import voxgig_struct as vs
from nofrixion_sdk import NofrixionSDK
from core import helpers
from test import runner


class TestNoFrixionMoneyMoovModelsOpenBankingTransactionDirect:

    def test_should_direct_list_no_frixion_money_moov_models_open_banking_transaction(self):
        setup = _no_frixion_money_moov_models_open_banking_transaction_direct_setup([
            {"id": "direct01"},
            {"id": "direct02"},
        ])
        _skip, _reason = runner.is_control_skipped("direct", "direct-list-no_frixion_money_moov_models_open_banking_transaction", "live" if setup["live"] else "unit")
        if _skip:
            # pytest already imported at module scope
            pytest.skip(_reason or "skipped via sdk-test-control.json")
            return
        if setup["live"]:
            for _live_key in ["account01", "no_frixion_money_moov_models_open_banking_transaction01"]:
                if setup["idmap"].get(_live_key) is None:
                    # pytest already imported at module scope
                    pytest.skip(f"live test needs {_live_key} via *_ENTID env var (synthetic IDs only)")
                    return

        client = setup["client"]

        params = {}
        if setup["live"]:
            params["account_id"] = setup["idmap"]["account01"]
        else:
            params["account_id"] = "direct01"
        if setup["live"]:
            params["id"] = setup["idmap"]["no_frixion_money_moov_models_open_banking_transaction01"]
        else:
            params["id"] = "direct01"

        result = client.direct({
            "path": "api/v1/openbanking/transactions/{id}/{account_id}",
            "method": "GET",
            "params": params,
        })
        if setup["live"]:
            # Live mode is lenient: synthetic IDs frequently 4xx and the
            # list-response shape varies wildly across public APIs. Skip
            # rather than fail when the call doesn't return a usable list.
            if result.get("err") is not None:
                pytest.skip(f"list call failed (likely synthetic IDs against live API): {result.get('err')}")
                return
            if not result.get("ok"):
                pytest.skip("list call not ok (likely synthetic IDs against live API)")
                return
            status = helpers.to_int(result["status"])
            if status < 200 or status >= 300:
                pytest.skip(f"expected 2xx status, got {status}")
                return
        else:
            assert result["ok"] is True
            assert helpers.to_int(result["status"]) == 200
            assert isinstance(result["data"], list)
            assert len(result["data"]) == 2
            assert len(setup["calls"]) == 1



def _no_frixion_money_moov_models_open_banking_transaction_direct_setup(mockres):
    runner.load_env_local()

    calls = []

    env = runner.env_override({
        "NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_OPEN_BANKING_TRANSACTION_ENTID": {},
        "NOFRIXION_TEST_LIVE": "FALSE",
        "NOFRIXION_APIKEY": "NONE",
    })

    live = env.get("NOFRIXION_TEST_LIVE") == "TRUE"

    if live:
        merged_opts = {
            "apikey": env.get("NOFRIXION_APIKEY"),
        }
        client = NofrixionSDK(merged_opts)
        return {
            "client": client,
            "calls": calls,
            "live": True,
            "idmap": {},
        }

    def mock_fetch(url, init):
        calls.append({"url": url, "init": init})
        return {
            "status": 200,
            "statusText": "OK",
            "headers": {},
            "json": lambda: mockres if mockres is not None else {"id": "direct01"},
            "body": "mock",
        }, None

    client = NofrixionSDK({
        "base": "http://localhost:8080",
        "system": {
            "fetch": mock_fetch,
        },
    })

    return {
        "client": client,
        "calls": calls,
        "live": False,
        "idmap": {},
    }
