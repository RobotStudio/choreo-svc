// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sensor/imu.proto

/*
Package msg is a generated protocol buffer package.

It is generated from these files:
	sensor/imu.proto

It has these top-level messages:
	IMU
	MagneticField
	NavSatStatus
	NavSatFix
*/
package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import choreo "github.com/RobotStudio/choreo-msg/msg"
import choreo1 "github.com/RobotStudio/choreo-msg/msg"
import choreo2 "github.com/RobotStudio/choreo-msg/msg"
import choreo3 "github.com/RobotStudio/choreo-msg/msg"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NavSatStatus_Status int32

const (
	NavSatStatus_NO_FIX   NavSatStatus_Status = 0
	NavSatStatus_FIX      NavSatStatus_Status = 1
	NavSatStatus_SBAS_FIX NavSatStatus_Status = 2
	NavSatStatus_GBAS_FIX NavSatStatus_Status = 3
)

var NavSatStatus_Status_name = map[int32]string{
	0: "NO_FIX",
	1: "FIX",
	2: "SBAS_FIX",
	3: "GBAS_FIX",
}
var NavSatStatus_Status_value = map[string]int32{
	"NO_FIX":   0,
	"FIX":      1,
	"SBAS_FIX": 2,
	"GBAS_FIX": 3,
}

func (x NavSatStatus_Status) String() string {
	return proto.EnumName(NavSatStatus_Status_name, int32(x))
}
func (NavSatStatus_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type NavSatStatus_Service int32

const (
	NavSatStatus_GPS     NavSatStatus_Service = 0
	NavSatStatus_GLONASS NavSatStatus_Service = 1
	NavSatStatus_COMPASS NavSatStatus_Service = 2
	NavSatStatus_GALILEO NavSatStatus_Service = 3
)

var NavSatStatus_Service_name = map[int32]string{
	0: "GPS",
	1: "GLONASS",
	2: "COMPASS",
	3: "GALILEO",
}
var NavSatStatus_Service_value = map[string]int32{
	"GPS":     0,
	"GLONASS": 1,
	"COMPASS": 2,
	"GALILEO": 3,
}

func (x NavSatStatus_Service) String() string {
	return proto.EnumName(NavSatStatus_Service_name, int32(x))
}
func (NavSatStatus_Service) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 1} }

type NavSatFix_PositionCovarianceType int32

const (
	NavSatFix_UNKNOWN          NavSatFix_PositionCovarianceType = 0
	NavSatFix_APPROXIMATED     NavSatFix_PositionCovarianceType = 1
	NavSatFix_DIAGONAL_UNKNOWN NavSatFix_PositionCovarianceType = 2
	NavSatFix_KNOWN            NavSatFix_PositionCovarianceType = 3
)

var NavSatFix_PositionCovarianceType_name = map[int32]string{
	0: "UNKNOWN",
	1: "APPROXIMATED",
	2: "DIAGONAL_UNKNOWN",
	3: "KNOWN",
}
var NavSatFix_PositionCovarianceType_value = map[string]int32{
	"UNKNOWN":          0,
	"APPROXIMATED":     1,
	"DIAGONAL_UNKNOWN": 2,
	"KNOWN":            3,
}

