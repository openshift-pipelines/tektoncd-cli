// Copyright 2018, OpenCensus Authors
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: opencensus/proto/agent/trace/v1/trace_service.proto

// NOTE: This proto is experimental and is subject to change at this point.
// Please do not use it at the moment.

package v1

import (
	context "context"
	v1 "github.com/census-instrumentation/opencensus-proto/gen-go/agent/common/v1"
	v12 "github.com/census-instrumentation/opencensus-proto/gen-go/resource/v1"
	v11 "github.com/census-instrumentation/opencensus-proto/gen-go/trace/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CurrentLibraryConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This is required only in the first message on the stream or if the
	// previous sent CurrentLibraryConfig message has a different Node (e.g.
	// when the same RPC is used to configure multiple Applications).
	Node *v1.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// Current configuration.
	Config *v11.TraceConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *CurrentLibraryConfig) Reset() {
	*x = CurrentLibraryConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opencensus_proto_agent_trace_v1_trace_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentLibraryConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentLibraryConfig) ProtoMessage() {}

func (x *CurrentLibraryConfig) ProtoReflect() protoreflect.Message {
	mi := &file_opencensus_proto_agent_trace_v1_trace_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentLibraryConfig.ProtoReflect.Descriptor instead.
func (*CurrentLibraryConfig) Descriptor() ([]byte, []int) {
	return file_opencensus_proto_agent_trace_v1_trace_service_proto_rawDescGZIP(), []int{0}
}

func (x *CurrentLibraryConfig) GetNode() *v1.Node {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *CurrentLibraryConfig) GetConfig() *v11.TraceConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type UpdatedLibraryConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This field is ignored when the RPC is used to configure only one Application.
	// This is required only in the first message on the stream or if the
	// previous sent UpdatedLibraryConfig message has a different Node.
	Node *v1.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// Requested updated configuration.
	Config *v11.TraceConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *UpdatedLibraryConfig) Reset() {
	*x = UpdatedLibraryConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opencensus_proto_agent_trace_v1_trace_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatedLibraryConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatedLibraryConfig) ProtoMessage() {}

