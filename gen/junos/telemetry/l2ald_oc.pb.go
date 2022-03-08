// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: l2ald_oc.proto

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

type VlansL2Al struct {
	Vlan                 []*VlansL2AlVlanList `protobuf:"bytes,151,rep,name=vlan" json:"vlan,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *VlansL2Al) Reset()         { *m = VlansL2Al{} }
func (m *VlansL2Al) String() string { return proto.CompactTextString(m) }
func (*VlansL2Al) ProtoMessage()    {}
func (*VlansL2Al) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d20593bdd4ba4e6, []int{0}
}
func (m *VlansL2Al) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VlansL2Al.Unmarshal(m, b)
}
func (m *VlansL2Al) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VlansL2Al.Marshal(b, m, deterministic)
}
func (m *VlansL2Al) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VlansL2Al.Merge(m, src)
}
func (m *VlansL2Al) XXX_Size() int {
	return xxx_messageInfo_VlansL2Al.Size(m)
}
func (m *VlansL2Al) XXX_DiscardUnknown() {
	xxx_messageInfo_VlansL2Al.DiscardUnknown(m)
}

var xxx_messageInfo_VlansL2Al proto.InternalMessageInfo

func (m *VlansL2Al) GetVlan() []*VlansL2AlVlanList {
	if m != nil {
		return m.Vlan
	}
	return nil
}

type VlansL2AlVlanList struct {
	VlanId               *uint32                       `protobuf:"varint,51,opt,name=vlan_id,json=vlanId" json:"vlan_id,omitempty"`
	State                *VlansL2AlVlanListStateType   `protobuf:"bytes,151,opt,name=state" json:"state,omitempty"`
	Members              *VlansL2AlVlanListMembersType `protobuf:"bytes,152,opt,name=members" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *VlansL2AlVlanList) Reset()         { *m = VlansL2AlVlanList{} }
func (m *VlansL2AlVlanList) String() string { return proto.CompactTextString(m) }
func (*VlansL2AlVlanList) ProtoMessage()    {}
func (*VlansL2AlVlanList) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d20593bdd4ba4e6, []int{0, 0}
}
func (m *VlansL2AlVlanList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VlansL2AlVlanList.Unmarshal(m, b)
}
func (m *VlansL2AlVlanList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VlansL2AlVlanList.Marshal(b, m, deterministic)
}
func (m *VlansL2AlVlanList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VlansL2AlVlanList.Merge(m, src)
}
func (m *VlansL2AlVlanList) XXX_Size() int {
	return xxx_messageInfo_VlansL2AlVlanList.Size(m)
}
func (m *VlansL2AlVlanList) XXX_DiscardUnknown() {
	xxx_messageInfo_VlansL2AlVlanList.DiscardUnknown(m)
}

var xxx_messageInfo_VlansL2AlVlanList proto.InternalMessageInfo

func (m *VlansL2AlVlanList) GetVlanId() uint32 {
	if m != nil && m.VlanId != nil {
		return *m.VlanId
	}
	return 0
}

func (m *VlansL2AlVlanList) GetState() *VlansL2AlVlanListStateType {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *VlansL2AlVlanList) GetMembers() *VlansL2AlVlanListMembersType {
	if m != nil {
		return m.Members
	}
	return nil
}

