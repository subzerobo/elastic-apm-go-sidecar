// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/spanv2.proto

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

type Span struct {
	Timestamp            int64         `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Id                   string        `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	TransactionId        string        `protobuf:"bytes,3,opt,name=transaction_id,proto3" json:"transaction_id,omitempty"`
	TraceId              string        `protobuf:"bytes,4,opt,name=trace_id,proto3" json:"trace_id,omitempty"`
	ParentId             string        `protobuf:"bytes,5,opt,name=parent_id,proto3" json:"parent_id,omitempty"`
	Start                float32       `protobuf:"fixed32,6,opt,name=start,proto3" json:"start,omitempty"`
	Subtype              string        `protobuf:"bytes,7,opt,name=subtype,proto3" json:"subtype,omitempty"`
	Action               string        `protobuf:"bytes,8,opt,name=action,proto3" json:"action,omitempty"`
	Context              *Span_Context `protobuf:"bytes,9,opt,name=context,proto3" json:"context,omitempty"`
	Duration             float32       `protobuf:"fixed32,10,opt,name=duration,proto3" json:"duration,omitempty"`
	Name                 string        `protobuf:"bytes,11,opt,name=name,proto3" json:"name,omitempty"`
	Stacktrace           []*StackTrace `protobuf:"bytes,12,rep,name=stacktrace,proto3" json:"stacktrace,omitempty"`
	Type                 string        `protobuf:"bytes,13,opt,name=type,proto3" json:"type,omitempty"`
	Sync                 bool          `protobuf:"varint,14,opt,name=sync,proto3" json:"sync,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Span) Reset()         { *m = Span{} }
func (m *Span) String() string { return proto.CompactTextString(m) }
func (*Span) ProtoMessage()    {}
func (*Span) Descriptor() ([]byte, []int) {
	return fileDescriptor_c25dc91180002fe9, []int{0}
}

func (m *Span) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Span.Unmarshal(m, b)
}
func (m *Span) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Span.Marshal(b, m, deterministic)
}
func (m *Span) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Span.Merge(m, src)
}
func (m *Span) XXX_Size() int {
	return xxx_messageInfo_Span.Size(m)
}
func (m *Span) XXX_DiscardUnknown() {
	xxx_messageInfo_Span.DiscardUnknown(m)
}

var xxx_messageInfo_Span proto.InternalMessageInfo

func (m *Span) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Span) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Span) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

func (m *Span) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

func (m *Span) GetParentId() string {
	if m != nil {
		return m.ParentId
	}
	return ""
}

func (m *Span) GetStart() float32 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Span) GetSubtype() string {
	if m != nil {
		return m.Subtype
	}
	return ""
}

func (m *Span) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Span) GetContext() *Span_Context {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Span) GetDuration() float32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Span) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Span) GetStacktrace() []*StackTrace {
	if m != nil {
		return m.Stacktrace
	}
	return nil
}

func (m *Span) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Span) GetSync() bool {
	if m != nil {
		return m.Sync
	}
	return false
}

type Span_Context struct {
	Db                   *Span_Context_DB   `protobuf:"bytes,1,opt,name=db,proto3" json:"db,omitempty"`
	Http                 *Span_Context_HTTP `protobuf:"bytes,2,opt,name=http,proto3" json:"http,omitempty"`
	Tags                 map[string]string  `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Span_Context) Reset()         { *m = Span_Context{} }
func (m *Span_Context) String() string { return proto.CompactTextString(m) }
func (*Span_Context) ProtoMessage()    {}
func (*Span_Context) Descriptor() ([]byte, []int) {
	return fileDescriptor_c25dc91180002fe9, []int{0, 0}
}

func (m *Span_Context) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Span_Context.Unmarshal(m, b)
}
func (m *Span_Context) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Span_Context.Marshal(b, m, deterministic)
}
func (m *Span_Context) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Span_Context.Merge(m, src)
}
func (m *Span_Context) XXX_Size() int {
	return xxx_messageInfo_Span_Context.Size(m)
}
func (m *Span_Context) XXX_DiscardUnknown() {
	xxx_messageInfo_Span_Context.DiscardUnknown(m)
}

var xxx_messageInfo_Span_Context proto.InternalMessageInfo

func (m *Span_Context) GetDb() *Span_Context_DB {
	if m != nil {
		return m.Db
	}
	return nil
}

func (m *Span_Context) GetHttp() *Span_Context_HTTP {
	if m != nil {
		return m.Http
	}
	return nil
}

