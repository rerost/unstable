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
	Interval             uint64              `protobuf:"varint,1,opt,name=interval,proto3" json:"interval,omitempty"`
	SlowResponseOption   *SlowResponseOption `protobuf:"bytes,2,opt,name=slow_response_option,json=slowResponseOption,proto3" json:"slow_response_option,omitempty"`
	ServerErrorOption    *ServerErrorOption  `protobuf:"bytes,3,opt,name=server_error_option,json=serverErrorOption,proto3" json:"server_error_option,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_c69b1ccf7898c492, []int{0}
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

func (m *Config) GetInterval() uint64 {
	if m != nil {
		return m.Interval
	}
	return 0
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
	Time                 uint64   `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SlowResponseOption) Reset()         { *m = SlowResponseOption{} }
func (m *SlowResponseOption) String() string { return proto.CompactTextString(m) }
func (*SlowResponseOption) ProtoMessage()    {}
func (*SlowResponseOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_c69b1ccf7898c492, []int{1}
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

func (m *SlowResponseOption) GetTime() uint64 {
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
	return fileDescriptor_config_c69b1ccf7898c492, []int{2}
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

func init() { proto.RegisterFile("config.proto", fileDescriptor_config_c69b1ccf7898c492) }

var fileDescriptor_config_c69b1ccf7898c492 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x89, 0x2e, 0x4b, 0x19, 0xbd, 0x74, 0x14, 0x29, 0x3d, 0x95, 0x3d, 0x15, 0xc4, 0x1c,
	0xf4, 0x05, 0x04, 0xf1, 0x2c, 0xc4, 0xb3, 0xac, 0xdd, 0x32, 0xd6, 0x40, 0x9a, 0x59, 0x66, 0xd2,
	0xf6, 0x89, 0x7d, 0x0f, 0xe9, 0xb0, 0x7a, 0x70, 0xa9, 0xb7, 0xcc, 0x90, 0xef, 0xff, 0xf2, 0x07,
	0x2e, 0xd7, 0x9c, 0x3f, 0xe2, 0xc6, 0xf7, 0xc2, 0x85, 0x71, 0xbe, 0xe6, 0xad, 0xdf, 0xc4, 0xf2,
	0xb9, 0xeb, 0xbc, 0x90, 0xb0, 0x16, 0xbf, 0xcb, 0x5a, 0x56, 0x5d, 0xa2, 0xe6, 0xcb, 0x41, 0xfd,
	0x64, 0x97, 0x71, 0x0e, 0x93, 0x98, 0x0b, 0xc9, 0x7e, 0x95, 0x66, 0x6e, 0xe1, 0x96, 0x55, 0xf8,
	0x9d, 0xf1, 0x1d, 0xae, 0x35, 0xf1, 0xa1, 0x15, 0xd2, 0x9e, 0xb3, 0x52, 0xcb, 0x7d, 0x89, 0x9c,
	0x67, 0x67, 0x0b, 0xb7, 0xbc, 0xb8, 0xf7, 0xfe, 0xb4, 0xc1, 0xbf, 0x26, 0x3e, 0x84, 0x01, 0x7b,
	0x31, 0x2a, 0xa0, 0x8e, 0x76, 0xf8, 0x06, 0x57, 0x4a, 0xb2, 0x27, 0x69, 0x49, 0x84, 0xe5, 0x47,
	0x70, 0x6e, 0x82, 0xbb, 0x7f, 0x05, 0x86, 0x3d, 0x1f, 0xa9, 0x21, 0x7f, 0xaa, 0x7f, 0x57, 0xcd,
	0x23, 0xe0, 0xf8, 0x21, 0x78, 0x03, 0x35, 0xe5, 0x63, 0x88, 0x15, 0x9e, 0x84, 0x61, 0x42, 0x84,
	0xaa, 0xc4, 0x2d, 0x59, 0xbd, 0x2a, 0xd8, 0xb9, 0xb9, 0x85, 0xe9, 0xc8, 0x74, 0x2a, 0xa0, 0xab,
	0xed, 0xe7, 0x1f, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x89, 0x22, 0xcf, 0x01, 0x89, 0x01, 0x00,
	0x00,
}
