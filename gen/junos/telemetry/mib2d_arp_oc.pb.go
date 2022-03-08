// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mib2d_arp_oc.proto

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

type ArpInformationMibArp struct {
	Ipv4                 *ArpInformationMibArpIpv4Type `protobuf:"bytes,151,opt,name=ipv4" json:"ipv4,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ArpInformationMibArp) Reset()         { *m = ArpInformationMibArp{} }
func (m *ArpInformationMibArp) String() string { return proto.CompactTextString(m) }
func (*ArpInformationMibArp) ProtoMessage()    {}
func (*ArpInformationMibArp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6609990a88c6913e, []int{0}
}
func (m *ArpInformationMibArp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArpInformationMibArp.Unmarshal(m, b)
}
func (m *ArpInformationMibArp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArpInformationMibArp.Marshal(b, m, deterministic)
}
func (m *ArpInformationMibArp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArpInformationMibArp.Merge(m, src)
}
func (m *ArpInformationMibArp) XXX_Size() int {
	return xxx_messageInfo_ArpInformationMibArp.Size(m)
}
func (m *ArpInformationMibArp) XXX_DiscardUnknown() {
	xxx_messageInfo_ArpInformationMibArp.DiscardUnknown(m)
}

var xxx_messageInfo_ArpInformationMibArp proto.InternalMessageInfo

func (m *ArpInformationMibArp) GetIpv4() *ArpInformationMibArpIpv4Type {
	if m != nil {
		return m.Ipv4
	}
	return nil
}

type ArpInformationMibArpIpv4Type struct {
	Neighbors            *ArpInformationMibArpIpv4TypeNeighborsType `protobuf:"bytes,151,opt,name=neighbors" json:"neighbors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *ArpInformationMibArpIpv4Type) Reset()         { *m = ArpInformationMibArpIpv4Type{} }
func (m *ArpInformationMibArpIpv4Type) String() string { return proto.CompactTextString(m) }
func (*ArpInformationMibArpIpv4Type) ProtoMessage()    {}
func (*ArpInformationMibArpIpv4Type) Descriptor() ([]byte, []int) {
	return fileDescriptor_6609990a88c6913e, []int{0, 0}
}
func (m *ArpInformationMibArpIpv4Type) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArpInformationMibArpIpv4Type.Unmarshal(m, b)
}
func (m *ArpInformationMibArpIpv4Type) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArpInformationMibArpIpv4Type.Marshal(b, m, deterministic)
}
func (m *ArpInformationMibArpIpv4Type) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArpInformationMibArpIpv4Type.Merge(m, src)
}
func (m *ArpInformationMibArpIpv4Type) XXX_Size() int {
	return xxx_messageInfo_ArpInformationMibArpIpv4Type.Size(m)
}
func (m *ArpInformationMibArpIpv4Type) XXX_DiscardUnknown() {
	xxx_messageInfo_ArpInformationMibArpIpv4Type.DiscardUnknown(m)
}

var xxx_messageInfo_ArpInformationMibArpIpv4Type proto.InternalMessageInfo

func (m *ArpInformationMibArpIpv4Type) GetNeighbors() *ArpInformationMibArpIpv4TypeNeighborsType {
	if m != nil {
		return m.Neighbors
	}
	return nil
}