func (m *Span_Context) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type Span_Context_DB struct {
	Instance             string   `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	Statement            string   `protobuf:"bytes,2,opt,name=statement,proto3" json:"statement,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	User                 string   `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Span_Context_DB) Reset()         { *m = Span_Context_DB{} }
func (m *Span_Context_DB) String() string { return proto.CompactTextString(m) }
func (*Span_Context_DB) ProtoMessage()    {}
func (*Span_Context_DB) Descriptor() ([]byte, []int) {
	return fileDescriptor_c25dc91180002fe9, []int{0, 0, 1}
}

func (m *Span_Context_DB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Span_Context_DB.Unmarshal(m, b)
}
func (m *Span_Context_DB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Span_Context_DB.Marshal(b, m, deterministic)
}
func (m *Span_Context_DB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Span_Context_DB.Merge(m, src)
}
func (m *Span_Context_DB) XXX_Size() int {
	return xxx_messageInfo_Span_Context_DB.Size(m)
}
func (m *Span_Context_DB) XXX_DiscardUnknown() {
	xxx_messageInfo_Span_Context_DB.DiscardUnknown(m)
}

var xxx_messageInfo_Span_Context_DB proto.InternalMessageInfo

func (m *Span_Context_DB) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *Span_Context_DB) GetStatement() string {
	if m != nil {
		return m.Statement
	}
	return ""
}

func (m *Span_Context_DB) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Span_Context_DB) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

