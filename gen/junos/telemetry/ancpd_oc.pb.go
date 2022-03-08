// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ancpd_oc.proto

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

type JunosAncp struct {
	System               *JunosAncpSystemType `protobuf:"bytes,151,opt,name=system" json:"system,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *JunosAncp) Reset()         { *m = JunosAncp{} }
func (m *JunosAncp) String() string { return proto.CompactTextString(m) }
func (*JunosAncp) ProtoMessage()    {}
func (*JunosAncp) Descriptor() ([]byte, []int) {
	return fileDescriptor_802aa389b0dda049, []int{0}
}
func (m *JunosAncp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosAncp.Unmarshal(m, b)
}
func (m *JunosAncp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosAncp.Marshal(b, m, deterministic)
}
func (m *JunosAncp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosAncp.Merge(m, src)
}
func (m *JunosAncp) XXX_Size() int {
	return xxx_messageInfo_JunosAncp.Size(m)
}
func (m *JunosAncp) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosAncp.DiscardUnknown(m)
}

var xxx_messageInfo_JunosAncp proto.InternalMessageInfo

func (m *JunosAncp) GetSystem() *JunosAncpSystemType {
	if m != nil {
		return m.System
	}
	return nil
}

type JunosAncpSystemType struct {
	SubscriberManagement *JunosAncpSystemTypeSubscriberManagementType `protobuf:"bytes,151,opt,name=subscriber_management,json=subscriberManagement" json:"subscriber_management,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *JunosAncpSystemType) Reset()         { *m = JunosAncpSystemType{} }
func (m *JunosAncpSystemType) String() string { return proto.CompactTextString(m) }
func (*JunosAncpSystemType) ProtoMessage()    {}
func (*JunosAncpSystemType) Descriptor() ([]byte, []int) {
	return fileDescriptor_802aa389b0dda049, []int{0, 0}
}
func (m *JunosAncpSystemType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosAncpSystemType.Unmarshal(m, b)
}
func (m *JunosAncpSystemType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosAncpSystemType.Marshal(b, m, deterministic)
}
func (m *JunosAncpSystemType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosAncpSystemType.Merge(m, src)
}
func (m *JunosAncpSystemType) XXX_Size() int {
	return xxx_messageInfo_JunosAncpSystemType.Size(m)
}
func (m *JunosAncpSystemType) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosAncpSystemType.DiscardUnknown(m)
}

var xxx_messageInfo_JunosAncpSystemType proto.InternalMessageInfo

func (m *JunosAncpSystemType) GetSubscriberManagement() *JunosAncpSystemTypeSubscriberManagementType {
	if m != nil {
		return m.SubscriberManagement
	}
	return nil
}

type JunosAncpSystemTypeSubscriberManagementType struct {
	AccessNetwork        *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkType `protobuf:"bytes,151,opt,name=access_network,json=accessNetwork" json:"access_network,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                      `json:"-"`
	XXX_unrecognized     []byte                                                        `json:"-"`
	XXX_sizecache        int32                                                         `json:"-"`
}

func (m *JunosAncpSystemTypeSubscriberManagementType) Reset() {
	*m = JunosAncpSystemTypeSubscriberManagementType{}
}
func (m *JunosAncpSystemTypeSubscriberManagementType) String() string {
	return proto.CompactTextString(m)
}
func (*JunosAncpSystemTypeSubscriberManagementType) ProtoMessage() {}
func (*JunosAncpSystemTypeSubscriberManagementType) Descriptor() ([]byte, []int) {
	return fileDescriptor_802aa389b0dda049, []int{0, 0, 0}
}
func (m *JunosAncpSystemTypeSubscriberManagementType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosAncpSystemTypeSubscriberManagementType.Unmarshal(m, b)
}
func (m *JunosAncpSystemTypeSubscriberManagementType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosAncpSystemTypeSubscriberManagementType.Marshal(b, m, deterministic)
}
func (m *JunosAncpSystemTypeSubscriberManagementType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosAncpSystemTypeSubscriberManagementType.Merge(m, src)
}
func (m *JunosAncpSystemTypeSubscriberManagementType) XXX_Size() int {
	return xxx_messageInfo_JunosAncpSystemTypeSubscriberManagementType.Size(m)
}
func (m *JunosAncpSystemTypeSubscriberManagementType) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosAncpSystemTypeSubscriberManagementType.DiscardUnknown(m)
}

var xxx_messageInfo_JunosAncpSystemTypeSubscriberManagementType proto.InternalMessageInfo

func (m *JunosAncpSystemTypeSubscriberManagementType) GetAccessNetwork() *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkType {
	if m != nil {
		return m.AccessNetwork
	}
	return nil
}

type JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkType struct {
	Ancp                 *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpType `protobuf:"bytes,151,opt,name=ancp" json:"ancp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                              `json:"-"`
	XXX_unrecognized     []byte                                                                `json:"-"`
	XXX_sizecache        int32                                                                 `json:"-"`
}

