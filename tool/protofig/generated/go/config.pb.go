// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: config.proto

package config

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MinioComponent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinioAccesskey *ConfigVal `protobuf:"bytes,1,opt,name=minio_accesskey,json=minioAccesskey,proto3" json:"minio_accesskey,omitempty"` // string
	MinioSecretkey *ConfigVal `protobuf:"bytes,2,opt,name=minio_secretkey,json=minioSecretkey,proto3" json:"minio_secretkey,omitempty"` // string
}

func (x *MinioComponent) Reset() {
	*x = MinioComponent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MinioComponent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MinioComponent) ProtoMessage() {}

func (x *MinioComponent) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MinioComponent.ProtoReflect.Descriptor instead.
func (*MinioComponent) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

func (x *MinioComponent) GetMinioAccesskey() *ConfigVal {
	if x != nil {
		return x.MinioAccesskey
	}
	return nil
}

func (x *MinioComponent) GetMinioSecretkey() *ConfigVal {
	if x != nil {
		return x.MinioSecretkey
	}
	return nil
}

type MaintemplateComponent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinioAccesskey *ConfigVal `protobuf:"bytes,1,opt,name=minio_accesskey,json=minioAccesskey,proto3" json:"minio_accesskey,omitempty"` // string
	MinioSecretkey *ConfigVal `protobuf:"bytes,2,opt,name=minio_secretkey,json=minioSecretkey,proto3" json:"minio_secretkey,omitempty"` // string
	MinioLocation  *ConfigVal `protobuf:"bytes,3,opt,name=minio_location,json=minioLocation,proto3" json:"minio_location,omitempty"`    // string
	MinioTimeout   *ConfigVal `protobuf:"bytes,4,opt,name=minio_timeout,json=minioTimeout,proto3" json:"minio_timeout,omitempty"`       // double
	MinioSsl       *ConfigVal `protobuf:"bytes,5,opt,name=minio_ssl,json=minioSsl,proto3" json:"minio_ssl,omitempty"`                   // bool
	MinioEnckey    *ConfigVal `protobuf:"bytes,6,opt,name=minio_enckey,json=minioEnckey,proto3" json:"minio_enckey,omitempty"`          // bytes
	MinioEndpoint  *ConfigVal `protobuf:"bytes,7,opt,name=minio_endpoint,json=minioEndpoint,proto3" json:"minio_endpoint,omitempty"`    // cidr
}

func (x *MaintemplateComponent) Reset() {
	*x = MaintemplateComponent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaintemplateComponent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaintemplateComponent) ProtoMessage() {}

func (x *MaintemplateComponent) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaintemplateComponent.ProtoReflect.Descriptor instead.
func (*MaintemplateComponent) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{1}
}

func (x *MaintemplateComponent) GetMinioAccesskey() *ConfigVal {
	if x != nil {
		return x.MinioAccesskey
	}
	return nil
}

func (x *MaintemplateComponent) GetMinioSecretkey() *ConfigVal {
	if x != nil {
		return x.MinioSecretkey
	}
	return nil
}

func (x *MaintemplateComponent) GetMinioLocation() *ConfigVal {
	if x != nil {
		return x.MinioLocation
	}
	return nil
}

func (x *MaintemplateComponent) GetMinioTimeout() *ConfigVal {
	if x != nil {
		return x.MinioTimeout
	}
	return nil
}

func (x *MaintemplateComponent) GetMinioSsl() *ConfigVal {
	if x != nil {
		return x.MinioSsl
	}
	return nil
}

func (x *MaintemplateComponent) GetMinioEnckey() *ConfigVal {
	if x != nil {
		return x.MinioEnckey
	}
	return nil
}

func (x *MaintemplateComponent) GetMinioEndpoint() *ConfigVal {
	if x != nil {
		return x.MinioEndpoint
	}
	return nil
}

