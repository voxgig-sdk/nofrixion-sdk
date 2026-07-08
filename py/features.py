# Nofrixion SDK feature factory

from feature.base_feature import NofrixionBaseFeature
from feature.test_feature import NofrixionTestFeature


def _make_feature(name):
    features = {
        "base": lambda: NofrixionBaseFeature(),
        "test": lambda: NofrixionTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
