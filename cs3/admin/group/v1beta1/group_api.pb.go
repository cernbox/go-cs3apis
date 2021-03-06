// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/admin/group/v1beta1/group_api.proto

package groupv1beta1

import (
	context "context"
	fmt "fmt"
	v1beta11 "github.com/cs3org/go-cs3apis/cs3/identity/group/v1beta1"
	v1beta13 "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	v1beta12 "github.com/cs3org/go-cs3apis/cs3/rpc/v1beta1"
	v1beta1 "github.com/cs3org/go-cs3apis/cs3/types/v1beta1"
	proto "github.com/golang/protobuf/proto"
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

type CreateGroupRequest struct {
	// OPTIONAL.
	// Opaque information. Allow to send any arbitrary data a service might use that is outside the API boundaries
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The information of group to be created.
	Group                *v1beta11.Group `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CreateGroupRequest) Reset()         { *m = CreateGroupRequest{} }
func (m *CreateGroupRequest) String() string { return proto.CompactTextString(m) }
func (*CreateGroupRequest) ProtoMessage()    {}
func (*CreateGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_019b61d7d70538b7, []int{0}
}

func (m *CreateGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateGroupRequest.Unmarshal(m, b)
}
func (m *CreateGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateGroupRequest.Marshal(b, m, deterministic)
}
func (m *CreateGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateGroupRequest.Merge(m, src)
}
func (m *CreateGroupRequest) XXX_Size() int {
	return xxx_messageInfo_CreateGroupRequest.Size(m)
}
func (m *CreateGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateGroupRequest proto.InternalMessageInfo

func (m *CreateGroupRequest) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *CreateGroupRequest) GetGroup() *v1beta11.Group {
	if m != nil {
		return m.Group
	}
	return nil
}

type CreateGroupResponse struct {
	// REQUIRED.
	// The response status.
	Status *v1beta12.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The group information.
	Group                *v1beta11.Group `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CreateGroupResponse) Reset()         { *m = CreateGroupResponse{} }
func (m *CreateGroupResponse) String() string { return proto.CompactTextString(m) }
func (*CreateGroupResponse) ProtoMessage()    {}
func (*CreateGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_019b61d7d70538b7, []int{1}
}

func (m *CreateGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateGroupResponse.Unmarshal(m, b)
}
func (m *CreateGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateGroupResponse.Marshal(b, m, deterministic)
}
func (m *CreateGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateGroupResponse.Merge(m, src)
}
func (m *CreateGroupResponse) XXX_Size() int {
	return xxx_messageInfo_CreateGroupResponse.Size(m)
}
func (m *CreateGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateGroupResponse proto.InternalMessageInfo

func (m *CreateGroupResponse) GetStatus() *v1beta12.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CreateGroupResponse) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *CreateGroupResponse) GetGroup() *v1beta11.Group {
	if m != nil {
		return m.Group
	}
	return nil
}

