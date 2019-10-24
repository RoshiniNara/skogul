// Code generated by protoc-gen-go. DO NOT EDIT.
// source: jpppd_oc.proto

package gen

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type JunosPpp struct {
	System               *JunosPppSystemType `protobuf:"bytes,151,opt,name=system" json:"system,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *JunosPpp) Reset()         { *m = JunosPpp{} }
func (m *JunosPpp) String() string { return proto.CompactTextString(m) }
func (*JunosPpp) ProtoMessage()    {}
func (*JunosPpp) Descriptor() ([]byte, []int) {
	return fileDescriptor_jpppd_oc_9d16cf3ec198406d, []int{0}
}
func (m *JunosPpp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosPpp.Unmarshal(m, b)
}
func (m *JunosPpp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosPpp.Marshal(b, m, deterministic)
}
func (dst *JunosPpp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosPpp.Merge(dst, src)
}
func (m *JunosPpp) XXX_Size() int {
	return xxx_messageInfo_JunosPpp.Size(m)
}
func (m *JunosPpp) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosPpp.DiscardUnknown(m)
}

var xxx_messageInfo_JunosPpp proto.InternalMessageInfo

func (m *JunosPpp) GetSystem() *JunosPppSystemType {
	if m != nil {
		return m.System
	}
	return nil
}

type JunosPppSystemType struct {
	SubscriberManagement *JunosPppSystemTypeSubscriberManagementType `protobuf:"bytes,151,opt,name=subscriber_management,json=subscriberManagement" json:"subscriber_management,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *JunosPppSystemType) Reset()         { *m = JunosPppSystemType{} }
func (m *JunosPppSystemType) String() string { return proto.CompactTextString(m) }
func (*JunosPppSystemType) ProtoMessage()    {}
func (*JunosPppSystemType) Descriptor() ([]byte, []int) {
	return fileDescriptor_jpppd_oc_9d16cf3ec198406d, []int{0, 0}
}
func (m *JunosPppSystemType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosPppSystemType.Unmarshal(m, b)
}
func (m *JunosPppSystemType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosPppSystemType.Marshal(b, m, deterministic)
}
func (dst *JunosPppSystemType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosPppSystemType.Merge(dst, src)
}
func (m *JunosPppSystemType) XXX_Size() int {
	return xxx_messageInfo_JunosPppSystemType.Size(m)
}
func (m *JunosPppSystemType) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosPppSystemType.DiscardUnknown(m)
}

var xxx_messageInfo_JunosPppSystemType proto.InternalMessageInfo

func (m *JunosPppSystemType) GetSubscriberManagement() *JunosPppSystemTypeSubscriberManagementType {
	if m != nil {
		return m.SubscriberManagement
	}
	return nil
}

type JunosPppSystemTypeSubscriberManagementType struct {
	ClientProtocols      *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsType `protobuf:"bytes,151,opt,name=client_protocols,json=clientProtocols" json:"client_protocols,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                       `json:"-"`
	XXX_unrecognized     []byte                                                         `json:"-"`
	XXX_sizecache        int32                                                          `json:"-"`
}

func (m *JunosPppSystemTypeSubscriberManagementType) Reset() {
	*m = JunosPppSystemTypeSubscriberManagementType{}
}
func (m *JunosPppSystemTypeSubscriberManagementType) String() string {
	return proto.CompactTextString(m)
}
func (*JunosPppSystemTypeSubscriberManagementType) ProtoMessage() {}
func (*JunosPppSystemTypeSubscriberManagementType) Descriptor() ([]byte, []int) {
	return fileDescriptor_jpppd_oc_9d16cf3ec198406d, []int{0, 0, 0}
}
func (m *JunosPppSystemTypeSubscriberManagementType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosPppSystemTypeSubscriberManagementType.Unmarshal(m, b)
}
func (m *JunosPppSystemTypeSubscriberManagementType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosPppSystemTypeSubscriberManagementType.Marshal(b, m, deterministic)
}
func (dst *JunosPppSystemTypeSubscriberManagementType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosPppSystemTypeSubscriberManagementType.Merge(dst, src)
}
func (m *JunosPppSystemTypeSubscriberManagementType) XXX_Size() int {
	return xxx_messageInfo_JunosPppSystemTypeSubscriberManagementType.Size(m)
}
func (m *JunosPppSystemTypeSubscriberManagementType) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosPppSystemTypeSubscriberManagementType.DiscardUnknown(m)
}

var xxx_messageInfo_JunosPppSystemTypeSubscriberManagementType proto.InternalMessageInfo

func (m *JunosPppSystemTypeSubscriberManagementType) GetClientProtocols() *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsType {
	if m != nil {
		return m.ClientProtocols
	}
	return nil
}

type JunosPppSystemTypeSubscriberManagementTypeClientProtocolsType struct {
	Ppp                  *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppType `protobuf:"bytes,151,opt,name=ppp" json:"ppp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                              `json:"-"`
	XXX_unrecognized     []byte                                                                `json:"-"`
	XXX_sizecache        int32                                                                 `json:"-"`
}