type ArpInformationMibArpIpv4TypeNeighborsType struct {
	Neighbor             []*ArpInformationMibArpIpv4TypeNeighborsTypeNeighborList `protobuf:"bytes,151,rep,name=neighbor" json:"neighbor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                 `json:"-"`
	XXX_unrecognized     []byte                                                   `json:"-"`
	XXX_sizecache        int32                                                    `json:"-"`
}

func (m *ArpInformationMibArpIpv4TypeNeighborsType) Reset() {
	*m = ArpInformationMibArpIpv4TypeNeighborsType{}
}
func (m *ArpInformationMibArpIpv4TypeNeighborsType) String() string {
	return proto.CompactTextString(m)
}
func (*ArpInformationMibArpIpv4TypeNeighborsType) ProtoMessage() {}
func (*ArpInformationMibArpIpv4TypeNeighborsType) Descriptor() ([]byte, []int) {
	return fileDescriptor_6609990a88c6913e, []int{0, 0, 0}
}
func (m *ArpInformationMibArpIpv4TypeNeighborsType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArpInformationMibArpIpv4TypeNeighborsType.Unmarshal(m, b)
}
func (m *ArpInformationMibArpIpv4TypeNeighborsType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArpInformationMibArpIpv4TypeNeighborsType.Marshal(b, m, deterministic)
}
func (m *ArpInformationMibArpIpv4TypeNeighborsType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArpInformationMibArpIpv4TypeNeighborsType.Merge(m, src)
}
func (m *ArpInformationMibArpIpv4TypeNeighborsType) XXX_Size() int {
	return xxx_messageInfo_ArpInformationMibArpIpv4TypeNeighborsType.Size(m)
}
func (m *ArpInformationMibArpIpv4TypeNeighborsType) XXX_DiscardUnknown() {
	xxx_messageInfo_ArpInformationMibArpIpv4TypeNeighborsType.DiscardUnknown(m)
}

var xxx_messageInfo_ArpInformationMibArpIpv4TypeNeighborsType proto.InternalMessageInfo

func (m *ArpInformationMibArpIpv4TypeNeighborsType) GetNeighbor() []*ArpInformationMibArpIpv4TypeNeighborsTypeNeighborList {
	if m != nil {
		return m.Neighbor
	}
	return nil
}

type ArpInformationMibArpIpv4TypeNeighborsTypeNeighborList struct {
	Ip                   *string                                                         `protobuf:"bytes,51,opt,name=ip" json:"ip,omitempty"`
	State                *ArpInformationMibArpIpv4TypeNeighborsTypeNeighborListStateType `protobuf:"bytes,151,opt,name=state" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                        `json:"-"`
	XXX_unrecognized     []byte                                                          `json:"-"`
	XXX_sizecache        int32                                                           `json:"-"`
}

func (m *ArpInformationMibArpIpv4TypeNeighborsTypeNeighborList) Reset() {
	*m = ArpInformationMibArpIpv4TypeNeighborsTypeNeighborList{}
}
func (m *ArpInformationMibArpIpv4TypeNeighborsTypeNeighborList) String() string {
	return proto.CompactTextString(m)
}
func (*ArpInformationMibArpIpv4TypeNeighborsTypeNeighborList) ProtoMessage() {}
func (*ArpInformationMibArpIpv4TypeNeighborsTypeNeighborList) Descriptor() ([]byte, []int) {
	return fileDescriptor_6609990a88c6913e, []int{0, 0, 0, 0}
}
func (m *ArpInformationMibArpIpv4TypeNeighborsTypeNeighborList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArpInformationMibArpIpv4TypeNeighborsTypeNeighborList.Unmarshal(m, b)
}
func (m *ArpInformationMibArpIpv4TypeNeighborsTypeNeighborList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArpInformationMibArpIpv4TypeNeighborsTypeNeighborList.Marshal(b, m, deterministic)
}
func (m *ArpInformationMibArpIpv4TypeNeighborsTypeNeighborList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArpInformationMibArpIpv4TypeNeighborsTypeNeighborList.Merge(m, src)
}
func (m *ArpInformationMibArpIpv4TypeNeighborsTypeNeighborList) XXX_Size() int {
	return xxx_messageInfo_ArpInformationMibArpIpv4TypeNeighborsTypeNeighborList.Size(m)
}
func (m *ArpInformationMibArpIpv4TypeNeighborsTypeNeighborList) XXX_DiscardUnknown() {
	xxx_messageInfo_ArpInformationMibArpIpv4TypeNeighborsTypeNeighborList.DiscardUnknown(m)
}

var xxx_messageInfo_ArpInformationMibArpIpv4TypeNeighborsTypeNeighborList proto.InternalMessageInfo

func (m *ArpInformationMibArpIpv4TypeNeighborsTypeNeighborList) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *ArpInformationMibArpIpv4TypeNeighborsTypeNeighborList) GetState() *ArpInformationMibArpIpv4TypeNeighborsTypeNeighborListStateType {
	if m != nil {
		return m.State
	}
	return nil
}

