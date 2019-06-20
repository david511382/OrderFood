// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/handler/models/resp/view.proto

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

type UpdateView struct {
	HTML                 []*KeyValue `protobuf:"bytes,1,rep,name=HTML,proto3" json:"HTML,omitempty"`
	Script               []*KeyValue `protobuf:"bytes,2,rep,name=Script,proto3" json:"Script,omitempty"`
	Css                  []*KeyValue `protobuf:"bytes,3,rep,name=Css,proto3" json:"Css,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UpdateView) Reset()         { *m = UpdateView{} }
func (m *UpdateView) String() string { return proto.CompactTextString(m) }
func (*UpdateView) ProtoMessage()    {}
func (*UpdateView) Descriptor() ([]byte, []int) {
	return fileDescriptor_85640e131710636d, []int{0}
}

func (m *UpdateView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateView.Unmarshal(m, b)
}
func (m *UpdateView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateView.Marshal(b, m, deterministic)
}
func (m *UpdateView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateView.Merge(m, src)
}
func (m *UpdateView) XXX_Size() int {
	return xxx_messageInfo_UpdateView.Size(m)
}
func (m *UpdateView) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateView.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateView proto.InternalMessageInfo

func (m *UpdateView) GetHTML() []*KeyValue {
	if m != nil {
		return m.HTML
	}
	return nil
}

func (m *UpdateView) GetScript() []*KeyValue {
	if m != nil {
		return m.Script
	}
	return nil
}

func (m *UpdateView) GetCss() []*KeyValue {
	if m != nil {
		return m.Css
	}
	return nil
}

type KeyValue struct {
	Key                  string   `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Data                 string   `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyValue) Reset()         { *m = KeyValue{} }
func (m *KeyValue) String() string { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()    {}
func (*KeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_85640e131710636d, []int{1}
}

func (m *KeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValue.Unmarshal(m, b)
}
func (m *KeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValue.Marshal(b, m, deterministic)
}
func (m *KeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValue.Merge(m, src)
}
func (m *KeyValue) XXX_Size() int {
	return xxx_messageInfo_KeyValue.Size(m)
}
func (m *KeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValue proto.InternalMessageInfo

func (m *KeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValue) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*UpdateView)(nil), "resp.UpdateView")
	proto.RegisterType((*KeyValue)(nil), "resp.KeyValue")
}

func init() { proto.RegisterFile("src/handler/models/resp/view.proto", fileDescriptor_85640e131710636d) }

var fileDescriptor_85640e131710636d = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2a, 0x2e, 0x4a, 0xd6,
	0xcf, 0x48, 0xcc, 0x4b, 0xc9, 0x49, 0x2d, 0xd2, 0xcf, 0xcd, 0x4f, 0x49, 0xcd, 0x29, 0xd6, 0x2f,
	0x4a, 0x2d, 0x2e, 0xd0, 0x2f, 0xcb, 0x4c, 0x2d, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x01, 0x09, 0x28, 0x55, 0x71, 0x71, 0x85, 0x16, 0xa4, 0x24, 0x96, 0xa4, 0x86, 0x65, 0xa6, 0x96,
	0x0b, 0x29, 0x71, 0xb1, 0x78, 0x84, 0xf8, 0xfa, 0x48, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0xf1,
	0xe9, 0x81, 0x94, 0xe8, 0x79, 0xa7, 0x56, 0x86, 0x25, 0xe6, 0x94, 0xa6, 0x06, 0x81, 0xe5, 0x84,
	0xd4, 0xb8, 0xd8, 0x82, 0x93, 0x8b, 0x32, 0x0b, 0x4a, 0x24, 0x98, 0xb0, 0xaa, 0x82, 0xca, 0x0a,
	0x29, 0x70, 0x31, 0x3b, 0x17, 0x17, 0x4b, 0x30, 0x63, 0x55, 0x04, 0x92, 0x52, 0x32, 0xe0, 0xe2,
	0x80, 0x09, 0x08, 0x09, 0x70, 0x31, 0x7b, 0xa7, 0x56, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06,
	0x81, 0x98, 0x42, 0x42, 0x5c, 0x2c, 0x2e, 0x89, 0x25, 0x89, 0x12, 0x4c, 0x60, 0x21, 0x30, 0x3b,
	0x89, 0x0d, 0xec, 0x74, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x69, 0xc0, 0x3e, 0xd4, 0xe0,
	0x00, 0x00, 0x00,
}
