// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cmerror_data.proto

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
// Juniper Error Item information
//
type ErrorData struct {
	// Identifier that uniquely identifies the source of
	// the error.
	Identifier  *string `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	Count       *uint64 `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	LastUpdated *uint64 `protobuf:"varint,3,opt,name=last_updated,json=lastUpdated" json:"last_updated,omitempty"`
	IsEnabled   *bool   `protobuf:"varint,4,opt,name=is_enabled,json=isEnabled" json:"is_enabled,omitempty"`
	// Additional Metadata for error processing
	ModuleId    *uint32 `protobuf:"varint,5,opt,name=module_id,json=moduleId" json:"module_id,omitempty"`
	ComponentId *uint32 `protobuf:"varint,6,opt,name=component_id,json=componentId" json:"component_id,omitempty"`
	ErrorCode   *uint32 `protobuf:"varint,7,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
	// Additional stats for each of the error
	OccurCount        *uint32 `protobuf:"varint,8,opt,name=occur_count,json=occurCount" json:"occur_count,omitempty"`
	ClearedCount      *uint32 `protobuf:"varint,9,opt,name=cleared_count,json=clearedCount" json:"cleared_count,omitempty"`
	LastClearedAt     *uint64 `protobuf:"varint,10,opt,name=last_cleared_at,json=lastClearedAt" json:"last_cleared_at,omitempty"`
	ActionCount       *uint32 `protobuf:"varint,11,opt,name=action_count,json=actionCount" json:"action_count,omitempty"`
	LastActionTakenAt *uint64 `protobuf:"varint,12,opt,name=last_action_taken_at,json=lastActionTakenAt" json:"last_action_taken_at,omitempty"`
	// Fru information
	FruType *string `protobuf:"bytes,13,opt,name=fru_type,json=fruType" json:"fru_type,omitempty"`
	FruSlot *uint32 `protobuf:"varint,14,opt,name=fru_slot,json=fruSlot" json:"fru_slot,omitempty"`
	// Help information regarding the error.
	Description          *string  `protobuf:"bytes,15,opt,name=description" json:"description,omitempty"`
	Help                 *string  `protobuf:"bytes,16,opt,name=help" json:"help,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorData) Reset()         { *m = ErrorData{} }
func (m *ErrorData) String() string { return proto.CompactTextString(m) }
func (*ErrorData) ProtoMessage()    {}
func (*ErrorData) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e62e0d982e26602, []int{0}
}
func (m *ErrorData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorData.Unmarshal(m, b)
}
func (m *ErrorData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorData.Marshal(b, m, deterministic)
}
func (m *ErrorData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorData.Merge(m, src)
}
func (m *ErrorData) XXX_Size() int {
	return xxx_messageInfo_ErrorData.Size(m)
}
func (m *ErrorData) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorData.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorData proto.InternalMessageInfo

func (m *ErrorData) GetIdentifier() string {
	if m != nil && m.Identifier != nil {
		return *m.Identifier
	}
	return ""
}

func (m *ErrorData) GetCount() uint64 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func (m *ErrorData) GetLastUpdated() uint64 {
	if m != nil && m.LastUpdated != nil {
		return *m.LastUpdated
	}
	return 0
}

func (m *ErrorData) GetIsEnabled() bool {
	if m != nil && m.IsEnabled != nil {
		return *m.IsEnabled
	}
	return false
}

func (m *ErrorData) GetModuleId() uint32 {
	if m != nil && m.ModuleId != nil {
		return *m.ModuleId
	}
	return 0
}

func (m *ErrorData) GetComponentId() uint32 {
	if m != nil && m.ComponentId != nil {
		return *m.ComponentId
	}
	return 0
}

func (m *ErrorData) GetErrorCode() uint32 {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return 0
}

func (m *ErrorData) GetOccurCount() uint32 {
	if m != nil && m.OccurCount != nil {
		return *m.OccurCount
	}
	return 0
}

func (m *ErrorData) GetClearedCount() uint32 {
	if m != nil && m.ClearedCount != nil {
		return *m.ClearedCount
	}
	return 0
}

func (m *ErrorData) GetLastClearedAt() uint64 {
	if m != nil && m.LastClearedAt != nil {
		return *m.LastClearedAt
	}
	return 0
}