func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkType) Reset() {
	*m = JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkType{}
}
func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkType) String() string {
	return proto.CompactTextString(m)
}
func (*JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkType) ProtoMessage() {}
func (*JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkType) Descriptor() ([]byte, []int) {
	return fileDescriptor_802aa389b0dda049, []int{0, 0, 0, 0}
}
func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkType.Unmarshal(m, b)
}
func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkType.Marshal(b, m, deterministic)
}
func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkType.Merge(m, src)
}
func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkType) XXX_Size() int {
	return xxx_messageInfo_JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkType.Size(m)
}
func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkType) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkType.DiscardUnknown(m)
}

var xxx_messageInfo_JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkType proto.InternalMessageInfo

func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkType) GetAncp() *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpType {
	if m != nil {
		return m.Ancp
	}
	return nil
}

type JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpType struct {
	Protocol             *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpTypeProtocolType `protobuf:"bytes,151,opt,name=protocol" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                          `json:"-"`
	XXX_unrecognized     []byte                                                                            `json:"-"`
	XXX_sizecache        int32                                                                             `json:"-"`
}

func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpType) Reset() {
	*m = JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpType{}
}
func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpType) String() string {
	return proto.CompactTextString(m)
}
func (*JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpType) ProtoMessage() {}
func (*JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpType) Descriptor() ([]byte, []int) {
	return fileDescriptor_802aa389b0dda049, []int{0, 0, 0, 0, 0}
}
func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpType.Unmarshal(m, b)
}
func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpType.Marshal(b, m, deterministic)
}
func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpType.Merge(m, src)
}
func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpType) XXX_Size() int {
	return xxx_messageInfo_JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpType.Size(m)
}
func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpType) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpType.DiscardUnknown(m)
}

var xxx_messageInfo_JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpType proto.InternalMessageInfo

func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpType) GetProtocol() *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpTypeProtocolType {
	if m != nil {
		return m.Protocol
	}
	return nil
}

type JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpTypeProtocolType struct {
	ConfiguredNeighborCount     *uint32  `protobuf:"varint,51,opt,name=configured_neighbor_count,json=configuredNeighborCount" json:"configured_neighbor_count,omitempty"`
	EstablishingNeighborCount   *uint32  `protobuf:"varint,52,opt,name=establishing_neighbor_count,json=establishingNeighborCount" json:"establishing_neighbor_count,omitempty"`
	EstablishedNeighborCount    *uint32  `protobuf:"varint,53,opt,name=established_neighbor_count,json=establishedNeighborCount" json:"established_neighbor_count,omitempty"`
	NotEstablishedNeighborCount *uint32  `protobuf:"varint,54,opt,name=not_established_neighbor_count,json=notEstablishedNeighborCount" json:"not_established_neighbor_count,omitempty"`
	TotalNeighborCount          *uint32  `protobuf:"varint,55,opt,name=total_neighbor_count,json=totalNeighborCount" json:"total_neighbor_count,omitempty"`
	MappedStaticSubscriberCount *uint32  `protobuf:"varint,56,opt,name=mapped_static_subscriber_count,json=mappedStaticSubscriberCount" json:"mapped_static_subscriber_count,omitempty"`
	PortUpCount                 *uint64  `protobuf:"varint,57,opt,name=port_up_count,json=portUpCount" json:"port_up_count,omitempty"`
	PortDownCount               *uint64  `protobuf:"varint,58,opt,name=port_down_count,json=portDownCount" json:"port_down_count,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpTypeProtocolType) Reset() {
	*m = JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpTypeProtocolType{}
}
func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpTypeProtocolType) String() string {
	return proto.CompactTextString(m)
}
func (*JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpTypeProtocolType) ProtoMessage() {
}
func (*JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpTypeProtocolType) Descriptor() ([]byte, []int) {
	return fileDescriptor_802aa389b0dda049, []int{0, 0, 0, 0, 0, 0}
}
func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpTypeProtocolType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpTypeProtocolType.Unmarshal(m, b)
}
func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpTypeProtocolType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpTypeProtocolType.Marshal(b, m, deterministic)
}
func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpTypeProtocolType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpTypeProtocolType.Merge(m, src)
}
func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpTypeProtocolType) XXX_Size() int {
	return xxx_messageInfo_JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpTypeProtocolType.Size(m)
}
func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpTypeProtocolType) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpTypeProtocolType.DiscardUnknown(m)
}

var xxx_messageInfo_JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpTypeProtocolType proto.InternalMessageInfo

func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpTypeProtocolType) GetConfiguredNeighborCount() uint32 {
	if m != nil && m.ConfiguredNeighborCount != nil {
		return *m.ConfiguredNeighborCount
	}
	return 0
}

