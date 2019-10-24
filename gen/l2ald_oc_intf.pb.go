// Code generated by protoc-gen-go. DO NOT EDIT.
// source: l2ald_oc_intf.proto

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

type InterfacesL2AlInterface struct {
	Interface            []*InterfacesL2AlInterfaceInterfaceList `protobuf:"bytes,151,rep,name=interface" json:"interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *InterfacesL2AlInterface) Reset()         { *m = InterfacesL2AlInterface{} }
func (m *InterfacesL2AlInterface) String() string { return proto.CompactTextString(m) }
func (*InterfacesL2AlInterface) ProtoMessage()    {}
func (*InterfacesL2AlInterface) Descriptor() ([]byte, []int) {
	return fileDescriptor_l2ald_oc_intf_d39471c2da058846, []int{0}
}
func (m *InterfacesL2AlInterface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfacesL2AlInterface.Unmarshal(m, b)
}
func (m *InterfacesL2AlInterface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfacesL2AlInterface.Marshal(b, m, deterministic)
}
func (dst *InterfacesL2AlInterface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfacesL2AlInterface.Merge(dst, src)
}
func (m *InterfacesL2AlInterface) XXX_Size() int {
	return xxx_messageInfo_InterfacesL2AlInterface.Size(m)
}
func (m *InterfacesL2AlInterface) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfacesL2AlInterface.DiscardUnknown(m)
}

var xxx_messageInfo_InterfacesL2AlInterface proto.InternalMessageInfo

func (m *InterfacesL2AlInterface) GetInterface() []*InterfacesL2AlInterfaceInterfaceList {
	if m != nil {
		return m.Interface
	}
	return nil
}

type InterfacesL2AlInterfaceInterfaceList struct {
	Name                 *string                                                `protobuf:"bytes,51,opt,name=name" json:"name,omitempty"`
	Subinterfaces        *InterfacesL2AlInterfaceInterfaceListSubinterfacesType `protobuf:"bytes,151,opt,name=subinterfaces" json:"subinterfaces,omitempty"`
	Ethernet             *InterfacesL2AlInterfaceInterfaceListEthernetType      `protobuf:"bytes,152,opt,name=ethernet" json:"ethernet,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                               `json:"-"`
	XXX_unrecognized     []byte                                                 `json:"-"`
	XXX_sizecache        int32                                                  `json:"-"`
}

func (m *InterfacesL2AlInterfaceInterfaceList) Reset()         { *m = InterfacesL2AlInterfaceInterfaceList{} }
func (m *InterfacesL2AlInterfaceInterfaceList) String() string { return proto.CompactTextString(m) }
func (*InterfacesL2AlInterfaceInterfaceList) ProtoMessage()    {}
func (*InterfacesL2AlInterfaceInterfaceList) Descriptor() ([]byte, []int) {
	return fileDescriptor_l2ald_oc_intf_d39471c2da058846, []int{0, 0}
}
func (m *InterfacesL2AlInterfaceInterfaceList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfacesL2AlInterfaceInterfaceList.Unmarshal(m, b)
}
func (m *InterfacesL2AlInterfaceInterfaceList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfacesL2AlInterfaceInterfaceList.Marshal(b, m, deterministic)
}
func (dst *InterfacesL2AlInterfaceInterfaceList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfacesL2AlInterfaceInterfaceList.Merge(dst, src)
}
func (m *InterfacesL2AlInterfaceInterfaceList) XXX_Size() int {
	return xxx_messageInfo_InterfacesL2AlInterfaceInterfaceList.Size(m)
}
func (m *InterfacesL2AlInterfaceInterfaceList) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfacesL2AlInterfaceInterfaceList.DiscardUnknown(m)
}

var xxx_messageInfo_InterfacesL2AlInterfaceInterfaceList proto.InternalMessageInfo

func (m *InterfacesL2AlInterfaceInterfaceList) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *InterfacesL2AlInterfaceInterfaceList) GetSubinterfaces() *InterfacesL2AlInterfaceInterfaceListSubinterfacesType {
	if m != nil {
		return m.Subinterfaces
	}
	return nil
}

