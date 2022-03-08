// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eventd.proto

package telemetry

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type JunosEvents struct {
	Events               *JunosEventsEventsType `protobuf:"bytes,151,opt,name=events" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *JunosEvents) Reset()         { *m = JunosEvents{} }
func (m *JunosEvents) String() string { return proto.CompactTextString(m) }
func (*JunosEvents) ProtoMessage()    {}
func (*JunosEvents) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5807bd979efd0cf, []int{0}
}
func (m *JunosEvents) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosEvents.Unmarshal(m, b)
}
func (m *JunosEvents) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosEvents.Marshal(b, m, deterministic)
}
func (m *JunosEvents) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosEvents.Merge(m, src)
}
func (m *JunosEvents) XXX_Size() int {
	return xxx_messageInfo_JunosEvents.Size(m)
}
func (m *JunosEvents) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosEvents.DiscardUnknown(m)
}

var xxx_messageInfo_JunosEvents proto.InternalMessageInfo

func (m *JunosEvents) GetEvents() *JunosEventsEventsType {
	if m != nil {
		return m.Events
	}
	return nil
}

type JunosEventsEventsType struct {
	Event                []*JunosEventsEventsTypeEventList `protobuf:"bytes,151,rep,name=event" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *JunosEventsEventsType) Reset()         { *m = JunosEventsEventsType{} }
func (m *JunosEventsEventsType) String() string { return proto.CompactTextString(m) }
func (*JunosEventsEventsType) ProtoMessage()    {}
func (*JunosEventsEventsType) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5807bd979efd0cf, []int{0, 0}
}
func (m *JunosEventsEventsType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosEventsEventsType.Unmarshal(m, b)
}
func (m *JunosEventsEventsType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosEventsEventsType.Marshal(b, m, deterministic)
}
func (m *JunosEventsEventsType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosEventsEventsType.Merge(m, src)
}
func (m *JunosEventsEventsType) XXX_Size() int {
	return xxx_messageInfo_JunosEventsEventsType.Size(m)
}
func (m *JunosEventsEventsType) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosEventsEventsType.DiscardUnknown(m)
}

var xxx_messageInfo_JunosEventsEventsType proto.InternalMessageInfo

func (m *JunosEventsEventsType) GetEvent() []*JunosEventsEventsTypeEventList {
	if m != nil {
		return m.Event
	}
	return nil
}

