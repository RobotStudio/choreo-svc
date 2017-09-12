// Code generated by protoc-gen-go. DO NOT EDIT.
// source: geometric/polygon.proto

package rs_choreo_msg_geometric

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import choreo "github.com/RobotStudio/choreo-msg/msg/primitive"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PolygonStamped struct {
	Header  *choreo.Header `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Polygon *Polygon       `protobuf:"bytes,2,opt,name=polygon" json:"polygon,omitempty"`
}

func (m *PolygonStamped) Reset()                    { *m = PolygonStamped{} }
func (m *PolygonStamped) String() string            { return proto.CompactTextString(m) }
func (*PolygonStamped) ProtoMessage()               {}
func (*PolygonStamped) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *PolygonStamped) GetHeader() *choreo.Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PolygonStamped) GetPolygon() *Polygon {
	if m != nil {
		return m.Polygon
	}
	return nil
}

type Polygon32Stamped struct {
	Header  *choreo.Header `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Polygon *Polygon32     `protobuf:"bytes,2,opt,name=polygon" json:"polygon,omitempty"`
}

func (m *Polygon32Stamped) Reset()                    { *m = Polygon32Stamped{} }
func (m *Polygon32Stamped) String() string            { return proto.CompactTextString(m) }
func (*Polygon32Stamped) ProtoMessage()               {}
func (*Polygon32Stamped) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *Polygon32Stamped) GetHeader() *choreo.Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Polygon32Stamped) GetPolygon() *Polygon32 {
	if m != nil {
		return m.Polygon
	}
	return nil
}

type Polygon struct {
	Points []*Point `protobuf:"bytes,1,rep,name=points" json:"points,omitempty"`
}

func (m *Polygon) Reset()                    { *m = Polygon{} }
func (m *Polygon) String() string            { return proto.CompactTextString(m) }
func (*Polygon) ProtoMessage()               {}
func (*Polygon) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *Polygon) GetPoints() []*Point {
	if m != nil {
		return m.Points
	}
	return nil
}

type Polygon32 struct {
	Points []*Point32 `protobuf:"bytes,1,rep,name=points" json:"points,omitempty"`
}

func (m *Polygon32) Reset()                    { *m = Polygon32{} }
func (m *Polygon32) String() string            { return proto.CompactTextString(m) }
func (*Polygon32) ProtoMessage()               {}
func (*Polygon32) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *Polygon32) GetPoints() []*Point32 {
	if m != nil {
		return m.Points
	}
	return nil
}

func init() {
	proto.RegisterType((*PolygonStamped)(nil), "choreo.PolygonStamped")
	proto.RegisterType((*Polygon32Stamped)(nil), "choreo.Polygon32Stamped")
	proto.RegisterType((*Polygon)(nil), "choreo.Polygon")
	proto.RegisterType((*Polygon32)(nil), "choreo.Polygon32")
}

func init() { proto.RegisterFile("geometric/polygon.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0x4b, 0xc4, 0x30,
	0x14, 0xc6, 0xa9, 0x42, 0x0f, 0x73, 0x78, 0xa7, 0x01, 0xb5, 0xdc, 0x74, 0x14, 0xd4, 0x13, 0x31,
	0x91, 0xd4, 0xcd, 0xcd, 0x45, 0xdd, 0xa4, 0xb7, 0xb9, 0x5d, 0xdb, 0x90, 0x06, 0x4c, 0x5f, 0x49,
	0xdf, 0x09, 0xfe, 0xf7, 0x62, 0x93, 0xb4, 0xda, 0xcd, 0x21, 0xcb, 0xf7, 0xbd, 0xdf, 0xf7, 0x25,
	0x79, 0xe4, 0x42, 0x49, 0x30, 0x12, 0xad, 0x2e, 0x79, 0x0b, 0x1f, 0x5f, 0x0a, 0x1a, 0xd6, 0x5a,
	0x40, 0xa0, 0x71, 0x59, 0x83, 0x95, 0xb0, 0x3a, 0x6f, 0xad, 0x36, 0x1a, 0xf5, 0xa7, 0xe4, 0xb5,
	0xdc, 0x55, 0xd2, 0x3a, 0x7f, 0x75, 0xf6, 0x1b, 0xd4, 0x0d, 0x3a, 0x39, 0x2d, 0xc9, 0xe2, 0xcd,
	0xe5, 0x6c, 0x71, 0x67, 0x5a, 0x59, 0xd1, 0x2b, 0x12, 0x3b, 0x30, 0x89, 0xd6, 0xd1, 0x66, 0x2e,
	0x16, 0xcc, 0x25, 0xb3, 0x97, 0x5e, 0xcd, 0xbd, 0x4b, 0x6f, 0xc8, 0xcc, 0xdf, 0x20, 0x39, 0xe8,
	0x07, 0x97, 0x61, 0xd0, 0x07, 0xe6, 0xc1, 0x4f, 0x15, 0x39, 0xf1, 0x5a, 0x26, 0xfe, 0x5b, 0x73,
	0x3b, 0xad, 0x39, 0x9d, 0xd4, 0x64, 0x62, 0x2c, 0xba, 0x27, 0x33, 0xaf, 0xd2, 0x4b, 0x12, 0xf7,
	0xef, 0xec, 0x92, 0x68, 0x7d, 0xb8, 0x99, 0x8b, 0xe3, 0x11, 0xd3, 0x0d, 0xe6, 0xde, 0x4c, 0x1f,
	0xc8, 0xd1, 0x90, 0x43, 0xaf, 0x27, 0xcc, 0xf2, 0x0f, 0x93, 0x89, 0x40, 0x3d, 0xbd, 0xbe, 0x3f,
	0x2b, 0x8d, 0xf5, 0xbe, 0x60, 0x25, 0x18, 0x9e, 0x43, 0x01, 0xb8, 0xc5, 0x7d, 0xa5, 0x81, 0x3b,
	0xe0, 0xce, 0x74, 0x8a, 0xff, 0x9c, 0xe1, 0xd3, 0x1f, 0x6d, 0x17, 0xc2, 0x4c, 0xa7, 0xd8, 0xa0,
	0x17, 0x71, 0xbf, 0x87, 0xec, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x1a, 0xdb, 0xbc, 0xd9, 0x01,
	0x00, 0x00,
}