// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: vod/business/vod_workflow.proto

package business

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

type VodStartWorkflowResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RunId string `protobuf:"bytes,1,opt,name=RunId,proto3" json:"RunId,omitempty"` // 工作流执行Id
}

func (x *VodStartWorkflowResult) Reset() {
	*x = VodStartWorkflowResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vod_business_vod_workflow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VodStartWorkflowResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VodStartWorkflowResult) ProtoMessage() {}

func (x *VodStartWorkflowResult) ProtoReflect() protoreflect.Message {
	mi := &file_vod_business_vod_workflow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VodStartWorkflowResult.ProtoReflect.Descriptor instead.
func (*VodStartWorkflowResult) Descriptor() ([]byte, []int) {
	return file_vod_business_vod_workflow_proto_rawDescGZIP(), []int{0}
}

func (x *VodStartWorkflowResult) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

type WorkflowParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OverrideParams *OverrideParams `protobuf:"bytes,1,opt,name=OverrideParams,proto3" json:"OverrideParams,omitempty"`                                                                                // 覆盖参数
	Condition      map[string]bool `protobuf:"bytes,2,rep,name=Condition,proto3" json:"Condition,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // 条件变量
}

func (x *WorkflowParams) Reset() {
	*x = WorkflowParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vod_business_vod_workflow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowParams) ProtoMessage() {}

func (x *WorkflowParams) ProtoReflect() protoreflect.Message {
	mi := &file_vod_business_vod_workflow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowParams.ProtoReflect.Descriptor instead.
func (*WorkflowParams) Descriptor() ([]byte, []int) {
	return file_vod_business_vod_workflow_proto_rawDescGZIP(), []int{1}
}

func (x *WorkflowParams) GetOverrideParams() *OverrideParams {
	if x != nil {
		return x.OverrideParams
	}
	return nil
}

func (x *WorkflowParams) GetCondition() map[string]bool {
	if x != nil {
		return x.Condition
	}
	return nil
}

type OverrideParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Logo           []*LogoOverride           `protobuf:"bytes,1,rep,name=Logo,proto3" json:"Logo,omitempty"`                     // 水印覆盖参数
	TranscodeVideo []*TranscodeVideoOverride `protobuf:"bytes,2,rep,name=TranscodeVideo,proto3" json:"TranscodeVideo,omitempty"` // 视频转码覆盖参数
	TranscodeAudio []*TranscodeAudioOverride `protobuf:"bytes,3,rep,name=TranscodeAudio,proto3" json:"TranscodeAudio,omitempty"` // 音频转码覆盖参数
	Snapshot       []*SnapshotOverride       `protobuf:"bytes,4,rep,name=Snapshot,proto3" json:"Snapshot,omitempty"`             // 截图覆盖参数
}

func (x *OverrideParams) Reset() {
	*x = OverrideParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vod_business_vod_workflow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OverrideParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverrideParams) ProtoMessage() {}

func (x *OverrideParams) ProtoReflect() protoreflect.Message {
	mi := &file_vod_business_vod_workflow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverrideParams.ProtoReflect.Descriptor instead.
func (*OverrideParams) Descriptor() ([]byte, []int) {
	return file_vod_business_vod_workflow_proto_rawDescGZIP(), []int{2}
}

func (x *OverrideParams) GetLogo() []*LogoOverride {
	if x != nil {
		return x.Logo
	}
	return nil
}

func (x *OverrideParams) GetTranscodeVideo() []*TranscodeVideoOverride {
	if x != nil {
		return x.TranscodeVideo
	}
	return nil
}

func (x *OverrideParams) GetTranscodeAudio() []*TranscodeAudioOverride {
	if x != nil {
		return x.TranscodeAudio
	}
	return nil
}

func (x *OverrideParams) GetSnapshot() []*SnapshotOverride {
	if x != nil {
		return x.Snapshot
	}
	return nil
}

type LogoOverride struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId []string          `protobuf:"bytes,1,rep,name=TemplateId,proto3" json:"TemplateId,omitempty"`                                                                             // 被覆盖的水印模板Id, 支持ALL
	Vars       map[string]string `protobuf:"bytes,2,rep,name=Vars,proto3" json:"Vars,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 自定义水印变量
}

func (x *LogoOverride) Reset() {
	*x = LogoOverride{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vod_business_vod_workflow_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoOverride) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoOverride) ProtoMessage() {}