type JunosEventsEventsTypeEventList struct {
	Id                   *string                                         `protobuf:"bytes,51,opt,name=id" json:"id,omitempty"`
	Type                 *string                                         `protobuf:"bytes,52,opt,name=type" json:"type,omitempty"`
	Timestamp            *JunosEventsEventsTypeEventListTimestampType    `protobuf:"bytes,151,opt,name=timestamp" json:"timestamp,omitempty"`
	Priority             *string                                         `protobuf:"bytes,53,opt,name=priority" json:"priority,omitempty"`
	Facility             *string                                         `protobuf:"bytes,54,opt,name=facility" json:"facility,omitempty"`
	Pid                  *uint32                                         `protobuf:"varint,55,opt,name=pid" json:"pid,omitempty"`
	Message              *string                                         `protobuf:"bytes,56,opt,name=message" json:"message,omitempty"`
	Daemon               *string                                         `protobuf:"bytes,57,opt,name=daemon" json:"daemon,omitempty"`
	Hostname             *string                                         `protobuf:"bytes,58,opt,name=hostname" json:"hostname,omitempty"`
	Lsname               *string                                         `protobuf:"bytes,59,opt,name=lsname" json:"lsname,omitempty"`
	Attributes           []*JunosEventsEventsTypeEventListAttributesList `protobuf:"bytes,152,rep,name=attributes" json:"attributes,omitempty"`
	Logoptions           *int32                                          `protobuf:"varint,60,opt,name=logoptions" json:"logoptions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                        `json:"-"`
	XXX_unrecognized     []byte                                          `json:"-"`
	XXX_sizecache        int32                                           `json:"-"`
}

func (m *JunosEventsEventsTypeEventList) Reset()         { *m = JunosEventsEventsTypeEventList{} }
func (m *JunosEventsEventsTypeEventList) String() string { return proto.CompactTextString(m) }
func (*JunosEventsEventsTypeEventList) ProtoMessage()    {}
func (*JunosEventsEventsTypeEventList) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5807bd979efd0cf, []int{0, 0, 0}
}
func (m *JunosEventsEventsTypeEventList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosEventsEventsTypeEventList.Unmarshal(m, b)
}
func (m *JunosEventsEventsTypeEventList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosEventsEventsTypeEventList.Marshal(b, m, deterministic)
}
func (m *JunosEventsEventsTypeEventList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosEventsEventsTypeEventList.Merge(m, src)
}
func (m *JunosEventsEventsTypeEventList) XXX_Size() int {
	return xxx_messageInfo_JunosEventsEventsTypeEventList.Size(m)
}
func (m *JunosEventsEventsTypeEventList) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosEventsEventsTypeEventList.DiscardUnknown(m)
}

var xxx_messageInfo_JunosEventsEventsTypeEventList proto.InternalMessageInfo

func (m *JunosEventsEventsTypeEventList) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *JunosEventsEventsTypeEventList) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *JunosEventsEventsTypeEventList) GetTimestamp() *JunosEventsEventsTypeEventListTimestampType {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *JunosEventsEventsTypeEventList) GetPriority() string {
	if m != nil && m.Priority != nil {
		return *m.Priority
	}
	return ""
}

func (m *JunosEventsEventsTypeEventList) GetFacility() string {
	if m != nil && m.Facility != nil {
		return *m.Facility
	}
	return ""
}

func (m *JunosEventsEventsTypeEventList) GetPid() uint32 {
	if m != nil && m.Pid != nil {
		return *m.Pid
	}
	return 0
}

func (m *JunosEventsEventsTypeEventList) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *JunosEventsEventsTypeEventList) GetDaemon() string {
	if m != nil && m.Daemon != nil {
		return *m.Daemon
	}
	return ""
}

func (m *JunosEventsEventsTypeEventList) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

func (m *JunosEventsEventsTypeEventList) GetLsname() string {
	if m != nil && m.Lsname != nil {
		return *m.Lsname
	}
	return ""
}

func (m *JunosEventsEventsTypeEventList) GetAttributes() []*JunosEventsEventsTypeEventListAttributesList {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *JunosEventsEventsTypeEventList) GetLogoptions() int32 {
	if m != nil && m.Logoptions != nil {
		return *m.Logoptions
	}
	return 0
}

type JunosEventsEventsTypeEventListTimestampType struct {
	Seconds              *uint32  `protobuf:"varint,51,opt,name=seconds" json:"seconds,omitempty"`
	Microseconds         *uint32  `protobuf:"varint,52,opt,name=microseconds" json:"microseconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JunosEventsEventsTypeEventListTimestampType) Reset() {
	*m = JunosEventsEventsTypeEventListTimestampType{}
}
func (m *JunosEventsEventsTypeEventListTimestampType) String() string {
	return proto.CompactTextString(m)
}
func (*JunosEventsEventsTypeEventListTimestampType) ProtoMessage() {}
func (*JunosEventsEventsTypeEventListTimestampType) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5807bd979efd0cf, []int{0, 0, 0, 0}
}
func (m *JunosEventsEventsTypeEventListTimestampType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosEventsEventsTypeEventListTimestampType.Unmarshal(m, b)
}
func (m *JunosEventsEventsTypeEventListTimestampType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosEventsEventsTypeEventListTimestampType.Marshal(b, m, deterministic)
}
func (m *JunosEventsEventsTypeEventListTimestampType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosEventsEventsTypeEventListTimestampType.Merge(m, src)
}
func (m *JunosEventsEventsTypeEventListTimestampType) XXX_Size() int {
	return xxx_messageInfo_JunosEventsEventsTypeEventListTimestampType.Size(m)
}
func (m *JunosEventsEventsTypeEventListTimestampType) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosEventsEventsTypeEventListTimestampType.DiscardUnknown(m)
}

var xxx_messageInfo_JunosEventsEventsTypeEventListTimestampType proto.InternalMessageInfo

func (m *JunosEventsEventsTypeEventListTimestampType) GetSeconds() uint32 {
	if m != nil && m.Seconds != nil {
		return *m.Seconds
	}
	return 0
}

func (m *JunosEventsEventsTypeEventListTimestampType) GetMicroseconds() uint32 {
	if m != nil && m.Microseconds != nil {
		return *m.Microseconds
	}
	return 0
}

