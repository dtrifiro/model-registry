"""Artifact types for model registry.

Artifacts represent pieces of data.
This could be datasets, models, metrics, or any other piece of data produced or consumed by an
execution, such as an experiment run.

Those types are used to map between proto types based on artifacts and Python objects.

TODO:
    * Move part of the description to API Reference docs (#120).
"""

from __future__ import annotations

from enum import Enum, unique
from typing import Optional
from uuid import uuid4

from attrs import define, field
from ml_metadata.proto import Artifact

from .base import Prefixable, ProtoBase


@unique
class ArtifactState(Enum):
    """State of an artifact."""

    UNKNOWN = Artifact.UNKNOWN
    PENDING = Artifact.PENDING
    LIVE = Artifact.LIVE
    MARKED_FOR_DELETION = Artifact.MARKED_FOR_DELETION
    DELETED = Artifact.DELETED
    ABANDONED = Artifact.ABANDONED
    REFERENCE = Artifact.REFERENCE


@define(slots=False)
class BaseArtifact(ProtoBase):
    """Abstract base class for all artifacts.

    Attributes:
        name (str): Name of the artifact.
        uri (str): URI of the artifact.
        state (ArtifactState): State of the artifact.
    """

    name: str
    uri: str
    state: ArtifactState = field(init=False, default=ArtifactState.UNKNOWN)

    @classmethod
    def get_proto_type(cls) -> type[Artifact]:
        return Artifact

    def map(self) -> Artifact:
        mlmd_obj = super().map()
        mlmd_obj.uri = self.uri
        mlmd_obj.state = ArtifactState[self.state.name].value
        return mlmd_obj

    @classmethod
    def unmap(cls, mlmd_obj: Artifact) -> BaseArtifact:
        py_obj = super().unmap(mlmd_obj)
        assert isinstance(
            py_obj, BaseArtifact
        ), f"Expected BaseArtifact, got {type(py_obj)}"
        py_obj.uri = mlmd_obj.uri
        py_obj.state = ArtifactState(mlmd_obj.state)
        return py_obj


@define(slots=False, auto_attribs=True)
class ModelArtifact(BaseArtifact, Prefixable):
    """Represents a Model.

    Attributes:
        name (str): Name of the model.
        uri (str): URI of the model.
        description (str, optional): Description of the object.
        external_id (str, optional): Customizable ID. Has to be unique among instances of the same type.
        model_format_name (str, optional): Name of the model format.
        model_format_version (str, optional): Version of the model format.
        storage_key (str, optional): Storage key of the model.
        storage_path (str, optional): Storage path of the model.
        service_account_name (str, optional): Service account name of the model.
    """

    # TODO: this could be an enum of valid formats
    model_format_name: Optional[str] = field(kw_only=True, default=None)
    model_format_version: Optional[str] = field(kw_only=True, default=None)
    storage_key: Optional[str] = field(kw_only=True, default=None)
    storage_path: Optional[str] = field(kw_only=True, default=None)
    service_account_name: Optional[str] = field(kw_only=True, default=None)

    @property
    def mlmd_name_prefix(self) -> str:
        return uuid4().hex

    def map(self) -> Artifact:
        mlmd_obj = super().map()
        props = {
            "modelFormatName": self.model_format_name,
            "modelFormatVersion": self.model_format_version,
            "storageKey": self.storage_key,
            "storagePath": self.storage_path,
            "serviceAccountName": self.service_account_name,
        }
        self._map_props(props, mlmd_obj.properties)
        return mlmd_obj

    @classmethod
    def unmap(cls, mlmd_obj: Artifact) -> ModelArtifact:
        py_obj = super().unmap(mlmd_obj)
        assert isinstance(
            py_obj, ModelArtifact
        ), f"Expected ModelArtifact, got {type(py_obj)}"
        py_obj.model_format_name = mlmd_obj.properties["modelFormatName"].string_value
        py_obj.model_format_version = mlmd_obj.properties[
            "modelFormatVersion"
        ].string_value
        py_obj.storage_key = mlmd_obj.properties["storageKey"].string_value
        py_obj.storage_path = mlmd_obj.properties["storagePath"].string_value
        py_obj.service_account_name = mlmd_obj.properties[
            "serviceAccountName"
        ].string_value
        return py_obj
