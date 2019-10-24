// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pfe_mpls_sr_te_ip_oc.proto

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

type MplsPfeMplsSrTeIp struct {
	SignalingProtocols   *MplsPfeMplsSrTeIpSignalingProtocolsType `protobuf:"bytes,151,opt,name=signaling_protocols,json=signalingProtocols" json:"signaling_protocols,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *MplsPfeMplsSrTeIp) Reset()         { *m = MplsPfeMplsSrTeIp{} }
func (m *MplsPfeMplsSrTeIp) String() string { return proto.CompactTextString(m) }
func (*MplsPfeMplsSrTeIp) ProtoMessage()    {}
func (*MplsPfeMplsSrTeIp) Descriptor() ([]byte, []int) {
	return fileDescriptor_pfe_mpls_sr_te_ip_oc_9dbf7accc912f551, []int{0}
}
func (m *MplsPfeMplsSrTeIp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsSrTeIp.Unmarshal(m, b)
}
func (m *MplsPfeMplsSrTeIp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsSrTeIp.Marshal(b, m, deterministic)
}
func (dst *MplsPfeMplsSrTeIp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsSrTeIp.Merge(dst, src)
}
func (m *MplsPfeMplsSrTeIp) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsSrTeIp.Size(m)
}
func (m *MplsPfeMplsSrTeIp) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsSrTeIp.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsSrTeIp proto.InternalMessageInfo

func (m *MplsPfeMplsSrTeIp) GetSignalingProtocols() *MplsPfeMplsSrTeIpSignalingProtocolsType {
	if m != nil {
		return m.SignalingProtocols
	}
	return nil
}

type MplsPfeMplsSrTeIpSignalingProtocolsType struct {
	SegmentRouting       *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingType `protobuf:"bytes,151,opt,name=segment_routing,json=segmentRouting" json:"segment_routing,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                   `json:"-"`
	XXX_unrecognized     []byte                                                     `json:"-"`
	XXX_sizecache        int32                                                      `json:"-"`
}

func (m *MplsPfeMplsSrTeIpSignalingProtocolsType) Reset() {
	*m = MplsPfeMplsSrTeIpSignalingProtocolsType{}
}
func (m *MplsPfeMplsSrTeIpSignalingProtocolsType) String() string { return proto.CompactTextString(m) }
func (*MplsPfeMplsSrTeIpSignalingProtocolsType) ProtoMessage()    {}
func (*MplsPfeMplsSrTeIpSignalingProtocolsType) Descriptor() ([]byte, []int) {
	return fileDescriptor_pfe_mpls_sr_te_ip_oc_9dbf7accc912f551, []int{0, 0}
}
func (m *MplsPfeMplsSrTeIpSignalingProtocolsType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsType.Unmarshal(m, b)
}
func (m *MplsPfeMplsSrTeIpSignalingProtocolsType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsType.Marshal(b, m, deterministic)
}
func (dst *MplsPfeMplsSrTeIpSignalingProtocolsType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsType.Merge(dst, src)
}
func (m *MplsPfeMplsSrTeIpSignalingProtocolsType) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsType.Size(m)
}
func (m *MplsPfeMplsSrTeIpSignalingProtocolsType) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsType.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsType proto.InternalMessageInfo

func (m *MplsPfeMplsSrTeIpSignalingProtocolsType) GetSegmentRouting() *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingType {
	if m != nil {
		return m.SegmentRouting
	}
	return nil
}

type MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingType struct {
	SrTeIpPolicies       *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesType `protobuf:"bytes,151,opt,name=sr_te_ip_policies,json=srTeIpPolicies" json:"sr_te_ip_policies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                     `json:"-"`
	XXX_unrecognized     []byte                                                                       `json:"-"`
	XXX_sizecache        int32                                                                        `json:"-"`
}

func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingType) Reset() {
	*m = MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingType{}
}
func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingType) String() string {
	return proto.CompactTextString(m)
}
func (*MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingType) ProtoMessage() {}
func (*MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingType) Descriptor() ([]byte, []int) {
	return fileDescriptor_pfe_mpls_sr_te_ip_oc_9dbf7accc912f551, []int{0, 0, 0}
}
func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingType.Unmarshal(m, b)
}
func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingType.Marshal(b, m, deterministic)
}
func (dst *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingType.Merge(dst, src)
}
func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingType) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingType.Size(m)
}
func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingType) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingType.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingType proto.InternalMessageInfo

func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingType) GetSrTeIpPolicies() *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesType {
	if m != nil {
		return m.SrTeIpPolicies
	}
	return nil
}

type MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesType struct {
	SrTeIpPolicy         []*MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyList `protobuf:"bytes,151,rep,name=sr_te_ip_policy,json=srTeIpPolicy" json:"sr_te_ip_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                                       `json:"-"`
	XXX_unrecognized     []byte                                                                                         `json:"-"`
	XXX_sizecache        int32                                                                                          `json:"-"`
}

func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesType) Reset() {
	*m = MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesType{}
}
func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesType) String() string {
	return proto.CompactTextString(m)
}
func (*MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesType) ProtoMessage() {}
func (*MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesType) Descriptor() ([]byte, []int) {
	return fileDescriptor_pfe_mpls_sr_te_ip_oc_9dbf7accc912f551, []int{0, 0, 0, 0}
}
func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesType.Unmarshal(m, b)
}
func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesType.Marshal(b, m, deterministic)
}
func (dst *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesType.Merge(dst, src)
}
func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesType) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesType.Size(m)
}
func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesType) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesType.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesType proto.InternalMessageInfo

func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesType) GetSrTeIpPolicy() []*MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyList {
	if m != nil {
		return m.SrTeIpPolicy
	}
	return nil
}

type MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyList struct {
	State                *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateType `protobuf:"bytes,151,opt,name=state" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                                              `json:"-"`
	XXX_unrecognized     []byte                                                                                                `json:"-"`
	XXX_sizecache        int32                                                                                                 `json:"-"`
}

func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyList) Reset() {
	*m = MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyList{}
}
func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyList) String() string {
	return proto.CompactTextString(m)
}
func (*MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyList) ProtoMessage() {
}
func (*MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyList) Descriptor() ([]byte, []int) {
	return fileDescriptor_pfe_mpls_sr_te_ip_oc_9dbf7accc912f551, []int{0, 0, 0, 0, 0}
}
func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyList.Unmarshal(m, b)
}
func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyList.Marshal(b, m, deterministic)
}
func (dst *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyList.Merge(dst, src)
}
func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyList) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyList.Size(m)
}
func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyList) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyList.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyList proto.InternalMessageInfo

func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyList) GetState() *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateType {
	if m != nil {
		return m.State
	}
	return nil
}

type MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateType struct {
	Counters             []*MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateTypeCountersList `protobuf:"bytes,151,rep,name=counters" json:"counters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                                                            `json:"-"`
	XXX_unrecognized     []byte                                                                                                              `json:"-"`
	XXX_sizecache        int32                                                                                                               `json:"-"`
}

func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateType) Reset() {
	*m = MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateType{}
}
func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateType) String() string {
	return proto.CompactTextString(m)
}
func (*MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateType) ProtoMessage() {
}
func (*MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateType) Descriptor() ([]byte, []int) {
	return fileDescriptor_pfe_mpls_sr_te_ip_oc_9dbf7accc912f551, []int{0, 0, 0, 0, 0, 0}
}
func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateType.Unmarshal(m, b)
}
func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateType.Marshal(b, m, deterministic)
}
func (dst *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateType.Merge(dst, src)
}
func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateType) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateType.Size(m)
}
func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateType) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateType.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateType proto.InternalMessageInfo

func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateType) GetCounters() []*MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateTypeCountersList {
	if m != nil {
		return m.Counters
	}
	return nil
}

type MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateTypeCountersList struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateTypeCountersList) Reset() {
	*m = MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateTypeCountersList{}
}
func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateTypeCountersList) String() string {
	return proto.CompactTextString(m)
}
func (*MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateTypeCountersList) ProtoMessage() {
}
func (*MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateTypeCountersList) Descriptor() ([]byte, []int) {
	return fileDescriptor_pfe_mpls_sr_te_ip_oc_9dbf7accc912f551, []int{0, 0, 0, 0, 0, 0, 0}
}
func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateTypeCountersList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateTypeCountersList.Unmarshal(m, b)
}
func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateTypeCountersList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateTypeCountersList.Marshal(b, m, deterministic)
}
func (dst *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateTypeCountersList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateTypeCountersList.Merge(dst, src)
}
func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateTypeCountersList) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateTypeCountersList.Size(m)
}
func (m *MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateTypeCountersList) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateTypeCountersList.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateTypeCountersList proto.InternalMessageInfo

var E_JnprMplsPfeMplsSrTeIpExt = &proto.ExtensionDesc{
	ExtendedType:  (*JuniperNetworksSensors)(nil),
	ExtensionType: (*MplsPfeMplsSrTeIp)(nil),
	Field:         102,
	Name:          "jnpr_mpls_pfe_mpls_sr_te_ip_ext",
	Tag:           "bytes,102,opt,name=jnpr_mpls_pfe_mpls_sr_te_ip_ext,json=jnprMplsPfeMplsSrTeIpExt",
	Filename:      "pfe_mpls_sr_te_ip_oc.proto",
}

func init() {
	proto.RegisterType((*MplsPfeMplsSrTeIp)(nil), "mpls_pfe_mpls_sr_te_ip")
	proto.RegisterType((*MplsPfeMplsSrTeIpSignalingProtocolsType)(nil), "mpls_pfe_mpls_sr_te_ip.signaling_protocols_type")
	proto.RegisterType((*MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingType)(nil), "mpls_pfe_mpls_sr_te_ip.signaling_protocols_type.segment_routing_type")
	proto.RegisterType((*MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesType)(nil), "mpls_pfe_mpls_sr_te_ip.signaling_protocols_type.segment_routing_type.sr_te_ip_policies_type")
	proto.RegisterType((*MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyList)(nil), "mpls_pfe_mpls_sr_te_ip.signaling_protocols_type.segment_routing_type.sr_te_ip_policies_type.sr_te_ip_policy_list")
	proto.RegisterType((*MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateType)(nil), "mpls_pfe_mpls_sr_te_ip.signaling_protocols_type.segment_routing_type.sr_te_ip_policies_type.sr_te_ip_policy_list.state_type")
	proto.RegisterType((*MplsPfeMplsSrTeIpSignalingProtocolsTypeSegmentRoutingTypeSrTeIpPoliciesTypeSrTeIpPolicyListStateTypeCountersList)(nil), "mpls_pfe_mpls_sr_te_ip.signaling_protocols_type.segment_routing_type.sr_te_ip_policies_type.sr_te_ip_policy_list.state_type.counters_list")
	proto.RegisterExtension(E_JnprMplsPfeMplsSrTeIpExt)
}

func init() {
	proto.RegisterFile("pfe_mpls_sr_te_ip_oc.proto", fileDescriptor_pfe_mpls_sr_te_ip_oc_9dbf7accc912f551)
}

