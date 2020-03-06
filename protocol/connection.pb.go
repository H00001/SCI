// Code generated by protoc-gen-go. DO NOT EDIT.
// source: connection.proto

package protocol

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

type MagicInbound struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MagicInbound) Reset()         { *m = MagicInbound{} }
func (m *MagicInbound) String() string { return proto.CompactTextString(m) }
func (*MagicInbound) ProtoMessage()    {}
func (*MagicInbound) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{0}
}

func (m *MagicInbound) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MagicInbound.Unmarshal(m, b)
}
func (m *MagicInbound) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MagicInbound.Marshal(b, m, deterministic)
}
func (m *MagicInbound) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MagicInbound.Merge(m, src)
}
func (m *MagicInbound) XXX_Size() int {
	return xxx_messageInfo_MagicInbound.Size(m)
}
func (m *MagicInbound) XXX_DiscardUnknown() {
	xxx_messageInfo_MagicInbound.DiscardUnknown(m)
}

var xxx_messageInfo_MagicInbound proto.InternalMessageInfo

func (m *MagicInbound) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type MagicOutbound struct {
	Magic                string   `protobuf:"bytes,1,opt,name=magic,proto3" json:"magic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MagicOutbound) Reset()         { *m = MagicOutbound{} }
func (m *MagicOutbound) String() string { return proto.CompactTextString(m) }
func (*MagicOutbound) ProtoMessage()    {}
func (*MagicOutbound) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{1}
}

func (m *MagicOutbound) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MagicOutbound.Unmarshal(m, b)
}
func (m *MagicOutbound) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MagicOutbound.Marshal(b, m, deterministic)
}
func (m *MagicOutbound) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MagicOutbound.Merge(m, src)
}
func (m *MagicOutbound) XXX_Size() int {
	return xxx_messageInfo_MagicOutbound.Size(m)
}
func (m *MagicOutbound) XXX_DiscardUnknown() {
	xxx_messageInfo_MagicOutbound.DiscardUnknown(m)
}

var xxx_messageInfo_MagicOutbound proto.InternalMessageInfo

func (m *MagicOutbound) GetMagic() string {
	if m != nil {
		return m.Magic
	}
	return ""
}

type Base64Result struct {
	B65                  string   `protobuf:"bytes,1,opt,name=b65,proto3" json:"b65,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Base64Result) Reset()         { *m = Base64Result{} }
func (m *Base64Result) String() string { return proto.CompactTextString(m) }
func (*Base64Result) ProtoMessage()    {}
func (*Base64Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{2}
}

func (m *Base64Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Base64Result.Unmarshal(m, b)
}
func (m *Base64Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Base64Result.Marshal(b, m, deterministic)
}
func (m *Base64Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Base64Result.Merge(m, src)
}
func (m *Base64Result) XXX_Size() int {
	return xxx_messageInfo_Base64Result.Size(m)
}
func (m *Base64Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Base64Result.DiscardUnknown(m)
}

var xxx_messageInfo_Base64Result proto.InternalMessageInfo

func (m *Base64Result) GetB65() string {
	if m != nil {
		return m.B65
	}
	return ""
}

type VerifyResult struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyResult) Reset()         { *m = VerifyResult{} }
func (m *VerifyResult) String() string { return proto.CompactTextString(m) }
func (*VerifyResult) ProtoMessage()    {}
func (*VerifyResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{3}
}

func (m *VerifyResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyResult.Unmarshal(m, b)
}
func (m *VerifyResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyResult.Marshal(b, m, deterministic)
}
func (m *VerifyResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyResult.Merge(m, src)
}
func (m *VerifyResult) XXX_Size() int {
	return xxx_messageInfo_VerifyResult.Size(m)
}
func (m *VerifyResult) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyResult.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyResult proto.InternalMessageInfo

func (m *VerifyResult) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*MagicInbound)(nil), "protocol.MagicInbound")
	proto.RegisterType((*MagicOutbound)(nil), "protocol.MagicOutbound")
	proto.RegisterType((*Base64Result)(nil), "protocol.Base64Result")
	proto.RegisterType((*VerifyResult)(nil), "protocol.VerifyResult")
}

func init() {
	proto.RegisterFile("connection.proto", fileDescriptor_51baa40a1cc6b48b)
}

