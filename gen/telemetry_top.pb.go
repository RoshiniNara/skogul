// Code generated by protoc-gen-go. DO NOT EDIT.
// source: telemetry_top.proto

package gen

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TelemetryFieldOptions struct {
	IsKey                *bool    `protobuf:"varint,1,opt,name=is_key,json=isKey" json:"is_key,omitempty"`
	IsTimestamp          *bool    `protobuf:"varint,2,opt,name=is_timestamp,json=isTimestamp" json:"is_timestamp,omitempty"`
	IsCounter            *bool    `protobuf:"varint,3,opt,name=is_counter,json=isCounter" json:"is_counter,omitempty"`
	IsGauge              *bool    `protobuf:"varint,4,opt,name=is_gauge,json=isGauge" json:"is_gauge,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TelemetryFieldOptions) Reset()         { *m = TelemetryFieldOptions{} }
func (m *TelemetryFieldOptions) String() string { return proto.CompactTextString(m) }
func (*TelemetryFieldOptions) ProtoMessage()    {}
func (*TelemetryFieldOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_telemetry_top_3f8148ea0c6fa289, []int{0}
}
func (m *TelemetryFieldOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetryFieldOptions.Unmarshal(m, b)
}
func (m *TelemetryFieldOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetryFieldOptions.Marshal(b, m, deterministic)
}
func (dst *TelemetryFieldOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetryFieldOptions.Merge(dst, src)
}
func (m *TelemetryFieldOptions) XXX_Size() int {
	return xxx_messageInfo_TelemetryFieldOptions.Size(m)
}
func (m *TelemetryFieldOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetryFieldOptions.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetryFieldOptions proto.InternalMessageInfo

func (m *TelemetryFieldOptions) GetIsKey() bool {
	if m != nil && m.IsKey != nil {
		return *m.IsKey
	}
	return false
}

func (m *TelemetryFieldOptions) GetIsTimestamp() bool {
	if m != nil && m.IsTimestamp != nil {
		return *m.IsTimestamp
	}
	return false
}

func (m *TelemetryFieldOptions) GetIsCounter() bool {
	if m != nil && m.IsCounter != nil {
		return *m.IsCounter
	}
	return false
}

func (m *TelemetryFieldOptions) GetIsGauge() bool {
	if m != nil && m.IsGauge != nil {
		return *m.IsGauge
	}
	return false
}

type TelemetryStream struct {
	// router name or export IP address
	SystemId *string `protobuf:"bytes,1,req,name=system_id,json=systemId" json:"system_id,omitempty"`
	// line card / RE (slot number). For RE, it will be 65535
	ComponentId *uint32 `protobuf:"varint,2,opt,name=component_id,json=componentId" json:"component_id,omitempty"`
	// PFE (if applicable)
	SubComponentId *uint32 `protobuf:"varint,3,opt,name=sub_component_id,json=subComponentId" json:"sub_component_id,omitempty"`
	// Overload sensor name with "senor name, internal path, external path
	// and component" seperated by ":". For RE sensors, component will be
	// daemon-name and for PFE sensors it will be "PFE".
	SensorName *string `protobuf:"bytes,4,opt,name=sensor_name,json=sensorName" json:"sensor_name,omitempty"`
	// sequence number, monotonically increasesing for each
	// system_id, component_id, sub_component_id + sensor_name.
	SequenceNumber *uint32 `protobuf:"varint,5,opt,name=sequence_number,json=sequenceNumber" json:"sequence_number,omitempty"`
	// timestamp (milliseconds since 00:00:00 UTC 1/1/1970)
	Timestamp *uint64 `protobuf:"varint,6,opt,name=timestamp" json:"timestamp,omitempty"`
	// major version
	VersionMajor *uint32 `protobuf:"varint,7,opt,name=version_major,json=versionMajor" json:"version_major,omitempty"`
	// minor version
	VersionMinor         *uint32            `protobuf:"varint,8,opt,name=version_minor,json=versionMinor" json:"version_minor,omitempty"`
	Ietf                 *IETFSensors       `protobuf:"bytes,100,opt,name=ietf" json:"ietf,omitempty"`
	Enterprise           *EnterpriseSensors `protobuf:"bytes,101,opt,name=enterprise" json:"enterprise,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TelemetryStream) Reset()         { *m = TelemetryStream{} }
func (m *TelemetryStream) String() string { return proto.CompactTextString(m) }
func (*TelemetryStream) ProtoMessage()    {}
func (*TelemetryStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_telemetry_top_3f8148ea0c6fa289, []int{1}
}
func (m *TelemetryStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetryStream.Unmarshal(m, b)
}
func (m *TelemetryStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetryStream.Marshal(b, m, deterministic)
}
func (dst *TelemetryStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetryStream.Merge(dst, src)
}
func (m *TelemetryStream) XXX_Size() int {
	return xxx_messageInfo_TelemetryStream.Size(m)
}
func (m *TelemetryStream) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetryStream.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetryStream proto.InternalMessageInfo

func (m *TelemetryStream) GetSystemId() string {
	if m != nil && m.SystemId != nil {
		return *m.SystemId
	}
	return ""
}

func (m *TelemetryStream) GetComponentId() uint32 {
	if m != nil && m.ComponentId != nil {
		return *m.ComponentId
	}
	return 0
}

func (m *TelemetryStream) GetSubComponentId() uint32 {
	if m != nil && m.SubComponentId != nil {
		return *m.SubComponentId
	}
	return 0
}

func (m *TelemetryStream) GetSensorName() string {
	if m != nil && m.SensorName != nil {
		return *m.SensorName
	}
	return ""
}

func (m *TelemetryStream) GetSequenceNumber() uint32 {
	if m != nil && m.SequenceNumber != nil {
		return *m.SequenceNumber
	}
	return 0
}

func (m *TelemetryStream) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *TelemetryStream) GetVersionMajor() uint32 {
	if m != nil && m.VersionMajor != nil {
		return *m.VersionMajor
	}
	return 0
}