type GcpComponent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GcpUser    *ConfigVal `protobuf:"bytes,1,opt,name=gcp_user,json=gcpUser,proto3" json:"gcp_user,omitempty"`          // string
	GcpProject *ConfigVal `protobuf:"bytes,2,opt,name=gcp_project,json=gcpProject,proto3" json:"gcp_project,omitempty"` // string
	GkeCluster *ConfigVal `protobuf:"bytes,3,opt,name=gke_cluster,json=gkeCluster,proto3" json:"gke_cluster,omitempty"` // string
	GkeZone    *ConfigVal `protobuf:"bytes,4,opt,name=gke_zone,json=gkeZone,proto3" json:"gke_zone,omitempty"`          // string
	GkeEmail   *ConfigVal `protobuf:"bytes,5,opt,name=gke_email,json=gkeEmail,proto3" json:"gke_email,omitempty"`       // email
}

func (x *GcpComponent) Reset() {
	*x = GcpComponent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcpComponent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcpComponent) ProtoMessage() {}

func (x *GcpComponent) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcpComponent.ProtoReflect.Descriptor instead.
func (*GcpComponent) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{2}
}

func (x *GcpComponent) GetGcpUser() *ConfigVal {
	if x != nil {
		return x.GcpUser
	}
	return nil
}

func (x *GcpComponent) GetGcpProject() *ConfigVal {
	if x != nil {
		return x.GcpProject
	}
	return nil
}

func (x *GcpComponent) GetGkeCluster() *ConfigVal {
	if x != nil {
		return x.GkeCluster
	}
	return nil
}

func (x *GcpComponent) GetGkeZone() *ConfigVal {
	if x != nil {
		return x.GkeZone
	}
	return nil
}

func (x *GcpComponent) GetGkeEmail() *ConfigVal {
	if x != nil {
		return x.GkeEmail
	}
	return nil
}

type JwtComponent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrivateKey *ConfigVal `protobuf:"bytes,1,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"` // string
	PublicKey  *ConfigVal `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`    // string
}

func (x *JwtComponent) Reset() {
	*x = JwtComponent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JwtComponent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JwtComponent) ProtoMessage() {}

func (x *JwtComponent) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JwtComponent.ProtoReflect.Descriptor instead.
func (*JwtComponent) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{3}
}

func (x *JwtComponent) GetPrivateKey() *ConfigVal {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

func (x *JwtComponent) GetPublicKey() *ConfigVal {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

type WorkflowComponent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GithubSha        *ConfigVal `protobuf:"bytes,2,opt,name=github_sha,json=githubSha,proto3" json:"github_sha,omitempty"`                      // string
	GithubRef        *ConfigVal `protobuf:"bytes,3,opt,name=github_ref,json=githubRef,proto3" json:"github_ref,omitempty"`                      // string
	Project          *ConfigVal `protobuf:"bytes,4,opt,name=project,proto3" json:"project,omitempty"`                                           // string
	RegistryHostname *ConfigVal `protobuf:"bytes,5,opt,name=registry_hostname,json=registryHostname,proto3" json:"registry_hostname,omitempty"` // cidr
	Url              *ConfigVal `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`                                                   // string
	Locales          *ConfigVal `protobuf:"bytes,7,opt,name=locales,proto3" json:"locales,omitempty"`                                           // repeated string
	FlutterChannel   *ConfigVal `protobuf:"bytes,8,opt,name=flutter_channel,json=flutterChannel,proto3" json:"flutter_channel,omitempty"`       // string
	ReleaseChannel   *ConfigVal `protobuf:"bytes,9,opt,name=release_channel,json=releaseChannel,proto3" json:"release_channel,omitempty"`       // string
}

func (x *WorkflowComponent) Reset() {
	*x = WorkflowComponent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowComponent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowComponent) ProtoMessage() {}

func (x *WorkflowComponent) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowComponent.ProtoReflect.Descriptor instead.
func (*WorkflowComponent) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{4}
}

func (x *WorkflowComponent) GetGithubSha() *ConfigVal {
	if x != nil {
		return x.GithubSha
	}
	return nil
}

func (x *WorkflowComponent) GetGithubRef() *ConfigVal {
	if x != nil {
		return x.GithubRef
	}
	return nil
}

func (x *WorkflowComponent) GetProject() *ConfigVal {
	if x != nil {
		return x.Project
	}
	return nil
}

func (x *WorkflowComponent) GetRegistryHostname() *ConfigVal {
	if x != nil {
		return x.RegistryHostname
	}
	return nil
}

