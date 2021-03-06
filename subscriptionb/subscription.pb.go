// Code generated by protoc-gen-go. DO NOT EDIT.
// source: subscription.proto

package subscriptionb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Channel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Channel) Reset()         { *m = Channel{} }
func (m *Channel) String() string { return proto.CompactTextString(m) }
func (*Channel) ProtoMessage()    {}
func (*Channel) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f8ad1a64b2bad6, []int{0}
}

func (m *Channel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Channel.Unmarshal(m, b)
}
func (m *Channel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Channel.Marshal(b, m, deterministic)
}
func (m *Channel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Channel.Merge(m, src)
}
func (m *Channel) XXX_Size() int {
	return xxx_messageInfo_Channel.Size(m)
}
func (m *Channel) XXX_DiscardUnknown() {
	xxx_messageInfo_Channel.DiscardUnknown(m)
}

var xxx_messageInfo_Channel proto.InternalMessageInfo

func (m *Channel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f8ad1a64b2bad6, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Subscription struct {
	Channel              *Channel `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Subscription) Reset()         { *m = Subscription{} }
func (m *Subscription) String() string { return proto.CompactTextString(m) }
func (*Subscription) ProtoMessage()    {}
func (*Subscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f8ad1a64b2bad6, []int{2}
}

func (m *Subscription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subscription.Unmarshal(m, b)
}
func (m *Subscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subscription.Marshal(b, m, deterministic)
}
func (m *Subscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subscription.Merge(m, src)
}
func (m *Subscription) XXX_Size() int {
	return xxx_messageInfo_Subscription.Size(m)
}
func (m *Subscription) XXX_DiscardUnknown() {
	xxx_messageInfo_Subscription.DiscardUnknown(m)
}

var xxx_messageInfo_Subscription proto.InternalMessageInfo

func (m *Subscription) GetChannel() *Channel {
	if m != nil {
		return m.Channel
	}
	return nil
}

func (m *Subscription) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type Status struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f8ad1a64b2bad6, []int{3}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ReadSubscriptionReq struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadSubscriptionReq) Reset()         { *m = ReadSubscriptionReq{} }
func (m *ReadSubscriptionReq) String() string { return proto.CompactTextString(m) }
func (*ReadSubscriptionReq) ProtoMessage()    {}
func (*ReadSubscriptionReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f8ad1a64b2bad6, []int{4}
}

func (m *ReadSubscriptionReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadSubscriptionReq.Unmarshal(m, b)
}
func (m *ReadSubscriptionReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadSubscriptionReq.Marshal(b, m, deterministic)
}
func (m *ReadSubscriptionReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadSubscriptionReq.Merge(m, src)
}
func (m *ReadSubscriptionReq) XXX_Size() int {
	return xxx_messageInfo_ReadSubscriptionReq.Size(m)
}
func (m *ReadSubscriptionReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadSubscriptionReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReadSubscriptionReq proto.InternalMessageInfo

func (m *ReadSubscriptionReq) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type ReadSubscriptionRes struct {
	Subscription         *Subscription `protobuf:"bytes,1,opt,name=Subscription,proto3" json:"Subscription,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReadSubscriptionRes) Reset()         { *m = ReadSubscriptionRes{} }
func (m *ReadSubscriptionRes) String() string { return proto.CompactTextString(m) }
func (*ReadSubscriptionRes) ProtoMessage()    {}
func (*ReadSubscriptionRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f8ad1a64b2bad6, []int{5}
}

func (m *ReadSubscriptionRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadSubscriptionRes.Unmarshal(m, b)
}
func (m *ReadSubscriptionRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadSubscriptionRes.Marshal(b, m, deterministic)
}
func (m *ReadSubscriptionRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadSubscriptionRes.Merge(m, src)
}
func (m *ReadSubscriptionRes) XXX_Size() int {
	return xxx_messageInfo_ReadSubscriptionRes.Size(m)
}
func (m *ReadSubscriptionRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadSubscriptionRes.DiscardUnknown(m)
}

var xxx_messageInfo_ReadSubscriptionRes proto.InternalMessageInfo

func (m *ReadSubscriptionRes) GetSubscription() *Subscription {
	if m != nil {
		return m.Subscription
	}
	return nil
}

type AddSubscriptionReq struct {
	Subscription         *Subscription `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AddSubscriptionReq) Reset()         { *m = AddSubscriptionReq{} }
func (m *AddSubscriptionReq) String() string { return proto.CompactTextString(m) }
func (*AddSubscriptionReq) ProtoMessage()    {}
func (*AddSubscriptionReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f8ad1a64b2bad6, []int{6}
}

func (m *AddSubscriptionReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddSubscriptionReq.Unmarshal(m, b)
}
func (m *AddSubscriptionReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddSubscriptionReq.Marshal(b, m, deterministic)
}
func (m *AddSubscriptionReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddSubscriptionReq.Merge(m, src)
}
func (m *AddSubscriptionReq) XXX_Size() int {
	return xxx_messageInfo_AddSubscriptionReq.Size(m)
}
func (m *AddSubscriptionReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddSubscriptionReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddSubscriptionReq proto.InternalMessageInfo

func (m *AddSubscriptionReq) GetSubscription() *Subscription {
	if m != nil {
		return m.Subscription
	}
	return nil
}

type AddSubscriptionRes struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddSubscriptionRes) Reset()         { *m = AddSubscriptionRes{} }
func (m *AddSubscriptionRes) String() string { return proto.CompactTextString(m) }
func (*AddSubscriptionRes) ProtoMessage()    {}
func (*AddSubscriptionRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f8ad1a64b2bad6, []int{7}
}

func (m *AddSubscriptionRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddSubscriptionRes.Unmarshal(m, b)
}
func (m *AddSubscriptionRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddSubscriptionRes.Marshal(b, m, deterministic)
}
func (m *AddSubscriptionRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddSubscriptionRes.Merge(m, src)
}
func (m *AddSubscriptionRes) XXX_Size() int {
	return xxx_messageInfo_AddSubscriptionRes.Size(m)
}
func (m *AddSubscriptionRes) XXX_DiscardUnknown() {
	xxx_messageInfo_AddSubscriptionRes.DiscardUnknown(m)
}

var xxx_messageInfo_AddSubscriptionRes proto.InternalMessageInfo

func (m *AddSubscriptionRes) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type DeleteSubscriptionReq struct {
	Subscription         *Subscription `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DeleteSubscriptionReq) Reset()         { *m = DeleteSubscriptionReq{} }
func (m *DeleteSubscriptionReq) String() string { return proto.CompactTextString(m) }
func (*DeleteSubscriptionReq) ProtoMessage()    {}
func (*DeleteSubscriptionReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f8ad1a64b2bad6, []int{8}
}

func (m *DeleteSubscriptionReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSubscriptionReq.Unmarshal(m, b)
}
func (m *DeleteSubscriptionReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSubscriptionReq.Marshal(b, m, deterministic)
}
func (m *DeleteSubscriptionReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSubscriptionReq.Merge(m, src)
}
func (m *DeleteSubscriptionReq) XXX_Size() int {
	return xxx_messageInfo_DeleteSubscriptionReq.Size(m)
}
func (m *DeleteSubscriptionReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSubscriptionReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSubscriptionReq proto.InternalMessageInfo

func (m *DeleteSubscriptionReq) GetSubscription() *Subscription {
	if m != nil {
		return m.Subscription
	}
	return nil
}

type DeleteSubscriptionRes struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteSubscriptionRes) Reset()         { *m = DeleteSubscriptionRes{} }
func (m *DeleteSubscriptionRes) String() string { return proto.CompactTextString(m) }
func (*DeleteSubscriptionRes) ProtoMessage()    {}
func (*DeleteSubscriptionRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f8ad1a64b2bad6, []int{9}
}

func (m *DeleteSubscriptionRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSubscriptionRes.Unmarshal(m, b)
}
func (m *DeleteSubscriptionRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSubscriptionRes.Marshal(b, m, deterministic)
}
func (m *DeleteSubscriptionRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSubscriptionRes.Merge(m, src)
}
func (m *DeleteSubscriptionRes) XXX_Size() int {
	return xxx_messageInfo_DeleteSubscriptionRes.Size(m)
}
func (m *DeleteSubscriptionRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSubscriptionRes.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSubscriptionRes proto.InternalMessageInfo

func (m *DeleteSubscriptionRes) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type ListSubscriptionsReq struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListSubscriptionsReq) Reset()         { *m = ListSubscriptionsReq{} }
func (m *ListSubscriptionsReq) String() string { return proto.CompactTextString(m) }
func (*ListSubscriptionsReq) ProtoMessage()    {}
func (*ListSubscriptionsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f8ad1a64b2bad6, []int{10}
}

func (m *ListSubscriptionsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSubscriptionsReq.Unmarshal(m, b)
}
func (m *ListSubscriptionsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSubscriptionsReq.Marshal(b, m, deterministic)
}
func (m *ListSubscriptionsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSubscriptionsReq.Merge(m, src)
}
func (m *ListSubscriptionsReq) XXX_Size() int {
	return xxx_messageInfo_ListSubscriptionsReq.Size(m)
}
func (m *ListSubscriptionsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSubscriptionsReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListSubscriptionsReq proto.InternalMessageInfo

func (m *ListSubscriptionsReq) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type ListSubscriptionsRes struct {
	Subscription         *Subscription `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListSubscriptionsRes) Reset()         { *m = ListSubscriptionsRes{} }
func (m *ListSubscriptionsRes) String() string { return proto.CompactTextString(m) }
func (*ListSubscriptionsRes) ProtoMessage()    {}
func (*ListSubscriptionsRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f8ad1a64b2bad6, []int{11}
}

func (m *ListSubscriptionsRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSubscriptionsRes.Unmarshal(m, b)
}
func (m *ListSubscriptionsRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSubscriptionsRes.Marshal(b, m, deterministic)
}
func (m *ListSubscriptionsRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSubscriptionsRes.Merge(m, src)
}
func (m *ListSubscriptionsRes) XXX_Size() int {
	return xxx_messageInfo_ListSubscriptionsRes.Size(m)
}
func (m *ListSubscriptionsRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSubscriptionsRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListSubscriptionsRes proto.InternalMessageInfo

func (m *ListSubscriptionsRes) GetSubscription() *Subscription {
	if m != nil {
		return m.Subscription
	}
	return nil
}

func init() {
	proto.RegisterType((*Channel)(nil), "subscription.Channel")
	proto.RegisterType((*User)(nil), "subscription.User")
	proto.RegisterType((*Subscription)(nil), "subscription.Subscription")
	proto.RegisterType((*Status)(nil), "subscription.Status")
	proto.RegisterType((*ReadSubscriptionReq)(nil), "subscription.ReadSubscriptionReq")
	proto.RegisterType((*ReadSubscriptionRes)(nil), "subscription.ReadSubscriptionRes")
	proto.RegisterType((*AddSubscriptionReq)(nil), "subscription.AddSubscriptionReq")
	proto.RegisterType((*AddSubscriptionRes)(nil), "subscription.AddSubscriptionRes")
	proto.RegisterType((*DeleteSubscriptionReq)(nil), "subscription.DeleteSubscriptionReq")
	proto.RegisterType((*DeleteSubscriptionRes)(nil), "subscription.DeleteSubscriptionRes")
	proto.RegisterType((*ListSubscriptionsReq)(nil), "subscription.ListSubscriptionsReq")
	proto.RegisterType((*ListSubscriptionsRes)(nil), "subscription.ListSubscriptionsRes")
}

func init() { proto.RegisterFile("subscription.proto", fileDescriptor_c4f8ad1a64b2bad6) }

var fileDescriptor_c4f8ad1a64b2bad6 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x41, 0x4f, 0xc2, 0x40,
	0x10, 0x85, 0x29, 0x12, 0xd0, 0x11, 0x45, 0x07, 0x30, 0xd8, 0x13, 0xae, 0x89, 0xf1, 0x60, 0xd0,
	0xe0, 0x59, 0x12, 0x51, 0x6f, 0x9e, 0x16, 0x91, 0x84, 0x8b, 0x42, 0xbb, 0xd1, 0x26, 0x04, 0xb0,
	0xd3, 0xfa, 0x2f, 0xfc, 0xcf, 0x86, 0xb5, 0x98, 0x1d, 0xba, 0x49, 0x21, 0x7a, 0xec, 0xce, 0x9b,
	0xef, 0xcd, 0xee, 0x9b, 0x14, 0x90, 0xe2, 0x31, 0x79, 0x61, 0x30, 0x8f, 0x82, 0xd9, 0xb4, 0x35,
	0x0f, 0x67, 0xd1, 0x0c, 0xcb, 0xe6, 0x99, 0x38, 0x86, 0xd2, 0xdd, 0xfb, 0x68, 0x3a, 0x55, 0x13,
	0xdc, 0x87, 0x7c, 0xe0, 0x37, 0x9c, 0xa6, 0x73, 0xbe, 0x23, 0xf3, 0x81, 0x2f, 0x8e, 0xa0, 0xd0,
	0x27, 0x15, 0xa6, 0xce, 0xdf, 0xa0, 0xdc, 0x33, 0x10, 0x78, 0x09, 0x25, 0xef, 0x07, 0xa1, 0x45,
	0xbb, 0xed, 0x7a, 0x8b, 0xd9, 0x26, 0x7c, 0xb9, 0x54, 0xe1, 0x19, 0x14, 0x62, 0x52, 0x61, 0x23,
	0xaf, 0xd5, 0xc8, 0xd5, 0x0b, 0x4b, 0xa9, 0xeb, 0x42, 0x40, 0xb1, 0x17, 0x8d, 0xa2, 0x98, 0xb0,
	0x01, 0x25, 0x8a, 0x3d, 0x4f, 0x11, 0x69, 0x8b, 0x6d, 0xb9, 0xfc, 0x14, 0x37, 0x50, 0x95, 0x6a,
	0xe4, 0x9b, 0x03, 0x49, 0xf5, 0xf1, 0x6b, 0xe1, 0x64, 0x58, 0xf4, 0x6d, 0xed, 0x84, 0x1d, 0x7e,
	0xc5, 0x04, 0xe3, 0x72, 0x0c, 0x6b, 0x62, 0x7a, 0xf1, 0x04, 0x78, 0xeb, 0xa7, 0x86, 0xea, 0x00,
	0x7b, 0xfb, 0x75, 0xa8, 0x2c, 0xab, 0xae, 0x85, 0x4a, 0x78, 0x01, 0x45, 0xd2, 0xaf, 0x94, 0xf0,
	0x6a, 0x2b, 0x3c, 0x5d, 0x93, 0x89, 0x46, 0x0c, 0xa0, 0x7e, 0xaf, 0x26, 0x2a, 0x52, 0xff, 0x3d,
	0xdc, 0x83, 0x1d, 0xbc, 0xe9, 0x7c, 0x1d, 0xa8, 0x3d, 0x06, 0x14, 0x99, 0x10, 0xda, 0x24, 0xd0,
	0x67, 0x6b, 0x3f, 0xfd, 0xf5, 0x7a, 0xed, 0xaf, 0x2d, 0xa8, 0x9a, 0xe5, 0x9e, 0x0a, 0x3f, 0x03,
	0x4f, 0xe1, 0x10, 0x0e, 0x56, 0x17, 0x08, 0x4f, 0x38, 0xd5, 0xb2, 0x9f, 0x6e, 0xa6, 0x84, 0x44,
	0x0e, 0x07, 0x50, 0x59, 0xc9, 0x1b, 0x9b, 0xbc, 0x2f, 0xbd, 0x64, 0x6e, 0x96, 0x62, 0x01, 0x7e,
	0x05, 0x4c, 0x67, 0x85, 0xa7, 0xbc, 0xd3, 0xba, 0x26, 0xee, 0x1a, 0xa2, 0x85, 0xc3, 0x0b, 0x1c,
	0xa6, 0x62, 0x40, 0xc1, 0x7b, 0x6d, 0x39, 0xbb, 0xd9, 0x1a, 0x12, 0xb9, 0x2b, 0xa7, 0x5b, 0x19,
	0xee, 0x99, 0xc2, 0xf1, 0xb8, 0xa8, 0xff, 0x6e, 0xd7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0c,
	0x72, 0xb9, 0x0e, 0xf3, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SubscriptionServiceClient is the client API for SubscriptionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SubscriptionServiceClient interface {
	ReadSubscription(ctx context.Context, in *ReadSubscriptionReq, opts ...grpc.CallOption) (*ReadSubscriptionRes, error)
	AddSubscription(ctx context.Context, in *AddSubscriptionReq, opts ...grpc.CallOption) (*AddSubscriptionRes, error)
	DeleteSubscription(ctx context.Context, in *DeleteSubscriptionReq, opts ...grpc.CallOption) (*DeleteSubscriptionRes, error)
	ListSubscriptions(ctx context.Context, in *ListSubscriptionsReq, opts ...grpc.CallOption) (SubscriptionService_ListSubscriptionsClient, error)
}

type subscriptionServiceClient struct {
	cc *grpc.ClientConn
}

func NewSubscriptionServiceClient(cc *grpc.ClientConn) SubscriptionServiceClient {
	return &subscriptionServiceClient{cc}
}

func (c *subscriptionServiceClient) ReadSubscription(ctx context.Context, in *ReadSubscriptionReq, opts ...grpc.CallOption) (*ReadSubscriptionRes, error) {
	out := new(ReadSubscriptionRes)
	err := c.cc.Invoke(ctx, "/subscription.SubscriptionService/ReadSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) AddSubscription(ctx context.Context, in *AddSubscriptionReq, opts ...grpc.CallOption) (*AddSubscriptionRes, error) {
	out := new(AddSubscriptionRes)
	err := c.cc.Invoke(ctx, "/subscription.SubscriptionService/AddSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) DeleteSubscription(ctx context.Context, in *DeleteSubscriptionReq, opts ...grpc.CallOption) (*DeleteSubscriptionRes, error) {
	out := new(DeleteSubscriptionRes)
	err := c.cc.Invoke(ctx, "/subscription.SubscriptionService/DeleteSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) ListSubscriptions(ctx context.Context, in *ListSubscriptionsReq, opts ...grpc.CallOption) (SubscriptionService_ListSubscriptionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SubscriptionService_serviceDesc.Streams[0], "/subscription.SubscriptionService/ListSubscriptions", opts...)
	if err != nil {
		return nil, err
	}
	x := &subscriptionServiceListSubscriptionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SubscriptionService_ListSubscriptionsClient interface {
	Recv() (*ListSubscriptionsRes, error)
	grpc.ClientStream
}

type subscriptionServiceListSubscriptionsClient struct {
	grpc.ClientStream
}

func (x *subscriptionServiceListSubscriptionsClient) Recv() (*ListSubscriptionsRes, error) {
	m := new(ListSubscriptionsRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SubscriptionServiceServer is the server API for SubscriptionService service.
type SubscriptionServiceServer interface {
	ReadSubscription(context.Context, *ReadSubscriptionReq) (*ReadSubscriptionRes, error)
	AddSubscription(context.Context, *AddSubscriptionReq) (*AddSubscriptionRes, error)
	DeleteSubscription(context.Context, *DeleteSubscriptionReq) (*DeleteSubscriptionRes, error)
	ListSubscriptions(*ListSubscriptionsReq, SubscriptionService_ListSubscriptionsServer) error
}

// UnimplementedSubscriptionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSubscriptionServiceServer struct {
}

func (*UnimplementedSubscriptionServiceServer) ReadSubscription(ctx context.Context, req *ReadSubscriptionReq) (*ReadSubscriptionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadSubscription not implemented")
}
func (*UnimplementedSubscriptionServiceServer) AddSubscription(ctx context.Context, req *AddSubscriptionReq) (*AddSubscriptionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSubscription not implemented")
}
func (*UnimplementedSubscriptionServiceServer) DeleteSubscription(ctx context.Context, req *DeleteSubscriptionReq) (*DeleteSubscriptionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSubscription not implemented")
}
func (*UnimplementedSubscriptionServiceServer) ListSubscriptions(req *ListSubscriptionsReq, srv SubscriptionService_ListSubscriptionsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListSubscriptions not implemented")
}

func RegisterSubscriptionServiceServer(s *grpc.Server, srv SubscriptionServiceServer) {
	s.RegisterService(&_SubscriptionService_serviceDesc, srv)
}

func _SubscriptionService_ReadSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadSubscriptionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).ReadSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subscription.SubscriptionService/ReadSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).ReadSubscription(ctx, req.(*ReadSubscriptionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_AddSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSubscriptionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).AddSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subscription.SubscriptionService/AddSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).AddSubscription(ctx, req.(*AddSubscriptionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_DeleteSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSubscriptionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).DeleteSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subscription.SubscriptionService/DeleteSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).DeleteSubscription(ctx, req.(*DeleteSubscriptionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_ListSubscriptions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListSubscriptionsReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SubscriptionServiceServer).ListSubscriptions(m, &subscriptionServiceListSubscriptionsServer{stream})
}

type SubscriptionService_ListSubscriptionsServer interface {
	Send(*ListSubscriptionsRes) error
	grpc.ServerStream
}

type subscriptionServiceListSubscriptionsServer struct {
	grpc.ServerStream
}

func (x *subscriptionServiceListSubscriptionsServer) Send(m *ListSubscriptionsRes) error {
	return x.ServerStream.SendMsg(m)
}

var _SubscriptionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "subscription.SubscriptionService",
	HandlerType: (*SubscriptionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadSubscription",
			Handler:    _SubscriptionService_ReadSubscription_Handler,
		},
		{
			MethodName: "AddSubscription",
			Handler:    _SubscriptionService_AddSubscription_Handler,
		},
		{
			MethodName: "DeleteSubscription",
			Handler:    _SubscriptionService_DeleteSubscription_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListSubscriptions",
			Handler:       _SubscriptionService_ListSubscriptions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "subscription.proto",
}
