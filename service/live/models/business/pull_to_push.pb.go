// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: live/business/pull_to_push.proto

package business

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

type CreatePullToPushTaskResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId string `protobuf:"bytes,1,opt,name=TaskId,proto3" json:"TaskId,omitempty"` //生成拉流转推任务ID
}

func (x *CreatePullToPushTaskResult) Reset() {
	*x = CreatePullToPushTaskResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_live_business_pull_to_push_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePullToPushTaskResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePullToPushTaskResult) ProtoMessage() {}

func (x *CreatePullToPushTaskResult) ProtoReflect() protoreflect.Message {
	mi := &file_live_business_pull_to_push_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePullToPushTaskResult.ProtoReflect.Descriptor instead.
func (*CreatePullToPushTaskResult) Descriptor() ([]byte, []int) {
	return file_live_business_pull_to_push_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePullToPushTaskResult) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

type ListPullToPushTaskResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List       []*TaskInfoItem `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`             //任务详情列表
	Pagination *Pagination     `protobuf:"bytes,2,opt,name=Pagination,proto3" json:"Pagination,omitempty"` //页码信息
}

func (x *ListPullToPushTaskResult) Reset() {
	*x = ListPullToPushTaskResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_live_business_pull_to_push_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPullToPushTaskResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPullToPushTaskResult) ProtoMessage() {}

func (x *ListPullToPushTaskResult) ProtoReflect() protoreflect.Message {
	mi := &file_live_business_pull_to_push_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPullToPushTaskResult.ProtoReflect.Descriptor instead.
func (*ListPullToPushTaskResult) Descriptor() ([]byte, []int) {
	return file_live_business_pull_to_push_proto_rawDescGZIP(), []int{1}
}

func (x *ListPullToPushTaskResult) GetList() []*TaskInfoItem {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *ListPullToPushTaskResult) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type TaskInfoItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string   `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`             //任务标题
	TaskId      string   `protobuf:"bytes,2,opt,name=TaskId,proto3" json:"TaskId,omitempty"`           //任务ID
	StartTime   string   `protobuf:"bytes,3,opt,name=StartTime,proto3" json:"StartTime,omitempty"`     //任务开始时间，UTC时间
	EndTime     string   `protobuf:"bytes,4,opt,name=EndTime,proto3" json:"EndTime,omitempty"`         //任务结束时间，UTC时间
	CallbackURL string   `protobuf:"bytes,5,opt,name=CallbackURL,proto3" json:"CallbackURL,omitempty"` //任务回调地址
	Type        int32    `protobuf:"varint,6,opt,name=Type,proto3" json:"Type,omitempty"`              //拉流转推类型，0：直播，1：点播
	CycleMode   int32    `protobuf:"varint,7,opt,name=CycleMode,proto3" json:"CycleMode,omitempty"`    //点播时，拉流地址的循环模式，Type=1时必选，-1：顺序循环
	DstAddr     string   `protobuf:"bytes,8,opt,name=DstAddr,proto3" json:"DstAddr,omitempty"`         //目标推流地址
	SrcAddr     string   `protobuf:"bytes,9,opt,name=SrcAddr,proto3" json:"SrcAddr,omitempty"`         //直播拉流地址
	SrcAddrS    []string `protobuf:"bytes,10,rep,name=SrcAddrS,proto3" json:"SrcAddrS,omitempty"`      //点播拉流地址列表
	Status      string   `protobuf:"bytes,11,opt,name=Status,proto3" json:"Status,omitempty"`          //任务状态：停用、未开始、生效中、已结束
	TaskStatus  int32    `protobuf:"varint,12,opt,name=TaskStatus,proto3" json:"TaskStatus,omitempty"` //任务状态：0：停用，1：未开始，2：生效中，3：已结束
}

func (x *TaskInfoItem) Reset() {
	*x = TaskInfoItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_live_business_pull_to_push_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskInfoItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskInfoItem) ProtoMessage() {}

func (x *TaskInfoItem) ProtoReflect() protoreflect.Message {
	mi := &file_live_business_pull_to_push_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskInfoItem.ProtoReflect.Descriptor instead.
func (*TaskInfoItem) Descriptor() ([]byte, []int) {
	return file_live_business_pull_to_push_proto_rawDescGZIP(), []int{2}
}

func (x *TaskInfoItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TaskInfoItem) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *TaskInfoItem) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *TaskInfoItem) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *TaskInfoItem) GetCallbackURL() string {
	if x != nil {
		return x.CallbackURL
	}
	return ""
}

func (x *TaskInfoItem) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *TaskInfoItem) GetCycleMode() int32 {
	if x != nil {
		return x.CycleMode
	}
	return 0
}

func (x *TaskInfoItem) GetDstAddr() string {
	if x != nil {
		return x.DstAddr
	}
	return ""
}

func (x *TaskInfoItem) GetSrcAddr() string {
	if x != nil {
		return x.SrcAddr
	}
	return ""
}