type ArpInformationMibArpIpv4TypeNeighborsTypeNeighborListStateType struct {
	Ip                   *string  `protobuf:"bytes,51,opt,name=ip" json:"ip,omitempty"`
	LinkLayerAddress     *string  `protobuf:"bytes,52,opt,name=link_layer_address,json=linkLayerAddress" json:"link_layer_address,omitempty"`
	Origin               *string  `protobuf:"bytes,53,opt,name=origin" json:"origin,omitempty"`
	HostName             *string  `protobuf:"bytes,61,opt,name=host_name,json=hostName" json:"host_name,omitempty"`
	TableId              *uint32  `protobuf:"varint,62,opt,name=table_id,json=tableId" json:"table_id,omitempty"`
	NeighborState        *string  `protobuf:"bytes,63,opt,name=neighbor_state,json=neighborState" json:"neighbor_state,omitempty"`
	Expiry               *uint32  `protobuf:"varint,64,opt,name=expiry" json:"expiry,omitempty"`
	IsPublish            *bool    `protobuf:"varint,65,opt,name=is_publish,json=isPublish" json:"is_publish,omitempty"`
	InterfaceName        *string  `protobuf:"bytes,66,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	LogicalRouterId      *uint32  `protobuf:"varint,67,opt,name=logical_router_id,json=logicalRouterId" json:"logical_router_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArpInformationMibArpIpv4TypeNeighborsTypeNeighborListStateType) Reset() {
	*m = ArpInformationMibArpIpv4TypeNeighborsTypeNeighborListStateType{}
}
func (m *ArpInformationMibArpIpv4TypeNeighborsTypeNeighborListStateType) String() string {
	return proto.CompactTextString(m)
}
func (*ArpInformationMibArpIpv4TypeNeighborsTypeNeighborListStateType) ProtoMessage() {}
func (*ArpInformationMibArpIpv4TypeNeighborsTypeNeighborListStateType) Descriptor() ([]byte, []int) {
	return fileDescriptor_6609990a88c6913e, []int{0, 0, 0, 0, 0}
}
func (m *ArpInformationMibArpIpv4TypeNeighborsTypeNeighborListStateType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArpInformationMibArpIpv4TypeNeighborsTypeNeighborListStateType.Unmarshal(m, b)
}
func (m *ArpInformationMibArpIpv4TypeNeighborsTypeNeighborListStateType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArpInformationMibArpIpv4TypeNeighborsTypeNeighborListStateType.Marshal(b, m, deterministic)
}
func (m *ArpInformationMibArpIpv4TypeNeighborsTypeNeighborListStateType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArpInformationMibArpIpv4TypeNeighborsTypeNeighborListStateType.Merge(m, src)
}
func (m *ArpInformationMibArpIpv4TypeNeighborsTypeNeighborListStateType) XXX_Size() int {
	return xxx_messageInfo_ArpInformationMibArpIpv4TypeNeighborsTypeNeighborListStateType.Size(m)
}
func (m *ArpInformationMibArpIpv4TypeNeighborsTypeNeighborListStateType) XXX_DiscardUnknown() {
	xxx_messageInfo_ArpInformationMibArpIpv4TypeNeighborsTypeNeighborListStateType.DiscardUnknown(m)
}

var xxx_messageInfo_ArpInformationMibArpIpv4TypeNeighborsTypeNeighborListStateType proto.InternalMessageInfo

func (m *ArpInformationMibArpIpv4TypeNeighborsTypeNeighborListStateType) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *ArpInformationMibArpIpv4TypeNeighborsTypeNeighborListStateType) GetLinkLayerAddress() string {
	if m != nil && m.LinkLayerAddress != nil {
		return *m.LinkLayerAddress
	}
	return ""
}

func (m *ArpInformationMibArpIpv4TypeNeighborsTypeNeighborListStateType) GetOrigin() string {
	if m != nil && m.Origin != nil {
		return *m.Origin
	}
	return ""
}

func (m *ArpInformationMibArpIpv4TypeNeighborsTypeNeighborListStateType) GetHostName() string {
	if m != nil && m.HostName != nil {
		return *m.HostName
	}
	return ""
}

func (m *ArpInformationMibArpIpv4TypeNeighborsTypeNeighborListStateType) GetTableId() uint32 {
	if m != nil && m.TableId != nil {
		return *m.TableId
	}
	return 0
}

func (m *ArpInformationMibArpIpv4TypeNeighborsTypeNeighborListStateType) GetNeighborState() string {
	if m != nil && m.NeighborState != nil {
		return *m.NeighborState
	}
	return ""
}

func (m *ArpInformationMibArpIpv4TypeNeighborsTypeNeighborListStateType) GetExpiry() uint32 {
	if m != nil && m.Expiry != nil {
		return *m.Expiry
	}
	return 0
}

