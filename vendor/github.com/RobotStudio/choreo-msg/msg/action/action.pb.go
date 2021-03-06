// Code generated by protoc-gen-go. DO NOT EDIT.
// source: action/action.proto

/*
Package action is a generated protocol buffer package.

It is generated from these files:
	action/action.proto

It has these top-level messages:
	GoalID
	GoalStatus
	GoalStatusArray
*/
package action

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import choreo "github.com/RobotStudio/choreo-msg/msg/primitive"
import choreo1 "github.com/RobotStudio/choreo-msg/msg/primitive"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GoalStatus_GoalStatusType int32

const (
	GoalStatus_PENDING    GoalStatus_GoalStatusType = 0
	GoalStatus_ACTIVE     GoalStatus_GoalStatusType = 1
	GoalStatus_PREEMPTED  GoalStatus_GoalStatusType = 2
	GoalStatus_SUCCEEDED  GoalStatus_GoalStatusType = 3
	GoalStatus_ABORTED    GoalStatus_GoalStatusType = 4
	GoalStatus_REJECTED   GoalStatus_GoalStatusType = 5
	GoalStatus_PREEMPTING GoalStatus_GoalStatusType = 6
	GoalStatus_RECALLING  GoalStatus_GoalStatusType = 7
	GoalStatus_RECALLED   GoalStatus_GoalStatusType = 8
	GoalStatus_LOST       GoalStatus_GoalStatusType = 9
)

var GoalStatus_GoalStatusType_name = map[int32]string{
	0: "PENDING",
	1: "ACTIVE",
	2: "PREEMPTED",
	3: "SUCCEEDED",
	4: "ABORTED",
	5: "REJECTED",
	6: "PREEMPTING",
	7: "RECALLING",
	8: "RECALLED",
	9: "LOST",
}
var GoalStatus_GoalStatusType_value = map[string]int32{
	"PENDING":    0,
	"ACTIVE":     1,
	"PREEMPTED":  2,
	"SUCCEEDED":  3,
	"ABORTED":    4,
	"REJECTED":   5,
	"PREEMPTING": 6,
	"RECALLING":  7,
	"RECALLED":   8,
	"LOST":       9,
}

func (x GoalStatus_GoalStatusType) String() string {
	return proto.EnumName(GoalStatus_GoalStatusType_name, int32(x))
}
func (GoalStatus_GoalStatusType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type GoalID struct {
	Stamp *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=stamp" json:"stamp,omitempty"`
	Id    *choreo1.String            `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
}

func (m *GoalID) Reset()                    { *m = GoalID{} }
func (m *GoalID) String() string            { return proto.CompactTextString(m) }
func (*GoalID) ProtoMessage()               {}
func (*GoalID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GoalID) GetStamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Stamp
	}
	return nil
}

func (m *GoalID) GetId() *choreo1.String {
	if m != nil {
		return m.Id
	}
	return nil
}

type GoalStatus struct {
	GoalId *GoalID                   `protobuf:"bytes,1,opt,name=goal_id,json=goalId" json:"goal_id,omitempty"`
	Status GoalStatus_GoalStatusType `protobuf:"varint,2,opt,name=status,enum=choreo.GoalStatus_GoalStatusType" json:"status,omitempty"`
	Text   *choreo1.String           `protobuf:"bytes,3,opt,name=text" json:"text,omitempty"`
}

func (m *GoalStatus) Reset()                    { *m = GoalStatus{} }
func (m *GoalStatus) String() string            { return proto.CompactTextString(m) }
func (*GoalStatus) ProtoMessage()               {}
func (*GoalStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GoalStatus) GetGoalId() *GoalID {
	if m != nil {
		return m.GoalId
	}
	return nil
}

func (m *GoalStatus) GetStatus() GoalStatus_GoalStatusType {
	if m != nil {
		return m.Status
	}
	return GoalStatus_PENDING
}

func (m *GoalStatus) GetText() *choreo1.String {
	if m != nil {
		return m.Text
	}
	return nil
}

type GoalStatusArray struct {
	Header     *choreo.Header `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	StatusList []*GoalStatus  `protobuf:"bytes,2,rep,name=status_list,json=statusList" json:"status_list,omitempty"`
}

func (m *GoalStatusArray) Reset()                    { *m = GoalStatusArray{} }
func (m *GoalStatusArray) String() string            { return proto.CompactTextString(m) }
func (*GoalStatusArray) ProtoMessage()               {}
func (*GoalStatusArray) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GoalStatusArray) GetHeader() *choreo.Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *GoalStatusArray) GetStatusList() []*GoalStatus {
	if m != nil {
		return m.StatusList
	}
	return nil
}

func init() {
	proto.RegisterType((*GoalID)(nil), "choreo.GoalID")
	proto.RegisterType((*GoalStatus)(nil), "choreo.GoalStatus")
	proto.RegisterType((*GoalStatusArray)(nil), "choreo.GoalStatusArray")
	proto.RegisterEnum("choreo.GoalStatus_GoalStatusType", GoalStatus_GoalStatusType_name, GoalStatus_GoalStatusType_value)
}

