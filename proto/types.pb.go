// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/types.proto

package proto // import "github.com/firecracker-microvm/firecracker-containerd/proto"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Message to store bundle/config.json bytes
type ExtraData struct {
	JsonSpec             []byte     `protobuf:"bytes,1,opt,name=JsonSpec,proto3" json:"JsonSpec,omitempty"`
	RuncOptions          *types.Any `protobuf:"bytes,2,opt,name=RuncOptions" json:"RuncOptions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ExtraData) Reset()         { *m = ExtraData{} }
func (m *ExtraData) String() string { return proto.CompactTextString(m) }
func (*ExtraData) ProtoMessage()    {}
func (*ExtraData) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_fd1a0f9dde9f57da, []int{0}
}
func (m *ExtraData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtraData.Unmarshal(m, b)
}
func (m *ExtraData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtraData.Marshal(b, m, deterministic)
}
func (dst *ExtraData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtraData.Merge(dst, src)
}
func (m *ExtraData) XXX_Size() int {
	return xxx_messageInfo_ExtraData.Size(m)
}
func (m *ExtraData) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtraData.DiscardUnknown(m)
}

var xxx_messageInfo_ExtraData proto.InternalMessageInfo

func (m *ExtraData) GetJsonSpec() []byte {
	if m != nil {
		return m.JsonSpec
	}
	return nil
}

func (m *ExtraData) GetRuncOptions() *types.Any {
	if m != nil {
		return m.RuncOptions
	}
	return nil
}

func init() {
	proto.RegisterType((*ExtraData)(nil), "firecracker.containerd.ExtraData")
}

func init() { proto.RegisterFile("proto/types.proto", fileDescriptor_types_fd1a0f9dde9f57da) }

var fileDescriptor_types_fd1a0f9dde9f57da = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x03, 0xb3, 0x85, 0xc4, 0xd2, 0x32, 0x8b, 0x52,
	0x93, 0x8b, 0x12, 0x93, 0xb3, 0x53, 0x8b, 0xf4, 0x92, 0xf3, 0xf3, 0x4a, 0x12, 0x33, 0xf3, 0x52,
	0x8b, 0x52, 0xa4, 0x24, 0xd3, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0xaa, 0x92, 0x4a, 0xd3,
	0xf4, 0x13, 0xf3, 0x2a, 0x21, 0x5a, 0x94, 0xe2, 0xb9, 0x38, 0x5d, 0x2b, 0x4a, 0x8a, 0x12, 0x5d,
	0x12, 0x4b, 0x12, 0x85, 0xa4, 0xb8, 0x38, 0xbc, 0x8a, 0xf3, 0xf3, 0x82, 0x0b, 0x52, 0x93, 0x25,
	0x18, 0x15, 0x18, 0x35, 0x78, 0x82, 0xe0, 0x7c, 0x21, 0x33, 0x2e, 0xee, 0xa0, 0xd2, 0xbc, 0x64,
	0xff, 0x82, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x11, 0x3d,
	0x88, 0xc9, 0x7a, 0x30, 0x93, 0xf5, 0x1c, 0xf3, 0x2a, 0x83, 0x90, 0x15, 0x3a, 0xd9, 0x46, 0x59,
	0xa7, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x23, 0x39, 0x50, 0x37, 0x37,
	0x33, 0xb9, 0x28, 0xbf, 0x0c, 0x55, 0x0c, 0xe1, 0x68, 0xa8, 0x63, 0xd9, 0xc0, 0x94, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0xad, 0x88, 0x4e, 0x21, 0xee, 0x00, 0x00, 0x00,
}
