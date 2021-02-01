// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: clustermanager.proto

package clustermanager

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

// Api Endpoints for ClusterManager service

func NewClusterManagerEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		&api.Endpoint{
			Name:    "ClusterManager.CreateCluster",
			Path:    []string{"/clustermanager/v1/cluster/{clusterID}"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ClusterManager.UpdateCluster",
			Path:    []string{"/clustermanager/v1/cluster/{clusterID}"},
			Method:  []string{"PUT"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ClusterManager.DeleteCluster",
			Path:    []string{"/clustermanager/v1/cluster/{clusterID}"},
			Method:  []string{"DELETE"},
			Body:    "",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ClusterManager.GetCluster",
			Path:    []string{"/clustermanager/v1/cluster/{clusterID}"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ClusterManager.ListCluster",
			Path:    []string{"/clustermanager/v1/cluster"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ClusterManager.GetClusterCredential",
			Path:    []string{"/clustermanager/v1/clustercredential/{serverKey}"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ClusterManager.UpdateClusterCredential",
			Path:    []string{"/clustermanager/v1/clustercredential/{serverKey}"},
			Method:  []string{"PUT"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ClusterManager.ListClusterCredential",
			Path:    []string{"/clustermanager/v1/clustercredential"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ClusterManager.InitFederationCluster",
			Path:    []string{"/clustermanager/v1/initfedcluster"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ClusterManager.AddFederatedCluster",
			Path:    []string{"/clustermanager/v1/addfederatedcluster"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ClusterManager.CreateNamespace",
			Path:    []string{"/clustermanager/v1/namespace/{federationClusterID}/{name}"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ClusterManager.UpdateNamespace",
			Path:    []string{"/clustermanager/v1/namespace/{federationClusterID}/{name}"},
			Method:  []string{"PUT"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ClusterManager.DeleteNamespace",
			Path:    []string{"/clustermanager/v1/namespace/{federationClusterID}/{name}"},
			Method:  []string{"DELETE"},
			Body:    "",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ClusterManager.GetNamespace",
			Path:    []string{"/clustermanager/v1/namespace/{federationClusterID}/{name}"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ClusterManager.ListNamespace",
			Path:    []string{"/clustermanager/v1/namespace"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ClusterManager.CreateNamespaceQuota",
			Path:    []string{"/clustermanager/v1/namespacequota/{federationClusterID}/{namespace}"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ClusterManager.UpdateNamespaceQuota",
			Path:    []string{"/clustermanager/v1/namespacequota/{federationClusterID}/{namespace}"},
			Method:  []string{"PUT"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ClusterManager.DeleteNamespaceQuota",
			Path:    []string{"/clustermanager/v1/namespacequota/{federationClusterID}/{namespace}"},
			Method:  []string{"DELETE"},
			Body:    "",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ClusterManager.GetNamespaceQuota",
			Path:    []string{"/clustermanager/v1/namespacequota/{federationClusterID}/{namespace}"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ClusterManager.ListNamespaceQuota",
			Path:    []string{"/clustermanager/v1/namespacequota"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ClusterManager.CreateNamespaceWithQuota",
			Path:    []string{"/clustermanager/v1/namespacewithquota"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for ClusterManager service

type ClusterManagerService interface {
	CreateCluster(ctx context.Context, in *CreateClusterReq, opts ...client.CallOption) (*CreateClusterResp, error)
	UpdateCluster(ctx context.Context, in *UpdateClusterReq, opts ...client.CallOption) (*UpdateClusterResp, error)
	DeleteCluster(ctx context.Context, in *DeleteClusterReq, opts ...client.CallOption) (*DeleteClusterResp, error)
	GetCluster(ctx context.Context, in *GetClusterReq, opts ...client.CallOption) (*GetClusterResp, error)
	ListCluster(ctx context.Context, in *ListClusterReq, opts ...client.CallOption) (*ListClusterResp, error)
	GetClusterCredential(ctx context.Context, in *GetClusterCredentialReq, opts ...client.CallOption) (*GetClusterCredentialResp, error)
	UpdateClusterCredential(ctx context.Context, in *UpdateClusterCredentialReq, opts ...client.CallOption) (*UpdateClusterCredentialResp, error)
	ListClusterCredential(ctx context.Context, in *ListClusterCredentialReq, opts ...client.CallOption) (*ListClusterCredentialResp, error)
	InitFederationCluster(ctx context.Context, in *InitFederationClusterReq, opts ...client.CallOption) (*InitFederationClusterResp, error)
	AddFederatedCluster(ctx context.Context, in *AddFederatedClusterReq, opts ...client.CallOption) (*AddFederatedClusterResp, error)
	CreateNamespace(ctx context.Context, in *CreateNamespaceReq, opts ...client.CallOption) (*CreateNamespaceResp, error)
	UpdateNamespace(ctx context.Context, in *UpdateNamespaceReq, opts ...client.CallOption) (*UpdateNamespaceResp, error)
	DeleteNamespace(ctx context.Context, in *DeleteNamespaceReq, opts ...client.CallOption) (*DeleteNamespaceResp, error)
	GetNamespace(ctx context.Context, in *GetNamespaceReq, opts ...client.CallOption) (*GetNamespaceResp, error)
	ListNamespace(ctx context.Context, in *ListNamespaceReq, opts ...client.CallOption) (*ListNamespaceResp, error)
	CreateNamespaceQuota(ctx context.Context, in *CreateNamespaceQuotaReq, opts ...client.CallOption) (*CreateNamespaceQuotaResp, error)
	UpdateNamespaceQuota(ctx context.Context, in *UpdateNamespaceQuotaReq, opts ...client.CallOption) (*UpdateNamespaceQuotaResp, error)
	DeleteNamespaceQuota(ctx context.Context, in *DeleteNamespaceQuotaReq, opts ...client.CallOption) (*DeleteNamespaceQuotaResp, error)
	GetNamespaceQuota(ctx context.Context, in *GetNamespaceQuotaReq, opts ...client.CallOption) (*GetNamespaceQuotaResp, error)
	ListNamespaceQuota(ctx context.Context, in *ListNamespaceQuotaReq, opts ...client.CallOption) (*ListNamespaceQuotaResp, error)
	CreateNamespaceWithQuota(ctx context.Context, in *CreateNamespaceWithQuotaReq, opts ...client.CallOption) (*CreateNamespaceWithQuotaResp, error)
}

type clusterManagerService struct {
	c    client.Client
	name string
}

func NewClusterManagerService(name string, c client.Client) ClusterManagerService {
	return &clusterManagerService{
		c:    c,
		name: name,
	}
}

func (c *clusterManagerService) CreateCluster(ctx context.Context, in *CreateClusterReq, opts ...client.CallOption) (*CreateClusterResp, error) {
	req := c.c.NewRequest(c.name, "ClusterManager.CreateCluster", in)
	out := new(CreateClusterResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterManagerService) UpdateCluster(ctx context.Context, in *UpdateClusterReq, opts ...client.CallOption) (*UpdateClusterResp, error) {
	req := c.c.NewRequest(c.name, "ClusterManager.UpdateCluster", in)
	out := new(UpdateClusterResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterManagerService) DeleteCluster(ctx context.Context, in *DeleteClusterReq, opts ...client.CallOption) (*DeleteClusterResp, error) {
	req := c.c.NewRequest(c.name, "ClusterManager.DeleteCluster", in)
	out := new(DeleteClusterResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterManagerService) GetCluster(ctx context.Context, in *GetClusterReq, opts ...client.CallOption) (*GetClusterResp, error) {
	req := c.c.NewRequest(c.name, "ClusterManager.GetCluster", in)
	out := new(GetClusterResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterManagerService) ListCluster(ctx context.Context, in *ListClusterReq, opts ...client.CallOption) (*ListClusterResp, error) {
	req := c.c.NewRequest(c.name, "ClusterManager.ListCluster", in)
	out := new(ListClusterResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterManagerService) GetClusterCredential(ctx context.Context, in *GetClusterCredentialReq, opts ...client.CallOption) (*GetClusterCredentialResp, error) {
	req := c.c.NewRequest(c.name, "ClusterManager.GetClusterCredential", in)
	out := new(GetClusterCredentialResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterManagerService) UpdateClusterCredential(ctx context.Context, in *UpdateClusterCredentialReq, opts ...client.CallOption) (*UpdateClusterCredentialResp, error) {
	req := c.c.NewRequest(c.name, "ClusterManager.UpdateClusterCredential", in)
	out := new(UpdateClusterCredentialResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterManagerService) ListClusterCredential(ctx context.Context, in *ListClusterCredentialReq, opts ...client.CallOption) (*ListClusterCredentialResp, error) {
	req := c.c.NewRequest(c.name, "ClusterManager.ListClusterCredential", in)
	out := new(ListClusterCredentialResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterManagerService) InitFederationCluster(ctx context.Context, in *InitFederationClusterReq, opts ...client.CallOption) (*InitFederationClusterResp, error) {
	req := c.c.NewRequest(c.name, "ClusterManager.InitFederationCluster", in)
	out := new(InitFederationClusterResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterManagerService) AddFederatedCluster(ctx context.Context, in *AddFederatedClusterReq, opts ...client.CallOption) (*AddFederatedClusterResp, error) {
	req := c.c.NewRequest(c.name, "ClusterManager.AddFederatedCluster", in)
	out := new(AddFederatedClusterResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterManagerService) CreateNamespace(ctx context.Context, in *CreateNamespaceReq, opts ...client.CallOption) (*CreateNamespaceResp, error) {
	req := c.c.NewRequest(c.name, "ClusterManager.CreateNamespace", in)
	out := new(CreateNamespaceResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterManagerService) UpdateNamespace(ctx context.Context, in *UpdateNamespaceReq, opts ...client.CallOption) (*UpdateNamespaceResp, error) {
	req := c.c.NewRequest(c.name, "ClusterManager.UpdateNamespace", in)
	out := new(UpdateNamespaceResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterManagerService) DeleteNamespace(ctx context.Context, in *DeleteNamespaceReq, opts ...client.CallOption) (*DeleteNamespaceResp, error) {
	req := c.c.NewRequest(c.name, "ClusterManager.DeleteNamespace", in)
	out := new(DeleteNamespaceResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterManagerService) GetNamespace(ctx context.Context, in *GetNamespaceReq, opts ...client.CallOption) (*GetNamespaceResp, error) {
	req := c.c.NewRequest(c.name, "ClusterManager.GetNamespace", in)
	out := new(GetNamespaceResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterManagerService) ListNamespace(ctx context.Context, in *ListNamespaceReq, opts ...client.CallOption) (*ListNamespaceResp, error) {
	req := c.c.NewRequest(c.name, "ClusterManager.ListNamespace", in)
	out := new(ListNamespaceResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterManagerService) CreateNamespaceQuota(ctx context.Context, in *CreateNamespaceQuotaReq, opts ...client.CallOption) (*CreateNamespaceQuotaResp, error) {
	req := c.c.NewRequest(c.name, "ClusterManager.CreateNamespaceQuota", in)
	out := new(CreateNamespaceQuotaResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterManagerService) UpdateNamespaceQuota(ctx context.Context, in *UpdateNamespaceQuotaReq, opts ...client.CallOption) (*UpdateNamespaceQuotaResp, error) {
	req := c.c.NewRequest(c.name, "ClusterManager.UpdateNamespaceQuota", in)
	out := new(UpdateNamespaceQuotaResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterManagerService) DeleteNamespaceQuota(ctx context.Context, in *DeleteNamespaceQuotaReq, opts ...client.CallOption) (*DeleteNamespaceQuotaResp, error) {
	req := c.c.NewRequest(c.name, "ClusterManager.DeleteNamespaceQuota", in)
	out := new(DeleteNamespaceQuotaResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterManagerService) GetNamespaceQuota(ctx context.Context, in *GetNamespaceQuotaReq, opts ...client.CallOption) (*GetNamespaceQuotaResp, error) {
	req := c.c.NewRequest(c.name, "ClusterManager.GetNamespaceQuota", in)
	out := new(GetNamespaceQuotaResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterManagerService) ListNamespaceQuota(ctx context.Context, in *ListNamespaceQuotaReq, opts ...client.CallOption) (*ListNamespaceQuotaResp, error) {
	req := c.c.NewRequest(c.name, "ClusterManager.ListNamespaceQuota", in)
	out := new(ListNamespaceQuotaResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterManagerService) CreateNamespaceWithQuota(ctx context.Context, in *CreateNamespaceWithQuotaReq, opts ...client.CallOption) (*CreateNamespaceWithQuotaResp, error) {
	req := c.c.NewRequest(c.name, "ClusterManager.CreateNamespaceWithQuota", in)
	out := new(CreateNamespaceWithQuotaResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ClusterManager service

type ClusterManagerHandler interface {
	CreateCluster(context.Context, *CreateClusterReq, *CreateClusterResp) error
	UpdateCluster(context.Context, *UpdateClusterReq, *UpdateClusterResp) error
	DeleteCluster(context.Context, *DeleteClusterReq, *DeleteClusterResp) error
	GetCluster(context.Context, *GetClusterReq, *GetClusterResp) error
	ListCluster(context.Context, *ListClusterReq, *ListClusterResp) error
	GetClusterCredential(context.Context, *GetClusterCredentialReq, *GetClusterCredentialResp) error
	UpdateClusterCredential(context.Context, *UpdateClusterCredentialReq, *UpdateClusterCredentialResp) error
	ListClusterCredential(context.Context, *ListClusterCredentialReq, *ListClusterCredentialResp) error
	InitFederationCluster(context.Context, *InitFederationClusterReq, *InitFederationClusterResp) error
	AddFederatedCluster(context.Context, *AddFederatedClusterReq, *AddFederatedClusterResp) error
	CreateNamespace(context.Context, *CreateNamespaceReq, *CreateNamespaceResp) error
	UpdateNamespace(context.Context, *UpdateNamespaceReq, *UpdateNamespaceResp) error
	DeleteNamespace(context.Context, *DeleteNamespaceReq, *DeleteNamespaceResp) error
	GetNamespace(context.Context, *GetNamespaceReq, *GetNamespaceResp) error
	ListNamespace(context.Context, *ListNamespaceReq, *ListNamespaceResp) error
	CreateNamespaceQuota(context.Context, *CreateNamespaceQuotaReq, *CreateNamespaceQuotaResp) error
	UpdateNamespaceQuota(context.Context, *UpdateNamespaceQuotaReq, *UpdateNamespaceQuotaResp) error
	DeleteNamespaceQuota(context.Context, *DeleteNamespaceQuotaReq, *DeleteNamespaceQuotaResp) error
	GetNamespaceQuota(context.Context, *GetNamespaceQuotaReq, *GetNamespaceQuotaResp) error
	ListNamespaceQuota(context.Context, *ListNamespaceQuotaReq, *ListNamespaceQuotaResp) error
	CreateNamespaceWithQuota(context.Context, *CreateNamespaceWithQuotaReq, *CreateNamespaceWithQuotaResp) error
}

func RegisterClusterManagerHandler(s server.Server, hdlr ClusterManagerHandler, opts ...server.HandlerOption) error {
	type clusterManager interface {
		CreateCluster(ctx context.Context, in *CreateClusterReq, out *CreateClusterResp) error
		UpdateCluster(ctx context.Context, in *UpdateClusterReq, out *UpdateClusterResp) error
		DeleteCluster(ctx context.Context, in *DeleteClusterReq, out *DeleteClusterResp) error
		GetCluster(ctx context.Context, in *GetClusterReq, out *GetClusterResp) error
		ListCluster(ctx context.Context, in *ListClusterReq, out *ListClusterResp) error
		GetClusterCredential(ctx context.Context, in *GetClusterCredentialReq, out *GetClusterCredentialResp) error
		UpdateClusterCredential(ctx context.Context, in *UpdateClusterCredentialReq, out *UpdateClusterCredentialResp) error
		ListClusterCredential(ctx context.Context, in *ListClusterCredentialReq, out *ListClusterCredentialResp) error
		InitFederationCluster(ctx context.Context, in *InitFederationClusterReq, out *InitFederationClusterResp) error
		AddFederatedCluster(ctx context.Context, in *AddFederatedClusterReq, out *AddFederatedClusterResp) error
		CreateNamespace(ctx context.Context, in *CreateNamespaceReq, out *CreateNamespaceResp) error
		UpdateNamespace(ctx context.Context, in *UpdateNamespaceReq, out *UpdateNamespaceResp) error
		DeleteNamespace(ctx context.Context, in *DeleteNamespaceReq, out *DeleteNamespaceResp) error
		GetNamespace(ctx context.Context, in *GetNamespaceReq, out *GetNamespaceResp) error
		ListNamespace(ctx context.Context, in *ListNamespaceReq, out *ListNamespaceResp) error
		CreateNamespaceQuota(ctx context.Context, in *CreateNamespaceQuotaReq, out *CreateNamespaceQuotaResp) error
		UpdateNamespaceQuota(ctx context.Context, in *UpdateNamespaceQuotaReq, out *UpdateNamespaceQuotaResp) error
		DeleteNamespaceQuota(ctx context.Context, in *DeleteNamespaceQuotaReq, out *DeleteNamespaceQuotaResp) error
		GetNamespaceQuota(ctx context.Context, in *GetNamespaceQuotaReq, out *GetNamespaceQuotaResp) error
		ListNamespaceQuota(ctx context.Context, in *ListNamespaceQuotaReq, out *ListNamespaceQuotaResp) error
		CreateNamespaceWithQuota(ctx context.Context, in *CreateNamespaceWithQuotaReq, out *CreateNamespaceWithQuotaResp) error
	}
	type ClusterManager struct {
		clusterManager
	}
	h := &clusterManagerHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterManager.CreateCluster",
		Path:    []string{"/clustermanager/v1/cluster/{clusterID}"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterManager.UpdateCluster",
		Path:    []string{"/clustermanager/v1/cluster/{clusterID}"},
		Method:  []string{"PUT"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterManager.DeleteCluster",
		Path:    []string{"/clustermanager/v1/cluster/{clusterID}"},
		Method:  []string{"DELETE"},
		Body:    "",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterManager.GetCluster",
		Path:    []string{"/clustermanager/v1/cluster/{clusterID}"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterManager.ListCluster",
		Path:    []string{"/clustermanager/v1/cluster"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterManager.GetClusterCredential",
		Path:    []string{"/clustermanager/v1/clustercredential/{serverKey}"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterManager.UpdateClusterCredential",
		Path:    []string{"/clustermanager/v1/clustercredential/{serverKey}"},
		Method:  []string{"PUT"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterManager.ListClusterCredential",
		Path:    []string{"/clustermanager/v1/clustercredential"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterManager.InitFederationCluster",
		Path:    []string{"/clustermanager/v1/initfedcluster"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterManager.AddFederatedCluster",
		Path:    []string{"/clustermanager/v1/addfederatedcluster"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterManager.CreateNamespace",
		Path:    []string{"/clustermanager/v1/namespace/{federationClusterID}/{name}"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterManager.UpdateNamespace",
		Path:    []string{"/clustermanager/v1/namespace/{federationClusterID}/{name}"},
		Method:  []string{"PUT"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterManager.DeleteNamespace",
		Path:    []string{"/clustermanager/v1/namespace/{federationClusterID}/{name}"},
		Method:  []string{"DELETE"},
		Body:    "",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterManager.GetNamespace",
		Path:    []string{"/clustermanager/v1/namespace/{federationClusterID}/{name}"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterManager.ListNamespace",
		Path:    []string{"/clustermanager/v1/namespace"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterManager.CreateNamespaceQuota",
		Path:    []string{"/clustermanager/v1/namespacequota/{federationClusterID}/{namespace}"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterManager.UpdateNamespaceQuota",
		Path:    []string{"/clustermanager/v1/namespacequota/{federationClusterID}/{namespace}"},
		Method:  []string{"PUT"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterManager.DeleteNamespaceQuota",
		Path:    []string{"/clustermanager/v1/namespacequota/{federationClusterID}/{namespace}"},
		Method:  []string{"DELETE"},
		Body:    "",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterManager.GetNamespaceQuota",
		Path:    []string{"/clustermanager/v1/namespacequota/{federationClusterID}/{namespace}"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterManager.ListNamespaceQuota",
		Path:    []string{"/clustermanager/v1/namespacequota"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterManager.CreateNamespaceWithQuota",
		Path:    []string{"/clustermanager/v1/namespacewithquota"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&ClusterManager{h}, opts...))
}

type clusterManagerHandler struct {
	ClusterManagerHandler
}

func (h *clusterManagerHandler) CreateCluster(ctx context.Context, in *CreateClusterReq, out *CreateClusterResp) error {
	return h.ClusterManagerHandler.CreateCluster(ctx, in, out)
}

func (h *clusterManagerHandler) UpdateCluster(ctx context.Context, in *UpdateClusterReq, out *UpdateClusterResp) error {
	return h.ClusterManagerHandler.UpdateCluster(ctx, in, out)
}

func (h *clusterManagerHandler) DeleteCluster(ctx context.Context, in *DeleteClusterReq, out *DeleteClusterResp) error {
	return h.ClusterManagerHandler.DeleteCluster(ctx, in, out)
}

func (h *clusterManagerHandler) GetCluster(ctx context.Context, in *GetClusterReq, out *GetClusterResp) error {
	return h.ClusterManagerHandler.GetCluster(ctx, in, out)
}

func (h *clusterManagerHandler) ListCluster(ctx context.Context, in *ListClusterReq, out *ListClusterResp) error {
	return h.ClusterManagerHandler.ListCluster(ctx, in, out)
}

func (h *clusterManagerHandler) GetClusterCredential(ctx context.Context, in *GetClusterCredentialReq, out *GetClusterCredentialResp) error {
	return h.ClusterManagerHandler.GetClusterCredential(ctx, in, out)
}

func (h *clusterManagerHandler) UpdateClusterCredential(ctx context.Context, in *UpdateClusterCredentialReq, out *UpdateClusterCredentialResp) error {
	return h.ClusterManagerHandler.UpdateClusterCredential(ctx, in, out)
}

func (h *clusterManagerHandler) ListClusterCredential(ctx context.Context, in *ListClusterCredentialReq, out *ListClusterCredentialResp) error {
	return h.ClusterManagerHandler.ListClusterCredential(ctx, in, out)
}

func (h *clusterManagerHandler) InitFederationCluster(ctx context.Context, in *InitFederationClusterReq, out *InitFederationClusterResp) error {
	return h.ClusterManagerHandler.InitFederationCluster(ctx, in, out)
}

func (h *clusterManagerHandler) AddFederatedCluster(ctx context.Context, in *AddFederatedClusterReq, out *AddFederatedClusterResp) error {
	return h.ClusterManagerHandler.AddFederatedCluster(ctx, in, out)
}

func (h *clusterManagerHandler) CreateNamespace(ctx context.Context, in *CreateNamespaceReq, out *CreateNamespaceResp) error {
	return h.ClusterManagerHandler.CreateNamespace(ctx, in, out)
}

func (h *clusterManagerHandler) UpdateNamespace(ctx context.Context, in *UpdateNamespaceReq, out *UpdateNamespaceResp) error {
	return h.ClusterManagerHandler.UpdateNamespace(ctx, in, out)
}

func (h *clusterManagerHandler) DeleteNamespace(ctx context.Context, in *DeleteNamespaceReq, out *DeleteNamespaceResp) error {
	return h.ClusterManagerHandler.DeleteNamespace(ctx, in, out)
}

func (h *clusterManagerHandler) GetNamespace(ctx context.Context, in *GetNamespaceReq, out *GetNamespaceResp) error {
	return h.ClusterManagerHandler.GetNamespace(ctx, in, out)
}

func (h *clusterManagerHandler) ListNamespace(ctx context.Context, in *ListNamespaceReq, out *ListNamespaceResp) error {
	return h.ClusterManagerHandler.ListNamespace(ctx, in, out)
}

func (h *clusterManagerHandler) CreateNamespaceQuota(ctx context.Context, in *CreateNamespaceQuotaReq, out *CreateNamespaceQuotaResp) error {
	return h.ClusterManagerHandler.CreateNamespaceQuota(ctx, in, out)
}

func (h *clusterManagerHandler) UpdateNamespaceQuota(ctx context.Context, in *UpdateNamespaceQuotaReq, out *UpdateNamespaceQuotaResp) error {
	return h.ClusterManagerHandler.UpdateNamespaceQuota(ctx, in, out)
}

func (h *clusterManagerHandler) DeleteNamespaceQuota(ctx context.Context, in *DeleteNamespaceQuotaReq, out *DeleteNamespaceQuotaResp) error {
	return h.ClusterManagerHandler.DeleteNamespaceQuota(ctx, in, out)
}

func (h *clusterManagerHandler) GetNamespaceQuota(ctx context.Context, in *GetNamespaceQuotaReq, out *GetNamespaceQuotaResp) error {
	return h.ClusterManagerHandler.GetNamespaceQuota(ctx, in, out)
}

func (h *clusterManagerHandler) ListNamespaceQuota(ctx context.Context, in *ListNamespaceQuotaReq, out *ListNamespaceQuotaResp) error {
	return h.ClusterManagerHandler.ListNamespaceQuota(ctx, in, out)
}

func (h *clusterManagerHandler) CreateNamespaceWithQuota(ctx context.Context, in *CreateNamespaceWithQuotaReq, out *CreateNamespaceWithQuotaResp) error {
	return h.ClusterManagerHandler.CreateNamespaceWithQuota(ctx, in, out)
}