func (x *WorkflowComponent) GetUrl() *ConfigVal {
	if x != nil {
		return x.Url
	}
	return nil
}

func (x *WorkflowComponent) GetLocales() *ConfigVal {
	if x != nil {
		return x.Locales
	}
	return nil
}

func (x *WorkflowComponent) GetFlutterChannel() *ConfigVal {
	if x != nil {
		return x.FlutterChannel
	}
	return nil
}

func (x *WorkflowComponent) GetReleaseChannel() *ConfigVal {
	if x != nil {
		return x.ReleaseChannel
	}
	return nil
}

var File_config_proto protoreflect.FileDescriptor

var file_config_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x0f, 0x62, 0x61, 0x73, 0x65, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x0e, 0x4d, 0x69, 0x6e, 0x69,
	0x6f, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x0f, 0x6d, 0x69,
	0x6e, 0x69, 0x6f, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x56, 0x61, 0x6c, 0x52, 0x0e, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x6b, 0x65, 0x79, 0x12, 0x3a, 0x0a, 0x0f, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x5f,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56,
	0x61, 0x6c, 0x52, 0x0e, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x6b,
	0x65, 0x79, 0x22, 0xa1, 0x03, 0x0a, 0x15, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x0f,
	0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x61, 0x6c, 0x52, 0x0e, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x6b, 0x65, 0x79, 0x12, 0x3a, 0x0a, 0x0f, 0x6d, 0x69, 0x6e, 0x69,
	0x6f, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x56, 0x61, 0x6c, 0x52, 0x0e, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x6b, 0x65, 0x79, 0x12, 0x38, 0x0a, 0x0e, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x5f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x61, 0x6c, 0x52,
	0x0d, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36,
	0x0a, 0x0d, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x61, 0x6c, 0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x5f,
	0x73, 0x73, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x61, 0x6c, 0x52, 0x08, 0x6d, 0x69,
	0x6e, 0x69, 0x6f, 0x53, 0x73, 0x6c, 0x12, 0x34, 0x0a, 0x0c, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x5f,
	0x65, 0x6e, 0x63, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x61, 0x6c, 0x52,
	0x0b, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x45, 0x6e, 0x63, 0x6b, 0x65, 0x79, 0x12, 0x38, 0x0a, 0x0e,
	0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x56, 0x61, 0x6c, 0x52, 0x0d, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x82, 0x02, 0x0a, 0x0c, 0x47, 0x63, 0x70, 0x43, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x67, 0x63, 0x70, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x61, 0x6c, 0x52, 0x07, 0x67, 0x63,
	0x70, 0x55, 0x73, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x0b, 0x67, 0x63, 0x70, 0x5f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x61, 0x6c, 0x52, 0x0a, 0x67,
	0x63, 0x70, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x32, 0x0a, 0x0b, 0x67, 0x6b, 0x65,
	0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x61,
	0x6c, 0x52, 0x0a, 0x67, 0x6b, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x2c, 0x0a,
	0x08, 0x67, 0x6b, 0x65, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56,
	0x61, 0x6c, 0x52, 0x07, 0x67, 0x6b, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x67,
	0x6b, 0x65, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x61,
	0x6c, 0x52, 0x08, 0x67, 0x6b, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x74, 0x0a, 0x0c, 0x4a,
	0x77, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x0b, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x56, 0x61, 0x6c, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12,
	0x30, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x56, 0x61, 0x6c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x22, 0xae, 0x03, 0x0a, 0x11, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x43, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x0a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x5f, 0x73, 0x68, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x61, 0x6c, 0x52, 0x09,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x53, 0x68, 0x61, 0x12, 0x30, 0x0a, 0x0a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x61, 0x6c,
	0x52, 0x09, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x52, 0x65, 0x66, 0x12, 0x2b, 0x0a, 0x07, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x61, 0x6c, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x3e, 0x0a, 0x11, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x56, 0x61, 0x6c, 0x52, 0x10, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x61, 0x6c, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x2b, 0x0a,
	0x07, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x61,
	0x6c, 0x52, 0x07, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x0f, 0x66, 0x6c,
	0x75, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x56, 0x61, 0x6c, 0x52, 0x0e, 0x66, 0x6c, 0x75, 0x74, 0x74, 0x65, 0x72, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x3a, 0x0a, 0x0f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56,
	0x61, 0x6c, 0x52, 0x0e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_proto_rawDescOnce sync.Once
	file_config_proto_rawDescData = file_config_proto_rawDesc
)