func (m *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsType) Reset() {
	*m = JunosPppSystemTypeSubscriberManagementTypeClientProtocolsType{}
}
func (m *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsType) String() string {
	return proto.CompactTextString(m)
}
func (*JunosPppSystemTypeSubscriberManagementTypeClientProtocolsType) ProtoMessage() {}
func (*JunosPppSystemTypeSubscriberManagementTypeClientProtocolsType) Descriptor() ([]byte, []int) {
	return fileDescriptor_jpppd_oc_9d16cf3ec198406d, []int{0, 0, 0, 0}
}
func (m *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosPppSystemTypeSubscriberManagementTypeClientProtocolsType.Unmarshal(m, b)
}
func (m *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosPppSystemTypeSubscriberManagementTypeClientProtocolsType.Marshal(b, m, deterministic)
}
func (dst *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosPppSystemTypeSubscriberManagementTypeClientProtocolsType.Merge(dst, src)
}
func (m *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsType) XXX_Size() int {
	return xxx_messageInfo_JunosPppSystemTypeSubscriberManagementTypeClientProtocolsType.Size(m)
}
func (m *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsType) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosPppSystemTypeSubscriberManagementTypeClientProtocolsType.DiscardUnknown(m)
}

var xxx_messageInfo_JunosPppSystemTypeSubscriberManagementTypeClientProtocolsType proto.InternalMessageInfo

func (m *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsType) GetPpp() *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppType {
	if m != nil {
		return m.Ppp
	}
	return nil
}

type JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppType struct {
	Statistics           *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppTypeStatisticsType `protobuf:"bytes,151,opt,name=statistics" json:"statistics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                            `json:"-"`
	XXX_unrecognized     []byte                                                                              `json:"-"`
	XXX_sizecache        int32                                                                               `json:"-"`
}

func (m *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppType) Reset() {
	*m = JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppType{}
}
func (m *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppType) String() string {
	return proto.CompactTextString(m)
}
func (*JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppType) ProtoMessage() {}
func (*JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppType) Descriptor() ([]byte, []int) {
	return fileDescriptor_jpppd_oc_9d16cf3ec198406d, []int{0, 0, 0, 0, 0}
}
func (m *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppType.Unmarshal(m, b)
}
func (m *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppType.Marshal(b, m, deterministic)
}
func (dst *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppType.Merge(dst, src)
}
func (m *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppType) XXX_Size() int {
	return xxx_messageInfo_JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppType.Size(m)
}
func (m *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppType) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppType.DiscardUnknown(m)
}

var xxx_messageInfo_JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppType proto.InternalMessageInfo

func (m *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppType) GetStatistics() *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppTypeStatisticsType {
	if m != nil {
		return m.Statistics
	}
	return nil
}

type JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppTypeStatisticsType struct {
	PppStatsTotalSubscriberSessions   *uint32  `protobuf:"varint,51,opt,name=ppp_stats_total_subscriber_sessions,json=pppStatsTotalSubscriberSessions" json:"ppp_stats_total_subscriber_sessions,omitempty"`
	PppStatsSessionsDisablePhase      *uint32  `protobuf:"varint,52,opt,name=ppp_stats_sessions_disable_phase,json=pppStatsSessionsDisablePhase" json:"ppp_stats_sessions_disable_phase,omitempty"`
	PppStatsSessionsEstablishPhase    *uint32  `protobuf:"varint,53,opt,name=ppp_stats_sessions_establish_phase,json=pppStatsSessionsEstablishPhase" json:"ppp_stats_sessions_establish_phase,omitempty"`
	PppStatsSessionsNetworkPhase      *uint32  `protobuf:"varint,54,opt,name=ppp_stats_sessions_network_phase,json=pppStatsSessionsNetworkPhase" json:"ppp_stats_sessions_network_phase,omitempty"`
	PppStatsSessionsAuthenticatePhase *uint32  `protobuf:"varint,55,opt,name=ppp_stats_sessions_authenticate_phase,json=pppStatsSessionsAuthenticatePhase" json:"ppp_stats_sessions_authenticate_phase,omitempty"`
	XXX_NoUnkeyedLiteral              struct{} `json:"-"`
	XXX_unrecognized                  []byte   `json:"-"`
	XXX_sizecache                     int32    `json:"-"`
}

func (m *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppTypeStatisticsType) Reset() {
	*m = JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppTypeStatisticsType{}
}
func (m *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppTypeStatisticsType) String() string {
	return proto.CompactTextString(m)
}
func (*JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppTypeStatisticsType) ProtoMessage() {
}
func (*JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppTypeStatisticsType) Descriptor() ([]byte, []int) {
	return fileDescriptor_jpppd_oc_9d16cf3ec198406d, []int{0, 0, 0, 0, 0, 0}
}
func (m *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppTypeStatisticsType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppTypeStatisticsType.Unmarshal(m, b)
}
func (m *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppTypeStatisticsType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppTypeStatisticsType.Marshal(b, m, deterministic)
}
func (dst *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppTypeStatisticsType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppTypeStatisticsType.Merge(dst, src)
}
func (m *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppTypeStatisticsType) XXX_Size() int {
	return xxx_messageInfo_JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppTypeStatisticsType.Size(m)
}
func (m *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppTypeStatisticsType) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppTypeStatisticsType.DiscardUnknown(m)
}

var xxx_messageInfo_JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppTypeStatisticsType proto.InternalMessageInfo

func (m *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppTypeStatisticsType) GetPppStatsTotalSubscriberSessions() uint32 {
	if m != nil && m.PppStatsTotalSubscriberSessions != nil {
		return *m.PppStatsTotalSubscriberSessions
	}
	return 0
}

func (m *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppTypeStatisticsType) GetPppStatsSessionsDisablePhase() uint32 {
	if m != nil && m.PppStatsSessionsDisablePhase != nil {
		return *m.PppStatsSessionsDisablePhase
	}
	return 0
}

func (m *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppTypeStatisticsType) GetPppStatsSessionsEstablishPhase() uint32 {
	if m != nil && m.PppStatsSessionsEstablishPhase != nil {
		return *m.PppStatsSessionsEstablishPhase
	}
	return 0
}

func (m *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppTypeStatisticsType) GetPppStatsSessionsNetworkPhase() uint32 {
	if m != nil && m.PppStatsSessionsNetworkPhase != nil {
		return *m.PppStatsSessionsNetworkPhase
	}
	return 0
}

func (m *JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppTypeStatisticsType) GetPppStatsSessionsAuthenticatePhase() uint32 {
	if m != nil && m.PppStatsSessionsAuthenticatePhase != nil {
		return *m.PppStatsSessionsAuthenticatePhase
	}
	return 0
}

var E_JnprJunosPppExt = &proto.ExtensionDesc{
	ExtendedType:  (*JuniperNetworksSensors)(nil),
	ExtensionType: (*JunosPpp)(nil),
	Field:         45,
	Name:          "jnpr_junos_ppp_ext",
	Tag:           "bytes,45,opt,name=jnpr_junos_ppp_ext,json=jnprJunosPppExt",
	Filename:      "jpppd_oc.proto",
}

func init() {
	proto.RegisterType((*JunosPpp)(nil), "junos_ppp")
	proto.RegisterType((*JunosPppSystemType)(nil), "junos_ppp.system_type")
	proto.RegisterType((*JunosPppSystemTypeSubscriberManagementType)(nil), "junos_ppp.system_type.subscriber_management_type")
	proto.RegisterType((*JunosPppSystemTypeSubscriberManagementTypeClientProtocolsType)(nil), "junos_ppp.system_type.subscriber_management_type.client_protocols_type")
	proto.RegisterType((*JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppType)(nil), "junos_ppp.system_type.subscriber_management_type.client_protocols_type.ppp_type")
	proto.RegisterType((*JunosPppSystemTypeSubscriberManagementTypeClientProtocolsTypePppTypeStatisticsType)(nil), "junos_ppp.system_type.subscriber_management_type.client_protocols_type.ppp_type.statistics_type")
	proto.RegisterExtension(E_JnprJunosPppExt)
}

func init() { proto.RegisterFile("jpppd_oc.proto", fileDescriptor_jpppd_oc_9d16cf3ec198406d) }

var fileDescriptor_jpppd_oc_9d16cf3ec198406d = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcf, 0x8a, 0x13, 0x41,
	0x10, 0xc6, 0x09, 0xab, 0xbb, 0x5a, 0x8b, 0x46, 0x5a, 0x57, 0xc3, 0x20, 0x1a, 0x57, 0x84, 0xbd,
	0x38, 0xe2, 0x7f, 0xf0, 0x26, 0xb8, 0x0a, 0x41, 0x65, 0x48, 0xbc, 0xb7, 0x9d, 0xd9, 0x66, 0xd3,
	0x71, 0xd2, 0x5d, 0x4c, 0x55, 0x70, 0x73, 0xf4, 0x01, 0x3c, 0x7b, 0xd2, 0x07, 0xf0, 0xe6, 0x63,
	0xf9, 0x16, 0xd2, 0xd3, 0xd3, 0x93, 0x30, 0x8c, 0x87, 0x85, 0x1c, 0xbb, 0xea, 0xfb, 0x7e, 0xdf,
	0xd7, 0x50, 0x70, 0x75, 0x8e, 0x88, 0x27, 0xd2, 0xe5, 0x29, 0x96, 0x8e, 0x5d, 0x72, 0x9d, 0x75,
	0xa1, 0x17, 0x9a, 0xcb, 0x95, 0x64, 0x87, 0x61, 0x78, 0xf8, 0x77, 0x0f, 0x2e, 0xcf, 0x97, 0xd6,
	0x91, 0x44, 0x44, 0xf1, 0x08, 0x76, 0x69, 0x45, 0xac, 0x17, 0x83, 0x1f, 0xbd, 0x61, 0xef, 0x68,
	0xff, 0xc9, 0xcd, 0xb4, 0x59, 0xa6, 0x61, 0x23, 0x79, 0x85, 0x7a, 0x5c, 0xcb, 0x92, 0xdf, 0x7b,
	0xb0, 0xbf, 0x31, 0x17, 0xa7, 0x70, 0x40, 0xcb, 0x29, 0xe5, 0xa5, 0x99, 0xea, 0x52, 0x2e, 0x94,
	0x55, 0xa7, 0x7a, 0xa1, 0x2d, 0x47, 0xde, 0xe3, 0x6e, 0x5e, 0xda, 0x69, 0x0a, 0x51, 0x37, 0xd6,
	0xbb, 0x0f, 0xcd, 0x2a, 0xf9, 0xb3, 0x0b, 0xc9, 0xff, 0x4d, 0x82, 0xe0, 0x5a, 0x5e, 0x18, 0xff,
	0xac, 0xbe, 0x99, 0xbb, 0x82, 0x62, 0x85, 0x77, 0xe7, 0xae, 0x90, 0xb6, 0x51, 0xa1, 0x58, 0x3f,
	0x8c, 0xb3, 0x38, 0x4d, 0x7e, 0x5e, 0x84, 0x83, 0x4e, 0xa9, 0xc8, 0x61, 0x07, 0x11, 0x63, 0x83,
	0x6c, 0x4b, 0x0d, 0x52, 0x44, 0x0c, 0x55, 0x3c, 0x3d, 0xf9, 0x75, 0x01, 0x2e, 0xc5, 0x89, 0xf8,
	0xd6, 0x03, 0x20, 0x56, 0x6c, 0x88, 0x4d, 0xde, 0xfc, 0xfd, 0xf3, 0xb6, 0x93, 0xd3, 0x75, 0x46,
	0x68, 0xb2, 0x11, 0x9a, 0x7c, 0xdf, 0x81, 0x7e, 0x6b, 0x2f, 0xde, 0xc3, 0x7d, 0xef, 0xf5, 0x63,
	0x92, 0xec, 0x58, 0x15, 0x72, 0x23, 0x9c, 0x34, 0x91, 0x71, 0x96, 0x06, 0x4f, 0x87, 0xbd, 0xa3,
	0x2b, 0xe3, 0xbb, 0x88, 0x38, 0xf1, 0xca, 0x4f, 0x5e, 0x38, 0x69, 0x74, 0x93, 0x5a, 0x26, 0xde,
	0xc2, 0x70, 0x4d, 0x8b, 0x66, 0x79, 0x62, 0x48, 0x4d, 0x0b, 0x2d, 0x71, 0xa6, 0x48, 0x0f, 0x9e,
	0x55, 0xa8, 0xdb, 0x11, 0x15, 0xbd, 0x6f, 0x82, 0x28, 0xf3, 0x1a, 0x31, 0x82, 0xc3, 0x0e, 0x8e,
	0x26, 0x56, 0xd3, 0xc2, 0xd0, 0xac, 0x26, 0x3d, 0xaf, 0x48, 0x77, 0xda, 0xa4, 0xe3, 0x28, 0x0b,
	0xac, 0xee, 0x4e, 0x56, 0xf3, 0x57, 0x57, 0x7e, 0xa9, 0x49, 0x2f, 0xba, 0x3b, 0x7d, 0x0c, 0xa2,
	0xc0, 0xc9, 0xe0, 0x41, 0x07, 0x47, 0x2d, 0x79, 0xa6, 0x2d, 0x9b, 0x5c, 0x71, 0xfc, 0xe0, 0xcb,
	0x0a, 0x76, 0xaf, 0x0d, 0x7b, 0xbd, 0xa1, 0xac, 0x88, 0xaf, 0x32, 0x10, 0x73, 0x8b, 0xa5, 0x6c,
	0x6e, 0x40, 0xea, 0x33, 0x16, 0xb7, 0xd2, 0xd1, 0xd2, 0x1a, 0xd4, 0x65, 0x9d, 0x4e, 0x13, 0x6d,
	0xc9, 0x95, 0x34, 0x78, 0x58, 0x9d, 0x0c, 0xac, 0x4f, 0x66, 0xdc, 0xf7, 0xf6, 0x91, 0x7f, 0x66,
	0x88, 0xc7, 0x67, 0xfc, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x30, 0xf3, 0x18, 0x0a, 0x63, 0x04, 0x00,
	0x00,
}
