#!/usr/bin/python
#
# Copyright 2019 Polyaxon, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: v1/project.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from protoc_gen_swagger.options import annotations_pb2 as protoc__gen__swagger_dot_options_dot_annotations__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='v1/project.proto',
  package='v1',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x10v1/project.proto\x12\x02v1\x1a\x1cgoogle/api/annotations.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a,protoc-gen-swagger/options/annotations.proto\"\xcf\x01\n\x07Project\x12\x0c\n\x04uuid\x18\x01 \x01(\t\x12\x0c\n\x04user\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x04 \x01(\t\x12.\n\ncreated_at\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12.\n\nupdated_at\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x11\n\tis_public\x18\x07 \x01(\t\x12\x12\n\nbookmarked\x18\x08 \x01(\x08\"A\n\x12ProjectBodyRequest\x12\r\n\x05owner\x18\x01 \x01(\t\x12\x1c\n\x07project\x18\x02 \x01(\x0b\x32\x0b.v1.Project\"c\n\x14ListProjectsResponse\x12\r\n\x05\x63ount\x18\x01 \x01(\x05\x12\x1c\n\x07results\x18\x02 \x03(\x0b\x32\x0b.v1.Project\x12\x10\n\x08previous\x18\x03 \x01(\t\x12\x0c\n\x04next\x18\x04 \x01(\tb\x06proto3')
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,google_dot_protobuf_dot_struct__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,protoc__gen__swagger_dot_options_dot_annotations__pb2.DESCRIPTOR,])




_PROJECT = _descriptor.Descriptor(
  name='Project',
  full_name='v1.Project',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='uuid', full_name='v1.Project.uuid', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='user', full_name='v1.Project.user', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='v1.Project.name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='v1.Project.description', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='created_at', full_name='v1.Project.created_at', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='updated_at', full_name='v1.Project.updated_at', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='is_public', full_name='v1.Project.is_public', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='bookmarked', full_name='v1.Project.bookmarked', index=7,
      number=8, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=164,
  serialized_end=371,
)


_PROJECTBODYREQUEST = _descriptor.Descriptor(
  name='ProjectBodyRequest',
  full_name='v1.ProjectBodyRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='owner', full_name='v1.ProjectBodyRequest.owner', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='project', full_name='v1.ProjectBodyRequest.project', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=373,
  serialized_end=438,
)


_LISTPROJECTSRESPONSE = _descriptor.Descriptor(
  name='ListProjectsResponse',
  full_name='v1.ListProjectsResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='count', full_name='v1.ListProjectsResponse.count', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='results', full_name='v1.ListProjectsResponse.results', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='previous', full_name='v1.ListProjectsResponse.previous', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='next', full_name='v1.ListProjectsResponse.next', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=440,
  serialized_end=539,
)

_PROJECT.fields_by_name['created_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_PROJECT.fields_by_name['updated_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_PROJECTBODYREQUEST.fields_by_name['project'].message_type = _PROJECT
_LISTPROJECTSRESPONSE.fields_by_name['results'].message_type = _PROJECT
DESCRIPTOR.message_types_by_name['Project'] = _PROJECT
DESCRIPTOR.message_types_by_name['ProjectBodyRequest'] = _PROJECTBODYREQUEST
DESCRIPTOR.message_types_by_name['ListProjectsResponse'] = _LISTPROJECTSRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Project = _reflection.GeneratedProtocolMessageType('Project', (_message.Message,), {
  'DESCRIPTOR' : _PROJECT,
  '__module__' : 'v1.project_pb2'
  # @@protoc_insertion_point(class_scope:v1.Project)
  })
_sym_db.RegisterMessage(Project)

ProjectBodyRequest = _reflection.GeneratedProtocolMessageType('ProjectBodyRequest', (_message.Message,), {
  'DESCRIPTOR' : _PROJECTBODYREQUEST,
  '__module__' : 'v1.project_pb2'
  # @@protoc_insertion_point(class_scope:v1.ProjectBodyRequest)
  })
_sym_db.RegisterMessage(ProjectBodyRequest)

ListProjectsResponse = _reflection.GeneratedProtocolMessageType('ListProjectsResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTPROJECTSRESPONSE,
  '__module__' : 'v1.project_pb2'
  # @@protoc_insertion_point(class_scope:v1.ListProjectsResponse)
  })
_sym_db.RegisterMessage(ListProjectsResponse)


# @@protoc_insertion_point(module_scope)