type DeleteGroupRequest struct {
	// OPTIONAL.
	// Opaque information. Allow to send any arbitrary data a service might use that is outside the API boundaries.
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The group to be deleted, given their ID.
	GroupId              *v1beta11.GroupId `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DeleteGroupRequest) Reset()         { *m = DeleteGroupRequest{} }
func (m *DeleteGroupRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteGroupRequest) ProtoMessage()    {}
func (*DeleteGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_019b61d7d70538b7, []int{2}
}

func (m *DeleteGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteGroupRequest.Unmarshal(m, b)
}
func (m *DeleteGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteGroupRequest.Marshal(b, m, deterministic)
}
func (m *DeleteGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteGroupRequest.Merge(m, src)
}
func (m *DeleteGroupRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteGroupRequest.Size(m)
}
func (m *DeleteGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteGroupRequest proto.InternalMessageInfo

func (m *DeleteGroupRequest) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *DeleteGroupRequest) GetGroupId() *v1beta11.GroupId {
	if m != nil {
		return m.GroupId
	}
	return nil
}

type DeleteGroupResponse struct {
	// REQUIRED.
	// The response status.
	Status *v1beta12.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque               *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DeleteGroupResponse) Reset()         { *m = DeleteGroupResponse{} }
func (m *DeleteGroupResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteGroupResponse) ProtoMessage()    {}
func (*DeleteGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_019b61d7d70538b7, []int{3}
}

func (m *DeleteGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteGroupResponse.Unmarshal(m, b)
}
func (m *DeleteGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteGroupResponse.Marshal(b, m, deterministic)
}
func (m *DeleteGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteGroupResponse.Merge(m, src)
}
func (m *DeleteGroupResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteGroupResponse.Size(m)
}
func (m *DeleteGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteGroupResponse proto.InternalMessageInfo

func (m *DeleteGroupResponse) GetStatus() *v1beta12.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DeleteGroupResponse) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

type AddUserToGroupRequest struct {
	// REQUIRED.
	// ID of the user to add to the group
	UserId *v1beta13.UserId `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// REQUIRED.
	// ID of the target group.
	GroupId *v1beta11.GroupId `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque               *v1beta1.Opaque `protobuf:"bytes,3,opt,name=opaque,proto3" json:"opaque,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AddUserToGroupRequest) Reset()         { *m = AddUserToGroupRequest{} }
func (m *AddUserToGroupRequest) String() string { return proto.CompactTextString(m) }
func (*AddUserToGroupRequest) ProtoMessage()    {}
func (*AddUserToGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_019b61d7d70538b7, []int{4}
}

func (m *AddUserToGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUserToGroupRequest.Unmarshal(m, b)
}
func (m *AddUserToGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUserToGroupRequest.Marshal(b, m, deterministic)
}
func (m *AddUserToGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUserToGroupRequest.Merge(m, src)
}
func (m *AddUserToGroupRequest) XXX_Size() int {
	return xxx_messageInfo_AddUserToGroupRequest.Size(m)
}
func (m *AddUserToGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUserToGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddUserToGroupRequest proto.InternalMessageInfo

func (m *AddUserToGroupRequest) GetUserId() *v1beta13.UserId {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *AddUserToGroupRequest) GetGroupId() *v1beta11.GroupId {
	if m != nil {
		return m.GroupId
	}
	return nil
}

func (m *AddUserToGroupRequest) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

type AddUserToGroupResponse struct {
	// REQUIRED.
	// The response status.
	Status *v1beta12.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque               *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AddUserToGroupResponse) Reset()         { *m = AddUserToGroupResponse{} }
func (m *AddUserToGroupResponse) String() string { return proto.CompactTextString(m) }
func (*AddUserToGroupResponse) ProtoMessage()    {}
func (*AddUserToGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_019b61d7d70538b7, []int{5}
}

func (m *AddUserToGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUserToGroupResponse.Unmarshal(m, b)
}
func (m *AddUserToGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUserToGroupResponse.Marshal(b, m, deterministic)
}
func (m *AddUserToGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUserToGroupResponse.Merge(m, src)
}
func (m *AddUserToGroupResponse) XXX_Size() int {
	return xxx_messageInfo_AddUserToGroupResponse.Size(m)
}
func (m *AddUserToGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUserToGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddUserToGroupResponse proto.InternalMessageInfo

func (m *AddUserToGroupResponse) GetStatus() *v1beta12.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AddUserToGroupResponse) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

type RemoveUserFromGroupRequest struct {
	// REQUIRED.
	// ID of the user to add to the group
	UserId *v1beta13.UserId `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// REQUIRED.
	// ID of the target group.
	GroupId *v1beta11.GroupId `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque               *v1beta1.Opaque `protobuf:"bytes,3,opt,name=opaque,proto3" json:"opaque,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RemoveUserFromGroupRequest) Reset()         { *m = RemoveUserFromGroupRequest{} }
func (m *RemoveUserFromGroupRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveUserFromGroupRequest) ProtoMessage()    {}
func (*RemoveUserFromGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_019b61d7d70538b7, []int{6}
}

func (m *RemoveUserFromGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveUserFromGroupRequest.Unmarshal(m, b)
}
func (m *RemoveUserFromGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveUserFromGroupRequest.Marshal(b, m, deterministic)
}
func (m *RemoveUserFromGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveUserFromGroupRequest.Merge(m, src)
}
func (m *RemoveUserFromGroupRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveUserFromGroupRequest.Size(m)
}
func (m *RemoveUserFromGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveUserFromGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveUserFromGroupRequest proto.InternalMessageInfo

func (m *RemoveUserFromGroupRequest) GetUserId() *v1beta13.UserId {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *RemoveUserFromGroupRequest) GetGroupId() *v1beta11.GroupId {
	if m != nil {
		return m.GroupId
	}
	return nil
}

func (m *RemoveUserFromGroupRequest) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

type RemoveUserFromGroupResponse struct {
	// REQUIRED.
	// The response status.
	Status *v1beta12.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque               *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RemoveUserFromGroupResponse) Reset()         { *m = RemoveUserFromGroupResponse{} }
func (m *RemoveUserFromGroupResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveUserFromGroupResponse) ProtoMessage()    {}
func (*RemoveUserFromGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_019b61d7d70538b7, []int{7}
}

func (m *RemoveUserFromGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveUserFromGroupResponse.Unmarshal(m, b)
}
func (m *RemoveUserFromGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveUserFromGroupResponse.Marshal(b, m, deterministic)
}
func (m *RemoveUserFromGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveUserFromGroupResponse.Merge(m, src)
}
func (m *RemoveUserFromGroupResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveUserFromGroupResponse.Size(m)
}
func (m *RemoveUserFromGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveUserFromGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveUserFromGroupResponse proto.InternalMessageInfo

func (m *RemoveUserFromGroupResponse) GetStatus() *v1beta12.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *RemoveUserFromGroupResponse) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateGroupRequest)(nil), "cs3.admin.group.v1beta1.CreateGroupRequest")
	proto.RegisterType((*CreateGroupResponse)(nil), "cs3.admin.group.v1beta1.CreateGroupResponse")
	proto.RegisterType((*DeleteGroupRequest)(nil), "cs3.admin.group.v1beta1.DeleteGroupRequest")
	proto.RegisterType((*DeleteGroupResponse)(nil), "cs3.admin.group.v1beta1.DeleteGroupResponse")
	proto.RegisterType((*AddUserToGroupRequest)(nil), "cs3.admin.group.v1beta1.AddUserToGroupRequest")
	proto.RegisterType((*AddUserToGroupResponse)(nil), "cs3.admin.group.v1beta1.AddUserToGroupResponse")
	proto.RegisterType((*RemoveUserFromGroupRequest)(nil), "cs3.admin.group.v1beta1.RemoveUserFromGroupRequest")
	proto.RegisterType((*RemoveUserFromGroupResponse)(nil), "cs3.admin.group.v1beta1.RemoveUserFromGroupResponse")
}

func init() {
	proto.RegisterFile("cs3/admin/group/v1beta1/group_api.proto", fileDescriptor_019b61d7d70538b7)
}

var fileDescriptor_019b61d7d70538b7 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x55, 0xcf, 0x6b, 0x13, 0x41,
	0x14, 0x66, 0x12, 0x4c, 0xcb, 0x8b, 0x7a, 0x98, 0xa0, 0xa9, 0x5b, 0x05, 0x8d, 0x07, 0x7f, 0x32,
	0x21, 0x5d, 0x41, 0xf0, 0x20, 0x24, 0x11, 0x4b, 0x4e, 0x86, 0xf5, 0xc7, 0x41, 0x02, 0xb2, 0xdd,
	0x1d, 0xda, 0x05, 0x93, 0xd9, 0xcc, 0xcc, 0x06, 0x02, 0x1e, 0xea, 0xc9, 0xff, 0xc3, 0xa3, 0x47,
	0xff, 0x08, 0x0f, 0x7a, 0xf5, 0x0f, 0x92, 0x79, 0x33, 0x0d, 0xd9, 0x76, 0xb7, 0x26, 0x14, 0x02,
	0x3d, 0x25, 0xcb, 0xf7, 0xbd, 0xf7, 0xbe, 0xef, 0xdb, 0x99, 0xb7, 0xf0, 0x20, 0x52, 0x7e, 0x3b,
	0x8c, 0xc7, 0xc9, 0xa4, 0x7d, 0x28, 0x45, 0x96, 0xb6, 0x67, 0x9d, 0x03, 0xae, 0xc3, 0x8e, 0x7d,
	0xfa, 0x14, 0xa6, 0x09, 0x4b, 0xa5, 0xd0, 0x82, 0x36, 0x23, 0xe5, 0x33, 0x24, 0x32, 0x84, 0x98,
	0x23, 0x7a, 0x8f, 0x4d, 0x87, 0x24, 0xe6, 0x13, 0x9d, 0xe8, 0xf9, 0xa9, 0x26, 0x92, 0x2b, 0x91,
	0xc9, 0x88, 0x2b, 0xdb, 0xc4, 0x7b, 0x94, 0xe3, 0x66, 0x8a, 0xcb, 0x52, 0xea, 0x6d, 0x43, 0x95,
	0x69, 0xb4, 0x20, 0x28, 0x1d, 0xea, 0xec, 0x04, 0xbd, 0x63, 0x50, 0x3d, 0x4f, 0xb9, 0x5a, 0xe0,
	0xf8, 0x64, 0xe1, 0xd6, 0x31, 0x01, 0xda, 0x97, 0x3c, 0xd4, 0x7c, 0xdf, 0xe8, 0x09, 0xf8, 0x34,
	0xe3, 0x4a, 0xd3, 0x0e, 0xd4, 0x44, 0x1a, 0x4e, 0x33, 0xbe, 0x43, 0xee, 0x92, 0x87, 0xf5, 0xbd,
	0x5b, 0xcc, 0x98, 0xb2, 0x85, 0xae, 0x0d, 0x7b, 0x83, 0x84, 0xc0, 0x11, 0xe9, 0x73, 0xb8, 0x82,
	0x96, 0x76, 0x2a, 0x58, 0x71, 0x0f, 0x2b, 0x4e, 0x1c, 0xe4, 0x93, 0x60, 0x76, 0x96, 0xe5, 0xb7,
	0x7e, 0x12, 0x68, 0xe4, 0x24, 0xa8, 0x54, 0x4c, 0x14, 0xa7, 0x6d, 0xa8, 0x59, 0x27, 0x4e, 0x43,
	0x13, 0x3b, 0xca, 0x34, 0x5a, 0xb4, 0x79, 0x8b, 0x70, 0xe0, 0x68, 0x4b, 0xa2, 0x2b, 0x6b, 0x8b,
	0xae, 0xae, 0x29, 0xfa, 0x1b, 0x01, 0xfa, 0x8a, 0x7f, 0xe6, 0x17, 0xcf, 0xed, 0x25, 0x6c, 0xdb,
	0x13, 0x94, 0xc4, 0x4e, 0xf7, 0xfd, 0xff, 0xaa, 0x18, 0xc4, 0xc1, 0xd6, 0xa1, 0xfd, 0xd3, 0x9a,
	0x43, 0x23, 0x27, 0x64, 0x73, 0xe9, 0xb5, 0x7e, 0x11, 0xb8, 0xd1, 0x8d, 0xe3, 0xf7, 0x8a, 0xcb,
	0x77, 0x22, 0x97, 0xc3, 0x0b, 0xd8, 0x32, 0x67, 0xd6, 0x78, 0x22, 0x45, 0xc9, 0x1a, 0x70, 0xd1,
	0xd5, 0xd4, 0x0f, 0xe2, 0xa0, 0x96, 0xe1, 0xef, 0x45, 0x03, 0x59, 0x32, 0x52, 0x5d, 0xd5, 0xc8,
	0x17, 0xb8, 0x79, 0xda, 0xc7, 0x06, 0x63, 0xfc, 0x43, 0xc0, 0x0b, 0xf8, 0x58, 0xcc, 0xb8, 0x51,
	0xf0, 0x5a, 0x8a, 0xf1, 0x65, 0xce, 0xf2, 0x2b, 0x81, 0xdd, 0x42, 0x37, 0x9b, 0x4b, 0x74, 0xef,
	0x6f, 0x15, 0xb6, 0x71, 0x6a, 0x77, 0x38, 0xa0, 0x47, 0x50, 0x5f, 0x5a, 0x2f, 0xf4, 0x09, 0x2b,
	0xd9, 0xcf, 0xec, 0xec, 0x1e, 0xf4, 0x9e, 0xae, 0x46, 0x76, 0xd6, 0x8e, 0xa0, 0xbe, 0x74, 0x15,
	0xcf, 0x99, 0x74, 0x76, 0x73, 0x9c, 0x33, 0xa9, 0xe8, 0x76, 0x4f, 0xe1, 0x7a, 0xfe, 0xc0, 0x52,
	0x56, 0x5a, 0x5f, 0x78, 0x43, 0xbd, 0xf6, 0xca, 0x7c, 0x37, 0xf2, 0x98, 0x40, 0xa3, 0xe0, 0xbd,
	0x52, 0xbf, 0xb4, 0x51, 0xf9, 0x99, 0xf6, 0x9e, 0xad, 0x57, 0x64, 0x25, 0xf4, 0x66, 0xb0, 0x1b,
	0x89, 0x71, 0x59, 0x69, 0xef, 0x9a, 0x7d, 0xe5, 0x69, 0x32, 0x34, 0x9f, 0xb6, 0x21, 0xf9, 0x78,
	0x15, 0x71, 0x07, 0x7f, 0xaf, 0x54, 0xfb, 0xdd, 0xfd, 0x1f, 0x95, 0x66, 0x5f, 0xf9, 0xac, 0x8b,
	0xe5, 0xc8, 0x67, 0x1f, 0x3a, 0x3d, 0x83, 0xff, 0x46, 0x64, 0x84, 0xc8, 0x08, 0x91, 0x91, 0x43,
	0x0e, 0x6a, 0xf8, 0xad, 0xf4, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x01, 0xd6, 0x73, 0xb8, 0x03,
	0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GroupAPIClient is the client API for GroupAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GroupAPIClient interface {
	// Create a group.
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error)
	// Delete a group.
	DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*DeleteGroupResponse, error)
	// Add a user to a group.
	AddUserToGroup(ctx context.Context, in *AddUserToGroupRequest, opts ...grpc.CallOption) (*AddUserToGroupResponse, error)
	// Remove a user from a group.
	RemoveUserFromGroup(ctx context.Context, in *RemoveUserFromGroupRequest, opts ...grpc.CallOption) (*RemoveUserFromGroupResponse, error)
}

type groupAPIClient struct {
	cc *grpc.ClientConn
}

func NewGroupAPIClient(cc *grpc.ClientConn) GroupAPIClient {
	return &groupAPIClient{cc}
}

func (c *groupAPIClient) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error) {
	out := new(CreateGroupResponse)
	err := c.cc.Invoke(ctx, "/cs3.admin.group.v1beta1.GroupAPI/CreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupAPIClient) DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*DeleteGroupResponse, error) {
	out := new(DeleteGroupResponse)
	err := c.cc.Invoke(ctx, "/cs3.admin.group.v1beta1.GroupAPI/DeleteGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupAPIClient) AddUserToGroup(ctx context.Context, in *AddUserToGroupRequest, opts ...grpc.CallOption) (*AddUserToGroupResponse, error) {
	out := new(AddUserToGroupResponse)
	err := c.cc.Invoke(ctx, "/cs3.admin.group.v1beta1.GroupAPI/AddUserToGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupAPIClient) RemoveUserFromGroup(ctx context.Context, in *RemoveUserFromGroupRequest, opts ...grpc.CallOption) (*RemoveUserFromGroupResponse, error) {
	out := new(RemoveUserFromGroupResponse)
	err := c.cc.Invoke(ctx, "/cs3.admin.group.v1beta1.GroupAPI/RemoveUserFromGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupAPIServer is the server API for GroupAPI service.
type GroupAPIServer interface {
	// Create a group.
	CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error)
	// Delete a group.
	DeleteGroup(context.Context, *DeleteGroupRequest) (*DeleteGroupResponse, error)
	// Add a user to a group.
	AddUserToGroup(context.Context, *AddUserToGroupRequest) (*AddUserToGroupResponse, error)
	// Remove a user from a group.
	RemoveUserFromGroup(context.Context, *RemoveUserFromGroupRequest) (*RemoveUserFromGroupResponse, error)
}

// UnimplementedGroupAPIServer can be embedded to have forward compatible implementations.
type UnimplementedGroupAPIServer struct {
}

func (*UnimplementedGroupAPIServer) CreateGroup(ctx context.Context, req *CreateGroupRequest) (*CreateGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (*UnimplementedGroupAPIServer) DeleteGroup(ctx context.Context, req *DeleteGroupRequest) (*DeleteGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroup not implemented")
}
func (*UnimplementedGroupAPIServer) AddUserToGroup(ctx context.Context, req *AddUserToGroupRequest) (*AddUserToGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserToGroup not implemented")
}
func (*UnimplementedGroupAPIServer) RemoveUserFromGroup(ctx context.Context, req *RemoveUserFromGroupRequest) (*RemoveUserFromGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUserFromGroup not implemented")
}

func RegisterGroupAPIServer(s *grpc.Server, srv GroupAPIServer) {
	s.RegisterService(&_GroupAPI_serviceDesc, srv)
}

func _GroupAPI_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupAPIServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.admin.group.v1beta1.GroupAPI/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupAPIServer).CreateGroup(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupAPI_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupAPIServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.admin.group.v1beta1.GroupAPI/DeleteGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupAPIServer).DeleteGroup(ctx, req.(*DeleteGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupAPI_AddUserToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserToGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupAPIServer).AddUserToGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.admin.group.v1beta1.GroupAPI/AddUserToGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupAPIServer).AddUserToGroup(ctx, req.(*AddUserToGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupAPI_RemoveUserFromGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUserFromGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupAPIServer).RemoveUserFromGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.admin.group.v1beta1.GroupAPI/RemoveUserFromGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupAPIServer).RemoveUserFromGroup(ctx, req.(*RemoveUserFromGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GroupAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cs3.admin.group.v1beta1.GroupAPI",
	HandlerType: (*GroupAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGroup",
			Handler:    _GroupAPI_CreateGroup_Handler,
		},
		{
			MethodName: "DeleteGroup",
			Handler:    _GroupAPI_DeleteGroup_Handler,
		},
		{
			MethodName: "AddUserToGroup",
			Handler:    _GroupAPI_AddUserToGroup_Handler,
		},
		{
			MethodName: "RemoveUserFromGroup",
			Handler:    _GroupAPI_RemoveUserFromGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cs3/admin/group/v1beta1/group_api.proto",
}