func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpTypeProtocolType) GetEstablishingNeighborCount() uint32 {
	if m != nil && m.EstablishingNeighborCount != nil {
		return *m.EstablishingNeighborCount
	}
	return 0
}

func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpTypeProtocolType) GetEstablishedNeighborCount() uint32 {
	if m != nil && m.EstablishedNeighborCount != nil {
		return *m.EstablishedNeighborCount
	}
	return 0
}

func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpTypeProtocolType) GetNotEstablishedNeighborCount() uint32 {
	if m != nil && m.NotEstablishedNeighborCount != nil {
		return *m.NotEstablishedNeighborCount
	}
	return 0
}

func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpTypeProtocolType) GetTotalNeighborCount() uint32 {
	if m != nil && m.TotalNeighborCount != nil {
		return *m.TotalNeighborCount
	}
	return 0
}

func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpTypeProtocolType) GetMappedStaticSubscriberCount() uint32 {
	if m != nil && m.MappedStaticSubscriberCount != nil {
		return *m.MappedStaticSubscriberCount
	}
	return 0
}

func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpTypeProtocolType) GetPortUpCount() uint64 {
	if m != nil && m.PortUpCount != nil {
		return *m.PortUpCount
	}
	return 0
}

func (m *JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpTypeProtocolType) GetPortDownCount() uint64 {
	if m != nil && m.PortDownCount != nil {
		return *m.PortDownCount
	}
	return 0
}

var E_JnprJunosAncpExt = &proto.ExtensionDesc{
	ExtendedType:  (*JuniperNetworksSensors)(nil),
	ExtensionType: (*JunosAncp)(nil),
	Field:         32,
	Name:          "jnpr_junos_ancp_ext",
	Tag:           "bytes,32,opt,name=jnpr_junos_ancp_ext",
	Filename:      "ancpd_oc.proto",
}

func init() {
	proto.RegisterType((*JunosAncp)(nil), "junos_ancp")
	proto.RegisterType((*JunosAncpSystemType)(nil), "junos_ancp.system_type")
	proto.RegisterType((*JunosAncpSystemTypeSubscriberManagementType)(nil), "junos_ancp.system_type.subscriber_management_type")
	proto.RegisterType((*JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkType)(nil), "junos_ancp.system_type.subscriber_management_type.access_network_type")
	proto.RegisterType((*JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpType)(nil), "junos_ancp.system_type.subscriber_management_type.access_network_type.ancp_type")
	proto.RegisterType((*JunosAncpSystemTypeSubscriberManagementTypeAccessNetworkTypeAncpTypeProtocolType)(nil), "junos_ancp.system_type.subscriber_management_type.access_network_type.ancp_type.protocol_type")
	proto.RegisterExtension(E_JnprJunosAncpExt)
}

func init() { proto.RegisterFile("ancpd_oc.proto", fileDescriptor_802aa389b0dda049) }