func (x *TaskInfoItem) GetSrcAddrS() []string {
	if x != nil {
		return x.SrcAddrS
	}
	return nil
}

func (x *TaskInfoItem) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TaskInfoItem) GetTaskStatus() int32 {
	if x != nil {
		return x.TaskStatus
	}
	return 0
}

var File_live_business_pull_to_push_proto protoreflect.FileDescriptor

var file_live_business_pull_to_push_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6c, 0x69, 0x76, 0x65, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f,
	0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x74, 0x6f, 0x5f, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1f, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4c,
	0x69, 0x76, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x1a, 0x23, 0x6c, 0x69, 0x76, 0x65, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x2f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x75, 0x6c, 0x6c, 0x54, 0x6f, 0x50, 0x75, 0x73, 0x68, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x22, 0xaa,
	0x01, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x6c, 0x6c, 0x54, 0x6f, 0x50, 0x75, 0x73,
	0x68, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x41, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x56, 0x6f, 0x6c, 0x63,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x49, 0x6e, 0x66, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4b,
	0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x4c, 0x69, 0x76, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd0, 0x02, 0x0a, 0x0c,
	0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x55, 0x52,
	0x4c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x55, 0x52, 0x4c, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x79, 0x63, 0x6c,
	0x65, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x43, 0x79, 0x63,
	0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x44, 0x73, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x44, 0x73, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x53, 0x72, 0x63, 0x41, 0x64, 0x64, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x53, 0x72, 0x63, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x72,
	0x63, 0x41, 0x64, 0x64, 0x72, 0x53, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x53, 0x72,
	0x63, 0x41, 0x64, 0x64, 0x72, 0x53, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0xd6,
	0x01, 0x0a, 0x2a, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x42, 0x0e, 0x50,
	0x75, 0x6c, 0x6c, 0x54, 0x6f, 0x50, 0x75, 0x73, 0x68, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x01, 0x5a,
	0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f, 0x6c, 0x63,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x6f, 0x6c, 0x63, 0x2d, 0x73, 0x64, 0x6b, 0x2d,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6c,
	0x69, 0x76, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0xa0, 0x01, 0x01, 0xd8, 0x01, 0x01, 0xc2, 0x02, 0x00, 0xca, 0x02, 0x21, 0x56,
	0x6f, 0x6c, 0x63, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x4c, 0x69, 0x76, 0x65,
	0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5c, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0xe2, 0x02, 0x24, 0x56, 0x6f, 0x6c, 0x63, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c,
	0x4c, 0x69, 0x76, 0x65, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_live_business_pull_to_push_proto_rawDescOnce sync.Once
	file_live_business_pull_to_push_proto_rawDescData = file_live_business_pull_to_push_proto_rawDesc
)

func file_live_business_pull_to_push_proto_rawDescGZIP() []byte {
	file_live_business_pull_to_push_proto_rawDescOnce.Do(func() {
		file_live_business_pull_to_push_proto_rawDescData = protoimpl.X.CompressGZIP(file_live_business_pull_to_push_proto_rawDescData)
	})
	return file_live_business_pull_to_push_proto_rawDescData
}

var file_live_business_pull_to_push_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_live_business_pull_to_push_proto_goTypes = []interface{}{
	(*CreatePullToPushTaskResult)(nil), // 0: Volcengine.Live.Models.Business.CreatePullToPushTaskResult
	(*ListPullToPushTaskResult)(nil),   // 1: Volcengine.Live.Models.Business.ListPullToPushTaskResult
	(*TaskInfoItem)(nil),               // 2: Volcengine.Live.Models.Business.TaskInfoItem
	(*Pagination)(nil),                 // 3: Volcengine.Live.Models.Business.Pagination
}
var file_live_business_pull_to_push_proto_depIdxs = []int32{
	2, // 0: Volcengine.Live.Models.Business.ListPullToPushTaskResult.List:type_name -> Volcengine.Live.Models.Business.TaskInfoItem
	3, // 1: Volcengine.Live.Models.Business.ListPullToPushTaskResult.Pagination:type_name -> Volcengine.Live.Models.Business.Pagination
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_live_business_pull_to_push_proto_init() }
func file_live_business_pull_to_push_proto_init() {
	if File_live_business_pull_to_push_proto != nil {
		return
	}
	file_live_business_snapshot_manage_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_live_business_pull_to_push_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePullToPushTaskResult); i {
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
		file_live_business_pull_to_push_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPullToPushTaskResult); i {
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
		file_live_business_pull_to_push_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskInfoItem); i {
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
			RawDescriptor: file_live_business_pull_to_push_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_live_business_pull_to_push_proto_goTypes,
		DependencyIndexes: file_live_business_pull_to_push_proto_depIdxs,
		MessageInfos:      file_live_business_pull_to_push_proto_msgTypes,
	}.Build()
	File_live_business_pull_to_push_proto = out.File
	file_live_business_pull_to_push_proto_rawDesc = nil
	file_live_business_pull_to_push_proto_goTypes = nil
	file_live_business_pull_to_push_proto_depIdxs = nil
}