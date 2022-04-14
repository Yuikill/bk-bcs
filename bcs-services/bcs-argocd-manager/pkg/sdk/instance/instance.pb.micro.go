// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: pkg/sdk/instance/instance.proto

package instance

import (
	fmt "fmt"
	_ "github.com/Tencent/bk-bcs/bcs-services/bcs-argocd-manager/pkg/apis/tkex/v1alpha1"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Instance service

func NewInstanceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "Instance.CreateArgocdInstance",
			Path:    []string{"/argocdmanager/v1/instance"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "Instance.UpdateArgocdInstance",
			Path:    []string{"/argocdmanager/v1/instance"},
			Method:  []string{"PUT"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "Instance.DeleteArgocdInstance",
			Path:    []string{"/argocdmanager/v1/instance/{name}"},
			Method:  []string{"DELETE"},
			Body:    "",
			Handler: "rpc",
		},
		{
			Name:    "Instance.GetArgocdInstance",
			Path:    []string{"/argocdmanager/v1/instance/{name}"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		{
			Name:    "Instance.ListArgocdInstances",
			Path:    []string{"/argocdmanager/v1/instances"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
	}
}

// Client API for Instance service

type InstanceService interface {
	CreateArgocdInstance(ctx context.Context, in *CreateArgocdInstanceRequest, opts ...client.CallOption) (*CreateArgocdInstanceResponse, error)
	UpdateArgocdInstance(ctx context.Context, in *UpdateArgocdInstanceRequest, opts ...client.CallOption) (*UpdateArgocdInstanceResponse, error)
	DeleteArgocdInstance(ctx context.Context, in *DeleteArgocdInstanceRequest, opts ...client.CallOption) (*DeleteArgocdInstanceResponse, error)
	GetArgocdInstance(ctx context.Context, in *GetArgocdInstanceRequest, opts ...client.CallOption) (*GetArgocdInstanceResponse, error)
	ListArgocdInstances(ctx context.Context, in *ListArgocdInstancesRequest, opts ...client.CallOption) (*ListArgocdInstancesResponse, error)
}

type instanceService struct {
	c    client.Client
	name string
}

func NewInstanceService(name string, c client.Client) InstanceService {
	return &instanceService{
		c:    c,
		name: name,
	}
}

func (c *instanceService) CreateArgocdInstance(ctx context.Context, in *CreateArgocdInstanceRequest, opts ...client.CallOption) (*CreateArgocdInstanceResponse, error) {
	req := c.c.NewRequest(c.name, "Instance.CreateArgocdInstance", in)
	out := new(CreateArgocdInstanceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceService) UpdateArgocdInstance(ctx context.Context, in *UpdateArgocdInstanceRequest, opts ...client.CallOption) (*UpdateArgocdInstanceResponse, error) {
	req := c.c.NewRequest(c.name, "Instance.UpdateArgocdInstance", in)
	out := new(UpdateArgocdInstanceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceService) DeleteArgocdInstance(ctx context.Context, in *DeleteArgocdInstanceRequest, opts ...client.CallOption) (*DeleteArgocdInstanceResponse, error) {
	req := c.c.NewRequest(c.name, "Instance.DeleteArgocdInstance", in)
	out := new(DeleteArgocdInstanceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceService) GetArgocdInstance(ctx context.Context, in *GetArgocdInstanceRequest, opts ...client.CallOption) (*GetArgocdInstanceResponse, error) {
	req := c.c.NewRequest(c.name, "Instance.GetArgocdInstance", in)
	out := new(GetArgocdInstanceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceService) ListArgocdInstances(ctx context.Context, in *ListArgocdInstancesRequest, opts ...client.CallOption) (*ListArgocdInstancesResponse, error) {
	req := c.c.NewRequest(c.name, "Instance.ListArgocdInstances", in)
	out := new(ListArgocdInstancesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Instance service

type InstanceHandler interface {
	CreateArgocdInstance(context.Context, *CreateArgocdInstanceRequest, *CreateArgocdInstanceResponse) error
	UpdateArgocdInstance(context.Context, *UpdateArgocdInstanceRequest, *UpdateArgocdInstanceResponse) error
	DeleteArgocdInstance(context.Context, *DeleteArgocdInstanceRequest, *DeleteArgocdInstanceResponse) error
	GetArgocdInstance(context.Context, *GetArgocdInstanceRequest, *GetArgocdInstanceResponse) error
	ListArgocdInstances(context.Context, *ListArgocdInstancesRequest, *ListArgocdInstancesResponse) error
}

func RegisterInstanceHandler(s server.Server, hdlr InstanceHandler, opts ...server.HandlerOption) error {
	type instance interface {
		CreateArgocdInstance(ctx context.Context, in *CreateArgocdInstanceRequest, out *CreateArgocdInstanceResponse) error
		UpdateArgocdInstance(ctx context.Context, in *UpdateArgocdInstanceRequest, out *UpdateArgocdInstanceResponse) error
		DeleteArgocdInstance(ctx context.Context, in *DeleteArgocdInstanceRequest, out *DeleteArgocdInstanceResponse) error
		GetArgocdInstance(ctx context.Context, in *GetArgocdInstanceRequest, out *GetArgocdInstanceResponse) error
		ListArgocdInstances(ctx context.Context, in *ListArgocdInstancesRequest, out *ListArgocdInstancesResponse) error
	}
	type Instance struct {
		instance
	}
	h := &instanceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Instance.CreateArgocdInstance",
		Path:    []string{"/argocdmanager/v1/instance"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Instance.UpdateArgocdInstance",
		Path:    []string{"/argocdmanager/v1/instance"},
		Method:  []string{"PUT"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Instance.DeleteArgocdInstance",
		Path:    []string{"/argocdmanager/v1/instance/{name}"},
		Method:  []string{"DELETE"},
		Body:    "",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Instance.GetArgocdInstance",
		Path:    []string{"/argocdmanager/v1/instance/{name}"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Instance.ListArgocdInstances",
		Path:    []string{"/argocdmanager/v1/instances"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&Instance{h}, opts...))
}

type instanceHandler struct {
	InstanceHandler
}

func (h *instanceHandler) CreateArgocdInstance(ctx context.Context, in *CreateArgocdInstanceRequest, out *CreateArgocdInstanceResponse) error {
	return h.InstanceHandler.CreateArgocdInstance(ctx, in, out)
}

func (h *instanceHandler) UpdateArgocdInstance(ctx context.Context, in *UpdateArgocdInstanceRequest, out *UpdateArgocdInstanceResponse) error {
	return h.InstanceHandler.UpdateArgocdInstance(ctx, in, out)
}

func (h *instanceHandler) DeleteArgocdInstance(ctx context.Context, in *DeleteArgocdInstanceRequest, out *DeleteArgocdInstanceResponse) error {
	return h.InstanceHandler.DeleteArgocdInstance(ctx, in, out)
}

func (h *instanceHandler) GetArgocdInstance(ctx context.Context, in *GetArgocdInstanceRequest, out *GetArgocdInstanceResponse) error {
	return h.InstanceHandler.GetArgocdInstance(ctx, in, out)
}

func (h *instanceHandler) ListArgocdInstances(ctx context.Context, in *ListArgocdInstancesRequest, out *ListArgocdInstancesResponse) error {
	return h.InstanceHandler.ListArgocdInstances(ctx, in, out)
}