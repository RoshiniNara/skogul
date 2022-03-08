// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pfed_oc.proto

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

type JunosPfe struct {
	NpuMemory            []*JunosPfeNpuMemoryList `protobuf:"bytes,151,rep,name=npu_memory,json=npuMemory" json:"npu_memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *JunosPfe) Reset()         { *m = JunosPfe{} }
func (m *JunosPfe) String() string { return proto.CompactTextString(m) }
func (*JunosPfe) ProtoMessage()    {}
func (*JunosPfe) Descriptor() ([]byte, []int) {
	return fileDescriptor_5714edf97ed06d9c, []int{0}
}
func (m *JunosPfe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosPfe.Unmarshal(m, b)
}
func (m *JunosPfe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosPfe.Marshal(b, m, deterministic)
}
func (m *JunosPfe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosPfe.Merge(m, src)
}
func (m *JunosPfe) XXX_Size() int {
	return xxx_messageInfo_JunosPfe.Size(m)
}
func (m *JunosPfe) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosPfe.DiscardUnknown(m)
}

var xxx_messageInfo_JunosPfe proto.InternalMessageInfo

func (m *JunosPfe) GetNpuMemory() []*JunosPfeNpuMemoryList {
	if m != nil {
		return m.NpuMemory
	}
	return nil
}

type JunosPfeNpuMemoryList struct {
	PfeName              *string  `protobuf:"bytes,51,opt,name=pfe_name,json=pfeName" json:"pfe_name,omitempty"`
	CombinedPoolName     *string  `protobuf:"bytes,52,opt,name=combined_pool_name,json=combinedPoolName" json:"combined_pool_name,omitempty"`
	CombinedSize         *uint32  `protobuf:"varint,53,opt,name=combined_size,json=combinedSize" json:"combined_size,omitempty"`
	CombinedUsageCount   *uint32  `protobuf:"varint,54,opt,name=combined_usage_count,json=combinedUsageCount" json:"combined_usage_count,omitempty"`
	CombinedUtilization  *uint32  `protobuf:"varint,55,opt,name=combined_utilization,json=combinedUtilization" json:"combined_utilization,omitempty"`
	GlobalPoolName       *string  `protobuf:"bytes,56,opt,name=global_pool_name,json=globalPoolName" json:"global_pool_name,omitempty"`
	GlobalUsageCount     *uint32  `protobuf:"varint,57,opt,name=global_usage_count,json=globalUsageCount" json:"global_usage_count,omitempty"`
	GlobalAllocCount     *uint32  `protobuf:"varint,58,opt,name=global_alloc_count,json=globalAllocCount" json:"global_alloc_count,omitempty"`
	GlobalFreeCount      *uint32  `protobuf:"varint,59,opt,name=global_free_count,json=globalFreeCount" json:"global_free_count,omitempty"`
	LocalPoolName        *string  `protobuf:"bytes,60,opt,name=local_pool_name,json=localPoolName" json:"local_pool_name,omitempty"`
	LocalUsageCount      *uint32  `protobuf:"varint,61,opt,name=local_usage_count,json=localUsageCount" json:"local_usage_count,omitempty"`
	LocalAllocCount      *uint32  `protobuf:"varint,62,opt,name=local_alloc_count,json=localAllocCount" json:"local_alloc_count,omitempty"`
	LocalFreeCount       *uint32  `protobuf:"varint,63,opt,name=local_free_count,json=localFreeCount" json:"local_free_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JunosPfeNpuMemoryList) Reset()         { *m = JunosPfeNpuMemoryList{} }
func (m *JunosPfeNpuMemoryList) String() string { return proto.CompactTextString(m) }
func (*JunosPfeNpuMemoryList) ProtoMessage()    {}
func (*JunosPfeNpuMemoryList) Descriptor() ([]byte, []int) {
	return fileDescriptor_5714edf97ed06d9c, []int{0, 0}
}
func (m *JunosPfeNpuMemoryList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosPfeNpuMemoryList.Unmarshal(m, b)
}
func (m *JunosPfeNpuMemoryList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosPfeNpuMemoryList.Marshal(b, m, deterministic)
}
func (m *JunosPfeNpuMemoryList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosPfeNpuMemoryList.Merge(m, src)
}
func (m *JunosPfeNpuMemoryList) XXX_Size() int {
	return xxx_messageInfo_JunosPfeNpuMemoryList.Size(m)
}
func (m *JunosPfeNpuMemoryList) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosPfeNpuMemoryList.DiscardUnknown(m)
}

var xxx_messageInfo_JunosPfeNpuMemoryList proto.InternalMessageInfo

