//
//Copyright The containerd Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: github.com/containerd/containerd/api/services/events/v1/events.proto

package events

import (
	_ "github.com/containerd/containerd/api/types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic string     `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Event *anypb.Any `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *PublishRequest) Reset() {
	*x = PublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_containerd_api_services_events_v1_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRequest) ProtoMessage() {}

func (x *PublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_containerd_api_services_events_v1_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRequest.ProtoReflect.Descriptor instead.
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_services_events_v1_events_proto_rawDescGZIP(), []int{0}
}

func (x *PublishRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *PublishRequest) GetEvent() *anypb.Any {
	if x != nil {
		return x.Event
	}
	return nil
}

type ForwardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Envelope *Envelope `protobuf:"bytes,1,opt,name=envelope,proto3" json:"envelope,omitempty"`
}

func (x *ForwardRequest) Reset() {
	*x = ForwardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_containerd_api_services_events_v1_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForwardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForwardRequest) ProtoMessage() {}

func (x *ForwardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_containerd_api_services_events_v1_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForwardRequest.ProtoReflect.Descriptor instead.
func (*ForwardRequest) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_services_events_v1_events_proto_rawDescGZIP(), []int{1}
}

func (x *ForwardRequest) GetEnvelope() *Envelope {
	if x != nil {
		return x.Envelope
	}
	return nil
}

type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filters []string `protobuf:"bytes,1,rep,name=filters,proto3" json:"filters,omitempty"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_containerd_api_services_events_v1_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_containerd_api_services_events_v1_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_services_events_v1_events_proto_rawDescGZIP(), []int{2}
}

func (x *SubscribeRequest) GetFilters() []string {
	if x != nil {
		return x.Filters
	}
	return nil
}

type Envelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Namespace string                 `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Topic     string                 `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	Event     *anypb.Any             `protobuf:"bytes,4,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *Envelope) Reset() {
	*x = Envelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_containerd_api_services_events_v1_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Envelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Envelope) ProtoMessage() {}

func (x *Envelope) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_containerd_api_services_events_v1_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Envelope.ProtoReflect.Descriptor instead.
func (*Envelope) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_services_events_v1_events_proto_rawDescGZIP(), []int{3}
}

func (x *Envelope) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Envelope) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Envelope) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Envelope) GetEvent() *anypb.Any {
	if x != nil {
		return x.Event
	}
	return nil
}

var File_github_com_containerd_containerd_api_services_events_v1_events_proto protoreflect.FileDescriptor

var file_github_com_containerd_containerd_api_services_events_v1_events_proto_rawDesc = []byte{
	0x0a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x70, 0x61, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x0e, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x12, 0x2a, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x55,
	0x0a, 0x0e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x43, 0x0a, 0x08, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x52, 0x08, 0x65, 0x6e, 0x76,
	0x65, 0x6c, 0x6f, 0x70, 0x65, 0x22, 0x2c, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x22, 0xaa, 0x01, 0x0a, 0x08, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65,
	0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x2a,
	0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x3a, 0x04, 0x80, 0xb9, 0x1f, 0x01,
	0x32, 0x95, 0x02, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x50, 0x0a, 0x07, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x2d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x50, 0x0a,
	0x07, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x12, 0x2d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x67, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x2f, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x30, 0x01, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2f, 0x76, 0x31, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_github_com_containerd_containerd_api_services_events_v1_events_proto_rawDescOnce sync.Once
	file_github_com_containerd_containerd_api_services_events_v1_events_proto_rawDescData = file_github_com_containerd_containerd_api_services_events_v1_events_proto_rawDesc
)

func file_github_com_containerd_containerd_api_services_events_v1_events_proto_rawDescGZIP() []byte {
	file_github_com_containerd_containerd_api_services_events_v1_events_proto_rawDescOnce.Do(func() {
		file_github_com_containerd_containerd_api_services_events_v1_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_containerd_containerd_api_services_events_v1_events_proto_rawDescData)
	})
	return file_github_com_containerd_containerd_api_services_events_v1_events_proto_rawDescData
}

var file_github_com_containerd_containerd_api_services_events_v1_events_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_containerd_containerd_api_services_events_v1_events_proto_goTypes = []interface{}{
	(*PublishRequest)(nil),        // 0: containerd.services.events.v1.PublishRequest
	(*ForwardRequest)(nil),        // 1: containerd.services.events.v1.ForwardRequest
	(*SubscribeRequest)(nil),      // 2: containerd.services.events.v1.SubscribeRequest
	(*Envelope)(nil),              // 3: containerd.services.events.v1.Envelope
	(*anypb.Any)(nil),             // 4: google.protobuf.Any
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 6: google.protobuf.Empty
}
var file_github_com_containerd_containerd_api_services_events_v1_events_proto_depIdxs = []int32{
	4, // 0: containerd.services.events.v1.PublishRequest.event:type_name -> google.protobuf.Any
	3, // 1: containerd.services.events.v1.ForwardRequest.envelope:type_name -> containerd.services.events.v1.Envelope
	5, // 2: containerd.services.events.v1.Envelope.timestamp:type_name -> google.protobuf.Timestamp
	4, // 3: containerd.services.events.v1.Envelope.event:type_name -> google.protobuf.Any
	0, // 4: containerd.services.events.v1.Events.Publish:input_type -> containerd.services.events.v1.PublishRequest
	1, // 5: containerd.services.events.v1.Events.Forward:input_type -> containerd.services.events.v1.ForwardRequest
	2, // 6: containerd.services.events.v1.Events.Subscribe:input_type -> containerd.services.events.v1.SubscribeRequest
	6, // 7: containerd.services.events.v1.Events.Publish:output_type -> google.protobuf.Empty
	6, // 8: containerd.services.events.v1.Events.Forward:output_type -> google.protobuf.Empty
	3, // 9: containerd.services.events.v1.Events.Subscribe:output_type -> containerd.services.events.v1.Envelope
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_github_com_containerd_containerd_api_services_events_v1_events_proto_init() }
func file_github_com_containerd_containerd_api_services_events_v1_events_proto_init() {
	if File_github_com_containerd_containerd_api_services_events_v1_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_containerd_containerd_api_services_events_v1_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishRequest); i {
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
		file_github_com_containerd_containerd_api_services_events_v1_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForwardRequest); i {
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
		file_github_com_containerd_containerd_api_services_events_v1_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
		file_github_com_containerd_containerd_api_services_events_v1_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Envelope); i {
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
			RawDescriptor: file_github_com_containerd_containerd_api_services_events_v1_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_containerd_containerd_api_services_events_v1_events_proto_goTypes,
		DependencyIndexes: file_github_com_containerd_containerd_api_services_events_v1_events_proto_depIdxs,
		MessageInfos:      file_github_com_containerd_containerd_api_services_events_v1_events_proto_msgTypes,
	}.Build()
	File_github_com_containerd_containerd_api_services_events_v1_events_proto = out.File
	file_github_com_containerd_containerd_api_services_events_v1_events_proto_rawDesc = nil
	file_github_com_containerd_containerd_api_services_events_v1_events_proto_goTypes = nil
	file_github_com_containerd_containerd_api_services_events_v1_events_proto_depIdxs = nil
}