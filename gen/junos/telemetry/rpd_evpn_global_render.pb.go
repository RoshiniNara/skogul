// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpd_evpn_global_render.proto

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

type JunosEvpn struct {
	Evpn                 *JunosEvpnEvpnType `protobuf:"bytes,171,opt,name=evpn" json:"evpn,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *JunosEvpn) Reset()         { *m = JunosEvpn{} }
func (m *JunosEvpn) String() string { return proto.CompactTextString(m) }
func (*JunosEvpn) ProtoMessage()    {}
func (*JunosEvpn) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce797971eeb39af0, []int{0}
}
func (m *JunosEvpn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosEvpn.Unmarshal(m, b)
}
func (m *JunosEvpn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosEvpn.Marshal(b, m, deterministic)
}
func (m *JunosEvpn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosEvpn.Merge(m, src)
}
func (m *JunosEvpn) XXX_Size() int {
	return xxx_messageInfo_JunosEvpn.Size(m)
}
func (m *JunosEvpn) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosEvpn.DiscardUnknown(m)
}

var xxx_messageInfo_JunosEvpn proto.InternalMessageInfo

func (m *JunosEvpn) GetEvpn() *JunosEvpnEvpnType {
	if m != nil {
		return m.Evpn
	}
	return nil
}

type JunosEvpnEvpnType struct {
	EvpnSmetForwarding   *JunosEvpnEvpnTypeEvpnSmetForwardingType `protobuf:"bytes,171,opt,name=evpn_smet_forwarding,json=evpnSmetForwarding" json:"evpn_smet_forwarding,omitempty"`
	L3Context            []*JunosEvpnEvpnTypeL3ContextList        `protobuf:"bytes,180,rep,name=l3_context,json=l3Context" json:"l3_context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *JunosEvpnEvpnType) Reset()         { *m = JunosEvpnEvpnType{} }
func (m *JunosEvpnEvpnType) String() string { return proto.CompactTextString(m) }
func (*JunosEvpnEvpnType) ProtoMessage()    {}
func (*JunosEvpnEvpnType) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce797971eeb39af0, []int{0, 0}
}
func (m *JunosEvpnEvpnType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosEvpnEvpnType.Unmarshal(m, b)
}
func (m *JunosEvpnEvpnType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosEvpnEvpnType.Marshal(b, m, deterministic)
}
func (m *JunosEvpnEvpnType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosEvpnEvpnType.Merge(m, src)
}
func (m *JunosEvpnEvpnType) XXX_Size() int {
	return xxx_messageInfo_JunosEvpnEvpnType.Size(m)
}
func (m *JunosEvpnEvpnType) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosEvpnEvpnType.DiscardUnknown(m)
}

var xxx_messageInfo_JunosEvpnEvpnType proto.InternalMessageInfo

func (m *JunosEvpnEvpnType) GetEvpnSmetForwarding() *JunosEvpnEvpnTypeEvpnSmetForwardingType {
	if m != nil {
		return m.EvpnSmetForwarding
	}
	return nil
}

func (m *JunosEvpnEvpnType) GetL3Context() []*JunosEvpnEvpnTypeL3ContextList {
	if m != nil {
		return m.L3Context
	}
	return nil
}

