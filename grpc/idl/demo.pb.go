// Code generated by protoc-gen-go. DO NOT EDIT.
// source: demo.proto

/*
Package idl is a generated protocol buffer package.

It is generated from these files:
	demo.proto

It has these top-level messages:
	HelloRequest
	HelloResponse
*/
package idl

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type SERVICE_NAME int32

const (
	SERVICE_NAME_DEMO_SERVICE SERVICE_NAME = 0
)

var SERVICE_NAME_name = map[int32]string{
	0: "DEMO_SERVICE",
}
var SERVICE_NAME_value = map[string]int32{
	"DEMO_SERVICE": 0,
}

func (x SERVICE_NAME) String() string {
	return proto.EnumName(SERVICE_NAME_name, int32(x))
}
func (SERVICE_NAME) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HelloResponse_Code int32

const (
	HelloResponse_OK    HelloResponse_Code = 0
	HelloResponse_ERROR HelloResponse_Code = 1
)

var HelloResponse_Code_name = map[int32]string{
	0: "OK",
	1: "ERROR",
}
var HelloResponse_Code_value = map[string]int32{
	"OK":    0,
	"ERROR": 1,
}

func (x HelloResponse_Code) String() string {
	return proto.EnumName(HelloResponse_Code_name, int32(x))
}
func (HelloResponse_Code) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type HelloRequest struct {
	Greet string `protobuf:"bytes,1,opt,name=greet" json:"greet,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetGreet() string {
	if m != nil {
		return m.Greet
	}
	return ""
}

type HelloResponse struct {
	Code  HelloResponse_Code `protobuf:"varint,1,opt,name=code,enum=idl.HelloResponse_Code" json:"code,omitempty"`
	Msg   string             `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	Reply string             `protobuf:"bytes,3,opt,name=reply" json:"reply,omitempty"`
}

func (m *HelloResponse) Reset()                    { *m = HelloResponse{} }
func (m *HelloResponse) String() string            { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()               {}
func (*HelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloResponse) GetCode() HelloResponse_Code {
	if m != nil {
		return m.Code
	}
	return HelloResponse_OK
}

func (m *HelloResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *HelloResponse) GetReply() string {
	if m != nil {
		return m.Reply
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "idl.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "idl.HelloResponse")
	proto.RegisterEnum("idl.SERVICE_NAME", SERVICE_NAME_name, SERVICE_NAME_value)
	proto.RegisterEnum("idl.HelloResponse_Code", HelloResponse_Code_name, HelloResponse_Code_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DemoService service

type DemoServiceClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type demoServiceClient struct {
	cc *grpc.ClientConn
}

func NewDemoServiceClient(cc *grpc.ClientConn) DemoServiceClient {
	return &demoServiceClient{cc}
}

func (c *demoServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := grpc.Invoke(ctx, "/idl.DemoService/Hello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DemoService service

type DemoServiceServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
}

func RegisterDemoServiceServer(s *grpc.Server, srv DemoServiceServer) {
	s.RegisterService(&_DemoService_serviceDesc, srv)
}

func _DemoService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.DemoService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DemoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "idl.DemoService",
	HandlerType: (*DemoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _DemoService_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}

func init() { proto.RegisterFile("demo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x49, 0xcd, 0xcd,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x4c, 0xc9, 0x51, 0x52, 0xe1, 0xe2, 0xf1,
	0x48, 0xcd, 0xc9, 0xc9, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe1, 0x62, 0x4d,
	0x2f, 0x4a, 0x4d, 0x2d, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x94, 0xea, 0xb9,
	0x78, 0xa1, 0xaa, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xb4, 0xb9, 0x58, 0x92, 0xf3, 0x53,
	0x52, 0xc1, 0xaa, 0xf8, 0x8c, 0xc4, 0xf5, 0x32, 0x53, 0x72, 0xf4, 0x50, 0x54, 0xe8, 0x39, 0xe7,
	0xa7, 0xa4, 0x06, 0x81, 0x15, 0x09, 0x09, 0x70, 0x31, 0xe7, 0x16, 0xa7, 0x4b, 0x30, 0x81, 0x4d,
	0x04, 0x31, 0x41, 0xb6, 0x14, 0xa5, 0x16, 0xe4, 0x54, 0x4a, 0x30, 0x43, 0x6c, 0x01, 0x73, 0x94,
	0x24, 0xb9, 0x58, 0x40, 0xba, 0x84, 0xd8, 0xb8, 0x98, 0xfc, 0xbd, 0x05, 0x18, 0x84, 0x38, 0xb9,
	0x58, 0x5d, 0x83, 0x82, 0xfc, 0x83, 0x04, 0x18, 0xb5, 0x14, 0xb8, 0x78, 0x82, 0x5d, 0x83, 0xc2,
	0x3c, 0x9d, 0x5d, 0xe3, 0xfd, 0x1c, 0x7d, 0x5d, 0x85, 0x04, 0xb8, 0x78, 0x5c, 0x5c, 0x7d, 0xfd,
	0xe3, 0xa1, 0x82, 0x02, 0x0c, 0x46, 0xb6, 0x5c, 0xdc, 0x2e, 0xa9, 0xb9, 0xf9, 0xc1, 0xa9, 0x45,
	0x65, 0x99, 0xc9, 0xa9, 0x42, 0x7a, 0x5c, 0xac, 0x60, 0xf7, 0x08, 0x09, 0x22, 0xbb, 0x0d, 0xec,
	0x47, 0x29, 0x21, 0x4c, 0xe7, 0x3a, 0xb1, 0x46, 0x81, 0x82, 0x23, 0x89, 0x0d, 0x1c, 0x34, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0e, 0xce, 0xd1, 0x9a, 0x28, 0x01, 0x00, 0x00,
}
