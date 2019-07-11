// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/balance/balance.proto

package balance

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

type Request struct {
	AccountId            string   `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_be989d5aed6f6a6c, []int{0}
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

func (m *Request) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

type Response struct {
	Completed            bool     `protobuf:"varint,1,opt,name=completed,proto3" json:"completed,omitempty"`
	Amount               int64    `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency             string   `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_be989d5aed6f6a6c, []int{1}
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

func (m *Response) GetCompleted() bool {
	if m != nil {
		return m.Completed
	}
	return false
}

func (m *Response) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Response) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

type Message struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_be989d5aed6f6a6c, []int{2}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "balance.Request")
	proto.RegisterType((*Response)(nil), "balance.Response")
	proto.RegisterType((*Message)(nil), "balance.Message")
}

func init() { proto.RegisterFile("proto/balance/balance.proto", fileDescriptor_be989d5aed6f6a6c) }

var fileDescriptor_be989d5aed6f6a6c = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xb1, 0x4f, 0x84, 0x30,
	0x14, 0xc6, 0x45, 0x22, 0x85, 0x37, 0xe1, 0x1b, 0x0c, 0x01, 0x4d, 0x48, 0x27, 0x26, 0x4c, 0x74,
	0x72, 0xd5, 0xc9, 0xc1, 0xa5, 0xb3, 0x89, 0x29, 0xe5, 0x45, 0xcd, 0x41, 0xcb, 0xd1, 0x32, 0xf0,
	0xdf, 0x5f, 0x80, 0x72, 0x77, 0x53, 0xfb, 0xfd, 0xf2, 0xde, 0xcb, 0xef, 0x83, 0x62, 0x18, 0x8d,
	0x33, 0xcf, 0x8d, 0xec, 0xa4, 0x56, 0xb4, 0xbf, 0xf5, 0x4a, 0x91, 0xf9, 0xc8, 0x2b, 0x60, 0x82,
	0x8e, 0x13, 0x59, 0x87, 0x4f, 0x00, 0x52, 0x29, 0x33, 0x69, 0xf7, 0xf3, 0xdf, 0x66, 0x41, 0x19,
	0x54, 0x89, 0x48, 0x3c, 0xf9, 0x6c, 0xf9, 0x37, 0xc4, 0x82, 0xec, 0x60, 0xb4, 0x25, 0x7c, 0x84,
	0x44, 0x99, 0x7e, 0xe8, 0xc8, 0xd1, 0x36, 0x19, 0x8b, 0x0b, 0xc0, 0x07, 0x88, 0x64, 0xbf, 0x6c,
	0x65, 0xb7, 0x65, 0x50, 0x85, 0xc2, 0x27, 0xcc, 0x21, 0x56, 0xd3, 0x38, 0x92, 0x56, 0x73, 0x16,
	0xae, 0xe7, 0xcf, 0x99, 0x17, 0xc0, 0xbe, 0xc8, 0x5a, 0xf9, 0x4b, 0x98, 0x42, 0x68, 0xe5, 0xec,
	0x05, 0x96, 0xef, 0xcb, 0x1b, 0xb0, 0xf7, 0xcd, 0x17, 0x6b, 0xb8, 0xfb, 0xf8, 0x23, 0x75, 0xc0,
	0xb4, 0xde, 0x1b, 0x79, 0xff, 0xfc, 0xfe, 0x8a, 0x6c, 0x9e, 0xfc, 0xa6, 0x89, 0xd6, 0xbe, 0xaf,
	0xa7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x96, 0x1f, 0xa2, 0x09, 0x0e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BalanceClient is the client API for Balance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BalanceClient interface {
	Check(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type balanceClient struct {
	cc *grpc.ClientConn
}

func NewBalanceClient(cc *grpc.ClientConn) BalanceClient {
	return &balanceClient{cc}
}

func (c *balanceClient) Check(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/balance.Balance/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BalanceServer is the server API for Balance service.
type BalanceServer interface {
	Check(context.Context, *Request) (*Response, error)
}

// UnimplementedBalanceServer can be embedded to have forward compatible implementations.
type UnimplementedBalanceServer struct {
}

func (*UnimplementedBalanceServer) Check(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}

func RegisterBalanceServer(s *grpc.Server, srv BalanceServer) {
	s.RegisterService(&_Balance_serviceDesc, srv)
}

func _Balance_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalanceServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/balance.Balance/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalanceServer).Check(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Balance_serviceDesc = grpc.ServiceDesc{
	ServiceName: "balance.Balance",
	HandlerType: (*BalanceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Balance_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/balance/balance.proto",
}
