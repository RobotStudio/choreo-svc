// Code generated by protoc-gen-go. DO NOT EDIT.
// source: geometric/vector.proto

/*
Package msg is a generated protocol buffer package.

It is generated from these files:
	geometric/vector.proto

It has these top-level messages:
	Vector3Stamped
	Vector3
*/
package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import choreo "github.com/RobotStudio/choreo-msg/msg"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Vector3Stamped struct {
	Header *choreo.Header `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Vector *Vector3       `protobuf:"bytes,2,opt,name=vector" json:"vector,omitempty"`
}

func (m *Vector3Stamped) Reset()                    { *m = Vector3Stamped{} }
func (m *Vector3Stamped) String() string            { return proto.CompactTextString(m) }
func (*Vector3Stamped) ProtoMessage()               {}
func (*Vector3Stamped) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Vector3Stamped) GetHeader() *choreo.Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Vector3Stamped) GetVector() *Vector3 {
	if m != nil {
		return m.Vector
	}
	return nil
}

type Vector3 struct {
	X float64 `protobuf:"fixed64,1,opt,name=x" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y" json:"y,omitempty"`
	Z float64 `protobuf:"fixed64,3,opt,name=z" json:"z,omitempty"`
}

func (m *Vector3) Reset()                    { *m = Vector3{} }
func (m *Vector3) String() string            { return proto.CompactTextString(m) }
func (*Vector3) ProtoMessage()               {}
func (*Vector3) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Vector3) GetX() float64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Vector3) GetY() float64 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Vector3) GetZ() float64 {
	if m != nil {
		return m.Z
	}
	return 0
}

func init() {
	proto.RegisterType((*Vector3Stamped)(nil), "choreo.Vector3Stamped")
	proto.RegisterType((*Vector3)(nil), "choreo.Vector3")
}

func init() { proto.RegisterFile("geometric/vector.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8f, 0x3f, 0x4b, 0xc5, 0x30,
	0x14, 0xc5, 0x89, 0x42, 0x84, 0x28, 0x4f, 0xc8, 0xf0, 0x28, 0x4e, 0xf2, 0x06, 0xff, 0x20, 0x26,
	0x60, 0x47, 0x37, 0x27, 0xe7, 0x14, 0x1c, 0xdc, 0xda, 0xf4, 0x92, 0x66, 0x08, 0xb7, 0xa4, 0xb7,
	0xa5, 0xed, 0xa7, 0x97, 0x26, 0x71, 0x3c, 0xe7, 0xfc, 0xee, 0x39, 0x5c, 0x71, 0x76, 0x80, 0x01,
	0x28, 0x7a, 0xab, 0x17, 0xb0, 0x84, 0x51, 0x8d, 0x11, 0x09, 0x25, 0xb7, 0x03, 0x46, 0xc0, 0x87,
	0xf3, 0x18, 0x7d, 0xf0, 0xe4, 0x17, 0xd0, 0x03, 0xb4, 0x3d, 0x94, 0xfc, 0xd2, 0x8a, 0xd3, 0x4f,
	0xe2, 0xeb, 0x86, 0xda, 0x30, 0x42, 0x2f, 0x9f, 0x04, 0xcf, 0x44, 0xc5, 0x1e, 0xd9, 0xcb, 0xed,
	0xc7, 0x49, 0xe5, 0x0a, 0xf5, 0x9d, 0x5c, 0x53, 0x52, 0xf9, 0x2c, 0x78, 0x5e, 0xaa, 0xae, 0x12,
	0x77, 0xff, 0xcf, 0x95, 0x3e, 0x53, 0xe2, 0x4b, 0x2d, 0x6e, 0x8a, 0x25, 0xef, 0x04, 0x5b, 0x53,
	0x2d, 0x33, 0x6c, 0x3d, 0xd4, 0x96, 0x8e, 0x99, 0x61, 0xdb, 0xa1, 0xf6, 0xea, 0x3a, 0xab, 0xfd,
	0xeb, 0xed, 0xf7, 0xd5, 0x79, 0x1a, 0xe6, 0x4e, 0x59, 0x0c, 0xda, 0x60, 0x87, 0xd4, 0xd0, 0xdc,
	0x7b, 0xd4, 0x79, 0xe5, 0x3d, 0x4c, 0x4e, 0x87, 0xc9, 0x7d, 0x86, 0xc9, 0x75, 0x3c, 0xfd, 0x52,
	0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x8c, 0xeb, 0xb6, 0x71, 0x05, 0x01, 0x00, 0x00,
}