func (m *TelemetryStream) GetVersionMinor() uint32 {
	if m != nil && m.VersionMinor != nil {
		return *m.VersionMinor
	}
	return 0
}

func (m *TelemetryStream) GetIetf() *IETFSensors {
	if m != nil {
		return m.Ietf
	}
	return nil
}

func (m *TelemetryStream) GetEnterprise() *EnterpriseSensors {
	if m != nil {
		return m.Enterprise
	}
	return nil
}

type IETFSensors struct {
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *IETFSensors) Reset()         { *m = IETFSensors{} }
func (m *IETFSensors) String() string { return proto.CompactTextString(m) }
func (*IETFSensors) ProtoMessage()    {}
func (*IETFSensors) Descriptor() ([]byte, []int) {
	return fileDescriptor_telemetry_top_3f8148ea0c6fa289, []int{2}
}

var extRange_IETFSensors = []proto.ExtensionRange{
	{Start: 1, End: 536870911},
}

func (*IETFSensors) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_IETFSensors
}
func (m *IETFSensors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IETFSensors.Unmarshal(m, b)
}
func (m *IETFSensors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IETFSensors.Marshal(b, m, deterministic)
}
func (dst *IETFSensors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IETFSensors.Merge(dst, src)
}
func (m *IETFSensors) XXX_Size() int {
	return xxx_messageInfo_IETFSensors.Size(m)
}
func (m *IETFSensors) XXX_DiscardUnknown() {
	xxx_messageInfo_IETFSensors.DiscardUnknown(m)
}

var xxx_messageInfo_IETFSensors proto.InternalMessageInfo

type EnterpriseSensors struct {
	XXX_NoUnkeyedLiteral         struct{} `jon:"-"`
	proto.XXX_InternalExtensions `jsn:"-"`
	XXX_unrecognized             []byte `jsn:"-"`
	XXX_sizecache                int32  `jsn:"-"`
}

func (m *EnterpriseSensors) Reset()         { *m = EnterpriseSensors{} }
func (m *EnterpriseSensors) String() string { return proto.CompactTextString(m) }
func (*EnterpriseSensors) ProtoMessage()    {}
func (*EnterpriseSensors) Descriptor() ([]byte, []int) {
	return fileDescriptor_telemetry_top_3f8148ea0c6fa289, []int{3}
}

var extRange_EnterpriseSensors = []proto.ExtensionRange{
	{Start: 1, End: 536870911},
}

func (*EnterpriseSensors) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_EnterpriseSensors
}
func (m *EnterpriseSensors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnterpriseSensors.Unmarshal(m, b)
}
func (m *EnterpriseSensors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnterpriseSensors.Marshal(b, m, deterministic)
}
func (dst *EnterpriseSensors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnterpriseSensors.Merge(dst, src)
}
func (m *EnterpriseSensors) XXX_Size() int {
	return xxx_messageInfo_EnterpriseSensors.Size(m)
}
func (m *EnterpriseSensors) XXX_DiscardUnknown() {
	xxx_messageInfo_EnterpriseSensors.DiscardUnknown(m)
}

var xxx_messageInfo_EnterpriseSensors proto.InternalMessageInfo

type JuniperNetworksSensors struct {
	XXX_NoUnkeyedLiteral         struct{} `jsn:"-"`
	proto.XXX_InternalExtensions `jon:"-"`
	XXX_unrecognized             []byte `jsn:"-"`
	XXX_sizecache                int32  `jsn:"-"`
}

