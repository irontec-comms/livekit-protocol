// Copyright 2023 LiveKit, Inc.
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
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: rpc/agent.proto

package rpc

import (
	livekit "github.com/livekit/protocol/livekit"
	_ "github.com/livekit/psrpc/protoc-gen-psrpc/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type JobTerminateReason int32

const (
	JobTerminateReason_TERMINATION_REQUESTED JobTerminateReason = 0
	JobTerminateReason_AGENT_LEFT_ROOM       JobTerminateReason = 1
)

// Enum value maps for JobTerminateReason.
var (
	JobTerminateReason_name = map[int32]string{
		0: "TERMINATION_REQUESTED",
		1: "AGENT_LEFT_ROOM",
	}
	JobTerminateReason_value = map[string]int32{
		"TERMINATION_REQUESTED": 0,
		"AGENT_LEFT_ROOM":       1,
	}
)

func (x JobTerminateReason) Enum() *JobTerminateReason {
	p := new(JobTerminateReason)
	*p = x
	return p
}

func (x JobTerminateReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobTerminateReason) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_agent_proto_enumTypes[0].Descriptor()
}

func (JobTerminateReason) Type() protoreflect.EnumType {
	return &file_rpc_agent_proto_enumTypes[0]
}

func (x JobTerminateReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobTerminateReason.Descriptor instead.
func (JobTerminateReason) EnumDescriptor() ([]byte, []int) {
	return file_rpc_agent_proto_rawDescGZIP(), []int{0}
}

type CheckEnabledRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckEnabledRequest) Reset() {
	*x = CheckEnabledRequest{}
	mi := &file_rpc_agent_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckEnabledRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckEnabledRequest) ProtoMessage() {}

func (x *CheckEnabledRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_agent_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckEnabledRequest.ProtoReflect.Descriptor instead.
func (*CheckEnabledRequest) Descriptor() ([]byte, []int) {
	return file_rpc_agent_proto_rawDescGZIP(), []int{0}
}

type CheckEnabledResponse struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	RoomEnabled        bool                   `protobuf:"varint,1,opt,name=room_enabled,json=roomEnabled,proto3" json:"room_enabled,omitempty"`
	PublisherEnabled   bool                   `protobuf:"varint,2,opt,name=publisher_enabled,json=publisherEnabled,proto3" json:"publisher_enabled,omitempty"`
	ParticipantEnabled bool                   `protobuf:"varint,5,opt,name=participant_enabled,json=participantEnabled,proto3" json:"participant_enabled,omitempty"`
	// Deprecated: Marked as deprecated in rpc/agent.proto.
	Namespaces    []string `protobuf:"bytes,3,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	AgentNames    []string `protobuf:"bytes,4,rep,name=agent_names,json=agentNames,proto3" json:"agent_names,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckEnabledResponse) Reset() {
	*x = CheckEnabledResponse{}
	mi := &file_rpc_agent_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckEnabledResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckEnabledResponse) ProtoMessage() {}

