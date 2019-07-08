// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messaging.proto

package messaging

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type KeywordsRequest struct {
	Txn                  []byte   `protobuf:"bytes,1,opt,name=txn,proto3" json:"txn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeywordsRequest) Reset()         { *m = KeywordsRequest{} }
func (m *KeywordsRequest) String() string { return proto.CompactTextString(m) }
func (*KeywordsRequest) ProtoMessage()    {}
func (*KeywordsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_42a1718997f046ec, []int{0}
}

func (m *KeywordsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordsRequest.Unmarshal(m, b)
}
func (m *KeywordsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordsRequest.Marshal(b, m, deterministic)
}
func (m *KeywordsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordsRequest.Merge(m, src)
}
func (m *KeywordsRequest) XXX_Size() int {
	return xxx_messageInfo_KeywordsRequest.Size(m)
}
func (m *KeywordsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordsRequest proto.InternalMessageInfo

func (m *KeywordsRequest) GetTxn() []byte {
	if m != nil {
		return m.Txn
	}
	return nil
}

type KeywordsResponse struct {
	Txn                  []byte   `protobuf:"bytes,1,opt,name=txn,proto3" json:"txn,omitempty"`
	Keyword              []string `protobuf:"bytes,2,rep,name=keyword,proto3" json:"keyword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeywordsResponse) Reset()         { *m = KeywordsResponse{} }
func (m *KeywordsResponse) String() string { return proto.CompactTextString(m) }
func (*KeywordsResponse) ProtoMessage()    {}
func (*KeywordsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_42a1718997f046ec, []int{1}
}

func (m *KeywordsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordsResponse.Unmarshal(m, b)
}
func (m *KeywordsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordsResponse.Marshal(b, m, deterministic)
}
func (m *KeywordsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordsResponse.Merge(m, src)
}
func (m *KeywordsResponse) XXX_Size() int {
	return xxx_messageInfo_KeywordsResponse.Size(m)
}
func (m *KeywordsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordsResponse proto.InternalMessageInfo

func (m *KeywordsResponse) GetTxn() []byte {
	if m != nil {
		return m.Txn
	}
	return nil
}

func (m *KeywordsResponse) GetKeyword() []string {
	if m != nil {
		return m.Keyword
	}
	return nil
}

type Request struct {
	Txn                  []byte   `protobuf:"bytes,1,opt,name=txn,proto3" json:"txn,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_42a1718997f046ec, []int{2}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetTxn() []byte {
	if m != nil {
		return m.Txn
	}
	return nil
}

func (m *Request) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Response struct {
	Txn                  []byte   `protobuf:"bytes,1,opt,name=txn,proto3" json:"txn,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_42a1718997f046ec, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetTxn() []byte {
	if m != nil {
		return m.Txn
	}
	return nil
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*KeywordsRequest)(nil), "KeywordsRequest")
	proto.RegisterType((*KeywordsResponse)(nil), "KeywordsResponse")
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Response)(nil), "Response")
}

func init() { proto.RegisterFile("messaging.proto", fileDescriptor_42a1718997f046ec) }

var fileDescriptor_42a1718997f046ec = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0xcf, 0xcc, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe6, 0xe2, 0xf7,
	0x4e, 0xad, 0x2c, 0xcf, 0x2f, 0x4a, 0x29, 0x0e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12,
	0xe0, 0x62, 0x2e, 0xa9, 0xc8, 0x93, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x02, 0x31, 0x95, 0xec,
	0xb8, 0x04, 0x10, 0x8a, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x31, 0x55, 0x09, 0x49, 0x70, 0xb1,
	0x67, 0x43, 0x54, 0x49, 0x30, 0x29, 0x30, 0x6b, 0x70, 0x06, 0xc1, 0xb8, 0x4a, 0xa6, 0x5c, 0xec,
	0x38, 0x0d, 0x07, 0x69, 0x83, 0x38, 0x2a, 0x55, 0x82, 0x49, 0x81, 0x11, 0xa4, 0x0d, 0xca, 0x55,
	0x32, 0xe3, 0xe2, 0xc0, 0x6f, 0x1d, 0x76, 0x7d, 0x46, 0x0e, 0x5c, 0x1c, 0x30, 0xe7, 0x0a, 0x99,
	0x70, 0x71, 0xbb, 0xa7, 0x96, 0xc0, 0xb9, 0x02, 0x7a, 0x68, 0xbe, 0x95, 0x12, 0xd4, 0x43, 0xf7,
	0x9a, 0x12, 0x83, 0x91, 0x21, 0x17, 0xbb, 0x2f, 0xc4, 0x30, 0x21, 0x35, 0x2e, 0x5e, 0x8f, 0xc4,
	0xbc, 0x94, 0x9c, 0x54, 0x98, 0x00, 0x87, 0x1e, 0x4c, 0x2b, 0xa7, 0x1e, 0x42, 0x4b, 0x12, 0x1b,
	0x38, 0x3c, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x93, 0xce, 0xd1, 0x82, 0x62, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KeywordsClient is the client API for Keywords service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeywordsClient interface {
	GetKeywords(ctx context.Context, in *KeywordsRequest, opts ...grpc.CallOption) (*KeywordsResponse, error)
}

type keywordsClient struct {
	cc *grpc.ClientConn
}

func NewKeywordsClient(cc *grpc.ClientConn) KeywordsClient {
	return &keywordsClient{cc}
}

func (c *keywordsClient) GetKeywords(ctx context.Context, in *KeywordsRequest, opts ...grpc.CallOption) (*KeywordsResponse, error) {
	out := new(KeywordsResponse)
	err := c.cc.Invoke(ctx, "/Keywords/GetKeywords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordsServer is the server API for Keywords service.
type KeywordsServer interface {
	GetKeywords(context.Context, *KeywordsRequest) (*KeywordsResponse, error)
}

// UnimplementedKeywordsServer can be embedded to have forward compatible implementations.
type UnimplementedKeywordsServer struct {
}

func (*UnimplementedKeywordsServer) GetKeywords(ctx context.Context, req *KeywordsRequest) (*KeywordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKeywords not implemented")
}

func RegisterKeywordsServer(s *grpc.Server, srv KeywordsServer) {
	s.RegisterService(&_Keywords_serviceDesc, srv)
}

func _Keywords_GetKeywords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeywordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordsServer).GetKeywords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Keywords/GetKeywords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordsServer).GetKeywords(ctx, req.(*KeywordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Keywords_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Keywords",
	HandlerType: (*KeywordsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKeywords",
			Handler:    _Keywords_GetKeywords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "messaging.proto",
}

// MessageClient is the client API for Message service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageClient interface {
	HandleMessage(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type messageClient struct {
	cc *grpc.ClientConn
}

func NewMessageClient(cc *grpc.ClientConn) MessageClient {
	return &messageClient{cc}
}

func (c *messageClient) HandleMessage(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Message/HandleMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServer is the server API for Message service.
type MessageServer interface {
	HandleMessage(context.Context, *Request) (*Response, error)
}

// UnimplementedMessageServer can be embedded to have forward compatible implementations.
type UnimplementedMessageServer struct {
}

func (*UnimplementedMessageServer) HandleMessage(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleMessage not implemented")
}

func RegisterMessageServer(s *grpc.Server, srv MessageServer) {
	s.RegisterService(&_Message_serviceDesc, srv)
}

func _Message_HandleMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).HandleMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Message/HandleMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).HandleMessage(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Message_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Message",
	HandlerType: (*MessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleMessage",
			Handler:    _Message_HandleMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "messaging.proto",
}