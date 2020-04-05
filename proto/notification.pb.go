// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notification.proto

package notifications

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Notification struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content              string               `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	UserId               string               `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Time                 *timestamp.Timestamp `protobuf:"bytes,5,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{0}
}

func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Notification) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Notification) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Notification) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Notification) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

type IDtoken struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IDtoken) Reset()         { *m = IDtoken{} }
func (m *IDtoken) String() string { return proto.CompactTextString(m) }
func (*IDtoken) ProtoMessage()    {}
func (*IDtoken) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{1}
}

func (m *IDtoken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IDtoken.Unmarshal(m, b)
}
func (m *IDtoken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IDtoken.Marshal(b, m, deterministic)
}
func (m *IDtoken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDtoken.Merge(m, src)
}
func (m *IDtoken) XXX_Size() int {
	return xxx_messageInfo_IDtoken.Size(m)
}
func (m *IDtoken) XXX_DiscardUnknown() {
	xxx_messageInfo_IDtoken.DiscardUnknown(m)
}

var xxx_messageInfo_IDtoken proto.InternalMessageInfo

func (m *IDtoken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type NotificationCreateRequest struct {
	Notification         *Notification `protobuf:"bytes,1,opt,name=notification,proto3" json:"notification,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *NotificationCreateRequest) Reset()         { *m = NotificationCreateRequest{} }
func (m *NotificationCreateRequest) String() string { return proto.CompactTextString(m) }
func (*NotificationCreateRequest) ProtoMessage()    {}
func (*NotificationCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{2}
}