func (m *ArpInformationMibArpIpv4TypeNeighborsTypeNeighborListStateType) GetIsPublish() bool {
	if m != nil && m.IsPublish != nil {
		return *m.IsPublish
	}
	return false
}

func (m *ArpInformationMibArpIpv4TypeNeighborsTypeNeighborListStateType) GetInterfaceName() string {
	if m != nil && m.InterfaceName != nil {
		return *m.InterfaceName
	}
	return ""
}

func (m *ArpInformationMibArpIpv4TypeNeighborsTypeNeighborListStateType) GetLogicalRouterId() uint32 {
	if m != nil && m.LogicalRouterId != nil {
		return *m.LogicalRouterId
	}
	return 0
}

var E_JnprArpInformationMibArpExt = &proto.ExtensionDesc{
	ExtendedType:  (*JuniperNetworksSensors)(nil),
	ExtensionType: (*ArpInformationMibArp)(nil),
	Field:         53,
	Name:          "jnpr_arp_information_mib_arp_ext",
	Tag:           "bytes,53,opt,name=jnpr_arp_information_mib_arp_ext",
	Filename:      "mib2d_arp_oc.proto",
}

func init() {
	proto.RegisterType((*ArpInformationMibArp)(nil), "arp_information_mib_arp")
	proto.RegisterType((*ArpInformationMibArpIpv4Type)(nil), "arp_information_mib_arp.ipv4_type")
	proto.RegisterType((*ArpInformationMibArpIpv4TypeNeighborsType)(nil), "arp_information_mib_arp.ipv4_type.neighbors_type")
	proto.RegisterType((*ArpInformationMibArpIpv4TypeNeighborsTypeNeighborList)(nil), "arp_information_mib_arp.ipv4_type.neighbors_type.neighbor_list")
	proto.RegisterType((*ArpInformationMibArpIpv4TypeNeighborsTypeNeighborListStateType)(nil), "arp_information_mib_arp.ipv4_type.neighbors_type.neighbor_list.state_type")
	proto.RegisterExtension(E_JnprArpInformationMibArpExt)
}

func init() { proto.RegisterFile("mib2d_arp_oc.proto", fileDescriptor_6609990a88c6913e) }