func (m *JunosPfeNpuMemoryList) GetPfeName() string {
	if m != nil && m.PfeName != nil {
		return *m.PfeName
	}
	return ""
}

func (m *JunosPfeNpuMemoryList) GetCombinedPoolName() string {
	if m != nil && m.CombinedPoolName != nil {
		return *m.CombinedPoolName
	}
	return ""
}

func (m *JunosPfeNpuMemoryList) GetCombinedSize() uint32 {
	if m != nil && m.CombinedSize != nil {
		return *m.CombinedSize
	}
	return 0
}

func (m *JunosPfeNpuMemoryList) GetCombinedUsageCount() uint32 {
	if m != nil && m.CombinedUsageCount != nil {
		return *m.CombinedUsageCount
	}
	return 0
}

func (m *JunosPfeNpuMemoryList) GetCombinedUtilization() uint32 {
	if m != nil && m.CombinedUtilization != nil {
		return *m.CombinedUtilization
	}
	return 0
}

func (m *JunosPfeNpuMemoryList) GetGlobalPoolName() string {
	if m != nil && m.GlobalPoolName != nil {
		return *m.GlobalPoolName
	}
	return ""
}

func (m *JunosPfeNpuMemoryList) GetGlobalUsageCount() uint32 {
	if m != nil && m.GlobalUsageCount != nil {
		return *m.GlobalUsageCount
	}
	return 0
}

func (m *JunosPfeNpuMemoryList) GetGlobalAllocCount() uint32 {
	if m != nil && m.GlobalAllocCount != nil {
		return *m.GlobalAllocCount
	}
	return 0
}

func (m *JunosPfeNpuMemoryList) GetGlobalFreeCount() uint32 {
	if m != nil && m.GlobalFreeCount != nil {
		return *m.GlobalFreeCount
	}
	return 0
}

func (m *JunosPfeNpuMemoryList) GetLocalPoolName() string {
	if m != nil && m.LocalPoolName != nil {
		return *m.LocalPoolName
	}
	return ""
}

func (m *JunosPfeNpuMemoryList) GetLocalUsageCount() uint32 {
	if m != nil && m.LocalUsageCount != nil {
		return *m.LocalUsageCount
	}
	return 0
}

func (m *JunosPfeNpuMemoryList) GetLocalAllocCount() uint32 {
	if m != nil && m.LocalAllocCount != nil {
		return *m.LocalAllocCount
	}
	return 0
}

func (m *JunosPfeNpuMemoryList) GetLocalFreeCount() uint32 {
	if m != nil && m.LocalFreeCount != nil {
		return *m.LocalFreeCount
	}
	return 0
}

var E_JnprJunosPfeExt = &proto.ExtensionDesc{
	ExtendedType:  (*JuniperNetworksSensors)(nil),
	ExtensionType: (*JunosPfe)(nil),
	Field:         57,
	Name:          "jnpr_junos_pfe_ext",
	Tag:           "bytes,57,opt,name=jnpr_junos_pfe_ext",
	Filename:      "pfed_oc.proto",
}

func init() {
	proto.RegisterType((*JunosPfe)(nil), "junos_pfe")
	proto.RegisterType((*JunosPfeNpuMemoryList)(nil), "junos_pfe.npu_memory_list")
	proto.RegisterExtension(E_JnprJunosPfeExt)
}

func init() { proto.RegisterFile("pfed_oc.proto", fileDescriptor_5714edf97ed06d9c) }

