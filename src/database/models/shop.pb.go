// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/database/models/shop.proto

// 生成的程式在 Golang 中將會屬於 `models` 套件。

package models

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

// Shop 。
type Shop struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Shop) Reset()         { *m = Shop{} }
func (m *Shop) String() string { return proto.CompactTextString(m) }
func (*Shop) ProtoMessage()    {}
func (*Shop) Descriptor() ([]byte, []int) {
	return fileDescriptor_d042450818a3e96a, []int{0}
}

func (m *Shop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Shop.Unmarshal(m, b)
}
func (m *Shop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Shop.Marshal(b, m, deterministic)
}
func (m *Shop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Shop.Merge(m, src)
}
func (m *Shop) XXX_Size() int {
	return xxx_messageInfo_Shop.Size(m)
}
func (m *Shop) XXX_DiscardUnknown() {
	xxx_messageInfo_Shop.DiscardUnknown(m)
}

var xxx_messageInfo_Shop proto.InternalMessageInfo

func (m *Shop) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Shop) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Shop)(nil), "models.Shop")
}

func init() { proto.RegisterFile("src/database/models/shop.proto", fileDescriptor_d042450818a3e96a) }

var fileDescriptor_d042450818a3e96a = []byte{
	// 106 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x2e, 0x4a, 0xd6,
	0x4f, 0x49, 0x2c, 0x49, 0x4c, 0x4a, 0x2c, 0x4e, 0xd5, 0xcf, 0xcd, 0x4f, 0x49, 0xcd, 0x29, 0xd6,
	0x2f, 0xce, 0xc8, 0x2f, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0x08, 0x29, 0x69,
	0x71, 0xb1, 0x04, 0x67, 0xe4, 0x17, 0x08, 0xf1, 0x71, 0x31, 0x79, 0xba, 0x48, 0x30, 0x2a, 0x30,
	0x6a, 0xb0, 0x06, 0x31, 0x79, 0xba, 0x08, 0x09, 0x71, 0xb1, 0xf8, 0x25, 0xe6, 0xa6, 0x4a, 0x30,
	0x29, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x49, 0x6c, 0x60, 0xad, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x25, 0xce, 0xf0, 0x95, 0x5c, 0x00, 0x00, 0x00,
}