func (m *ErrorData) GetActionCount() uint32 {
	if m != nil && m.ActionCount != nil {
		return *m.ActionCount
	}
	return 0
}

func (m *ErrorData) GetLastActionTakenAt() uint64 {
	if m != nil && m.LastActionTakenAt != nil {
		return *m.LastActionTakenAt
	}
	return 0
}

func (m *ErrorData) GetFruType() string {
	if m != nil && m.FruType != nil {
		return *m.FruType
	}
	return ""
}

func (m *ErrorData) GetFruSlot() uint32 {
	if m != nil && m.FruSlot != nil {
		return *m.FruSlot
	}
	return 0
}

func (m *ErrorData) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *ErrorData) GetHelp() string {
	if m != nil && m.Help != nil {
		return *m.Help
	}
	return ""
}

type ErrorResourceInfo struct {
	// resource id e.g pfe identifier
	Id *uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// resource type e.g pfe/pic
	Type *string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	// resource state e.g enabled /disabled
	State                *uint32  `protobuf:"varint,3,opt,name=state" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorResourceInfo) Reset()         { *m = ErrorResourceInfo{} }
func (m *ErrorResourceInfo) String() string { return proto.CompactTextString(m) }
func (*ErrorResourceInfo) ProtoMessage()    {}
func (*ErrorResourceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e62e0d982e26602, []int{1}
}
func (m *ErrorResourceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorResourceInfo.Unmarshal(m, b)
}
func (m *ErrorResourceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorResourceInfo.Marshal(b, m, deterministic)
}
func (m *ErrorResourceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorResourceInfo.Merge(m, src)
}
func (m *ErrorResourceInfo) XXX_Size() int {
	return xxx_messageInfo_ErrorResourceInfo.Size(m)
}
func (m *ErrorResourceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorResourceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorResourceInfo proto.InternalMessageInfo

func (m *ErrorResourceInfo) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *ErrorResourceInfo) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *ErrorResourceInfo) GetState() uint32 {
	if m != nil && m.State != nil {
		return *m.State
	}
	return 0
}

//
// Top-level CmerrorData message
//
type CmerrorData struct {
	// collection of error items
	ErrorItem []*ErrorData `protobuf:"bytes,1,rep,name=error_item,json=errorItem" json:"error_item,omitempty"`
	// Fru slot identifier
	FruSlot *uint32 `protobuf:"varint,2,opt,name=fru_slot,json=fruSlot" json:"fru_slot,omitempty"`
	FruType *string `protobuf:"bytes,3,opt,name=fru_type,json=fruType" json:"fru_type,omitempty"`
	// collection of resource states from fru
	ResourceItem         []*ErrorResourceInfo `protobuf:"bytes,4,rep,name=resource_item,json=resourceItem" json:"resource_item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CmerrorData) Reset()         { *m = CmerrorData{} }
func (m *CmerrorData) String() string { return proto.CompactTextString(m) }
func (*CmerrorData) ProtoMessage()    {}
func (*CmerrorData) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e62e0d982e26602, []int{2}
}
func (m *CmerrorData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CmerrorData.Unmarshal(m, b)
}
func (m *CmerrorData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CmerrorData.Marshal(b, m, deterministic)
}
func (m *CmerrorData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CmerrorData.Merge(m, src)
}
func (m *CmerrorData) XXX_Size() int {
	return xxx_messageInfo_CmerrorData.Size(m)
}
func (m *CmerrorData) XXX_DiscardUnknown() {
	xxx_messageInfo_CmerrorData.DiscardUnknown(m)
}

var xxx_messageInfo_CmerrorData proto.InternalMessageInfo

func (m *CmerrorData) GetErrorItem() []*ErrorData {
	if m != nil {
		return m.ErrorItem
	}
	return nil
}

func (m *CmerrorData) GetFruSlot() uint32 {
	if m != nil && m.FruSlot != nil {
		return *m.FruSlot
	}
	return 0
}

func (m *CmerrorData) GetFruType() string {
	if m != nil && m.FruType != nil {
		return *m.FruType
	}
	return ""
}

func (m *CmerrorData) GetResourceItem() []*ErrorResourceInfo {
	if m != nil {
		return m.ResourceItem
	}
	return nil
}

