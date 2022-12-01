// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: gorush.proto

package proto

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NotificationRequest_Priority int32

const (
	NotificationRequest_NORMAL NotificationRequest_Priority = 0
	NotificationRequest_HIGH   NotificationRequest_Priority = 1
)

// Enum value maps for NotificationRequest_Priority.
var (
	NotificationRequest_Priority_name = map[int32]string{
		0: "NORMAL",
		1: "HIGH",
	}
	NotificationRequest_Priority_value = map[string]int32{
		"NORMAL": 0,
		"HIGH":   1,
	}
)

func (x NotificationRequest_Priority) Enum() *NotificationRequest_Priority {
	p := new(NotificationRequest_Priority)
	*p = x
	return p
}

func (x NotificationRequest_Priority) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationRequest_Priority) Descriptor() protoreflect.EnumDescriptor {
	return file_gorush_proto_enumTypes[0].Descriptor()
}

func (NotificationRequest_Priority) Type() protoreflect.EnumType {
	return &file_gorush_proto_enumTypes[0]
}

func (x NotificationRequest_Priority) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationRequest_Priority.Descriptor instead.
func (NotificationRequest_Priority) EnumDescriptor() ([]byte, []int) {
	return file_gorush_proto_rawDescGZIP(), []int{1, 0}
}

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN     HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING     HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING HealthCheckResponse_ServingStatus = 2
)

// Enum value maps for HealthCheckResponse_ServingStatus.
var (
	HealthCheckResponse_ServingStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "SERVING",
		2: "NOT_SERVING",
	}
	HealthCheckResponse_ServingStatus_value = map[string]int32{
		"UNKNOWN":     0,
		"SERVING":     1,
		"NOT_SERVING": 2,
	}
)

func (x HealthCheckResponse_ServingStatus) Enum() *HealthCheckResponse_ServingStatus {
	p := new(HealthCheckResponse_ServingStatus)
	*p = x
	return p
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HealthCheckResponse_ServingStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_gorush_proto_enumTypes[1].Descriptor()
}

func (HealthCheckResponse_ServingStatus) Type() protoreflect.EnumType {
	return &file_gorush_proto_enumTypes[1]
}

func (x HealthCheckResponse_ServingStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HealthCheckResponse_ServingStatus.Descriptor instead.
func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return file_gorush_proto_rawDescGZIP(), []int{4, 0}
}

type Alert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title        string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Body         string   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Subtitle     string   `protobuf:"bytes,3,opt,name=subtitle,proto3" json:"subtitle,omitempty"`
	Action       string   `protobuf:"bytes,4,opt,name=action,proto3" json:"action,omitempty"`
	ActionLocKey string   `protobuf:"bytes,5,opt,name=actionLocKey,proto3" json:"actionLocKey,omitempty"`
	LaunchImage  string   `protobuf:"bytes,6,opt,name=launchImage,proto3" json:"launchImage,omitempty"`
	LocKey       string   `protobuf:"bytes,7,opt,name=locKey,proto3" json:"locKey,omitempty"`
	TitleLocKey  string   `protobuf:"bytes,8,opt,name=titleLocKey,proto3" json:"titleLocKey,omitempty"`
	LocArgs      []string `protobuf:"bytes,9,rep,name=locArgs,proto3" json:"locArgs,omitempty"`
	TitleLocArgs []string `protobuf:"bytes,10,rep,name=titleLocArgs,proto3" json:"titleLocArgs,omitempty"`
}

func (x *Alert) Reset() {
	*x = Alert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gorush_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alert) ProtoMessage() {}

func (x *Alert) ProtoReflect() protoreflect.Message {
	mi := &file_gorush_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alert.ProtoReflect.Descriptor instead.
func (*Alert) Descriptor() ([]byte, []int) {
	return file_gorush_proto_rawDescGZIP(), []int{0}
}

func (x *Alert) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Alert) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Alert) GetSubtitle() string {
	if x != nil {
		return x.Subtitle
	}
	return ""
}

func (x *Alert) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *Alert) GetActionLocKey() string {
	if x != nil {
		return x.ActionLocKey
	}
	return ""
}

func (x *Alert) GetLaunchImage() string {
	if x != nil {
		return x.LaunchImage
	}
	return ""
}

func (x *Alert) GetLocKey() string {
	if x != nil {
		return x.LocKey
	}
	return ""
}

func (x *Alert) GetTitleLocKey() string {
	if x != nil {
		return x.TitleLocKey
	}
	return ""
}

func (x *Alert) GetLocArgs() []string {
	if x != nil {
		return x.LocArgs
	}
	return nil
}

func (x *Alert) GetTitleLocArgs() []string {
	if x != nil {
		return x.TitleLocArgs
	}
	return nil
}

type NotificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tokens           []string                     `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
	Platform         int32                        `protobuf:"varint,2,opt,name=platform,proto3" json:"platform,omitempty"`
	Message          string                       `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Title            string                       `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Topic            string                       `protobuf:"bytes,5,opt,name=topic,proto3" json:"topic,omitempty"`
	Key              string                       `protobuf:"bytes,6,opt,name=key,proto3" json:"key,omitempty"`
	Badge            int32                        `protobuf:"varint,7,opt,name=badge,proto3" json:"badge,omitempty"`
	Category         string                       `protobuf:"bytes,8,opt,name=category,proto3" json:"category,omitempty"`
	Alert            *Alert                       `protobuf:"bytes,9,opt,name=alert,proto3" json:"alert,omitempty"`
	Sound            string                       `protobuf:"bytes,10,opt,name=sound,proto3" json:"sound,omitempty"`
	ContentAvailable bool                         `protobuf:"varint,11,opt,name=contentAvailable,proto3" json:"contentAvailable,omitempty"`
	ThreadID         string                       `protobuf:"bytes,12,opt,name=threadID,proto3" json:"threadID,omitempty"`
	MutableContent   bool                         `protobuf:"varint,13,opt,name=mutableContent,proto3" json:"mutableContent,omitempty"`
	Data             *structpb.Struct             `protobuf:"bytes,14,opt,name=data,proto3" json:"data,omitempty"`
	Image            string                       `protobuf:"bytes,15,opt,name=image,proto3" json:"image,omitempty"`
	Priority         NotificationRequest_Priority `protobuf:"varint,16,opt,name=priority,proto3,enum=proto.NotificationRequest_Priority" json:"priority,omitempty"`
	ID               string                       `protobuf:"bytes,17,opt,name=ID,proto3" json:"ID,omitempty"`
	PushType         string                       `protobuf:"bytes,18,opt,name=pushType,proto3" json:"pushType,omitempty"`
}

func (x *NotificationRequest) Reset() {
	*x = NotificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gorush_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationRequest) ProtoMessage() {}

func (x *NotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gorush_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationRequest.ProtoReflect.Descriptor instead.
func (*NotificationRequest) Descriptor() ([]byte, []int) {
	return file_gorush_proto_rawDescGZIP(), []int{1}
}

func (x *NotificationRequest) GetTokens() []string {
	if x != nil {
		return x.Tokens
	}
	return nil
}

func (x *NotificationRequest) GetPlatform() int32 {
	if x != nil {
		return x.Platform
	}
	return 0
}

func (x *NotificationRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *NotificationRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NotificationRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *NotificationRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *NotificationRequest) GetBadge() int32 {
	if x != nil {
		return x.Badge
	}
	return 0
}

func (x *NotificationRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *NotificationRequest) GetAlert() *Alert {
	if x != nil {
		return x.Alert
	}
	return nil
}

func (x *NotificationRequest) GetSound() string {
	if x != nil {
		return x.Sound
	}
	return ""
}

func (x *NotificationRequest) GetContentAvailable() bool {
	if x != nil {
		return x.ContentAvailable
	}
	return false
}

func (x *NotificationRequest) GetThreadID() string {
	if x != nil {
		return x.ThreadID
	}
	return ""
}

func (x *NotificationRequest) GetMutableContent() bool {
	if x != nil {
		return x.MutableContent
	}
	return false
}

func (x *NotificationRequest) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *NotificationRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *NotificationRequest) GetPriority() NotificationRequest_Priority {
	if x != nil {
		return x.Priority
	}
	return NotificationRequest_NORMAL
}

func (x *NotificationRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *NotificationRequest) GetPushType() string {
	if x != nil {
		return x.PushType
	}
	return ""
}

type NotificationReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool  `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Counts  int32 `protobuf:"varint,2,opt,name=counts,proto3" json:"counts,omitempty"`
}

func (x *NotificationReply) Reset() {
	*x = NotificationReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gorush_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationReply) ProtoMessage() {}

func (x *NotificationReply) ProtoReflect() protoreflect.Message {
	mi := &file_gorush_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationReply.ProtoReflect.Descriptor instead.
func (*NotificationReply) Descriptor() ([]byte, []int) {
	return file_gorush_proto_rawDescGZIP(), []int{2}
}

func (x *NotificationReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *NotificationReply) GetCounts() int32 {
	if x != nil {
		return x.Counts
	}
	return 0
}

type HealthCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *HealthCheckRequest) Reset() {
	*x = HealthCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gorush_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckRequest) ProtoMessage() {}

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gorush_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckRequest.ProtoReflect.Descriptor instead.
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_gorush_proto_rawDescGZIP(), []int{3}
}

func (x *HealthCheckRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type HealthCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,proto3,enum=proto.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
}