var fileDescriptor_802aa389b0dda049 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xdf, 0x6a, 0x13, 0x41,
	0x14, 0xc6, 0x09, 0xc6, 0x5a, 0x4f, 0x48, 0x95, 0x4d, 0x25, 0xe9, 0x16, 0x24, 0xf4, 0x42, 0x72,
	0xb5, 0x94, 0xfa, 0x3f, 0x88, 0x20, 0xb5, 0x5e, 0x14, 0x2c, 0x92, 0xe0, 0xad, 0xc3, 0x66, 0x72,
	0x4c, 0xb6, 0x26, 0x67, 0x86, 0x99, 0xb3, 0xa4, 0xc1, 0x07, 0xf1, 0x09, 0x7c, 0x20, 0xaf, 0xf5,
	0xd2, 0x2b, 0x9f, 0x42, 0x76, 0x66, 0x77, 0xb3, 0x69, 0x1b, 0x41, 0xe8, 0xe5, 0x9e, 0xef, 0xf7,
	0x7d, 0xdf, 0x81, 0x39, 0x0b, 0x3b, 0x31, 0x49, 0x3d, 0x16, 0x4a, 0x46, 0xda, 0x28, 0x56, 0x61,
	0x8b, 0x71, 0x86, 0x73, 0x64, 0xb3, 0x14, 0xac, 0xb4, 0x1f, 0x1e, 0x7c, 0xdf, 0x06, 0x38, 0x4f,
	0x49, 0x59, 0x91, 0xd1, 0xc1, 0x21, 0x6c, 0xd9, 0xa5, 0x65, 0x9c, 0x77, 0xbe, 0xd5, 0xba, 0xb5,
	0x5e, 0xe3, 0xa8, 0x1d, 0xad, 0xd4, 0xc8, 0x4b, 0x82, 0x97, 0x1a, 0x07, 0x39, 0x17, 0xfe, 0xbe,
	0x03, 0x8d, 0xca, 0x3c, 0x98, 0xc2, 0x03, 0x9b, 0x8e, 0xac, 0x34, 0xc9, 0x08, 0x8d, 0x98, 0xc7,
	0x14, 0x4f, 0x70, 0x8e, 0xc4, 0x45, 0xe0, 0xd1, 0x86, 0xc0, 0xe8, 0x5a, 0x97, 0xef, 0xda, 0x5d,
	0x69, 0xef, 0x4b, 0x29, 0xfc, 0xb3, 0x05, 0xe1, 0x66, 0x53, 0x40, 0xb0, 0x13, 0x4b, 0x89, 0xd6,
	0x0a, 0x42, 0x5e, 0x28, 0xf3, 0xa5, 0xd8, 0xe0, 0xdd, 0xff, 0x6f, 0x10, 0xad, 0x27, 0xf9, 0xad,
	0x9a, 0x7e, 0x78, 0xe6, 0x67, 0xe1, 0xcf, 0xdb, 0xd0, 0xba, 0x06, 0x0b, 0x10, 0xea, 0x59, 0x53,
	0xd1, 0xfe, 0xe1, 0x66, 0xda, 0xa3, 0xcc, 0xef, 0xf7, 0x70, 0xf1, 0xe1, 0xaf, 0x3a, 0xdc, 0x2d,
	0x67, 0xc1, 0x57, 0xd8, 0x76, 0xef, 0x2b, 0xd5, 0xac, 0x28, 0xfe, 0x74, 0xd3, 0xc5, 0x51, 0xd1,
	0xe0, 0xd7, 0x28, 0x0b, 0xc3, 0x1f, 0xb7, 0xa0, 0xb9, 0xa6, 0x05, 0x7d, 0xd8, 0x93, 0x8a, 0x3e,
	0x27, 0x93, 0xd4, 0xe0, 0x58, 0x10, 0x26, 0x93, 0xe9, 0x48, 0x19, 0x21, 0x55, 0x4a, 0xdc, 0x79,
	0xdc, 0xad, 0xf5, 0x9a, 0x83, 0xf6, 0x0a, 0x38, 0xcb, 0xf5, 0xe3, 0x4c, 0x0e, 0x5e, 0xc3, 0x3e,
	0x5a, 0x8e, 0x47, 0xb3, 0xc4, 0x4e, 0x13, 0x9a, 0x5c, 0x76, 0x3f, 0x71, 0xee, 0xbd, 0x2a, 0xb2,
	0xee, 0x7f, 0x05, 0x61, 0x29, 0x5e, 0x2d, 0x7f, 0xea, 0xec, 0x9d, 0x0a, 0xb1, 0xee, 0x3e, 0x86,
	0x87, 0xa4, 0x58, 0xfc, 0x23, 0xe1, 0x99, 0x4b, 0xd8, 0x27, 0xc5, 0x27, 0x9b, 0x42, 0x0e, 0x61,
	0x97, 0x15, 0xc7, 0xb3, 0xcb, 0xd6, 0xe7, 0xce, 0x1a, 0x38, 0xed, 0x4a, 0xed, 0x3c, 0xd6, 0x1a,
	0xc7, 0xc2, 0x72, 0xcc, 0x89, 0x14, 0x95, 0x47, 0xf2, 0xde, 0x17, 0xbe, 0xd6, 0x53, 0x43, 0x07,
	0x0d, 0x4b, 0xc6, 0x87, 0x1c, 0x40, 0x53, 0x2b, 0xc3, 0x22, 0xd5, 0xb9, 0xe7, 0x65, 0xb7, 0xd6,
	0xab, 0x0f, 0x1a, 0xd9, 0xf0, 0xa3, 0xf6, 0xcc, 0x23, 0xb8, 0xe7, 0x98, 0xb1, 0x5a, 0x50, 0x4e,
	0xf5, 0x1d, 0xe5, 0xac, 0x6f, 0xd5, 0x82, 0x1c, 0xd7, 0x1f, 0x42, 0xeb, 0x9c, 0xb4, 0x11, 0xab,
	0x1b, 0x12, 0x78, 0xc1, 0x41, 0x3b, 0x3a, 0x4d, 0x29, 0xd1, 0x68, 0xf2, 0xbf, 0xc0, 0x0e, 0x91,
	0xac, 0x32, 0xb6, 0xd3, 0x75, 0x37, 0xd7, 0xa8, 0xdc, 0xdc, 0xe0, 0x7e, 0x16, 0x70, 0x9a, 0x7d,
	0xbf, 0x21, 0xa9, 0x4f, 0x2e, 0xf8, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x66, 0x42, 0x20, 0x50,
	0xa2, 0x04, 0x00, 0x00,
}
