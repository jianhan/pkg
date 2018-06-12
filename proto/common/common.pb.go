// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/common/common.proto

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	proto/common/common.proto

It has these top-level messages:
	Timestamp
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Timestamp struct {
	CreatedAt uint64 `protobuf:"varint,1,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt uint64 `protobuf:"varint,2,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *Timestamp) Reset()                    { *m = Timestamp{} }
func (m *Timestamp) String() string            { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()               {}
func (*Timestamp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Timestamp) GetCreatedAt() uint64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Timestamp) GetUpdatedAt() uint64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*Timestamp)(nil), "go.micro.srv.pkg.common.Timestamp")
}

func init() { proto.RegisterFile("proto/common/common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 127 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x83, 0x52, 0x7a, 0x60, 0x31, 0x21, 0xf1, 0xf4,
	0x7c, 0xbd, 0xdc, 0xcc, 0xe4, 0xa2, 0x7c, 0xbd, 0xe2, 0xa2, 0x32, 0xbd, 0x82, 0xec, 0x74, 0x3d,
	0x88, 0xb4, 0x92, 0x27, 0x17, 0x67, 0x48, 0x66, 0x6e, 0x6a, 0x71, 0x49, 0x62, 0x6e, 0x81, 0x90,
	0x2c, 0x17, 0x57, 0x72, 0x51, 0x6a, 0x62, 0x49, 0x6a, 0x4a, 0x7c, 0x62, 0x89, 0x04, 0xa3, 0x02,
	0xa3, 0x06, 0x4b, 0x10, 0x27, 0x54, 0xc4, 0xb1, 0x04, 0x24, 0x5d, 0x5a, 0x90, 0x02, 0x93, 0x66,
	0x82, 0x48, 0x43, 0x45, 0x1c, 0x4b, 0x9c, 0x38, 0xa2, 0xd8, 0x20, 0x86, 0x26, 0xb1, 0x81, 0x2d,
	0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x63, 0x4a, 0x63, 0x39, 0x91, 0x00, 0x00, 0x00,
}