func (x NavSatFix_PositionCovarianceType) String() string {
	return proto.EnumName(NavSatFix_PositionCovarianceType_name, int32(x))
}
func (NavSatFix_PositionCovarianceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

type IMU struct {
	Header                       *choreo1.Header     `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Orientation                  *choreo2.Quaternion `protobuf:"bytes,2,opt,name=orientation" json:"orientation,omitempty"`
	OrientationCovariance        []*choreo.Float64   `protobuf:"bytes,3,rep,name=orientation_covariance,json=orientationCovariance" json:"orientation_covariance,omitempty"`
	AngularVelocity              *choreo3.Vector3    `protobuf:"bytes,4,opt,name=angular_velocity,json=angularVelocity" json:"angular_velocity,omitempty"`
	AngularVelocityCovariance    []*choreo.Float64   `protobuf:"bytes,5,rep,name=angular_velocity_covariance,json=angularVelocityCovariance" json:"angular_velocity_covariance,omitempty"`
	LinearAcceleration           *choreo3.Vector3    `protobuf:"bytes,6,opt,name=linear_acceleration,json=linearAcceleration" json:"linear_acceleration,omitempty"`
	LinearAccelerationCovariance []*choreo.Float64   `protobuf:"bytes,7,rep,name=linear_acceleration_covariance,json=linearAccelerationCovariance" json:"linear_acceleration_covariance,omitempty"`
}

func (m *IMU) Reset()                    { *m = IMU{} }
func (m *IMU) String() string            { return proto.CompactTextString(m) }
func (*IMU) ProtoMessage()               {}
func (*IMU) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IMU) GetHeader() *choreo1.Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *IMU) GetOrientation() *choreo2.Quaternion {
	if m != nil {
		return m.Orientation
	}
	return nil
}

func (m *IMU) GetOrientationCovariance() []*choreo.Float64 {
	if m != nil {
		return m.OrientationCovariance
	}
	return nil
}

func (m *IMU) GetAngularVelocity() *choreo3.Vector3 {
	if m != nil {
		return m.AngularVelocity
	}
	return nil
}

func (m *IMU) GetAngularVelocityCovariance() []*choreo.Float64 {
	if m != nil {
		return m.AngularVelocityCovariance
	}
	return nil
}

func (m *IMU) GetLinearAcceleration() *choreo3.Vector3 {
	if m != nil {
		return m.LinearAcceleration
	}
	return nil
}

func (m *IMU) GetLinearAccelerationCovariance() []*choreo.Float64 {
	if m != nil {
		return m.LinearAccelerationCovariance
	}
	return nil
}

type MagneticField struct {
	Header                  *choreo1.Header   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	MagneticField           *choreo3.Vector3  `protobuf:"bytes,2,opt,name=magnetic_field,json=magneticField" json:"magnetic_field,omitempty"`
	MagneticFieldCovariance []*choreo.Float64 `protobuf:"bytes,3,rep,name=magnetic_field_covariance,json=magneticFieldCovariance" json:"magnetic_field_covariance,omitempty"`
}

func (m *MagneticField) Reset()                    { *m = MagneticField{} }
func (m *MagneticField) String() string            { return proto.CompactTextString(m) }
func (*MagneticField) ProtoMessage()               {}
func (*MagneticField) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MagneticField) GetHeader() *choreo1.Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *MagneticField) GetMagneticField() *choreo3.Vector3 {
	if m != nil {
		return m.MagneticField
	}
	return nil
}

func (m *MagneticField) GetMagneticFieldCovariance() []*choreo.Float64 {
	if m != nil {
		return m.MagneticFieldCovariance
	}
	return nil
}

type NavSatStatus struct {
	NavSatStatus  NavSatStatus_Status  `protobuf:"varint,1,opt,name=nav_sat_status,json=navSatStatus,enum=choreo.NavSatStatus_Status" json:"nav_sat_status,omitempty"`
	NavSatService NavSatStatus_Service `protobuf:"varint,2,opt,name=nav_sat_service,json=navSatService,enum=choreo.NavSatStatus_Service" json:"nav_sat_service,omitempty"`
}

func (m *NavSatStatus) Reset()                    { *m = NavSatStatus{} }
func (m *NavSatStatus) String() string            { return proto.CompactTextString(m) }
func (*NavSatStatus) ProtoMessage()               {}
func (*NavSatStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *NavSatStatus) GetNavSatStatus() NavSatStatus_Status {
	if m != nil {
		return m.NavSatStatus
	}
	return NavSatStatus_NO_FIX
}

func (m *NavSatStatus) GetNavSatService() NavSatStatus_Service {
	if m != nil {
		return m.NavSatService
	}
	return NavSatStatus_GPS
}

type NavSatFix struct {
	Header                 *choreo1.Header                  `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	NavSatStatus           *NavSatStatus                    `protobuf:"bytes,2,opt,name=nav_sat_status,json=navSatStatus" json:"nav_sat_status,omitempty"`
	Latitude               *choreo.Float64                  `protobuf:"bytes,3,opt,name=latitude" json:"latitude,omitempty"`
	Longitude              *choreo.Float64                  `protobuf:"bytes,4,opt,name=longitude" json:"longitude,omitempty"`
	Altitude               *choreo.Float64                  `protobuf:"bytes,5,opt,name=altitude" json:"altitude,omitempty"`
	PositionCovariance     []*choreo.Float64                `protobuf:"bytes,6,rep,name=position_covariance,json=positionCovariance" json:"position_covariance,omitempty"`
	PositionCovarianceType NavSatFix_PositionCovarianceType `protobuf:"varint,7,opt,name=position_covariance_type,json=positionCovarianceType,enum=choreo.NavSatFix_PositionCovarianceType" json:"position_covariance_type,omitempty"`
}

func (m *NavSatFix) Reset()                    { *m = NavSatFix{} }
func (m *NavSatFix) String() string            { return proto.CompactTextString(m) }
func (*NavSatFix) ProtoMessage()               {}
func (*NavSatFix) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *NavSatFix) GetHeader() *choreo1.Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *NavSatFix) GetNavSatStatus() *NavSatStatus {
	if m != nil {
		return m.NavSatStatus
	}
	return nil
}