func (x *CheckEnabledResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_agent_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckEnabledResponse.ProtoReflect.Descriptor instead.
func (*CheckEnabledResponse) Descriptor() ([]byte, []int) {
	return file_rpc_agent_proto_rawDescGZIP(), []int{1}
}

func (x *CheckEnabledResponse) GetRoomEnabled() bool {
	if x != nil {
		return x.RoomEnabled
	}
	return false
}

func (x *CheckEnabledResponse) GetPublisherEnabled() bool {
	if x != nil {
		return x.PublisherEnabled
	}
	return false
}

func (x *CheckEnabledResponse) GetParticipantEnabled() bool {
	if x != nil {
		return x.ParticipantEnabled
	}
	return false
}

// Deprecated: Marked as deprecated in rpc/agent.proto.
func (x *CheckEnabledResponse) GetNamespaces() []string {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

func (x *CheckEnabledResponse) GetAgentNames() []string {
	if x != nil {
		return x.AgentNames
	}
	return nil
}

type JobRequestResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	State         *livekit.JobState      `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobRequestResponse) Reset() {
	*x = JobRequestResponse{}
	mi := &file_rpc_agent_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobRequestResponse) ProtoMessage() {}

func (x *JobRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_agent_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobRequestResponse.ProtoReflect.Descriptor instead.
func (*JobRequestResponse) Descriptor() ([]byte, []int) {
	return file_rpc_agent_proto_rawDescGZIP(), []int{2}
}

func (x *JobRequestResponse) GetState() *livekit.JobState {
	if x != nil {
		return x.State
	}
	return nil
}

type JobTerminateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	JobId         string                 `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Reason        JobTerminateReason     `protobuf:"varint,2,opt,name=reason,proto3,enum=rpc.JobTerminateReason" json:"reason,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobTerminateRequest) Reset() {
	*x = JobTerminateRequest{}
	mi := &file_rpc_agent_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobTerminateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobTerminateRequest) ProtoMessage() {}

func (x *JobTerminateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_agent_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobTerminateRequest.ProtoReflect.Descriptor instead.
func (*JobTerminateRequest) Descriptor() ([]byte, []int) {
	return file_rpc_agent_proto_rawDescGZIP(), []int{3}
}

func (x *JobTerminateRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *JobTerminateRequest) GetReason() JobTerminateReason {
	if x != nil {
		return x.Reason
	}
	return JobTerminateReason_TERMINATION_REQUESTED
}

type JobTerminateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	State         *livekit.JobState      `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobTerminateResponse) Reset() {
	*x = JobTerminateResponse{}
	mi := &file_rpc_agent_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobTerminateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobTerminateResponse) ProtoMessage() {}

func (x *JobTerminateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_agent_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobTerminateResponse.ProtoReflect.Descriptor instead.
func (*JobTerminateResponse) Descriptor() ([]byte, []int) {
	return file_rpc_agent_proto_rawDescGZIP(), []int{4}
}

func (x *JobTerminateResponse) GetState() *livekit.JobState {
	if x != nil {
		return x.State
	}
	return nil
}

var File_rpc_agent_proto protoreflect.FileDescriptor

const file_rpc_agent_proto_rawDesc = "" +
	"\n" +
	"\x0frpc/agent.proto\x12\x03rpc\x1a\x1bgoogle/protobuf/empty.proto\x1a\roptions.proto\x1a\x13livekit_agent.proto\"\x15\n" +
	"\x13CheckEnabledRequest\"\xdc\x01\n" +
	"\x14CheckEnabledResponse\x12!\n" +
	"\froom_enabled\x18\x01 \x01(\bR\vroomEnabled\x12+\n" +
	"\x11publisher_enabled\x18\x02 \x01(\bR\x10publisherEnabled\x12/\n" +
	"\x13participant_enabled\x18\x05 \x01(\bR\x12participantEnabled\x12\"\n" +
	"\n" +
	"namespaces\x18\x03 \x03(\tB\x02\x18\x01R\n" +
	"namespaces\x12\x1f\n" +
	"\vagent_names\x18\x04 \x03(\tR\n" +
	"agentNames\"=\n" +
	"\x12JobRequestResponse\x12'\n" +
	"\x05state\x18\x01 \x01(\v2\x11.livekit.JobStateR\x05state\"]\n" +
	"\x13JobTerminateRequest\x12\x15\n" +
	"\x06job_id\x18\x01 \x01(\tR\x05jobId\x12/\n" +
	"\x06reason\x18\x02 \x01(\x0e2\x17.rpc.JobTerminateReasonR\x06reason\"?\n" +
	"\x14JobTerminateResponse\x12'\n" +
	"\x05state\x18\x01 \x01(\v2\x11.livekit.JobStateR\x05state*D\n" +
	"\x12JobTerminateReason\x12\x19\n" +
	"\x15TERMINATION_REQUESTED\x10\x00\x12\x13\n" +
	"\x0fAGENT_LEFT_ROOM\x10\x012\xee\x02\n" +
	"\rAgentInternal\x12K\n" +
	"\fCheckEnabled\x12\x18.rpc.CheckEnabledRequest\x1a\x19.rpc.CheckEnabledResponse\"\x06\xb2\x89\x01\x02(\x01\x12T\n" +
	"\n" +
	"JobRequest\x12\f.livekit.Job\x1a\x17.rpc.JobRequestResponse\"\x1f\xb2\x89\x01\x1b\x10\x01\x1a\x15\x12\tnamespace\x12\bjob_type0\x01\x12U\n" +
	"\fJobTerminate\x12\x18.rpc.JobTerminateRequest\x1a\x19.rpc.JobTerminateResponse\"\x10\xb2\x89\x01\f\x10\x01\x1a\b\x12\x06job_id\x12c\n" +
	"\x10WorkerRegistered\x12\x16.google.protobuf.Empty\x1a\x16.google.protobuf.Empty\"\x1f\xb2\x89\x01\x1b\b\x01\x10\x01\x1a\x13\x12\x11handler_namespace(\x01B!Z\x1fgithub.com/livekit/protocol/rpcb\x06proto3"

var (
	file_rpc_agent_proto_rawDescOnce sync.Once
	file_rpc_agent_proto_rawDescData []byte
)

func file_rpc_agent_proto_rawDescGZIP() []byte {
	file_rpc_agent_proto_rawDescOnce.Do(func() {
		file_rpc_agent_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_rpc_agent_proto_rawDesc), len(file_rpc_agent_proto_rawDesc)))
	})
	return file_rpc_agent_proto_rawDescData
}

