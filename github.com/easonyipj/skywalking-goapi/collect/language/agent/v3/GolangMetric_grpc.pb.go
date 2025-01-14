// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v3

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	v3 "skywalking.apache.org/repo/goapi/collect/common/v3"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GolangMetricReportServiceClient is the client API for GolangMetricReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GolangMetricReportServiceClient interface {
	Collect(ctx context.Context, in *GolangMetricCollection, opts ...grpc.CallOption) (*v3.Commands, error)
}

type golangMetricReportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGolangMetricReportServiceClient(cc grpc.ClientConnInterface) GolangMetricReportServiceClient {
	return &golangMetricReportServiceClient{cc}
}

func (c *golangMetricReportServiceClient) Collect(ctx context.Context, in *GolangMetricCollection, opts ...grpc.CallOption) (*v3.Commands, error) {
	out := new(v3.Commands)
	err := c.cc.Invoke(ctx, "/skywalking.v3.GolangMetricReportService/collect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GolangMetricReportServiceServer is the server API for GolangMetricReportService service.
// All implementations must embed UnimplementedGolangMetricReportServiceServer
// for forward compatibility
type GolangMetricReportServiceServer interface {
	Collect(context.Context, *GolangMetricCollection) (*v3.Commands, error)
	mustEmbedUnimplementedGolangMetricReportServiceServer()
}

// UnimplementedGolangMetricReportServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGolangMetricReportServiceServer struct {
}

func (UnimplementedGolangMetricReportServiceServer) Collect(context.Context, *GolangMetricCollection) (*v3.Commands, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Collect not implemented")
}
func (UnimplementedGolangMetricReportServiceServer) mustEmbedUnimplementedGolangMetricReportServiceServer() {
}

// UnsafeGolangMetricReportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GolangMetricReportServiceServer will
// result in compilation errors.
type UnsafeGolangMetricReportServiceServer interface {
	mustEmbedUnimplementedGolangMetricReportServiceServer()
}

func RegisterGolangMetricReportServiceServer(s grpc.ServiceRegistrar, srv GolangMetricReportServiceServer) {
	s.RegisterService(&GolangMetricReportService_ServiceDesc, srv)
}

func _GolangMetricReportService_Collect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GolangMetricCollection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GolangMetricReportServiceServer).Collect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skywalking.v3.GolangMetricReportService/collect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GolangMetricReportServiceServer).Collect(ctx, req.(*GolangMetricCollection))
	}
	return interceptor(ctx, in, info, handler)
}

// GolangMetricReportService_ServiceDesc is the grpc.ServiceDesc for GolangMetricReportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GolangMetricReportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "skywalking.v3.GolangMetricReportService",
	HandlerType: (*GolangMetricReportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "collect",
			Handler:    _GolangMetricReportService_Collect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "language-agent/GolangMetric.proto",
}