func file_config_proto_rawDescGZIP() []byte {
	file_config_proto_rawDescOnce.Do(func() {
		file_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_proto_rawDescData)
	})
	return file_config_proto_rawDescData
}

var file_config_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_config_proto_goTypes = []interface{}{
	(*MinioComponent)(nil),        // 0: config.MinioComponent
	(*MaintemplateComponent)(nil), // 1: config.MaintemplateComponent
	(*GcpComponent)(nil),          // 2: config.GcpComponent
	(*JwtComponent)(nil),          // 3: config.JwtComponent
	(*WorkflowComponent)(nil),     // 4: config.WorkflowComponent
	(*ConfigVal)(nil),             // 5: config.ConfigVal
}
var file_config_proto_depIdxs = []int32{
	5,  // 0: config.MinioComponent.minio_accesskey:type_name -> config.ConfigVal
	5,  // 1: config.MinioComponent.minio_secretkey:type_name -> config.ConfigVal
	5,  // 2: config.MaintemplateComponent.minio_accesskey:type_name -> config.ConfigVal
	5,  // 3: config.MaintemplateComponent.minio_secretkey:type_name -> config.ConfigVal
	5,  // 4: config.MaintemplateComponent.minio_location:type_name -> config.ConfigVal
	5,  // 5: config.MaintemplateComponent.minio_timeout:type_name -> config.ConfigVal
	5,  // 6: config.MaintemplateComponent.minio_ssl:type_name -> config.ConfigVal
	5,  // 7: config.MaintemplateComponent.minio_enckey:type_name -> config.ConfigVal
	5,  // 8: config.MaintemplateComponent.minio_endpoint:type_name -> config.ConfigVal
	5,  // 9: config.GcpComponent.gcp_user:type_name -> config.ConfigVal
	5,  // 10: config.GcpComponent.gcp_project:type_name -> config.ConfigVal
	5,  // 11: config.GcpComponent.gke_cluster:type_name -> config.ConfigVal
	5,  // 12: config.GcpComponent.gke_zone:type_name -> config.ConfigVal
	5,  // 13: config.GcpComponent.gke_email:type_name -> config.ConfigVal
	5,  // 14: config.JwtComponent.private_key:type_name -> config.ConfigVal
	5,  // 15: config.JwtComponent.public_key:type_name -> config.ConfigVal
	5,  // 16: config.WorkflowComponent.github_sha:type_name -> config.ConfigVal
	5,  // 17: config.WorkflowComponent.github_ref:type_name -> config.ConfigVal
	5,  // 18: config.WorkflowComponent.project:type_name -> config.ConfigVal
	5,  // 19: config.WorkflowComponent.registry_hostname:type_name -> config.ConfigVal
	5,  // 20: config.WorkflowComponent.url:type_name -> config.ConfigVal
	5,  // 21: config.WorkflowComponent.locales:type_name -> config.ConfigVal
	5,  // 22: config.WorkflowComponent.flutter_channel:type_name -> config.ConfigVal
	5,  // 23: config.WorkflowComponent.release_channel:type_name -> config.ConfigVal
	24, // [24:24] is the sub-list for method output_type
	24, // [24:24] is the sub-list for method input_type
	24, // [24:24] is the sub-list for extension type_name
	24, // [24:24] is the sub-list for extension extendee
	0,  // [0:24] is the sub-list for field type_name
}

func init() { file_config_proto_init() }
func file_config_proto_init() {
	if File_config_proto != nil {
		return
	}
	file_baseproto_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MinioComponent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaintemplateComponent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcpComponent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JwtComponent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowComponent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_proto_goTypes,
		DependencyIndexes: file_config_proto_depIdxs,
		MessageInfos:      file_config_proto_msgTypes,
	}.Build()
	File_config_proto = out.File
	file_config_proto_rawDesc = nil
	file_config_proto_goTypes = nil
	file_config_proto_depIdxs = nil
}
