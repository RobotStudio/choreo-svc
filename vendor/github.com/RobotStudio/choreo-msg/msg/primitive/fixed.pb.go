// Code generated by protoc-gen-go. DO NOT EDIT.
// source: primitive/fixed.proto

package primitive

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Fixed32Array struct {
	Data []*Fixed32 `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
}

func (m *Fixed32Array) Reset()                    { *m = Fixed32Array{} }
func (m *Fixed32Array) String() string            { return proto.CompactTextString(m) }
func (*Fixed32Array) ProtoMessage()               {}
func (*Fixed32Array) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Fixed32Array) GetData() []*Fixed32 {
	if m != nil {
		return m.Data
	}
	return nil
}

type Fixed32 struct {
	Data uint32 `protobuf:"varint,1,opt,name=data" json:"data,omitempty"`
}

func (m *Fixed32) Reset()                    { *m = Fixed32{} }
func (m *Fixed32) String() string            { return proto.CompactTextString(m) }
func (*Fixed32) ProtoMessage()               {}
func (*Fixed32) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *Fixed32) GetData() uint32 {
	if m != nil {
		return m.Data
	}
	return 0
}

type Fixed64Array struct {
	Data []*Fixed64 `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
}

func (m *Fixed64Array) Reset()                    { *m = Fixed64Array{} }
func (m *Fixed64Array) String() string            { return proto.CompactTextString(m) }
func (*Fixed64Array) ProtoMessage()               {}
func (*Fixed64Array) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *Fixed64Array) GetData() []*Fixed64 {
	if m != nil {
		return m.Data
	}
	return nil
}

type Fixed64 struct {
	Data uint64 `protobuf:"varint,1,opt,name=data" json:"data,omitempty"`
}

func (m *Fixed64) Reset()                    { *m = Fixed64{} }
func (m *Fixed64) String() string            { return proto.CompactTextString(m) }
func (*Fixed64) ProtoMessage()               {}
func (*Fixed64) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *Fixed64) GetData() uint64 {
	if m != nil {
		return m.Data
	}
	return 0
}

func init() {
	proto.RegisterType((*Fixed32Array)(nil), "choreo.Fixed32Array")
	proto.RegisterType((*Fixed32)(nil), "choreo.Fixed32")
	proto.RegisterType((*Fixed64Array)(nil), "choreo.Fixed64Array")
	proto.RegisterType((*Fixed64)(nil), "choreo.Fixed64")
}

func init() { proto.RegisterFile("primitive/fixed.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x28, 0xca, 0xcc,
	0xcd, 0x2c, 0xc9, 0x2c, 0x4b, 0xd5, 0x4f, 0xcb, 0xac, 0x48, 0x4d, 0xd1, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x4b, 0xce, 0xc8, 0x2f, 0x4a, 0xcd, 0x57, 0x32, 0xe6, 0xe2, 0x71, 0x03, 0x09,
	0x1b, 0x1b, 0x39, 0x16, 0x15, 0x25, 0x56, 0x0a, 0x29, 0x73, 0xb1, 0xa4, 0x24, 0x96, 0x24, 0x4a,
	0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0xf1, 0xeb, 0x41, 0x94, 0xe9, 0x41, 0xd5, 0x04, 0x81, 0x25,
	0x95, 0x64, 0xb9, 0xd8, 0xa1, 0x02, 0x42, 0x42, 0x70, 0xf5, 0x8c, 0x1a, 0xbc, 0x50, 0x69, 0x98,
	0x99, 0x66, 0x26, 0x84, 0xcd, 0x34, 0x33, 0x41, 0x33, 0xd3, 0xcc, 0x04, 0xc5, 0x4c, 0x16, 0x88,
	0xb4, 0x93, 0x61, 0x94, 0x7e, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e,
	0x50, 0x7e, 0x52, 0x7e, 0x49, 0x70, 0x49, 0x69, 0x4a, 0x66, 0xbe, 0x3e, 0xc4, 0x34, 0xdd, 0xdc,
	0xe2, 0x74, 0x7d, 0x10, 0x86, 0x7b, 0x37, 0x89, 0x0d, 0xec, 0x53, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x2d, 0x56, 0x0f, 0xf2, 0x02, 0x01, 0x00, 0x00,
}
