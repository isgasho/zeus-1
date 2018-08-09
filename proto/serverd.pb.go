// Code generated by protoc-gen-go. DO NOT EDIT.
// source: serverd.proto

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 消息请求体
type ServerdRequest struct {
}

func (m *ServerdRequest) Reset()                    { *m = ServerdRequest{} }
func (m *ServerdRequest) String() string            { return proto1.CompactTextString(m) }
func (*ServerdRequest) ProtoMessage()               {}
func (*ServerdRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// 消息响应
type ServerdReply struct {
	Code    string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *ServerdReply) Reset()                    { *m = ServerdReply{} }
func (m *ServerdReply) String() string            { return proto1.CompactTextString(m) }
func (*ServerdReply) ProtoMessage()               {}
func (*ServerdReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ServerdReply) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *ServerdReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto1.RegisterType((*ServerdRequest)(nil), "proto.ServerdRequest")
	proto1.RegisterType((*ServerdReply)(nil), "proto.ServerdReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ServerdService service

type ServerdServiceClient interface {
	// 发送消息
	SayRestart(ctx context.Context, in *ServerdRequest, opts ...grpc.CallOption) (*ServerdReply, error)
}

type serverdServiceClient struct {
	cc *grpc.ClientConn
}

func NewServerdServiceClient(cc *grpc.ClientConn) ServerdServiceClient {
	return &serverdServiceClient{cc}
}

func (c *serverdServiceClient) SayRestart(ctx context.Context, in *ServerdRequest, opts ...grpc.CallOption) (*ServerdReply, error) {
	out := new(ServerdReply)
	err := grpc.Invoke(ctx, "/proto.ServerdService/SayRestart", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ServerdService service

type ServerdServiceServer interface {
	// 发送消息
	SayRestart(context.Context, *ServerdRequest) (*ServerdReply, error)
}

func RegisterServerdServiceServer(s *grpc.Server, srv ServerdServiceServer) {
	s.RegisterService(&_ServerdService_serviceDesc, srv)
}

func _ServerdService_SayRestart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerdServiceServer).SayRestart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ServerdService/SayRestart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerdServiceServer).SayRestart(ctx, req.(*ServerdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServerdService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ServerdService",
	HandlerType: (*ServerdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayRestart",
			Handler:    _ServerdService_SayRestart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "serverd.proto",
}

func init() { proto1.RegisterFile("serverd.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0x4a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x02, 0x5c,
	0x7c, 0xc1, 0x10, 0xf1, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x25, 0x1b, 0x2e, 0x1e, 0xb8,
	0x48, 0x41, 0x4e, 0xa5, 0x90, 0x10, 0x17, 0x4b, 0x72, 0x7e, 0x4a, 0xaa, 0x04, 0xa3, 0x02, 0xa3,
	0x06, 0x67, 0x10, 0x98, 0x2d, 0x24, 0xc1, 0xc5, 0x9e, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x9e, 0x2a,
	0xc1, 0x04, 0x16, 0x86, 0x71, 0x8d, 0x7c, 0xe0, 0xe6, 0x81, 0xa8, 0xcc, 0xe4, 0x54, 0x21, 0x2b,
	0x2e, 0xae, 0xe0, 0xc4, 0xca, 0xa0, 0xd4, 0xe2, 0x92, 0xc4, 0xa2, 0x12, 0x21, 0x51, 0x88, 0xf5,
	0x7a, 0xa8, 0x96, 0x4a, 0x09, 0xa3, 0x0b, 0x17, 0xe4, 0x54, 0x2a, 0x31, 0x38, 0x89, 0x72, 0xf1,
	0xe5, 0xa5, 0x96, 0xe8, 0x99, 0x1a, 0x67, 0x96, 0x40, 0x14, 0x2c, 0x62, 0x62, 0x8a, 0xf2, 0x4b,
	0x62, 0x03, 0xb3, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xf2, 0x3a, 0x39, 0xd3, 0x00,
	0x00, 0x00,
}