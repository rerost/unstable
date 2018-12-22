// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config.proto

package com_github_rerost_unstable

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Config struct {
	Interval             int64               `protobuf:"varint,1,opt,name=interval,proto3" json:"interval,omitempty"`
	ServerError          bool                `protobuf:"varint,2,opt,name=server_error,json=serverError,proto3" json:"server_error,omitempty"`
	SlowResponseOption   *SlowResponseOption `protobuf:"bytes,3,opt,name=slow_response_option,json=slowResponseOption,proto3" json:"slow_response_option,omitempty"`
	ServerErrorOption    *ServerErrorOption  `protobuf:"bytes,4,opt,name=server_error_option,json=serverErrorOption,proto3" json:"server_error_option,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_700606528256b7d5, []int{0}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (dst *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(dst, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetInterval() int64 {
	if m != nil {
		return m.Interval
	}
	return 0
}

func (m *Config) GetServerError() bool {
	if m != nil {
		return m.ServerError
	}
	return false
}

func (m *Config) GetSlowResponseOption() *SlowResponseOption {
	if m != nil {
		return m.SlowResponseOption
	}
	return nil
}

func (m *Config) GetServerErrorOption() *ServerErrorOption {
	if m != nil {
		return m.ServerErrorOption
	}
	return nil
}

type SlowResponseOption struct {
	Enable               bool     `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	Time                 int64    `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SlowResponseOption) Reset()         { *m = SlowResponseOption{} }
func (m *SlowResponseOption) String() string { return proto.CompactTextString(m) }
func (*SlowResponseOption) ProtoMessage()    {}
func (*SlowResponseOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_700606528256b7d5, []int{1}
}
func (m *SlowResponseOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SlowResponseOption.Unmarshal(m, b)
}
func (m *SlowResponseOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SlowResponseOption.Marshal(b, m, deterministic)
}
func (dst *SlowResponseOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SlowResponseOption.Merge(dst, src)
}
func (m *SlowResponseOption) XXX_Size() int {
	return xxx_messageInfo_SlowResponseOption.Size(m)
}
func (m *SlowResponseOption) XXX_DiscardUnknown() {
	xxx_messageInfo_SlowResponseOption.DiscardUnknown(m)
}

var xxx_messageInfo_SlowResponseOption proto.InternalMessageInfo

func (m *SlowResponseOption) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *SlowResponseOption) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type ServerErrorOption struct {
	Enable               bool     `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerErrorOption) Reset()         { *m = ServerErrorOption{} }
func (m *ServerErrorOption) String() string { return proto.CompactTextString(m) }
func (*ServerErrorOption) ProtoMessage()    {}
func (*ServerErrorOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_700606528256b7d5, []int{2}
}
func (m *ServerErrorOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerErrorOption.Unmarshal(m, b)
}
func (m *ServerErrorOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerErrorOption.Marshal(b, m, deterministic)
}
func (dst *ServerErrorOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerErrorOption.Merge(dst, src)
}
func (m *ServerErrorOption) XXX_Size() int {
	return xxx_messageInfo_ServerErrorOption.Size(m)
}
func (m *ServerErrorOption) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerErrorOption.DiscardUnknown(m)
}

var xxx_messageInfo_ServerErrorOption proto.InternalMessageInfo

func (m *ServerErrorOption) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func init() {
	proto.RegisterType((*Config)(nil), "com.github.rerost.unstable.Config")
	proto.RegisterType((*SlowResponseOption)(nil), "com.github.rerost.unstable.SlowResponseOption")
	proto.RegisterType((*ServerErrorOption)(nil), "com.github.rerost.unstable.ServerErrorOption")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor_config_700606528256b7d5) }

var fileDescriptor_config_700606528256b7d5 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x49, 0x53, 0x42, 0x98, 0xf6, 0xd2, 0x51, 0x24, 0xf4, 0x14, 0x73, 0x0a, 0x88, 0x7b,
	0xd0, 0x2f, 0x20, 0x88, 0x67, 0x61, 0x3d, 0x4b, 0x6c, 0xca, 0x58, 0x17, 0xd2, 0x9d, 0x30, 0xb3,
	0x6d, 0xbf, 0x82, 0x1f, 0x5b, 0xba, 0xc6, 0x3f, 0x18, 0xea, 0x6d, 0xe7, 0xb1, 0xf3, 0xde, 0x8f,
	0x37, 0x30, 0x5f, 0xb3, 0x7f, 0x75, 0x1b, 0xd3, 0x0b, 0x07, 0xc6, 0xe5, 0x9a, 0xb7, 0x66, 0xe3,
	0xc2, 0xdb, 0xae, 0x35, 0x42, 0xc2, 0x1a, 0xcc, 0xce, 0x6b, 0x58, 0xb5, 0x1d, 0x55, 0xef, 0x13,
	0xc8, 0xee, 0xe3, 0x67, 0x5c, 0x42, 0xee, 0x7c, 0x20, 0xd9, 0xaf, 0xba, 0x22, 0x29, 0x93, 0x3a,
	0xb5, 0xdf, 0x33, 0x5e, 0xc2, 0x5c, 0x49, 0xf6, 0x24, 0x0d, 0x89, 0xb0, 0x14, 0x93, 0x32, 0xa9,
	0x73, 0x3b, 0xfb, 0xd4, 0x1e, 0x8e, 0x12, 0xbe, 0xc0, 0xb9, 0x76, 0x7c, 0x68, 0x84, 0xb4, 0x67,
	0xaf, 0xd4, 0x70, 0x1f, 0x1c, 0xfb, 0x22, 0x2d, 0x93, 0x7a, 0x76, 0x63, 0xcc, 0x69, 0x08, 0xf3,
	0xd4, 0xf1, 0xc1, 0x0e, 0x6b, 0x8f, 0x71, 0xcb, 0xa2, 0x8e, 0x34, 0x7c, 0x86, 0xb3, 0xdf, 0x10,
	0x5f, 0x01, 0xd3, 0x18, 0x70, 0xfd, 0x6f, 0xc0, 0x0f, 0xe7, 0xe0, 0xbf, 0xd0, 0xbf, 0x52, 0x75,
	0x07, 0x38, 0x06, 0xc1, 0x0b, 0xc8, 0xc8, 0x1f, 0x4d, 0x62, 0x27, 0xb9, 0x1d, 0x26, 0x44, 0x98,
	0x06, 0xb7, 0xa5, 0xd8, 0x44, 0x6a, 0xe3, 0xbb, 0xba, 0x82, 0xc5, 0x28, 0xe9, 0x94, 0x41, 0x9b,
	0xc5, 0xe3, 0xdc, 0x7e, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x2d, 0xf1, 0x6a, 0xac, 0x01, 0x00,
	0x00,
}