type Span_Context_HTTP struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	StatusCode           int32    `protobuf:"varint,2,opt,name=status_code,proto3" json:"status_code,omitempty"`
	Method               string   `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Span_Context_HTTP) Reset()         { *m = Span_Context_HTTP{} }
func (m *Span_Context_HTTP) String() string { return proto.CompactTextString(m) }
func (*Span_Context_HTTP) ProtoMessage()    {}
func (*Span_Context_HTTP) Descriptor() ([]byte, []int) {
	return fileDescriptor_c25dc91180002fe9, []int{0, 0, 2}
}

func (m *Span_Context_HTTP) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Span_Context_HTTP.Unmarshal(m, b)
}
func (m *Span_Context_HTTP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Span_Context_HTTP.Marshal(b, m, deterministic)
}
func (m *Span_Context_HTTP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Span_Context_HTTP.Merge(m, src)
}
func (m *Span_Context_HTTP) XXX_Size() int {
	return xxx_messageInfo_Span_Context_HTTP.Size(m)
}
func (m *Span_Context_HTTP) XXX_DiscardUnknown() {
	xxx_messageInfo_Span_Context_HTTP.DiscardUnknown(m)
}

var xxx_messageInfo_Span_Context_HTTP proto.InternalMessageInfo

func (m *Span_Context_HTTP) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Span_Context_HTTP) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *Span_Context_HTTP) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func init() {
	proto.RegisterType((*Span)(nil), "protos.Span")
	proto.RegisterType((*Span_Context)(nil), "protos.Span.Context")
	proto.RegisterMapType((map[string]string)(nil), "protos.Span.Context.TagsEntry")
	proto.RegisterType((*Span_Context_DB)(nil), "protos.Span.Context.DB")
	proto.RegisterType((*Span_Context_HTTP)(nil), "protos.Span.Context.HTTP")
}

func init() { proto.RegisterFile("protos/spanv2.proto", fileDescriptor_c25dc91180002fe9) }

var fileDescriptor_c25dc91180002fe9 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xcd, 0x8e, 0xd4, 0x3c,
	0x10, 0x54, 0x7e, 0xe6, 0xaf, 0xf3, 0x7d, 0x23, 0x64, 0x56, 0xac, 0x89, 0x10, 0x8a, 0x38, 0x40,
	0x2e, 0x04, 0x29, 0x1c, 0x40, 0x1c, 0x97, 0x45, 0xe2, 0x88, 0xcc, 0xdc, 0x57, 0x9e, 0xc4, 0xec,
	0x46, 0x3b, 0x71, 0x22, 0xbb, 0xb3, 0x62, 0x5e, 0x81, 0x17, 0xe1, 0x35, 0x91, 0xdb, 0x49, 0x66,
	0x41, 0x73, 0xeb, 0xea, 0xaa, 0x74, 0x57, 0xb7, 0x3b, 0xf0, 0xb4, 0x37, 0x1d, 0x76, 0xf6, 0x9d,
	0xed, 0xa5, 0x7e, 0x28, 0x0b, 0x42, 0x6c, 0xe9, 0x93, 0xe9, 0xe5, 0x44, 0xa2, 0xac, 0xee, 0xd1,
	0xc8, 0x4a, 0x79, 0xc1, 0xab, 0xdf, 0x4b, 0x88, 0xbf, 0xf7, 0x52, 0xb3, 0x17, 0xb0, 0xc1, 0xa6,
	0x55, 0x16, 0x65, 0xdb, 0xf3, 0x20, 0x0b, 0xf2, 0x48, 0x9c, 0x12, 0x6c, 0x0b, 0x61, 0x53, 0xf3,
	0x30, 0x0b, 0xf2, 0x8d, 0x08, 0x9b, 0x9a, 0xbd, 0x86, 0x2d, 0x1a, 0xa9, 0xad, 0xac, 0xb0, 0xe9,
	0xf4, 0x4d, 0x53, 0xf3, 0x88, 0xb8, 0x7f, 0xb2, 0x2c, 0x85, 0x35, 0x75, 0x73, 0x8a, 0x98, 0x14,
	0x33, 0x76, 0x1d, 0x7b, 0x69, 0x94, 0x46, 0x47, 0x2e, 0x88, 0x3c, 0x25, 0xd8, 0x05, 0x2c, 0x2c,
	0x4a, 0x83, 0x7c, 0x99, 0x05, 0x79, 0x28, 0x3c, 0x60, 0x1c, 0x56, 0x76, 0xd8, 0xe3, 0xb1, 0x57,
	0x7c, 0x45, 0x5f, 0x4c, 0x90, 0x3d, 0x83, 0xa5, 0x6f, 0xcb, 0xd7, 0x44, 0x8c, 0x88, 0x15, 0xb0,
	0xaa, 0x3a, 0x8d, 0xea, 0x27, 0xf2, 0x4d, 0x16, 0xe4, 0x49, 0x79, 0xe1, 0x27, 0xb7, 0x85, 0x1b,
	0xbb, 0xf8, 0xec, 0x39, 0x31, 0x89, 0x9c, 0xe3, 0x7a, 0x30, 0x92, 0x2a, 0x01, 0xb5, 0x9e, 0x31,
	0x63, 0x10, 0x6b, 0xd9, 0x2a, 0x9e, 0x50, 0x07, 0x8a, 0x59, 0x09, 0x70, 0x5a, 0x2a, 0xff, 0x2f,
	0x8b, 0xf2, 0xa4, 0x64, 0x73, 0x0b, 0xc7, 0xec, 0x1c, 0x23, 0x1e, 0xa9, 0x5c, 0x1d, 0x1a, 0xe1,
	0x7f, 0x5f, 0x87, 0xfc, 0x33, 0x88, 0xed, 0x51, 0x57, 0x7c, 0x9b, 0x05, 0xf9, 0x5a, 0x50, 0x9c,
	0xfe, 0x8a, 0x60, 0x35, 0x1a, 0x64, 0x6f, 0x20, 0xac, 0xf7, 0xf4, 0x30, 0x49, 0x79, 0x79, 0x6e,
	0x84, 0xe2, 0xfa, 0x4a, 0x84, 0xf5, 0x9e, 0xbd, 0x85, 0xf8, 0x0e, 0xb1, 0xa7, 0xc7, 0x4a, 0xca,
	0xe7, 0x67, 0xa5, 0x5f, 0x77, 0xbb, 0x6f, 0x82, 0x64, 0xac, 0x84, 0x18, 0xe5, 0xad, 0xe5, 0x11,
	0x39, 0x7f, 0x79, 0x56, 0xbe, 0x93, 0xb7, 0xf6, 0x8b, 0x46, 0x73, 0x14, 0xa4, 0x4d, 0x3f, 0xc0,
	0x66, 0x4e, 0xb1, 0x27, 0x10, 0xdd, 0xab, 0x23, 0x39, 0xdb, 0x08, 0x17, 0xba, 0xa7, 0x7b, 0x90,
	0x87, 0x41, 0x8d, 0xf7, 0xe2, 0xc1, 0xa7, 0xf0, 0x63, 0x90, 0xfe, 0x80, 0xf0, 0xfa, 0xca, 0xad,
	0xb8, 0xd1, 0x16, 0xa5, 0xae, 0xd4, 0xf8, 0xd9, 0x8c, 0xdd, 0x51, 0x58, 0x94, 0xa8, 0x5a, 0xa5,
	0x71, 0xfc, 0xfe, 0x94, 0x98, 0x17, 0x17, 0xfd, 0xbd, 0xb8, 0xc1, 0x2a, 0x33, 0x9e, 0x17, 0xc5,
	0xa9, 0x80, 0xd8, 0x8d, 0xe8, 0xbc, 0x0d, 0xe6, 0x30, 0x79, 0x1b, 0xcc, 0x81, 0x65, 0x90, 0xb8,
	0x72, 0x83, 0xbd, 0xa9, 0xba, 0xda, 0x3b, 0x5c, 0x88, 0xc7, 0x29, 0x77, 0x48, 0xad, 0xc2, 0xbb,
	0x6e, 0x3a, 0xe9, 0x11, 0xed, 0xfd, 0xaf, 0xf4, 0xfe, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe3,
	0x1e, 0x9a, 0xa4, 0x68, 0x03, 0x00, 0x00,
}
