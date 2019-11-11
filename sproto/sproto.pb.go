// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sproto.proto

package sproto

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type Address struct {
	Address1             string   `protobuf:"bytes,1,opt,name=Address1,proto3" json:"Address1,omitempty"`
	Address2             string   `protobuf:"bytes,2,opt,name=Address2,proto3" json:"Address2,omitempty"`
	City                 string   `protobuf:"bytes,3,opt,name=City,proto3" json:"City,omitempty"`
	State                string   `protobuf:"bytes,4,opt,name=State,proto3" json:"State,omitempty"`
	Country              string   `protobuf:"bytes,5,opt,name=Country,proto3" json:"Country,omitempty"`
	ZipCode              string   `protobuf:"bytes,6,opt,name=ZipCode,proto3" json:"ZipCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9dd1c2b06379676, []int{0}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetAddress1() string {
	if m != nil {
		return m.Address1
	}
	return ""
}

func (m *Address) GetAddress2() string {
	if m != nil {
		return m.Address2
	}
	return ""
}

func (m *Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Address) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Address) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Address) GetZipCode() string {
	if m != nil {
		return m.ZipCode
	}
	return ""
}

type MsgCodeResult struct {
	MsgCode              string   `protobuf:"bytes,1,opt,name=MsgCode,proto3" json:"MsgCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *MsgCodeResult) Reset()         { *m = MsgCodeResult{} }
func (m *MsgCodeResult) String() string { return proto.CompactTextString(m) }
func (*MsgCodeResult) ProtoMessage()    {}
func (*MsgCodeResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9dd1c2b06379676, []int{1}
}

func (m *MsgCodeResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgCodeResult.Unmarshal(m, b)
}
func (m *MsgCodeResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgCodeResult.Marshal(b, m, deterministic)
}
func (m *MsgCodeResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCodeResult.Merge(m, src)
}
func (m *MsgCodeResult) XXX_Size() int {
	return xxx_messageInfo_MsgCodeResult.Size(m)
}
func (m *MsgCodeResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCodeResult.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCodeResult proto.InternalMessageInfo

func (m *MsgCodeResult) GetMsgCode() string {
	if m != nil {
		return m.MsgCode
	}
	return ""
}

type RouteDTO struct {
	ID                   string         `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Permission_ID        string         `protobuf:"bytes,2,opt,name=Permission_ID,json=PermissionID,proto3" json:"Permission_ID,omitempty"`
	Area                 string         `protobuf:"bytes,3,opt,name=Area,proto3" json:"Area,omitempty"`
	Controller           string         `protobuf:"bytes,4,opt,name=Controller,proto3" json:"Controller,omitempty"`
	Action               string         `protobuf:"bytes,5,opt,name=Action,proto3" json:"Action,omitempty"`
	Permission           *PermissionDTO `protobuf:"bytes,6,opt,name=Permission,proto3" json:"Permission,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-" bson:"-"`
	XXX_unrecognized     []byte         `json:"-" bson:"-"`
	XXX_sizecache        int32          `json:"-" bson:"-"`
}

func (m *RouteDTO) Reset()         { *m = RouteDTO{} }
func (m *RouteDTO) String() string { return proto.CompactTextString(m) }
func (*RouteDTO) ProtoMessage()    {}
func (*RouteDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9dd1c2b06379676, []int{2}
}

func (m *RouteDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteDTO.Unmarshal(m, b)
}
func (m *RouteDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteDTO.Marshal(b, m, deterministic)
}
func (m *RouteDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteDTO.Merge(m, src)
}
func (m *RouteDTO) XXX_Size() int {
	return xxx_messageInfo_RouteDTO.Size(m)
}
func (m *RouteDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteDTO.DiscardUnknown(m)
}

var xxx_messageInfo_RouteDTO proto.InternalMessageInfo

func (m *RouteDTO) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *RouteDTO) GetPermission_ID() string {
	if m != nil {
		return m.Permission_ID
	}
	return ""
}

func (m *RouteDTO) GetArea() string {
	if m != nil {
		return m.Area
	}
	return ""
}

func (m *RouteDTO) GetController() string {
	if m != nil {
		return m.Controller
	}
	return ""
}

func (m *RouteDTO) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *RouteDTO) GetPermission() *PermissionDTO {
	if m != nil {
		return m.Permission
	}
	return nil
}

type PermissionDTO struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	IsAllowGuest         bool     `protobuf:"varint,3,opt,name=IsAllowGuest,proto3" json:"IsAllowGuest,omitempty"`
	IsAllowAnyUser       bool     `protobuf:"varint,4,opt,name=IsAllowAnyUser,proto3" json:"IsAllowAnyUser,omitempty"`
	AllowedRoles         int64    `protobuf:"varint,5,opt,name=AllowedRoles,proto3" json:"AllowedRoles,omitempty"`
	Level                int32    `protobuf:"varint,6,opt,name=Level,proto3" json:"Level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *PermissionDTO) Reset()         { *m = PermissionDTO{} }
func (m *PermissionDTO) String() string { return proto.CompactTextString(m) }
func (*PermissionDTO) ProtoMessage()    {}
func (*PermissionDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9dd1c2b06379676, []int{3}
}

func (m *PermissionDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PermissionDTO.Unmarshal(m, b)
}
func (m *PermissionDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PermissionDTO.Marshal(b, m, deterministic)
}
func (m *PermissionDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PermissionDTO.Merge(m, src)
}
func (m *PermissionDTO) XXX_Size() int {
	return xxx_messageInfo_PermissionDTO.Size(m)
}
func (m *PermissionDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_PermissionDTO.DiscardUnknown(m)
}

var xxx_messageInfo_PermissionDTO proto.InternalMessageInfo

func (m *PermissionDTO) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *PermissionDTO) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PermissionDTO) GetIsAllowGuest() bool {
	if m != nil {
		return m.IsAllowGuest
	}
	return false
}

func (m *PermissionDTO) GetIsAllowAnyUser() bool {
	if m != nil {
		return m.IsAllowAnyUser
	}
	return false
}

func (m *PermissionDTO) GetAllowedRoles() int64 {
	if m != nil {
		return m.AllowedRoles
	}
	return 0
}

func (m *PermissionDTO) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func init() {
	proto.RegisterType((*Address)(nil), "sproto.Address")
	proto.RegisterType((*MsgCodeResult)(nil), "sproto.MsgCodeResult")
	proto.RegisterType((*RouteDTO)(nil), "sproto.RouteDTO")
	proto.RegisterType((*PermissionDTO)(nil), "sproto.PermissionDTO")
}

func init() { proto.RegisterFile("sproto.proto", fileDescriptor_c9dd1c2b06379676) }

var fileDescriptor_c9dd1c2b06379676 = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4d, 0x8b, 0xd3, 0x40,
	0x18, 0xc7, 0x49, 0xb6, 0xcd, 0xc6, 0xc7, 0xee, 0x1e, 0x06, 0x95, 0xe0, 0x41, 0x97, 0x08, 0xa2,
	0x97, 0x14, 0x2b, 0x7e, 0x80, 0x98, 0xc0, 0x12, 0x7c, 0xd9, 0x65, 0x5a, 0x11, 0xf6, 0x22, 0x6d,
	0xf2, 0xb4, 0x06, 0xa6, 0x33, 0x65, 0x5e, 0x2c, 0xf9, 0x34, 0xde, 0xfd, 0x06, 0x5e, 0xfd, 0x54,
	0x1e, 0x65, 0x26, 0x53, 0x9b, 0x8a, 0x97, 0xe1, 0xf9, 0xff, 0xfe, 0x0f, 0xc3, 0xfc, 0x18, 0x98,
	0xa8, 0x9d, 0x14, 0x5a, 0x64, 0xee, 0x24, 0x51, 0x9f, 0xd2, 0xef, 0x01, 0x9c, 0xe7, 0x4d, 0x23,
	0x51, 0x29, 0xf2, 0x18, 0x62, 0x3f, 0xbe, 0x4a, 0x82, 0xab, 0xe0, 0xc5, 0x3d, 0xfa, 0x37, 0x0f,
	0xba, 0x59, 0x12, 0x9e, 0x74, 0x33, 0x42, 0x60, 0x54, 0xb4, 0xba, 0x4b, 0xce, 0x1c, 0x77, 0x33,
	0x79, 0x00, 0xe3, 0xb9, 0x5e, 0x6a, 0x4c, 0x46, 0x0e, 0xf6, 0x81, 0x24, 0x70, 0x5e, 0x08, 0xc3,
	0xb5, 0xec, 0x92, 0xb1, 0xe3, 0x87, 0x68, 0x9b, 0xbb, 0x76, 0x57, 0x88, 0x06, 0x93, 0xa8, 0x6f,
	0x7c, 0x4c, 0x5f, 0xc2, 0xc5, 0x07, 0xb5, 0xb1, 0x23, 0x45, 0x65, 0x98, 0xb6, 0xab, 0x1e, 0xf8,
	0x57, 0x1e, 0x62, 0xfa, 0x2b, 0x80, 0x98, 0x0a, 0xa3, 0xb1, 0x5c, 0xdc, 0x90, 0x4b, 0x08, 0xab,
	0xd2, 0x6f, 0x84, 0x55, 0x49, 0x9e, 0xc1, 0xc5, 0x2d, 0xca, 0x6d, 0xab, 0x54, 0x2b, 0xf8, 0x97,
	0xaa, 0xf4, 0x1a, 0x93, 0x23, 0xac, 0x4a, 0xab, 0x92, 0x4b, 0x5c, 0x1e, 0x54, 0xec, 0x4c, 0x9e,
	0x00, 0x14, 0x82, 0x6b, 0x29, 0x18, 0x43, 0xe9, 0x7d, 0x06, 0x84, 0x3c, 0x82, 0x28, 0xaf, 0x75,
	0x2b, 0xb8, 0x77, 0xf2, 0x89, 0xbc, 0x01, 0x38, 0xde, 0xed, 0xac, 0xee, 0xcf, 0x1e, 0x66, 0xfe,
	0x17, 0x8e, 0x4d, 0xb9, 0xb8, 0xa1, 0x83, 0xc5, 0xf4, 0x67, 0x30, 0x7c, 0xe8, 0xff, 0x4c, 0x08,
	0x8c, 0x3e, 0x2e, 0xb7, 0xe8, 0x05, 0xdc, 0x4c, 0x52, 0x98, 0x54, 0x2a, 0x67, 0x4c, 0xec, 0xaf,
	0x0d, 0x2a, 0xed, 0x04, 0x62, 0x7a, 0xc2, 0xc8, 0x73, 0xb8, 0xf4, 0x39, 0xe7, 0xdd, 0x27, 0xe5,
	0x65, 0x62, 0xfa, 0x0f, 0xb5, 0x77, 0xb9, 0x8c, 0x0d, 0x15, 0x0c, 0x95, 0xd3, 0x3a, 0xa3, 0x27,
	0xcc, 0xfe, 0xef, 0x7b, 0xfc, 0x86, 0xcc, 0x79, 0x8d, 0x69, 0x1f, 0xde, 0x5e, 0xdf, 0x3d, 0xdd,
	0xb4, 0xfa, 0xab, 0x59, 0x65, 0xb5, 0xd8, 0x4e, 0x55, 0xc7, 0xeb, 0xb5, 0xd1, 0x46, 0xe2, 0x74,
	0x23, 0xa6, 0xbd, 0xf8, 0xef, 0x20, 0xf8, 0x11, 0x5e, 0xcd, 0x3b, 0x5e, 0xcf, 0xc5, 0x5a, 0x67,
	0xb7, 0x96, 0xad, 0xcc, 0x3a, 0xfb, 0x8c, 0x8c, 0xbd, 0xe3, 0x62, 0xcf, 0x17, 0xdd, 0x0e, 0xd5,
	0x2a, 0x72, 0xcb, 0xaf, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xc9, 0x86, 0x8a, 0xb5, 0x02,
	0x00, 0x00,
}
