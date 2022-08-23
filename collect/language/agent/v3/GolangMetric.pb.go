//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: language-agent/GolangMetric.proto

package v3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	v3 "skywalking.apache.org/repo/goapi/collect/common/v3"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GolangMetricCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics         []*GolangMetric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
	Service         string          `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	ServiceInstance string          `protobuf:"bytes,3,opt,name=serviceInstance,proto3" json:"serviceInstance,omitempty"`
}

func (x *GolangMetricCollection) Reset() {
	*x = GolangMetricCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_language_agent_GolangMetric_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GolangMetricCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GolangMetricCollection) ProtoMessage() {}

func (x *GolangMetricCollection) ProtoReflect() protoreflect.Message {
	mi := &file_language_agent_GolangMetric_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GolangMetricCollection.ProtoReflect.Descriptor instead.
func (*GolangMetricCollection) Descriptor() ([]byte, []int) {
	return file_language_agent_GolangMetric_proto_rawDescGZIP(), []int{0}
}

func (x *GolangMetricCollection) GetMetrics() []*GolangMetric {
	if x != nil {
		return x.Metrics
	}
	return nil
}

func (x *GolangMetricCollection) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *GolangMetricCollection) GetServiceInstance() string {
	if x != nil {
		return x.ServiceInstance
	}
	return ""
}

type GolangMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time         int64   `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	HeapAlloc    int64   `protobuf:"varint,2,opt,name=heapAlloc,proto3" json:"heapAlloc,omitempty"`
	StackInUse   int64   `protobuf:"varint,3,opt,name=stackInUse,proto3" json:"stackInUse,omitempty"`
	GcNum        int64   `protobuf:"varint,4,opt,name=gcNum,proto3" json:"gcNum,omitempty"`
	GcPauseTime  float32 `protobuf:"fixed32,5,opt,name=gcPauseTime,proto3" json:"gcPauseTime,omitempty"`
	GoroutineNum int64   `protobuf:"varint,6,opt,name=goroutineNum,proto3" json:"goroutineNum,omitempty"`
	ThreadNum    int64   `protobuf:"varint,7,opt,name=threadNum,proto3" json:"threadNum,omitempty"`
	CpuUsedRate  float32 `protobuf:"fixed32,8,opt,name=cpuUsedRate,proto3" json:"cpuUsedRate,omitempty"`
	MemUsedRate  float32 `protobuf:"fixed32,9,opt,name=memUsedRate,proto3" json:"memUsedRate,omitempty"`
}

func (x *GolangMetric) Reset() {
	*x = GolangMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_language_agent_GolangMetric_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GolangMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GolangMetric) ProtoMessage() {}

func (x *GolangMetric) ProtoReflect() protoreflect.Message {
	mi := &file_language_agent_GolangMetric_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GolangMetric.ProtoReflect.Descriptor instead.
func (*GolangMetric) Descriptor() ([]byte, []int) {
	return file_language_agent_GolangMetric_proto_rawDescGZIP(), []int{1}
}

func (x *GolangMetric) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *GolangMetric) GetHeapAlloc() int64 {
	if x != nil {
		return x.HeapAlloc
	}
	return 0
}

func (x *GolangMetric) GetStackInUse() int64 {
	if x != nil {
		return x.StackInUse
	}
	return 0
}

func (x *GolangMetric) GetGcNum() int64 {
	if x != nil {
		return x.GcNum
	}
	return 0
}

func (x *GolangMetric) GetGcPauseTime() float32 {
	if x != nil {
		return x.GcPauseTime
	}
	return 0
}

func (x *GolangMetric) GetGoroutineNum() int64 {
	if x != nil {
		return x.GoroutineNum
	}
	return 0
}

func (x *GolangMetric) GetThreadNum() int64 {
	if x != nil {
		return x.ThreadNum
	}
	return 0
}