func (m *JuniperNetworksSensors) Reset()         { *m = JuniperNetworksSensors{} }
func (m *JuniperNetworksSensors) String() string { return proto.CompactTextString(m) }
func (*JuniperNetworksSensors) ProtoMessage()    {}
func (*JuniperNetworksSensors) Descriptor() ([]byte, []int) {
	return fileDescriptor_telemetry_top_3f8148ea0c6fa289, []int{4}
}

var extRange_JuniperNetworksSensors = []proto.ExtensionRange{
	{Start: 1, End: 536870911},
}

func (*JuniperNetworksSensors) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_JuniperNetworksSensors
}
func (m *JuniperNetworksSensors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JuniperNetworksSensors.Unmarshal(m, b)
}
func (m *JuniperNetworksSensors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JuniperNetworksSensors.Marshal(b, m, deterministic)
}
func (dst *JuniperNetworksSensors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JuniperNetworksSensors.Merge(dst, src)
}
func (m *JuniperNetworksSensors) XXX_Size() int {
	return xxx_messageInfo_JuniperNetworksSensors.Size(m)
}
func (m *JuniperNetworksSensors) XXX_DiscardUnknown() {
	xxx_messageInfo_JuniperNetworksSensors.DiscardUnknown(m)
}

var xxx_messageInfo_JuniperNetworksSensors proto.InternalMessageInfo

var E_TelemetryOptions = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*TelemetryFieldOptions)(nil),
	Field:         1024,
	Name:          "telemetry_options",
	Tag:           "bytes,1024,opt,name=telemetry_options,json=telemetryOptions",
	Filename:      "telemetry_top.proto",
}

var E_JuniperNetworks = &proto.ExtensionDesc{
	ExtendedType:  (*EnterpriseSensors)(nil),
	ExtensionType: (*JuniperNetworksSensors)(nil),
	Field:         2636,
	Name:          "juniperNetworks",
	Tag:           "bytes,2636,opt,name=juniperNetworks",
	Filename:      "telemetry_top.proto",
}

func init() {
	proto.RegisterType((*TelemetryFieldOptions)(nil), "TelemetryFieldOptions")
	proto.RegisterType((*TelemetryStream)(nil), "TelemetryStream")
	proto.RegisterType((*IETFSensors)(nil), "IETFSensors")
	proto.RegisterType((*EnterpriseSensors)(nil), "EnterpriseSensors")
	proto.RegisterType((*JuniperNetworksSensors)(nil), "JuniperNetworksSensors")
	proto.RegisterExtension(E_TelemetryOptions)
	proto.RegisterExtension(E_JuniperNetworks)
}

func init() { proto.RegisterFile("telemetry_top.proto", fileDescriptor_telemetry_top_3f8148ea0c6fa289) }

