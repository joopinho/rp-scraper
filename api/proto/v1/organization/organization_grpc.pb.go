// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: organization.proto

package organization

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	OrganizationService_QueryProfile_FullMethodName = "/organization.OrganizationService/QueryProfile"
)

// OrganizationServiceClient is the client API for OrganizationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrganizationServiceClient interface {
	QueryProfile(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*ProfileResponse, error)
}

type organizationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrganizationServiceClient(cc grpc.ClientConnInterface) OrganizationServiceClient {
	return &organizationServiceClient{cc}
}

func (c *organizationServiceClient) QueryProfile(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*ProfileResponse, error) {
	out := new(ProfileResponse)
	err := c.cc.Invoke(ctx, OrganizationService_QueryProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrganizationServiceServer is the server API for OrganizationService service.
// All implementations must embed UnimplementedOrganizationServiceServer
// for forward compatibility
type OrganizationServiceServer interface {
	QueryProfile(context.Context, *QueryRequest) (*ProfileResponse, error)
	mustEmbedUnimplementedOrganizationServiceServer()
}

// UnimplementedOrganizationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrganizationServiceServer struct {
}

func (UnimplementedOrganizationServiceServer) QueryProfile(context.Context, *QueryRequest) (*ProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryProfile not implemented")
}
func (UnimplementedOrganizationServiceServer) mustEmbedUnimplementedOrganizationServiceServer() {}

// UnsafeOrganizationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrganizationServiceServer will
// result in compilation errors.
type UnsafeOrganizationServiceServer interface {
	mustEmbedUnimplementedOrganizationServiceServer()
}

func RegisterOrganizationServiceServer(s grpc.ServiceRegistrar, srv OrganizationServiceServer) {
	s.RegisterService(&OrganizationService_ServiceDesc, srv)
}

func _OrganizationService_QueryProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServiceServer).QueryProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrganizationService_QueryProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServiceServer).QueryProfile(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrganizationService_ServiceDesc is the grpc.ServiceDesc for OrganizationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrganizationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "organization.OrganizationService",
	HandlerType: (*OrganizationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryProfile",
			Handler:    _OrganizationService_QueryProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "organization.proto",
}