var fileDescriptor_6609990a88c6913e = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xd1, 0x6a, 0xd4, 0x40,
	0x14, 0x86, 0xc9, 0xb6, 0xab, 0x9b, 0x53, 0x5a, 0x75, 0x44, 0x1b, 0x53, 0x84, 0x50, 0x10, 0x82,
	0x48, 0xc0, 0xda, 0x22, 0x08, 0xda, 0x6e, 0xc5, 0x8b, 0x2d, 0x5a, 0x4a, 0x7a, 0x2b, 0x0c, 0x93,
	0xcd, 0x74, 0xf7, 0xd8, 0x64, 0x66, 0x98, 0x99, 0xd5, 0xdd, 0x5b, 0x5f, 0xc2, 0x57, 0xf0, 0x51,
	0x04, 0x1f, 0xc0, 0xa7, 0x11, 0x64, 0x26, 0xbb, 0x09, 0xbd, 0x58, 0x44, 0xbc, 0x3c, 0xff, 0x7f,
	0xce, 0x77, 0xfe, 0xcc, 0x09, 0x90, 0x1a, 0x8b, 0x83, 0x92, 0x32, 0xad, 0xa8, 0x1c, 0x67, 0x4a,
	0x4b, 0x2b, 0xe3, 0xfb, 0x96, 0x57, 0xbc, 0xe6, 0x56, 0x2f, 0xa8, 0x95, 0xaa, 0x11, 0xf7, 0x7f,
	0xf7, 0x61, 0xd7, 0x75, 0xa1, 0xb8, 0x92, 0xba, 0x66, 0x16, 0xa5, 0xa0, 0x35, 0x16, 0x6e, 0x92,
	0xbc, 0x84, 0x4d, 0x54, 0x9f, 0x0f, 0xa3, 0x6f, 0x41, 0x12, 0xa4, 0x5b, 0x07, 0xfb, 0xd9, 0x9a,
	0xc6, 0xcc, 0x75, 0x51, 0xbb, 0x50, 0x3c, 0xf7, 0x03, 0xf1, 0xf7, 0x3e, 0x84, 0xad, 0x46, 0x2e,
	0x20, 0x14, 0x1c, 0x27, 0xd3, 0x42, 0x6a, 0xb3, 0x62, 0x3d, 0xff, 0x3b, 0x2b, 0x6b, 0x87, 0x1a,
	0x74, 0x07, 0x89, 0x7f, 0x6c, 0xc2, 0xce, 0x4d, 0x97, 0x7c, 0x84, 0xc1, 0x4a, 0x71, 0x3b, 0x36,
	0xd2, 0xad, 0x83, 0xe3, 0x7f, 0xde, 0xd1, 0x96, 0xb4, 0x42, 0x63, 0xf3, 0x96, 0x18, 0xff, 0xda,
	0x80, 0xed, 0x1b, 0x1e, 0x79, 0x00, 0x3d, 0x54, 0xd1, 0x8b, 0x24, 0x48, 0xc3, 0xd3, 0xfe, 0xd7,
	0x93, 0xde, 0x20, 0xc8, 0x7b, 0xa8, 0x08, 0x83, 0xbe, 0xb1, 0xcc, 0xf2, 0xd5, 0x77, 0x9e, 0xfd,
	0x67, 0x86, 0xcc, 0xd3, 0x9a, 0x07, 0x68, 0xc8, 0xf1, 0xcf, 0x1e, 0x40, 0xa7, 0x92, 0x9d, 0x2e,
	0x88, 0x4f, 0xf0, 0x0c, 0x48, 0x85, 0xe2, 0x9a, 0x56, 0x6c, 0xc1, 0x35, 0x65, 0x65, 0xa9, 0xb9,
	0x31, 0xd1, 0xa1, 0xf7, 0xef, 0x3a, 0xe7, 0xbd, 0x33, 0x86, 0x8d, 0x4e, 0x1e, 0xc2, 0x2d, 0xa9,
	0x71, 0x82, 0x22, 0x3a, 0xf2, 0x1d, 0xcb, 0x8a, 0xec, 0x41, 0x38, 0x95, 0xc6, 0x52, 0xc1, 0x6a,
	0x1e, 0xbd, 0xf6, 0xd6, 0xc0, 0x09, 0xe7, 0xac, 0xe6, 0xe4, 0x11, 0x0c, 0x2c, 0x2b, 0x2a, 0x4e,
	0xb1, 0x8c, 0xde, 0x24, 0x41, 0xba, 0x9d, 0xdf, 0xf6, 0xf5, 0xa8, 0x24, 0x4f, 0xba, 0xc3, 0xd0,
	0xe6, 0x21, 0x8e, 0xfd, 0x70, 0xfb, 0x7a, 0x97, 0x4e, 0x74, 0x6b, 0xf9, 0x5c, 0xa1, 0x5e, 0x44,
	0x27, 0x7e, 0x7e, 0x59, 0x91, 0xc7, 0x00, 0x68, 0xa8, 0x9a, 0x15, 0x15, 0x9a, 0x69, 0x34, 0x4c,
	0x82, 0x74, 0x90, 0x87, 0x68, 0x2e, 0x1a, 0xc1, 0xd1, 0x51, 0x58, 0xae, 0xaf, 0xd8, 0x98, 0x37,
	0xd1, 0x4e, 0x1b, 0x7a, 0xab, 0xfa, 0x7c, 0x4f, 0xe1, 0x5e, 0x25, 0x27, 0x38, 0x66, 0x15, 0xd5,
	0x72, 0x66, 0xb9, 0x76, 0x41, 0xdf, 0xfa, 0x45, 0x77, 0x96, 0x46, 0xee, 0xf5, 0x51, 0xf9, 0x6a,
	0x0e, 0xc9, 0x27, 0xa1, 0x34, 0x5d, 0x73, 0x26, 0xca, 0xe7, 0x96, 0xec, 0x66, 0x67, 0x33, 0x81,
	0x8a, 0xeb, 0x73, 0x6e, 0xbf, 0x48, 0x7d, 0x6d, 0x2e, 0xb9, 0x30, 0xee, 0x67, 0x3e, 0xf2, 0x37,
	0x8e, 0xd6, 0xdd, 0x38, 0xdf, 0x73, 0xe8, 0xa1, 0x56, 0xa3, 0xce, 0xfb, 0x80, 0xc5, 0x50, 0xab,
	0x77, 0x73, 0xfb, 0x27, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x74, 0x03, 0x2d, 0xa3, 0x03, 0x00, 0x00,
}