var fileDescriptor_51baa40a1cc6b48b = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0xce, 0xcf, 0xcb,
	0x4b, 0x4d, 0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53,
	0xc9, 0xf9, 0x39, 0x4a, 0x4a, 0x5c, 0x3c, 0xbe, 0x89, 0xe9, 0x99, 0xc9, 0x9e, 0x79, 0x49, 0xf9,
	0xa5, 0x79, 0x29, 0x42, 0x42, 0x5c, 0x2c, 0xa5, 0xa5, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0x9c, 0x41, 0x60, 0xb6, 0x92, 0x2a, 0x17, 0x2f, 0x58, 0x8d, 0x7f, 0x69, 0x09, 0x44, 0x91, 0x08,
	0x17, 0x6b, 0x2e, 0x48, 0x00, 0xaa, 0x0a, 0xc2, 0x51, 0x52, 0xe0, 0xe2, 0x71, 0x4a, 0x2c, 0x4e,
	0x35, 0x33, 0x09, 0x4a, 0x2d, 0x2e, 0xcd, 0x29, 0x11, 0x12, 0xe0, 0x62, 0x4e, 0x32, 0x33, 0x85,
	0xaa, 0x01, 0x31, 0x95, 0x54, 0xb8, 0x78, 0xc2, 0x52, 0x8b, 0x32, 0xd3, 0x2a, 0xa1, 0x2a, 0x44,
	0xb8, 0x58, 0x4b, 0xf2, 0xb3, 0x53, 0xf3, 0x60, 0xe6, 0x80, 0x39, 0x46, 0x1d, 0x8c, 0x5c, 0x6c,
	0x10, 0x65, 0x42, 0xf6, 0x5c, 0x3c, 0x8e, 0xc9, 0x85, 0xa5, 0x99, 0x45, 0xa9, 0x60, 0x07, 0x08,
	0x89, 0xe9, 0xc1, 0x1c, 0xae, 0x87, 0xec, 0x6a, 0x29, 0x71, 0x34, 0x71, 0xb8, 0x4b, 0x6d, 0xb9,
	0xb8, 0xa1, 0x06, 0x38, 0x27, 0xe6, 0xa0, 0xe8, 0x47, 0x76, 0xaa, 0x14, 0x92, 0x38, 0xb2, 0x03,
	0x93, 0xd8, 0xc0, 0xc2, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x92, 0x75, 0xa5, 0x42,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// VerifyClient is the client API for Verify service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VerifyClient interface {
	AcquireMagic(ctx context.Context, in *MagicInbound, opts ...grpc.CallOption) (*MagicOutbound, error)
	AcquireCalc(ctx context.Context, in *Base64Result, opts ...grpc.CallOption) (*VerifyResult, error)
}

type verifyClient struct {
	cc grpc.ClientConnInterface
}

func NewVerifyClient(cc grpc.ClientConnInterface) VerifyClient {
	return &verifyClient{cc}
}

func (c *verifyClient) AcquireMagic(ctx context.Context, in *MagicInbound, opts ...grpc.CallOption) (*MagicOutbound, error) {
	out := new(MagicOutbound)
	err := c.cc.Invoke(ctx, "/protocol.Verify/AcquireMagic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verifyClient) AcquireCalc(ctx context.Context, in *Base64Result, opts ...grpc.CallOption) (*VerifyResult, error) {
	out := new(VerifyResult)
	err := c.cc.Invoke(ctx, "/protocol.Verify/AcquireCalc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VerifyServer is the server API for Verify service.
type VerifyServer interface {
	AcquireMagic(context.Context, *MagicInbound) (*MagicOutbound, error)
	AcquireCalc(context.Context, *Base64Result) (*VerifyResult, error)
}

// UnimplementedVerifyServer can be embedded to have forward compatible implementations.
type UnimplementedVerifyServer struct {
}

func (*UnimplementedVerifyServer) AcquireMagic(ctx context.Context, req *MagicInbound) (*MagicOutbound, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcquireMagic not implemented")
}
func (*UnimplementedVerifyServer) AcquireCalc(ctx context.Context, req *Base64Result) (*VerifyResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcquireCalc not implemented")
}

func RegisterVerifyServer(s *grpc.Server, srv VerifyServer) {
	s.RegisterService(&_Verify_serviceDesc, srv)
}

func _Verify_AcquireMagic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MagicInbound)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifyServer).AcquireMagic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Verify/AcquireMagic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifyServer).AcquireMagic(ctx, req.(*MagicInbound))
	}
	return interceptor(ctx, in, info, handler)
}

func _Verify_AcquireCalc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Base64Result)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifyServer).AcquireCalc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Verify/AcquireCalc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifyServer).AcquireCalc(ctx, req.(*Base64Result))
	}
	return interceptor(ctx, in, info, handler)
}

var _Verify_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.Verify",
	HandlerType: (*VerifyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AcquireMagic",
			Handler:    _Verify_AcquireMagic_Handler,
		},
		{
			MethodName: "AcquireCalc",
			Handler:    _Verify_AcquireCalc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "connection.proto",
}
