// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package beacon

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// BeaconClient is the client API for Beacon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BeaconClient interface {
	Handle(ctx context.Context, in *BeaconIncoming, opts ...grpc.CallOption) (*BeaconOutgoing, error)
}

type beaconClient struct {
	cc grpc.ClientConnInterface
}

func NewBeaconClient(cc grpc.ClientConnInterface) BeaconClient {
	return &beaconClient{cc}
}

func (c *beaconClient) Handle(ctx context.Context, in *BeaconIncoming, opts ...grpc.CallOption) (*BeaconOutgoing, error) {
	out := new(BeaconOutgoing)
	err := c.cc.Invoke(ctx, "/beacon.Beacon/Handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BeaconServer is the server API for Beacon service.
// All implementations must embed UnimplementedBeaconServer
// for forward compatibility
type BeaconServer interface {
	Handle(context.Context, *BeaconIncoming) (*BeaconOutgoing, error)
	mustEmbedUnimplementedBeaconServer()
}

// UnimplementedBeaconServer must be embedded to have forward compatible implementations.
type UnimplementedBeaconServer struct {
}

func (UnimplementedBeaconServer) Handle(context.Context, *BeaconIncoming) (*BeaconOutgoing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}
func (UnimplementedBeaconServer) mustEmbedUnimplementedBeaconServer() {}

// UnsafeBeaconServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BeaconServer will
// result in compilation errors.
type UnsafeBeaconServer interface {
	mustEmbedUnimplementedBeaconServer()
}

func RegisterBeaconServer(s grpc.ServiceRegistrar, srv BeaconServer) {
	s.RegisterService(&_Beacon_serviceDesc, srv)
}

func _Beacon_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BeaconIncoming)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/beacon.Beacon/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconServer).Handle(ctx, req.(*BeaconIncoming))
	}
	return interceptor(ctx, in, info, handler)
}

var _Beacon_serviceDesc = grpc.ServiceDesc{
	ServiceName: "beacon.Beacon",
	HandlerType: (*BeaconServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Beacon_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "beacon.proto",
}