func (x *UpdatedLibraryConfig) ProtoReflect() protoreflect.Message {
	mi := &file_opencensus_proto_agent_trace_v1_trace_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatedLibraryConfig.ProtoReflect.Descriptor instead.
func (*UpdatedLibraryConfig) Descriptor() ([]byte, []int) {
	return file_opencensus_proto_agent_trace_v1_trace_service_proto_rawDescGZIP(), []int{1}
}

func (x *UpdatedLibraryConfig) GetNode() *v1.Node {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *UpdatedLibraryConfig) GetConfig() *v11.TraceConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type ExportTraceServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This is required only in the first message on the stream or if the
	// previous sent ExportTraceServiceRequest message has a different Node (e.g.
	// when the same RPC is used to send Spans from multiple Applications).
	Node *v1.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// A list of Spans that belong to the last received Node.
	Spans []*v11.Span `protobuf:"bytes,2,rep,name=spans,proto3" json:"spans,omitempty"`
	// The resource for the spans in this message that do not have an explicit
	// resource set.
	// If unset, the most recently set resource in the RPC stream applies. It is
	// valid to never be set within a stream, e.g. when no resource info is known.
	Resource *v12.Resource `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *ExportTraceServiceRequest) Reset() {
	*x = ExportTraceServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opencensus_proto_agent_trace_v1_trace_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportTraceServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportTraceServiceRequest) ProtoMessage() {}

func (x *ExportTraceServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_opencensus_proto_agent_trace_v1_trace_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportTraceServiceRequest.ProtoReflect.Descriptor instead.
func (*ExportTraceServiceRequest) Descriptor() ([]byte, []int) {
	return file_opencensus_proto_agent_trace_v1_trace_service_proto_rawDescGZIP(), []int{2}
}

func (x *ExportTraceServiceRequest) GetNode() *v1.Node {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *ExportTraceServiceRequest) GetSpans() []*v11.Span {
	if x != nil {
		return x.Spans
	}
	return nil
}

func (x *ExportTraceServiceRequest) GetResource() *v12.Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

type ExportTraceServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ExportTraceServiceResponse) Reset() {
	*x = ExportTraceServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opencensus_proto_agent_trace_v1_trace_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportTraceServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportTraceServiceResponse) ProtoMessage() {}

func (x *ExportTraceServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_opencensus_proto_agent_trace_v1_trace_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportTraceServiceResponse.ProtoReflect.Descriptor instead.
func (*ExportTraceServiceResponse) Descriptor() ([]byte, []int) {
	return file_opencensus_proto_agent_trace_v1_trace_service_proto_rawDescGZIP(), []int{3}
}

var File_opencensus_proto_agent_trace_v1_trace_service_proto protoreflect.FileDescriptor

var file_opencensus_proto_agent_trace_v1_trace_service_proto_rawDesc = []byte{
	0x0a, 0x33, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x73, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x73,
	0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x73, 0x75,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x25, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x6f, 0x70, 0x65, 0x6e, 0x63,
	0x65, 0x6e, 0x73, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a, 0x14, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x3a, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x3e, 0x0a, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x92, 0x01, 0x0a,
	0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3a, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64,
	0x65, 0x12, 0x3e, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x22, 0xd2, 0x01, 0x0a, 0x19, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x3a, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x35, 0x0a, 0x05, 0x73,
	0x70, 0x61, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x52, 0x05, 0x73, 0x70, 0x61,
	0x6e, 0x73, 0x12, 0x42, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x73, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x1c, 0x0a, 0x1a, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x54, 0x72, 0x61, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0x96, 0x02, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x63, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7c, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x35, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x35, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x65, 0x6e,
	0x73, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00, 0x28,
	0x01, 0x30, 0x01, 0x12, 0x87, 0x01, 0x0a, 0x06, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x3a,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x72, 0x61, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x54, 0x72, 0x61, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0xa9, 0x01,
	0x0a, 0x22, 0x69, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x54, 0x72, 0x61, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2d, 0x69, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x65,
	0x6e, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65,
	0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0xea, 0x02, 0x23, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x65, 0x6e, 0x73, 0x75, 0x73,
	0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x3a, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x3a, 0x3a,
	0x54, 0x72, 0x61, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_opencensus_proto_agent_trace_v1_trace_service_proto_rawDescOnce sync.Once
	file_opencensus_proto_agent_trace_v1_trace_service_proto_rawDescData = file_opencensus_proto_agent_trace_v1_trace_service_proto_rawDesc
)

func file_opencensus_proto_agent_trace_v1_trace_service_proto_rawDescGZIP() []byte {
	file_opencensus_proto_agent_trace_v1_trace_service_proto_rawDescOnce.Do(func() {
		file_opencensus_proto_agent_trace_v1_trace_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_opencensus_proto_agent_trace_v1_trace_service_proto_rawDescData)
	})
	return file_opencensus_proto_agent_trace_v1_trace_service_proto_rawDescData
}

var file_opencensus_proto_agent_trace_v1_trace_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_opencensus_proto_agent_trace_v1_trace_service_proto_goTypes = []interface{}{
	(*CurrentLibraryConfig)(nil),       // 0: opencensus.proto.agent.trace.v1.CurrentLibraryConfig
	(*UpdatedLibraryConfig)(nil),       // 1: opencensus.proto.agent.trace.v1.UpdatedLibraryConfig
	(*ExportTraceServiceRequest)(nil),  // 2: opencensus.proto.agent.trace.v1.ExportTraceServiceRequest
	(*ExportTraceServiceResponse)(nil), // 3: opencensus.proto.agent.trace.v1.ExportTraceServiceResponse
	(*v1.Node)(nil),                    // 4: opencensus.proto.agent.common.v1.Node
	(*v11.TraceConfig)(nil),            // 5: opencensus.proto.trace.v1.TraceConfig
	(*v11.Span)(nil),                   // 6: opencensus.proto.trace.v1.Span
	(*v12.Resource)(nil),               // 7: opencensus.proto.resource.v1.Resource
}
var file_opencensus_proto_agent_trace_v1_trace_service_proto_depIdxs = []int32{
	4, // 0: opencensus.proto.agent.trace.v1.CurrentLibraryConfig.node:type_name -> opencensus.proto.agent.common.v1.Node
	5, // 1: opencensus.proto.agent.trace.v1.CurrentLibraryConfig.config:type_name -> opencensus.proto.trace.v1.TraceConfig
	4, // 2: opencensus.proto.agent.trace.v1.UpdatedLibraryConfig.node:type_name -> opencensus.proto.agent.common.v1.Node
	5, // 3: opencensus.proto.agent.trace.v1.UpdatedLibraryConfig.config:type_name -> opencensus.proto.trace.v1.TraceConfig
	4, // 4: opencensus.proto.agent.trace.v1.ExportTraceServiceRequest.node:type_name -> opencensus.proto.agent.common.v1.Node
	6, // 5: opencensus.proto.agent.trace.v1.ExportTraceServiceRequest.spans:type_name -> opencensus.proto.trace.v1.Span
	7, // 6: opencensus.proto.agent.trace.v1.ExportTraceServiceRequest.resource:type_name -> opencensus.proto.resource.v1.Resource
	0, // 7: opencensus.proto.agent.trace.v1.TraceService.Config:input_type -> opencensus.proto.agent.trace.v1.CurrentLibraryConfig
	2, // 8: opencensus.proto.agent.trace.v1.TraceService.Export:input_type -> opencensus.proto.agent.trace.v1.ExportTraceServiceRequest
	1, // 9: opencensus.proto.agent.trace.v1.TraceService.Config:output_type -> opencensus.proto.agent.trace.v1.UpdatedLibraryConfig
	3, // 10: opencensus.proto.agent.trace.v1.TraceService.Export:output_type -> opencensus.proto.agent.trace.v1.ExportTraceServiceResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_opencensus_proto_agent_trace_v1_trace_service_proto_init() }
func file_opencensus_proto_agent_trace_v1_trace_service_proto_init() {
	if File_opencensus_proto_agent_trace_v1_trace_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_opencensus_proto_agent_trace_v1_trace_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrentLibraryConfig); i {
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
		file_opencensus_proto_agent_trace_v1_trace_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatedLibraryConfig); i {
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
		file_opencensus_proto_agent_trace_v1_trace_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportTraceServiceRequest); i {
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
		file_opencensus_proto_agent_trace_v1_trace_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportTraceServiceResponse); i {
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
			RawDescriptor: file_opencensus_proto_agent_trace_v1_trace_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_opencensus_proto_agent_trace_v1_trace_service_proto_goTypes,
		DependencyIndexes: file_opencensus_proto_agent_trace_v1_trace_service_proto_depIdxs,
		MessageInfos:      file_opencensus_proto_agent_trace_v1_trace_service_proto_msgTypes,
	}.Build()
	File_opencensus_proto_agent_trace_v1_trace_service_proto = out.File
	file_opencensus_proto_agent_trace_v1_trace_service_proto_rawDesc = nil
	file_opencensus_proto_agent_trace_v1_trace_service_proto_goTypes = nil
	file_opencensus_proto_agent_trace_v1_trace_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TraceServiceClient is the client API for TraceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TraceServiceClient interface {
	// After initialization, this RPC must be kept alive for the entire life of
	// the application. The agent pushes configs down to applications via a
	// stream.
	Config(ctx context.Context, opts ...grpc.CallOption) (TraceService_ConfigClient, error)
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(ctx context.Context, opts ...grpc.CallOption) (TraceService_ExportClient, error)
}

type traceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTraceServiceClient(cc grpc.ClientConnInterface) TraceServiceClient {
	return &traceServiceClient{cc}
}

func (c *traceServiceClient) Config(ctx context.Context, opts ...grpc.CallOption) (TraceService_ConfigClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TraceService_serviceDesc.Streams[0], "/opencensus.proto.agent.trace.v1.TraceService/Config", opts...)
	if err != nil {
		return nil, err
	}
	x := &traceServiceConfigClient{stream}
	return x, nil
}

type TraceService_ConfigClient interface {
	Send(*CurrentLibraryConfig) error
	Recv() (*UpdatedLibraryConfig, error)
	grpc.ClientStream
}

type traceServiceConfigClient struct {
	grpc.ClientStream
}

func (x *traceServiceConfigClient) Send(m *CurrentLibraryConfig) error {
	return x.ClientStream.SendMsg(m)
}

func (x *traceServiceConfigClient) Recv() (*UpdatedLibraryConfig, error) {
	m := new(UpdatedLibraryConfig)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *traceServiceClient) Export(ctx context.Context, opts ...grpc.CallOption) (TraceService_ExportClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TraceService_serviceDesc.Streams[1], "/opencensus.proto.agent.trace.v1.TraceService/Export", opts...)
	if err != nil {
		return nil, err
	}
	x := &traceServiceExportClient{stream}
	return x, nil
}

type TraceService_ExportClient interface {
	Send(*ExportTraceServiceRequest) error
	Recv() (*ExportTraceServiceResponse, error)
	grpc.ClientStream
}

type traceServiceExportClient struct {
	grpc.ClientStream
}

func (x *traceServiceExportClient) Send(m *ExportTraceServiceRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *traceServiceExportClient) Recv() (*ExportTraceServiceResponse, error) {
	m := new(ExportTraceServiceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TraceServiceServer is the server API for TraceService service.
type TraceServiceServer interface {
	// After initialization, this RPC must be kept alive for the entire life of
	// the application. The agent pushes configs down to applications via a
	// stream.
	Config(TraceService_ConfigServer) error
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(TraceService_ExportServer) error
}

// UnimplementedTraceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTraceServiceServer struct {
}

func (*UnimplementedTraceServiceServer) Config(TraceService_ConfigServer) error {
	return status.Errorf(codes.Unimplemented, "method Config not implemented")
}
func (*UnimplementedTraceServiceServer) Export(TraceService_ExportServer) error {
	return status.Errorf(codes.Unimplemented, "method Export not implemented")
}

func RegisterTraceServiceServer(s *grpc.Server, srv TraceServiceServer) {
	s.RegisterService(&_TraceService_serviceDesc, srv)
}

func _TraceService_Config_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TraceServiceServer).Config(&traceServiceConfigServer{stream})
}

type TraceService_ConfigServer interface {
	Send(*UpdatedLibraryConfig) error
	Recv() (*CurrentLibraryConfig, error)
	grpc.ServerStream
}

type traceServiceConfigServer struct {
	grpc.ServerStream
}

func (x *traceServiceConfigServer) Send(m *UpdatedLibraryConfig) error {
	return x.ServerStream.SendMsg(m)
}

func (x *traceServiceConfigServer) Recv() (*CurrentLibraryConfig, error) {
	m := new(CurrentLibraryConfig)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TraceService_Export_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TraceServiceServer).Export(&traceServiceExportServer{stream})
}

type TraceService_ExportServer interface {
	Send(*ExportTraceServiceResponse) error
	Recv() (*ExportTraceServiceRequest, error)
	grpc.ServerStream
}

type traceServiceExportServer struct {
	grpc.ServerStream
}

func (x *traceServiceExportServer) Send(m *ExportTraceServiceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *traceServiceExportServer) Recv() (*ExportTraceServiceRequest, error) {
	m := new(ExportTraceServiceRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TraceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "opencensus.proto.agent.trace.v1.TraceService",
	HandlerType: (*TraceServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Config",
			Handler:       _TraceService_Config_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Export",
			Handler:       _TraceService_Export_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "opencensus/proto/agent/trace/v1/trace_service.proto",
}
