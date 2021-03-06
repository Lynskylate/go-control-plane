// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/eds.proto

package envoy_api_v2

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type EdsDummy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EdsDummy) Reset()         { *m = EdsDummy{} }
func (m *EdsDummy) String() string { return proto.CompactTextString(m) }
func (*EdsDummy) ProtoMessage()    {}
func (*EdsDummy) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c8760f38742c17f, []int{0}
}

func (m *EdsDummy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EdsDummy.Unmarshal(m, b)
}
func (m *EdsDummy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EdsDummy.Marshal(b, m, deterministic)
}
func (m *EdsDummy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EdsDummy.Merge(m, src)
}
func (m *EdsDummy) XXX_Size() int {
	return xxx_messageInfo_EdsDummy.Size(m)
}
func (m *EdsDummy) XXX_DiscardUnknown() {
	xxx_messageInfo_EdsDummy.DiscardUnknown(m)
}

var xxx_messageInfo_EdsDummy proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EdsDummy)(nil), "envoy.api.v2.EdsDummy")
}

func init() { proto.RegisterFile("envoy/api/v2/eds.proto", fileDescriptor_3c8760f38742c17f) }

var fileDescriptor_3c8760f38742c17f = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0xbf, 0xf4, 0x13, 0x15, 0xb2, 0x50, 0x91, 0x32, 0x50, 0x08, 0x55, 0x41, 0x85, 0xa1,
	0x62, 0x48, 0x50, 0xbb, 0x75, 0xac, 0x5a, 0xe6, 0x42, 0x17, 0xd6, 0xdb, 0xfa, 0xd2, 0x5a, 0x4a,
	0x6c, 0x63, 0x3b, 0x86, 0x88, 0x8d, 0x89, 0x8d, 0x81, 0x85, 0xf7, 0x82, 0x47, 0xe0, 0x09, 0x78,
	0x00, 0x84, 0x9a, 0x3f, 0xa5, 0x01, 0x95, 0x89, 0x2d, 0xf1, 0xef, 0xdc, 0x73, 0x7d, 0xef, 0x31,
	0xd9, 0x41, 0x6e, 0x45, 0x12, 0x80, 0x64, 0x81, 0xed, 0x04, 0x48, 0xb5, 0x2f, 0x95, 0x30, 0xc2,
	0xdd, 0x4a, 0xcf, 0x7d, 0x90, 0xcc, 0xb7, 0x1d, 0xaf, 0x51, 0x52, 0x51, 0xa6, 0xa7, 0xc2, 0xa2,
	0x4a, 0x32, 0xad, 0xd7, 0x98, 0x09, 0x31, 0x0b, 0x31, 0xc5, 0xc0, 0xb9, 0x30, 0x60, 0x98, 0xe0,
	0xb9, 0x93, 0xd7, 0xcc, 0x69, 0xfa, 0x37, 0x89, 0xaf, 0x02, 0x1a, 0xab, 0x54, 0xb0, 0x8e, 0xdf,
	0x28, 0x90, 0x12, 0xd5, 0xb2, 0x3e, 0xa6, 0x12, 0x56, 0x7d, 0x83, 0x88, 0xcd, 0x14, 0x18, 0xcc,
	0x79, 0xdd, 0x42, 0xc8, 0x28, 0x18, 0x0c, 0x8a, 0x8f, 0x1c, 0xec, 0x97, 0x47, 0xe3, 0x54, 0x0a,
	0xc6, 0x4d, 0x06, 0x5b, 0x84, 0x6c, 0x0e, 0xa9, 0x1e, 0xc4, 0x51, 0x94, 0x74, 0x5e, 0x2b, 0x64,
	0x77, 0x98, 0xe3, 0x41, 0x31, 0xdb, 0x18, 0x95, 0x65, 0x53, 0x74, 0x2f, 0xc9, 0xf6, 0xd8, 0x28,
	0x84, 0xa8, 0x50, 0x68, 0xb7, 0xe9, 0xaf, 0x2e, 0xc7, 0x5f, 0x96, 0x5c, 0xe0, 0x75, 0x8c, 0xda,
	0x78, 0x07, 0x6b, 0xb9, 0x96, 0x82, 0x6b, 0x6c, 0xfd, 0x6b, 0x3b, 0xa7, 0x8e, 0x0b, 0xa4, 0x36,
	0xc0, 0xd0, 0xc0, 0x97, 0xf1, 0xd1, 0xb7, 0xc2, 0x05, 0xfd, 0xe1, 0x7e, 0xfc, 0xbb, 0xa8, 0xd4,
	0xe2, 0x8e, 0xd4, 0xce, 0xd0, 0x4c, 0xe7, 0x7f, 0x78, 0xf7, 0xf6, 0xfd, 0xcb, 0xdb, 0x53, 0x65,
	0xaf, 0x55, 0x2f, 0xbd, 0x84, 0x5e, 0xb1, 0x5e, 0x9d, 0xe2, 0xff, 0x3d, 0xe7, 0xa4, 0x7f, 0xfe,
	0xfe, 0xfc, 0xf1, 0xb8, 0x71, 0xe8, 0x36, 0x33, 0x47, 0x9d, 0x2d, 0xd4, 0x5f, 0x06, 0x61, 0xbb,
	0x10, 0xca, 0x39, 0x10, 0x8f, 0x89, 0xac, 0xa9, 0x54, 0xe2, 0x36, 0x29, 0xf5, 0xef, 0x2f, 0x42,
	0x1a, 0x2d, 0x02, 0x1b, 0x39, 0x0f, 0x8e, 0x33, 0xaa, 0x4e, 0xaa, 0x69, 0x7c, 0xdd, 0xcf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x79, 0x81, 0x45, 0x8c, 0xb8, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EndpointDiscoveryServiceClient is the client API for EndpointDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EndpointDiscoveryServiceClient interface {
	StreamEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_StreamEndpointsClient, error)
	DeltaEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_DeltaEndpointsClient, error)
	FetchEndpoints(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type endpointDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewEndpointDiscoveryServiceClient(cc *grpc.ClientConn) EndpointDiscoveryServiceClient {
	return &endpointDiscoveryServiceClient{cc}
}

func (c *endpointDiscoveryServiceClient) StreamEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_StreamEndpointsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EndpointDiscoveryService_serviceDesc.Streams[0], "/envoy.api.v2.EndpointDiscoveryService/StreamEndpoints", opts...)
	if err != nil {
		return nil, err
	}
	x := &endpointDiscoveryServiceStreamEndpointsClient{stream}
	return x, nil
}