func (m *NotificationCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationCreateRequest.Unmarshal(m, b)
}
func (m *NotificationCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationCreateRequest.Marshal(b, m, deterministic)
}
func (m *NotificationCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationCreateRequest.Merge(m, src)
}
func (m *NotificationCreateRequest) XXX_Size() int {
	return xxx_messageInfo_NotificationCreateRequest.Size(m)
}
func (m *NotificationCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationCreateRequest proto.InternalMessageInfo

func (m *NotificationCreateRequest) GetNotification() *Notification {
	if m != nil {
		return m.Notification
	}
	return nil
}

type NotificationCreateResponse struct {
	Notification         *Notification `protobuf:"bytes,1,opt,name=notification,proto3" json:"notification,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *NotificationCreateResponse) Reset()         { *m = NotificationCreateResponse{} }
func (m *NotificationCreateResponse) String() string { return proto.CompactTextString(m) }
func (*NotificationCreateResponse) ProtoMessage()    {}
func (*NotificationCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{3}
}

func (m *NotificationCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationCreateResponse.Unmarshal(m, b)
}
func (m *NotificationCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationCreateResponse.Marshal(b, m, deterministic)
}
func (m *NotificationCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationCreateResponse.Merge(m, src)
}
func (m *NotificationCreateResponse) XXX_Size() int {
	return xxx_messageInfo_NotificationCreateResponse.Size(m)
}
func (m *NotificationCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationCreateResponse proto.InternalMessageInfo

func (m *NotificationCreateResponse) GetNotification() *Notification {
	if m != nil {
		return m.Notification
	}
	return nil
}

type NotificationReadRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotificationReadRequest) Reset()         { *m = NotificationReadRequest{} }
func (m *NotificationReadRequest) String() string { return proto.CompactTextString(m) }
func (*NotificationReadRequest) ProtoMessage()    {}
func (*NotificationReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{4}
}

func (m *NotificationReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationReadRequest.Unmarshal(m, b)
}
func (m *NotificationReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationReadRequest.Marshal(b, m, deterministic)
}
func (m *NotificationReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationReadRequest.Merge(m, src)
}
func (m *NotificationReadRequest) XXX_Size() int {
	return xxx_messageInfo_NotificationReadRequest.Size(m)
}
func (m *NotificationReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationReadRequest proto.InternalMessageInfo

func (m *NotificationReadRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type NotificationReadResponse struct {
	Notification         *Notification `protobuf:"bytes,1,opt,name=notification,proto3" json:"notification,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *NotificationReadResponse) Reset()         { *m = NotificationReadResponse{} }
func (m *NotificationReadResponse) String() string { return proto.CompactTextString(m) }
func (*NotificationReadResponse) ProtoMessage()    {}
func (*NotificationReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{5}
}

func (m *NotificationReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationReadResponse.Unmarshal(m, b)
}
func (m *NotificationReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationReadResponse.Marshal(b, m, deterministic)
}
func (m *NotificationReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationReadResponse.Merge(m, src)
}
func (m *NotificationReadResponse) XXX_Size() int {
	return xxx_messageInfo_NotificationReadResponse.Size(m)
}
func (m *NotificationReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationReadResponse proto.InternalMessageInfo

func (m *NotificationReadResponse) GetNotification() *Notification {
	if m != nil {
		return m.Notification
	}
	return nil
}

type NotificationUpdateRequest struct {
	Notification         *Notification `protobuf:"bytes,1,opt,name=notification,proto3" json:"notification,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *NotificationUpdateRequest) Reset()         { *m = NotificationUpdateRequest{} }
func (m *NotificationUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*NotificationUpdateRequest) ProtoMessage()    {}
func (*NotificationUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{6}
}

func (m *NotificationUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationUpdateRequest.Unmarshal(m, b)
}
func (m *NotificationUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationUpdateRequest.Marshal(b, m, deterministic)
}
func (m *NotificationUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationUpdateRequest.Merge(m, src)
}
func (m *NotificationUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_NotificationUpdateRequest.Size(m)
}
func (m *NotificationUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationUpdateRequest proto.InternalMessageInfo

func (m *NotificationUpdateRequest) GetNotification() *Notification {
	if m != nil {
		return m.Notification
	}
	return nil
}

type NotificationUpdateResponse struct {
	Notification         *Notification `protobuf:"bytes,1,opt,name=notification,proto3" json:"notification,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *NotificationUpdateResponse) Reset()         { *m = NotificationUpdateResponse{} }
func (m *NotificationUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*NotificationUpdateResponse) ProtoMessage()    {}
func (*NotificationUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{7}
}

func (m *NotificationUpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationUpdateResponse.Unmarshal(m, b)
}
func (m *NotificationUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationUpdateResponse.Marshal(b, m, deterministic)
}
func (m *NotificationUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationUpdateResponse.Merge(m, src)
}
func (m *NotificationUpdateResponse) XXX_Size() int {
	return xxx_messageInfo_NotificationUpdateResponse.Size(m)
}
func (m *NotificationUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationUpdateResponse proto.InternalMessageInfo

func (m *NotificationUpdateResponse) GetNotification() *Notification {
	if m != nil {
		return m.Notification
	}
	return nil
}

type NotificationDeleteRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotificationDeleteRequest) Reset()         { *m = NotificationDeleteRequest{} }
func (m *NotificationDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*NotificationDeleteRequest) ProtoMessage()    {}
func (*NotificationDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{8}
}

func (m *NotificationDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationDeleteRequest.Unmarshal(m, b)
}
func (m *NotificationDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationDeleteRequest.Marshal(b, m, deterministic)
}
func (m *NotificationDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationDeleteRequest.Merge(m, src)
}
func (m *NotificationDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_NotificationDeleteRequest.Size(m)
}
func (m *NotificationDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationDeleteRequest proto.InternalMessageInfo

func (m *NotificationDeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type NotificationDeleteResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotificationDeleteResponse) Reset()         { *m = NotificationDeleteResponse{} }
func (m *NotificationDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*NotificationDeleteResponse) ProtoMessage()    {}
func (*NotificationDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{9}
}

func (m *NotificationDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationDeleteResponse.Unmarshal(m, b)
}
func (m *NotificationDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationDeleteResponse.Marshal(b, m, deterministic)
}
func (m *NotificationDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationDeleteResponse.Merge(m, src)
}
func (m *NotificationDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_NotificationDeleteResponse.Size(m)
}
func (m *NotificationDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationDeleteResponse proto.InternalMessageInfo

func (m *NotificationDeleteResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type NotificationsListRequest struct {
	Idtoken              *IDtoken `protobuf:"bytes,1,opt,name=idtoken,proto3" json:"idtoken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotificationsListRequest) Reset()         { *m = NotificationsListRequest{} }
func (m *NotificationsListRequest) String() string { return proto.CompactTextString(m) }
func (*NotificationsListRequest) ProtoMessage()    {}
func (*NotificationsListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{10}
}

func (m *NotificationsListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationsListRequest.Unmarshal(m, b)
}
func (m *NotificationsListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationsListRequest.Marshal(b, m, deterministic)
}
func (m *NotificationsListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationsListRequest.Merge(m, src)
}
func (m *NotificationsListRequest) XXX_Size() int {
	return xxx_messageInfo_NotificationsListRequest.Size(m)
}
func (m *NotificationsListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationsListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationsListRequest proto.InternalMessageInfo

func (m *NotificationsListRequest) GetIdtoken() *IDtoken {
	if m != nil {
		return m.Idtoken
	}
	return nil
}

type NotificationsListResponse struct {
	Notification         *Notification `protobuf:"bytes,1,opt,name=notification,proto3" json:"notification,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *NotificationsListResponse) Reset()         { *m = NotificationsListResponse{} }
func (m *NotificationsListResponse) String() string { return proto.CompactTextString(m) }
func (*NotificationsListResponse) ProtoMessage()    {}
func (*NotificationsListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{11}
}

func (m *NotificationsListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationsListResponse.Unmarshal(m, b)
}
func (m *NotificationsListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationsListResponse.Marshal(b, m, deterministic)
}
func (m *NotificationsListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationsListResponse.Merge(m, src)
}
func (m *NotificationsListResponse) XXX_Size() int {
	return xxx_messageInfo_NotificationsListResponse.Size(m)
}
func (m *NotificationsListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationsListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationsListResponse proto.InternalMessageInfo

func (m *NotificationsListResponse) GetNotification() *Notification {
	if m != nil {
		return m.Notification
	}
	return nil
}

func init() {
	proto.RegisterType((*Notification)(nil), "notifications.Notification")
	proto.RegisterType((*IDtoken)(nil), "notifications.IDtoken")
	proto.RegisterType((*NotificationCreateRequest)(nil), "notifications.NotificationCreateRequest")
	proto.RegisterType((*NotificationCreateResponse)(nil), "notifications.NotificationCreateResponse")
	proto.RegisterType((*NotificationReadRequest)(nil), "notifications.NotificationReadRequest")
	proto.RegisterType((*NotificationReadResponse)(nil), "notifications.NotificationReadResponse")
	proto.RegisterType((*NotificationUpdateRequest)(nil), "notifications.NotificationUpdateRequest")
	proto.RegisterType((*NotificationUpdateResponse)(nil), "notifications.NotificationUpdateResponse")
	proto.RegisterType((*NotificationDeleteRequest)(nil), "notifications.NotificationDeleteRequest")
	proto.RegisterType((*NotificationDeleteResponse)(nil), "notifications.NotificationDeleteResponse")
	proto.RegisterType((*NotificationsListRequest)(nil), "notifications.NotificationsListRequest")
	proto.RegisterType((*NotificationsListResponse)(nil), "notifications.NotificationsListResponse")
}

func init() {
	proto.RegisterFile("notification.proto", fileDescriptor_736a457d4a5efa07)
}

var fileDescriptor_736a457d4a5efa07 = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x41, 0xcf, 0xd2, 0x40,
	0x10, 0x4d, 0x91, 0x8f, 0xea, 0x80, 0x46, 0x57, 0x23, 0x6b, 0x3d, 0x40, 0x7a, 0xd0, 0x12, 0x93,
	0x42, 0x6a, 0xe2, 0xd5, 0x83, 0x5c, 0x48, 0x88, 0x87, 0xaa, 0x27, 0x35, 0xa6, 0xb4, 0x03, 0x6e,
	0x84, 0x6e, 0xed, 0x6e, 0xfd, 0x2b, 0xfe, 0x02, 0xff, 0xa7, 0x61, 0xdb, 0x0d, 0xdd, 0xc5, 0x14,
	0x13, 0xf1, 0xc6, 0xec, 0xbc, 0x99, 0xf7, 0xe6, 0xf1, 0x00, 0x48, 0xce, 0x25, 0xdb, 0xb2, 0x34,
	0x91, 0x8c, 0xe7, 0x61, 0x51, 0x72, 0xc9, 0xc9, 0xdd, 0xf6, 0x9b, 0xf0, 0x26, 0x3b, 0xce, 0x77,
	0x7b, 0x9c, 0xab, 0xe6, 0xa6, 0xda, 0xce, 0x25, 0x3b, 0xa0, 0x90, 0xc9, 0xa1, 0xa8, 0xf1, 0xfe,
	0x4f, 0x07, 0x46, 0x6f, 0x5b, 0x23, 0xe4, 0x1e, 0xf4, 0x58, 0x46, 0x9d, 0xa9, 0x13, 0xdc, 0x89,
	0x7b, 0x2c, 0x23, 0x8f, 0xe0, 0x46, 0x32, 0xb9, 0x47, 0xda, 0x53, 0x4f, 0x75, 0x41, 0x28, 0xb8,
	0x29, 0xcf, 0x25, 0xe6, 0x92, 0xde, 0x52, 0xef, 0xba, 0x24, 0x63, 0x70, 0x2b, 0x81, 0xe5, 0x17,
	0x96, 0xd1, 0xbe, 0xea, 0x0c, 0x8e, 0xe5, 0x2a, 0x23, 0x21, 0xf4, 0x8f, 0xe4, 0xf4, 0x66, 0xea,
	0x04, 0xc3, 0xc8, 0x0b, 0x6b, 0x65, 0xa1, 0x56, 0x16, 0xbe, 0xd7, 0xca, 0x62, 0x85, 0xf3, 0x27,
	0xe0, 0xae, 0x96, 0x92, 0x7f, 0xc3, 0x5c, 0x69, 0x38, 0x7e, 0x68, 0x64, 0xd5, 0x85, 0xff, 0x09,
	0x9e, 0xb4, 0x95, 0xbf, 0x29, 0x31, 0x91, 0x18, 0xe3, 0xf7, 0x0a, 0x85, 0x24, 0xaf, 0x61, 0xd4,
	0x76, 0x42, 0x4d, 0x0e, 0xa3, 0xa7, 0xa1, 0x61, 0x4f, 0xd8, 0x9e, 0x8f, 0x8d, 0x01, 0xff, 0x33,
	0x78, 0x7f, 0xda, 0x2e, 0x0a, 0x9e, 0x0b, 0xfc, 0xf7, 0xf5, 0x33, 0x18, 0x1b, 0x5d, 0x4c, 0x32,
	0x2d, 0xdd, 0xfa, 0x06, 0xfc, 0x8f, 0x40, 0xcf, 0xa1, 0xd7, 0xd2, 0x61, 0x99, 0xf8, 0xa1, 0xc8,
	0xfe, 0xa3, 0x89, 0x7a, 0xfb, 0xb5, 0xc4, 0xbf, 0x30, 0xc5, 0x2f, 0x71, 0x8f, 0x27, 0xf1, 0xb6,
	0x8d, 0xaf, 0x4c, 0x2d, 0x1a, 0xdc, 0x68, 0xa1, 0xe0, 0x8a, 0x2a, 0x4d, 0x51, 0x08, 0x35, 0x72,
	0x3b, 0xd6, 0xa5, 0xbf, 0x36, 0xed, 0x17, 0x6b, 0x26, 0xa4, 0xe6, 0x58, 0x80, 0xcb, 0xb2, 0x53,
	0x34, 0x87, 0xd1, 0x63, 0x4b, 0x7c, 0x93, 0xe0, 0x58, 0xc3, 0x6c, 0xbf, 0x9b, 0x6d, 0x57, 0x32,
	0x24, 0xfa, 0xd5, 0x87, 0x87, 0xed, 0xf6, 0x3b, 0x2c, 0x7f, 0xb0, 0x14, 0x09, 0x03, 0x72, 0x1e,
	0x66, 0x12, 0x74, 0x2c, 0x36, 0x7e, 0x4d, 0xde, 0xec, 0x2f, 0x90, 0xcd, 0x0d, 0x29, 0xdc, 0xb7,
	0xd3, 0x4a, 0x9e, 0x75, 0x5d, 0x70, 0x4a, 0xbe, 0xf7, 0xfc, 0x22, 0xae, 0x21, 0xb1, 0xee, 0xa9,
	0x73, 0xd5, 0x79, 0x8f, 0x11, 0xec, 0xce, 0x7b, 0xac, 0x90, 0x5a, 0x54, 0x75, 0x6c, 0x3a, 0xa9,
	0x8c, 0x18, 0x76, 0x52, 0x59, 0x19, 0xfc, 0x0a, 0x0f, 0xce, 0xb2, 0x41, 0xba, 0x3c, 0x69, 0x67,
	0xd1, 0x0b, 0x2e, 0x03, 0x6b, 0x9e, 0x85, 0xb3, 0x19, 0xa8, 0x7f, 0xdd, 0x97, 0xbf, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x6f, 0x88, 0x7a, 0x94, 0x42, 0x06, 0x00, 0x00,
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
	NotificationCreate(ctx context.Context, in *NotificationCreateRequest, opts ...grpc.CallOption) (*NotificationCreateResponse, error)
	NotificationRead(ctx context.Context, in *NotificationReadRequest, opts ...grpc.CallOption) (*NotificationReadResponse, error)
	NotificationUpdate(ctx context.Context, in *NotificationUpdateRequest, opts ...grpc.CallOption) (*NotificationUpdateResponse, error)
	NotificationDelete(ctx context.Context, in *NotificationDeleteRequest, opts ...grpc.CallOption) (*NotificationDeleteResponse, error)
	NotificationsList(ctx context.Context, in *NotificationsListRequest, opts ...grpc.CallOption) (NotificationService_NotificationsListClient, error)
}

type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) NotificationCreate(ctx context.Context, in *NotificationCreateRequest, opts ...grpc.CallOption) (*NotificationCreateResponse, error) {
	out := new(NotificationCreateResponse)
	err := c.cc.Invoke(ctx, "/notifications.NotificationService/NotificationCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) NotificationRead(ctx context.Context, in *NotificationReadRequest, opts ...grpc.CallOption) (*NotificationReadResponse, error) {
	out := new(NotificationReadResponse)
	err := c.cc.Invoke(ctx, "/notifications.NotificationService/NotificationRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) NotificationUpdate(ctx context.Context, in *NotificationUpdateRequest, opts ...grpc.CallOption) (*NotificationUpdateResponse, error) {
	out := new(NotificationUpdateResponse)
	err := c.cc.Invoke(ctx, "/notifications.NotificationService/NotificationUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) NotificationDelete(ctx context.Context, in *NotificationDeleteRequest, opts ...grpc.CallOption) (*NotificationDeleteResponse, error) {
	out := new(NotificationDeleteResponse)
	err := c.cc.Invoke(ctx, "/notifications.NotificationService/NotificationDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) NotificationsList(ctx context.Context, in *NotificationsListRequest, opts ...grpc.CallOption) (NotificationService_NotificationsListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NotificationService_serviceDesc.Streams[0], "/notifications.NotificationService/NotificationsList", opts...)
	if err != nil {
		return nil, err
	}
	x := &notificationServiceNotificationsListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NotificationService_NotificationsListClient interface {
	Recv() (*NotificationsListResponse, error)
	grpc.ClientStream
}

type notificationServiceNotificationsListClient struct {
	grpc.ClientStream
}

func (x *notificationServiceNotificationsListClient) Recv() (*NotificationsListResponse, error) {
	m := new(NotificationsListResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NotificationServiceServer is the server API for NotificationService service.
type NotificationServiceServer interface {
	NotificationCreate(context.Context, *NotificationCreateRequest) (*NotificationCreateResponse, error)
	NotificationRead(context.Context, *NotificationReadRequest) (*NotificationReadResponse, error)
	NotificationUpdate(context.Context, *NotificationUpdateRequest) (*NotificationUpdateResponse, error)
	NotificationDelete(context.Context, *NotificationDeleteRequest) (*NotificationDeleteResponse, error)
	NotificationsList(*NotificationsListRequest, NotificationService_NotificationsListServer) error
}

// UnimplementedNotificationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNotificationServiceServer struct {
}

func (*UnimplementedNotificationServiceServer) NotificationCreate(ctx context.Context, req *NotificationCreateRequest) (*NotificationCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotificationCreate not implemented")
}
func (*UnimplementedNotificationServiceServer) NotificationRead(ctx context.Context, req *NotificationReadRequest) (*NotificationReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotificationRead not implemented")
}
func (*UnimplementedNotificationServiceServer) NotificationUpdate(ctx context.Context, req *NotificationUpdateRequest) (*NotificationUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotificationUpdate not implemented")
}
func (*UnimplementedNotificationServiceServer) NotificationDelete(ctx context.Context, req *NotificationDeleteRequest) (*NotificationDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotificationDelete not implemented")
}
func (*UnimplementedNotificationServiceServer) NotificationsList(req *NotificationsListRequest, srv NotificationService_NotificationsListServer) error {
	return status.Errorf(codes.Unimplemented, "method NotificationsList not implemented")
}

func RegisterNotificationServiceServer(s *grpc.Server, srv NotificationServiceServer) {
	s.RegisterService(&_NotificationService_serviceDesc, srv)
}

func _NotificationService_NotificationCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).NotificationCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.NotificationService/NotificationCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).NotificationCreate(ctx, req.(*NotificationCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_NotificationRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).NotificationRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.NotificationService/NotificationRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).NotificationRead(ctx, req.(*NotificationReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_NotificationUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).NotificationUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.NotificationService/NotificationUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).NotificationUpdate(ctx, req.(*NotificationUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_NotificationDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).NotificationDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.NotificationService/NotificationDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).NotificationDelete(ctx, req.(*NotificationDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_NotificationsList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NotificationsListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NotificationServiceServer).NotificationsList(m, &notificationServiceNotificationsListServer{stream})
}

type NotificationService_NotificationsListServer interface {
	Send(*NotificationsListResponse) error
	grpc.ServerStream
}

type notificationServiceNotificationsListServer struct {
	grpc.ServerStream
}

func (x *notificationServiceNotificationsListServer) Send(m *NotificationsListResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _NotificationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "notifications.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NotificationCreate",
			Handler:    _NotificationService_NotificationCreate_Handler,
		},
		{
			MethodName: "NotificationRead",
			Handler:    _NotificationService_NotificationRead_Handler,
		},
		{
			MethodName: "NotificationUpdate",
			Handler:    _NotificationService_NotificationUpdate_Handler,
		},
		{
			MethodName: "NotificationDelete",
			Handler:    _NotificationService_NotificationDelete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "NotificationsList",
			Handler:       _NotificationService_NotificationsList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "notification.proto",
}