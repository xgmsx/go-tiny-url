// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.2
// source: shortener_v1.proto

package shortener_v1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateLinkRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateLinkRequest) Reset() {
	*x = CreateLinkRequest{}
	mi := &file_shortener_v1_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLinkRequest) ProtoMessage() {}

func (x *CreateLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_v1_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLinkRequest.ProtoReflect.Descriptor instead.
func (*CreateLinkRequest) Descriptor() ([]byte, []int) {
	return file_shortener_v1_proto_rawDescGZIP(), []int{0}
}

func (x *CreateLinkRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type CreateLinkResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Alias         string                 `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
	ExpiredAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateLinkResponse) Reset() {
	*x = CreateLinkResponse{}
	mi := &file_shortener_v1_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLinkResponse) ProtoMessage() {}

func (x *CreateLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_v1_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLinkResponse.ProtoReflect.Descriptor instead.
func (*CreateLinkResponse) Descriptor() ([]byte, []int) {
	return file_shortener_v1_proto_rawDescGZIP(), []int{1}
}

func (x *CreateLinkResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CreateLinkResponse) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *CreateLinkResponse) GetExpiredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiredAt
	}
	return nil
}

type FetchLinkRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Alias         string                 `protobuf:"bytes,1,opt,name=alias,proto3" json:"alias,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FetchLinkRequest) Reset() {
	*x = FetchLinkRequest{}
	mi := &file_shortener_v1_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchLinkRequest) ProtoMessage() {}

func (x *FetchLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_v1_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchLinkRequest.ProtoReflect.Descriptor instead.
func (*FetchLinkRequest) Descriptor() ([]byte, []int) {
	return file_shortener_v1_proto_rawDescGZIP(), []int{2}
}

func (x *FetchLinkRequest) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

type FetchLinkResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Alias         string                 `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
	ExpiredAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FetchLinkResponse) Reset() {
	*x = FetchLinkResponse{}
	mi := &file_shortener_v1_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchLinkResponse) ProtoMessage() {}

func (x *FetchLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_v1_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchLinkResponse.ProtoReflect.Descriptor instead.
func (*FetchLinkResponse) Descriptor() ([]byte, []int) {
	return file_shortener_v1_proto_rawDescGZIP(), []int{3}
}

func (x *FetchLinkResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *FetchLinkResponse) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *FetchLinkResponse) GetExpiredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiredAt
	}
	return nil
}

var File_shortener_v1_proto protoreflect.FileDescriptor

var file_shortener_v1_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x77, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x28, 0x0a, 0x10, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x6e, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x22, 0x76, 0x0a,
	0x11, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x64, 0x41, 0x74, 0x32, 0xaa, 0x01, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x12, 0x4f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e,
	0x6b, 0x12, 0x1f, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x09, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x6e,
	0x6b, 0x12, 0x1e, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f, 0x76, 0x31,
	0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f, 0x76, 0x31,
	0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shortener_v1_proto_rawDescOnce sync.Once
	file_shortener_v1_proto_rawDescData = file_shortener_v1_proto_rawDesc
)

func file_shortener_v1_proto_rawDescGZIP() []byte {
	file_shortener_v1_proto_rawDescOnce.Do(func() {
		file_shortener_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_shortener_v1_proto_rawDescData)
	})
	return file_shortener_v1_proto_rawDescData
}

var file_shortener_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_shortener_v1_proto_goTypes = []any{
	(*CreateLinkRequest)(nil),     // 0: shortener_v1.CreateLinkRequest
	(*CreateLinkResponse)(nil),    // 1: shortener_v1.CreateLinkResponse
	(*FetchLinkRequest)(nil),      // 2: shortener_v1.FetchLinkRequest
	(*FetchLinkResponse)(nil),     // 3: shortener_v1.FetchLinkResponse
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_shortener_v1_proto_depIdxs = []int32{
	4, // 0: shortener_v1.CreateLinkResponse.expired_at:type_name -> google.protobuf.Timestamp
	4, // 1: shortener_v1.FetchLinkResponse.expired_at:type_name -> google.protobuf.Timestamp
	0, // 2: shortener_v1.Shortener.CreateLink:input_type -> shortener_v1.CreateLinkRequest
	2, // 3: shortener_v1.Shortener.FetchLink:input_type -> shortener_v1.FetchLinkRequest
	1, // 4: shortener_v1.Shortener.CreateLink:output_type -> shortener_v1.CreateLinkResponse
	3, // 5: shortener_v1.Shortener.FetchLink:output_type -> shortener_v1.FetchLinkResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_shortener_v1_proto_init() }
func file_shortener_v1_proto_init() {
	if File_shortener_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shortener_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shortener_v1_proto_goTypes,
		DependencyIndexes: file_shortener_v1_proto_depIdxs,
		MessageInfos:      file_shortener_v1_proto_msgTypes,
	}.Build()
	File_shortener_v1_proto = out.File
	file_shortener_v1_proto_rawDesc = nil
	file_shortener_v1_proto_goTypes = nil
	file_shortener_v1_proto_depIdxs = nil
}
