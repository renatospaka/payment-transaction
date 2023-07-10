// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.6.1
// source: reprocess.proto

package pb

import (
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

type AuthorizationReprocessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorizationId string  `protobuf:"bytes,1,opt,name=authorization_id,json=authorizationId,proto3" json:"authorization_id,omitempty"`
	ClientId        string  `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	TransactionId   string  `protobuf:"bytes,3,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	Value           float32 `protobuf:"fixed32,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *AuthorizationReprocessRequest) Reset() {
	*x = AuthorizationReprocessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reprocess_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizationReprocessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizationReprocessRequest) ProtoMessage() {}

func (x *AuthorizationReprocessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reprocess_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizationReprocessRequest.ProtoReflect.Descriptor instead.
func (*AuthorizationReprocessRequest) Descriptor() ([]byte, []int) {
	return file_reprocess_proto_rawDescGZIP(), []int{0}
}

func (x *AuthorizationReprocessRequest) GetAuthorizationId() string {
	if x != nil {
		return x.AuthorizationId
	}
	return ""
}

func (x *AuthorizationReprocessRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *AuthorizationReprocessRequest) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *AuthorizationReprocessRequest) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type AuthorizationReprocessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorizationId string  `protobuf:"bytes,1,opt,name=authorization_id,json=authorizationId,proto3" json:"authorization_id,omitempty"`
	ClientId        string  `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	TransactionId   string  `protobuf:"bytes,3,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	Status          string  `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Value           float32 `protobuf:"fixed32,5,opt,name=value,proto3" json:"value,omitempty"`
	ErrorMessage    string  `protobuf:"bytes,6,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *AuthorizationReprocessResponse) Reset() {
	*x = AuthorizationReprocessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reprocess_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizationReprocessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizationReprocessResponse) ProtoMessage() {}

func (x *AuthorizationReprocessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reprocess_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizationReprocessResponse.ProtoReflect.Descriptor instead.
func (*AuthorizationReprocessResponse) Descriptor() ([]byte, []int) {
	return file_reprocess_proto_rawDescGZIP(), []int{1}
}

func (x *AuthorizationReprocessResponse) GetAuthorizationId() string {
	if x != nil {
		return x.AuthorizationId
	}
	return ""
}

func (x *AuthorizationReprocessResponse) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *AuthorizationReprocessResponse) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *AuthorizationReprocessResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *AuthorizationReprocessResponse) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *AuthorizationReprocessResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_reprocess_proto protoreflect.FileDescriptor

var file_reprocess_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x65, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0xa4, 0x01, 0x0a, 0x1d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xe2, 0x01, 0x0a, 0x1e, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reprocess_proto_rawDescOnce sync.Once
	file_reprocess_proto_rawDescData = file_reprocess_proto_rawDesc
)

func file_reprocess_proto_rawDescGZIP() []byte {
	file_reprocess_proto_rawDescOnce.Do(func() {
		file_reprocess_proto_rawDescData = protoimpl.X.CompressGZIP(file_reprocess_proto_rawDescData)
	})
	return file_reprocess_proto_rawDescData
}

var file_reprocess_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_reprocess_proto_goTypes = []interface{}{
	(*AuthorizationReprocessRequest)(nil),  // 0: authorization.AuthorizationReprocessRequest
	(*AuthorizationReprocessResponse)(nil), // 1: authorization.AuthorizationReprocessResponse
}
var file_reprocess_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_reprocess_proto_init() }
func file_reprocess_proto_init() {
	if File_reprocess_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reprocess_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizationReprocessRequest); i {
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
		file_reprocess_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizationReprocessResponse); i {
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
			RawDescriptor: file_reprocess_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_reprocess_proto_goTypes,
		DependencyIndexes: file_reprocess_proto_depIdxs,
		MessageInfos:      file_reprocess_proto_msgTypes,
	}.Build()
	File_reprocess_proto = out.File
	file_reprocess_proto_rawDesc = nil
	file_reprocess_proto_goTypes = nil
	file_reprocess_proto_depIdxs = nil
}
