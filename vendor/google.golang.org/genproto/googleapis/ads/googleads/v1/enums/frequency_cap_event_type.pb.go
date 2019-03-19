// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/frequency_cap_event_type.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The type of event that the cap applies to (e.g. impression).
type FrequencyCapEventTypeEnum_FrequencyCapEventType int32

const (
	// Not specified.
	FrequencyCapEventTypeEnum_UNSPECIFIED FrequencyCapEventTypeEnum_FrequencyCapEventType = 0
	// Used for return value only. Represents value unknown in this version.
	FrequencyCapEventTypeEnum_UNKNOWN FrequencyCapEventTypeEnum_FrequencyCapEventType = 1
	// The cap applies on ad impressions.
	FrequencyCapEventTypeEnum_IMPRESSION FrequencyCapEventTypeEnum_FrequencyCapEventType = 2
	// The cap applies on video ad views.
	FrequencyCapEventTypeEnum_VIDEO_VIEW FrequencyCapEventTypeEnum_FrequencyCapEventType = 3
)

var FrequencyCapEventTypeEnum_FrequencyCapEventType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "IMPRESSION",
	3: "VIDEO_VIEW",
}
var FrequencyCapEventTypeEnum_FrequencyCapEventType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"IMPRESSION":  2,
	"VIDEO_VIEW":  3,
}

func (x FrequencyCapEventTypeEnum_FrequencyCapEventType) String() string {
	return proto.EnumName(FrequencyCapEventTypeEnum_FrequencyCapEventType_name, int32(x))
}
func (FrequencyCapEventTypeEnum_FrequencyCapEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_frequency_cap_event_type_d5459adaf007f3b3, []int{0, 0}
}

// Container for enum describing the type of event that the cap applies to.
type FrequencyCapEventTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrequencyCapEventTypeEnum) Reset()         { *m = FrequencyCapEventTypeEnum{} }
func (m *FrequencyCapEventTypeEnum) String() string { return proto.CompactTextString(m) }
func (*FrequencyCapEventTypeEnum) ProtoMessage()    {}
func (*FrequencyCapEventTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_frequency_cap_event_type_d5459adaf007f3b3, []int{0}
}
func (m *FrequencyCapEventTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrequencyCapEventTypeEnum.Unmarshal(m, b)
}
func (m *FrequencyCapEventTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrequencyCapEventTypeEnum.Marshal(b, m, deterministic)
}
func (dst *FrequencyCapEventTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrequencyCapEventTypeEnum.Merge(dst, src)
}
func (m *FrequencyCapEventTypeEnum) XXX_Size() int {
	return xxx_messageInfo_FrequencyCapEventTypeEnum.Size(m)
}
func (m *FrequencyCapEventTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FrequencyCapEventTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FrequencyCapEventTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FrequencyCapEventTypeEnum)(nil), "google.ads.googleads.v1.enums.FrequencyCapEventTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.FrequencyCapEventTypeEnum_FrequencyCapEventType", FrequencyCapEventTypeEnum_FrequencyCapEventType_name, FrequencyCapEventTypeEnum_FrequencyCapEventType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/frequency_cap_event_type.proto", fileDescriptor_frequency_cap_event_type_d5459adaf007f3b3)
}

var fileDescriptor_frequency_cap_event_type_d5459adaf007f3b3 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x75, 0x1d, 0x28, 0x64, 0xa0, 0xa5, 0xe0, 0xc1, 0xe1, 0x0e, 0xdb, 0x0f, 0x48, 0x29, 0xde,
	0xa2, 0x97, 0x6e, 0xcb, 0x46, 0x11, 0xbb, 0xe2, 0x5c, 0x07, 0x52, 0x28, 0x71, 0x8d, 0x61, 0xb0,
	0x25, 0xb1, 0xe9, 0x06, 0xfb, 0x3b, 0x1e, 0xfd, 0x29, 0xfe, 0x14, 0x2f, 0xfe, 0x05, 0x49, 0x62,
	0x7b, 0x9a, 0x5e, 0xc2, 0x4b, 0xde, 0xf7, 0x5e, 0xde, 0xf7, 0xc0, 0x1d, 0x13, 0x82, 0x6d, 0xa8,
	0x4f, 0x0a, 0xe5, 0x5b, 0xa8, 0xd1, 0x3e, 0xf0, 0x29, 0xdf, 0x6d, 0x95, 0xff, 0x5a, 0xd2, 0xb7,
	0x1d, 0xe5, 0xab, 0x43, 0xbe, 0x22, 0x32, 0xa7, 0x7b, 0xca, 0xab, 0xbc, 0x3a, 0x48, 0x0a, 0x65,
	0x29, 0x2a, 0xe1, 0xf5, 0xac, 0x04, 0x92, 0x42, 0xc1, 0x46, 0x0d, 0xf7, 0x01, 0x34, 0xea, 0xee,
	0x75, 0x6d, 0x2e, 0xd7, 0x3e, 0xe1, 0x5c, 0x54, 0xa4, 0x5a, 0x0b, 0xae, 0xac, 0x78, 0x50, 0x82,
	0xab, 0x49, 0x6d, 0x3f, 0x22, 0x12, 0x6b, 0xf3, 0xa7, 0x83, 0xa4, 0x98, 0xef, 0xb6, 0x83, 0x05,
	0xb8, 0x3c, 0x4a, 0x7a, 0x17, 0xa0, 0xb3, 0x88, 0xe7, 0x09, 0x1e, 0x45, 0x93, 0x08, 0x8f, 0xdd,
	0x13, 0xaf, 0x03, 0xce, 0x16, 0xf1, 0x7d, 0x3c, 0x5b, 0xc6, 0x6e, 0xcb, 0x3b, 0x07, 0x20, 0x7a,
	0x48, 0x1e, 0xf1, 0x7c, 0x1e, 0xcd, 0x62, 0xd7, 0xd1, 0xf7, 0x34, 0x1a, 0xe3, 0x59, 0x9e, 0x46,
	0x78, 0xe9, 0xb6, 0x87, 0xdf, 0x2d, 0xd0, 0x5f, 0x89, 0x2d, 0xfc, 0x37, 0xf7, 0xb0, 0x7b, 0xf4,
	0xeb, 0x44, 0xa7, 0x4e, 0x5a, 0xcf, 0xc3, 0x5f, 0x31, 0x13, 0x1b, 0xc2, 0x19, 0x14, 0x25, 0xf3,
	0x19, 0xe5, 0x66, 0xa7, 0xba, 0x42, 0xb9, 0x56, 0x7f, 0x34, 0x7a, 0x6b, 0xce, 0x77, 0xa7, 0x3d,
	0x0d, 0xc3, 0x0f, 0xa7, 0x37, 0xb5, 0x56, 0x61, 0xa1, 0xa0, 0x85, 0x1a, 0xa5, 0x01, 0xd4, 0x1d,
	0xa8, 0xcf, 0x9a, 0xcf, 0xc2, 0x42, 0x65, 0x0d, 0x9f, 0xa5, 0x41, 0x66, 0xf8, 0x2f, 0xa7, 0x6f,
	0x1f, 0x11, 0x0a, 0x0b, 0x85, 0x50, 0x33, 0x81, 0x50, 0x1a, 0x20, 0x64, 0x66, 0x5e, 0x4e, 0x4d,
	0xb0, 0x9b, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x76, 0x40, 0x80, 0xe2, 0xe9, 0x01, 0x00, 0x00,
}