var fileDescriptor_telemetry_top_3f8148ea0c6fa289 = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x6f, 0xda, 0x40,
	0x14, 0xc4, 0x65, 0xfe, 0x24, 0xe6, 0x41, 0x02, 0xd9, 0x36, 0xc9, 0xb6, 0x12, 0x92, 0xeb, 0x48,
	0x2d, 0xea, 0xc1, 0x48, 0x39, 0xe6, 0x54, 0x35, 0x4a, 0x2a, 0x5a, 0x95, 0x4a, 0x0e, 0xea, 0xd5,
	0x32, 0xf8, 0x05, 0x6d, 0x82, 0xbd, 0xee, 0xbe, 0x75, 0x2b, 0x6e, 0xa8, 0xd7, 0x7e, 0x8f, 0x7e,
	0x9a, 0x7e, 0xa8, 0xca, 0x6b, 0x0c, 0x26, 0xe1, 0xb8, 0x33, 0xbf, 0x37, 0x1a, 0xf0, 0xc0, 0x0b,
	0x8d, 0x0b, 0x8c, 0x51, 0xab, 0x65, 0xa0, 0x65, 0xea, 0xa5, 0x4a, 0x6a, 0xf9, 0xda, 0x99, 0x4b,
	0x39, 0x5f, 0xe0, 0xd0, 0xbc, 0xa6, 0xd9, 0xfd, 0x30, 0x42, 0x9a, 0x29, 0x91, 0x6a, 0xa9, 0x0a,
	0xc2, 0xfd, 0x63, 0xc1, 0xe9, 0xa4, 0xbc, 0xbc, 0x15, 0xb8, 0x88, 0xbe, 0xa5, 0x5a, 0xc8, 0x84,
	0xd8, 0x29, 0x1c, 0x08, 0x0a, 0x1e, 0x71, 0xc9, 0x2d, 0xc7, 0x1a, 0xd8, 0x7e, 0x53, 0xd0, 0x17,
	0x5c, 0xb2, 0x37, 0xd0, 0x11, 0x14, 0x68, 0x11, 0x23, 0xe9, 0x30, 0x4e, 0x79, 0xcd, 0x98, 0x6d,
	0x41, 0x93, 0x52, 0x62, 0x7d, 0x00, 0x41, 0xc1, 0x4c, 0x66, 0x89, 0x46, 0xc5, 0xeb, 0x06, 0x68,
	0x09, 0xba, 0x2e, 0x04, 0xf6, 0x0a, 0x6c, 0x41, 0xc1, 0x3c, 0xcc, 0xe6, 0xc8, 0x1b, 0xc6, 0x3c,
	0x14, 0xf4, 0x29, 0x7f, 0xba, 0x7f, 0xeb, 0xd0, 0xdd, 0xb4, 0xb9, 0xd3, 0x0a, 0xc3, 0x98, 0xb9,
	0xd0, 0xa2, 0x25, 0x69, 0x8c, 0x03, 0x11, 0x71, 0xcb, 0xa9, 0x0d, 0x5a, 0x1f, 0x9b, 0xbf, 0x3f,
	0xd4, 0x6c, 0xcb, 0xb7, 0x0b, 0x7d, 0x14, 0xb1, 0x01, 0x74, 0x66, 0x32, 0x4e, 0x65, 0x82, 0x89,
	0xce, 0xb1, 0xbc, 0xd4, 0x51, 0x89, 0xb5, 0x37, 0xd6, 0x28, 0x62, 0x43, 0xe8, 0x51, 0x36, 0x0d,
	0x76, 0xe8, 0x7a, 0x95, 0x3e, 0xa6, 0x6c, 0x7a, 0x5d, 0x39, 0x78, 0x0b, 0x6d, 0xc2, 0x84, 0xa4,
	0x0a, 0x92, 0x30, 0x2e, 0x0a, 0x6f, 0x0a, 0x40, 0xe1, 0x8c, 0xc3, 0x18, 0xd9, 0x3b, 0xe8, 0x12,
	0xfe, 0xc8, 0x30, 0x99, 0x61, 0x90, 0x64, 0xf1, 0x14, 0x15, 0x6f, 0xe6, 0xb9, 0xfe, 0x71, 0x29,
	0x8f, 0x8d, 0xca, 0x2e, 0xa0, 0xb5, 0xfd, 0xf7, 0x0e, 0x1c, 0x6b, 0xd0, 0x30, 0x71, 0x3d, 0xcb,
	0xdf, 0xea, 0xec, 0x02, 0x8e, 0x7e, 0xa2, 0x22, 0x21, 0x93, 0x20, 0x0e, 0x1f, 0xa4, 0xe2, 0x87,
	0x26, 0xab, 0xb3, 0x16, 0xbf, 0xe6, 0xda, 0x0e, 0x24, 0x12, 0xa9, 0xb8, 0xbd, 0x0b, 0xe5, 0x1a,
	0x73, 0xa0, 0x21, 0x50, 0xdf, 0xf3, 0xc8, 0xb1, 0x06, 0xed, 0xcb, 0x8e, 0x37, 0xba, 0x99, 0xdc,
	0xde, 0x99, 0xda, 0xe4, 0x1b, 0x87, 0x5d, 0x02, 0x60, 0xfe, 0x61, 0x52, 0x25, 0x08, 0x39, 0x1a,
	0x8e, 0x79, 0x37, 0x1b, 0xa9, 0xa4, 0x2b, 0x94, 0x7b, 0x0e, 0xed, 0x4a, 0xd0, 0x7b, 0xdb, 0xb6,
	0x7a, 0xab, 0xd5, 0x6a, 0x55, 0x73, 0xfb, 0x70, 0xf2, 0xec, 0xb2, 0x62, 0xbb, 0x70, 0xf6, 0x39,
	0x4b, 0x44, 0x8a, 0x6a, 0x8c, 0xfa, 0x97, 0x54, 0x8f, 0xf4, 0x8c, 0xb9, 0x9a, 0xc1, 0xc9, 0x76,
	0xcb, 0x72, 0xbd, 0xc6, 0xbe, 0x57, 0x4c, 0xd9, 0x2b, 0xa7, 0xec, 0x55, 0xc7, 0xca, 0x57, 0xb6,
	0xa9, 0x7d, 0xe6, 0xed, 0xdd, 0xb2, 0xdf, 0xdb, 0x04, 0xae, 0x95, 0xab, 0xef, 0xd0, 0x7d, 0xd8,
	0x2d, 0xc2, 0xf6, 0xfc, 0x66, 0xfe, 0xef, 0xa5, 0xc9, 0x3d, 0xf7, 0xf6, 0xb7, 0xf6, 0x9f, 0x86,
	0xfc, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x4e, 0xee, 0x3d, 0x87, 0x03, 0x00, 0x00,
}