var E_JnprCmerrorDataExt = &proto.ExtensionDesc{
	ExtendedType:  (*JuniperNetworksSensors)(nil),
	ExtensionType: (*CmerrorData)(nil),
	Field:         21,
	Name:          "jnpr_cmerror_data_ext",
	Tag:           "bytes,21,opt,name=jnpr_cmerror_data_ext",
	Filename:      "cmerror_data.proto",
}

func init() {
	proto.RegisterType((*ErrorData)(nil), "ErrorData")
	proto.RegisterType((*ErrorResourceInfo)(nil), "ErrorResourceInfo")
	proto.RegisterType((*CmerrorData)(nil), "CmerrorData")
	proto.RegisterExtension(E_JnprCmerrorDataExt)
}

func init() { proto.RegisterFile("cmerror_data.proto", fileDescriptor_9e62e0d982e26602) }

var fileDescriptor_9e62e0d982e26602 = []byte{
	// 543 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x93, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0x95, 0x74, 0x97, 0x6e, 0x26, 0x49, 0xff, 0x98, 0x56, 0x18, 0x2a, 0xa4, 0x50, 0x09,
	0x14, 0x90, 0xd8, 0x43, 0x0f, 0x20, 0x71, 0xa2, 0x2c, 0x3d, 0x2c, 0x12, 0x1c, 0xd2, 0x22, 0x71,
	0x8b, 0x4c, 0x3c, 0x2b, 0x4c, 0x93, 0x38, 0x72, 0x26, 0xa2, 0xbd, 0xf2, 0x32, 0xbc, 0x0a, 0x8f,
	0x85, 0x6c, 0x67, 0xab, 0x54, 0xdc, 0xb2, 0xf3, 0xfb, 0xe6, 0xb3, 0xbf, 0xf1, 0x2c, 0xb0, 0xaa,
	0x41, 0x63, 0xb4, 0x29, 0xa5, 0x20, 0xb1, 0xec, 0x8c, 0x26, 0xfd, 0xe4, 0x21, 0x61, 0x8d, 0x0d,
	0x92, 0xb9, 0x2d, 0x49, 0x77, 0xbe, 0x78, 0xfa, 0x77, 0x06, 0xd1, 0x85, 0x55, 0x7e, 0x14, 0x24,
	0xd8, 0x73, 0x00, 0x25, 0xb1, 0x25, 0xb5, 0x51, 0x68, 0x78, 0x90, 0x05, 0x79, 0xf4, 0x61, 0xfe,
	0xfb, 0x7d, 0xb8, 0x08, 0x8a, 0x09, 0x60, 0x27, 0x30, 0xaf, 0xf4, 0xd0, 0x12, 0x0f, 0xb3, 0x20,
	0x9f, 0x39, 0x05, 0x0f, 0x0a, 0x5f, 0x63, 0x39, 0x24, 0xb5, 0xe8, 0xa9, 0x1c, 0x3a, 0x29, 0x08,
	0x25, 0xdf, 0xb9, 0xd3, 0x1c, 0x04, 0x45, 0x6c, 0xd1, 0x57, 0x4f, 0xd8, 0x53, 0x00, 0xd5, 0x97,
	0xd8, 0x8a, 0xef, 0x35, 0x4a, 0x3e, 0xcb, 0x82, 0x7c, 0x51, 0x44, 0xaa, 0xbf, 0xf0, 0x05, 0x76,
	0x02, 0x51, 0xa3, 0xe5, 0x50, 0x63, 0xa9, 0x24, 0x9f, 0x67, 0x41, 0x9e, 0x16, 0x0b, 0x5f, 0x58,
	0x4b, 0xf6, 0x0c, 0x92, 0x4a, 0x37, 0x9d, 0x6e, 0xb1, 0x25, 0xcb, 0x1f, 0x38, 0x1e, 0xdf, 0xd5,
	0xd6, 0xce, 0xde, 0xcf, 0xa0, 0xd2, 0x12, 0xf9, 0xae, 0x13, 0x44, 0xae, 0xb2, 0xd2, 0x12, 0xd9,
	0x0b, 0x88, 0x75, 0x55, 0x0d, 0x16, 0xdb, 0x28, 0x0b, 0xcb, 0xb7, 0x51, 0xc0, 0x91, 0x95, 0xcb,
	0xf3, 0x0a, 0xd2, 0xaa, 0x46, 0x61, 0x50, 0x8e, 0xca, 0x68, 0xaa, 0x4c, 0x46, 0xe6, 0xb5, 0xaf,
	0x61, 0xdf, 0x65, 0xdf, 0x36, 0x08, 0xe2, 0x30, 0x8d, 0x9f, 0x5a, 0xba, 0xf2, 0xf0, 0xdc, 0x8d,
	0x4a, 0x54, 0xa4, 0x74, 0x3b, 0x3a, 0xc7, 0x53, 0xe7, 0xd8, 0x23, 0x6f, 0xfc, 0x06, 0x8e, 0x9c,
	0xf1, 0x28, 0x27, 0x71, 0x8d, 0xad, 0x75, 0x4f, 0xa6, 0xee, 0x87, 0x56, 0x72, 0xee, 0x14, 0x57,
	0x56, 0x70, 0x4e, 0xec, 0x31, 0x2c, 0x36, 0x66, 0x28, 0xe9, 0xb6, 0x43, 0x9e, 0xda, 0xe7, 0x2c,
	0x76, 0x37, 0x66, 0xb8, 0xba, 0xed, 0x70, 0x8b, 0xfa, 0x5a, 0x13, 0xdf, 0x73, 0xc3, 0xb1, 0xe8,
	0xb2, 0xd6, 0xc4, 0x32, 0x88, 0x25, 0xf6, 0x95, 0x51, 0x9d, 0xf5, 0xe2, 0xfb, 0xae, 0x71, 0x5a,
	0x62, 0x0c, 0x66, 0x3f, 0xb0, 0xee, 0xf8, 0x81, 0x43, 0xee, 0xfb, 0xf4, 0x33, 0x1c, 0xba, 0x4d,
	0x2a, 0xb0, 0xd7, 0x83, 0xa9, 0x70, 0xdd, 0x6e, 0x34, 0xdb, 0x83, 0x50, 0x49, 0xb7, 0x49, 0x69,
	0x11, 0x2a, 0x69, 0x1b, 0xdd, 0x65, 0x42, 0xdf, 0x68, 0xbf, 0xd9, 0x11, 0xcc, 0x7b, 0x12, 0x84,
	0x6e, 0x55, 0xd2, 0xc2, 0xff, 0x38, 0xfd, 0x13, 0x40, 0xbc, 0xf2, 0x5b, 0xec, 0x76, 0xf3, 0xe5,
	0xf6, 0x39, 0x15, 0x61, 0xc3, 0x83, 0x6c, 0x27, 0x8f, 0xcf, 0x60, 0x79, 0xb7, 0xbb, 0xe3, 0xd3,
	0xae, 0x09, 0x9b, 0x7b, 0xd1, 0xc2, 0xfb, 0xd1, 0xa6, 0x03, 0xd9, 0xb9, 0x3f, 0x90, 0xb7, 0x90,
	0x9a, 0xf1, 0xea, 0xfe, 0x8c, 0x99, 0x3b, 0x83, 0x2d, 0xff, 0x4b, 0x55, 0x24, 0x5b, 0xa1, 0x3d,
	0xee, 0xdd, 0x37, 0x38, 0xfe, 0xd9, 0x76, 0xa6, 0x9c, 0xfe, 0xe7, 0x4a, 0xbc, 0x21, 0xf6, 0x68,
	0xf9, 0x69, 0x68, 0x55, 0x87, 0xe6, 0x0b, 0xd2, 0x2f, 0x6d, 0xae, 0xfb, 0x4b, 0x6c, 0x7b, 0x6d,
	0x7a, 0x7e, 0x9c, 0x05, 0x79, 0x7c, 0x96, 0x2c, 0x27, 0xf9, 0x0a, 0x66, 0x3d, 0x26, 0x85, 0x8b,
	0x1b, 0xfa, 0x17, 0x00, 0x00, 0xff, 0xff, 0x55, 0xfa, 0x54, 0x59, 0xc7, 0x03, 0x00, 0x00,
}
