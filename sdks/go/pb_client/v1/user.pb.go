// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/user.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Auth specification
type User struct {
	// Username
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// Email
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	// Role
	Role                 string   `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_644e625c59382b7a, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "v1.User")
}

func init() { proto.RegisterFile("v1/user.proto", fileDescriptor_644e625c59382b7a) }

var fileDescriptor_644e625c59382b7a = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x33, 0xd4, 0x2f,
	0x2d, 0x4e, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x33, 0x94, 0x92, 0x49,
	0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49,
	0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0xa8, 0x90, 0xd2, 0x01, 0x53, 0xc9, 0xba, 0xe9, 0xa9, 0x79,
	0xba, 0xc5, 0xe5, 0x89, 0xe9, 0xe9, 0xa9, 0x45, 0xfa, 0xf9, 0x05, 0x60, 0x15, 0x98, 0xaa, 0x95,
	0x7c, 0xb8, 0x58, 0x42, 0x8b, 0x53, 0x8b, 0x84, 0xa4, 0xb8, 0x38, 0x40, 0xb6, 0xe4, 0x25, 0xe6,
	0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xf9, 0x42, 0x22, 0x5c, 0xac, 0xa9, 0xb9,
	0x89, 0x99, 0x39, 0x12, 0x4c, 0x60, 0x09, 0x08, 0x47, 0x48, 0x88, 0x8b, 0xa5, 0x28, 0x3f, 0x27,
	0x55, 0x82, 0x19, 0x2c, 0x08, 0x66, 0x27, 0xb1, 0x81, 0x0d, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x86, 0x1a, 0x17, 0x6b, 0xb5, 0x00, 0x00, 0x00,
}