type VlansL2AlVlanListStateType struct {
	VlanId               *uint32  `protobuf:"varint,52,opt,name=vlan_id,json=vlanId" json:"vlan_id,omitempty"`
	Name                 *string  `protobuf:"bytes,51,opt,name=name" json:"name,omitempty"`
	Status               *string  `protobuf:"bytes,53,opt,name=status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VlansL2AlVlanListStateType) Reset()         { *m = VlansL2AlVlanListStateType{} }
func (m *VlansL2AlVlanListStateType) String() string { return proto.CompactTextString(m) }
func (*VlansL2AlVlanListStateType) ProtoMessage()    {}
func (*VlansL2AlVlanListStateType) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d20593bdd4ba4e6, []int{0, 0, 0}
}
func (m *VlansL2AlVlanListStateType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VlansL2AlVlanListStateType.Unmarshal(m, b)
}
func (m *VlansL2AlVlanListStateType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VlansL2AlVlanListStateType.Marshal(b, m, deterministic)
}
func (m *VlansL2AlVlanListStateType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VlansL2AlVlanListStateType.Merge(m, src)
}
func (m *VlansL2AlVlanListStateType) XXX_Size() int {
	return xxx_messageInfo_VlansL2AlVlanListStateType.Size(m)
}
func (m *VlansL2AlVlanListStateType) XXX_DiscardUnknown() {
	xxx_messageInfo_VlansL2AlVlanListStateType.DiscardUnknown(m)
}

var xxx_messageInfo_VlansL2AlVlanListStateType proto.InternalMessageInfo

func (m *VlansL2AlVlanListStateType) GetVlanId() uint32 {
	if m != nil && m.VlanId != nil {
		return *m.VlanId
	}
	return 0
}

func (m *VlansL2AlVlanListStateType) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *VlansL2AlVlanListStateType) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ""
}

type VlansL2AlVlanListMembersType struct {
	Member               []*VlansL2AlVlanListMembersTypeMemberList `protobuf:"bytes,151,rep,name=member" json:"member,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *VlansL2AlVlanListMembersType) Reset()         { *m = VlansL2AlVlanListMembersType{} }
func (m *VlansL2AlVlanListMembersType) String() string { return proto.CompactTextString(m) }
func (*VlansL2AlVlanListMembersType) ProtoMessage()    {}
func (*VlansL2AlVlanListMembersType) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d20593bdd4ba4e6, []int{0, 0, 1}
}
func (m *VlansL2AlVlanListMembersType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VlansL2AlVlanListMembersType.Unmarshal(m, b)
}
func (m *VlansL2AlVlanListMembersType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VlansL2AlVlanListMembersType.Marshal(b, m, deterministic)
}
func (m *VlansL2AlVlanListMembersType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VlansL2AlVlanListMembersType.Merge(m, src)
}
func (m *VlansL2AlVlanListMembersType) XXX_Size() int {
	return xxx_messageInfo_VlansL2AlVlanListMembersType.Size(m)
}
func (m *VlansL2AlVlanListMembersType) XXX_DiscardUnknown() {
	xxx_messageInfo_VlansL2AlVlanListMembersType.DiscardUnknown(m)
}

var xxx_messageInfo_VlansL2AlVlanListMembersType proto.InternalMessageInfo

func (m *VlansL2AlVlanListMembersType) GetMember() []*VlansL2AlVlanListMembersTypeMemberList {
	if m != nil {
		return m.Member
	}
	return nil
}

type VlansL2AlVlanListMembersTypeMemberList struct {
	InterfaceRef         *VlansL2AlVlanListMembersTypeMemberListInterfaceRefType `protobuf:"bytes,151,opt,name=interface_ref,json=interfaceRef" json:"interface_ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                `json:"-"`
	XXX_unrecognized     []byte                                                  `json:"-"`
	XXX_sizecache        int32                                                   `json:"-"`
}

func (m *VlansL2AlVlanListMembersTypeMemberList) Reset() {
	*m = VlansL2AlVlanListMembersTypeMemberList{}
}
func (m *VlansL2AlVlanListMembersTypeMemberList) String() string { return proto.CompactTextString(m) }
func (*VlansL2AlVlanListMembersTypeMemberList) ProtoMessage()    {}
func (*VlansL2AlVlanListMembersTypeMemberList) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d20593bdd4ba4e6, []int{0, 0, 1, 0}
}
func (m *VlansL2AlVlanListMembersTypeMemberList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VlansL2AlVlanListMembersTypeMemberList.Unmarshal(m, b)
}
func (m *VlansL2AlVlanListMembersTypeMemberList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VlansL2AlVlanListMembersTypeMemberList.Marshal(b, m, deterministic)
}
func (m *VlansL2AlVlanListMembersTypeMemberList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VlansL2AlVlanListMembersTypeMemberList.Merge(m, src)
}
func (m *VlansL2AlVlanListMembersTypeMemberList) XXX_Size() int {
	return xxx_messageInfo_VlansL2AlVlanListMembersTypeMemberList.Size(m)
}
func (m *VlansL2AlVlanListMembersTypeMemberList) XXX_DiscardUnknown() {
	xxx_messageInfo_VlansL2AlVlanListMembersTypeMemberList.DiscardUnknown(m)
}

var xxx_messageInfo_VlansL2AlVlanListMembersTypeMemberList proto.InternalMessageInfo

func (m *VlansL2AlVlanListMembersTypeMemberList) GetInterfaceRef() *VlansL2AlVlanListMembersTypeMemberListInterfaceRefType {
	if m != nil {
		return m.InterfaceRef
	}
	return nil
}

type VlansL2AlVlanListMembersTypeMemberListInterfaceRefType struct {
	State                *VlansL2AlVlanListMembersTypeMemberListInterfaceRefTypeStateType `protobuf:"bytes,151,opt,name=state" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                         `json:"-"`
	XXX_unrecognized     []byte                                                           `json:"-"`
	XXX_sizecache        int32                                                            `json:"-"`
}