type EndpointDiscoveryService_StreamEndpointsClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type endpointDiscoveryServiceStreamEndpointsClient struct {
	grpc.ClientStream
}

func (x *endpointDiscoveryServiceStreamEndpointsClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceStreamEndpointsClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *endpointDiscoveryServiceClient) DeltaEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_DeltaEndpointsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EndpointDiscoveryService_serviceDesc.Streams[1], "/envoy.api.v2.EndpointDiscoveryService/DeltaEndpoints", opts...)
	if err != nil {
		return nil, err
	}
	x := &endpointDiscoveryServiceDeltaEndpointsClient{stream}
	return x, nil
}

type EndpointDiscoveryService_DeltaEndpointsClient interface {
	Send(*DeltaDiscoveryRequest) error
	Recv() (*DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type endpointDiscoveryServiceDeltaEndpointsClient struct {
	grpc.ClientStream
}

func (x *endpointDiscoveryServiceDeltaEndpointsClient) Send(m *DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceDeltaEndpointsClient) Recv() (*DeltaDiscoveryResponse, error) {
	m := new(DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *endpointDiscoveryServiceClient) FetchEndpoints(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.api.v2.EndpointDiscoveryService/FetchEndpoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndpointDiscoveryServiceServer is the server API for EndpointDiscoveryService service.
type EndpointDiscoveryServiceServer interface {
	StreamEndpoints(EndpointDiscoveryService_StreamEndpointsServer) error
	DeltaEndpoints(EndpointDiscoveryService_DeltaEndpointsServer) error
	FetchEndpoints(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
}

// UnimplementedEndpointDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEndpointDiscoveryServiceServer struct {
}

func (*UnimplementedEndpointDiscoveryServiceServer) StreamEndpoints(srv EndpointDiscoveryService_StreamEndpointsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamEndpoints not implemented")
}
func (*UnimplementedEndpointDiscoveryServiceServer) DeltaEndpoints(srv EndpointDiscoveryService_DeltaEndpointsServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaEndpoints not implemented")
}
func (*UnimplementedEndpointDiscoveryServiceServer) FetchEndpoints(ctx context.Context, req *DiscoveryRequest) (*DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchEndpoints not implemented")
}

func RegisterEndpointDiscoveryServiceServer(s *grpc.Server, srv EndpointDiscoveryServiceServer) {
	s.RegisterService(&_EndpointDiscoveryService_serviceDesc, srv)
}

func _EndpointDiscoveryService_StreamEndpoints_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EndpointDiscoveryServiceServer).StreamEndpoints(&endpointDiscoveryServiceStreamEndpointsServer{stream})
}

type EndpointDiscoveryService_StreamEndpointsServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type endpointDiscoveryServiceStreamEndpointsServer struct {
	grpc.ServerStream
}

func (x *endpointDiscoveryServiceStreamEndpointsServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceStreamEndpointsServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EndpointDiscoveryService_DeltaEndpoints_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EndpointDiscoveryServiceServer).DeltaEndpoints(&endpointDiscoveryServiceDeltaEndpointsServer{stream})
}

type EndpointDiscoveryService_DeltaEndpointsServer interface {
	Send(*DeltaDiscoveryResponse) error
	Recv() (*DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type endpointDiscoveryServiceDeltaEndpointsServer struct {
	grpc.ServerStream
}

func (x *endpointDiscoveryServiceDeltaEndpointsServer) Send(m *DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceDeltaEndpointsServer) Recv() (*DeltaDiscoveryRequest, error) {
	m := new(DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EndpointDiscoveryService_FetchEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointDiscoveryServiceServer).FetchEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v2.EndpointDiscoveryService/FetchEndpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointDiscoveryServiceServer).FetchEndpoints(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EndpointDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.EndpointDiscoveryService",
	HandlerType: (*EndpointDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchEndpoints",
			Handler:    _EndpointDiscoveryService_FetchEndpoints_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEndpoints",
			Handler:       _EndpointDiscoveryService_StreamEndpoints_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaEndpoints",
			Handler:       _EndpointDiscoveryService_DeltaEndpoints_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v2/eds.proto",
}