func (m *InterfacesL2AlInterfaceInterfaceList) GetEthernet() *InterfacesL2AlInterfaceInterfaceListEthernetType {
	if m != nil {
		return m.Ethernet
	}
	return nil
}

type InterfacesL2AlInterfaceInterfaceListSubinterfacesType struct {
	Subinterface         []*InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList `protobuf:"bytes,151,rep,name=subinterface" json:"subinterface,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                 `json:"-"`
	XXX_unrecognized     []byte                                                                   `json:"-"`
	XXX_sizecache        int32                                                                    `json:"-"`
}

func (m *InterfacesL2AlInterfaceInterfaceListSubinterfacesType) Reset() {
	*m = InterfacesL2AlInterfaceInterfaceListSubinterfacesType{}
}
func (m *InterfacesL2AlInterfaceInterfaceListSubinterfacesType) String() string {
	return proto.CompactTextString(m)
}
func (*InterfacesL2AlInterfaceInterfaceListSubinterfacesType) ProtoMessage() {}
func (*InterfacesL2AlInterfaceInterfaceListSubinterfacesType) Descriptor() ([]byte, []int) {
	return fileDescriptor_l2ald_oc_intf_d39471c2da058846, []int{0, 0, 0}
}
func (m *InterfacesL2AlInterfaceInterfaceListSubinterfacesType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListSubinterfacesType.Unmarshal(m, b)
}
func (m *InterfacesL2AlInterfaceInterfaceListSubinterfacesType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListSubinterfacesType.Marshal(b, m, deterministic)
}
func (dst *InterfacesL2AlInterfaceInterfaceListSubinterfacesType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListSubinterfacesType.Merge(dst, src)
}
func (m *InterfacesL2AlInterfaceInterfaceListSubinterfacesType) XXX_Size() int {
	return xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListSubinterfacesType.Size(m)
}
func (m *InterfacesL2AlInterfaceInterfaceListSubinterfacesType) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListSubinterfacesType.DiscardUnknown(m)
}

var xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListSubinterfacesType proto.InternalMessageInfo

func (m *InterfacesL2AlInterfaceInterfaceListSubinterfacesType) GetSubinterface() []*InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList {
	if m != nil {
		return m.Subinterface
	}
	return nil
}

type InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList struct {
	Index                *uint32                                                                        `protobuf:"varint,51,opt,name=index" json:"index,omitempty"`
	Vlan                 *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType `protobuf:"bytes,151,opt,name=vlan" json:"vlan,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                       `json:"-"`
	XXX_unrecognized     []byte                                                                         `json:"-"`
	XXX_sizecache        int32                                                                          `json:"-"`
}

func (m *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList) Reset() {
	*m = InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList{}
}
func (m *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList) String() string {
	return proto.CompactTextString(m)
}
func (*InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList) ProtoMessage() {}
func (*InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList) Descriptor() ([]byte, []int) {
	return fileDescriptor_l2ald_oc_intf_d39471c2da058846, []int{0, 0, 0, 0}
}
func (m *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList.Unmarshal(m, b)
}
func (m *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList.Marshal(b, m, deterministic)
}
func (dst *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList.Merge(dst, src)
}
func (m *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList) XXX_Size() int {
	return xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList.Size(m)
}
func (m *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList.DiscardUnknown(m)
}

var xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList proto.InternalMessageInfo

func (m *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList) GetIndex() uint32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList) GetVlan() *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType {
	if m != nil {
		return m.Vlan
	}
	return nil
}

type InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType struct {
	State                *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType `protobuf:"bytes,152,opt,name=state" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                                `json:"-"`
	XXX_unrecognized     []byte                                                                                  `json:"-"`
	XXX_sizecache        int32                                                                                   `json:"-"`
}

func (m *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType) Reset() {
	*m = InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType{}
}
func (m *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType) String() string {
	return proto.CompactTextString(m)
}
func (*InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType) ProtoMessage() {}
func (*InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType) Descriptor() ([]byte, []int) {
	return fileDescriptor_l2ald_oc_intf_d39471c2da058846, []int{0, 0, 0, 0, 0}
}
func (m *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType.Unmarshal(m, b)
}
func (m *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType.Marshal(b, m, deterministic)
}
func (dst *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType.Merge(dst, src)
}
func (m *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType) XXX_Size() int {
	return xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType.Size(m)
}
func (m *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType.DiscardUnknown(m)
}

var xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType proto.InternalMessageInfo

func (m *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType) GetState() *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType {
	if m != nil {
		return m.State
	}
	return nil
}

type InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType struct {
	VlanId               *uint32  `protobuf:"varint,51,opt,name=vlan_id,json=vlanId" json:"vlan_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType) Reset() {
	*m = InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType{}
}
func (m *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType) String() string {
	return proto.CompactTextString(m)
}
func (*InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType) ProtoMessage() {
}
func (*InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType) Descriptor() ([]byte, []int) {
	return fileDescriptor_l2ald_oc_intf_d39471c2da058846, []int{0, 0, 0, 0, 0, 0}
}
func (m *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType.Unmarshal(m, b)
}
func (m *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType.Marshal(b, m, deterministic)
}
func (dst *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType.Merge(dst, src)
}
func (m *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType) XXX_Size() int {
	return xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType.Size(m)
}
func (m *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType.DiscardUnknown(m)
}

var xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType proto.InternalMessageInfo

func (m *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType) GetVlanId() uint32 {
	if m != nil && m.VlanId != nil {
		return *m.VlanId
	}
	return 0
}

type InterfacesL2AlInterfaceInterfaceListEthernetType struct {
	SwitchedVlan         *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType `protobuf:"bytes,151,opt,name=switched_vlan,json=switchedVlan" json:"switched_vlan,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                          `json:"-"`
	XXX_unrecognized     []byte                                                            `json:"-"`
	XXX_sizecache        int32                                                             `json:"-"`
}

func (m *InterfacesL2AlInterfaceInterfaceListEthernetType) Reset() {
	*m = InterfacesL2AlInterfaceInterfaceListEthernetType{}
}
func (m *InterfacesL2AlInterfaceInterfaceListEthernetType) String() string {
	return proto.CompactTextString(m)
}
func (*InterfacesL2AlInterfaceInterfaceListEthernetType) ProtoMessage() {}
func (*InterfacesL2AlInterfaceInterfaceListEthernetType) Descriptor() ([]byte, []int) {
	return fileDescriptor_l2ald_oc_intf_d39471c2da058846, []int{0, 0, 1}
}
func (m *InterfacesL2AlInterfaceInterfaceListEthernetType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListEthernetType.Unmarshal(m, b)
}
func (m *InterfacesL2AlInterfaceInterfaceListEthernetType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListEthernetType.Marshal(b, m, deterministic)
}
func (dst *InterfacesL2AlInterfaceInterfaceListEthernetType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListEthernetType.Merge(dst, src)
}
func (m *InterfacesL2AlInterfaceInterfaceListEthernetType) XXX_Size() int {
	return xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListEthernetType.Size(m)
}
func (m *InterfacesL2AlInterfaceInterfaceListEthernetType) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListEthernetType.DiscardUnknown(m)
}

var xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListEthernetType proto.InternalMessageInfo

func (m *InterfacesL2AlInterfaceInterfaceListEthernetType) GetSwitchedVlan() *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType {
	if m != nil {
		return m.SwitchedVlan
	}
	return nil
}

type InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType struct {
	State                *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType `protobuf:"bytes,151,opt,name=state" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                   `json:"-"`
	XXX_unrecognized     []byte                                                                     `json:"-"`
	XXX_sizecache        int32                                                                      `json:"-"`
}

