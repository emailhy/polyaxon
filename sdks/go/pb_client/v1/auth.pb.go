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
// source: v1/auth.proto

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
type Auth struct {
	// UUID
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Auth) Reset()         { *m = Auth{} }
func (m *Auth) String() string { return proto.CompactTextString(m) }
func (*Auth) ProtoMessage()    {}
func (*Auth) Descriptor() ([]byte, []int) {
	return fileDescriptor_15263a302b6905f5, []int{0}
}

func (m *Auth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Auth.Unmarshal(m, b)
}
func (m *Auth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Auth.Marshal(b, m, deterministic)
}
func (m *Auth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auth.Merge(m, src)
}
func (m *Auth) XXX_Size() int {
	return xxx_messageInfo_Auth.Size(m)
}
func (m *Auth) XXX_DiscardUnknown() {
	xxx_messageInfo_Auth.DiscardUnknown(m)
}

var xxx_messageInfo_Auth proto.InternalMessageInfo

func (m *Auth) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// Login
type CredsBodyRequest struct {
	// User username or email
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// Project where the experiement will be assigned
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CredsBodyRequest) Reset()         { *m = CredsBodyRequest{} }
func (m *CredsBodyRequest) String() string { return proto.CompactTextString(m) }
func (*CredsBodyRequest) ProtoMessage()    {}
func (*CredsBodyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15263a302b6905f5, []int{1}
}

func (m *CredsBodyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CredsBodyRequest.Unmarshal(m, b)
}
func (m *CredsBodyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CredsBodyRequest.Marshal(b, m, deterministic)
}
func (m *CredsBodyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredsBodyRequest.Merge(m, src)
}
func (m *CredsBodyRequest) XXX_Size() int {
	return xxx_messageInfo_CredsBodyRequest.Size(m)
}
func (m *CredsBodyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CredsBodyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CredsBodyRequest proto.InternalMessageInfo

func (m *CredsBodyRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CredsBodyRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*Auth)(nil), "v1.Auth")
	proto.RegisterType((*CredsBodyRequest)(nil), "v1.CredsBodyRequest")
}

func init() { proto.RegisterFile("v1/auth.proto", fileDescriptor_15263a302b6905f5) }

var fileDescriptor_15263a302b6905f5 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0xcd, 0xc1, 0x0a, 0x82, 0x40,
	0x10, 0xc6, 0x71, 0x94, 0x8a, 0x5a, 0x08, 0x42, 0x3a, 0x88, 0x78, 0x08, 0x4f, 0x1d, 0xb2, 0x45,
	0x7a, 0x82, 0xea, 0xd6, 0xd1, 0x37, 0xd8, 0x72, 0x58, 0xa5, 0xda, 0xd9, 0x76, 0x66, 0x95, 0xde,
	0x3e, 0xd2, 0xea, 0xd2, 0x69, 0xf8, 0xf3, 0xfd, 0x60, 0xc4, 0xbc, 0x2d, 0xa4, 0xf2, 0x5c, 0x6f,
	0xad, 0x43, 0xc6, 0x28, 0x6c, 0x8b, 0x24, 0xd5, 0x88, 0xfa, 0x06, 0x52, 0xd9, 0x46, 0x2a, 0x63,
	0x90, 0x15, 0x37, 0x68, 0x68, 0x10, 0xc9, 0xa6, 0x3f, 0x97, 0x5c, 0x83, 0xc9, 0xa9, 0x53, 0x5a,
	0x83, 0x93, 0x68, 0x7b, 0xf1, 0xaf, 0xb3, 0x54, 0x8c, 0xf6, 0x9e, 0xeb, 0x68, 0x29, 0xc6, 0x8c,
	0x57, 0x30, 0x71, 0xb0, 0x0a, 0xd6, 0xb3, 0x72, 0x88, 0xec, 0x24, 0x16, 0x47, 0x07, 0x15, 0x1d,
	0xb0, 0x7a, 0x96, 0xf0, 0xf0, 0x40, 0x1c, 0x25, 0x62, 0xea, 0x09, 0x9c, 0x51, 0x77, 0xf8, 0xe0,
	0x5f, 0xbf, 0x37, 0xab, 0x88, 0x3a, 0x74, 0x55, 0x1c, 0x0e, 0xdb, 0xb7, 0xcf, 0x93, 0xfe, 0xe1,
	0xee, 0x15, 0x00, 0x00, 0xff, 0xff, 0x76, 0xf2, 0x5c, 0x05, 0xd1, 0x00, 0x00, 0x00,
}
