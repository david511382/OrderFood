// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/handler/models/reqs/menu.proto

// 生成的程式在 Golang 中將會屬於 `resp` 套件。

package reqs

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type MenuSelection struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Price                int32    `protobuf:"varint,2,opt,name=Price,proto3" json:"Price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MenuSelection) Reset()         { *m = MenuSelection{} }
func (m *MenuSelection) String() string { return proto.CompactTextString(m) }
func (*MenuSelection) ProtoMessage()    {}
func (*MenuSelection) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec784fbf625bb9ca, []int{0}
}

func (m *MenuSelection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MenuSelection.Unmarshal(m, b)
}
func (m *MenuSelection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MenuSelection.Marshal(b, m, deterministic)
}
func (m *MenuSelection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MenuSelection.Merge(m, src)
}
func (m *MenuSelection) XXX_Size() int {
	return xxx_messageInfo_MenuSelection.Size(m)
}
func (m *MenuSelection) XXX_DiscardUnknown() {
	xxx_messageInfo_MenuSelection.DiscardUnknown(m)
}

var xxx_messageInfo_MenuSelection proto.InternalMessageInfo

func (m *MenuSelection) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MenuSelection) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

type MenuItem struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Price                int32    `protobuf:"varint,2,opt,name=Price,proto3" json:"Price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MenuItem) Reset()         { *m = MenuItem{} }
func (m *MenuItem) String() string { return proto.CompactTextString(m) }
func (*MenuItem) ProtoMessage()    {}
func (*MenuItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec784fbf625bb9ca, []int{1}
}

func (m *MenuItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MenuItem.Unmarshal(m, b)
}
func (m *MenuItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MenuItem.Marshal(b, m, deterministic)
}
func (m *MenuItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MenuItem.Merge(m, src)
}
func (m *MenuItem) XXX_Size() int {
	return xxx_messageInfo_MenuItem.Size(m)
}
func (m *MenuItem) XXX_DiscardUnknown() {
	xxx_messageInfo_MenuItem.DiscardUnknown(m)
}

var xxx_messageInfo_MenuItem proto.InternalMessageInfo

func (m *MenuItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MenuItem) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

type MenuOption struct {
	ShopID               int32            `protobuf:"varint,1,opt,name=ShopID,proto3" json:"ShopID,omitempty"`
	SelectNum            int32            `protobuf:"varint,2,opt,name=SelectNum,proto3" json:"SelectNum,omitempty"`
	Items                []*MenuItem      `protobuf:"bytes,3,rep,name=Items,proto3" json:"Items,omitempty"`
	Selections           []*MenuSelection `protobuf:"bytes,4,rep,name=Selections,proto3" json:"Selections,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *MenuOption) Reset()         { *m = MenuOption{} }
func (m *MenuOption) String() string { return proto.CompactTextString(m) }
func (*MenuOption) ProtoMessage()    {}
func (*MenuOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec784fbf625bb9ca, []int{2}
}

func (m *MenuOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MenuOption.Unmarshal(m, b)
}
func (m *MenuOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MenuOption.Marshal(b, m, deterministic)
}
func (m *MenuOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MenuOption.Merge(m, src)
}
func (m *MenuOption) XXX_Size() int {
	return xxx_messageInfo_MenuOption.Size(m)
}
func (m *MenuOption) XXX_DiscardUnknown() {
	xxx_messageInfo_MenuOption.DiscardUnknown(m)
}

var xxx_messageInfo_MenuOption proto.InternalMessageInfo

func (m *MenuOption) GetShopID() int32 {
	if m != nil {
		return m.ShopID
	}
	return 0
}

func (m *MenuOption) GetSelectNum() int32 {
	if m != nil {
		return m.SelectNum
	}
	return 0
}

func (m *MenuOption) GetItems() []*MenuItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *MenuOption) GetSelections() []*MenuSelection {
	if m != nil {
		return m.Selections
	}
	return nil
}

func init() {
	proto.RegisterType((*MenuSelection)(nil), "reqs.MenuSelection")
	proto.RegisterType((*MenuItem)(nil), "reqs.MenuItem")
	proto.RegisterType((*MenuOption)(nil), "reqs.MenuOption")
}

func init() { proto.RegisterFile("src/handler/models/reqs/menu.proto", fileDescriptor_ec784fbf625bb9ca) }

var fileDescriptor_ec784fbf625bb9ca = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2a, 0x2e, 0x4a, 0xd6,
	0xcf, 0x48, 0xcc, 0x4b, 0xc9, 0x49, 0x2d, 0xd2, 0xcf, 0xcd, 0x4f, 0x49, 0xcd, 0x29, 0xd6, 0x2f,
	0x4a, 0x2d, 0x2c, 0xd6, 0xcf, 0x4d, 0xcd, 0x2b, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x01, 0x09, 0x28, 0x59, 0x72, 0xf1, 0xfa, 0xa6, 0xe6, 0x95, 0x06, 0xa7, 0xe6, 0xa4, 0x26, 0x97,
	0x64, 0xe6, 0xe7, 0x09, 0x09, 0x71, 0xb1, 0xf8, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x06, 0x81, 0xd9, 0x42, 0x22, 0x5c, 0xac, 0x01, 0x45, 0x99, 0xc9, 0xa9, 0x12, 0x4c, 0x0a,
	0x8c, 0x1a, 0xac, 0x41, 0x10, 0x8e, 0x92, 0x09, 0x17, 0x07, 0x48, 0xab, 0x67, 0x49, 0x6a, 0x2e,
	0x09, 0xba, 0xe6, 0x32, 0x72, 0x71, 0x81, 0xb4, 0xf9, 0x17, 0x80, 0xad, 0x13, 0xe3, 0x62, 0x0b,
	0xce, 0xc8, 0x2f, 0xf0, 0x74, 0x01, 0x6b, 0x65, 0x0d, 0x82, 0xf2, 0x84, 0x64, 0xb8, 0x38, 0x21,
	0x6e, 0xf2, 0x2b, 0xcd, 0x85, 0x1a, 0x80, 0x10, 0x10, 0x52, 0xe1, 0x62, 0x05, 0x59, 0x5b, 0x2c,
	0xc1, 0xac, 0xc0, 0xac, 0xc1, 0x6d, 0xc4, 0xa7, 0x07, 0xf2, 0x8b, 0x1e, 0xcc, 0x35, 0x41, 0x10,
	0x49, 0x21, 0x63, 0x2e, 0x2e, 0xb8, 0xbf, 0x8a, 0x25, 0x58, 0xc0, 0x4a, 0x85, 0x11, 0x4a, 0xe1,
	0x72, 0x41, 0x48, 0xca, 0x92, 0xd8, 0xc0, 0xa1, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x30,
	0xe6, 0x63, 0x66, 0x43, 0x01, 0x00, 0x00,
}