func (m *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType) Reset() {
	*m = InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType{}
}
func (m *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType) String() string {
	return proto.CompactTextString(m)
}
func (*InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType) ProtoMessage() {}
func (*InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType) Descriptor() ([]byte, []int) {
	return fileDescriptor_l2ald_oc_intf_d39471c2da058846, []int{0, 0, 1, 0}
}
func (m *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType.Unmarshal(m, b)
}
func (m *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType.Marshal(b, m, deterministic)
}
func (dst *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType.Merge(dst, src)
}
func (m *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType) XXX_Size() int {
	return xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType.Size(m)
}
func (m *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType.DiscardUnknown(m)
}

var xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType proto.InternalMessageInfo

func (m *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType) GetState() *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType {
	if m != nil {
		return m.State
	}
	return nil
}

type InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType struct {
	InterfaceMode        *string  `protobuf:"bytes,51,opt,name=interface_mode,json=interfaceMode" json:"interface_mode,omitempty"`
	NativeVlan           *uint32  `protobuf:"varint,52,opt,name=native_vlan,json=nativeVlan" json:"native_vlan,omitempty"`
	AccessVlan           *uint32  `protobuf:"varint,53,opt,name=access_vlan,json=accessVlan" json:"access_vlan,omitempty"`
	TrunkVlans           []string `protobuf:"bytes,54,rep,name=trunk_vlans,json=trunkVlans" json:"trunk_vlans,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType) Reset() {
	*m = InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType{}
}
func (m *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType) String() string {
	return proto.CompactTextString(m)
}
func (*InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType) ProtoMessage() {}
func (*InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType) Descriptor() ([]byte, []int) {
	return fileDescriptor_l2ald_oc_intf_d39471c2da058846, []int{0, 0, 1, 0, 0}
}
func (m *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType.Unmarshal(m, b)
}
func (m *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType.Marshal(b, m, deterministic)
}
func (dst *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType.Merge(dst, src)
}
func (m *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType) XXX_Size() int {
	return xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType.Size(m)
}
func (m *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType.DiscardUnknown(m)
}

var xxx_messageInfo_InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType proto.InternalMessageInfo

func (m *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType) GetInterfaceMode() string {
	if m != nil && m.InterfaceMode != nil {
		return *m.InterfaceMode
	}
	return ""
}

func (m *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType) GetNativeVlan() uint32 {
	if m != nil && m.NativeVlan != nil {
		return *m.NativeVlan
	}
	return 0
}

func (m *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType) GetAccessVlan() uint32 {
	if m != nil && m.AccessVlan != nil {
		return *m.AccessVlan
	}
	return 0
}

func (m *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType) GetTrunkVlans() []string {
	if m != nil {
		return m.TrunkVlans
	}
	return nil
}

var E_JnprInterfacesL2AlInterfaceExt = &proto.ExtensionDesc{
	ExtendedType:  (*JuniperNetworksSensors)(nil),
	ExtensionType: (*InterfacesL2AlInterface)(nil),
	Field:         49,
	Name:          "jnpr_interfaces_l2al_interface_ext",
	Tag:           "bytes,49,opt,name=jnpr_interfaces_l2al_interface_ext,json=jnprInterfacesL2alInterfaceExt",
	Filename:      "l2ald_oc_intf.proto",
}

func init() {
	proto.RegisterType((*InterfacesL2AlInterface)(nil), "interfaces_l2al_interface")
	proto.RegisterType((*InterfacesL2AlInterfaceInterfaceList)(nil), "interfaces_l2al_interface.interface_list")
	proto.RegisterType((*InterfacesL2AlInterfaceInterfaceListSubinterfacesType)(nil), "interfaces_l2al_interface.interface_list.subinterfaces_type")
	proto.RegisterType((*InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList)(nil), "interfaces_l2al_interface.interface_list.subinterfaces_type.subinterface_list")
	proto.RegisterType((*InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType)(nil), "interfaces_l2al_interface.interface_list.subinterfaces_type.subinterface_list.vlan_type")
	proto.RegisterType((*InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType)(nil), "interfaces_l2al_interface.interface_list.subinterfaces_type.subinterface_list.vlan_type.state_type")
	proto.RegisterType((*InterfacesL2AlInterfaceInterfaceListEthernetType)(nil), "interfaces_l2al_interface.interface_list.ethernet_type")
	proto.RegisterType((*InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType)(nil), "interfaces_l2al_interface.interface_list.ethernet_type.switched_vlan_type")
	proto.RegisterType((*InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType)(nil), "interfaces_l2al_interface.interface_list.ethernet_type.switched_vlan_type.state_type")
	proto.RegisterExtension(E_JnprInterfacesL2AlInterfaceExt)
}

func init() { proto.RegisterFile("l2ald_oc_intf.proto", fileDescriptor_l2ald_oc_intf_d39471c2da058846) }

var fileDescriptor_l2ald_oc_intf_d39471c2da058846 = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xdf, 0x6a, 0xd4, 0x40,
	0x14, 0xc6, 0xc9, 0x6e, 0x57, 0x37, 0x67, 0x1b, 0xc1, 0xe9, 0x45, 0xd3, 0x08, 0xba, 0x14, 0x0a,
	0xeb, 0x4d, 0xc0, 0xf5, 0x1f, 0x88, 0x17, 0x22, 0x08, 0xb6, 0x68, 0x2f, 0xe2, 0x1f, 0xbc, 0x1b,
	0x62, 0x72, 0x4a, 0x63, 0xb3, 0x93, 0x30, 0x73, 0xb6, 0xdd, 0x7a, 0x23, 0xf8, 0x10, 0x2a, 0x78,
	0xe1, 0x1b, 0xf8, 0x2a, 0xfa, 0x00, 0xde, 0xfb, 0x1a, 0x32, 0x33, 0x4d, 0xb2, 0x61, 0x5d, 0x28,
	0xba, 0xbd, 0x9c, 0xef, 0x9c, 0xf3, 0x7d, 0x67, 0x7e, 0xc9, 0xc0, 0x46, 0x3e, 0x8e, 0xf3, 0x94,
	0x17, 0x09, 0xcf, 0x04, 0x1d, 0x84, 0xa5, 0x2c, 0xa8, 0x08, 0x36, 0x08, 0x73, 0x9c, 0x20, 0xc9,
	0x53, 0x4e, 0x45, 0x69, 0xc5, 0xed, 0x1f, 0x2e, 0x6c, 0x65, 0x82, 0x50, 0x1e, 0xc4, 0x09, 0x2a,
	0xae, 0xe7, 0x78, 0x7d, 0x66, 0x4f, 0xc1, 0xad, 0x0f, 0xfe, 0x67, 0x67, 0xd8, 0x1d, 0x0d, 0xc6,
	0x37, 0xc3, 0xa5, 0xfd, 0x4d, 0x85, 0xe7, 0x99, 0xa2, 0xa8, 0x19, 0x0e, 0x7e, 0xf7, 0xe1, 0x4a,
	0xbb, 0xca, 0xb6, 0x60, 0x4d, 0xc4, 0x13, 0xf4, 0x6f, 0x0f, 0x9d, 0x91, 0xfb, 0xb8, 0xf7, 0xf1,
	0x51, 0xa7, 0xef, 0x44, 0x46, 0x62, 0x09, 0x78, 0x6a, 0xfa, 0xb6, 0xc9, 0xd1, 0xd9, 0xce, 0x68,
	0x30, 0x7e, 0x78, 0xee, 0xec, 0xb0, 0x35, 0xcf, 0xe9, 0xb4, 0xc4, 0xa8, 0xed, 0xc9, 0x5e, 0x42,
	0x1f, 0xe9, 0x10, 0xa5, 0x40, 0xf2, 0xbf, 0x58, 0xff, 0xfb, 0xe7, 0xf7, 0xaf, 0x46, 0xad, 0x75,
	0xed, 0x14, 0xfc, 0xea, 0x02, 0x5b, 0xcc, 0x66, 0x0a, 0xd6, 0xe7, 0xd5, 0x0a, 0xe6, 0xfe, 0xff,
	0x5c, 0xa8, 0x25, 0x59, 0xe2, 0xad, 0x90, 0xe0, 0x67, 0x07, 0xae, 0x2e, 0xf4, 0xb0, 0x6b, 0xd0,
	0xcb, 0x44, 0x8a, 0x33, 0x03, 0xde, 0xab, 0xc0, 0x5b, 0x8d, 0x4d, 0x60, 0xed, 0x38, 0x8f, 0x45,
	0x05, 0xfc, 0xcd, 0x6a, 0xf7, 0x0b, 0xb5, 0xb7, 0x25, 0x66, 0x62, 0x82, 0xef, 0x0e, 0xb8, 0xb5,
	0xc6, 0xde, 0x43, 0x4f, 0x51, 0x4c, 0x58, 0x7d, 0x8e, 0xe4, 0xa2, 0xd2, 0x43, 0x13, 0x63, 0x17,
	0xb1, 0x91, 0xc1, 0x0e, 0x40, 0x23, 0xb2, 0x4d, 0xb8, 0x6c, 0x9a, 0xb3, 0xd4, 0x52, 0x8a, 0x2e,
	0xe9, 0xe3, 0x6e, 0x1a, 0x7c, 0xeb, 0x82, 0xd7, 0xfa, 0xf4, 0xac, 0x04, 0x4f, 0x9d, 0x64, 0x94,
	0x1c, 0x62, 0xca, 0xe7, 0xd1, 0xed, 0xfd, 0xe3, 0xbf, 0x14, 0xb6, 0xdc, 0xec, 0x8e, 0xeb, 0x95,
	0xf6, 0x5a, 0x43, 0xfb, 0xda, 0x01, 0xb6, 0xd8, 0xc4, 0xf2, 0x8a, 0xde, 0xd9, 0x02, 0xaf, 0x56,
	0xb7, 0xc0, 0x5f, 0x78, 0x7d, 0x72, 0x5a, 0xc0, 0x76, 0xe6, 0x9f, 0xf7, 0xa4, 0x48, 0xcf, 0x9e,
	0x75, 0xe4, 0xd5, 0xea, 0xf3, 0x22, 0x45, 0x76, 0x03, 0x06, 0x22, 0xa6, 0xec, 0x18, 0x2d, 0xaa,
	0x3b, 0x86, 0x2d, 0x58, 0x49, 0xdf, 0x4d, 0x37, 0xc4, 0x49, 0x82, 0x4a, 0xd9, 0x86, 0xbb, 0xb6,
	0xc1, 0x4a, 0x55, 0x03, 0xc9, 0xa9, 0x38, 0x32, 0x75, 0xe5, 0xdf, 0x1b, 0x76, 0x47, 0x6e, 0x04,
	0x46, 0xd2, 0x75, 0xf5, 0xe0, 0x03, 0x6c, 0xbf, 0x13, 0xa5, 0xe4, 0x4b, 0x2f, 0xcf, 0x71, 0x46,
	0x6c, 0x33, 0xdc, 0x9b, 0x8a, 0xac, 0x44, 0xb9, 0x8f, 0x74, 0x52, 0xc8, 0x23, 0xf5, 0x02, 0x85,
	0x2a, 0xa4, 0xf2, 0x6f, 0x19, 0x76, 0xc1, 0x72, 0x76, 0xd1, 0x75, 0x6d, 0xbf, 0x5b, 0x97, 0x9f,
	0x8d, 0xe3, 0xbc, 0x3e, 0x3d, 0x99, 0xd1, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x19, 0x9b, 0x69,
	0x7d, 0x7d, 0x05, 0x00, 0x00,
}