func (x *GolangMetric) GetCpuUsedRate() float32 {
	if x != nil {
		return x.CpuUsedRate
	}
	return 0
}

func (x *GolangMetric) GetMemUsedRate() float32 {
	if x != nil {
		return x.MemUsedRate
	}
	return 0
}

var File_language_agent_GolangMetric_proto protoreflect.FileDescriptor

var file_language_agent_GolangMetric_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2f, 0x47, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x33, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x16, 0x47, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x33, 0x2e, 0x47, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x9e, 0x02,
	0x0a, 0x0c, 0x47, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x65, 0x61, 0x70, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x68, 0x65, 0x61, 0x70, 0x41, 0x6c, 0x6c, 0x6f, 0x63,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x55, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x67, 0x63, 0x4e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x67, 0x63, 0x4e, 0x75, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x67, 0x63, 0x50, 0x61, 0x75, 0x73,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x67, 0x63, 0x50,
	0x61, 0x75, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x67, 0x6f, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x67, 0x6f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x4e, 0x75, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x4e, 0x75, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x70,
	0x75, 0x55, 0x73, 0x65, 0x64, 0x52, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0b, 0x63, 0x70, 0x75, 0x55, 0x73, 0x65, 0x64, 0x52, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x6d, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x64, 0x52, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x64, 0x52, 0x61, 0x74, 0x65, 0x32, 0x68,
	0x0a, 0x19, 0x47, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x07, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x25, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x17, 0x2e,
	0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x22, 0x00, 0x42, 0x71, 0x0a, 0x33, 0x6f, 0x72, 0x67, 0x2e,
	0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x33, 0x5a,
	0x3a, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63,
	0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x67, 0x6f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_language_agent_GolangMetric_proto_rawDescOnce sync.Once
	file_language_agent_GolangMetric_proto_rawDescData = file_language_agent_GolangMetric_proto_rawDesc
)

func file_language_agent_GolangMetric_proto_rawDescGZIP() []byte {
	file_language_agent_GolangMetric_proto_rawDescOnce.Do(func() {
		file_language_agent_GolangMetric_proto_rawDescData = protoimpl.X.CompressGZIP(file_language_agent_GolangMetric_proto_rawDescData)
	})
	return file_language_agent_GolangMetric_proto_rawDescData
}

var file_language_agent_GolangMetric_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_language_agent_GolangMetric_proto_goTypes = []interface{}{
	(*GolangMetricCollection)(nil), // 0: skywalking.v3.GolangMetricCollection
	(*GolangMetric)(nil),           // 1: skywalking.v3.GolangMetric
	(*v3.Commands)(nil),            // 2: skywalking.v3.Commands
}
var file_language_agent_GolangMetric_proto_depIdxs = []int32{
	1, // 0: skywalking.v3.GolangMetricCollection.metrics:type_name -> skywalking.v3.GolangMetric
	0, // 1: skywalking.v3.GolangMetricReportService.collect:input_type -> skywalking.v3.GolangMetricCollection
	2, // 2: skywalking.v3.GolangMetricReportService.collect:output_type -> skywalking.v3.Commands
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_language_agent_GolangMetric_proto_init() }
func file_language_agent_GolangMetric_proto_init() {
	if File_language_agent_GolangMetric_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_language_agent_GolangMetric_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GolangMetricCollection); i {
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
		file_language_agent_GolangMetric_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GolangMetric); i {
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
			RawDescriptor: file_language_agent_GolangMetric_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_language_agent_GolangMetric_proto_goTypes,
		DependencyIndexes: file_language_agent_GolangMetric_proto_depIdxs,
		MessageInfos:      file_language_agent_GolangMetric_proto_msgTypes,
	}.Build()
	File_language_agent_GolangMetric_proto = out.File
	file_language_agent_GolangMetric_proto_rawDesc = nil
	file_language_agent_GolangMetric_proto_goTypes = nil
	file_language_agent_GolangMetric_proto_depIdxs = nil
}