func (x *LogoOverride) ProtoReflect() protoreflect.Message {
	mi := &file_vod_business_vod_workflow_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoOverride.ProtoReflect.Descriptor instead.
func (*LogoOverride) Descriptor() ([]byte, []int) {
	return file_vod_business_vod_workflow_proto_rawDescGZIP(), []int{3}
}

func (x *LogoOverride) GetTemplateId() []string {
	if x != nil {
		return x.TemplateId
	}
	return nil
}

func (x *LogoOverride) GetVars() map[string]string {
	if x != nil {
		return x.Vars
	}
	return nil
}

type TranscodeVideoOverride struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId []string `protobuf:"bytes,1,rep,name=TemplateId,proto3" json:"TemplateId,omitempty"` // 被覆盖的视频模板Id, 支持ALL
	Clip       *Clip    `protobuf:"bytes,2,opt,name=Clip,proto3" json:"Clip,omitempty"`             // 裁剪参数
}

func (x *TranscodeVideoOverride) Reset() {
	*x = TranscodeVideoOverride{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vod_business_vod_workflow_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranscodeVideoOverride) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranscodeVideoOverride) ProtoMessage() {}

func (x *TranscodeVideoOverride) ProtoReflect() protoreflect.Message {
	mi := &file_vod_business_vod_workflow_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranscodeVideoOverride.ProtoReflect.Descriptor instead.
func (*TranscodeVideoOverride) Descriptor() ([]byte, []int) {
	return file_vod_business_vod_workflow_proto_rawDescGZIP(), []int{4}
}

func (x *TranscodeVideoOverride) GetTemplateId() []string {
	if x != nil {
		return x.TemplateId
	}
	return nil
}

func (x *TranscodeVideoOverride) GetClip() *Clip {
	if x != nil {
		return x.Clip
	}
	return nil
}

type Clip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime int32 `protobuf:"varint,1,opt,name=StartTime,proto3" json:"StartTime,omitempty"` // 开始时间 ms
	EndTime   int32 `protobuf:"varint,2,opt,name=EndTime,proto3" json:"EndTime,omitempty"`     // 结束时间 ms
}

func (x *Clip) Reset() {
	*x = Clip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vod_business_vod_workflow_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Clip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Clip) ProtoMessage() {}

func (x *Clip) ProtoReflect() protoreflect.Message {
	mi := &file_vod_business_vod_workflow_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Clip.ProtoReflect.Descriptor instead.
func (*Clip) Descriptor() ([]byte, []int) {
	return file_vod_business_vod_workflow_proto_rawDescGZIP(), []int{5}
}

func (x *Clip) GetStartTime() int32 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *Clip) GetEndTime() int32 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

type TranscodeAudioOverride struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId []string `protobuf:"bytes,1,rep,name=TemplateId,proto3" json:"TemplateId,omitempty"` // 被覆盖的音频模板Id, 支持ALL
	Clip       *Clip    `protobuf:"bytes,2,opt,name=Clip,proto3" json:"Clip,omitempty"`
}

func (x *TranscodeAudioOverride) Reset() {
	*x = TranscodeAudioOverride{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vod_business_vod_workflow_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranscodeAudioOverride) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranscodeAudioOverride) ProtoMessage() {}

func (x *TranscodeAudioOverride) ProtoReflect() protoreflect.Message {
	mi := &file_vod_business_vod_workflow_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranscodeAudioOverride.ProtoReflect.Descriptor instead.
func (*TranscodeAudioOverride) Descriptor() ([]byte, []int) {
	return file_vod_business_vod_workflow_proto_rawDescGZIP(), []int{6}
}

func (x *TranscodeAudioOverride) GetTemplateId() []string {
	if x != nil {
		return x.TemplateId
	}
	return nil
}

func (x *TranscodeAudioOverride) GetClip() *Clip {
	if x != nil {
		return x.Clip
	}
	return nil
}

type SnapshotOverride struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId     []string `protobuf:"bytes,1,rep,name=TemplateId,proto3" json:"TemplateId,omitempty"`                 // 被覆盖的截图模板Id, 支持ALL
	OffsetTime     int32    `protobuf:"varint,2,opt,name=OffsetTime,proto3" json:"OffsetTime,omitempty"`                // 截图时间, 单位ms, AIDynpost和Sprite类型不支持
	OffsetTimeList []int32  `protobuf:"varint,3,rep,packed,name=OffsetTimeList,proto3" json:"OffsetTimeList,omitempty"` // 多Dynpost类型截取时间，单位ms
}

func (x *SnapshotOverride) Reset() {
	*x = SnapshotOverride{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vod_business_vod_workflow_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotOverride) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotOverride) ProtoMessage() {}