type JunosEvpnEvpnTypeEvpnSmetForwardingType struct {
	Enabled              *bool    `protobuf:"varint,71,opt,name=enabled" json:"enabled,omitempty"`
	NexthopLimit         *uint32  `protobuf:"varint,72,opt,name=nexthop_limit,json=nexthopLimit" json:"nexthop_limit,omitempty"`
	NexthopUsage         *uint32  `protobuf:"varint,73,opt,name=nexthop_usage,json=nexthopUsage" json:"nexthop_usage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JunosEvpnEvpnTypeEvpnSmetForwardingType) Reset() {
	*m = JunosEvpnEvpnTypeEvpnSmetForwardingType{}
}
func (m *JunosEvpnEvpnTypeEvpnSmetForwardingType) String() string { return proto.CompactTextString(m) }
func (*JunosEvpnEvpnTypeEvpnSmetForwardingType) ProtoMessage()    {}
func (*JunosEvpnEvpnTypeEvpnSmetForwardingType) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce797971eeb39af0, []int{0, 0, 0}
}
func (m *JunosEvpnEvpnTypeEvpnSmetForwardingType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosEvpnEvpnTypeEvpnSmetForwardingType.Unmarshal(m, b)
}
func (m *JunosEvpnEvpnTypeEvpnSmetForwardingType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosEvpnEvpnTypeEvpnSmetForwardingType.Marshal(b, m, deterministic)
}
func (m *JunosEvpnEvpnTypeEvpnSmetForwardingType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosEvpnEvpnTypeEvpnSmetForwardingType.Merge(m, src)
}
func (m *JunosEvpnEvpnTypeEvpnSmetForwardingType) XXX_Size() int {
	return xxx_messageInfo_JunosEvpnEvpnTypeEvpnSmetForwardingType.Size(m)
}
func (m *JunosEvpnEvpnTypeEvpnSmetForwardingType) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosEvpnEvpnTypeEvpnSmetForwardingType.DiscardUnknown(m)
}

var xxx_messageInfo_JunosEvpnEvpnTypeEvpnSmetForwardingType proto.InternalMessageInfo

func (m *JunosEvpnEvpnTypeEvpnSmetForwardingType) GetEnabled() bool {
	if m != nil && m.Enabled != nil {
		return *m.Enabled
	}
	return false
}

func (m *JunosEvpnEvpnTypeEvpnSmetForwardingType) GetNexthopLimit() uint32 {
	if m != nil && m.NexthopLimit != nil {
		return *m.NexthopLimit
	}
	return 0
}

func (m *JunosEvpnEvpnTypeEvpnSmetForwardingType) GetNexthopUsage() uint32 {
	if m != nil && m.NexthopUsage != nil {
		return *m.NexthopUsage
	}
	return 0
}

type JunosEvpnEvpnTypeL3ContextList struct {
	ContextName          *string                                               `protobuf:"bytes,71,opt,name=context_name,json=contextName" json:"context_name,omitempty"`
	Encapsulation        *string                                               `protobuf:"bytes,72,opt,name=encapsulation" json:"encapsulation,omitempty"`
	AdvertisementMode    *string                                               `protobuf:"bytes,73,opt,name=advertisement_mode,json=advertisementMode" json:"advertisement_mode,omitempty"`
	MulticastRoutingMode *string                                               `protobuf:"bytes,74,opt,name=multicast_routing_mode,json=multicastRoutingMode" json:"multicast_routing_mode,omitempty"`
	IpPrefixDatabase     []*JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList `protobuf:"bytes,170,rep,name=ip_prefix_database,json=ipPrefixDatabase" json:"ip_prefix_database,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                              `json:"-"`
	XXX_unrecognized     []byte                                                `json:"-"`
	XXX_sizecache        int32                                                 `json:"-"`
}

func (m *JunosEvpnEvpnTypeL3ContextList) Reset()         { *m = JunosEvpnEvpnTypeL3ContextList{} }
func (m *JunosEvpnEvpnTypeL3ContextList) String() string { return proto.CompactTextString(m) }
func (*JunosEvpnEvpnTypeL3ContextList) ProtoMessage()    {}
func (*JunosEvpnEvpnTypeL3ContextList) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce797971eeb39af0, []int{0, 0, 1}
}
func (m *JunosEvpnEvpnTypeL3ContextList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosEvpnEvpnTypeL3ContextList.Unmarshal(m, b)
}
func (m *JunosEvpnEvpnTypeL3ContextList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosEvpnEvpnTypeL3ContextList.Marshal(b, m, deterministic)
}
func (m *JunosEvpnEvpnTypeL3ContextList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosEvpnEvpnTypeL3ContextList.Merge(m, src)
}
func (m *JunosEvpnEvpnTypeL3ContextList) XXX_Size() int {
	return xxx_messageInfo_JunosEvpnEvpnTypeL3ContextList.Size(m)
}
func (m *JunosEvpnEvpnTypeL3ContextList) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosEvpnEvpnTypeL3ContextList.DiscardUnknown(m)
}

var xxx_messageInfo_JunosEvpnEvpnTypeL3ContextList proto.InternalMessageInfo

func (m *JunosEvpnEvpnTypeL3ContextList) GetContextName() string {
	if m != nil && m.ContextName != nil {
		return *m.ContextName
	}
	return ""
}

func (m *JunosEvpnEvpnTypeL3ContextList) GetEncapsulation() string {
	if m != nil && m.Encapsulation != nil {
		return *m.Encapsulation
	}
	return ""
}

func (m *JunosEvpnEvpnTypeL3ContextList) GetAdvertisementMode() string {
	if m != nil && m.AdvertisementMode != nil {
		return *m.AdvertisementMode
	}
	return ""
}

func (m *JunosEvpnEvpnTypeL3ContextList) GetMulticastRoutingMode() string {
	if m != nil && m.MulticastRoutingMode != nil {
		return *m.MulticastRoutingMode
	}
	return ""
}

func (m *JunosEvpnEvpnTypeL3ContextList) GetIpPrefixDatabase() []*JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList {
	if m != nil {
		return m.IpPrefixDatabase
	}
	return nil
}

type JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList struct {
	IpPrefix             *string  `protobuf:"bytes,71,opt,name=ip_prefix,json=ipPrefix" json:"ip_prefix,omitempty"`
	RouteStatus          *string  `protobuf:"bytes,72,opt,name=route_status,json=routeStatus" json:"route_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList) Reset() {
	*m = JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList{}
}
func (m *JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList) String() string {
	return proto.CompactTextString(m)
}
func (*JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList) ProtoMessage() {}
func (*JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce797971eeb39af0, []int{0, 0, 1, 0}
}
func (m *JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList.Unmarshal(m, b)
}
func (m *JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList.Marshal(b, m, deterministic)
}
func (m *JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList.Merge(m, src)
}
func (m *JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList) XXX_Size() int {
	return xxx_messageInfo_JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList.Size(m)
}
func (m *JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList.DiscardUnknown(m)
}

var xxx_messageInfo_JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList proto.InternalMessageInfo

func (m *JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList) GetIpPrefix() string {
	if m != nil && m.IpPrefix != nil {
		return *m.IpPrefix
	}
	return ""
}

func (m *JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList) GetRouteStatus() string {
	if m != nil && m.RouteStatus != nil {
		return *m.RouteStatus
	}
	return ""
}

var E_JnprJunosEvpnExt = &proto.ExtensionDesc{
	ExtendedType:  (*JuniperNetworksSensors)(nil),
	ExtensionType: (*JunosEvpn)(nil),
	Field:         113,
	Name:          "jnpr_junos_evpn_ext",
	Tag:           "bytes,113,opt,name=jnpr_junos_evpn_ext",
	Filename:      "rpd_evpn_global_render.proto",
}

func init() {
	proto.RegisterType((*JunosEvpn)(nil), "junos_evpn")
	proto.RegisterType((*JunosEvpnEvpnType)(nil), "junos_evpn.evpn_type")
	proto.RegisterType((*JunosEvpnEvpnTypeEvpnSmetForwardingType)(nil), "junos_evpn.evpn_type.evpn_smet_forwarding_type")
	proto.RegisterType((*JunosEvpnEvpnTypeL3ContextList)(nil), "junos_evpn.evpn_type.l3_context_list")
	proto.RegisterType((*JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList)(nil), "junos_evpn.evpn_type.l3_context_list.ip_prefix_database_list")
	proto.RegisterExtension(E_JnprJunosEvpnExt)
}

func init() { proto.RegisterFile("rpd_evpn_global_render.proto", fileDescriptor_ce797971eeb39af0) }

var fileDescriptor_ce797971eeb39af0 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x65, 0x4a, 0xa1, 0x9e, 0x34, 0xa2, 0x6c, 0x0b, 0x35, 0x11, 0x87, 0x50, 0x40, 0x8a,
	0x90, 0x30, 0x52, 0xcb, 0x09, 0x09, 0x09, 0x01, 0xe1, 0x23, 0x82, 0x0a, 0x39, 0xe2, 0xbc, 0x6c,
	0xe2, 0x69, 0xd8, 0xb2, 0xde, 0x5d, 0x76, 0xc7, 0x69, 0x7a, 0x42, 0xe2, 0xca, 0x9d, 0x87, 0x80,
	0x2b, 0x0f, 0xc2, 0x23, 0x21, 0xaf, 0xf3, 0xd1, 0x40, 0x2a, 0x71, 0xb2, 0xe6, 0xf7, 0x9f, 0xf9,
	0xcf, 0xec, 0x8c, 0xe1, 0xa6, 0xb3, 0x39, 0xc7, 0xb1, 0xd5, 0x7c, 0xa4, 0xcc, 0x40, 0x28, 0xee,
	0x50, 0xe7, 0xe8, 0x52, 0xeb, 0x0c, 0x99, 0xd6, 0x36, 0xa1, 0xc2, 0x02, 0xc9, 0x9d, 0x72, 0x32,
	0xb6, 0x86, 0x7b, 0xdf, 0x2e, 0x01, 0x1c, 0x97, 0xda, 0xf8, 0x50, 0xc7, 0xee, 0xc1, 0xc5, 0xea,
	0x9b, 0xfc, 0x8c, 0xda, 0x51, 0xa7, 0xb1, 0x7f, 0x2d, 0x5d, 0x68, 0x69, 0x30, 0xa6, 0x53, 0x8b,
	0x59, 0xc8, 0x69, 0xfd, 0x5e, 0x87, 0x78, 0xce, 0xd8, 0x00, 0x76, 0x42, 0xe0, 0x0b, 0x24, 0x7e,
	0x64, 0xdc, 0x89, 0x70, 0xb9, 0xd4, 0xa3, 0x99, 0xd3, 0x83, 0x95, 0x4e, 0xe9, 0xaa, 0x92, 0xba,
	0x07, 0xab, 0xa4, 0x7e, 0x81, 0xf4, 0x62, 0x2e, 0xb0, 0x2e, 0x80, 0x3a, 0xe0, 0x43, 0xa3, 0x09,
	0x27, 0x94, 0xfc, 0x8a, 0xda, 0x6b, 0x9d, 0xc6, 0xfe, 0xdd, 0xd5, 0xce, 0x8b, 0x44, 0xae, 0xa4,
	0xa7, 0x2c, 0x56, 0x07, 0xcf, 0xea, 0xb8, 0xf5, 0x05, 0x6e, 0x9c, 0xdb, 0x97, 0x25, 0x70, 0x19,
	0xb5, 0x18, 0x28, 0xcc, 0x93, 0x97, 0xed, 0xa8, 0xb3, 0x91, 0xcd, 0x42, 0x76, 0x1b, 0x9a, 0x1a,
	0x27, 0xf4, 0xd1, 0x58, 0xae, 0x64, 0x21, 0x29, 0x79, 0xd5, 0x8e, 0x3a, 0xcd, 0x6c, 0x73, 0x0a,
	0xdf, 0x54, 0xec, 0x6c, 0x52, 0xe9, 0xc5, 0x08, 0x93, 0xd7, 0x4b, 0x49, 0xef, 0x2b, 0xd6, 0xfa,
	0xbe, 0x06, 0x57, 0xfe, 0x9a, 0x8f, 0x75, 0x60, 0x73, 0x16, 0x6b, 0x51, 0x60, 0x68, 0x1e, 0x3f,
	0x5d, 0xff, 0xfa, 0xe4, 0xc2, 0x46, 0x94, 0x35, 0xa6, 0xd2, 0xa1, 0x28, 0x90, 0xdd, 0x81, 0x26,
	0xea, 0xa1, 0xb0, 0xbe, 0x54, 0x82, 0xa4, 0xd1, 0x61, 0x8e, 0x38, 0x5b, 0x86, 0xec, 0x3e, 0x30,
	0x91, 0x8f, 0xd1, 0x91, 0xf4, 0x58, 0xa0, 0x26, 0x5e, 0x98, 0xbc, 0x9e, 0x26, 0xce, 0xae, 0x2e,
	0x29, 0x6f, 0x4d, 0x8e, 0xec, 0x21, 0x5c, 0x2f, 0x4a, 0x45, 0x72, 0x28, 0x3c, 0x71, 0x67, 0x4a,
	0xaa, 0x16, 0x12, 0x4a, 0x7a, 0xa1, 0x64, 0x67, 0xae, 0x66, 0xb5, 0x18, 0xaa, 0x14, 0x30, 0x69,
	0xb9, 0x75, 0x78, 0x24, 0x27, 0x3c, 0x17, 0x24, 0x06, 0xc2, 0x63, 0xf2, 0xa3, 0x3e, 0xcc, 0xe3,
	0xff, 0x3a, 0x4c, 0xfa, 0xaf, 0x41, 0x7d, 0xb0, 0x2d, 0x69, 0xdf, 0x05, 0xfe, 0x7c, 0x8a, 0x5b,
	0x1f, 0x60, 0xf7, 0x9c, 0x64, 0xb6, 0x07, 0xf1, 0x5c, 0x5a, 0x5e, 0xdd, 0xc6, 0xcc, 0x86, 0xdd,
	0x82, 0xcd, 0xea, 0x61, 0xc8, 0x3d, 0x09, 0x2a, 0xfd, 0x74, 0x6d, 0x8d, 0xc0, 0xfa, 0x01, 0x3d,
	0xea, 0xc3, 0xf6, 0xb1, 0xb6, 0x8e, 0x2f, 0x06, 0xe7, 0x38, 0x21, 0xb6, 0x9b, 0xf6, 0x4a, 0x2d,
	0x2d, 0xba, 0x43, 0xa4, 0x13, 0xe3, 0x3e, 0xf9, 0x3e, 0x6a, 0x6f, 0x9c, 0x4f, 0x3e, 0x87, 0x7f,
	0xbb, 0x71, 0xe6, 0xa1, 0xd9, 0x56, 0x65, 0xd0, 0xab, 0xe2, 0xee, 0xd8, 0xea, 0xee, 0x84, 0xfe,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x28, 0xd1, 0x68, 0x2d, 0x96, 0x03, 0x00, 0x00,
}
