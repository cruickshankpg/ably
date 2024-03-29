// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/stateful/stateful.proto

package proto

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

type GenerateSequenceRequest struct {
	ConnectionID         string   `protobuf:"bytes,1,opt,name=connectionID,proto3" json:"connectionID,omitempty"`
	SequenceLength       uint32   `protobuf:"varint,2,opt,name=sequenceLength,proto3" json:"sequenceLength,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateSequenceRequest) Reset()         { *m = GenerateSequenceRequest{} }
func (m *GenerateSequenceRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateSequenceRequest) ProtoMessage()    {}
func (*GenerateSequenceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_53b71fbd071323fc, []int{0}
}

func (m *GenerateSequenceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateSequenceRequest.Unmarshal(m, b)
}
func (m *GenerateSequenceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateSequenceRequest.Marshal(b, m, deterministic)
}
func (m *GenerateSequenceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateSequenceRequest.Merge(m, src)
}
func (m *GenerateSequenceRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateSequenceRequest.Size(m)
}
func (m *GenerateSequenceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateSequenceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateSequenceRequest proto.InternalMessageInfo

func (m *GenerateSequenceRequest) GetConnectionID() string {
	if m != nil {
		return m.ConnectionID
	}
	return ""
}

func (m *GenerateSequenceRequest) GetSequenceLength() uint32 {
	if m != nil {
		return m.SequenceLength
	}
	return 0
}

type ReconnectSequenceRequest struct {
	ConnectionID         string   `protobuf:"bytes,1,opt,name=connectionID,proto3" json:"connectionID,omitempty"`
	LastReceivedIndex    uint32   `protobuf:"varint,2,opt,name=lastReceivedIndex,proto3" json:"lastReceivedIndex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReconnectSequenceRequest) Reset()         { *m = ReconnectSequenceRequest{} }
func (m *ReconnectSequenceRequest) String() string { return proto.CompactTextString(m) }
func (*ReconnectSequenceRequest) ProtoMessage()    {}
func (*ReconnectSequenceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_53b71fbd071323fc, []int{1}
}

func (m *ReconnectSequenceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReconnectSequenceRequest.Unmarshal(m, b)
}
func (m *ReconnectSequenceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReconnectSequenceRequest.Marshal(b, m, deterministic)
}
func (m *ReconnectSequenceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReconnectSequenceRequest.Merge(m, src)
}
func (m *ReconnectSequenceRequest) XXX_Size() int {
	return xxx_messageInfo_ReconnectSequenceRequest.Size(m)
}
func (m *ReconnectSequenceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReconnectSequenceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReconnectSequenceRequest proto.InternalMessageInfo

func (m *ReconnectSequenceRequest) GetConnectionID() string {
	if m != nil {
		return m.ConnectionID
	}
	return ""
}

func (m *ReconnectSequenceRequest) GetLastReceivedIndex() uint32 {
	if m != nil {
		return m.LastReceivedIndex
	}
	return 0
}

type Generated struct {
	Number               uint32   `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	FinalItem            bool     `protobuf:"varint,2,opt,name=finalItem,proto3" json:"finalItem,omitempty"`
	Checksum             []byte   `protobuf:"bytes,3,opt,name=checksum,proto3" json:"checksum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Generated) Reset()         { *m = Generated{} }
func (m *Generated) String() string { return proto.CompactTextString(m) }
func (*Generated) ProtoMessage()    {}
func (*Generated) Descriptor() ([]byte, []int) {
	return fileDescriptor_53b71fbd071323fc, []int{2}
}

func (m *Generated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Generated.Unmarshal(m, b)
}
func (m *Generated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Generated.Marshal(b, m, deterministic)
}
func (m *Generated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Generated.Merge(m, src)
}
func (m *Generated) XXX_Size() int {
	return xxx_messageInfo_Generated.Size(m)
}
func (m *Generated) XXX_DiscardUnknown() {
	xxx_messageInfo_Generated.DiscardUnknown(m)
}

var xxx_messageInfo_Generated proto.InternalMessageInfo

func (m *Generated) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Generated) GetFinalItem() bool {
	if m != nil {
		return m.FinalItem
	}
	return false
}

func (m *Generated) GetChecksum() []byte {
	if m != nil {
		return m.Checksum
	}
	return nil
}

func init() {
	proto.RegisterType((*GenerateSequenceRequest)(nil), "ablyStatefulServer.GenerateSequenceRequest")
	proto.RegisterType((*ReconnectSequenceRequest)(nil), "ablyStatefulServer.ReconnectSequenceRequest")
	proto.RegisterType((*Generated)(nil), "ablyStatefulServer.Generated")
}

func init() { proto.RegisterFile("protos/stateful/stateful.proto", fileDescriptor_53b71fbd071323fc) }

var fileDescriptor_53b71fbd071323fc = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xd1, 0x4a, 0xfb, 0x30,
	0x14, 0xc6, 0xc9, 0xff, 0x8f, 0x73, 0x3d, 0x6c, 0xe2, 0xce, 0x85, 0x2b, 0x43, 0xa5, 0xf4, 0x42,
	0x0a, 0x8e, 0x2a, 0xfa, 0x06, 0x22, 0x48, 0x41, 0xbc, 0x48, 0xef, 0x04, 0xc1, 0x36, 0x3d, 0x73,
	0xc5, 0x36, 0xd5, 0x24, 0x1d, 0xfa, 0xb4, 0xbe, 0x8a, 0x18, 0xb2, 0x89, 0xab, 0x53, 0xf0, 0xea,
	0xe4, 0x7c, 0x27, 0xe7, 0xfb, 0xc8, 0x8f, 0xc0, 0xe1, 0x93, 0x6a, 0x4c, 0xa3, 0x4f, 0xb4, 0xc9,
	0x0c, 0xcd, 0xda, 0x6a, 0x75, 0x88, 0xed, 0x00, 0x31, 0xcb, 0xab, 0xd7, 0xd4, 0x69, 0x29, 0xa9,
	0x05, 0xa9, 0x90, 0x60, 0x7c, 0x45, 0x92, 0x54, 0x66, 0x28, 0xa5, 0xe7, 0x96, 0xa4, 0x20, 0xfe,
	0x51, 0xb5, 0xc1, 0x10, 0x06, 0xa2, 0x91, 0x92, 0x84, 0x29, 0x1b, 0x99, 0x5c, 0xfa, 0x2c, 0x60,
	0x91, 0xc7, 0xbf, 0x68, 0x78, 0x04, 0x3b, 0xda, 0xad, 0x5d, 0x93, 0x7c, 0x30, 0x73, 0xff, 0x5f,
	0xc0, 0xa2, 0x21, 0x5f, 0x53, 0xc3, 0x0a, 0x7c, 0x4e, 0x6e, 0xf3, 0x2f, 0x39, 0x53, 0x18, 0x55,
	0x99, 0x36, 0x9c, 0x04, 0x95, 0x0b, 0x2a, 0x12, 0x59, 0xd0, 0x8b, 0x8b, 0xea, 0x0e, 0xc2, 0x3b,
	0xf0, 0x96, 0x8f, 0x2a, 0x70, 0x0f, 0x7a, 0xb2, 0xad, 0x73, 0x52, 0xd6, 0x78, 0xc8, 0x5d, 0x87,
	0xfb, 0xe0, 0xcd, 0x4a, 0x99, 0x55, 0x89, 0xa1, 0xda, 0x5a, 0xf5, 0xf9, 0xa7, 0x80, 0x13, 0xe8,
	0x8b, 0x39, 0x89, 0x47, 0xdd, 0xd6, 0xfe, 0xff, 0x80, 0x45, 0x03, 0xbe, 0xea, 0xcf, 0xde, 0x18,
	0x8c, 0x97, 0x18, 0x6f, 0xac, 0x99, 0x4b, 0x6b, 0x14, 0xde, 0xc3, 0xee, 0x3a, 0x4f, 0x3c, 0x8e,
	0xbb, 0xe0, 0xe3, 0x0d, 0xd4, 0x27, 0x07, 0x3f, 0x5d, 0x2e, 0x4e, 0x19, 0xe6, 0x30, 0xea, 0xa0,
	0xc4, 0xe9, 0x77, 0x5b, 0x9b, 0x88, 0xff, 0x9a, 0x71, 0xb1, 0x7d, 0xbb, 0x65, 0xbf, 0x4c, 0xde,
	0xb3, 0xe5, 0xfc, 0x3d, 0x00, 0x00, 0xff, 0xff, 0x23, 0xf1, 0xc1, 0x04, 0x5b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StatefulNumberGeneratorClient is the client API for StatefulNumberGenerator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StatefulNumberGeneratorClient interface {
	GenerateSequence(ctx context.Context, in *GenerateSequenceRequest, opts ...grpc.CallOption) (StatefulNumberGenerator_GenerateSequenceClient, error)
	ReconnectSequence(ctx context.Context, in *ReconnectSequenceRequest, opts ...grpc.CallOption) (StatefulNumberGenerator_ReconnectSequenceClient, error)
}

type statefulNumberGeneratorClient struct {
	cc *grpc.ClientConn
}

func NewStatefulNumberGeneratorClient(cc *grpc.ClientConn) StatefulNumberGeneratorClient {
	return &statefulNumberGeneratorClient{cc}
}

func (c *statefulNumberGeneratorClient) GenerateSequence(ctx context.Context, in *GenerateSequenceRequest, opts ...grpc.CallOption) (StatefulNumberGenerator_GenerateSequenceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StatefulNumberGenerator_serviceDesc.Streams[0], "/ablyStatefulServer.StatefulNumberGenerator/GenerateSequence", opts...)
	if err != nil {
		return nil, err
	}
	x := &statefulNumberGeneratorGenerateSequenceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StatefulNumberGenerator_GenerateSequenceClient interface {
	Recv() (*Generated, error)
	grpc.ClientStream
}

type statefulNumberGeneratorGenerateSequenceClient struct {
	grpc.ClientStream
}

func (x *statefulNumberGeneratorGenerateSequenceClient) Recv() (*Generated, error) {
	m := new(Generated)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *statefulNumberGeneratorClient) ReconnectSequence(ctx context.Context, in *ReconnectSequenceRequest, opts ...grpc.CallOption) (StatefulNumberGenerator_ReconnectSequenceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StatefulNumberGenerator_serviceDesc.Streams[1], "/ablyStatefulServer.StatefulNumberGenerator/ReconnectSequence", opts...)
	if err != nil {
		return nil, err
	}
	x := &statefulNumberGeneratorReconnectSequenceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StatefulNumberGenerator_ReconnectSequenceClient interface {
	Recv() (*Generated, error)
	grpc.ClientStream
}

type statefulNumberGeneratorReconnectSequenceClient struct {
	grpc.ClientStream
}

func (x *statefulNumberGeneratorReconnectSequenceClient) Recv() (*Generated, error) {
	m := new(Generated)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StatefulNumberGeneratorServer is the server API for StatefulNumberGenerator service.
type StatefulNumberGeneratorServer interface {
	GenerateSequence(*GenerateSequenceRequest, StatefulNumberGenerator_GenerateSequenceServer) error
	ReconnectSequence(*ReconnectSequenceRequest, StatefulNumberGenerator_ReconnectSequenceServer) error
}

// UnimplementedStatefulNumberGeneratorServer can be embedded to have forward compatible implementations.
type UnimplementedStatefulNumberGeneratorServer struct {
}

func (*UnimplementedStatefulNumberGeneratorServer) GenerateSequence(req *GenerateSequenceRequest, srv StatefulNumberGenerator_GenerateSequenceServer) error {
	return status.Errorf(codes.Unimplemented, "method GenerateSequence not implemented")
}
func (*UnimplementedStatefulNumberGeneratorServer) ReconnectSequence(req *ReconnectSequenceRequest, srv StatefulNumberGenerator_ReconnectSequenceServer) error {
	return status.Errorf(codes.Unimplemented, "method ReconnectSequence not implemented")
}

func RegisterStatefulNumberGeneratorServer(s *grpc.Server, srv StatefulNumberGeneratorServer) {
	s.RegisterService(&_StatefulNumberGenerator_serviceDesc, srv)
}

func _StatefulNumberGenerator_GenerateSequence_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GenerateSequenceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StatefulNumberGeneratorServer).GenerateSequence(m, &statefulNumberGeneratorGenerateSequenceServer{stream})
}

type StatefulNumberGenerator_GenerateSequenceServer interface {
	Send(*Generated) error
	grpc.ServerStream
}

type statefulNumberGeneratorGenerateSequenceServer struct {
	grpc.ServerStream
}

func (x *statefulNumberGeneratorGenerateSequenceServer) Send(m *Generated) error {
	return x.ServerStream.SendMsg(m)
}

func _StatefulNumberGenerator_ReconnectSequence_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReconnectSequenceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StatefulNumberGeneratorServer).ReconnectSequence(m, &statefulNumberGeneratorReconnectSequenceServer{stream})
}

type StatefulNumberGenerator_ReconnectSequenceServer interface {
	Send(*Generated) error
	grpc.ServerStream
}

type statefulNumberGeneratorReconnectSequenceServer struct {
	grpc.ServerStream
}

func (x *statefulNumberGeneratorReconnectSequenceServer) Send(m *Generated) error {
	return x.ServerStream.SendMsg(m)
}

var _StatefulNumberGenerator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ablyStatefulServer.StatefulNumberGenerator",
	HandlerType: (*StatefulNumberGeneratorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GenerateSequence",
			Handler:       _StatefulNumberGenerator_GenerateSequence_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ReconnectSequence",
			Handler:       _StatefulNumberGenerator_ReconnectSequence_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/stateful/stateful.proto",
}
