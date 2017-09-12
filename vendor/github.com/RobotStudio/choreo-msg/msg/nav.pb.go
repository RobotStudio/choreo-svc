// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nav.proto

package rs_choreo_msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import choreo3 "github.com/RobotStudio/choreo-msg/msg/primitive"
import choreo4 "github.com/RobotStudio/choreo-msg/msg/primitive"
import choreo "github.com/RobotStudio/choreo-msg/msg/primitive"
import choreo5 "github.com/RobotStudio/choreo-msg/msg/primitive"
import choreo1 "github.com/RobotStudio/choreo-msg/msg/primitive"
import choreo6 "github.com/RobotStudio/choreo-msg/msg/geometric"
import choreo8 "github.com/RobotStudio/choreo-msg/msg/geometric"
import choreo10 "github.com/RobotStudio/choreo-msg/msg/geometric"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GridCells struct {
	Header     *choreo.Header   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	CellWidth  *choreo4.Float64 `protobuf:"bytes,2,opt,name=cell_width,json=cellWidth" json:"cell_width,omitempty"`
	CellHeight *choreo4.Float64 `protobuf:"bytes,3,opt,name=cell_height,json=cellHeight" json:"cell_height,omitempty"`
	Cells      []*choreo6.Point `protobuf:"bytes,4,rep,name=cells" json:"cells,omitempty"`
}

func (m *GridCells) Reset()                    { *m = GridCells{} }
func (m *GridCells) String() string            { return proto.CompactTextString(m) }
func (*GridCells) ProtoMessage()               {}
func (*GridCells) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *GridCells) GetHeader() *choreo.Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *GridCells) GetCellWidth() *choreo4.Float64 {
	if m != nil {
		return m.CellWidth
	}
	return nil
}

func (m *GridCells) GetCellHeight() *choreo4.Float64 {
	if m != nil {
		return m.CellHeight
	}
	return nil
}

func (m *GridCells) GetCells() []*choreo6.Point {
	if m != nil {
		return m.Cells
	}
	return nil
}

type MapMetaData struct {
	MapLoadTime *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=map_load_time,json=mapLoadTime" json:"map_load_time,omitempty"`
	Resolution  *choreo4.Float32           `protobuf:"bytes,2,opt,name=resolution" json:"resolution,omitempty"`
	Width       *choreo5.UInt32            `protobuf:"bytes,3,opt,name=width" json:"width,omitempty"`
	Height      *choreo5.UInt32            `protobuf:"bytes,4,opt,name=height" json:"height,omitempty"`
	Origin      *choreo8.Pose              `protobuf:"bytes,5,opt,name=origin" json:"origin,omitempty"`
}

func (m *MapMetaData) Reset()                    { *m = MapMetaData{} }
func (m *MapMetaData) String() string            { return proto.CompactTextString(m) }
func (*MapMetaData) ProtoMessage()               {}
func (*MapMetaData) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *MapMetaData) GetMapLoadTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.MapLoadTime
	}
	return nil
}

func (m *MapMetaData) GetResolution() *choreo4.Float32 {
	if m != nil {
		return m.Resolution
	}
	return nil
}

func (m *MapMetaData) GetWidth() *choreo5.UInt32 {
	if m != nil {
		return m.Width
	}
	return nil
}

func (m *MapMetaData) GetHeight() *choreo5.UInt32 {
	if m != nil {
		return m.Height
	}
	return nil
}

func (m *MapMetaData) GetOrigin() *choreo8.Pose {
	if m != nil {
		return m.Origin
	}
	return nil
}

type OccupancyGrid struct {
	Header *choreo.Header `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Info   *MapMetaData   `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
	Data   *choreo3.Bytes `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (m *OccupancyGrid) Reset()                    { *m = OccupancyGrid{} }
func (m *OccupancyGrid) String() string            { return proto.CompactTextString(m) }
func (*OccupancyGrid) ProtoMessage()               {}
func (*OccupancyGrid) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *OccupancyGrid) GetHeader() *choreo.Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *OccupancyGrid) GetInfo() *MapMetaData {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *OccupancyGrid) GetData() *choreo3.Bytes {
	if m != nil {
		return m.Data
	}
	return nil
}

type Odometry struct {
	Header       *choreo.Header                `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ChildFrameId *choreo1.String               `protobuf:"bytes,2,opt,name=child_frame_id,json=childFrameId" json:"child_frame_id,omitempty"`
	Pose         *choreo8.PoseWithCovariance   `protobuf:"bytes,3,opt,name=pose" json:"pose,omitempty"`
	Twist        *choreo10.TwistWithCovariance `protobuf:"bytes,4,opt,name=twist" json:"twist,omitempty"`
}