func (x *HealthCheckResponse) Reset() {
	*x = HealthCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gorush_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResponse) ProtoMessage() {}

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gorush_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResponse.ProtoReflect.Descriptor instead.
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_gorush_proto_rawDescGZIP(), []int{4}
}

func (x *HealthCheckResponse) GetStatus() HealthCheckResponse_ServingStatus {
	if x != nil {
		return x.Status
	}
	return HealthCheckResponse_UNKNOWN
}

var File_gorush_proto protoreflect.FileDescriptor

var file_gorush_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x67, 0x6f, 0x72, 0x75, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x02, 0x0a, 0x05, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x4b, 0x65, 0x79, 0x12,
	0x20, 0x0a, 0x0b, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6c,
	0x6f, 0x63, 0x41, 0x72, 0x67, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f,
	0x63, 0x41, 0x72, 0x67, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x4c, 0x6f,
	0x63, 0x41, 0x72, 0x67, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x41, 0x72, 0x67, 0x73, 0x22, 0xcf, 0x04, 0x0a, 0x13, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x62, 0x61, 0x64, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x61,
	0x64, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x22, 0x0a, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x05, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x41, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49,
	0x44, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49,
	0x44, 0x12, 0x26, 0x0a, 0x0e, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x6d, 0x75, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x3f, 0x0a, 0x08,
	0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x75, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x75, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x22, 0x20, 0x0a, 0x08, 0x50, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x49, 0x47, 0x48, 0x10, 0x01, 0x22, 0x45, 0x0a, 0x11, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x22, 0x2e, 0x0a, 0x12, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x13, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3a, 0x0a, 0x0d,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x45,
	0x52, 0x56, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x54, 0x5f, 0x53,
	0x45, 0x52, 0x56, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x32, 0x48, 0x0a, 0x06, 0x47, 0x6f, 0x72, 0x75,
	0x73, 0x68, 0x12, 0x3e, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x32, 0x48, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x3e, 0x0a, 0x05,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gorush_proto_rawDescOnce sync.Once
	file_gorush_proto_rawDescData = file_gorush_proto_rawDesc
)

func file_gorush_proto_rawDescGZIP() []byte {
	file_gorush_proto_rawDescOnce.Do(func() {
		file_gorush_proto_rawDescData = protoimpl.X.CompressGZIP(file_gorush_proto_rawDescData)
	})
	return file_gorush_proto_rawDescData
}

var (
	file_gorush_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
	file_gorush_proto_msgTypes  = make([]protoimpl.MessageInfo, 5)
	file_gorush_proto_goTypes   = []interface{}{
		(NotificationRequest_Priority)(0),      // 0: proto.NotificationRequest.Priority
		(HealthCheckResponse_ServingStatus)(0), // 1: proto.HealthCheckResponse.ServingStatus
		(*Alert)(nil),                          // 2: proto.Alert
		(*NotificationRequest)(nil),            // 3: proto.NotificationRequest
		(*NotificationReply)(nil),              // 4: proto.NotificationReply
		(*HealthCheckRequest)(nil),             // 5: proto.HealthCheckRequest
		(*HealthCheckResponse)(nil),            // 6: proto.HealthCheckResponse
		(*structpb.Struct)(nil),                // 7: google.protobuf.Struct
	}
)

var file_gorush_proto_depIdxs = []int32{
	2, // 0: proto.NotificationRequest.alert:type_name -> proto.Alert
	7, // 1: proto.NotificationRequest.data:type_name -> google.protobuf.Struct
	0, // 2: proto.NotificationRequest.priority:type_name -> proto.NotificationRequest.Priority
	1, // 3: proto.HealthCheckResponse.status:type_name -> proto.HealthCheckResponse.ServingStatus
	3, // 4: proto.Gorush.Send:input_type -> proto.NotificationRequest
	5, // 5: proto.Health.Check:input_type -> proto.HealthCheckRequest
	4, // 6: proto.Gorush.Send:output_type -> proto.NotificationReply
	6, // 7: proto.Health.Check:output_type -> proto.HealthCheckResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_gorush_proto_init() }
func file_gorush_proto_init() {
	if File_gorush_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gorush_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Alert); i {
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
		file_gorush_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationRequest); i {
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
		file_gorush_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationReply); i {
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
		file_gorush_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckRequest); i {
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
		file_gorush_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckResponse); i {
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
			RawDescriptor: file_gorush_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_gorush_proto_goTypes,
		DependencyIndexes: file_gorush_proto_depIdxs,
		EnumInfos:         file_gorush_proto_enumTypes,
		MessageInfos:      file_gorush_proto_msgTypes,
	}.Build()
	File_gorush_proto = out.File
	file_gorush_proto_rawDesc = nil
	file_gorush_proto_goTypes = nil
	file_gorush_proto_depIdxs = nil
}
