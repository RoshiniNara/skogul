// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pfe_mpls_ldp_p2mp_lsp_branch_oc.proto

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

type MplsPfeMplsLdpP2MpLspBranch struct {
	SignalingProtocols   *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsType `protobuf:"bytes,151,opt,name=signaling_protocols,json=signalingProtocols" json:"signaling_protocols,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                           `json:"-"`
	XXX_unrecognized     []byte                                             `json:"-"`
	XXX_sizecache        int32                                              `json:"-"`
}

func (m *MplsPfeMplsLdpP2MpLspBranch) Reset()         { *m = MplsPfeMplsLdpP2MpLspBranch{} }
func (m *MplsPfeMplsLdpP2MpLspBranch) String() string { return proto.CompactTextString(m) }
func (*MplsPfeMplsLdpP2MpLspBranch) ProtoMessage()    {}
func (*MplsPfeMplsLdpP2MpLspBranch) Descriptor() ([]byte, []int) {
	return fileDescriptor_pfe_mpls_ldp_p2mp_lsp_branch_oc_b7a4f771e13a68b0, []int{0}
}
func (m *MplsPfeMplsLdpP2MpLspBranch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranch.Unmarshal(m, b)
}
func (m *MplsPfeMplsLdpP2MpLspBranch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranch.Marshal(b, m, deterministic)
}
func (dst *MplsPfeMplsLdpP2MpLspBranch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranch.Merge(dst, src)
}
func (m *MplsPfeMplsLdpP2MpLspBranch) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranch.Size(m)
}
func (m *MplsPfeMplsLdpP2MpLspBranch) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranch.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranch proto.InternalMessageInfo

func (m *MplsPfeMplsLdpP2MpLspBranch) GetSignalingProtocols() *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsType {
	if m != nil {
		return m.SignalingProtocols
	}
	return nil
}

type MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsType struct {
	Ldp                  *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpType `protobuf:"bytes,151,opt,name=ldp" json:"ldp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                  `json:"-"`
	XXX_unrecognized     []byte                                                    `json:"-"`
	XXX_sizecache        int32                                                     `json:"-"`
}

func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsType) Reset() {
	*m = MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsType{}
}
func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsType) String() string {
	return proto.CompactTextString(m)
}
func (*MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsType) ProtoMessage() {}
func (*MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsType) Descriptor() ([]byte, []int) {
	return fileDescriptor_pfe_mpls_ldp_p2mp_lsp_branch_oc_b7a4f771e13a68b0, []int{0, 0}
}
func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsType.Unmarshal(m, b)
}
func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsType.Marshal(b, m, deterministic)
}
func (dst *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsType.Merge(dst, src)
}
func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsType) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsType.Size(m)
}
func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsType) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsType.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsType proto.InternalMessageInfo

func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsType) GetLdp() *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpType {
	if m != nil {
		return m.Ldp
	}
	return nil
}

type MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpType struct {
	P2MpLspBranches      *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesType `protobuf:"bytes,151,opt,name=p2mp_lsp_branches,json=p2mpLspBranches" json:"p2mp_lsp_branches,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                     `json:"-"`
	XXX_unrecognized     []byte                                                                       `json:"-"`
	XXX_sizecache        int32                                                                        `json:"-"`
}

func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpType) Reset() {
	*m = MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpType{}
}
func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpType) String() string {
	return proto.CompactTextString(m)
}
func (*MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpType) ProtoMessage() {}
func (*MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpType) Descriptor() ([]byte, []int) {
	return fileDescriptor_pfe_mpls_ldp_p2mp_lsp_branch_oc_b7a4f771e13a68b0, []int{0, 0, 0}
}
func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpType.Unmarshal(m, b)
}
func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpType.Marshal(b, m, deterministic)
}
func (dst *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpType.Merge(dst, src)
}
func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpType) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpType.Size(m)
}
func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpType) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpType.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpType proto.InternalMessageInfo

func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpType) GetP2MpLspBranches() *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesType {
	if m != nil {
		return m.P2MpLspBranches
	}
	return nil
}

type MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesType struct {
	P2MpLspBranch        []*MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchList `protobuf:"bytes,151,rep,name=p2mp_lsp_branch,json=p2mpLspBranch" json:"p2mp_lsp_branch,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                                        `json:"-"`
	XXX_unrecognized     []byte                                                                                          `json:"-"`
	XXX_sizecache        int32                                                                                           `json:"-"`
}