func (m *VlansL2AlVlanListMembersTypeMemberListInterfaceRefType) Reset() {
	*m = VlansL2AlVlanListMembersTypeMemberListInterfaceRefType{}
}
func (m *VlansL2AlVlanListMembersTypeMemberListInterfaceRefType) String() string {
	return proto.CompactTextString(m)
}
func (*VlansL2AlVlanListMembersTypeMemberListInterfaceRefType) ProtoMessage() {}
func (*VlansL2AlVlanListMembersTypeMemberListInterfaceRefType) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d20593bdd4ba4e6, []int{0, 0, 1, 0, 0}
}
func (m *VlansL2AlVlanListMembersTypeMemberListInterfaceRefType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VlansL2AlVlanListMembersTypeMemberListInterfaceRefType.Unmarshal(m, b)
}
func (m *VlansL2AlVlanListMembersTypeMemberListInterfaceRefType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VlansL2AlVlanListMembersTypeMemberListInterfaceRefType.Marshal(b, m, deterministic)
}
func (m *VlansL2AlVlanListMembersTypeMemberListInterfaceRefType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VlansL2AlVlanListMembersTypeMemberListInterfaceRefType.Merge(m, src)
}
func (m *VlansL2AlVlanListMembersTypeMemberListInterfaceRefType) XXX_Size() int {
	return xxx_messageInfo_VlansL2AlVlanListMembersTypeMemberListInterfaceRefType.Size(m)
}
func (m *VlansL2AlVlanListMembersTypeMemberListInterfaceRefType) XXX_DiscardUnknown() {
	xxx_messageInfo_VlansL2AlVlanListMembersTypeMemberListInterfaceRefType.DiscardUnknown(m)
}

var xxx_messageInfo_VlansL2AlVlanListMembersTypeMemberListInterfaceRefType proto.InternalMessageInfo

func (m *VlansL2AlVlanListMembersTypeMemberListInterfaceRefType) GetState() *VlansL2AlVlanListMembersTypeMemberListInterfaceRefTypeStateType {
	if m != nil {
		return m.State
	}
	return nil
}

