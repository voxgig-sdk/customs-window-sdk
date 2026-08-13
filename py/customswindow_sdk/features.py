# CustomsWindow SDK feature factory

from customswindow_sdk.feature.base_feature import CustomsWindowBaseFeature
from customswindow_sdk.feature.test_feature import CustomsWindowTestFeature


def _make_feature(name):
    features = {
        "base": lambda: CustomsWindowBaseFeature(),
        "test": lambda: CustomsWindowTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