func (x *SnapshotOverride) ProtoReflect() protoreflect.Message {
	mi := &file_vod_business_vod_workflow_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotOverride.ProtoReflect.Descriptor instead.
func (*SnapshotOverride) Descriptor() ([]byte, []int) {
	return file_vod_business_vod_workflow_proto_rawDescGZIP(), []int{7}
}

func (x *SnapshotOverride) GetTemplateId() []string {
	if x != nil {
		return x.TemplateId
	}
	return nil
}

func (x *SnapshotOverride) GetOffsetTime() int32 {
	if x != nil {
		return x.OffsetTime
	}
	return 0
}

func (x *SnapshotOverride) GetOffsetTimeList() []int32 {
	if x != nil {
		return x.OffsetTimeList
	}
	return nil
}

var File_vod_business_vod_workflow_proto protoreflect.FileDescriptor

var file_vod_business_vod_workflow_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x76, 0x6f, 0x64, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x76,
	0x6f, 0x64, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1e, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x56, 0x6f, 0x64, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x22, 0x2e, 0x0a, 0x16, 0x56, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x52,
	0x75, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x52, 0x75, 0x6e, 0x49,
	0x64, 0x22, 0x83, 0x02, 0x0a, 0x0e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x12, 0x56, 0x0a, 0x0e, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x56,
	0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x56, 0x6f, 0x64, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x4f, 0x76,
	0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x0e, 0x4f, 0x76,
	0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x5b, 0x0a, 0x09,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x3d, 0x2e, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x56, 0x6f, 0x64, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x3c, 0x0a, 0x0e, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xe0, 0x02, 0x0a, 0x0e, 0x4f, 0x76, 0x65, 0x72,
	0x72, 0x69, 0x64, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x40, 0x0a, 0x04, 0x4c, 0x6f,
	0x67, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x56, 0x6f, 0x6c, 0x63, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x56, 0x6f, 0x64,
	0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x4f, 0x76,
	0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x52, 0x04, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x5e, 0x0a, 0x0e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x56, 0x6f, 0x64, 0x2e, 0x42, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x52, 0x0e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x5e, 0x0a, 0x0e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x56, 0x6f, 0x64, 0x2e, 0x42, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x41,
	0x75, 0x64, 0x69, 0x6f, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x52, 0x0e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x4c, 0x0a, 0x08,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30,
	0x2e, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x56, 0x6f, 0x64, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65,
	0x52, 0x08, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x22, 0xb3, 0x01, 0x0a, 0x0c, 0x4c,
	0x6f, 0x67, 0x6f, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0a, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x4a, 0x0a, 0x04, 0x56,
	0x61, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x56, 0x6f, 0x6c, 0x63,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x56, 0x6f,
	0x64, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x4f,
	0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x56, 0x61, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x04, 0x56, 0x61, 0x72, 0x73, 0x1a, 0x37, 0x0a, 0x09, 0x56, 0x61, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x72, 0x0a, 0x16, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x04, 0x43, 0x6c,
	0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x56, 0x6f, 0x6c, 0x63, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x56, 0x6f, 0x64,
	0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x43, 0x6c, 0x69, 0x70, 0x52, 0x04,
	0x43, 0x6c, 0x69, 0x70, 0x22, 0x3e, 0x0a, 0x04, 0x43, 0x6c, 0x69, 0x70, 0x12, 0x1c, 0x0a, 0x09,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x45, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x72, 0x0a, 0x16, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64,
	0x65, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0a, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x38,
	0x0a, 0x04, 0x43, 0x6c, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x56,
	0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x56, 0x6f, 0x64, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x43, 0x6c,
	0x69, 0x70, 0x52, 0x04, 0x43, 0x6c, 0x69, 0x70, 0x22, 0x7a, 0x0a, 0x10, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0a, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x0e, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0xb6, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x79, 0x74,
	0x65, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x76, 0x6f, 0x64, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x42, 0x0b, 0x56, 0x6f,
	0x64, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2f, 0x76, 0x6f, 0x6c, 0x63, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x76, 0x6f, 0x64, 0x2f, 0x62, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0xa0, 0x01, 0x01, 0xd8, 0x01, 0x01, 0xca, 0x02, 0x1e, 0x56,
	0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x5c, 0x56, 0x6f, 0x64, 0x5c, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0xe2, 0x02, 0x1d,
	0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vod_business_vod_workflow_proto_rawDescOnce sync.Once
	file_vod_business_vod_workflow_proto_rawDescData = file_vod_business_vod_workflow_proto_rawDesc
)

