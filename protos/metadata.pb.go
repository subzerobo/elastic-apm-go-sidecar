// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/metadata.proto

package protos

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

type MetaData struct {
	Service              *Service `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Process              *Process `protobuf:"bytes,2,opt,name=process,proto3" json:"process,omitempty"`
	System               *System  `protobuf:"bytes,3,opt,name=system,proto3" json:"system,omitempty"`
	User                 *User    `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetaData) Reset()         { *m = MetaData{} }
func (m *MetaData) String() string { return proto.CompactTextString(m) }
func (*MetaData) ProtoMessage()    {}
func (*MetaData) Descriptor() ([]byte, []int) {
	return fileDescriptor_615ae332be27cf42, []int{0}
}

func (m *MetaData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaData.Unmarshal(m, b)
}
func (m *MetaData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaData.Marshal(b, m, deterministic)
}
func (m *MetaData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaData.Merge(m, src)
}
func (m *MetaData) XXX_Size() int {
	return xxx_messageInfo_MetaData.Size(m)
}
func (m *MetaData) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaData.DiscardUnknown(m)
}

var xxx_messageInfo_MetaData proto.InternalMessageInfo

func (m *MetaData) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *MetaData) GetProcess() *Process {
	if m != nil {
		return m.Process
	}
	return nil
}

func (m *MetaData) GetSystem() *System {
	if m != nil {
		return m.System
	}
	return nil
}

func (m *MetaData) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*MetaData)(nil), "protos.MetaData")
}

func init() { proto.RegisterFile("protos/metadata.proto", fileDescriptor_615ae332be27cf42) }

var fileDescriptor_615ae332be27cf42 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x0a, 0xc2, 0x30,
	0x0c, 0x40, 0x99, 0x8e, 0x29, 0x55, 0x14, 0xab, 0x42, 0xf1, 0x34, 0x3c, 0x88, 0x5e, 0x26, 0xe8,
	0x2f, 0x78, 0x15, 0xa4, 0xe2, 0x07, 0xd4, 0x99, 0x83, 0x87, 0xb1, 0xd1, 0x44, 0xc1, 0x5f, 0xf2,
	0x2b, 0x65, 0x4d, 0x3a, 0xf0, 0xd8, 0xf7, 0x1e, 0x69, 0x88, 0x5a, 0x36, 0xbe, 0xa6, 0x1a, 0xf7,
	0x15, 0x90, 0x7b, 0x38, 0x72, 0x45, 0x78, 0xeb, 0x8c, 0xf1, 0x6a, 0x21, 0x1a, 0xc1, 0xbf, 0x9f,
	0x25, 0xb0, 0xed, 0x68, 0xe3, 0xeb, 0x12, 0x10, 0x85, 0xce, 0x63, 0xfb, 0x41, 0x82, 0x4a, 0xe0,
	0x4c, 0xe0, 0x0b, 0xc1, 0x33, 0x5a, 0x7f, 0x13, 0x35, 0x3c, 0x03, 0xb9, 0x93, 0x23, 0xa7, 0x77,
	0x6a, 0x20, 0xb3, 0x4d, 0x92, 0x27, 0xdb, 0xd1, 0x61, 0xca, 0x15, 0x16, 0x57, 0xc6, 0x36, 0xfa,
	0x36, 0x95, 0x0f, 0x4d, 0xef, 0x3f, 0xbd, 0x30, 0xb6, 0xd1, 0xeb, 0x8d, 0xca, 0x78, 0x0b, 0xd3,
	0x0f, 0xe5, 0xa4, 0x1b, 0x1a, 0xa8, 0x15, 0xab, 0x73, 0x95, 0xb6, 0x8b, 0x99, 0x34, 0x54, 0xe3,
	0x58, 0xdd, 0x10, 0xbc, 0x0d, 0xe6, 0xce, 0x87, 0x38, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x80,
	0xf0, 0x01, 0xbf, 0x28, 0x01, 0x00, 0x00,
}
