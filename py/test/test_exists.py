# ProjectName SDK exists test

import pytest
from nofrixion_sdk import NofrixionSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = NofrixionSDK.test(None, None)
        assert testsdk is not None
