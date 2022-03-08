// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ldp_p2mp_lsp_stats.proto

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

//
// Top-level message
//
type LdpP2MpLspStats struct {
	// List of LDP P2MP LSP stats record
	LdpP2MpLspStatsRecords []*LdpP2MpLspRecord `protobuf:"bytes,1,rep,name=ldp_p2mp_lsp_stats_records,json=ldpP2mpLspStatsRecords" json:"ldp_p2mp_lsp_stats_records,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}            `json:"-"`
	XXX_unrecognized       []byte              `json:"-"`
	XXX_sizecache          int32               `json:"-"`
}

func (m *LdpP2MpLspStats) Reset()         { *m = LdpP2MpLspStats{} }
func (m *LdpP2MpLspStats) String() string { return proto.CompactTextString(m) }
func (*LdpP2MpLspStats) ProtoMessage()    {}
func (*LdpP2MpLspStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_de3095a4de0e0ad1, []int{0}
}
func (m *LdpP2MpLspStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpP2MpLspStats.Unmarshal(m, b)
}
func (m *LdpP2MpLspStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpP2MpLspStats.Marshal(b, m, deterministic)
}
func (m *LdpP2MpLspStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpP2MpLspStats.Merge(m, src)
}
func (m *LdpP2MpLspStats) XXX_Size() int {
	return xxx_messageInfo_LdpP2MpLspStats.Size(m)
}
func (m *LdpP2MpLspStats) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpP2MpLspStats.DiscardUnknown(m)
}

var xxx_messageInfo_LdpP2MpLspStats proto.InternalMessageInfo

func (m *LdpP2MpLspStats) GetLdpP2MpLspStatsRecords() []*LdpP2MpLspRecord {
	if m != nil {
		return m.LdpP2MpLspStatsRecords
	}
	return nil
}

type LdpP2MpLspRecord struct {
	// IP prefix
	RootAddress *string `protobuf:"bytes,1,req,name=root_address,json=rootAddress" json:"root_address,omitempty"`
	LspId       *uint32 `protobuf:"varint,2,opt,name=lsp_id,json=lspId" json:"lsp_id,omitempty"`
	// Source Ip address
	SourceIp *string `protobuf:"bytes,3,opt,name=source_ip,json=sourceIp" json:"source_ip,omitempty"`
	// Group Ip address
	GroupIp *string `protobuf:"bytes,4,opt,name=group_ip,json=groupIp" json:"group_ip,omitempty"`
	// Instance Identifier for cases when RPD creates multiple instances
	InstanceIdentifier *uint32 `protobuf:"varint,5,opt,name=instance_identifier,json=instanceIdentifier" json:"instance_identifier,omitempty"`
	// Name of the counter.
	CounterName *string `protobuf:"bytes,6,opt,name=counter_name,json=counterName" json:"counter_name,omitempty"`
	// Statistics
	Stats                *LabelDistributionProtocolP2MpLspStats `protobuf:"bytes,7,opt,name=stats" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *LdpP2MpLspRecord) Reset()         { *m = LdpP2MpLspRecord{} }
func (m *LdpP2MpLspRecord) String() string { return proto.CompactTextString(m) }
func (*LdpP2MpLspRecord) ProtoMessage()    {}
func (*LdpP2MpLspRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_de3095a4de0e0ad1, []int{1}
}
func (m *LdpP2MpLspRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpP2MpLspRecord.Unmarshal(m, b)
}
func (m *LdpP2MpLspRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpP2MpLspRecord.Marshal(b, m, deterministic)
}
func (m *LdpP2MpLspRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpP2MpLspRecord.Merge(m, src)
}
func (m *LdpP2MpLspRecord) XXX_Size() int {
	return xxx_messageInfo_LdpP2MpLspRecord.Size(m)
}
func (m *LdpP2MpLspRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpP2MpLspRecord.DiscardUnknown(m)
}

var xxx_messageInfo_LdpP2MpLspRecord proto.InternalMessageInfo

func (m *LdpP2MpLspRecord) GetRootAddress() string {
	if m != nil && m.RootAddress != nil {
		return *m.RootAddress
	}
	return ""
}

func (m *LdpP2MpLspRecord) GetLspId() uint32 {
	if m != nil && m.LspId != nil {
		return *m.LspId
	}
	return 0
}

func (m *LdpP2MpLspRecord) GetSourceIp() string {
	if m != nil && m.SourceIp != nil {
		return *m.SourceIp
	}
	return ""
}

func (m *LdpP2MpLspRecord) GetGroupIp() string {
	if m != nil && m.GroupIp != nil {
		return *m.GroupIp
	}
	return ""
}

func (m *LdpP2MpLspRecord) GetInstanceIdentifier() uint32 {
	if m != nil && m.InstanceIdentifier != nil {
		return *m.InstanceIdentifier
	}
	return 0
}

