// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: msg.proto

package msg

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

// CheckClient is the client API for Check service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CheckClient interface {
	SearchID(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	SearchIP4(ctx context.Context, in *IP4Request, opts ...grpc.CallOption) (*SearchResponse, error)
	SearchIP6(ctx context.Context, in *IP6Request, opts ...grpc.CallOption) (*SearchResponse, error)
	SearchURL(ctx context.Context, in *URLRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	SearchDomain(ctx context.Context, in *DomainRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	SearchDecision(ctx context.Context, in *DecisionRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	SearchTextDecision(ctx context.Context, in *TextDecisionRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	SearchSubnet4(ctx context.Context, in *Subnet4Request, opts ...grpc.CallOption) (*SearchResponse, error)
	SearchSubnet6(ctx context.Context, in *Subnet6Request, opts ...grpc.CallOption) (*SearchResponse, error)
	Stat(ctx context.Context, in *StatRequest, opts ...grpc.CallOption) (*StatResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongResponse, error)
}

type checkClient struct {
	cc grpc.ClientConnInterface
}

func NewCheckClient(cc grpc.ClientConnInterface) CheckClient {
	return &checkClient{cc}
}

func (c *checkClient) SearchID(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/msg.Check/SearchID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkClient) SearchIP4(ctx context.Context, in *IP4Request, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/msg.Check/SearchIP4", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkClient) SearchIP6(ctx context.Context, in *IP6Request, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/msg.Check/SearchIP6", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkClient) SearchURL(ctx context.Context, in *URLRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/msg.Check/SearchURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkClient) SearchDomain(ctx context.Context, in *DomainRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/msg.Check/SearchDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkClient) SearchDecision(ctx context.Context, in *DecisionRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/msg.Check/SearchDecision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkClient) SearchTextDecision(ctx context.Context, in *TextDecisionRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/msg.Check/SearchTextDecision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkClient) SearchSubnet4(ctx context.Context, in *Subnet4Request, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/msg.Check/SearchSubnet4", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkClient) SearchSubnet6(ctx context.Context, in *Subnet6Request, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/msg.Check/SearchSubnet6", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkClient) Stat(ctx context.Context, in *StatRequest, opts ...grpc.CallOption) (*StatResponse, error) {
	out := new(StatResponse)
	err := c.cc.Invoke(ctx, "/msg.Check/Stat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongResponse, error) {
	out := new(PongResponse)
	err := c.cc.Invoke(ctx, "/msg.Check/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckServer is the server API for Check service.
// All implementations must embed UnimplementedCheckServer
// for forward compatibility
type CheckServer interface {
	SearchID(context.Context, *IDRequest) (*SearchResponse, error)
	SearchIP4(context.Context, *IP4Request) (*SearchResponse, error)
	SearchIP6(context.Context, *IP6Request) (*SearchResponse, error)
	SearchURL(context.Context, *URLRequest) (*SearchResponse, error)
	SearchDomain(context.Context, *DomainRequest) (*SearchResponse, error)
	SearchDecision(context.Context, *DecisionRequest) (*SearchResponse, error)
	SearchTextDecision(context.Context, *TextDecisionRequest) (*SearchResponse, error)
	SearchSubnet4(context.Context, *Subnet4Request) (*SearchResponse, error)
	SearchSubnet6(context.Context, *Subnet6Request) (*SearchResponse, error)
	Stat(context.Context, *StatRequest) (*StatResponse, error)
	Ping(context.Context, *PingRequest) (*PongResponse, error)
	mustEmbedUnimplementedCheckServer()
}

// UnimplementedCheckServer must be embedded to have forward compatible implementations.
type UnimplementedCheckServer struct {
}

func (UnimplementedCheckServer) SearchID(context.Context, *IDRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchID not implemented")
}
func (UnimplementedCheckServer) SearchIP4(context.Context, *IP4Request) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchIP4 not implemented")
}
func (UnimplementedCheckServer) SearchIP6(context.Context, *IP6Request) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchIP6 not implemented")
}
func (UnimplementedCheckServer) SearchURL(context.Context, *URLRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchURL not implemented")
}
func (UnimplementedCheckServer) SearchDomain(context.Context, *DomainRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchDomain not implemented")
}
func (UnimplementedCheckServer) SearchDecision(context.Context, *DecisionRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchDecision not implemented")
}
func (UnimplementedCheckServer) SearchTextDecision(context.Context, *TextDecisionRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchTextDecision not implemented")
}
func (UnimplementedCheckServer) SearchSubnet4(context.Context, *Subnet4Request) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchSubnet4 not implemented")
}
func (UnimplementedCheckServer) SearchSubnet6(context.Context, *Subnet6Request) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchSubnet6 not implemented")
}
func (UnimplementedCheckServer) Stat(context.Context, *StatRequest) (*StatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedCheckServer) Ping(context.Context, *PingRequest) (*PongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedCheckServer) mustEmbedUnimplementedCheckServer() {}

// UnsafeCheckServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CheckServer will
// result in compilation errors.
type UnsafeCheckServer interface {
	mustEmbedUnimplementedCheckServer()
}

func RegisterCheckServer(s grpc.ServiceRegistrar, srv CheckServer) {
	s.RegisterService(&Check_ServiceDesc, srv)
}

func _Check_SearchID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckServer).SearchID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.Check/SearchID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckServer).SearchID(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Check_SearchIP4_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IP4Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckServer).SearchIP4(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.Check/SearchIP4",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckServer).SearchIP4(ctx, req.(*IP4Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Check_SearchIP6_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IP6Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckServer).SearchIP6(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.Check/SearchIP6",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckServer).SearchIP6(ctx, req.(*IP6Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Check_SearchURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(URLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckServer).SearchURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.Check/SearchURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckServer).SearchURL(ctx, req.(*URLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Check_SearchDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckServer).SearchDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.Check/SearchDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckServer).SearchDomain(ctx, req.(*DomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Check_SearchDecision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckServer).SearchDecision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.Check/SearchDecision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckServer).SearchDecision(ctx, req.(*DecisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Check_SearchTextDecision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextDecisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckServer).SearchTextDecision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.Check/SearchTextDecision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckServer).SearchTextDecision(ctx, req.(*TextDecisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Check_SearchSubnet4_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Subnet4Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckServer).SearchSubnet4(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.Check/SearchSubnet4",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckServer).SearchSubnet4(ctx, req.(*Subnet4Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Check_SearchSubnet6_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Subnet6Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckServer).SearchSubnet6(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.Check/SearchSubnet6",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckServer).SearchSubnet6(ctx, req.(*Subnet6Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Check_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.Check/Stat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckServer).Stat(ctx, req.(*StatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Check_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.Check/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Check_ServiceDesc is the grpc.ServiceDesc for Check service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Check_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "msg.Check",
	HandlerType: (*CheckServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchID",
			Handler:    _Check_SearchID_Handler,
		},
		{
			MethodName: "SearchIP4",
			Handler:    _Check_SearchIP4_Handler,
		},
		{
			MethodName: "SearchIP6",
			Handler:    _Check_SearchIP6_Handler,
		},
		{
			MethodName: "SearchURL",
			Handler:    _Check_SearchURL_Handler,
		},
		{
			MethodName: "SearchDomain",
			Handler:    _Check_SearchDomain_Handler,
		},
		{
			MethodName: "SearchDecision",
			Handler:    _Check_SearchDecision_Handler,
		},
		{
			MethodName: "SearchTextDecision",
			Handler:    _Check_SearchTextDecision_Handler,
		},
		{
			MethodName: "SearchSubnet4",
			Handler:    _Check_SearchSubnet4_Handler,
		},
		{
			MethodName: "SearchSubnet6",
			Handler:    _Check_SearchSubnet6_Handler,
		},
		{
			MethodName: "Stat",
			Handler:    _Check_Stat_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Check_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "msg.proto",
}
