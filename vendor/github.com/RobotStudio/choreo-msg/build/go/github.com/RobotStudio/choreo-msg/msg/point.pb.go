// Code generated by protoc-gen-go. DO NOT EDIT.
// source: geometric/point.proto

/*
Package msg is a generated protocol buffer package.

It is generated from these files:
	geometric/point.proto

It has these top-level messages:
	Point32Stamped
	PointStamped
	Point32
	Point
*/
package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import choreo "github.com/RobotStudio/choreo-msg/msg"
import choreo1 "github.com/RobotStudio/choreo-msg/msg"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Point32Stamped struct {
	Header *choreo.Header `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Point  *Point32       `protobuf:"bytes,2,opt,name=point" json:"point,omitempty"`
}

func (m *Point32Stamped) Reset()                    { *m = Point32Stamped{} }
func (m *Point32Stamped) String() string            { return proto.CompactTextString(m) }
func (*Point32Stamped) ProtoMessage()               {}
func (*Point32Stamped) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Point32Stamped) GetHeader() *choreo.Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Point32Stamped) GetPoint() *Point32 {
	if m != nil {
		return m.Point
	}
	return nil
}

type PointStamped struct {
	Header *choreo.Header `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Point  *Point         `protobuf:"bytes,2,opt,name=point" json:"point,omitempty"`
}

func (m *PointStamped) Reset()                    { *m = PointStamped{} }
func (m *PointStamped) String() string            { return proto.CompactTextString(m) }
func (*PointStamped) ProtoMessage()               {}
func (*PointStamped) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PointStamped) GetHeader() *choreo.Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PointStamped) GetPoint() *Point {
	if m != nil {
		return m.Point
	}
	return nil
}

type Point32 struct {
	X *choreo1.Float32 `protobuf:"bytes,1,opt,name=x" json:"x,omitempty"`
	Y *choreo1.Float32 `protobuf:"bytes,2,opt,name=y" json:"y,omitempty"`
	Z *choreo1.Float32 `protobuf:"bytes,3,opt,name=z" json:"z,omitempty"`
}

func (m *Point32) Reset()                    { *m = Point32{} }
func (m *Point32) String() string            { return proto.CompactTextString(m) }
func (*Point32) ProtoMessage()               {}
func (*Point32) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Point32) GetX() *choreo1.Float32 {
	if m != nil {
		return m.X
	}
	return nil
}

func (m *Point32) GetY() *choreo1.Float32 {
	if m != nil {
		return m.Y
	}
	return nil
}

func (m *Point32) GetZ() *choreo1.Float32 {
	if m != nil {
		return m.Z
	}
	return nil
}

type Point struct {
	X *choreo1.Float64 `protobuf:"bytes,1,opt,name=x" json:"x,omitempty"`
	Y *choreo1.Float64 `protobuf:"bytes,2,opt,name=y" json:"y,omitempty"`
	Z *choreo1.Float64 `protobuf:"bytes,3,opt,name=z" json:"z,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Point) GetX() *choreo1.Float64 {
	if m != nil {
		return m.X
	}
	return nil
}

func (m *Point) GetY() *choreo1.Float64 {
	if m != nil {
		return m.Y
	}
	return nil
}

func (m *Point) GetZ() *choreo1.Float64 {
	if m != nil {
		return m.Z
	}
	return nil
}

func init() {
	proto.RegisterType((*Point32Stamped)(nil), "choreo.Point32Stamped")
	proto.RegisterType((*PointStamped)(nil), "choreo.PointStamped")
	proto.RegisterType((*Point32)(nil), "choreo.Point32")
	proto.RegisterType((*Point)(nil), "choreo.Point")
}

func init() { proto.RegisterFile("geometric/point.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0x89, 0xd2, 0x08, 0xab, 0x56, 0x08, 0x54, 0x42, 0x41, 0x90, 0x88, 0xa2, 0x88, 0x59,
	0x48, 0xc4, 0x8b, 0x37, 0x0f, 0xe2, 0x51, 0xd2, 0x9b, 0x1e, 0x24, 0x1f, 0xdb, 0xcd, 0x82, 0xeb,
	0x2c, 0x9b, 0xa9, 0xb4, 0xf9, 0xf5, 0xb2, 0x1f, 0xad, 0x54, 0x68, 0xc0, 0xeb, 0x3c, 0xef, 0x3c,
	0xef, 0x2c, 0x4b, 0x26, 0x9c, 0x81, 0x64, 0xa8, 0x45, 0x4d, 0x15, 0x88, 0x2f, 0x4c, 0x95, 0x06,
	0x84, 0x28, 0xac, 0x5b, 0xd0, 0x0c, 0xa6, 0xa7, 0x4a, 0x0b, 0x29, 0x50, 0x7c, 0x33, 0xda, 0xb2,
	0xb2, 0x61, 0xda, 0xf1, 0xe9, 0xe4, 0x77, 0x3e, 0xff, 0x84, 0xd2, 0xaf, 0x25, 0x1f, 0x64, 0xfc,
	0x6a, 0x2c, 0x79, 0x36, 0xc3, 0x52, 0x2a, 0xd6, 0x44, 0x57, 0x24, 0x74, 0x8b, 0x71, 0x70, 0x1e,
	0x5c, 0x1f, 0x66, 0xe3, 0xd4, 0x99, 0xd3, 0x17, 0x3b, 0x2d, 0x3c, 0x8d, 0x2e, 0xc9, 0xc8, 0xf6,
	0xc7, 0x7b, 0x36, 0x76, 0xb2, 0x8e, 0x79, 0x5d, 0xe1, 0x68, 0xf2, 0x4e, 0x8e, 0xec, 0xe4, 0xbf,
	0xfa, 0x8b, 0x6d, 0xfd, 0xf1, 0x96, 0x7e, 0x2d, 0x9f, 0x93, 0x03, 0x5f, 0x17, 0x9d, 0x91, 0x60,
	0xe9, 0x95, 0x9b, 0x53, 0x9e, 0xcd, 0x43, 0xf3, 0xac, 0x08, 0x96, 0x06, 0xaf, 0xfe, 0x5e, 0xba,
	0xc1, 0x2b, 0x83, 0xfb, 0x78, 0x7f, 0x07, 0xee, 0x93, 0x86, 0x8c, 0x6c, 0xcf, 0xee, 0x96, 0x87,
	0xfb, 0xc1, 0x16, 0x83, 0x07, 0x5a, 0x0c, 0xee, 0x9f, 0x6e, 0xdf, 0x6e, 0xb8, 0xc0, 0x76, 0x51,
	0xa5, 0x35, 0x48, 0x5a, 0x40, 0x05, 0x38, 0xc3, 0x45, 0x23, 0x80, 0xba, 0xec, 0x9d, 0xec, 0x38,
	0x95, 0x1d, 0x7f, 0x94, 0x1d, 0xaf, 0x42, 0xfb, 0x7f, 0xf9, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xad, 0x8a, 0x96, 0x16, 0x0f, 0x02, 0x00, 0x00,
}