type VlansL2AlVlanListMembersTypeMemberListInterfaceRefTypeStateType struct {
	Interface            *string  `protobuf:"bytes,51,opt,name=interface" json:"interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VlansL2AlVlanListMembersTypeMemberListInterfaceRefTypeStateType) Reset() {
	*m = VlansL2AlVlanListMembersTypeMemberListInterfaceRefTypeStateType{}
}
func (m *VlansL2AlVlanListMembersTypeMemberListInterfaceRefTypeStateType) String() string {
	return proto.CompactTextString(m)
}
func (*VlansL2AlVlanListMembersTypeMemberListInterfaceRefTypeStateType) ProtoMessage() {}
func (*VlansL2AlVlanListMembersTypeMemberListInterfaceRefTypeStateType) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d20593bdd4ba4e6, []int{0, 0, 1, 0, 0, 0}
}
func (m *VlansL2AlVlanListMembersTypeMemberListInterfaceRefTypeStateType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VlansL2AlVlanListMembersTypeMemberListInterfaceRefTypeStateType.Unmarshal(m, b)
}
func (m *VlansL2AlVlanListMembersTypeMemberListInterfaceRefTypeStateType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VlansL2AlVlanListMembersTypeMemberListInterfaceRefTypeStateType.Marshal(b, m, deterministic)
}
func (m *VlansL2AlVlanListMembersTypeMemberListInterfaceRefTypeStateType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VlansL2AlVlanListMembersTypeMemberListInterfaceRefTypeStateType.Merge(m, src)
}
func (m *VlansL2AlVlanListMembersTypeMemberListInterfaceRefTypeStateType) XXX_Size() int {
	return xxx_messageInfo_VlansL2AlVlanListMembersTypeMemberListInterfaceRefTypeStateType.Size(m)
}
func (m *VlansL2AlVlanListMembersTypeMemberListInterfaceRefTypeStateType) XXX_DiscardUnknown() {
	xxx_messageInfo_VlansL2AlVlanListMembersTypeMemberListInterfaceRefTypeStateType.DiscardUnknown(m)
}

var xxx_messageInfo_VlansL2AlVlanListMembersTypeMemberListInterfaceRefTypeStateType proto.InternalMessageInfo

func (m *VlansL2AlVlanListMembersTypeMemberListInterfaceRefTypeStateType) GetInterface() string {
	if m != nil && m.Interface != nil {
		return *m.Interface
	}
	return ""
}

var E_JnprVlansL2AlExt = &proto.ExtensionDesc{
	ExtendedType:  (*JuniperNetworksSensors)(nil),
	ExtensionType: (*VlansL2Al)(nil),
	Field:         50,
	Name:          "jnpr_vlans_l2al_ext",
	Tag:           "bytes,50,opt,name=jnpr_vlans_l2al_ext",
	Filename:      "l2ald_oc.proto",
}

func init() {
	proto.RegisterType((*VlansL2Al)(nil), "vlans_l2al")
	proto.RegisterType((*VlansL2AlVlanList)(nil), "vlans_l2al.vlan_list")
	proto.RegisterType((*VlansL2AlVlanListStateType)(nil), "vlans_l2al.vlan_list.state_type")
	proto.RegisterType((*VlansL2AlVlanListMembersType)(nil), "vlans_l2al.vlan_list.members_type")
	proto.RegisterType((*VlansL2AlVlanListMembersTypeMemberList)(nil), "vlans_l2al.vlan_list.members_type.member_list")
	proto.RegisterType((*VlansL2AlVlanListMembersTypeMemberListInterfaceRefType)(nil), "vlans_l2al.vlan_list.members_type.member_list.interface_ref_type")
	proto.RegisterType((*VlansL2AlVlanListMembersTypeMemberListInterfaceRefTypeStateType)(nil), "vlans_l2al.vlan_list.members_type.member_list.interface_ref_type.state_type")
	proto.RegisterExtension(E_JnprVlansL2AlExt)
}

func init() { proto.RegisterFile("l2ald_oc.proto", fileDescriptor_3d20593bdd4ba4e6) }

var fileDescriptor_3d20593bdd4ba4e6 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x51, 0x4f, 0x4b, 0xfb, 0x40,
	0x14, 0x24, 0xfd, 0xf5, 0xcf, 0xaf, 0x2f, 0xad, 0xc8, 0x16, 0x6d, 0x08, 0x22, 0xa5, 0xa7, 0xd2,
	0x43, 0x0e, 0x51, 0x11, 0x04, 0x41, 0x05, 0x11, 0xa5, 0x08, 0xa6, 0xe0, 0x75, 0x89, 0xed, 0xab,
	0x44, 0xf3, 0x8f, 0xdd, 0xad, 0xb6, 0x57, 0xbf, 0x84, 0x9e, 0xbc, 0x79, 0xf3, 0xcb, 0x79, 0xf3,
	0x28, 0xbb, 0x49, 0x93, 0xd6, 0x16, 0x44, 0xbc, 0xed, 0xcc, 0xce, 0xcc, 0x1b, 0xde, 0x83, 0x35,
	0xdf, 0x76, 0xfd, 0x21, 0x8d, 0x06, 0x56, 0xcc, 0x22, 0x11, 0x99, 0x0d, 0x81, 0x3e, 0x06, 0x28,
	0xd8, 0x94, 0x8a, 0x28, 0x4e, 0xc8, 0xf6, 0x7b, 0x09, 0xe0, 0xc1, 0x77, 0x43, 0x4e, 0xa5, 0x9a,
	0x74, 0xa1, 0x28, 0x91, 0xf1, 0xac, 0xb5, 0xfe, 0x75, 0x74, 0x7b, 0xc3, 0xca, 0xff, 0xd4, 0x93,
	0xfa, 0x1e, 0x17, 0x8e, 0xd2, 0x98, 0x1f, 0x45, 0xa8, 0x66, 0x1c, 0xd9, 0x86, 0x8a, 0x02, 0xde,
	0xd0, 0xd8, 0x69, 0x69, 0x9d, 0xfa, 0x49, 0xe9, 0xe9, 0xa8, 0xf0, 0x5f, 0x73, 0xca, 0x92, 0x3d,
	0x1f, 0x92, 0x7d, 0x28, 0x71, 0xe1, 0x0a, 0x94, 0xd1, 0x5a, 0x47, 0xb7, 0x5b, 0x2b, 0xa3, 0x2d,
	0xa5, 0xa1, 0x62, 0x1a, 0xa3, 0x93, 0xe8, 0xc9, 0x21, 0x54, 0x02, 0x0c, 0x6e, 0x90, 0x71, 0xe3,
	0x25, 0xb1, 0xb6, 0x57, 0x5b, 0x53, 0x55, 0x62, 0x9e, 0x79, 0xcc, 0x2b, 0x80, 0x3c, 0x93, 0x34,
	0xf3, 0x96, 0xbb, 0xb2, 0x65, 0x56, 0x8f, 0x40, 0x31, 0x74, 0x03, 0x54, 0xdd, 0xab, 0x8e, 0x7a,
	0x93, 0x4d, 0x28, 0x4b, 0xeb, 0x98, 0x1b, 0x7b, 0x8a, 0x4d, 0x91, 0xf9, 0x59, 0x80, 0xda, 0xfc,
	0x30, 0x72, 0x06, 0xe5, 0x04, 0xcf, 0xf6, 0x66, 0xfd, 0xdc, 0x30, 0x05, 0xc9, 0x42, 0x53, 0xbb,
	0xf9, 0x5a, 0x00, 0x7d, 0x8e, 0x27, 0xb7, 0x50, 0xf7, 0x42, 0x81, 0x6c, 0xe4, 0x0e, 0x90, 0x32,
	0x1c, 0xcd, 0x96, 0x77, 0xfc, 0xbb, 0x7c, 0x6b, 0x21, 0x24, 0x59, 0x50, 0x2d, 0xe3, 0x1c, 0x1c,
	0x99, 0x6f, 0x1a, 0x90, 0x65, 0x11, 0x19, 0x7c, 0x3b, 0x5a, 0xef, 0xcf, 0x73, 0x97, 0x0f, 0x6c,
	0x76, 0x17, 0x2e, 0xb4, 0x05, 0xd5, 0xcc, 0x95, 0x5e, 0x23, 0x27, 0x0e, 0xfa, 0xd0, 0xb8, 0x0b,
	0x63, 0x46, 0xf3, 0x1a, 0x14, 0x27, 0x82, 0x34, 0xad, 0x8b, 0x71, 0xe8, 0xc5, 0xc8, 0x2e, 0x51,
	0x3c, 0x46, 0xec, 0x9e, 0xf7, 0x31, 0xe4, 0x11, 0xe3, 0x86, 0xad, 0x6a, 0xeb, 0x73, 0xb5, 0x9d,
	0x75, 0x19, 0x70, 0x2d, 0x71, 0xcf, 0x76, 0xfd, 0xd3, 0x89, 0xf8, 0x0a, 0x00, 0x00, 0xff, 0xff,
	0xa9, 0xef, 0x30, 0x7a, 0x29, 0x03, 0x00, 0x00,
}
