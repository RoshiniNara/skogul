// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pfe_export_mon_oc.proto

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

type JunosPfeExportMonStats struct {
	System               *JunosPfeExportMonStatsSystemType `protobuf:"bytes,151,opt,name=system" json:"system,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *JunosPfeExportMonStats) Reset()         { *m = JunosPfeExportMonStats{} }
func (m *JunosPfeExportMonStats) String() string { return proto.CompactTextString(m) }
func (*JunosPfeExportMonStats) ProtoMessage()    {}
func (*JunosPfeExportMonStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_daeba95696a69f7b, []int{0}
}
func (m *JunosPfeExportMonStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosPfeExportMonStats.Unmarshal(m, b)
}
func (m *JunosPfeExportMonStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosPfeExportMonStats.Marshal(b, m, deterministic)
}
func (m *JunosPfeExportMonStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosPfeExportMonStats.Merge(m, src)
}
func (m *JunosPfeExportMonStats) XXX_Size() int {
	return xxx_messageInfo_JunosPfeExportMonStats.Size(m)
}
func (m *JunosPfeExportMonStats) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosPfeExportMonStats.DiscardUnknown(m)
}

var xxx_messageInfo_JunosPfeExportMonStats proto.InternalMessageInfo

func (m *JunosPfeExportMonStats) GetSystem() *JunosPfeExportMonStatsSystemType {
	if m != nil {
		return m.System
	}
	return nil
}

type JunosPfeExportMonStatsSystemType struct {
	Linecard             *JunosPfeExportMonStatsSystemTypeLinecardType `protobuf:"bytes,151,opt,name=linecard" json:"linecard,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *JunosPfeExportMonStatsSystemType) Reset()         { *m = JunosPfeExportMonStatsSystemType{} }
func (m *JunosPfeExportMonStatsSystemType) String() string { return proto.CompactTextString(m) }
func (*JunosPfeExportMonStatsSystemType) ProtoMessage()    {}
func (*JunosPfeExportMonStatsSystemType) Descriptor() ([]byte, []int) {
	return fileDescriptor_daeba95696a69f7b, []int{0, 0}
}
func (m *JunosPfeExportMonStatsSystemType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosPfeExportMonStatsSystemType.Unmarshal(m, b)
}
func (m *JunosPfeExportMonStatsSystemType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosPfeExportMonStatsSystemType.Marshal(b, m, deterministic)
}
func (m *JunosPfeExportMonStatsSystemType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosPfeExportMonStatsSystemType.Merge(m, src)
}
func (m *JunosPfeExportMonStatsSystemType) XXX_Size() int {
	return xxx_messageInfo_JunosPfeExportMonStatsSystemType.Size(m)
}
func (m *JunosPfeExportMonStatsSystemType) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosPfeExportMonStatsSystemType.DiscardUnknown(m)
}

var xxx_messageInfo_JunosPfeExportMonStatsSystemType proto.InternalMessageInfo

func (m *JunosPfeExportMonStatsSystemType) GetLinecard() *JunosPfeExportMonStatsSystemTypeLinecardType {
	if m != nil {
		return m.Linecard
	}
	return nil
}

type JunosPfeExportMonStatsSystemTypeLinecardType struct {
	ExportMonitor        *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorType `protobuf:"bytes,151,opt,name=export_monitor,json=exportMonitor" json:"export_monitor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                       `json:"-"`
	XXX_unrecognized     []byte                                                         `json:"-"`
	XXX_sizecache        int32                                                          `json:"-"`
}

func (m *JunosPfeExportMonStatsSystemTypeLinecardType) Reset() {
	*m = JunosPfeExportMonStatsSystemTypeLinecardType{}
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardType) String() string {
	return proto.CompactTextString(m)
}
func (*JunosPfeExportMonStatsSystemTypeLinecardType) ProtoMessage() {}
func (*JunosPfeExportMonStatsSystemTypeLinecardType) Descriptor() ([]byte, []int) {
	return fileDescriptor_daeba95696a69f7b, []int{0, 0, 0}
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardType.Unmarshal(m, b)
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardType.Marshal(b, m, deterministic)
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardType.Merge(m, src)
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardType) XXX_Size() int {
	return xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardType.Size(m)
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardType) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardType.DiscardUnknown(m)
}

var xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardType proto.InternalMessageInfo

func (m *JunosPfeExportMonStatsSystemTypeLinecardType) GetExportMonitor() *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorType {
	if m != nil {
		return m.ExportMonitor
	}
	return nil
}

type JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorType struct {
	ExportInfo           []*JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList `protobuf:"bytes,152,rep,name=export_info,json=exportInfo" json:"export_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                       `json:"-"`
	XXX_unrecognized     []byte                                                                         `json:"-"`
	XXX_sizecache        int32                                                                          `json:"-"`
}

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorType) Reset() {
	*m = JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorType{}
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorType) String() string {
	return proto.CompactTextString(m)
}
func (*JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorType) ProtoMessage() {}
func (*JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorType) Descriptor() ([]byte, []int) {
	return fileDescriptor_daeba95696a69f7b, []int{0, 0, 0, 0}
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorType.Unmarshal(m, b)
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorType.Marshal(b, m, deterministic)
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorType.Merge(m, src)
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorType) XXX_Size() int {
	return xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorType.Size(m)
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorType) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorType.DiscardUnknown(m)
}

var xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorType proto.InternalMessageInfo

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorType) GetExportInfo() []*JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList {
	if m != nil {
		return m.ExportInfo
	}
	return nil
}

type JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList struct {
	Resource             *string                                                                                    `protobuf:"bytes,153,opt,name=resource" json:"resource,omitempty"`
	RepInterval          *uint32                                                                                    `protobuf:"varint,154,opt,name=rep_interval,json=repInterval" json:"rep_interval,omitempty"`
	PayloadSize          *uint32                                                                                    `protobuf:"varint,155,opt,name=payload_size,json=payloadSize" json:"payload_size,omitempty"`
	QosOptions           *uint32                                                                                    `protobuf:"varint,156,opt,name=qos_options,json=qosOptions" json:"qos_options,omitempty"`
	FcOptions            *uint32                                                                                    `protobuf:"varint,157,opt,name=fc_options,json=fcOptions" json:"fc_options,omitempty"`
	PlpOptions           *uint32                                                                                    `protobuf:"varint,158,opt,name=plp_options,json=plpOptions" json:"plp_options,omitempty"`
	Server               *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListServerType     `protobuf:"bytes,159,opt,name=server" json:"server,omitempty"`
	NumClients           *uint32                                                                                    `protobuf:"varint,162,opt,name=num_clients,json=numClients" json:"num_clients,omitempty"`
	Clients              []*JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListClientsList  `protobuf:"bytes,163,rep,name=clients" json:"clients,omitempty"`
	Accounting           *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListAccountingType `protobuf:"bytes,167,opt,name=accounting" json:"accounting,omitempty"`
	Wraps                []*JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListWrapsList    `protobuf:"bytes,177,rep,name=wraps" json:"wraps,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                                   `json:"-"`
	XXX_unrecognized     []byte                                                                                     `json:"-"`
	XXX_sizecache        int32                                                                                      `json:"-"`
}

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList) Reset() {
	*m = JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList{}
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList) String() string {
	return proto.CompactTextString(m)
}
func (*JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList) ProtoMessage() {}
func (*JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList) Descriptor() ([]byte, []int) {
	return fileDescriptor_daeba95696a69f7b, []int{0, 0, 0, 0, 0}
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList.Unmarshal(m, b)
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList.Marshal(b, m, deterministic)
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList.Merge(m, src)
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList) XXX_Size() int {
	return xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList.Size(m)
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList.DiscardUnknown(m)
}

var xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList proto.InternalMessageInfo

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList) GetResource() string {
	if m != nil && m.Resource != nil {
		return *m.Resource
	}
	return ""
}

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList) GetRepInterval() uint32 {
	if m != nil && m.RepInterval != nil {
		return *m.RepInterval
	}
	return 0
}

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList) GetPayloadSize() uint32 {
	if m != nil && m.PayloadSize != nil {
		return *m.PayloadSize
	}
	return 0
}

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList) GetQosOptions() uint32 {
	if m != nil && m.QosOptions != nil {
		return *m.QosOptions
	}
	return 0
}

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList) GetFcOptions() uint32 {
	if m != nil && m.FcOptions != nil {
		return *m.FcOptions
	}
	return 0
}

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList) GetPlpOptions() uint32 {
	if m != nil && m.PlpOptions != nil {
		return *m.PlpOptions
	}
	return 0
}

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList) GetServer() *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListServerType {
	if m != nil {
		return m.Server
	}
	return nil
}

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList) GetNumClients() uint32 {
	if m != nil && m.NumClients != nil {
		return *m.NumClients
	}
	return 0
}

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList) GetClients() []*JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListClientsList {
	if m != nil {
		return m.Clients
	}
	return nil
}

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList) GetAccounting() *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListAccountingType {
	if m != nil {
		return m.Accounting
	}
	return nil
}

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList) GetWraps() []*JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListWrapsList {
	if m != nil {
		return m.Wraps
	}
	return nil
}

type JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListServerType struct {
	Ip                   *string  `protobuf:"bytes,160,opt,name=ip" json:"ip,omitempty"`
	Port                 *uint32  `protobuf:"varint,161,opt,name=port" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListServerType) Reset() {
	*m = JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListServerType{}
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListServerType) String() string {
	return proto.CompactTextString(m)
}
func (*JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListServerType) ProtoMessage() {
}
func (*JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListServerType) Descriptor() ([]byte, []int) {
	return fileDescriptor_daeba95696a69f7b, []int{0, 0, 0, 0, 0, 0}
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListServerType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListServerType.Unmarshal(m, b)
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListServerType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListServerType.Marshal(b, m, deterministic)
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListServerType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListServerType.Merge(m, src)
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListServerType) XXX_Size() int {
	return xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListServerType.Size(m)
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListServerType) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListServerType.DiscardUnknown(m)
}

var xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListServerType proto.InternalMessageInfo

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListServerType) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListServerType) GetPort() uint32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

type JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListClientsList struct {
	Index                *uint32  `protobuf:"varint,164,opt,name=index" json:"index,omitempty"`
	Ip                   *string  `protobuf:"bytes,165,opt,name=ip" json:"ip,omitempty"`
	Port                 *uint32  `protobuf:"varint,166,opt,name=port" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListClientsList) Reset() {
	*m = JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListClientsList{}
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListClientsList) String() string {
	return proto.CompactTextString(m)
}
func (*JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListClientsList) ProtoMessage() {
}
func (*JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListClientsList) Descriptor() ([]byte, []int) {
	return fileDescriptor_daeba95696a69f7b, []int{0, 0, 0, 0, 0, 1}
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListClientsList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListClientsList.Unmarshal(m, b)
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListClientsList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListClientsList.Marshal(b, m, deterministic)
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListClientsList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListClientsList.Merge(m, src)
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListClientsList) XXX_Size() int {
	return xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListClientsList.Size(m)
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListClientsList) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListClientsList.DiscardUnknown(m)
}

var xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListClientsList proto.InternalMessageInfo

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListClientsList) GetIndex() uint32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListClientsList) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListClientsList) GetPort() uint32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

type JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListAccountingType struct {
	SuccessfulReaps            *uint64  `protobuf:"varint,168,opt,name=successful_reaps,json=successfulReaps" json:"successful_reaps,omitempty"`
	FailedReaps                *uint64  `protobuf:"varint,169,opt,name=failed_reaps,json=failedReaps" json:"failed_reaps,omitempty"`
	ReapsInLastRepInterval     *uint32  `protobuf:"varint,170,opt,name=reaps_in_last_rep_interval,json=reapsInLastRepInterval" json:"reaps_in_last_rep_interval,omitempty"`
	PacketsInLastRepInterval   *uint32  `protobuf:"varint,171,opt,name=packets_in_last_rep_interval,json=packetsInLastRepInterval" json:"packets_in_last_rep_interval,omitempty"`
	InstancesInLastRepInterval *uint32  `protobuf:"varint,172,opt,name=instances_in_last_rep_interval,json=instancesInLastRepInterval" json:"instances_in_last_rep_interval,omitempty"`
	TotalPackets               *uint64  `protobuf:"varint,173,opt,name=total_packets,json=totalPackets" json:"total_packets,omitempty"`
	NumWraps                   *uint64  `protobuf:"varint,174,opt,name=num_wraps,json=numWraps" json:"num_wraps,omitempty"`
	AverageReaps               *uint64  `protobuf:"varint,175,opt,name=average_reaps,json=averageReaps" json:"average_reaps,omitempty"`
	LastPacketSize             *uint32  `protobuf:"varint,176,opt,name=last_packet_size,json=lastPacketSize" json:"last_packet_size,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListAccountingType) Reset() {
	*m = JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListAccountingType{}
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListAccountingType) String() string {
	return proto.CompactTextString(m)
}
func (*JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListAccountingType) ProtoMessage() {
}
func (*JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListAccountingType) Descriptor() ([]byte, []int) {
	return fileDescriptor_daeba95696a69f7b, []int{0, 0, 0, 0, 0, 2}
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListAccountingType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListAccountingType.Unmarshal(m, b)
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListAccountingType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListAccountingType.Marshal(b, m, deterministic)
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListAccountingType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListAccountingType.Merge(m, src)
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListAccountingType) XXX_Size() int {
	return xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListAccountingType.Size(m)
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListAccountingType) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListAccountingType.DiscardUnknown(m)
}

var xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListAccountingType proto.InternalMessageInfo

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListAccountingType) GetSuccessfulReaps() uint64 {
	if m != nil && m.SuccessfulReaps != nil {
		return *m.SuccessfulReaps
	}
	return 0
}

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListAccountingType) GetFailedReaps() uint64 {
	if m != nil && m.FailedReaps != nil {
		return *m.FailedReaps
	}
	return 0
}

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListAccountingType) GetReapsInLastRepInterval() uint32 {
	if m != nil && m.ReapsInLastRepInterval != nil {
		return *m.ReapsInLastRepInterval
	}
	return 0
}

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListAccountingType) GetPacketsInLastRepInterval() uint32 {
	if m != nil && m.PacketsInLastRepInterval != nil {
		return *m.PacketsInLastRepInterval
	}
	return 0
}

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListAccountingType) GetInstancesInLastRepInterval() uint32 {
	if m != nil && m.InstancesInLastRepInterval != nil {
		return *m.InstancesInLastRepInterval
	}
	return 0
}

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListAccountingType) GetTotalPackets() uint64 {
	if m != nil && m.TotalPackets != nil {
		return *m.TotalPackets
	}
	return 0
}

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListAccountingType) GetNumWraps() uint64 {
	if m != nil && m.NumWraps != nil {
		return *m.NumWraps
	}
	return 0
}

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListAccountingType) GetAverageReaps() uint64 {
	if m != nil && m.AverageReaps != nil {
		return *m.AverageReaps
	}
	return 0
}

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListAccountingType) GetLastPacketSize() uint32 {
	if m != nil && m.LastPacketSize != nil {
		return *m.LastPacketSize
	}
	return 0
}

type JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListWrapsList struct {
	Index                *uint32  `protobuf:"varint,178,opt,name=index" json:"index,omitempty"`
	Time                 *string  `protobuf:"bytes,179,opt,name=time" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListWrapsList) Reset() {
	*m = JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListWrapsList{}
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListWrapsList) String() string {
	return proto.CompactTextString(m)
}
func (*JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListWrapsList) ProtoMessage() {
}
func (*JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListWrapsList) Descriptor() ([]byte, []int) {
	return fileDescriptor_daeba95696a69f7b, []int{0, 0, 0, 0, 0, 3}
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListWrapsList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListWrapsList.Unmarshal(m, b)
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListWrapsList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListWrapsList.Marshal(b, m, deterministic)
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListWrapsList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListWrapsList.Merge(m, src)
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListWrapsList) XXX_Size() int {
	return xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListWrapsList.Size(m)
}
func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListWrapsList) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListWrapsList.DiscardUnknown(m)
}

var xxx_messageInfo_JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListWrapsList proto.InternalMessageInfo

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListWrapsList) GetIndex() uint32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListWrapsList) GetTime() string {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return ""
}

var E_JnprJunosPfeExportMonStatsExt = &proto.ExtensionDesc{
	ExtendedType:  (*JuniperNetworksSensors)(nil),
	ExtensionType: (*JunosPfeExportMonStats)(nil),
	Field:         129,
	Name:          "jnpr_junos_pfe_export_mon_stats_ext",
	Tag:           "bytes,129,opt,name=jnpr_junos_pfe_export_mon_stats_ext",
	Filename:      "pfe_export_mon_oc.proto",
}

func init() {
	proto.RegisterType((*JunosPfeExportMonStats)(nil), "junos_pfe_export_mon_stats")
	proto.RegisterType((*JunosPfeExportMonStatsSystemType)(nil), "junos_pfe_export_mon_stats.system_type")
	proto.RegisterType((*JunosPfeExportMonStatsSystemTypeLinecardType)(nil), "junos_pfe_export_mon_stats.system_type.linecard_type")
	proto.RegisterType((*JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorType)(nil), "junos_pfe_export_mon_stats.system_type.linecard_type.export_monitor_type")
	proto.RegisterType((*JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoList)(nil), "junos_pfe_export_mon_stats.system_type.linecard_type.export_monitor_type.export_info_list")
	proto.RegisterType((*JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListServerType)(nil), "junos_pfe_export_mon_stats.system_type.linecard_type.export_monitor_type.export_info_list.server_type")
	proto.RegisterType((*JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListClientsList)(nil), "junos_pfe_export_mon_stats.system_type.linecard_type.export_monitor_type.export_info_list.clients_list")
	proto.RegisterType((*JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListAccountingType)(nil), "junos_pfe_export_mon_stats.system_type.linecard_type.export_monitor_type.export_info_list.accounting_type")
	proto.RegisterType((*JunosPfeExportMonStatsSystemTypeLinecardTypeExportMonitorTypeExportInfoListWrapsList)(nil), "junos_pfe_export_mon_stats.system_type.linecard_type.export_monitor_type.export_info_list.wraps_list")
	proto.RegisterExtension(E_JnprJunosPfeExportMonStatsExt)
}

func init() { proto.RegisterFile("pfe_export_mon_oc.proto", fileDescriptor_daeba95696a69f7b) }

var fileDescriptor_daeba95696a69f7b = []byte{
	// 787 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x4b, 0x6f, 0x1c, 0x45,
	0x10, 0xd6, 0x3a, 0x76, 0xb0, 0x6b, 0xbc, 0xb6, 0xd5, 0x2b, 0x91, 0xd6, 0x60, 0xa2, 0x95, 0x41,
	0x22, 0x70, 0x98, 0x43, 0x10, 0x17, 0x38, 0x24, 0x22, 0xca, 0x61, 0x23, 0x1e, 0xd1, 0xf8, 0x80,
	0xc4, 0xa5, 0x35, 0x9a, 0xad, 0xb1, 0xda, 0x99, 0xed, 0x6e, 0x77, 0xf7, 0x38, 0xbb, 0x51, 0x24,
	0x94, 0x5c, 0xf9, 0x01, 0x3c, 0x12, 0xde, 0xef, 0x37, 0x84, 0x03, 0xe2, 0xbf, 0x71, 0x40, 0x3d,
	0x35, 0xb3, 0xbb, 0xb6, 0xb3, 0x08, 0x21, 0xf6, 0x38, 0x5f, 0x7d, 0xf5, 0x7d, 0xb5, 0xdf, 0x54,
	0xcd, 0xc2, 0x05, 0x53, 0xa0, 0xc0, 0xb1, 0xd1, 0xd6, 0x8b, 0x91, 0x56, 0x42, 0xe7, 0x89, 0xb1,
	0xda, 0xeb, 0xb8, 0xe7, 0xb1, 0xc4, 0x11, 0x7a, 0x3b, 0x11, 0x5e, 0x1b, 0x02, 0xf7, 0xfe, 0xda,
	0x82, 0xf8, 0xb0, 0x52, 0xda, 0x89, 0x53, 0x6d, 0xce, 0x67, 0xde, 0xb1, 0xab, 0x70, 0xde, 0x4d,
	0x9c, 0xc7, 0x11, 0x7f, 0xaf, 0xd3, 0xef, 0x5c, 0x8a, 0x2e, 0x3f, 0x97, 0x2c, 0x66, 0x27, 0x44,
	0x15, 0x7e, 0x62, 0x30, 0x6d, 0xfa, 0xe2, 0x07, 0x5b, 0x10, 0xcd, 0xe1, 0x2c, 0x85, 0xf5, 0x52,
	0x2a, 0xcc, 0x33, 0x3b, 0x6c, 0x35, 0x5f, 0xfa, 0x97, 0x9a, 0x49, 0xdb, 0x48, 0x0e, 0x53, 0x9d,
	0xf8, 0x8f, 0x2e, 0x74, 0x4f, 0xd4, 0xd8, 0x11, 0x6c, 0xcd, 0x94, 0xa4, 0xd7, 0xb6, 0xf5, 0x1a,
	0xfc, 0x27, 0xaf, 0xe4, 0xa4, 0x18, 0xf9, 0x77, 0x09, 0x7c, 0x9d, 0xb0, 0xf8, 0xd1, 0x26, 0xf4,
	0x1e, 0x43, 0x63, 0x77, 0x21, 0x6a, 0x60, 0xa9, 0x0a, 0xcd, 0xdf, 0xef, 0xf4, 0xcf, 0x5d, 0x8a,
	0x2e, 0xbf, 0xfd, 0xbf, 0xcd, 0x91, 0xcc, 0xa9, 0x8b, 0x52, 0x3a, 0x9f, 0x02, 0x21, 0x03, 0x55,
	0xe8, 0xf8, 0x61, 0x04, 0x3b, 0xa7, 0x09, 0x6c, 0x0f, 0xd6, 0x2d, 0x3a, 0x5d, 0xd9, 0x1c, 0xf9,
	0x07, 0x21, 0x97, 0x8d, 0x57, 0xd7, 0xee, 0x5f, 0x5d, 0x59, 0xef, 0xa4, 0x53, 0x9c, 0xed, 0xc1,
	0xa6, 0x45, 0x23, 0xa4, 0xf2, 0x68, 0x8f, 0xb3, 0x92, 0x7f, 0x18, 0x78, 0xdd, 0x34, 0xb2, 0x68,
	0x06, 0x0d, 0x16, 0x38, 0x26, 0x9b, 0x94, 0x3a, 0x1b, 0x0a, 0x27, 0xef, 0x20, 0x7f, 0xd0, 0x70,
	0x1a, 0x70, 0x5f, 0xde, 0x41, 0xd6, 0x87, 0xe8, 0x48, 0x3b, 0xa1, 0x8d, 0x97, 0x5a, 0x39, 0xfe,
	0x90, 0x28, 0x70, 0xa4, 0xdd, 0x9b, 0x04, 0xb1, 0x8b, 0x00, 0x45, 0x3e, 0x25, 0x7c, 0x44, 0x84,
	0x8d, 0x22, 0x6f, 0xeb, 0x7d, 0x88, 0x4c, 0x69, 0xa6, 0x84, 0x8f, 0x1b, 0x05, 0x53, 0x9a, 0x96,
	0xf1, 0x0e, 0x9c, 0x77, 0x68, 0x8f, 0xd1, 0xf2, 0x4f, 0xe8, 0x2d, 0x17, 0xcb, 0x4b, 0x37, 0x21,
	0xa7, 0x76, 0xc9, 0xeb, 0x87, 0x30, 0xa2, 0xaa, 0x46, 0x22, 0x2f, 0x25, 0x2a, 0xef, 0xf8, 0xe7,
	0xcd, 0x88, 0xaa, 0x1a, 0x5d, 0x23, 0x88, 0xdd, 0xef, 0xc0, 0x13, 0x6d, 0xf9, 0x0b, 0x5a, 0x81,
	0x83, 0x25, 0x0e, 0xd9, 0x58, 0xd1, 0x3e, 0xb4, 0xc6, 0xec, 0xdd, 0x0e, 0x40, 0x96, 0xe7, 0xba,
	0x52, 0x5e, 0xaa, 0x03, 0xfe, 0x0d, 0x85, 0x75, 0xb8, 0xc4, 0x39, 0x66, 0x6e, 0x14, 0xd8, 0x9c,
	0x3d, 0xbb, 0x0b, 0x6b, 0xb7, 0x6d, 0x66, 0x1c, 0xff, 0x8d, 0xf2, 0xc0, 0x25, 0xce, 0x51, 0x1b,
	0x51, 0x1a, 0x64, 0x1a, 0xbf, 0x08, 0xd1, 0xdc, 0x9b, 0x64, 0xdb, 0xb0, 0x22, 0x0d, 0xff, 0xb4,
	0x3e, 0x86, 0x74, 0x45, 0x1a, 0xd6, 0x83, 0xd5, 0xa0, 0xc1, 0x3f, 0xa3, 0x77, 0x59, 0x3f, 0xc4,
	0x29, 0x6c, 0xce, 0x27, 0xcb, 0x76, 0x61, 0x4d, 0xaa, 0x21, 0x8e, 0xf9, 0x97, 0x35, 0xab, 0xbd,
	0x22, 0x02, 0x1b, 0xcd, 0xaf, 0xce, 0x6a, 0x7e, 0x3d, 0xaf, 0xf9, 0xe7, 0x39, 0xd8, 0x3e, 0x15,
	0x13, 0x7b, 0x01, 0x76, 0x5c, 0x95, 0xe7, 0xe8, 0x5c, 0x51, 0x95, 0xc2, 0x62, 0x48, 0xe9, 0xdb,
	0xd0, 0xb4, 0x9a, 0x6e, 0xcf, 0x0a, 0x69, 0xc0, 0xc3, 0x11, 0x16, 0x99, 0x2c, 0x71, 0xd8, 0xf0,
	0xbe, 0x23, 0x5e, 0x44, 0x20, 0x71, 0x5e, 0x81, 0xb8, 0x2e, 0x0a, 0xa9, 0x44, 0x99, 0x39, 0x2f,
	0x4e, 0x9c, 0xf6, 0xf7, 0x34, 0xce, 0x93, 0x35, 0x65, 0xa0, 0x5e, 0xcb, 0x9c, 0x4f, 0xe7, 0xae,
	0xfc, 0x0a, 0xec, 0x9a, 0x2c, 0xbf, 0x85, 0x7e, 0x41, 0xfb, 0x0f, 0xd4, 0xce, 0x1b, 0xd2, 0x59,
	0x81, 0x6b, 0x70, 0x51, 0x2a, 0xe7, 0x33, 0x95, 0xe3, 0x02, 0x89, 0x1f, 0x49, 0x22, 0x9e, 0xd2,
	0xce, 0x8a, 0x3c, 0x0b, 0x5d, 0xaf, 0x7d, 0x56, 0x8a, 0xc6, 0x86, 0xff, 0x44, 0xbf, 0x73, 0xb3,
	0x46, 0x6f, 0x12, 0xc8, 0x76, 0x61, 0x23, 0x1c, 0x22, 0xed, 0xd5, 0xcf, 0xc4, 0x58, 0x57, 0xd5,
	0xe8, 0xad, 0x00, 0x04, 0x8d, 0xec, 0x18, 0x6d, 0x76, 0x80, 0x4d, 0x56, 0xbf, 0x34, 0x1a, 0x0d,
	0x4a, 0x61, 0x3d, 0x0f, 0x3b, 0xf5, 0x84, 0x64, 0x44, 0x5f, 0xb6, 0x5f, 0x69, 0xc0, 0xad, 0x50,
	0x20, 0xaf, 0xf0, 0x71, 0x8b, 0xaf, 0x00, 0xcc, 0x36, 0x6b, 0xb6, 0x0d, 0x8f, 0x1e, 0xb7, 0x0d,
	0x3d, 0x58, 0xf5, 0x72, 0x84, 0xfc, 0x77, 0xda, 0x87, 0xfa, 0xe1, 0xe5, 0x7b, 0x1d, 0x78, 0xe6,
	0x50, 0x19, 0x2b, 0x16, 0xaf, 0xbe, 0xc0, 0xb1, 0x67, 0x17, 0x92, 0x1b, 0x95, 0x92, 0x06, 0xed,
	0x1b, 0xe8, 0x6f, 0x6b, 0x7b, 0xcb, 0xed, 0xa3, 0x72, 0xda, 0x3a, 0x7e, 0x8f, 0x6e, 0xf8, 0xa9,
	0x7f, 0xb8, 0x9d, 0xf4, 0xe9, 0xe0, 0x70, 0x23, 0xd4, 0x6f, 0x16, 0x78, 0xbd, 0xfd, 0xd3, 0xda,
	0x0f, 0xb5, 0xeb, 0x63, 0xff, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbd, 0xe8, 0x1d, 0x7b, 0x31,
	0x08, 0x00, 0x00,
}