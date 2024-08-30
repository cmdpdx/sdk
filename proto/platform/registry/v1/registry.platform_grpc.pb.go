// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: registry.platform.proto

package v1

import (
	v1 "chainguard.dev/sdk/proto/platform/tenant/v1"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Registry_CreateRepo_FullMethodName                = "/chainguard.platform.registry.Registry/CreateRepo"
	Registry_UpdateRepo_FullMethodName                = "/chainguard.platform.registry.Registry/UpdateRepo"
	Registry_ListRepos_FullMethodName                 = "/chainguard.platform.registry.Registry/ListRepos"
	Registry_DeleteRepo_FullMethodName                = "/chainguard.platform.registry.Registry/DeleteRepo"
	Registry_CreateTag_FullMethodName                 = "/chainguard.platform.registry.Registry/CreateTag"
	Registry_UpdateTag_FullMethodName                 = "/chainguard.platform.registry.Registry/UpdateTag"
	Registry_DeleteTag_FullMethodName                 = "/chainguard.platform.registry.Registry/DeleteTag"
	Registry_ListTags_FullMethodName                  = "/chainguard.platform.registry.Registry/ListTags"
	Registry_ListTagHistory_FullMethodName            = "/chainguard.platform.registry.Registry/ListTagHistory"
	Registry_GetSbom_FullMethodName                   = "/chainguard.platform.registry.Registry/GetSbom"
	Registry_GetImageConfig_FullMethodName            = "/chainguard.platform.registry.Registry/GetImageConfig"
	Registry_GetArchs_FullMethodName                  = "/chainguard.platform.registry.Registry/GetArchs"
	Registry_GetSize_FullMethodName                   = "/chainguard.platform.registry.Registry/GetSize"
	Registry_GetRawSbom_FullMethodName                = "/chainguard.platform.registry.Registry/GetRawSbom"
	Registry_GetVulnReport_FullMethodName             = "/chainguard.platform.registry.Registry/GetVulnReport"
	Registry_ListManifestMetadata_FullMethodName      = "/chainguard.platform.registry.Registry/ListManifestMetadata"
	Registry_GetPackageVersionMetadata_FullMethodName = "/chainguard.platform.registry.Registry/GetPackageVersionMetadata"
)

