// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/handler/models/resp/menu.proto

// 生成的程式在 Golang 中將會屬於 `resp` 套件。

package resp

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

// MenuItem 。
type MenuItem struct {
	Name                 string       `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	SizePrice            []*SizePrice `protobuf:"bytes,2,rep,name=SizePrice,proto3" json:"SizePrice,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *MenuItem) Reset()         { *m = MenuItem{} }
func (m *MenuItem) String() string { return proto.CompactTextString(m) }
func (*MenuItem) ProtoMessage()    {}
func (*MenuItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_9605b9fdee46970e, []int{0}
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

func (m *MenuItem) GetSizePrice() []*SizePrice {
	if m != nil {
		return m.SizePrice
	}
	return nil
}

// SizePrice 。
type SizePrice struct {
	Size                 string   `protobuf:"bytes,1,opt,name=Size,proto3" json:"Size,omitempty"`
	Price                int32    `protobuf:"varint,2,opt,name=Price,proto3" json:"Price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SizePrice) Reset()         { *m = SizePrice{} }
func (m *SizePrice) String() string { return proto.CompactTextString(m) }
func (*SizePrice) ProtoMessage()    {}
func (*SizePrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_9605b9fdee46970e, []int{1}
}

func (m *SizePrice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SizePrice.Unmarshal(m, b)
}
func (m *SizePrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SizePrice.Marshal(b, m, deterministic)
}
func (m *SizePrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SizePrice.Merge(m, src)
}
func (m *SizePrice) XXX_Size() int {
	return xxx_messageInfo_SizePrice.Size(m)
}
func (m *SizePrice) XXX_DiscardUnknown() {
	xxx_messageInfo_SizePrice.DiscardUnknown(m)
}

var xxx_messageInfo_SizePrice proto.InternalMessageInfo

func (m *SizePrice) GetSize() string {
	if m != nil {
		return m.Size
	}
	return ""
}

func (m *SizePrice) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

// Size 。
type Size struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Size                 string   `protobuf:"bytes,2,opt,name=Size,proto3" json:"Size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Size) Reset()         { *m = Size{} }
func (m *Size) String() string { return proto.CompactTextString(m) }
func (*Size) ProtoMessage()    {}
func (*Size) Descriptor() ([]byte, []int) {
	return fileDescriptor_9605b9fdee46970e, []int{2}
}

func (m *Size) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Size.Unmarshal(m, b)
}
func (m *Size) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Size.Marshal(b, m, deterministic)
}
func (m *Size) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Size.Merge(m, src)
}
func (m *Size) XXX_Size() int {
	return xxx_messageInfo_Size.Size(m)
}
func (m *Size) XXX_DiscardUnknown() {
	xxx_messageInfo_Size.DiscardUnknown(m)
}

var xxx_messageInfo_Size proto.InternalMessageInfo

func (m *Size) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Size) GetSize() string {
	if m != nil {
		return m.Size
	}
	return ""
}

