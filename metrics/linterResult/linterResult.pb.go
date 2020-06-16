// Copyright 2020 Google LLC. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: linterResult.proto

package linter

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Path    string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Line    int32  `protobuf:"varint,5,opt,name=line,proto3" json:"line,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_linterResult_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_linterResult_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_linterResult_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Message) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Message) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Message) GetLine() int32 {
	if x != nil {
		return x.Line
	}
	return 0
}

type Linter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LinterResults []*Message `protobuf:"bytes,1,rep,name=linterResults,proto3" json:"linterResults,omitempty"`
}

func (x *Linter) Reset() {
	*x = Linter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_linterResult_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Linter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Linter) ProtoMessage() {}

func (x *Linter) ProtoReflect() protoreflect.Message {
	mi := &file_linterResult_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Linter.ProtoReflect.Descriptor instead.
func (*Linter) Descriptor() ([]byte, []int) {
	return file_linterResult_proto_rawDescGZIP(), []int{1}
}

func (x *Linter) GetLinterResults() []*Message {
	if x != nil {
		return x.LinterResults
	}
	return nil
}

var File_linterResult_proto protoreflect.FileDescriptor

var file_linterResult_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x22, 0x5f, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x3f, 0x0a,
	0x06, 0x4c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x0d, 0x6c, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x6c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x0d, 0x6c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x42, 0x08,
	0x5a, 0x06, 0x6c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_linterResult_proto_rawDescOnce sync.Once
	file_linterResult_proto_rawDescData = file_linterResult_proto_rawDesc
)

func file_linterResult_proto_rawDescGZIP() []byte {
	file_linterResult_proto_rawDescOnce.Do(func() {
		file_linterResult_proto_rawDescData = protoimpl.X.CompressGZIP(file_linterResult_proto_rawDescData)
	})
	return file_linterResult_proto_rawDescData
}

var file_linterResult_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_linterResult_proto_goTypes = []interface{}{
	(*Message)(nil), // 0: linter.Message
	(*Linter)(nil),  // 1: linter.Linter
}
var file_linterResult_proto_depIdxs = []int32{
	0, // 0: linter.Linter.linterResults:type_name -> linter.Message
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_linterResult_proto_init() }
func file_linterResult_proto_init() {
	if File_linterResult_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_linterResult_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_linterResult_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Linter); i {
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
			RawDescriptor: file_linterResult_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_linterResult_proto_goTypes,
		DependencyIndexes: file_linterResult_proto_depIdxs,
		MessageInfos:      file_linterResult_proto_msgTypes,
	}.Build()
	File_linterResult_proto = out.File
	file_linterResult_proto_rawDesc = nil
	file_linterResult_proto_goTypes = nil
	file_linterResult_proto_depIdxs = nil
}
