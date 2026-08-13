# CustomsWindow SDK exists test

import pytest
from customswindow_sdk import CustomsWindowSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = CustomsWindowSDK.test(None, None)
        assert testsdk is not None