var fileDescriptor_5714edf97ed06d9c = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0x5d, 0x6b, 0xd4, 0x40,
	0x14, 0x86, 0x59, 0xb5, 0xd8, 0x3d, 0x75, 0xbb, 0xdb, 0xa9, 0x60, 0xd8, 0xab, 0x45, 0x41, 0x42,
	0x29, 0x8b, 0xd6, 0xef, 0xf5, 0x5b, 0xd1, 0x8b, 0x82, 0x65, 0x49, 0xf1, 0x7a, 0x48, 0xd3, 0x93,
	0x32, 0x75, 0x32, 0x67, 0x98, 0x4c, 0xb0, 0xbb, 0x97, 0xfe, 0x03, 0xaf, 0xfc, 0x49, 0xfe, 0x2d,
	0xc9, 0x4c, 0x76, 0x32, 0xf6, 0x72, 0xce, 0xfb, 0x9c, 0x37, 0x0f, 0xe1, 0xc0, 0x48, 0x97, 0x78,
	0xce, 0xa9, 0x98, 0x6b, 0x43, 0x96, 0xa6, 0xfb, 0x16, 0x25, 0x56, 0x68, 0xcd, 0x8a, 0x5b, 0xd2,
	0x7e, 0x78, 0xff, 0xf7, 0x16, 0x0c, 0x2f, 0x1b, 0x45, 0x35, 0xd7, 0x25, 0xb2, 0x05, 0x80, 0xd2,
	0x0d, 0xaf, 0xb0, 0x22, 0xb3, 0x4a, 0xfe, 0x0c, 0x66, 0x37, 0xd3, 0x9d, 0xa3, 0xe9, 0x3c, 0x00,
	0xf3, 0x3e, 0xe5, 0x52, 0xd4, 0x36, 0x1b, 0x2a, 0xdd, 0x7c, 0x73, 0xef, 0xe9, 0xdf, 0x5b, 0x30,
	0xbe, 0x16, 0xb3, 0x19, 0x6c, 0xeb, 0x12, 0xb9, 0xca, 0x2b, 0x4c, 0x9e, 0xcc, 0x06, 0xe9, 0xf0,
	0xd3, 0xd6, 0xaf, 0x0f, 0x37, 0xb6, 0x07, 0xd9, 0x6d, 0x5d, 0xe2, 0x49, 0x5e, 0x21, 0x3b, 0x04,
	0x56, 0x50, 0x75, 0x26, 0x14, 0x9e, 0x73, 0x4d, 0x24, 0x3d, 0xfb, 0xb4, 0x65, 0xb3, 0xc9, 0x26,
	0x59, 0x12, 0x49, 0x47, 0x3f, 0x80, 0x51, 0xa0, 0x6b, 0xb1, 0xc6, 0xe4, 0xd9, 0x6c, 0x90, 0x8e,
	0xb2, 0x3b, 0x9b, 0xe1, 0xa9, 0x58, 0x23, 0x7b, 0x04, 0x77, 0x03, 0xd4, 0xd4, 0xf9, 0x05, 0xf2,
	0x82, 0x1a, 0x65, 0x93, 0xe7, 0x8e, 0x0d, 0x9f, 0xfb, 0xde, 0x46, 0x9f, 0xdb, 0x84, 0x3d, 0x8e,
	0x37, 0xac, 0x90, 0x62, 0x9d, 0x5b, 0x41, 0x2a, 0x79, 0xe1, 0x36, 0xf6, 0xc3, 0x46, 0x1f, 0xb1,
	0x14, 0x26, 0x17, 0x92, 0xce, 0x72, 0x19, 0x59, 0xbf, 0x74, 0xd6, 0xbb, 0x7e, 0x1e, 0x9c, 0x0f,
	0x81, 0x75, 0x64, 0x2c, 0xf3, 0xca, 0x55, 0x77, 0x1d, 0x91, 0x4a, 0x4f, 0xe7, 0x52, 0x52, 0xd1,
	0xd1, 0x8b, 0x98, 0xfe, 0xd8, 0x06, 0x9e, 0x3e, 0x80, 0xbd, 0x8e, 0x2e, 0x0d, 0x6e, 0xaa, 0x5f,
	0x3b, 0x78, 0xec, 0x83, 0xaf, 0x06, 0xbb, 0xe6, 0x87, 0x30, 0x96, 0x54, 0xfc, 0x27, 0xfc, 0xc6,
	0x09, 0x8f, 0xdc, 0x38, 0xf8, 0x1e, 0xc0, 0x9e, 0xe7, 0x62, 0xdd, 0xb7, 0xbe, 0xd3, 0x05, 0x91,
	0x6d, 0x60, 0x63, 0xd9, 0x77, 0x11, 0x1b, 0xb9, 0xa6, 0x30, 0xf1, 0x6c, 0xa4, 0xfa, 0xde, 0xa1,
	0xbb, 0x6e, 0x1e, 0x4c, 0x17, 0x4b, 0x60, 0x97, 0x4a, 0x1b, 0x1e, 0xce, 0x8e, 0xe3, 0x95, 0x65,
	0xf7, 0xe6, 0xc7, 0x8d, 0x12, 0x1a, 0xcd, 0x09, 0xda, 0x9f, 0x64, 0x7e, 0xd4, 0xa7, 0xa8, 0x6a,
	0x32, 0xb5, 0xfb, 0x95, 0x3b, 0x47, 0xd0, 0x5f, 0x69, 0x36, 0x6e, 0xd7, 0x8f, 0xdb, 0xe7, 0xb2,
	0xc4, 0x2f, 0x57, 0xf6, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x74, 0x5a, 0x2b, 0x64, 0x0a, 0x03,
	0x00, 0x00,
}
