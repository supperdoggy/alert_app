// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: alert_proto/alert.proto

package alert_proto

import (
	context "context"
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

type Alert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag                string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Address            string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Title              string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Body               string `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	ExpirationDatetime int64  `protobuf:"varint,5,opt,name=expiration_datetime,json=expirationDatetime,proto3" json:"expiration_datetime,omitempty"`
}

func (x *Alert) Reset() {
	*x = Alert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alert_proto_alert_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alert) ProtoMessage() {}

func (x *Alert) ProtoReflect() protoreflect.Message {
	mi := &file_alert_proto_alert_proto_msgTypes[0]
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
	return file_alert_proto_alert_proto_rawDescGZIP(), []int{0}
}

func (x *Alert) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Alert) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
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

func (x *Alert) GetExpirationDatetime() int64 {
	if x != nil {
		return x.ExpirationDatetime
	}
	return 0
}

type GetAlertsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetAlertsRequest) Reset() {
	*x = GetAlertsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alert_proto_alert_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlertsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlertsRequest) ProtoMessage() {}

func (x *GetAlertsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alert_proto_alert_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlertsRequest.ProtoReflect.Descriptor instead.
func (*GetAlertsRequest) Descriptor() ([]byte, []int) {
	return file_alert_proto_alert_proto_rawDescGZIP(), []int{1}
}

func (x *GetAlertsRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetAlertsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetAlertsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alert *Alert `protobuf:"bytes,1,opt,name=alert,proto3" json:"alert,omitempty"`
}

func (x *GetAlertsResponse) Reset() {
	*x = GetAlertsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alert_proto_alert_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlertsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlertsResponse) ProtoMessage() {}

func (x *GetAlertsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_alert_proto_alert_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlertsResponse.ProtoReflect.Descriptor instead.
func (*GetAlertsResponse) Descriptor() ([]byte, []int) {
	return file_alert_proto_alert_proto_rawDescGZIP(), []int{2}
}

func (x *GetAlertsResponse) GetAlert() *Alert {
	if x != nil {
		return x.Alert
	}
	return nil
}

type NewAlertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alert *Alert `protobuf:"bytes,1,opt,name=alert,proto3" json:"alert,omitempty"`
}

func (x *NewAlertRequest) Reset() {
	*x = NewAlertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alert_proto_alert_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewAlertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewAlertRequest) ProtoMessage() {}

func (x *NewAlertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alert_proto_alert_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewAlertRequest.ProtoReflect.Descriptor instead.
func (*NewAlertRequest) Descriptor() ([]byte, []int) {
	return file_alert_proto_alert_proto_rawDescGZIP(), []int{3}
}

func (x *NewAlertRequest) GetAlert() *Alert {
	if x != nil {
		return x.Alert
	}
	return nil
}

type NewAlertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alert *Alert `protobuf:"bytes,1,opt,name=alert,proto3" json:"alert,omitempty"`
}

func (x *NewAlertResponse) Reset() {
	*x = NewAlertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alert_proto_alert_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewAlertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewAlertResponse) ProtoMessage() {}

func (x *NewAlertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_alert_proto_alert_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewAlertResponse.ProtoReflect.Descriptor instead.
func (*NewAlertResponse) Descriptor() ([]byte, []int) {
	return file_alert_proto_alert_proto_rawDescGZIP(), []int{4}
}

func (x *NewAlertResponse) GetAlert() *Alert {
	if x != nil {
		return x.Alert
	}
	return nil
}

type GetAlertStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAlertStreamRequest) Reset() {
	*x = GetAlertStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alert_proto_alert_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlertStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlertStreamRequest) ProtoMessage() {}

func (x *GetAlertStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alert_proto_alert_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlertStreamRequest.ProtoReflect.Descriptor instead.
func (*GetAlertStreamRequest) Descriptor() ([]byte, []int) {
	return file_alert_proto_alert_proto_rawDescGZIP(), []int{5}
}

type GetAlertStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alert *Alert `protobuf:"bytes,1,opt,name=alert,proto3" json:"alert,omitempty"`
}

func (x *GetAlertStreamResponse) Reset() {
	*x = GetAlertStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alert_proto_alert_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlertStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlertStreamResponse) ProtoMessage() {}

func (x *GetAlertStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_alert_proto_alert_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlertStreamResponse.ProtoReflect.Descriptor instead.
func (*GetAlertStreamResponse) Descriptor() ([]byte, []int) {
	return file_alert_proto_alert_proto_rawDescGZIP(), []int{6}
}

func (x *GetAlertStreamResponse) GetAlert() *Alert {
	if x != nil {
		return x.Alert
	}
	return nil
}

type GetAllActiveAlertsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetAllActiveAlertsRequest) Reset() {
	*x = GetAllActiveAlertsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alert_proto_alert_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllActiveAlertsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllActiveAlertsRequest) ProtoMessage() {}

func (x *GetAllActiveAlertsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alert_proto_alert_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllActiveAlertsRequest.ProtoReflect.Descriptor instead.
func (*GetAllActiveAlertsRequest) Descriptor() ([]byte, []int) {
	return file_alert_proto_alert_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllActiveAlertsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetAllActiveAlertsRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetAllActiveAlertsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alerts []*Alert `protobuf:"bytes,1,rep,name=alerts,proto3" json:"alerts,omitempty"`
}

func (x *GetAllActiveAlertsResponse) Reset() {
	*x = GetAllActiveAlertsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alert_proto_alert_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllActiveAlertsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllActiveAlertsResponse) ProtoMessage() {}

func (x *GetAllActiveAlertsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_alert_proto_alert_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllActiveAlertsResponse.ProtoReflect.Descriptor instead.
func (*GetAllActiveAlertsResponse) Descriptor() ([]byte, []int) {
	return file_alert_proto_alert_proto_rawDescGZIP(), []int{8}
}

func (x *GetAllActiveAlertsResponse) GetAlerts() []*Alert {
	if x != nil {
		return x.Alerts
	}
	return nil
}

var File_alert_proto_alert_proto protoreflect.FileDescriptor

var file_alert_proto_alert_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x70, 0x62, 0x22, 0x8e, 0x01, 0x0a, 0x05, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x12, 0x2f, 0x0a, 0x13, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x12, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x74,
	0x69, 0x6d, 0x65, 0x22, 0x41, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x70, 0x62, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x05, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x22, 0x37, 0x0a, 0x0f, 0x4e, 0x65, 0x77, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x52, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x22, 0x38, 0x0a, 0x10, 0x4e, 0x65,
	0x77, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24,
	0x0a, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x05, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3e, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x70, 0x62,
	0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x22, 0x4a, 0x0a,
	0x19, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x44, 0x0a, 0x1a, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x70,
	0x62, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x32,
	0xbe, 0x01, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x12, 0x19, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x5f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x73, 0x12, 0x22, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x32, 0x91, 0x02, 0x0a, 0x14, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x4e, 0x65, 0x77,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x18, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x70, 0x62, 0x2e,
	0x4e, 0x65, 0x77, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x4e, 0x65, 0x77, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1e,
	0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x30, 0x01, 0x12, 0x5f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x12, 0x22, 0x2e, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_alert_proto_alert_proto_rawDescOnce sync.Once
	file_alert_proto_alert_proto_rawDescData = file_alert_proto_alert_proto_rawDesc
)

func file_alert_proto_alert_proto_rawDescGZIP() []byte {
	file_alert_proto_alert_proto_rawDescOnce.Do(func() {
		file_alert_proto_alert_proto_rawDescData = protoimpl.X.CompressGZIP(file_alert_proto_alert_proto_rawDescData)
	})
	return file_alert_proto_alert_proto_rawDescData
}

var file_alert_proto_alert_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_alert_proto_alert_proto_goTypes = []interface{}{
	(*Alert)(nil),                      // 0: alertpb.Alert
	(*GetAlertsRequest)(nil),           // 1: alertpb.GetAlertsRequest
	(*GetAlertsResponse)(nil),          // 2: alertpb.GetAlertsResponse
	(*NewAlertRequest)(nil),            // 3: alertpb.NewAlertRequest
	(*NewAlertResponse)(nil),           // 4: alertpb.NewAlertResponse
	(*GetAlertStreamRequest)(nil),      // 5: alertpb.GetAlertStreamRequest
	(*GetAlertStreamResponse)(nil),     // 6: alertpb.GetAlertStreamResponse
	(*GetAllActiveAlertsRequest)(nil),  // 7: alertpb.GetAllActiveAlertsRequest
	(*GetAllActiveAlertsResponse)(nil), // 8: alertpb.GetAllActiveAlertsResponse
}
var file_alert_proto_alert_proto_depIdxs = []int32{
	0,  // 0: alertpb.GetAlertsResponse.alert:type_name -> alertpb.Alert
	0,  // 1: alertpb.NewAlertRequest.alert:type_name -> alertpb.Alert
	0,  // 2: alertpb.NewAlertResponse.alert:type_name -> alertpb.Alert
	0,  // 3: alertpb.GetAlertStreamResponse.alert:type_name -> alertpb.Alert
	0,  // 4: alertpb.GetAllActiveAlertsResponse.alerts:type_name -> alertpb.Alert
	1,  // 5: alertpb.NotificationService.GetAlerts:input_type -> alertpb.GetAlertsRequest
	7,  // 6: alertpb.NotificationService.GetAllActiveAlerts:input_type -> alertpb.GetAllActiveAlertsRequest
	3,  // 7: alertpb.AlertDatabaseService.NewAlert:input_type -> alertpb.NewAlertRequest
	5,  // 8: alertpb.AlertDatabaseService.GetAlertStream:input_type -> alertpb.GetAlertStreamRequest
	7,  // 9: alertpb.AlertDatabaseService.GetAllActiveAlerts:input_type -> alertpb.GetAllActiveAlertsRequest
	2,  // 10: alertpb.NotificationService.GetAlerts:output_type -> alertpb.GetAlertsResponse
	8,  // 11: alertpb.NotificationService.GetAllActiveAlerts:output_type -> alertpb.GetAllActiveAlertsResponse
	4,  // 12: alertpb.AlertDatabaseService.NewAlert:output_type -> alertpb.NewAlertResponse
	6,  // 13: alertpb.AlertDatabaseService.GetAlertStream:output_type -> alertpb.GetAlertStreamResponse
	8,  // 14: alertpb.AlertDatabaseService.GetAllActiveAlerts:output_type -> alertpb.GetAllActiveAlertsResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_alert_proto_alert_proto_init() }
func file_alert_proto_alert_proto_init() {
	if File_alert_proto_alert_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_alert_proto_alert_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_alert_proto_alert_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAlertsRequest); i {
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
		file_alert_proto_alert_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAlertsResponse); i {
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
		file_alert_proto_alert_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewAlertRequest); i {
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
		file_alert_proto_alert_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewAlertResponse); i {
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
		file_alert_proto_alert_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAlertStreamRequest); i {
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
		file_alert_proto_alert_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAlertStreamResponse); i {
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
		file_alert_proto_alert_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllActiveAlertsRequest); i {
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
		file_alert_proto_alert_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllActiveAlertsResponse); i {
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
			RawDescriptor: file_alert_proto_alert_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_alert_proto_alert_proto_goTypes,
		DependencyIndexes: file_alert_proto_alert_proto_depIdxs,
		MessageInfos:      file_alert_proto_alert_proto_msgTypes,
	}.Build()
	File_alert_proto_alert_proto = out.File
	file_alert_proto_alert_proto_rawDesc = nil
	file_alert_proto_alert_proto_goTypes = nil
	file_alert_proto_alert_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotificationServiceClient interface {
	GetAlerts(ctx context.Context, in *GetAlertsRequest, opts ...grpc.CallOption) (NotificationService_GetAlertsClient, error)
	GetAllActiveAlerts(ctx context.Context, in *GetAllActiveAlertsRequest, opts ...grpc.CallOption) (*GetAllActiveAlertsResponse, error)
}

type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) GetAlerts(ctx context.Context, in *GetAlertsRequest, opts ...grpc.CallOption) (NotificationService_GetAlertsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NotificationService_serviceDesc.Streams[0], "/alertpb.NotificationService/GetAlerts", opts...)
	if err != nil {
		return nil, err
	}
	x := &notificationServiceGetAlertsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NotificationService_GetAlertsClient interface {
	Recv() (*GetAlertsResponse, error)
	grpc.ClientStream
}

type notificationServiceGetAlertsClient struct {
	grpc.ClientStream
}

func (x *notificationServiceGetAlertsClient) Recv() (*GetAlertsResponse, error) {
	m := new(GetAlertsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *notificationServiceClient) GetAllActiveAlerts(ctx context.Context, in *GetAllActiveAlertsRequest, opts ...grpc.CallOption) (*GetAllActiveAlertsResponse, error) {
	out := new(GetAllActiveAlertsResponse)
	err := c.cc.Invoke(ctx, "/alertpb.NotificationService/GetAllActiveAlerts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
type NotificationServiceServer interface {
	GetAlerts(*GetAlertsRequest, NotificationService_GetAlertsServer) error
	GetAllActiveAlerts(context.Context, *GetAllActiveAlertsRequest) (*GetAllActiveAlertsResponse, error)
}

// UnimplementedNotificationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNotificationServiceServer struct {
}

func (*UnimplementedNotificationServiceServer) GetAlerts(*GetAlertsRequest, NotificationService_GetAlertsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAlerts not implemented")
}
func (*UnimplementedNotificationServiceServer) GetAllActiveAlerts(context.Context, *GetAllActiveAlertsRequest) (*GetAllActiveAlertsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllActiveAlerts not implemented")
}

func RegisterNotificationServiceServer(s *grpc.Server, srv NotificationServiceServer) {
	s.RegisterService(&_NotificationService_serviceDesc, srv)
}

func _NotificationService_GetAlerts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAlertsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NotificationServiceServer).GetAlerts(m, &notificationServiceGetAlertsServer{stream})
}

type NotificationService_GetAlertsServer interface {
	Send(*GetAlertsResponse) error
	grpc.ServerStream
}

type notificationServiceGetAlertsServer struct {
	grpc.ServerStream
}

func (x *notificationServiceGetAlertsServer) Send(m *GetAlertsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _NotificationService_GetAllActiveAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllActiveAlertsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).GetAllActiveAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alertpb.NotificationService/GetAllActiveAlerts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).GetAllActiveAlerts(ctx, req.(*GetAllActiveAlertsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NotificationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "alertpb.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllActiveAlerts",
			Handler:    _NotificationService_GetAllActiveAlerts_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAlerts",
			Handler:       _NotificationService_GetAlerts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "alert_proto/alert.proto",
}

// AlertDatabaseServiceClient is the client API for AlertDatabaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlertDatabaseServiceClient interface {
	NewAlert(ctx context.Context, in *NewAlertRequest, opts ...grpc.CallOption) (*NewAlertResponse, error)
	GetAlertStream(ctx context.Context, in *GetAlertStreamRequest, opts ...grpc.CallOption) (AlertDatabaseService_GetAlertStreamClient, error)
	GetAllActiveAlerts(ctx context.Context, in *GetAllActiveAlertsRequest, opts ...grpc.CallOption) (*GetAllActiveAlertsResponse, error)
}

type alertDatabaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAlertDatabaseServiceClient(cc grpc.ClientConnInterface) AlertDatabaseServiceClient {
	return &alertDatabaseServiceClient{cc}
}

func (c *alertDatabaseServiceClient) NewAlert(ctx context.Context, in *NewAlertRequest, opts ...grpc.CallOption) (*NewAlertResponse, error) {
	out := new(NewAlertResponse)
	err := c.cc.Invoke(ctx, "/alertpb.AlertDatabaseService/NewAlert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertDatabaseServiceClient) GetAlertStream(ctx context.Context, in *GetAlertStreamRequest, opts ...grpc.CallOption) (AlertDatabaseService_GetAlertStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AlertDatabaseService_serviceDesc.Streams[0], "/alertpb.AlertDatabaseService/GetAlertStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &alertDatabaseServiceGetAlertStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AlertDatabaseService_GetAlertStreamClient interface {
	Recv() (*GetAlertStreamResponse, error)
	grpc.ClientStream
}

type alertDatabaseServiceGetAlertStreamClient struct {
	grpc.ClientStream
}

func (x *alertDatabaseServiceGetAlertStreamClient) Recv() (*GetAlertStreamResponse, error) {
	m := new(GetAlertStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *alertDatabaseServiceClient) GetAllActiveAlerts(ctx context.Context, in *GetAllActiveAlertsRequest, opts ...grpc.CallOption) (*GetAllActiveAlertsResponse, error) {
	out := new(GetAllActiveAlertsResponse)
	err := c.cc.Invoke(ctx, "/alertpb.AlertDatabaseService/GetAllActiveAlerts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlertDatabaseServiceServer is the server API for AlertDatabaseService service.
type AlertDatabaseServiceServer interface {
	NewAlert(context.Context, *NewAlertRequest) (*NewAlertResponse, error)
	GetAlertStream(*GetAlertStreamRequest, AlertDatabaseService_GetAlertStreamServer) error
	GetAllActiveAlerts(context.Context, *GetAllActiveAlertsRequest) (*GetAllActiveAlertsResponse, error)
}

// UnimplementedAlertDatabaseServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAlertDatabaseServiceServer struct {
}

func (*UnimplementedAlertDatabaseServiceServer) NewAlert(context.Context, *NewAlertRequest) (*NewAlertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewAlert not implemented")
}
func (*UnimplementedAlertDatabaseServiceServer) GetAlertStream(*GetAlertStreamRequest, AlertDatabaseService_GetAlertStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAlertStream not implemented")
}
func (*UnimplementedAlertDatabaseServiceServer) GetAllActiveAlerts(context.Context, *GetAllActiveAlertsRequest) (*GetAllActiveAlertsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllActiveAlerts not implemented")
}

func RegisterAlertDatabaseServiceServer(s *grpc.Server, srv AlertDatabaseServiceServer) {
	s.RegisterService(&_AlertDatabaseService_serviceDesc, srv)
}

func _AlertDatabaseService_NewAlert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertDatabaseServiceServer).NewAlert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alertpb.AlertDatabaseService/NewAlert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertDatabaseServiceServer).NewAlert(ctx, req.(*NewAlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertDatabaseService_GetAlertStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAlertStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AlertDatabaseServiceServer).GetAlertStream(m, &alertDatabaseServiceGetAlertStreamServer{stream})
}

type AlertDatabaseService_GetAlertStreamServer interface {
	Send(*GetAlertStreamResponse) error
	grpc.ServerStream
}

type alertDatabaseServiceGetAlertStreamServer struct {
	grpc.ServerStream
}

func (x *alertDatabaseServiceGetAlertStreamServer) Send(m *GetAlertStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _AlertDatabaseService_GetAllActiveAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllActiveAlertsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertDatabaseServiceServer).GetAllActiveAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alertpb.AlertDatabaseService/GetAllActiveAlerts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertDatabaseServiceServer).GetAllActiveAlerts(ctx, req.(*GetAllActiveAlertsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AlertDatabaseService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "alertpb.AlertDatabaseService",
	HandlerType: (*AlertDatabaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewAlert",
			Handler:    _AlertDatabaseService_NewAlert_Handler,
		},
		{
			MethodName: "GetAllActiveAlerts",
			Handler:    _AlertDatabaseService_GetAllActiveAlerts_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAlertStream",
			Handler:       _AlertDatabaseService_GetAlertStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "alert_proto/alert.proto",
}
