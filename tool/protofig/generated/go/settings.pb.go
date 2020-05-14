// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: settings.proto

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

type ProtoModuleConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModuleID string             `protobuf:"bytes,1,opt,name=moduleID,proto3" json:"moduleID,omitempty"`
	Configs  map[string]*Config `protobuf:"bytes,2,rep,name=configs,proto3" json:"configs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ProtoModuleConfig) Reset() {
	*x = ProtoModuleConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoModuleConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoModuleConfig) ProtoMessage() {}

func (x *ProtoModuleConfig) ProtoReflect() protoreflect.Message {
	mi := &file_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoModuleConfig.ProtoReflect.Descriptor instead.
func (*ProtoModuleConfig) Descriptor() ([]byte, []int) {
	return file_settings_proto_rawDescGZIP(), []int{0}
}

func (x *ProtoModuleConfig) GetModuleID() string {
	if x != nil {
		return x.ModuleID
	}
	return ""
}

func (x *ProtoModuleConfig) GetConfigs() map[string]*Config {
	if x != nil {
		return x.Configs
	}
	return nil
}

var File_settings_proto protoreflect.FileDescriptor

var file_settings_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x0f, 0x62, 0x61, 0x73, 0x65, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x01, 0x0a, 0x11, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x40, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x1a, 0x4a, 0x0a,
	0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x24, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_settings_proto_rawDescOnce sync.Once
	file_settings_proto_rawDescData = file_settings_proto_rawDesc
)

func file_settings_proto_rawDescGZIP() []byte {
	file_settings_proto_rawDescOnce.Do(func() {
		file_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_settings_proto_rawDescData)
	})
	return file_settings_proto_rawDescData
}

var file_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_settings_proto_goTypes = []interface{}{
	(*ProtoModuleConfig)(nil), // 0: config.ProtoModuleConfig
	nil,                       // 1: config.ProtoModuleConfig.ConfigsEntry
	(*Config)(nil),            // 2: config.Config
}
var file_settings_proto_depIdxs = []int32{
	1, // 0: config.ProtoModuleConfig.configs:type_name -> config.ProtoModuleConfig.ConfigsEntry
	2, // 1: config.ProtoModuleConfig.ConfigsEntry.value:type_name -> config.Config
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_settings_proto_init() }
func file_settings_proto_init() {
	if File_settings_proto != nil {
		return
	}
	file_baseproto_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoModuleConfig); i {
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
			RawDescriptor: file_settings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_settings_proto_goTypes,
		DependencyIndexes: file_settings_proto_depIdxs,
		MessageInfos:      file_settings_proto_msgTypes,
	}.Build()
	File_settings_proto = out.File
	file_settings_proto_rawDesc = nil
	file_settings_proto_goTypes = nil
	file_settings_proto_depIdxs = nil
}