var file_rpc_agent_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rpc_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_rpc_agent_proto_goTypes = []any{
	(JobTerminateReason)(0),      // 0: rpc.JobTerminateReason
	(*CheckEnabledRequest)(nil),  // 1: rpc.CheckEnabledRequest
	(*CheckEnabledResponse)(nil), // 2: rpc.CheckEnabledResponse
	(*JobRequestResponse)(nil),   // 3: rpc.JobRequestResponse
	(*JobTerminateRequest)(nil),  // 4: rpc.JobTerminateRequest
	(*JobTerminateResponse)(nil), // 5: rpc.JobTerminateResponse
	(*livekit.JobState)(nil),     // 6: livekit.JobState
	(*livekit.Job)(nil),          // 7: livekit.Job
	(*emptypb.Empty)(nil),        // 8: google.protobuf.Empty
}
var file_rpc_agent_proto_depIdxs = []int32{
	6, // 0: rpc.JobRequestResponse.state:type_name -> livekit.JobState
	0, // 1: rpc.JobTerminateRequest.reason:type_name -> rpc.JobTerminateReason
	6, // 2: rpc.JobTerminateResponse.state:type_name -> livekit.JobState
	1, // 3: rpc.AgentInternal.CheckEnabled:input_type -> rpc.CheckEnabledRequest
	7, // 4: rpc.AgentInternal.JobRequest:input_type -> livekit.Job
	4, // 5: rpc.AgentInternal.JobTerminate:input_type -> rpc.JobTerminateRequest
	8, // 6: rpc.AgentInternal.WorkerRegistered:input_type -> google.protobuf.Empty
	2, // 7: rpc.AgentInternal.CheckEnabled:output_type -> rpc.CheckEnabledResponse
	3, // 8: rpc.AgentInternal.JobRequest:output_type -> rpc.JobRequestResponse
	5, // 9: rpc.AgentInternal.JobTerminate:output_type -> rpc.JobTerminateResponse
	8, // 10: rpc.AgentInternal.WorkerRegistered:output_type -> google.protobuf.Empty
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_rpc_agent_proto_init() }
func file_rpc_agent_proto_init() {
	if File_rpc_agent_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_rpc_agent_proto_rawDesc), len(file_rpc_agent_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_agent_proto_goTypes,
		DependencyIndexes: file_rpc_agent_proto_depIdxs,
		EnumInfos:         file_rpc_agent_proto_enumTypes,
		MessageInfos:      file_rpc_agent_proto_msgTypes,
	}.Build()
	File_rpc_agent_proto = out.File
	file_rpc_agent_proto_goTypes = nil
	file_rpc_agent_proto_depIdxs = nil
}
