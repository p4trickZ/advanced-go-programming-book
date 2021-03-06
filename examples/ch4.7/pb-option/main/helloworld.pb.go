// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

package main // import "main"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message0 struct {
	Name                 *string  `protobuf:"bytes,1,opt,name=name,def=gopher" json:"name,omitempty"`
	Age                  *int32   `protobuf:"varint,2,opt,name=age,def=10" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message0) Reset()         { *m = Message0{} }
func (m *Message0) String() string { return proto.CompactTextString(m) }
func (*Message0) ProtoMessage()    {}
func (*Message0) Descriptor() ([]byte, []int) {
	return fileDescriptor_helloworld_a2a557ac07ecc4b1, []int{0}
}
func (m *Message0) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message0.Unmarshal(m, b)
}
func (m *Message0) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message0.Marshal(b, m, deterministic)
}
func (dst *Message0) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message0.Merge(dst, src)
}
func (m *Message0) XXX_Size() int {
	return xxx_messageInfo_Message0.Size(m)
}
func (m *Message0) XXX_DiscardUnknown() {
	xxx_messageInfo_Message0.DiscardUnknown(m)
}

var xxx_messageInfo_Message0 proto.InternalMessageInfo

const Default_Message0_Name string = "gopher"
const Default_Message0_Age int32 = 10

func (m *Message0) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return Default_Message0_Name
}

func (m *Message0) GetAge() int32 {
	if m != nil && m.Age != nil {
		return *m.Age
	}
	return Default_Message0_Age
}

type Message2 struct {
	Name                 *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Age                  *int32   `protobuf:"varint,2,opt,name=age" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message2) Reset()         { *m = Message2{} }
func (m *Message2) String() string { return proto.CompactTextString(m) }
func (*Message2) ProtoMessage()    {}
func (*Message2) Descriptor() ([]byte, []int) {
	return fileDescriptor_helloworld_a2a557ac07ecc4b1, []int{1}
}
func (m *Message2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message2.Unmarshal(m, b)
}
func (m *Message2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message2.Marshal(b, m, deterministic)
}
func (dst *Message2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message2.Merge(dst, src)
}
func (m *Message2) XXX_Size() int {
	return xxx_messageInfo_Message2.Size(m)
}
func (m *Message2) XXX_DiscardUnknown() {
	xxx_messageInfo_Message2.DiscardUnknown(m)
}

var xxx_messageInfo_Message2 proto.InternalMessageInfo

func (m *Message2) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Message2) GetAge() int32 {
	if m != nil && m.Age != nil {
		return *m.Age
	}
	return 0
}

var E_Email = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50000,
	Name:          "main.email",
	Tag:           "bytes,50000,opt,name=email",
	Filename:      "helloworld.proto",
}

var E_DefaultString = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50000,
	Name:          "main.default_string",
	Tag:           "bytes,50000,opt,name=default_string,json=defaultString",
	Filename:      "helloworld.proto",
}

var E_DefaultInt = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*int32)(nil),
	Field:         50001,
	Name:          "main.default_int",
	Tag:           "varint,50001,opt,name=default_int,json=defaultInt",
	Filename:      "helloworld.proto",
}

func init() {
	proto.RegisterType((*Message0)(nil), "main.Message0")
	proto.RegisterType((*Message2)(nil), "main.Message2")
	proto.RegisterExtension(E_Email)
	proto.RegisterExtension(E_DefaultString)
	proto.RegisterExtension(E_DefaultInt)
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor_helloworld_a2a557ac07ecc4b1) }

var fileDescriptor_helloworld_a2a557ac07ecc4b1 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0xd9, 0xd6, 0x8a, 0x8b, 0x28, 0x12, 0x54, 0xc2, 0x50, 0x29, 0xc3, 0xc3, 0x4e, 0xdd,
	0x14, 0x4f, 0xd5, 0xc3, 0xe8, 0x61, 0xe0, 0x41, 0x84, 0x7a, 0xf3, 0x22, 0xb1, 0xfd, 0x96, 0x06,
	0xd2, 0x7c, 0x25, 0x49, 0xf1, 0xbe, 0x83, 0xf8, 0x87, 0xfa, 0x3f, 0xf4, 0x1f, 0x49, 0xdb, 0x4d,
	0xa6, 0x08, 0x5e, 0x02, 0x79, 0x79, 0x9f, 0x87, 0x17, 0x3e, 0x72, 0x98, 0x83, 0x52, 0xf8, 0x8a,
	0x46, 0x65, 0x61, 0x69, 0xd0, 0x21, 0xf5, 0x0a, 0x2e, 0xf5, 0x28, 0x10, 0x88, 0x42, 0xc1, 0xb4,
	0xcd, 0x5e, 0xaa, 0xe5, 0x34, 0x03, 0x9b, 0x1a, 0x59, 0x3a, 0x34, 0x5d, 0x6f, 0x7c, 0x4b, 0x76,
	0xef, 0xc1, 0x5a, 0x2e, 0x60, 0x46, 0x47, 0xc4, 0xd3, 0xbc, 0x00, 0xd6, 0x0b, 0x7a, 0x93, 0x61,
	0xb4, 0x23, 0xb0, 0xcc, 0xc1, 0x24, 0x6d, 0x46, 0x8f, 0xc8, 0x80, 0x0b, 0x60, 0xfd, 0xa0, 0x37,
	0xf1, 0xa3, 0xfe, 0xe5, 0x2c, 0x69, 0xbe, 0xe3, 0xf8, 0x9b, 0xbe, 0xa2, 0xe7, 0xdb, 0x74, 0x4c,
	0x56, 0x35, 0xfb, 0x69, 0x38, 0xd9, 0x32, 0xc4, 0xde, 0x7b, 0xcd, 0x48, 0xeb, 0x88, 0xae, 0x89,
	0x0f, 0x05, 0x97, 0x8a, 0x9e, 0x86, 0xdd, 0xda, 0x70, 0xb3, 0x36, 0x5c, 0x48, 0x05, 0x0f, 0xa5,
	0x93, 0xa8, 0x2d, 0xfb, 0x78, 0x1b, 0x34, 0xe2, 0xa4, 0x2b, 0x47, 0x0b, 0x72, 0x90, 0xc1, 0x92,
	0x57, 0xca, 0x3d, 0x5b, 0x67, 0xa4, 0x16, 0xf4, 0xec, 0x0f, 0x1c, 0x54, 0xf6, 0x9b, 0xdf, 0x5f,
	0x63, 0x8f, 0x2d, 0x15, 0xcd, 0xc9, 0xde, 0xc6, 0x23, 0xb5, 0xfb, 0x4f, 0xf2, 0xd9, 0x4a, 0xfc,
	0x84, 0xac, 0x99, 0x3b, 0xed, 0xe2, 0x8b, 0x55, 0xcd, 0x8e, 0xd3, 0x9c, 0x4b, 0x9b, 0x57, 0x36,
	0xe7, 0x7a, 0x2e, 0x9a, 0x81, 0x61, 0x8a, 0xc5, 0xd3, 0xb0, 0x39, 0xc2, 0x4d, 0xf3, 0x7c, 0x05,
	0x00, 0x00, 0xff, 0xff, 0xd2, 0x01, 0xe8, 0x5f, 0xa2, 0x01, 0x00, 0x00,
}