func (m *Odometry) Reset()                    { *m = Odometry{} }
func (m *Odometry) String() string            { return proto.CompactTextString(m) }
func (*Odometry) ProtoMessage()               {}
func (*Odometry) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *Odometry) GetHeader() *choreo.Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Odometry) GetChildFrameId() *choreo1.String {
	if m != nil {
		return m.ChildFrameId
	}
	return nil
}

func (m *Odometry) GetPose() *choreo8.PoseWithCovariance {
	if m != nil {
		return m.Pose
	}
	return nil
}

func (m *Odometry) GetTwist() *choreo10.TwistWithCovariance {
	if m != nil {
		return m.Twist
	}
	return nil
}

type Path struct {
	Header *choreo.Header  `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Poses  []*choreo8.Pose `protobuf:"bytes,2,rep,name=poses" json:"poses,omitempty"`
}

func (m *Path) Reset()                    { *m = Path{} }
func (m *Path) String() string            { return proto.CompactTextString(m) }
func (*Path) ProtoMessage()               {}
func (*Path) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *Path) GetHeader() *choreo.Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Path) GetPoses() []*choreo8.Pose {
	if m != nil {
		return m.Poses
	}
	return nil
}

type GetPlanRequest struct {
	Start     *choreo8.PoseStamped `protobuf:"bytes,1,opt,name=start" json:"start,omitempty"`
	Goal      *choreo8.PoseStamped `protobuf:"bytes,2,opt,name=goal" json:"goal,omitempty"`
	Tolerance *choreo4.Float32     `protobuf:"bytes,3,opt,name=tolerance" json:"tolerance,omitempty"`
}

func (m *GetPlanRequest) Reset()                    { *m = GetPlanRequest{} }
func (m *GetPlanRequest) String() string            { return proto.CompactTextString(m) }
func (*GetPlanRequest) ProtoMessage()               {}
func (*GetPlanRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *GetPlanRequest) GetStart() *choreo8.PoseStamped {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *GetPlanRequest) GetGoal() *choreo8.PoseStamped {
	if m != nil {
		return m.Goal
	}
	return nil
}

func (m *GetPlanRequest) GetTolerance() *choreo4.Float32 {
	if m != nil {
		return m.Tolerance
	}
	return nil
}

type SetMapRequest struct {
	Map         *OccupancyGrid                     `protobuf:"bytes,1,opt,name=map" json:"map,omitempty"`
	InitialPose *choreo8.PoseWithCovarianceStamped `protobuf:"bytes,2,opt,name=initial_pose,json=initialPose" json:"initial_pose,omitempty"`
}

func (m *SetMapRequest) Reset()                    { *m = SetMapRequest{} }
func (m *SetMapRequest) String() string            { return proto.CompactTextString(m) }
func (*SetMapRequest) ProtoMessage()               {}
func (*SetMapRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *SetMapRequest) GetMap() *OccupancyGrid {
	if m != nil {
		return m.Map
	}
	return nil
}

func (m *SetMapRequest) GetInitialPose() *choreo8.PoseWithCovarianceStamped {
	if m != nil {
		return m.InitialPose
	}
	return nil
}

func init() {
	proto.RegisterType((*GridCells)(nil), "choreo.GridCells")
	proto.RegisterType((*MapMetaData)(nil), "choreo.MapMetaData")
	proto.RegisterType((*OccupancyGrid)(nil), "choreo.OccupancyGrid")
	proto.RegisterType((*Odometry)(nil), "choreo.Odometry")
	proto.RegisterType((*Path)(nil), "choreo.Path")
	proto.RegisterType((*GetPlanRequest)(nil), "choreo.GetPlanRequest")
	proto.RegisterType((*SetMapRequest)(nil), "choreo.SetMapRequest")
}

func init() { proto.RegisterFile("nav.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 668 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcb, 0x4e, 0xdb, 0x4e,
	0x14, 0xc6, 0x65, 0x72, 0xd1, 0x3f, 0x27, 0x84, 0xbf, 0x34, 0x94, 0x2a, 0x4a, 0x17, 0x05, 0x17,
	0x15, 0xba, 0xc0, 0x6e, 0x13, 0xda, 0x4d, 0xa5, 0x2e, 0x00, 0x71, 0x91, 0x8a, 0x40, 0x13, 0x2a,
	0xa4, 0x6e, 0xa2, 0x89, 0x3d, 0xb1, 0x47, 0xb2, 0x3d, 0xee, 0xcc, 0x04, 0xc4, 0xa6, 0x9b, 0x3e,
	0x42, 0xdf, 0xa6, 0xaf, 0xd0, 0xb7, 0xe9, 0x13, 0x54, 0x73, 0x31, 0x09, 0x90, 0x4a, 0x59, 0xe6,
	0x3b, 0xbf, 0x33, 0x39, 0xdf, 0xf9, 0x8e, 0xa1, 0x55, 0x90, 0x9b, 0xa0, 0x14, 0x5c, 0x71, 0xd4,
	0x8c, 0x52, 0x2e, 0x28, 0xef, 0xbd, 0x4c, 0x38, 0x4f, 0x32, 0x1a, 0x1a, 0x75, 0x3c, 0x9d, 0x84,
	0x8a, 0xe5, 0x54, 0x2a, 0x92, 0x97, 0x16, 0xec, 0x6d, 0x94, 0x82, 0xe5, 0x4c, 0xb1, 0x1b, 0x1a,
	0x8e, 0xef, 0x14, 0x95, 0x4f, 0xe5, 0x49, 0xc6, 0x89, 0x72, 0xf2, 0xf3, 0x99, 0x9c, 0x52, 0x12,
	0x53, 0xe1, 0xf4, 0xf5, 0x99, 0xce, 0x8a, 0x05, 0xb0, 0x54, 0x82, 0x15, 0x49, 0xf5, 0x76, 0x42,
	0x79, 0x4e, 0x95, 0x60, 0x51, 0x58, 0xf2, 0x19, 0xfe, 0x6c, 0x5e, 0x96, 0xf4, 0x29, 0xac, 0x6e,
	0x99, 0x74, 0xb0, 0xff, 0xcb, 0x83, 0xd6, 0x89, 0x60, 0xf1, 0x21, 0xcd, 0x32, 0x89, 0x5e, 0x43,
	0xd3, 0x8e, 0xd3, 0xf5, 0x36, 0xbd, 0xdd, 0x76, 0x7f, 0x2d, 0xb0, 0xf6, 0x83, 0x53, 0xa3, 0x62,
	0x57, 0x45, 0x01, 0x40, 0x44, 0xb3, 0x6c, 0x74, 0xcb, 0x62, 0x95, 0x76, 0x57, 0x0c, 0xfb, 0x7f,
	0xc5, 0x1e, 0x6b, 0x9f, 0x1f, 0xf6, 0x71, 0x4b, 0x23, 0xd7, 0x9a, 0x40, 0x6f, 0xa1, 0x6d, 0xf8,
	0x94, 0xb2, 0x24, 0x55, 0xdd, 0xda, 0xe2, 0x06, 0xf3, 0xe6, 0xa9, 0x41, 0xd0, 0x2b, 0x68, 0xe8,
	0x5f, 0xb2, 0x5b, 0xdf, 0xac, 0xed, 0xb6, 0xfb, 0x9d, 0x8a, 0xbd, 0xd4, 0x46, 0xb1, 0xad, 0xf9,
	0x7f, 0x3c, 0x68, 0x9f, 0x93, 0xf2, 0x9c, 0x2a, 0x72, 0x44, 0x14, 0x41, 0x9f, 0xa0, 0x93, 0x93,
	0x72, 0x94, 0x71, 0x12, 0x8f, 0x74, 0x3e, 0xce, 0x45, 0x2f, 0xb0, 0xe1, 0x05, 0x55, 0x78, 0xc1,
	0x55, 0x15, 0x1e, 0x6e, 0xe7, 0xa4, 0xfc, 0xcc, 0x49, 0xac, 0x15, 0x14, 0x02, 0x08, 0x2a, 0x79,
	0x36, 0x55, 0x8c, 0x17, 0x0b, 0x6d, 0x0d, 0xfa, 0x78, 0x0e, 0x41, 0xdb, 0xd0, 0xb0, 0x2b, 0xa8,
	0x3d, 0x5c, 0xd7, 0x97, 0xb3, 0x42, 0xa3, 0xb6, 0x68, 0xb7, 0x6a, 0x8c, 0xd7, 0x17, 0x62, 0xae,
	0x8a, 0xb6, 0xa1, 0xc9, 0x05, 0x4b, 0x58, 0xd1, 0x6d, 0x18, 0x6e, 0x75, 0x66, 0x5a, 0x52, 0xec,
	0x6a, 0xfe, 0x0f, 0x0f, 0x3a, 0x17, 0x51, 0x34, 0x2d, 0x49, 0x11, 0xdd, 0xe9, 0xe8, 0x96, 0x4e,
	0x6d, 0x07, 0xea, 0xac, 0x98, 0x70, 0x67, 0x6c, 0xbd, 0xa2, 0xe6, 0x36, 0x88, 0x0d, 0x80, 0xb6,
	0xa0, 0x1e, 0x13, 0x45, 0x9c, 0xab, 0xfb, 0xdd, 0x1f, 0xe8, 0xbb, 0xc6, 0xa6, 0xe4, 0xff, 0xf6,
	0xe0, 0xbf, 0x8b, 0xd8, 0x5c, 0xd4, 0xdd, 0xd2, 0x03, 0xec, 0xc3, 0x5a, 0x94, 0xb2, 0x2c, 0x1e,
	0x4d, 0x04, 0xc9, 0xe9, 0x88, 0xc5, 0x6e, 0x94, 0x7b, 0x7e, 0x68, 0xce, 0x1b, 0xaf, 0x1a, 0xea,
	0x58, 0x43, 0x67, 0x31, 0x0a, 0xa0, 0xae, 0xef, 0xd8, 0x4d, 0xd3, 0x9b, 0x5f, 0xca, 0x35, 0x53,
	0xe9, 0x21, 0xbf, 0x21, 0x82, 0x91, 0x22, 0xa2, 0xd8, 0x70, 0xe8, 0x1d, 0x34, 0xcc, 0x85, 0xbb,
	0x6d, 0xbf, 0xa8, 0x1a, 0xae, 0xb4, 0xf8, 0xa8, 0xc3, 0x92, 0x3e, 0x86, 0xfa, 0x25, 0xa9, 0x92,
	0x5a, 0xc2, 0x88, 0x0f, 0x0d, 0xfd, 0x57, 0xb2, 0xbb, 0x62, 0xae, 0xf3, 0x61, 0x50, 0xb6, 0xe4,
	0xff, 0xf4, 0x60, 0xed, 0x84, 0xaa, 0xcb, 0x8c, 0x14, 0x98, 0x7e, 0x9b, 0x52, 0xa9, 0xd0, 0x1b,
	0x68, 0x48, 0x45, 0x84, 0x72, 0xaf, 0xaf, 0xcf, 0xb7, 0x0d, 0xf5, 0x39, 0xd2, 0x18, 0x5b, 0x42,
	0x67, 0x95, 0x70, 0x92, 0x3d, 0xce, 0x6a, 0x9e, 0x34, 0x00, 0xda, 0x83, 0x96, 0xe2, 0x19, 0x15,
	0xda, 0xce, 0xc2, 0x0f, 0x6b, 0xd0, 0xc7, 0x33, 0xc2, 0xff, 0x0e, 0x9d, 0x21, 0x55, 0xe7, 0xa4,
	0xac, 0x66, 0xda, 0x81, 0x5a, 0x4e, 0x4a, 0x37, 0xd1, 0x46, 0xd5, 0xf9, 0xe0, 0xc0, 0xb0, 0x26,
	0xd0, 0x11, 0xac, 0xb2, 0x82, 0x29, 0x46, 0xb2, 0x91, 0x89, 0xc3, 0x4e, 0xb6, 0xf5, 0xef, 0x38,
	0xaa, 0x39, 0xdb, 0xae, 0x4d, 0x13, 0x07, 0xef, 0xbf, 0x0e, 0x12, 0xa6, 0xd2, 0xe9, 0x38, 0x88,
	0x78, 0x1e, 0x62, 0x3e, 0xe6, 0x6a, 0xa8, 0xa6, 0x31, 0xe3, 0xa1, 0x7d, 0x67, 0x2f, 0x97, 0x49,
	0x98, 0xcb, 0xe4, 0xa3, 0x90, 0xd5, 0xcb, 0xb9, 0x4c, 0xc6, 0x4d, 0xf3, 0xe9, 0x0e, 0xfe, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x4b, 0xe5, 0xb5, 0xcf, 0x9a, 0x05, 0x00, 0x00,
}
