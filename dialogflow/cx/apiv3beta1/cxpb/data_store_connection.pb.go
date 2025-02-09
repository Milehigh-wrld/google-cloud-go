// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: google/cloud/dialogflow/cx/v3beta1/data_store_connection.proto

package cxpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Type of a data store.
// Determines how search is performed in the data store.
type DataStoreType int32

const (
	// Not specified. This value indicates that the data store type is not
	// specified, so it will not be used during search.
	DataStoreType_DATA_STORE_TYPE_UNSPECIFIED DataStoreType = 0
	// A data store that contains public web content.
	DataStoreType_PUBLIC_WEB DataStoreType = 1
	// A data store that contains unstructured private data.
	DataStoreType_UNSTRUCTURED DataStoreType = 2
	// A data store that contains structured data (for example FAQ).
	DataStoreType_STRUCTURED DataStoreType = 3
)

// Enum value maps for DataStoreType.
var (
	DataStoreType_name = map[int32]string{
		0: "DATA_STORE_TYPE_UNSPECIFIED",
		1: "PUBLIC_WEB",
		2: "UNSTRUCTURED",
		3: "STRUCTURED",
	}
	DataStoreType_value = map[string]int32{
		"DATA_STORE_TYPE_UNSPECIFIED": 0,
		"PUBLIC_WEB":                  1,
		"UNSTRUCTURED":                2,
		"STRUCTURED":                  3,
	}
)

func (x DataStoreType) Enum() *DataStoreType {
	p := new(DataStoreType)
	*p = x
	return p
}

func (x DataStoreType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataStoreType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto_enumTypes[0].Descriptor()
}

func (DataStoreType) Type() protoreflect.EnumType {
	return &file_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto_enumTypes[0]
}

func (x DataStoreType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataStoreType.Descriptor instead.
func (DataStoreType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto_rawDescGZIP(), []int{0}
}

// A data store connection. It represents a data store in Discovery Engine and
// the type of the contents it contains.
type DataStoreConnection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of the connected data store.
	DataStoreType DataStoreType `protobuf:"varint,1,opt,name=data_store_type,json=dataStoreType,proto3,enum=google.cloud.dialogflow.cx.v3beta1.DataStoreType" json:"data_store_type,omitempty"`
	// The full name of the referenced data store.
	// Formats:
	// `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}`
	// `projects/{project}/locations/{location}/dataStores/{data_store}`
	DataStore string `protobuf:"bytes,2,opt,name=data_store,json=dataStore,proto3" json:"data_store,omitempty"`
}

func (x *DataStoreConnection) Reset() {
	*x = DataStoreConnection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataStoreConnection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataStoreConnection) ProtoMessage() {}

func (x *DataStoreConnection) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataStoreConnection.ProtoReflect.Descriptor instead.
func (*DataStoreConnection) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto_rawDescGZIP(), []int{0}
}

func (x *DataStoreConnection) GetDataStoreType() DataStoreType {
	if x != nil {
		return x.DataStoreType
	}
	return DataStoreType_DATA_STORE_TYPE_UNSPECIFIED
}

func (x *DataStoreConnection) GetDataStore() string {
	if x != nil {
		return x.DataStore
	}
	return ""
}

var File_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto protoreflect.FileDescriptor

var file_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x78, 0x2f, 0x76, 0x33, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x22, 0x8f, 0x01, 0x0a, 0x13, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x59, 0x0a, 0x0f,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x63, 0x78, 0x2e, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74,
	0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2a, 0x62, 0x0a, 0x0d, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x44, 0x41, 0x54, 0x41, 0x5f,
	0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x55, 0x42, 0x4c,
	0x49, 0x43, 0x5f, 0x57, 0x45, 0x42, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x4e, 0x53, 0x54,
	0x52, 0x55, 0x43, 0x54, 0x55, 0x52, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x54,
	0x52, 0x55, 0x43, 0x54, 0x55, 0x52, 0x45, 0x44, 0x10, 0x03, 0x42, 0xd2, 0x01, 0x0a, 0x26, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x18, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x36, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f,
	0x77, 0x2f, 0x63, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x63, 0x78, 0x70, 0x62, 0x3b, 0x63, 0x78, 0x70, 0x62, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x02, 0x44,
	0x46, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x78, 0x2e, 0x56,
	0x33, 0x42, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c,
	0x6f, 0x77, 0x3a, 0x3a, 0x43, 0x58, 0x3a, 0x3a, 0x56, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto_rawDescOnce sync.Once
	file_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto_rawDescData = file_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto_rawDesc
)

func file_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto_rawDescGZIP() []byte {
	file_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto_rawDescOnce.Do(func() {
		file_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto_rawDescData)
	})
	return file_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto_rawDescData
}

var file_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto_goTypes = []interface{}{
	(DataStoreType)(0),          // 0: google.cloud.dialogflow.cx.v3beta1.DataStoreType
	(*DataStoreConnection)(nil), // 1: google.cloud.dialogflow.cx.v3beta1.DataStoreConnection
}
var file_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto_depIdxs = []int32{
	0, // 0: google.cloud.dialogflow.cx.v3beta1.DataStoreConnection.data_store_type:type_name -> google.cloud.dialogflow.cx.v3beta1.DataStoreType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto_init() }
func file_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto_init() {
	if File_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataStoreConnection); i {
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
			RawDescriptor: file_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto_goTypes,
		DependencyIndexes: file_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto_depIdxs,
		EnumInfos:         file_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto_enumTypes,
		MessageInfos:      file_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto_msgTypes,
	}.Build()
	File_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto = out.File
	file_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto_rawDesc = nil
	file_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto_goTypes = nil
	file_google_cloud_dialogflow_cx_v3beta1_data_store_connection_proto_depIdxs = nil
}
