// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sample.proto

package com_github_rerost_unstable_example_api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetSampleResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSampleResponse) Reset()         { *m = GetSampleResponse{} }
func (m *GetSampleResponse) String() string { return proto.CompactTextString(m) }
func (*GetSampleResponse) ProtoMessage()    {}
func (*GetSampleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_sample_3957a6d61e90482d, []int{0}
}
func (m *GetSampleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSampleResponse.Unmarshal(m, b)
}
func (m *GetSampleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSampleResponse.Marshal(b, m, deterministic)
}
func (dst *GetSampleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSampleResponse.Merge(dst, src)
}
func (m *GetSampleResponse) XXX_Size() int {
	return xxx_messageInfo_GetSampleResponse.Size(m)
}
func (m *GetSampleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSampleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSampleResponse proto.InternalMessageInfo

func (m *GetSampleResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*GetSampleResponse)(nil), "com.github.rerost.unstable.example.api.GetSampleResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SampleServiceClient is the client API for SampleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SampleServiceClient interface {
	GetSample(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetSampleResponse, error)
}

type sampleServiceClient struct {
	cc *grpc.ClientConn
}

func NewSampleServiceClient(cc *grpc.ClientConn) SampleServiceClient {
	return &sampleServiceClient{cc}
}

func (c *sampleServiceClient) GetSample(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetSampleResponse, error) {
	out := new(GetSampleResponse)
	err := c.cc.Invoke(ctx, "/com.github.rerost.unstable.example.api.SampleService/GetSample", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SampleServiceServer is the server API for SampleService service.
type SampleServiceServer interface {
	GetSample(context.Context, *empty.Empty) (*GetSampleResponse, error)
}

func RegisterSampleServiceServer(s *grpc.Server, srv SampleServiceServer) {
	s.RegisterService(&_SampleService_serviceDesc, srv)
}

func _SampleService_GetSample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleServiceServer).GetSample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.rerost.unstable.example.api.SampleService/GetSample",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleServiceServer).GetSample(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _SampleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.github.rerost.unstable.example.api.SampleService",
	HandlerType: (*SampleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSample",
			Handler:    _SampleService_GetSample_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sample.proto",
}

func init() { proto.RegisterFile("sample.proto", fileDescriptor_sample_3957a6d61e90482d) }

var fileDescriptor_sample_3957a6d61e90482d = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0xcc, 0x2d,
	0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x4b, 0xce, 0xcf, 0xd5, 0x4b, 0xcf,
	0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x2b, 0x4a, 0x2d, 0xca, 0x2f, 0x2e, 0xd1, 0x2b, 0xcd, 0x2b, 0x2e,
	0x49, 0x4c, 0xca, 0x49, 0xd5, 0x4b, 0xad, 0x80, 0xa8, 0x4c, 0x2c, 0xc8, 0x94, 0x92, 0x4e, 0xcf,
	0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0xeb, 0x4a, 0x2a, 0x4d, 0xd3, 0x4f, 0xcd, 0x2d, 0x28, 0xa9,
	0x84, 0x18, 0xa2, 0xa4, 0xcb, 0x25, 0xe8, 0x9e, 0x5a, 0x12, 0x0c, 0x56, 0x1d, 0x94, 0x5a, 0x5c,
	0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc1, 0xc5, 0x9e, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x9e, 0x2a,
	0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe3, 0x1a, 0x15, 0x72, 0xf1, 0x42, 0xd4, 0x06, 0xa7,
	0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0x25, 0x70, 0x71, 0xc2, 0xf5, 0x0b, 0x89, 0xe9, 0x41, 0xac,
	0xd2, 0x83, 0x59, 0xa5, 0xe7, 0x0a, 0xb2, 0x4a, 0xca, 0x52, 0x8f, 0x38, 0xa7, 0xea, 0x61, 0x38,
	0x45, 0x89, 0x21, 0x89, 0x0d, 0x6c, 0x98, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x36, 0x29,
	0xd5, 0xfd, 0x00, 0x00, 0x00,
}