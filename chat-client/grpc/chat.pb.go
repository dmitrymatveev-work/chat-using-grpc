// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

package chat

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

type IntroRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntroRequest) Reset()         { *m = IntroRequest{} }
func (m *IntroRequest) String() string { return proto.CompactTextString(m) }
func (*IntroRequest) ProtoMessage()    {}
func (*IntroRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_08414bdb60c06f38, []int{0}
}
func (m *IntroRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntroRequest.Unmarshal(m, b)
}
func (m *IntroRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntroRequest.Marshal(b, m, deterministic)
}
func (dst *IntroRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntroRequest.Merge(dst, src)
}
func (m *IntroRequest) XXX_Size() int {
	return xxx_messageInfo_IntroRequest.Size(m)
}
func (m *IntroRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IntroRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IntroRequest proto.InternalMessageInfo

func (m *IntroRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type IntroResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntroResponse) Reset()         { *m = IntroResponse{} }
func (m *IntroResponse) String() string { return proto.CompactTextString(m) }
func (*IntroResponse) ProtoMessage()    {}
func (*IntroResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_08414bdb60c06f38, []int{1}
}
func (m *IntroResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntroResponse.Unmarshal(m, b)
}
func (m *IntroResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntroResponse.Marshal(b, m, deterministic)
}
func (dst *IntroResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntroResponse.Merge(dst, src)
}
func (m *IntroResponse) XXX_Size() int {
	return xxx_messageInfo_IntroResponse.Size(m)
}
func (m *IntroResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IntroResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IntroResponse proto.InternalMessageInfo

func (m *IntroResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*IntroRequest)(nil), "IntroRequest")
	proto.RegisterType((*IntroResponse)(nil), "IntroResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatClient interface {
	Introduce(ctx context.Context, in *IntroRequest, opts ...grpc.CallOption) (*IntroResponse, error)
}

type chatClient struct {
	cc *grpc.ClientConn
}

func NewChatClient(cc *grpc.ClientConn) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) Introduce(ctx context.Context, in *IntroRequest, opts ...grpc.CallOption) (*IntroResponse, error) {
	out := new(IntroResponse)
	err := c.cc.Invoke(ctx, "/Chat/Introduce", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServer is the server API for Chat service.
type ChatServer interface {
	Introduce(context.Context, *IntroRequest) (*IntroResponse, error)
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_Introduce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).Introduce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Chat/Introduce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).Introduce(ctx, req.(*IntroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Introduce",
			Handler:    _Chat_Introduce_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_chat_08414bdb60c06f38) }

var fileDescriptor_chat_08414bdb60c06f38 = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0x48, 0x2c,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xd2, 0xe2, 0xe2, 0xf1, 0xcc, 0x2b, 0x29, 0xca, 0x0f,
	0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe2, 0xe2, 0x28, 0x2d, 0x4e, 0x2d, 0xca, 0x4b,
	0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x95, 0x34, 0xb9, 0x78, 0xa1,
	0x6a, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x24, 0xb8, 0xd8, 0x73, 0x53, 0x8b, 0x8b, 0x13,
	0xd3, 0x61, 0x6a, 0x61, 0x5c, 0x23, 0x13, 0x2e, 0x16, 0xe7, 0x8c, 0xc4, 0x12, 0x21, 0x1d, 0x2e,
	0x4e, 0xb0, 0x96, 0x94, 0xd2, 0xe4, 0x54, 0x21, 0x5e, 0x3d, 0x64, 0xab, 0xa4, 0xf8, 0xf4, 0x50,
	0x4c, 0x53, 0x62, 0x48, 0x62, 0x03, 0xbb, 0xc9, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xea, 0xc1,
	0xa3, 0x44, 0xa1, 0x00, 0x00, 0x00,
}