func (m *NavSatFix) GetLatitude() *choreo.Float64 {
	if m != nil {
		return m.Latitude
	}
	return nil
}

func (m *NavSatFix) GetLongitude() *choreo.Float64 {
	if m != nil {
		return m.Longitude
	}
	return nil
}

func (m *NavSatFix) GetAltitude() *choreo.Float64 {
	if m != nil {
		return m.Altitude
	}
	return nil
}

func (m *NavSatFix) GetPositionCovariance() []*choreo.Float64 {
	if m != nil {
		return m.PositionCovariance
	}
	return nil
}

func (m *NavSatFix) GetPositionCovarianceType() NavSatFix_PositionCovarianceType {
	if m != nil {
		return m.PositionCovarianceType
	}
	return NavSatFix_UNKNOWN
}

func init() {
	proto.RegisterType((*IMU)(nil), "choreo.IMU")
	proto.RegisterType((*MagneticField)(nil), "choreo.MagneticField")
	proto.RegisterType((*NavSatStatus)(nil), "choreo.NavSatStatus")
	proto.RegisterType((*NavSatFix)(nil), "choreo.NavSatFix")
	proto.RegisterEnum("choreo.NavSatStatus_Status", NavSatStatus_Status_name, NavSatStatus_Status_value)
	proto.RegisterEnum("choreo.NavSatStatus_Service", NavSatStatus_Service_name, NavSatStatus_Service_value)
	proto.RegisterEnum("choreo.NavSatFix_PositionCovarianceType", NavSatFix_PositionCovarianceType_name, NavSatFix_PositionCovarianceType_value)
}

