# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: user.proto
# Protobuf Python Version: 5.29.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    29,
    0,
    '',
    'user.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\nuser.proto\x1a\x1bgoogle/protobuf/empty.proto\"2\n\x08PageInfo\x12\x13\n\x0bpage_number\x18\x01 \x01(\r\x12\x11\n\tpage_size\x18\x02 \x01(\r\"D\n\x0e\x43reateUserInfo\x12\x0e\n\x06mobile\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\x12\x10\n\x08nickname\x18\x03 \x01(\t\"\x1f\n\rMobileRequest\x12\x0e\n\x06mobile\x18\x01 \x01(\t\"\x17\n\tIdRequest\x12\n\n\x02id\x18\x01 \x01(\x05\"P\n\x0eUpdateUserInfo\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x10\n\x08nickName\x18\x02 \x01(\t\x12\x0e\n\x06gender\x18\x03 \x01(\t\x12\x10\n\x08\x62irthday\x18\x04 \x01(\x04\"B\n\x10UserListResponse\x12\r\n\x05total\x18\x01 \x01(\x05\x12\x1f\n\x04\x64\x61ta\x18\x02 \x03(\x0b\x32\x11.UserInfoResponse\"\x82\x01\n\x10UserInfoResponse\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x10\n\x08password\x18\x02 \x01(\t\x12\x0e\n\x06mobile\x18\x03 \x01(\t\x12\x10\n\x08\x62irthday\x18\x04 \x01(\t\x12\x0e\n\x06gender\x18\x05 \x01(\t\x12\x10\n\x08nickname\x18\x06 \x01(\t\x12\x0c\n\x04role\x18\x07 \x01(\t2\x80\x02\n\x04User\x12+\n\x0bGetUserList\x12\t.PageInfo\x1a\x11.UserListResponse\x12\x34\n\x0fGetUserByMobile\x12\x0e.MobileRequest\x1a\x11.UserInfoResponse\x12,\n\x0bGetUserById\x12\n.IdRequest\x1a\x11.UserInfoResponse\x12\x30\n\nCreateUser\x12\x0f.CreateUserInfo\x1a\x11.UserInfoResponse\x12\x35\n\nUpdateUser\x12\x0f.UpdateUserInfo\x1a\x16.google.protobuf.EmptyB\tZ\x07.;protob\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'user_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\007.;proto'
  _globals['_PAGEINFO']._serialized_start=43
  _globals['_PAGEINFO']._serialized_end=93
  _globals['_CREATEUSERINFO']._serialized_start=95
  _globals['_CREATEUSERINFO']._serialized_end=163
  _globals['_MOBILEREQUEST']._serialized_start=165
  _globals['_MOBILEREQUEST']._serialized_end=196
  _globals['_IDREQUEST']._serialized_start=198
  _globals['_IDREQUEST']._serialized_end=221
  _globals['_UPDATEUSERINFO']._serialized_start=223
  _globals['_UPDATEUSERINFO']._serialized_end=303
  _globals['_USERLISTRESPONSE']._serialized_start=305
  _globals['_USERLISTRESPONSE']._serialized_end=371
  _globals['_USERINFORESPONSE']._serialized_start=374
  _globals['_USERINFORESPONSE']._serialized_end=504
  _globals['_USER']._serialized_start=507
  _globals['_USER']._serialized_end=763
# @@protoc_insertion_point(module_scope)