func file_vod_business_vod_workflow_proto_rawDescGZIP() []byte {
	file_vod_business_vod_workflow_proto_rawDescOnce.Do(func() {
		file_vod_business_vod_workflow_proto_rawDescData = protoimpl.X.CompressGZIP(file_vod_business_vod_workflow_proto_rawDescData)
	})
	return file_vod_business_vod_workflow_proto_rawDescData
}

var file_vod_business_vod_workflow_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_vod_business_vod_workflow_proto_goTypes = []interface{}{
	(*VodStartWorkflowResult)(nil), // 0: Volcengine.Models.Vod.Business.VodStartWorkflowResult
	(*WorkflowParams)(nil),         // 1: Volcengine.Models.Vod.Business.WorkflowParams
	(*OverrideParams)(nil),         // 2: Volcengine.Models.Vod.Business.OverrideParams
	(*LogoOverride)(nil),           // 3: Volcengine.Models.Vod.Business.LogoOverride
	(*TranscodeVideoOverride)(nil), // 4: Volcengine.Models.Vod.Business.TranscodeVideoOverride
	(*Clip)(nil),                   // 5: Volcengine.Models.Vod.Business.Clip
	(*TranscodeAudioOverride)(nil), // 6: Volcengine.Models.Vod.Business.TranscodeAudioOverride
	(*SnapshotOverride)(nil),       // 7: Volcengine.Models.Vod.Business.SnapshotOverride
	nil,                            // 8: Volcengine.Models.Vod.Business.WorkflowParams.ConditionEntry
	nil,                            // 9: Volcengine.Models.Vod.Business.LogoOverride.VarsEntry
}
var file_vod_business_vod_workflow_proto_depIdxs = []int32{
	2, // 0: Volcengine.Models.Vod.Business.WorkflowParams.OverrideParams:type_name -> Volcengine.Models.Vod.Business.OverrideParams
	8, // 1: Volcengine.Models.Vod.Business.WorkflowParams.Condition:type_name -> Volcengine.Models.Vod.Business.WorkflowParams.ConditionEntry
	3, // 2: Volcengine.Models.Vod.Business.OverrideParams.Logo:type_name -> Volcengine.Models.Vod.Business.LogoOverride
	4, // 3: Volcengine.Models.Vod.Business.OverrideParams.TranscodeVideo:type_name -> Volcengine.Models.Vod.Business.TranscodeVideoOverride
	6, // 4: Volcengine.Models.Vod.Business.OverrideParams.TranscodeAudio:type_name -> Volcengine.Models.Vod.Business.TranscodeAudioOverride
	7, // 5: Volcengine.Models.Vod.Business.OverrideParams.Snapshot:type_name -> Volcengine.Models.Vod.Business.SnapshotOverride
	9, // 6: Volcengine.Models.Vod.Business.LogoOverride.Vars:type_name -> Volcengine.Models.Vod.Business.LogoOverride.VarsEntry
	5, // 7: Volcengine.Models.Vod.Business.TranscodeVideoOverride.Clip:type_name -> Volcengine.Models.Vod.Business.Clip
	5, // 8: Volcengine.Models.Vod.Business.TranscodeAudioOverride.Clip:type_name -> Volcengine.Models.Vod.Business.Clip
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_vod_business_vod_workflow_proto_init() }
func file_vod_business_vod_workflow_proto_init() {
	if File_vod_business_vod_workflow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vod_business_vod_workflow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VodStartWorkflowResult); i {
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
		file_vod_business_vod_workflow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowParams); i {
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
		file_vod_business_vod_workflow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OverrideParams); i {
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
		file_vod_business_vod_workflow_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoOverride); i {
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
		file_vod_business_vod_workflow_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranscodeVideoOverride); i {
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
		file_vod_business_vod_workflow_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Clip); i {
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
		file_vod_business_vod_workflow_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranscodeAudioOverride); i {
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
		file_vod_business_vod_workflow_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotOverride); i {
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
			RawDescriptor: file_vod_business_vod_workflow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vod_business_vod_workflow_proto_goTypes,
		DependencyIndexes: file_vod_business_vod_workflow_proto_depIdxs,
		MessageInfos:      file_vod_business_vod_workflow_proto_msgTypes,
	}.Build()
	File_vod_business_vod_workflow_proto = out.File
	file_vod_business_vod_workflow_proto_rawDesc = nil
	file_vod_business_vod_workflow_proto_goTypes = nil
	file_vod_business_vod_workflow_proto_depIdxs = nil
}