func init() { proto.RegisterFile("sensor/imu.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 695 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x5d, 0x4f, 0xe2, 0x4c,
	0x14, 0xc7, 0x85, 0x2a, 0xe8, 0x01, 0xb1, 0x99, 0x47, 0x79, 0xf0, 0x25, 0x4f, 0x0c, 0x17, 0x4f,
	0xdc, 0x18, 0x21, 0x41, 0x63, 0xa2, 0x7b, 0xb3, 0xf5, 0x05, 0x96, 0x08, 0x94, 0x6d, 0xd5, 0x35,
	0x7b, 0xd3, 0x0c, 0x65, 0xac, 0x93, 0xb4, 0x1d, 0xb6, 0x9d, 0x36, 0xfa, 0x21, 0xf6, 0x1b, 0xed,
	0x97, 0xd8, 0x6f, 0xb3, 0x97, 0x9b, 0xbe, 0x41, 0x81, 0x6e, 0xe2, 0x15, 0x3d, 0x67, 0xfe, 0xe7,
	0xc7, 0xff, 0xcc, 0xe9, 0x29, 0x88, 0x2e, 0xb1, 0x5d, 0xe6, 0x34, 0xa9, 0xe5, 0x35, 0x26, 0x0e,
	0xe3, 0x0c, 0x15, 0xf4, 0x17, 0xe6, 0x10, 0xb6, 0xb7, 0x33, 0x71, 0xa8, 0x45, 0x39, 0xf5, 0x49,
	0xf3, 0xd9, 0x64, 0x98, 0x47, 0xc7, 0x7b, 0xd5, 0x59, 0xfa, 0x85, 0xe0, 0x31, 0x71, 0xe2, 0xfc,
	0x9e, 0x41, 0x98, 0x45, 0xb8, 0x43, 0xf5, 0xe6, 0x77, 0x0f, 0x73, 0xe2, 0xd8, 0x94, 0xd9, 0x49,
	0xcd, 0xec, 0xcc, 0x27, 0x3a, 0x67, 0x71, 0x4d, 0xfd, 0x97, 0x00, 0x42, 0xb7, 0xff, 0x80, 0xfe,
	0x87, 0x42, 0xc4, 0xaa, 0xe5, 0x0e, 0x73, 0x47, 0xa5, 0x56, 0xa5, 0x11, 0x79, 0x68, 0x7c, 0x0e,
	0xb3, 0x4a, 0x7c, 0x8a, 0xce, 0xa0, 0xc4, 0x1c, 0x4a, 0x6c, 0x8e, 0x39, 0x65, 0x76, 0x2d, 0x1f,
	0x8a, 0x51, 0x22, 0xfe, 0x32, 0xfd, 0x5b, 0x25, 0x2d, 0x43, 0x6d, 0xa8, 0xa6, 0x42, 0x4d, 0x67,
	0x3e, 0x76, 0x28, 0xb6, 0x75, 0x52, 0x13, 0x0e, 0x85, 0xa3, 0x52, 0x6b, 0x2b, 0x01, 0xb4, 0x83,
	0x36, 0xcf, 0xcf, 0x94, 0x9d, 0x94, 0xfc, 0x7a, 0xaa, 0x46, 0x97, 0x20, 0x62, 0xdb, 0xf0, 0x4c,
	0xec, 0x68, 0x3e, 0x31, 0x99, 0x4e, 0xf9, 0x5b, 0x6d, 0x35, 0xb4, 0x30, 0x25, 0x3c, 0x86, 0xdd,
	0x9d, 0x2a, 0x5b, 0xb1, 0xf0, 0x31, 0xd6, 0x21, 0x19, 0xf6, 0x17, 0x6b, 0xd3, 0x46, 0xd6, 0xb2,
	0x8d, 0xec, 0x2e, 0x60, 0x52, 0x66, 0x3e, 0xc1, 0x3f, 0x26, 0xb5, 0x09, 0x76, 0x34, 0xac, 0xeb,
	0xc4, 0x24, 0x4e, 0x74, 0x25, 0x85, 0x6c, 0x3f, 0x28, 0xd2, 0x4a, 0x29, 0x29, 0x7a, 0x80, 0xff,
	0x32, 0x08, 0x69, 0x57, 0xc5, 0x6c, 0x57, 0x07, 0xcb, 0xb0, 0x99, 0xb1, 0xfa, 0xcf, 0x1c, 0x6c,
	0xf6, 0xb1, 0x61, 0x13, 0x4e, 0xf5, 0x36, 0x25, 0xe6, 0xf8, 0xdd, 0xd3, 0x3d, 0x87, 0x8a, 0x15,
	0x17, 0x6a, 0xcf, 0x41, 0x65, 0x3c, 0xe0, 0xa5, 0x6e, 0x36, 0xad, 0x39, 0xfe, 0x1d, 0xec, 0xce,
	0xd7, 0xbd, 0x63, 0xc4, 0xff, 0xce, 0x21, 0x52, 0xf6, 0x7f, 0xe4, 0xa1, 0x3c, 0xc0, 0xbe, 0x8a,
	0xb9, 0xca, 0x31, 0xf7, 0x5c, 0x24, 0x41, 0xc5, 0xc6, 0xbe, 0xe6, 0x62, 0xae, 0xb9, 0x61, 0x26,
	0xec, 0xa2, 0xd2, 0xda, 0x4f, 0x90, 0x69, 0x75, 0x23, 0xfa, 0x51, 0xca, 0x76, 0x1a, 0x71, 0x03,
	0x5b, 0x53, 0x04, 0x71, 0x7c, 0xaa, 0x93, 0xb0, 0xb3, 0x4a, 0xeb, 0x20, 0x9b, 0x11, 0x69, 0x94,
	0xcd, 0x18, 0x12, 0x85, 0xf5, 0x0b, 0x28, 0xc4, 0x3c, 0x80, 0xc2, 0x40, 0xd6, 0xda, 0xdd, 0x27,
	0x71, 0x05, 0x15, 0x41, 0x08, 0x1e, 0x72, 0xa8, 0x0c, 0xeb, 0xea, 0x95, 0xa4, 0x86, 0xe9, 0x7c,
	0x10, 0x75, 0x92, 0x48, 0xa8, 0x5f, 0x40, 0x31, 0xa6, 0x04, 0xfa, 0xce, 0x50, 0x15, 0x57, 0x50,
	0x09, 0x8a, 0x9d, 0x9e, 0x3c, 0x90, 0x54, 0x55, 0xcc, 0x05, 0xc1, 0xb5, 0xdc, 0x1f, 0x06, 0x41,
	0x3e, 0x3c, 0x91, 0x7a, 0xdd, 0xde, 0xad, 0x2c, 0x0a, 0xf5, 0xdf, 0x02, 0x6c, 0x44, 0xee, 0xda,
	0xf4, 0xf5, 0xdd, 0xa3, 0xbc, 0x5c, 0xba, 0xb4, 0x68, 0x94, 0xdb, 0x59, 0x0d, 0x2f, 0xdc, 0xd6,
	0x31, 0xac, 0x9b, 0x98, 0x53, 0xee, 0x8d, 0x83, 0xe9, 0xe5, 0xb2, 0xa6, 0x37, 0x15, 0xa0, 0x13,
	0xd8, 0x30, 0x99, 0x6d, 0x44, 0xea, 0xd5, 0x6c, 0xf5, 0x4c, 0x11, 0xb0, 0xb1, 0x19, 0xb3, 0xd7,
	0xfe, 0xc2, 0x4e, 0x04, 0xc1, 0x8a, 0x4d, 0x98, 0x4b, 0x17, 0xb7, 0xa2, 0x90, 0xfd, 0x46, 0xa1,
	0x44, 0x9b, 0x5a, 0xd2, 0x11, 0xd4, 0x32, 0x08, 0x1a, 0x7f, 0x9b, 0x04, 0xcb, 0x15, 0xbc, 0x01,
	0x47, 0xf3, 0x17, 0xd2, 0xa6, 0xaf, 0x8d, 0xe1, 0x12, 0xe7, 0xfe, 0x6d, 0x42, 0x94, 0xea, 0x24,
	0x33, 0x5f, 0x7f, 0x82, 0x6a, 0x76, 0x45, 0x30, 0xc7, 0x87, 0xc1, 0xdd, 0x40, 0xfe, 0x3a, 0x10,
	0x57, 0x90, 0x08, 0x65, 0x69, 0x38, 0x54, 0xe4, 0xa7, 0x6e, 0x5f, 0xba, 0xbf, 0xbd, 0x11, 0x73,
	0x68, 0x1b, 0xc4, 0x9b, 0xae, 0xd4, 0x91, 0x07, 0x52, 0x4f, 0x4b, 0x74, 0x79, 0xb4, 0x01, 0x6b,
	0xd1, 0xa3, 0x70, 0x75, 0xfc, 0xed, 0x83, 0x41, 0xf9, 0x8b, 0x37, 0x6a, 0xe8, 0xcc, 0x6a, 0x2a,
	0x6c, 0xc4, 0xb8, 0xca, 0xbd, 0x31, 0x65, 0xcd, 0xc8, 0xf3, 0x89, 0xe5, 0x1a, 0x4d, 0xcb, 0x35,
	0x3e, 0x5a, 0xae, 0x31, 0x2a, 0x84, 0x5f, 0xf4, 0xd3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x67,
	0x06, 0xb6, 0xf0, 0x50, 0x06, 0x00, 0x00,
}