type JunosEventsEventsTypeEventListAttributesList struct {
	Key                  *string  `protobuf:"bytes,51,opt,name=key" json:"key,omitempty"`
	Value                *string  `protobuf:"bytes,52,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JunosEventsEventsTypeEventListAttributesList) Reset() {
	*m = JunosEventsEventsTypeEventListAttributesList{}
}
func (m *JunosEventsEventsTypeEventListAttributesList) String() string {
	return proto.CompactTextString(m)
}
func (*JunosEventsEventsTypeEventListAttributesList) ProtoMessage() {}
func (*JunosEventsEventsTypeEventListAttributesList) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5807bd979efd0cf, []int{0, 0, 0, 1}
}
func (m *JunosEventsEventsTypeEventListAttributesList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosEventsEventsTypeEventListAttributesList.Unmarshal(m, b)
}
func (m *JunosEventsEventsTypeEventListAttributesList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosEventsEventsTypeEventListAttributesList.Marshal(b, m, deterministic)
}
func (m *JunosEventsEventsTypeEventListAttributesList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosEventsEventsTypeEventListAttributesList.Merge(m, src)
}
func (m *JunosEventsEventsTypeEventListAttributesList) XXX_Size() int {
	return xxx_messageInfo_JunosEventsEventsTypeEventListAttributesList.Size(m)
}
func (m *JunosEventsEventsTypeEventListAttributesList) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosEventsEventsTypeEventListAttributesList.DiscardUnknown(m)
}

var xxx_messageInfo_JunosEventsEventsTypeEventListAttributesList proto.InternalMessageInfo

func (m *JunosEventsEventsTypeEventListAttributesList) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *JunosEventsEventsTypeEventListAttributesList) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

var E_JnprJunosEventsExt = &proto.ExtensionDesc{
	ExtendedType:  (*JuniperNetworksSensors)(nil),
	ExtensionType: (*JunosEvents)(nil),
	Field:         42,
	Name:          "jnpr_junos_events_ext",
	Tag:           "bytes,42,opt,name=jnpr_junos_events_ext",
	Filename:      "eventd.proto",
}

func init() {
	proto.RegisterType((*JunosEvents)(nil), "junos_events")
	proto.RegisterType((*JunosEventsEventsType)(nil), "junos_events.events_type")
	proto.RegisterType((*JunosEventsEventsTypeEventList)(nil), "junos_events.events_type.event_list")
	proto.RegisterType((*JunosEventsEventsTypeEventListTimestampType)(nil), "junos_events.events_type.event_list.timestamp_type")
	proto.RegisterType((*JunosEventsEventsTypeEventListAttributesList)(nil), "junos_events.events_type.event_list.attributes_list")
	proto.RegisterExtension(E_JnprJunosEventsExt)
}

func init() { proto.RegisterFile("eventd.proto", fileDescriptor_d5807bd979efd0cf) }

var fileDescriptor_d5807bd979efd0cf = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x55, 0xba, 0x4d, 0x69, 0x67, 0x77, 0x0b, 0x32, 0x1f, 0x35, 0x39, 0xa0, 0x55, 0xc5, 0x21,
	0xe2, 0x90, 0xc3, 0x76, 0xf9, 0x6a, 0x39, 0x54, 0x48, 0xbd, 0xf4, 0xd0, 0x83, 0x11, 0x07, 0x4e,
	0x51, 0xd8, 0x0c, 0xc5, 0x6d, 0x62, 0x5b, 0xf6, 0x6c, 0xe9, 0x5e, 0x39, 0xf2, 0x07, 0xe0, 0xcf,
	0xf0, 0x1b, 0xf8, 0x4b, 0xc8, 0x76, 0xd3, 0xcd, 0x22, 0x21, 0xf5, 0x94, 0x79, 0xef, 0xf9, 0xbd,
	0x78, 0xc6, 0x03, 0x23, 0xbc, 0x42, 0x45, 0x75, 0x61, 0xac, 0x26, 0x9d, 0x3d, 0x24, 0x6c, 0xb0,
	0x45, 0xb2, 0xcb, 0x92, 0xb4, 0x89, 0xe4, 0xfe, 0x9f, 0x14, 0x46, 0x17, 0x0b, 0xa5, 0x5d, 0x19,
	0xce, 0x3a, 0x36, 0x85, 0xad, 0x58, 0xf1, 0x9f, 0xc9, 0x24, 0xc9, 0x87, 0xd3, 0xa7, 0x45, 0x5f,
	0x2f, 0xe2, 0xa7, 0xa4, 0xa5, 0x41, 0x71, 0x73, 0x32, 0xfb, 0x91, 0xc2, 0xb0, 0xc7, 0xb3, 0x23,
	0x48, 0x03, 0xf4, 0x11, 0x83, 0x7c, 0x38, 0x7d, 0xfe, 0xdf, 0x88, 0x58, 0x97, 0x8d, 0x74, 0x24,
	0xa2, 0x27, 0xfb, 0xbd, 0x09, 0xb0, 0x62, 0xd9, 0x2e, 0x6c, 0xc8, 0x9a, 0x1f, 0x4c, 0x92, 0x7c,
	0x47, 0x6c, 0xc8, 0x9a, 0x31, 0xd8, 0xf4, 0x46, 0x3e, 0x0b, 0x4c, 0xa8, 0x99, 0x80, 0x1d, 0x92,
	0x2d, 0x3a, 0xaa, 0x5a, 0xd3, 0x5d, 0xfb, 0xe0, 0x2e, 0xff, 0x2c, 0x6e, 0x6d, 0xb1, 0xa1, 0x55,
	0x0c, 0xcb, 0x60, 0xdb, 0x58, 0xa9, 0xad, 0xa4, 0x25, 0x7f, 0x19, 0xfe, 0x75, 0x8b, 0xbd, 0xf6,
	0xa5, 0x9a, 0xcb, 0xc6, 0x6b, 0xaf, 0xa2, 0xd6, 0x61, 0xf6, 0x00, 0x06, 0x46, 0xd6, 0xfc, 0xf5,
	0x24, 0xc9, 0xc7, 0xc2, 0x97, 0x8c, 0xc3, 0xbd, 0x16, 0x9d, 0xab, 0xce, 0x91, 0xbf, 0x09, 0x87,
	0x3b, 0xc8, 0x9e, 0xc0, 0x56, 0x5d, 0x61, 0xab, 0x15, 0x7f, 0x1b, 0x84, 0x1b, 0xe4, 0xf3, 0xbf,
	0x6a, 0x47, 0xaa, 0x6a, 0x91, 0x1f, 0xc6, 0xfc, 0x0e, 0x7b, 0x4f, 0xe3, 0x82, 0x72, 0x14, 0x3d,
	0x11, 0xb1, 0x8f, 0x00, 0x15, 0x91, 0x95, 0x9f, 0x17, 0x84, 0x8e, 0xff, 0x8a, 0x83, 0x9f, 0xdd,
	0x69, 0x08, 0x2b, 0x5f, 0x7c, 0x88, 0x5e, 0x10, 0x7b, 0x06, 0xd0, 0xe8, 0x73, 0x6d, 0x48, 0x6a,
	0xe5, 0xf8, 0xbb, 0x49, 0x92, 0xa7, 0xa2, 0xc7, 0x64, 0x67, 0xb0, 0xbb, 0x3e, 0x43, 0xdf, 0xae,
	0xc3, 0xb9, 0x56, 0xb5, 0x0b, 0xaf, 0x36, 0x16, 0x1d, 0x64, 0xfb, 0x30, 0x6a, 0xe5, 0xdc, 0xea,
	0x4e, 0x9e, 0x05, 0x79, 0x8d, 0xcb, 0x8e, 0xe1, 0xfe, 0x3f, 0xd7, 0x61, 0x7b, 0x30, 0xb8, 0xc4,
	0x65, 0x5c, 0x81, 0xf7, 0xe9, 0xf7, 0xe3, 0x8d, 0xed, 0x44, 0x78, 0x86, 0x3d, 0x82, 0xf4, 0xaa,
	0x6a, 0x16, 0xdd, 0x2e, 0x44, 0x70, 0xf8, 0x09, 0x1e, 0x5f, 0x28, 0x63, 0xcb, 0x7e, 0xe7, 0x25,
	0x5e, 0x13, 0xdb, 0x2b, 0x4e, 0x17, 0x4a, 0x1a, 0xb4, 0x67, 0x48, 0xdf, 0xb4, 0xbd, 0x74, 0x1f,
	0x50, 0x39, 0x6d, 0x1d, 0x7f, 0x11, 0x16, 0x66, 0xbc, 0x36, 0x2b, 0xc1, 0x7c, 0xc8, 0xa9, 0x67,
	0x4e, 0x02, 0x71, 0x72, 0x4d, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x54, 0x74, 0x79, 0x8e, 0x50,
	0x03, 0x00, 0x00,
}