func init() { proto.RegisterFile("action/action.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x89, 0x93, 0x6e, 0xd2, 0x09, 0x04, 0x6b, 0x91, 0x50, 0x94, 0x03, 0x14, 0x1f, 0xa0,
	0x07, 0x58, 0xa3, 0xf4, 0xc4, 0x31, 0xb5, 0x57, 0x25, 0xc8, 0xb4, 0xd1, 0xda, 0x70, 0xe8, 0xa5,
	0xb2, 0xe3, 0xc5, 0x59, 0xc9, 0xce, 0x5a, 0xf6, 0x1a, 0xd1, 0x47, 0xe1, 0x2d, 0x78, 0x44, 0xb4,
	0xbb, 0x2e, 0x14, 0xd2, 0x83, 0x65, 0xcf, 0xcc, 0xff, 0xfd, 0x33, 0x1a, 0x0f, 0x3c, 0x4b, 0xb7,
	0x4a, 0xc8, 0xbd, 0x6f, 0x5f, 0xa4, 0x6e, 0xa4, 0x92, 0x18, 0x6d, 0x77, 0xb2, 0xe1, 0x72, 0xf1,
	0xb2, 0x90, 0xb2, 0x28, 0xb9, 0x6f, 0xb2, 0x59, 0xf7, 0xcd, 0x57, 0xa2, 0xe2, 0xad, 0x4a, 0xab,
	0xda, 0x0a, 0x17, 0xcf, 0xeb, 0x46, 0x54, 0x42, 0x89, 0xef, 0xdc, 0xdf, 0xf1, 0x34, 0xe7, 0xcd,
	0x61, 0xbe, 0x55, 0x8d, 0xd8, 0x17, 0x36, 0xef, 0x5d, 0x03, 0xba, 0x90, 0x69, 0xb9, 0x0e, 0xf1,
	0x7b, 0x38, 0x32, 0x46, 0xf3, 0xc1, 0xc9, 0xe0, 0x74, 0xba, 0x5c, 0x10, 0xdb, 0x8a, 0xdc, 0xb5,
	0x22, 0xc9, 0x5d, 0x2b, 0x66, 0x85, 0xf8, 0x05, 0x38, 0x22, 0x9f, 0x3b, 0x46, 0x3e, 0x23, 0x76,
	0x42, 0x12, 0x1b, 0x77, 0xe6, 0x88, 0xdc, 0xfb, 0xe5, 0x00, 0x68, 0xf3, 0x58, 0xa5, 0xaa, 0x6b,
	0xf1, 0x1b, 0x18, 0x17, 0x32, 0x2d, 0x6f, 0x44, 0xde, 0xb7, 0xf8, 0xc3, 0xd8, 0x09, 0x18, 0xd2,
	0xe5, 0x75, 0x8e, 0x3f, 0x00, 0x6a, 0x0d, 0x62, 0xbc, 0x67, 0xcb, 0x57, 0xf7, 0x75, 0xd6, 0xec,
	0xde, 0x67, 0x72, 0x5b, 0x73, 0xd6, 0x03, 0xd8, 0x83, 0x91, 0xe2, 0x3f, 0xd4, 0x7c, 0xf8, 0xe0,
	0x50, 0xa6, 0xe6, 0xfd, 0x1c, 0xc0, 0xec, 0x5f, 0x1c, 0x4f, 0x61, 0xbc, 0xa1, 0x97, 0xe1, 0xfa,
	0xf2, 0xc2, 0x7d, 0x84, 0x01, 0xd0, 0x2a, 0x48, 0xd6, 0x5f, 0xa9, 0x3b, 0xc0, 0x4f, 0xe0, 0x78,
	0xc3, 0x28, 0xfd, 0xbc, 0x49, 0x68, 0xe8, 0x3a, 0x3a, 0x8c, 0xbf, 0x04, 0x01, 0xa5, 0x21, 0x0d,
	0xdd, 0xa1, 0xc6, 0x56, 0xe7, 0x57, 0x4c, 0xd7, 0x46, 0xf8, 0x31, 0x4c, 0x18, 0xfd, 0x44, 0x03,
	0x1d, 0x1d, 0xe1, 0x19, 0x40, 0x0f, 0x6a, 0x53, 0xa4, 0x49, 0x46, 0x83, 0x55, 0x14, 0xe9, 0x70,
	0x6c, 0xc5, 0x3a, 0xa4, 0xa1, 0x3b, 0xc1, 0x13, 0x18, 0x45, 0x57, 0x71, 0xe2, 0x1e, 0x7b, 0x7b,
	0x78, 0xfa, 0x77, 0xb4, 0x55, 0xd3, 0xa4, 0xb7, 0xf8, 0x35, 0x20, 0xfb, 0x27, 0xff, 0xdf, 0xda,
	0x47, 0x93, 0x65, 0x7d, 0x15, 0x9f, 0xc1, 0xd4, 0x2e, 0xe1, 0xa6, 0x14, 0xad, 0x9a, 0x3b, 0x27,
	0xc3, 0xd3, 0xe9, 0x12, 0x1f, 0xae, 0x8e, 0x81, 0x95, 0x45, 0xa2, 0x55, 0xe7, 0xe4, 0xfa, 0x6d,
	0x21, 0xd4, 0xae, 0xcb, 0xc8, 0x56, 0x56, 0x3e, 0x93, 0x99, 0x54, 0xb1, 0xea, 0x72, 0x21, 0x7d,
	0xcb, 0xbd, 0xab, 0xda, 0xc2, 0xd7, 0x8f, 0xbd, 0xc6, 0x0c, 0x99, 0x6b, 0x38, 0xfb, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0xfa, 0xd0, 0xfa, 0x69, 0xa5, 0x02, 0x00, 0x00,
}