// RegistryClient is the client API for Registry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegistryClient interface {
	CreateRepo(ctx context.Context, in *CreateRepoRequest, opts ...grpc.CallOption) (*Repo, error)
	UpdateRepo(ctx context.Context, in *Repo, opts ...grpc.CallOption) (*Repo, error)
	ListRepos(ctx context.Context, in *RepoFilter, opts ...grpc.CallOption) (*RepoList, error)
	DeleteRepo(ctx context.Context, in *DeleteRepoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateTag(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (*Tag, error)
	UpdateTag(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*Tag, error)
	DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListTags(ctx context.Context, in *TagFilter, opts ...grpc.CallOption) (*TagList, error)
	ListTagHistory(ctx context.Context, in *TagHistoryFilter, opts ...grpc.CallOption) (*TagHistoryList, error)
	GetSbom(ctx context.Context, in *SbomRequest, opts ...grpc.CallOption) (*v1.Sbom2, error)
	GetImageConfig(ctx context.Context, in *ImageConfigRequest, opts ...grpc.CallOption) (*ImageConfig, error)
	GetArchs(ctx context.Context, in *ArchRequest, opts ...grpc.CallOption) (*Archs, error)
	GetSize(ctx context.Context, in *SizeRequest, opts ...grpc.CallOption) (*Size, error)
	GetRawSbom(ctx context.Context, in *RawSbomRequest, opts ...grpc.CallOption) (*RawSbom, error)
	GetVulnReport(ctx context.Context, in *VulnReportRequest, opts ...grpc.CallOption) (*v1.VulnReport, error)
	ListManifestMetadata(ctx context.Context, in *ManifestMetadataFilter, opts ...grpc.CallOption) (*ManifestMetadataList, error)
	GetPackageVersionMetadata(ctx context.Context, in *PackageVersionMetadataRequest, opts ...grpc.CallOption) (*PackageVersionMetadata, error)
}

type registryClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistryClient(cc grpc.ClientConnInterface) RegistryClient {
	return &registryClient{cc}
}

func (c *registryClient) CreateRepo(ctx context.Context, in *CreateRepoRequest, opts ...grpc.CallOption) (*Repo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Repo)
	err := c.cc.Invoke(ctx, Registry_CreateRepo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) UpdateRepo(ctx context.Context, in *Repo, opts ...grpc.CallOption) (*Repo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Repo)
	err := c.cc.Invoke(ctx, Registry_UpdateRepo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) ListRepos(ctx context.Context, in *RepoFilter, opts ...grpc.CallOption) (*RepoList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RepoList)
	err := c.cc.Invoke(ctx, Registry_ListRepos_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) DeleteRepo(ctx context.Context, in *DeleteRepoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Registry_DeleteRepo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) CreateTag(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (*Tag, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Tag)
	err := c.cc.Invoke(ctx, Registry_CreateTag_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) UpdateTag(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*Tag, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Tag)
	err := c.cc.Invoke(ctx, Registry_UpdateTag_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Registry_DeleteTag_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) ListTags(ctx context.Context, in *TagFilter, opts ...grpc.CallOption) (*TagList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TagList)
	err := c.cc.Invoke(ctx, Registry_ListTags_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) ListTagHistory(ctx context.Context, in *TagHistoryFilter, opts ...grpc.CallOption) (*TagHistoryList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TagHistoryList)
	err := c.cc.Invoke(ctx, Registry_ListTagHistory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) GetSbom(ctx context.Context, in *SbomRequest, opts ...grpc.CallOption) (*v1.Sbom2, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.Sbom2)
	err := c.cc.Invoke(ctx, Registry_GetSbom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) GetImageConfig(ctx context.Context, in *ImageConfigRequest, opts ...grpc.CallOption) (*ImageConfig, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ImageConfig)
	err := c.cc.Invoke(ctx, Registry_GetImageConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) GetArchs(ctx context.Context, in *ArchRequest, opts ...grpc.CallOption) (*Archs, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Archs)
	err := c.cc.Invoke(ctx, Registry_GetArchs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) GetSize(ctx context.Context, in *SizeRequest, opts ...grpc.CallOption) (*Size, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Size)
	err := c.cc.Invoke(ctx, Registry_GetSize_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) GetRawSbom(ctx context.Context, in *RawSbomRequest, opts ...grpc.CallOption) (*RawSbom, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RawSbom)
	err := c.cc.Invoke(ctx, Registry_GetRawSbom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) GetVulnReport(ctx context.Context, in *VulnReportRequest, opts ...grpc.CallOption) (*v1.VulnReport, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.VulnReport)
	err := c.cc.Invoke(ctx, Registry_GetVulnReport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) ListManifestMetadata(ctx context.Context, in *ManifestMetadataFilter, opts ...grpc.CallOption) (*ManifestMetadataList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ManifestMetadataList)
	err := c.cc.Invoke(ctx, Registry_ListManifestMetadata_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) GetPackageVersionMetadata(ctx context.Context, in *PackageVersionMetadataRequest, opts ...grpc.CallOption) (*PackageVersionMetadata, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PackageVersionMetadata)
	err := c.cc.Invoke(ctx, Registry_GetPackageVersionMetadata_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistryServer is the server API for Registry service.
// All implementations must embed UnimplementedRegistryServer
// for forward compatibility.
type RegistryServer interface {
	CreateRepo(context.Context, *CreateRepoRequest) (*Repo, error)
	UpdateRepo(context.Context, *Repo) (*Repo, error)
	ListRepos(context.Context, *RepoFilter) (*RepoList, error)
	DeleteRepo(context.Context, *DeleteRepoRequest) (*emptypb.Empty, error)
	CreateTag(context.Context, *CreateTagRequest) (*Tag, error)
	UpdateTag(context.Context, *Tag) (*Tag, error)
	DeleteTag(context.Context, *DeleteTagRequest) (*emptypb.Empty, error)
	ListTags(context.Context, *TagFilter) (*TagList, error)
	ListTagHistory(context.Context, *TagHistoryFilter) (*TagHistoryList, error)
	GetSbom(context.Context, *SbomRequest) (*v1.Sbom2, error)
	GetImageConfig(context.Context, *ImageConfigRequest) (*ImageConfig, error)
	GetArchs(context.Context, *ArchRequest) (*Archs, error)
	GetSize(context.Context, *SizeRequest) (*Size, error)
	GetRawSbom(context.Context, *RawSbomRequest) (*RawSbom, error)
	GetVulnReport(context.Context, *VulnReportRequest) (*v1.VulnReport, error)
	ListManifestMetadata(context.Context, *ManifestMetadataFilter) (*ManifestMetadataList, error)
	GetPackageVersionMetadata(context.Context, *PackageVersionMetadataRequest) (*PackageVersionMetadata, error)
	mustEmbedUnimplementedRegistryServer()
}

// UnimplementedRegistryServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRegistryServer struct{}

func (UnimplementedRegistryServer) CreateRepo(context.Context, *CreateRepoRequest) (*Repo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRepo not implemented")
}
func (UnimplementedRegistryServer) UpdateRepo(context.Context, *Repo) (*Repo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRepo not implemented")
}
func (UnimplementedRegistryServer) ListRepos(context.Context, *RepoFilter) (*RepoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRepos not implemented")
}
func (UnimplementedRegistryServer) DeleteRepo(context.Context, *DeleteRepoRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRepo not implemented")
}
func (UnimplementedRegistryServer) CreateTag(context.Context, *CreateTagRequest) (*Tag, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTag not implemented")
}
func (UnimplementedRegistryServer) UpdateTag(context.Context, *Tag) (*Tag, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTag not implemented")
}
func (UnimplementedRegistryServer) DeleteTag(context.Context, *DeleteTagRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTag not implemented")
}
func (UnimplementedRegistryServer) ListTags(context.Context, *TagFilter) (*TagList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTags not implemented")
}
func (UnimplementedRegistryServer) ListTagHistory(context.Context, *TagHistoryFilter) (*TagHistoryList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTagHistory not implemented")
}
func (UnimplementedRegistryServer) GetSbom(context.Context, *SbomRequest) (*v1.Sbom2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSbom not implemented")
}
func (UnimplementedRegistryServer) GetImageConfig(context.Context, *ImageConfigRequest) (*ImageConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImageConfig not implemented")
}
func (UnimplementedRegistryServer) GetArchs(context.Context, *ArchRequest) (*Archs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArchs not implemented")
}
func (UnimplementedRegistryServer) GetSize(context.Context, *SizeRequest) (*Size, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSize not implemented")
}
func (UnimplementedRegistryServer) GetRawSbom(context.Context, *RawSbomRequest) (*RawSbom, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRawSbom not implemented")
}
func (UnimplementedRegistryServer) GetVulnReport(context.Context, *VulnReportRequest) (*v1.VulnReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVulnReport not implemented")
}
func (UnimplementedRegistryServer) ListManifestMetadata(context.Context, *ManifestMetadataFilter) (*ManifestMetadataList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListManifestMetadata not implemented")
}
func (UnimplementedRegistryServer) GetPackageVersionMetadata(context.Context, *PackageVersionMetadataRequest) (*PackageVersionMetadata, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPackageVersionMetadata not implemented")
}
func (UnimplementedRegistryServer) mustEmbedUnimplementedRegistryServer() {}
func (UnimplementedRegistryServer) testEmbeddedByValue()                  {}

// UnsafeRegistryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegistryServer will
// result in compilation errors.
type UnsafeRegistryServer interface {
	mustEmbedUnimplementedRegistryServer()
}

func RegisterRegistryServer(s grpc.ServiceRegistrar, srv RegistryServer) {
	// If the following call pancis, it indicates UnimplementedRegistryServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Registry_ServiceDesc, srv)
}

func _Registry_CreateRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).CreateRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_CreateRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).CreateRepo(ctx, req.(*CreateRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_UpdateRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Repo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).UpdateRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_UpdateRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).UpdateRepo(ctx, req.(*Repo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_ListRepos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepoFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ListRepos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_ListRepos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ListRepos(ctx, req.(*RepoFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_DeleteRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).DeleteRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_DeleteRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).DeleteRepo(ctx, req.(*DeleteRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_CreateTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).CreateTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_CreateTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).CreateTag(ctx, req.(*CreateTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_UpdateTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Tag)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).UpdateTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_UpdateTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).UpdateTag(ctx, req.(*Tag))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_DeleteTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).DeleteTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_DeleteTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).DeleteTag(ctx, req.(*DeleteTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_ListTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TagFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ListTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_ListTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ListTags(ctx, req.(*TagFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_ListTagHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TagHistoryFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ListTagHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_ListTagHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ListTagHistory(ctx, req.(*TagHistoryFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_GetSbom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SbomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).GetSbom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_GetSbom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).GetSbom(ctx, req.(*SbomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_GetImageConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).GetImageConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_GetImageConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).GetImageConfig(ctx, req.(*ImageConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_GetArchs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).GetArchs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_GetArchs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).GetArchs(ctx, req.(*ArchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_GetSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).GetSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_GetSize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).GetSize(ctx, req.(*SizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_GetRawSbom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawSbomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).GetRawSbom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_GetRawSbom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).GetRawSbom(ctx, req.(*RawSbomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_GetVulnReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VulnReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).GetVulnReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_GetVulnReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).GetVulnReport(ctx, req.(*VulnReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_ListManifestMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManifestMetadataFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ListManifestMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_ListManifestMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ListManifestMetadata(ctx, req.(*ManifestMetadataFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_GetPackageVersionMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PackageVersionMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).GetPackageVersionMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_GetPackageVersionMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).GetPackageVersionMetadata(ctx, req.(*PackageVersionMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Registry_ServiceDesc is the grpc.ServiceDesc for Registry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Registry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chainguard.platform.registry.Registry",
	HandlerType: (*RegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRepo",
			Handler:    _Registry_CreateRepo_Handler,
		},
		{
			MethodName: "UpdateRepo",
			Handler:    _Registry_UpdateRepo_Handler,
		},
		{
			MethodName: "ListRepos",
			Handler:    _Registry_ListRepos_Handler,
		},
		{
			MethodName: "DeleteRepo",
			Handler:    _Registry_DeleteRepo_Handler,
		},
		{
			MethodName: "CreateTag",
			Handler:    _Registry_CreateTag_Handler,
		},
		{
			MethodName: "UpdateTag",
			Handler:    _Registry_UpdateTag_Handler,
		},
		{
			MethodName: "DeleteTag",
			Handler:    _Registry_DeleteTag_Handler,
		},
		{
			MethodName: "ListTags",
			Handler:    _Registry_ListTags_Handler,
		},
		{
			MethodName: "ListTagHistory",
			Handler:    _Registry_ListTagHistory_Handler,
		},
		{
			MethodName: "GetSbom",
			Handler:    _Registry_GetSbom_Handler,
		},
		{
			MethodName: "GetImageConfig",
			Handler:    _Registry_GetImageConfig_Handler,
		},
		{
			MethodName: "GetArchs",
			Handler:    _Registry_GetArchs_Handler,
		},
		{
			MethodName: "GetSize",
			Handler:    _Registry_GetSize_Handler,
		},
		{
			MethodName: "GetRawSbom",
			Handler:    _Registry_GetRawSbom_Handler,
		},
		{
			MethodName: "GetVulnReport",
			Handler:    _Registry_GetVulnReport_Handler,
		},
		{
			MethodName: "ListManifestMetadata",
			Handler:    _Registry_ListManifestMetadata_Handler,
		},
		{
			MethodName: "GetPackageVersionMetadata",
			Handler:    _Registry_GetPackageVersionMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry.platform.proto",
}
