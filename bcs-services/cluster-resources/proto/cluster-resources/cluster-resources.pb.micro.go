// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/cluster-resources/cluster-resources.proto

package cluster_resources

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for ClusterResources service

func NewClusterResourcesEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		&api.Endpoint{
			Name:    "ClusterResources.Echo",
			Path:    []string{"/clusterresources/v1/echo"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ClusterResources.Ping",
			Path:    []string{"/clusterresources/v1/ping"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ClusterResources.Healthz",
			Path:    []string{"/clusterresources/v1/healthz"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
	}
}

// Client API for ClusterResources service

type ClusterResourcesService interface {
	// 基础类接口
	Echo(ctx context.Context, in *EchoReq, opts ...client.CallOption) (*EchoResp, error)
	Ping(ctx context.Context, in *PingReq, opts ...client.CallOption) (*PingResp, error)
	Healthz(ctx context.Context, in *HealthzReq, opts ...client.CallOption) (*HealthzResp, error)
}

type clusterResourcesService struct {
	c    client.Client
	name string
}

func NewClusterResourcesService(name string, c client.Client) ClusterResourcesService {
	return &clusterResourcesService{
		c:    c,
		name: name,
	}
}

func (c *clusterResourcesService) Echo(ctx context.Context, in *EchoReq, opts ...client.CallOption) (*EchoResp, error) {
	req := c.c.NewRequest(c.name, "ClusterResources.Echo", in)
	out := new(EchoResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterResourcesService) Ping(ctx context.Context, in *PingReq, opts ...client.CallOption) (*PingResp, error) {
	req := c.c.NewRequest(c.name, "ClusterResources.Ping", in)
	out := new(PingResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterResourcesService) Healthz(ctx context.Context, in *HealthzReq, opts ...client.CallOption) (*HealthzResp, error) {
	req := c.c.NewRequest(c.name, "ClusterResources.Healthz", in)
	out := new(HealthzResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ClusterResources service

type ClusterResourcesHandler interface {
	// 基础类接口
	Echo(context.Context, *EchoReq, *EchoResp) error
	Ping(context.Context, *PingReq, *PingResp) error
	Healthz(context.Context, *HealthzReq, *HealthzResp) error
}

func RegisterClusterResourcesHandler(s server.Server, hdlr ClusterResourcesHandler, opts ...server.HandlerOption) error {
	type clusterResources interface {
		Echo(ctx context.Context, in *EchoReq, out *EchoResp) error
		Ping(ctx context.Context, in *PingReq, out *PingResp) error
		Healthz(ctx context.Context, in *HealthzReq, out *HealthzResp) error
	}
	type ClusterResources struct {
		clusterResources
	}
	h := &clusterResourcesHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterResources.Echo",
		Path:    []string{"/clusterresources/v1/echo"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterResources.Ping",
		Path:    []string{"/clusterresources/v1/ping"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterResources.Healthz",
		Path:    []string{"/clusterresources/v1/healthz"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&ClusterResources{h}, opts...))
}

type clusterResourcesHandler struct {
	ClusterResourcesHandler
}

func (h *clusterResourcesHandler) Echo(ctx context.Context, in *EchoReq, out *EchoResp) error {
	return h.ClusterResourcesHandler.Echo(ctx, in, out)
}

func (h *clusterResourcesHandler) Ping(ctx context.Context, in *PingReq, out *PingResp) error {
	return h.ClusterResourcesHandler.Ping(ctx, in, out)
}

func (h *clusterResourcesHandler) Healthz(ctx context.Context, in *HealthzReq, out *HealthzResp) error {
	return h.ClusterResourcesHandler.Healthz(ctx, in, out)
}