// Code generated by protoc-gen-go. DO NOT EDIT.
// source: geometric/vector.proto

package geometric

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import choreo "github.com/RobotStudio/choreo-msg/msg/primitive"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Vector3Stamped struct {
	Header *choreo.Header `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Vector *Vector3       `protobuf:"bytes,2,opt,name=vector" json:"vector,omitempty"`
}

func (m *Vector3Stamped) Reset()                    { *m = Vector3Stamped{} }
func (m *Vector3Stamped) String() string            { return proto.CompactTextString(m) }
func (*Vector3Stamped) ProtoMessage()               {}
func (*Vector3Stamped) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

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
func (*Vector3) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

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

func init() { proto.RegisterFile("geometric/vector.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0x31, 0x4b, 0xc7, 0x30,
	0x14, 0xc4, 0x89, 0x42, 0x84, 0x28, 0x15, 0x32, 0x94, 0xe2, 0x24, 0x1d, 0xd4, 0xc5, 0x04, 0xed,
	0x37, 0x70, 0x72, 0x4e, 0xc1, 0xc1, 0xad, 0x4d, 0x1f, 0x6d, 0x86, 0xf0, 0x4a, 0x7c, 0x2d, 0x6d,
	0x3f, 0xbd, 0x34, 0xc9, 0xff, 0x3f, 0x64, 0xb8, 0x77, 0xbf, 0xdc, 0x71, 0xa2, 0x1c, 0x01, 0x3d,
	0x50, 0x70, 0x56, 0xaf, 0x60, 0x09, 0x83, 0x9a, 0x03, 0x12, 0x4a, 0x6e, 0x27, 0x0c, 0x80, 0x4f,
	0xe5, 0x1c, 0x9c, 0x77, 0xe4, 0x56, 0xd0, 0x13, 0x74, 0x03, 0x64, 0xbf, 0xee, 0x44, 0xf1, 0x13,
	0xf9, 0xa6, 0xa5, 0xce, 0xcf, 0x30, 0xc8, 0x17, 0xc1, 0x13, 0x51, 0xb1, 0x67, 0xf6, 0x76, 0xff,
	0x59, 0xa8, 0x14, 0xa1, 0xbe, 0xe3, 0xd5, 0x64, 0x57, 0xbe, 0x0a, 0x9e, 0x9a, 0xaa, 0x9b, 0xc8,
	0x3d, 0x5e, 0xb8, 0x9c, 0x67, 0xb2, 0x5d, 0x37, 0xe2, 0x2e, 0x9f, 0xe4, 0x83, 0x60, 0x5b, 0x8c,
	0x65, 0x86, 0x6d, 0xa7, 0xda, 0xe3, 0x67, 0x66, 0xd8, 0x7e, 0xaa, 0xa3, 0xba, 0x4d, 0xea, 0xf8,
	0xfa, 0xf8, 0xd5, 0xa3, 0xa3, 0x69, 0xe9, 0x95, 0x45, 0xaf, 0x0d, 0xf6, 0x48, 0x2d, 0x2d, 0x83,
	0x43, 0x9d, 0x5a, 0xde, 0xfd, 0xdf, 0xa8, 0xcf, 0x77, 0xdd, 0xdd, 0xf3, 0xb8, 0xa8, 0xf9, 0x0f,
	0x00, 0x00, 0xff, 0xff, 0x67, 0xc4, 0xf9, 0xb4, 0x0b, 0x01, 0x00, 0x00,
}
