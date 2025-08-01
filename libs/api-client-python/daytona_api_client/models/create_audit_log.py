# coding: utf-8

"""
    Daytona

    Daytona AI platform API Docs

    The version of the OpenAPI document: 1.0
    Contact: support@daytona.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import pprint
import re  # noqa: F401
import json

from pydantic import BaseModel, ConfigDict, Field, StrictStr, field_validator
from typing import Any, ClassVar, Dict, List, Optional
from typing import Optional, Set
from typing_extensions import Self

class CreateAuditLog(BaseModel):
    """
    CreateAuditLog
    """ # noqa: E501
    actor_id: StrictStr = Field(alias="actorId")
    actor_email: StrictStr = Field(alias="actorEmail")
    organization_id: Optional[StrictStr] = Field(default=None, alias="organizationId")
    action: StrictStr
    target_type: Optional[StrictStr] = Field(default=None, alias="targetType")
    target_id: Optional[StrictStr] = Field(default=None, alias="targetId")
    additional_properties: Dict[str, Any] = {}
    __properties: ClassVar[List[str]] = ["actorId", "actorEmail", "organizationId", "action", "targetType", "targetId"]

    @field_validator('action')
    def action_validate_enum(cls, value):
        """Validates the enum"""
        if value not in set(['create', 'read', 'update', 'delete', 'login', 'set_default', 'update_role', 'update_assigned_roles', 'update_quota', 'suspend', 'unsuspend', 'accept', 'decline', 'link_account', 'unlink_account', 'leave_organization', 'regenerate_key_pair', 'update_scheduling', 'start', 'stop', 'replace_labels', 'create_backup', 'update_public_status', 'set_auto_stop_interval', 'set_auto_archive_interval', 'set_auto_delete_interval', 'archive', 'get_port_preview_url', 'set_general_status', 'activate', 'deactivate', 'toolbox_delete_file', 'toolbox_download_file', 'toolbox_create_folder', 'toolbox_move_file', 'toolbox_set_file_permissions', 'toolbox_replace_in_files', 'toolbox_upload_file', 'toolbox_bulk_upload_files', 'toolbox_git_add_files', 'toolbox_git_create_branch', 'toolbox_git_delete_branch', 'toolbox_git_clone_repository', 'toolbox_git_commit_changes', 'toolbox_git_pull_changes', 'toolbox_git_push_changes', 'toolbox_git_checkout_branch', 'toolbox_execute_command', 'toolbox_create_session', 'toolbox_session_execute_command', 'toolbox_delete_session', 'toolbox_computer_use_start', 'toolbox_computer_use_stop', 'toolbox_computer_use_restart_process']):
            raise ValueError("must be one of enum values ('create', 'read', 'update', 'delete', 'login', 'set_default', 'update_role', 'update_assigned_roles', 'update_quota', 'suspend', 'unsuspend', 'accept', 'decline', 'link_account', 'unlink_account', 'leave_organization', 'regenerate_key_pair', 'update_scheduling', 'start', 'stop', 'replace_labels', 'create_backup', 'update_public_status', 'set_auto_stop_interval', 'set_auto_archive_interval', 'set_auto_delete_interval', 'archive', 'get_port_preview_url', 'set_general_status', 'activate', 'deactivate', 'toolbox_delete_file', 'toolbox_download_file', 'toolbox_create_folder', 'toolbox_move_file', 'toolbox_set_file_permissions', 'toolbox_replace_in_files', 'toolbox_upload_file', 'toolbox_bulk_upload_files', 'toolbox_git_add_files', 'toolbox_git_create_branch', 'toolbox_git_delete_branch', 'toolbox_git_clone_repository', 'toolbox_git_commit_changes', 'toolbox_git_pull_changes', 'toolbox_git_push_changes', 'toolbox_git_checkout_branch', 'toolbox_execute_command', 'toolbox_create_session', 'toolbox_session_execute_command', 'toolbox_delete_session', 'toolbox_computer_use_start', 'toolbox_computer_use_stop', 'toolbox_computer_use_restart_process')")
        return value

    @field_validator('target_type')
    def target_type_validate_enum(cls, value):
        """Validates the enum"""
        if value is None:
            return value

        if value not in set(['api_key', 'organization', 'organization_invitation', 'organization_role', 'organization_user', 'docker_registry', 'runner', 'sandbox', 'snapshot', 'user', 'volume']):
            raise ValueError("must be one of enum values ('api_key', 'organization', 'organization_invitation', 'organization_role', 'organization_user', 'docker_registry', 'runner', 'sandbox', 'snapshot', 'user', 'volume')")
        return value

    model_config = ConfigDict(
        populate_by_name=True,
        validate_assignment=True,
        protected_namespaces=(),
    )


    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.model_dump(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        # TODO: pydantic v2: use .model_dump_json(by_alias=True, exclude_unset=True) instead
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> Optional[Self]:
        """Create an instance of CreateAuditLog from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        """Return the dictionary representation of the model using alias.

        This has the following differences from calling pydantic's
        `self.model_dump(by_alias=True)`:

        * `None` is only added to the output dict for nullable fields that
          were set at model initialization. Other fields with value `None`
          are ignored.
        * Fields in `self.additional_properties` are added to the output dict.
        """
        excluded_fields: Set[str] = set([
            "additional_properties",
        ])

        _dict = self.model_dump(
            by_alias=True,
            exclude=excluded_fields,
            exclude_none=True,
        )
        # puts key-value pairs in additional_properties in the top level
        if self.additional_properties is not None:
            for _key, _value in self.additional_properties.items():
                _dict[_key] = _value

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of CreateAuditLog from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "actorId": obj.get("actorId"),
            "actorEmail": obj.get("actorEmail"),
            "organizationId": obj.get("organizationId"),
            "action": obj.get("action"),
            "targetType": obj.get("targetType"),
            "targetId": obj.get("targetId")
        })
        # store additional fields in additional_properties
        for _key in obj.keys():
            if _key not in cls.__properties:
                _obj.additional_properties[_key] = obj.get(_key)

        return _obj


