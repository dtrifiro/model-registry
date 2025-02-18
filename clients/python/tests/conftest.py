from ml_metadata.proto import ConnectionConfig
from model_registry.store.wrapper import MLMDStore
from pytest import fixture


@fixture
def store_wrapper() -> MLMDStore:
    config = ConnectionConfig()
    config.fake_database.SetInParent()
    return MLMDStore(config)