func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesType) Reset() {
	*m = MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesType{}
}
func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesType) String() string {
	return proto.CompactTextString(m)
}
func (*MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesType) ProtoMessage() {}
func (*MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesType) Descriptor() ([]byte, []int) {
	return fileDescriptor_pfe_mpls_ldp_p2mp_lsp_branch_oc_b7a4f771e13a68b0, []int{0, 0, 0, 0}
}
func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesType.Unmarshal(m, b)
}
func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesType.Marshal(b, m, deterministic)
}
func (dst *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesType.Merge(dst, src)
}
func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesType) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesType.Size(m)
}
func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesType) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesType.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesType proto.InternalMessageInfo

func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesType) GetP2MpLspBranch() []*MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchList {
	if m != nil {
		return m.P2MpLspBranch
	}
	return nil
}

type MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchList struct {
	State                *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateType `protobuf:"bytes,151,opt,name=state" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                                               `json:"-"`
	XXX_unrecognized     []byte                                                                                                 `json:"-"`
	XXX_sizecache        int32                                                                                                  `json:"-"`
}

func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchList) Reset() {
	*m = MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchList{}
}
func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchList) String() string {
	return proto.CompactTextString(m)
}
func (*MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchList) ProtoMessage() {
}
func (*MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchList) Descriptor() ([]byte, []int) {
	return fileDescriptor_pfe_mpls_ldp_p2mp_lsp_branch_oc_b7a4f771e13a68b0, []int{0, 0, 0, 0, 0}
}
func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchList.Unmarshal(m, b)
}
func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchList.Marshal(b, m, deterministic)
}
func (dst *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchList.Merge(dst, src)
}
func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchList) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchList.Size(m)
}
func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchList) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchList.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchList proto.InternalMessageInfo

func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchList) GetState() *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateType {
	if m != nil {
		return m.State
	}
	return nil
}

type MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateType struct {
	Counters             []*MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateTypeCountersList `protobuf:"bytes,151,rep,name=counters" json:"counters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                                                             `json:"-"`
	XXX_unrecognized     []byte                                                                                                               `json:"-"`
	XXX_sizecache        int32                                                                                                                `json:"-"`
}

func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateType) Reset() {
	*m = MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateType{}
}
func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateType) String() string {
	return proto.CompactTextString(m)
}
func (*MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateType) ProtoMessage() {
}
func (*MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateType) Descriptor() ([]byte, []int) {
	return fileDescriptor_pfe_mpls_ldp_p2mp_lsp_branch_oc_b7a4f771e13a68b0, []int{0, 0, 0, 0, 0, 0}
}
func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateType.Unmarshal(m, b)
}
func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateType.Marshal(b, m, deterministic)
}
func (dst *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateType.Merge(dst, src)
}
func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateType) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateType.Size(m)
}
func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateType) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateType.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateType proto.InternalMessageInfo

func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateType) GetCounters() []*MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateTypeCountersList {
	if m != nil {
		return m.Counters
	}
	return nil
}

type MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateTypeCountersList struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateTypeCountersList) Reset() {
	*m = MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateTypeCountersList{}
}
func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateTypeCountersList) String() string {
	return proto.CompactTextString(m)
}
func (*MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateTypeCountersList) ProtoMessage() {
}
func (*MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateTypeCountersList) Descriptor() ([]byte, []int) {
	return fileDescriptor_pfe_mpls_ldp_p2mp_lsp_branch_oc_b7a4f771e13a68b0, []int{0, 0, 0, 0, 0, 0, 0}
}
func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateTypeCountersList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateTypeCountersList.Unmarshal(m, b)
}
func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateTypeCountersList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateTypeCountersList.Marshal(b, m, deterministic)
}
func (dst *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateTypeCountersList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateTypeCountersList.Merge(dst, src)
}
func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateTypeCountersList) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateTypeCountersList.Size(m)
}
func (m *MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateTypeCountersList) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateTypeCountersList.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateTypeCountersList proto.InternalMessageInfo