var fileDescriptor_pfe_mpls_sr_te_ip_oc_9dbf7accc912f551 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x91, 0xbd, 0x6a, 0x1b, 0x41,
	0x14, 0x85, 0x99, 0x24, 0x4a, 0xc2, 0x55, 0x12, 0x91, 0x91, 0x90, 0x96, 0x69, 0x12, 0x52, 0xa9,
	0x1a, 0x82, 0xca, 0xf4, 0x2a, 0x12, 0x88, 0x11, 0x2b, 0x97, 0x86, 0x41, 0x5e, 0xae, 0x96, 0xb1,
	0x57, 0x33, 0xe3, 0x99, 0x11, 0xd6, 0x82, 0x1f, 0xc0, 0xb8, 0x75, 0x61, 0xfc, 0x06, 0xae, 0xfc,
	0x1c, 0xae, 0xfd, 0x02, 0xf6, 0x5b, 0xb8, 0x35, 0xda, 0xd5, 0xae, 0x2d, 0x69, 0x55, 0x18, 0xfc,
	0x53, 0x2d, 0x9c, 0x7b, 0xce, 0xbd, 0xdf, 0x9e, 0x01, 0x66, 0xc6, 0x28, 0x26, 0x26, 0x71, 0xc2,
	0x59, 0xe1, 0x51, 0x48, 0x23, 0x74, 0xc4, 0x8d, 0xd5, 0x5e, 0xb3, 0xa6, 0xc7, 0x04, 0x27, 0xe8,
	0x6d, 0x2a, 0xbc, 0x36, 0xb9, 0xf8, 0xeb, 0xe6, 0x13, 0xb4, 0x33, 0xff, 0x5a, 0x90, 0xee, 0x42,
	0xd3, 0xc9, 0x58, 0x8d, 0x12, 0xa9, 0x62, 0x91, 0xb9, 0x23, 0x9d, 0xb8, 0xe0, 0x8c, 0xfc, 0x24,
	0xdd, 0x7a, 0xef, 0x37, 0xaf, 0x8e, 0xf1, 0x8a, 0x8c, 0xf0, 0xa9, 0xc1, 0x90, 0x96, 0x93, 0x41,
	0x31, 0x60, 0xb7, 0x1f, 0x21, 0xd8, 0x14, 0xa0, 0x1a, 0x1a, 0x0e, 0xe3, 0x09, 0x2a, 0x2f, 0xac,
	0x9e, 0x7a, 0xa9, 0xe2, 0xe2, 0x78, 0xff, 0xa9, 0xc7, 0xf9, 0xca, 0xa2, 0x9c, 0xe8, 0xdb, 0x42,
	0x0d, 0x73, 0x91, 0x5d, 0xd7, 0xa0, 0x55, 0x65, 0xa4, 0xc7, 0x04, 0xbe, 0x97, 0x85, 0x1a, 0x9d,
	0xc8, 0x48, 0x62, 0xd9, 0xc4, 0xce, 0xb3, 0xc0, 0xf0, 0xb5, 0xfd, 0x05, 0xa3, 0xdd, 0xc6, 0xbf,
	0x66, 0xb0, 0x10, 0xd9, 0xe5, 0x07, 0x68, 0x57, 0x5b, 0xe9, 0x39, 0x81, 0xc6, 0xf2, 0x28, 0x9d,
	0x33, 0xbe, 0xef, 0xd6, 0x7b, 0x07, 0x2f, 0xc9, 0xb8, 0x22, 0xa7, 0x22, 0x91, 0xce, 0x87, 0x5f,
	0x1e, 0x81, 0xa7, 0xec, 0xee, 0x1d, 0xb4, 0xaa, 0x6c, 0xf4, 0x94, 0x40, 0xcd, 0xf9, 0x91, 0xc7,
	0xa2, 0xce, 0xa3, 0x57, 0x47, 0xe5, 0xd9, 0xfd, 0xbc, 0xee, 0x9c, 0x85, 0x5d, 0x11, 0x80, 0x07,
	0x95, 0x5e, 0x10, 0xf8, 0x1c, 0xe9, 0xa9, 0xf2, 0x68, 0x5d, 0x51, 0xe9, 0x09, 0x79, 0x4b, 0x50,
	0x5e, 0xe0, 0xe4, 0x65, 0x97, 0x74, 0xac, 0x01, 0x5f, 0x97, 0x46, 0x7f, 0x1c, 0xfc, 0xd8, 0x53,
	0xc6, 0x8a, 0x6a, 0x58, 0x81, 0x33, 0x4f, 0x3b, 0xfc, 0xdf, 0x54, 0x49, 0x83, 0x76, 0x0b, 0xfd,
	0xa1, 0xb6, 0xfb, 0x6e, 0x88, 0xca, 0x69, 0xeb, 0x82, 0x71, 0xf6, 0x22, 0x9d, 0x0d, 0x3f, 0x1a,
	0x06, 0xf3, 0xc5, 0xff, 0x4d, 0xe2, 0x06, 0x63, 0x9c, 0x7f, 0x86, 0xd9, 0x8b, 0xf7, 0x67, 0xfe,
	0x3e, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x17, 0xd1, 0xcb, 0x88, 0x04, 0x00, 0x00,
}