// KindOption 。
type KindOption struct {
	Kind                 string   `protobuf:"bytes,1,opt,name=Kind,proto3" json:"Kind,omitempty"`
	Price                int32    `protobuf:"varint,2,opt,name=Price,proto3" json:"Price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KindOption) Reset()         { *m = KindOption{} }
func (m *KindOption) String() string { return proto.CompactTextString(m) }
func (*KindOption) ProtoMessage()    {}
func (*KindOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_9605b9fdee46970e, []int{3}
}

func (m *KindOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KindOption.Unmarshal(m, b)
}
func (m *KindOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KindOption.Marshal(b, m, deterministic)
}
func (m *KindOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KindOption.Merge(m, src)
}
func (m *KindOption) XXX_Size() int {
	return xxx_messageInfo_KindOption.Size(m)
}
func (m *KindOption) XXX_DiscardUnknown() {
	xxx_messageInfo_KindOption.DiscardUnknown(m)
}

var xxx_messageInfo_KindOption proto.InternalMessageInfo

func (m *KindOption) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *KindOption) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

// MenuKind 。
type MenuKind struct {
	Items                []*MenuItem   `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
	RequiredSelection    []*IntMap     `protobuf:"bytes,3,rep,name=RequiredSelection,proto3" json:"RequiredSelection,omitempty"`
	CheckOption          []*KindOption `protobuf:"bytes,4,rep,name=CheckOption,proto3" json:"CheckOption,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MenuKind) Reset()         { *m = MenuKind{} }
func (m *MenuKind) String() string { return proto.CompactTextString(m) }
func (*MenuKind) ProtoMessage()    {}
func (*MenuKind) Descriptor() ([]byte, []int) {
	return fileDescriptor_9605b9fdee46970e, []int{4}
}

func (m *MenuKind) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MenuKind.Unmarshal(m, b)
}
func (m *MenuKind) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MenuKind.Marshal(b, m, deterministic)
}
func (m *MenuKind) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MenuKind.Merge(m, src)
}
func (m *MenuKind) XXX_Size() int {
	return xxx_messageInfo_MenuKind.Size(m)
}
func (m *MenuKind) XXX_DiscardUnknown() {
	xxx_messageInfo_MenuKind.DiscardUnknown(m)
}

var xxx_messageInfo_MenuKind proto.InternalMessageInfo

func (m *MenuKind) GetItems() []*MenuItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *MenuKind) GetRequiredSelection() []*IntMap {
	if m != nil {
		return m.RequiredSelection
	}
	return nil
}

func (m *MenuKind) GetCheckOption() []*KindOption {
	if m != nil {
		return m.CheckOption
	}
	return nil
}

// IntPair 。
type IntPair struct {
	Key                  string   `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Value                int32    `protobuf:"varint,2,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntPair) Reset()         { *m = IntPair{} }
func (m *IntPair) String() string { return proto.CompactTextString(m) }
func (*IntPair) ProtoMessage()    {}
func (*IntPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_9605b9fdee46970e, []int{5}
}

func (m *IntPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntPair.Unmarshal(m, b)
}
func (m *IntPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntPair.Marshal(b, m, deterministic)
}
func (m *IntPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntPair.Merge(m, src)
}
func (m *IntPair) XXX_Size() int {
	return xxx_messageInfo_IntPair.Size(m)
}
func (m *IntPair) XXX_DiscardUnknown() {
	xxx_messageInfo_IntPair.DiscardUnknown(m)
}

var xxx_messageInfo_IntPair proto.InternalMessageInfo

func (m *IntPair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *IntPair) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// IntMap 。
type IntMap struct {
	Pairs                []*IntPair `protobuf:"bytes,1,rep,name=Pairs,proto3" json:"Pairs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *IntMap) Reset()         { *m = IntMap{} }
func (m *IntMap) String() string { return proto.CompactTextString(m) }
func (*IntMap) ProtoMessage()    {}
func (*IntMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_9605b9fdee46970e, []int{6}
}

func (m *IntMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntMap.Unmarshal(m, b)
}
func (m *IntMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntMap.Marshal(b, m, deterministic)
}
func (m *IntMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntMap.Merge(m, src)
}
func (m *IntMap) XXX_Size() int {
	return xxx_messageInfo_IntMap.Size(m)
}
func (m *IntMap) XXX_DiscardUnknown() {
	xxx_messageInfo_IntMap.DiscardUnknown(m)
}

var xxx_messageInfo_IntMap proto.InternalMessageInfo

func (m *IntMap) GetPairs() []*IntPair {
	if m != nil {
		return m.Pairs
	}
	return nil
}

func init() {
	proto.RegisterType((*MenuItem)(nil), "resp.MenuItem")
	proto.RegisterType((*SizePrice)(nil), "resp.SizePrice")
	proto.RegisterType((*Size)(nil), "resp.Size")
	proto.RegisterType((*KindOption)(nil), "resp.KindOption")
	proto.RegisterType((*MenuKind)(nil), "resp.MenuKind")
	proto.RegisterType((*IntPair)(nil), "resp.IntPair")
	proto.RegisterType((*IntMap)(nil), "resp.IntMap")
}

func init() { proto.RegisterFile("src/handler/models/resp/menu.proto", fileDescriptor_9605b9fdee46970e) }

var fileDescriptor_9605b9fdee46970e = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x41, 0x4f, 0xbb, 0x40,
	0x10, 0xc5, 0x03, 0x2d, 0xfd, 0xff, 0x3b, 0xd5, 0x5a, 0x37, 0x1e, 0x38, 0x36, 0xab, 0x87, 0xc6,
	0xa4, 0x25, 0xd6, 0xe8, 0xc1, 0xab, 0xbd, 0x90, 0x06, 0x6d, 0xb6, 0x89, 0x77, 0x84, 0x49, 0xba,
	0x11, 0x16, 0x5c, 0xe0, 0xa0, 0x9f, 0xc6, 0x8f, 0x6a, 0x66, 0x57, 0xa0, 0x89, 0xf1, 0x36, 0x33,
	0xfb, 0xde, 0x9b, 0xf9, 0x01, 0xf0, 0x4a, 0x27, 0xc1, 0x21, 0x56, 0x69, 0x86, 0x3a, 0xc8, 0x8b,
	0x14, 0xb3, 0x2a, 0xd0, 0x58, 0x95, 0x41, 0x8e, 0xaa, 0x59, 0x95, 0xba, 0xa8, 0x0b, 0x36, 0xa4,
	0x01, 0x8f, 0xe0, 0x7f, 0x84, 0xaa, 0x09, 0x6b, 0xcc, 0x19, 0x83, 0xe1, 0x53, 0x9c, 0xa3, 0xef,
	0xcc, 0x9d, 0xc5, 0x58, 0x98, 0x9a, 0x2d, 0x61, 0xbc, 0x97, 0x9f, 0xb8, 0xd3, 0x32, 0x41, 0xdf,
	0x9d, 0x0f, 0x16, 0x93, 0xf5, 0xd9, 0x8a, 0x9c, 0xab, 0x6e, 0x2c, 0x7a, 0x05, 0xbf, 0x3b, 0x92,
	0x53, 0x1e, 0x35, 0x6d, 0x1e, 0xd5, 0xec, 0x02, 0xbc, 0x36, 0xcb, 0x59, 0x78, 0xc2, 0x36, 0xfc,
	0xda, 0x2a, 0xd9, 0x14, 0xdc, 0x70, 0x63, 0xf4, 0x9e, 0x70, 0xc3, 0x4d, 0x97, 0xe0, 0xf6, 0x09,
	0xfc, 0x1e, 0x60, 0x2b, 0x55, 0xfa, 0x5c, 0xd6, 0xb2, 0x50, 0xa4, 0xa0, 0xae, 0xdd, 0x41, 0xf5,
	0x1f, 0x3b, 0xbe, 0x1c, 0x8b, 0x6a, 0x24, 0x57, 0xe0, 0x11, 0x72, 0xe5, 0x3b, 0x06, 0x69, 0x6a,
	0x91, 0xda, 0x2f, 0x21, 0xec, 0x23, 0x7b, 0x80, 0x73, 0x81, 0xef, 0x8d, 0xd4, 0x98, 0xee, 0x31,
	0xc3, 0x84, 0x36, 0xfa, 0x03, 0xe3, 0x38, 0xb1, 0x8e, 0x50, 0xd5, 0x51, 0x5c, 0x8a, 0xdf, 0x32,
	0xb6, 0x86, 0xc9, 0xe3, 0x01, 0x93, 0x37, 0x7b, 0xa7, 0x3f, 0x34, 0xae, 0x99, 0x75, 0xf5, 0xf7,
	0x8b, 0x63, 0x11, 0xbf, 0x81, 0x7f, 0xa1, 0xaa, 0x77, 0xb1, 0xd4, 0x6c, 0x06, 0x83, 0x2d, 0x7e,
	0xfc, 0x60, 0x51, 0x49, 0x54, 0x2f, 0x71, 0xd6, 0x74, 0x54, 0xa6, 0xe1, 0x4b, 0x18, 0xd9, 0x1b,
	0xd8, 0x25, 0x78, 0xe4, 0x6c, 0x91, 0x4e, 0xbb, 0x03, 0x69, 0x2a, 0xec, 0xdb, 0xeb, 0xc8, 0xfc,
	0xfb, 0xdb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x10, 0x40, 0xff, 0x8e, 0x21, 0x02, 0x00, 0x00,
}