var E_JnprMplsPfeMplsLdpP2MpLspBranchExt = &proto.ExtensionDesc{
	ExtendedType:  (*JuniperNetworksSensors)(nil),
	ExtensionType: (*MplsPfeMplsLdpP2MpLspBranch)(nil),
	Field:         90,
	Name:          "jnpr_mpls_pfe_mpls_ldp_p2mp_lsp_branch_ext",
	Tag:           "bytes,90,opt,name=jnpr_mpls_pfe_mpls_ldp_p2mp_lsp_branch_ext,json=jnprMplsPfeMplsLdpP2mpLspBranchExt",
	Filename:      "pfe_mpls_ldp_p2mp_lsp_branch_oc.proto",
}

func init() {
	proto.RegisterType((*MplsPfeMplsLdpP2MpLspBranch)(nil), "mpls_pfe_mpls_ldp_p2mp_lsp_branch")
	proto.RegisterType((*MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsType)(nil), "mpls_pfe_mpls_ldp_p2mp_lsp_branch.signaling_protocols_type")
	proto.RegisterType((*MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpType)(nil), "mpls_pfe_mpls_ldp_p2mp_lsp_branch.signaling_protocols_type.ldp_type")
	proto.RegisterType((*MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesType)(nil), "mpls_pfe_mpls_ldp_p2mp_lsp_branch.signaling_protocols_type.ldp_type.p2mp_lsp_branches_type")
	proto.RegisterType((*MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchList)(nil), "mpls_pfe_mpls_ldp_p2mp_lsp_branch.signaling_protocols_type.ldp_type.p2mp_lsp_branches_type.p2mp_lsp_branch_list")
	proto.RegisterType((*MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateType)(nil), "mpls_pfe_mpls_ldp_p2mp_lsp_branch.signaling_protocols_type.ldp_type.p2mp_lsp_branches_type.p2mp_lsp_branch_list.state_type")
	proto.RegisterType((*MplsPfeMplsLdpP2MpLspBranchSignalingProtocolsTypeLdpTypeP2MpLspBranchesTypeP2MpLspBranchListStateTypeCountersList)(nil), "mpls_pfe_mpls_ldp_p2mp_lsp_branch.signaling_protocols_type.ldp_type.p2mp_lsp_branches_type.p2mp_lsp_branch_list.state_type.counters_list")
	proto.RegisterExtension(E_JnprMplsPfeMplsLdpP2MpLspBranchExt)
}

func init() {
	proto.RegisterFile("pfe_mpls_ldp_p2mp_lsp_branch_oc.proto", fileDescriptor_pfe_mpls_ldp_p2mp_lsp_branch_oc_b7a4f771e13a68b0)
}

