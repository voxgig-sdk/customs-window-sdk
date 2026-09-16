# CustomsWindow SDK feature factory

from customswindow_sdk.feature.base_feature import CustomsWindowBaseFeature
from customswindow_sdk.feature.debug_feature import CustomsWindowDebugFeature
from customswindow_sdk.feature.idempotency_feature import CustomsWindowIdempotencyFeature
from customswindow_sdk.feature.metrics_feature import CustomsWindowMetricsFeature
from customswindow_sdk.feature.paging_feature import CustomsWindowPagingFeature
from customswindow_sdk.feature.ratelimit_feature import CustomsWindowRatelimitFeature
from customswindow_sdk.feature.retry_feature import CustomsWindowRetryFeature
from customswindow_sdk.feature.test_feature import CustomsWindowTestFeature
from customswindow_sdk.feature.timeout_feature import CustomsWindowTimeoutFeature


_FEATURES = {
    "base": lambda: CustomsWindowBaseFeature(),
    "debug": lambda: CustomsWindowDebugFeature(),
    "idempotency": lambda: CustomsWindowIdempotencyFeature(),
    "metrics": lambda: CustomsWindowMetricsFeature(),
    "paging": lambda: CustomsWindowPagingFeature(),
    "ratelimit": lambda: CustomsWindowRatelimitFeature(),
    "retry": lambda: CustomsWindowRetryFeature(),
    "test": lambda: CustomsWindowTestFeature(),
    "timeout": lambda: CustomsWindowTimeoutFeature(),
}


def _make_feature(name):
    factory = _FEATURES.get(name)
    if factory is not None:
        return factory()
    return _FEATURES["base"]()


# True when this SDK was generated with the named feature class - the
# constructor's tolerance for extend-carried features reads this (an
# active name with no generated class must not become a BaseFeature
# stray when an extend instance carries it).
def _has_feature(name):
    return name in _FEATURES