func (m *LdpP2MpLspRecord) GetCounterName() string {
	if m != nil && m.CounterName != nil {
		return *m.CounterName
	}
	return ""
}

func (m *LdpP2MpLspRecord) GetStats() *LabelDistributionProtocolP2MpLspStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

type LabelDistributionProtocolP2MpLspStats struct {
	// Packet and Byte statistics
	Packets *uint64 `protobuf:"varint,1,opt,name=packets" json:"packets,omitempty"`
	Bytes   *uint64 `protobuf:"varint,2,opt,name=bytes" json:"bytes,omitempty"`
	// Rates of the above counters
	PacketRate           *uint64  `protobuf:"varint,3,opt,name=packet_rate,json=packetRate" json:"packet_rate,omitempty"`
	ByteRate             *uint64  `protobuf:"varint,4,opt,name=byte_rate,json=byteRate" json:"byte_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LabelDistributionProtocolP2MpLspStats) Reset()         { *m = LabelDistributionProtocolP2MpLspStats{} }
func (m *LabelDistributionProtocolP2MpLspStats) String() string { return proto.CompactTextString(m) }
func (*LabelDistributionProtocolP2MpLspStats) ProtoMessage()    {}
func (*LabelDistributionProtocolP2MpLspStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_de3095a4de0e0ad1, []int{2}
}
func (m *LabelDistributionProtocolP2MpLspStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabelDistributionProtocolP2MpLspStats.Unmarshal(m, b)
}
func (m *LabelDistributionProtocolP2MpLspStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabelDistributionProtocolP2MpLspStats.Marshal(b, m, deterministic)
}
func (m *LabelDistributionProtocolP2MpLspStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelDistributionProtocolP2MpLspStats.Merge(m, src)
}
func (m *LabelDistributionProtocolP2MpLspStats) XXX_Size() int {
	return xxx_messageInfo_LabelDistributionProtocolP2MpLspStats.Size(m)
}
func (m *LabelDistributionProtocolP2MpLspStats) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelDistributionProtocolP2MpLspStats.DiscardUnknown(m)
}

var xxx_messageInfo_LabelDistributionProtocolP2MpLspStats proto.InternalMessageInfo

func (m *LabelDistributionProtocolP2MpLspStats) GetPackets() uint64 {
	if m != nil && m.Packets != nil {
		return *m.Packets
	}
	return 0
}

func (m *LabelDistributionProtocolP2MpLspStats) GetBytes() uint64 {
	if m != nil && m.Bytes != nil {
		return *m.Bytes
	}
	return 0
}

func (m *LabelDistributionProtocolP2MpLspStats) GetPacketRate() uint64 {
	if m != nil && m.PacketRate != nil {
		return *m.PacketRate
	}
	return 0
}

func (m *LabelDistributionProtocolP2MpLspStats) GetByteRate() uint64 {
	if m != nil && m.ByteRate != nil {
		return *m.ByteRate
	}
	return 0
}

var E_JnprLdpP2MpLspStatsExt = &proto.ExtensionDesc{
	ExtendedType:  (*JuniperNetworksSensors)(nil),
	ExtensionType: (*LdpP2MpLspStats)(nil),
	Field:         155,
	Name:          "jnpr_ldp_p2mp_lsp_stats_ext",
	Tag:           "bytes,155,opt,name=jnpr_ldp_p2mp_lsp_stats_ext",
	Filename:      "ldp_p2mp_lsp_stats.proto",
}

func init() {
	proto.RegisterType((*LdpP2MpLspStats)(nil), "LdpP2mpLspStats")
	proto.RegisterType((*LdpP2MpLspRecord)(nil), "LdpP2mpLspRecord")
	proto.RegisterType((*LabelDistributionProtocolP2MpLspStats)(nil), "LabelDistributionProtocolP2mpLspStats")
	proto.RegisterExtension(E_JnprLdpP2MpLspStatsExt)
}

func init() { proto.RegisterFile("ldp_p2mp_lsp_stats.proto", fileDescriptor_de3095a4de0e0ad1) }

var fileDescriptor_de3095a4de0e0ad1 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6e, 0xd4, 0x30,
	0x18, 0x85, 0x95, 0xe9, 0xa4, 0x33, 0x75, 0x40, 0x14, 0x57, 0x2a, 0x56, 0x8b, 0x44, 0x34, 0x12,
	0x55, 0x56, 0x59, 0xcc, 0x82, 0x05, 0x62, 0x01, 0x08, 0x16, 0x83, 0x86, 0xaa, 0x4a, 0x0f, 0x60,
	0x3c, 0xc9, 0x4f, 0x65, 0x9a, 0xd8, 0xbf, 0xec, 0x3f, 0xa2, 0xdd, 0x72, 0x0d, 0x4e, 0xc1, 0x0d,
	0x38, 0x1a, 0x72, 0xcc, 0x4c, 0xa7, 0x03, 0x8b, 0x6e, 0xdf, 0xfb, 0xde, 0x8b, 0xf3, 0x6c, 0x26,
	0xda, 0x06, 0x25, 0xce, 0x3b, 0x94, 0xad, 0x47, 0xe9, 0x49, 0x91, 0x2f, 0xd1, 0x59, 0xb2, 0x27,
	0x47, 0x04, 0x2d, 0x74, 0x40, 0xee, 0x56, 0x92, 0xc5, 0x28, 0xce, 0xbe, 0xb0, 0x27, 0xcb, 0x06,
	0x2f, 0xe6, 0x1d, 0x2e, 0x3d, 0x5e, 0x06, 0x9a, 0x7f, 0x66, 0x27, 0xff, 0x76, 0x48, 0x07, 0xb5,
	0x75, 0x8d, 0x17, 0x49, 0xbe, 0x57, 0x64, 0xf3, 0xa7, 0xe5, 0x5d, 0xaa, 0x1a, 0x9c, 0xea, 0xb8,
	0xbd, 0xdf, 0x13, 0x65, 0x3f, 0xfb, 0x3d, 0x62, 0x87, 0xbb, 0x30, 0x2f, 0xd8, 0x23, 0x67, 0x2d,
	0x49, 0xd5, 0x34, 0x0e, 0x7c, 0x68, 0x1d, 0x15, 0x07, 0xef, 0xd3, 0x1f, 0x6f, 0x47, 0xd3, 0xa4,
	0xca, 0x82, 0xf5, 0x2e, 0x3a, 0xfc, 0x39, 0xdb, 0x0f, 0x87, 0xd0, 0x8d, 0x18, 0xe5, 0x49, 0xf1,
	0x78, 0xcd, 0xa4, 0xad, 0xc7, 0x45, 0xc3, 0x67, 0xec, 0xc0, 0xdb, 0xde, 0xd5, 0x20, 0x35, 0x8a,
	0xbd, 0x3c, 0xb9, 0x2b, 0x99, 0x46, 0x7d, 0x81, 0x3c, 0x67, 0xd3, 0x2b, 0x67, 0x7b, 0x0c, 0xc8,
	0x78, 0x1b, 0x99, 0x0c, 0xf2, 0x02, 0xf9, 0x2b, 0x76, 0xa4, 0x8d, 0x27, 0x65, 0x42, 0x4f, 0x03,
	0x86, 0xf4, 0x57, 0x0d, 0x4e, 0xa4, 0xdb, 0x1f, 0xe4, 0x6b, 0x62, 0xb1, 0x01, 0xc2, 0x5f, 0xd4,
	0xb6, 0x37, 0x04, 0x4e, 0x1a, 0xd5, 0x81, 0xd8, 0xdf, 0x6e, 0xcf, 0xfe, 0x5a, 0xe7, 0xaa, 0x03,
	0xfe, 0x86, 0xa5, 0xc3, 0x8c, 0x62, 0x92, 0x27, 0x45, 0x36, 0x3f, 0x2b, 0x97, 0x6a, 0x05, 0xed,
	0x07, 0xed, 0xc9, 0xe9, 0x55, 0x4f, 0xda, 0x9a, 0x8b, 0x70, 0x1d, 0xb5, 0x6d, 0xef, 0x4d, 0x18,
	0x43, 0xb3, 0x5f, 0x09, 0x7b, 0xf9, 0xa0, 0x00, 0x7f, 0xc1, 0x26, 0xa8, 0xea, 0x6b, 0xa0, 0x30,
	0x69, 0x52, 0x8c, 0x87, 0xc3, 0x88, 0xa4, 0x5a, 0xab, 0xfc, 0x94, 0xa5, 0xab, 0x5b, 0x02, 0x3f,
	0xac, 0xb9, 0xb1, 0xa3, 0xc6, 0xcf, 0x58, 0x16, 0x39, 0xe9, 0x14, 0xc1, 0xb0, 0xe7, 0x06, 0x61,
	0xd1, 0xa9, 0x14, 0x41, 0x58, 0x3d, 0x04, 0x22, 0x35, 0xde, 0xa6, 0xa6, 0x41, 0x0f, 0xcc, 0xeb,
	0x2b, 0x76, 0xfa, 0xcd, 0xa0, 0x93, 0xff, 0x79, 0x4a, 0x70, 0x43, 0xfc, 0x59, 0xf9, 0xa9, 0x37,
	0x1a, 0xc1, 0x9d, 0x03, 0x7d, 0xb7, 0xee, 0xda, 0x5f, 0x82, 0xf1, 0xd6, 0x79, 0xf1, 0x33, 0x19,
	0x16, 0x3a, 0x2c, 0x77, 0x9e, 0x65, 0x75, 0x1c, 0xea, 0x76, 0xc4, 0x8f, 0x37, 0xf4, 0x27, 0x00,
	0x00, 0xff, 0xff, 0x90, 0x6c, 0x67, 0x57, 0xf1, 0x02, 0x00, 0x00,
}