var fileDescriptor_pfe_mpls_ldp_p2mp_lsp_branch_oc_b7a4f771e13a68b0 = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0xbf, 0x4a, 0x2b, 0x41,
	0x14, 0xc6, 0x99, 0x9b, 0x9b, 0x7b, 0xc3, 0x09, 0x21, 0xdc, 0xc9, 0x45, 0xc3, 0x56, 0x1a, 0x10,
	0x82, 0xc5, 0x14, 0x29, 0xb5, 0x13, 0x6d, 0x24, 0x4a, 0x88, 0x85, 0xa0, 0xc5, 0x10, 0x37, 0x93,
	0xb8, 0x3a, 0x3b, 0x73, 0x98, 0x99, 0x60, 0x82, 0x2f, 0xa0, 0xad, 0x16, 0xe2, 0x13, 0xd8, 0xf8,
	0x04, 0xd6, 0x82, 0xa5, 0x4f, 0x63, 0x2f, 0xbb, 0xeb, 0x46, 0x13, 0xa3, 0x29, 0xfc, 0x57, 0x2d,
	0x7c, 0xe7, 0x9c, 0xef, 0xfb, 0x9d, 0xb3, 0x03, 0x0b, 0xd8, 0x11, 0x3c, 0x44, 0x69, 0xb9, 0x6c,
	0x23, 0xc7, 0x5a, 0x88, 0x5c, 0x5a, 0xe4, 0x7b, 0xa6, 0xa5, 0xfc, 0x7d, 0xae, 0x7d, 0x86, 0x46,
	0x3b, 0xed, 0x95, 0x9c, 0x90, 0x22, 0x14, 0xce, 0x0c, 0xb8, 0xd3, 0x98, 0x88, 0x95, 0xfb, 0xbf,
	0x30, 0x1f, 0x8f, 0xbe, 0xe7, 0x41, 0x43, 0x28, 0xd9, 0xa0, 0xab, 0x5a, 0x32, 0x50, 0x5d, 0x1e,
	0x0f, 0xfa, 0x5a, 0xda, 0xf2, 0x05, 0x99, 0x23, 0xd5, 0x7c, 0x6d, 0x99, 0x4d, 0x75, 0x60, 0x13,
	0xc6, 0xb9, 0x1b, 0xa0, 0x68, 0xd2, 0x61, 0xa5, 0x91, 0x16, 0xbc, 0x9b, 0x3f, 0x50, 0x7e, 0x6b,
	0x80, 0x6e, 0x43, 0x46, 0xb6, 0x31, 0xcd, 0x5e, 0xfd, 0x40, 0x36, 0x8b, 0x9a, 0x63, 0x88, 0xc8,
	0xd1, 0xbb, 0xcd, 0x42, 0x2e, 0x55, 0xe8, 0x09, 0x81, 0x7f, 0x63, 0x3e, 0x62, 0xb8, 0xf0, 0xee,
	0x67, 0x84, 0xb2, 0x57, 0xf6, 0x09, 0x4b, 0x31, 0xd2, 0xeb, 0x16, 0x57, 0x9e, 0x54, 0xef, 0xfa,
	0x37, 0xcc, 0x4c, 0xee, 0xa5, 0x97, 0x04, 0x8a, 0x63, 0xa5, 0x88, 0x31, 0x53, 0xcd, 0xd7, 0xf0,
	0x0b, 0x19, 0xc7, 0x65, 0x2e, 0x03, 0xeb, 0x9a, 0x85, 0x11, 0x70, 0xef, 0xe1, 0x17, 0xfc, 0x9f,
	0xd4, 0x47, 0xcf, 0x08, 0x64, 0xad, 0x6b, 0x39, 0x91, 0x9e, 0xf3, 0xf8, 0xbb, 0x51, 0x59, 0x1c,
	0x9f, 0x9c, 0x3b, 0x41, 0xf1, 0xee, 0x08, 0xc0, 0xb3, 0x4a, 0xaf, 0x08, 0xe4, 0x7c, 0xdd, 0x53,
	0x4e, 0x18, 0x9b, 0x5e, 0xf4, 0x94, 0xfc, 0x20, 0x27, 0x4b, 0x69, 0x92, 0x5b, 0x0f, 0xe1, 0xbc,
	0x22, 0x14, 0x46, 0x4a, 0x4b, 0xe7, 0x04, 0x16, 0x0f, 0x14, 0x1a, 0x3e, 0x15, 0x96, 0x8b, 0xbe,
	0xa3, 0xb3, 0x6c, 0xbd, 0xa7, 0x02, 0x14, 0x66, 0x53, 0xb8, 0x23, 0x6d, 0x0e, 0xed, 0x96, 0x50,
	0x56, 0x1b, 0x5b, 0xde, 0x89, 0xff, 0x4d, 0x65, 0xfa, 0xce, 0xcd, 0x4a, 0x14, 0xb7, 0x81, 0xd2,
	0x36, 0x3a, 0x22, 0xfa, 0xd4, 0xdb, 0xd8, 0x78, 0xf9, 0x14, 0xd6, 0xfa, 0xee, 0x31, 0x00, 0x00,
	0xff, 0xff, 0xa4, 0xcb, 0x66, 0x68, 0xa4, 0x04, 0x00, 0x00,
}
