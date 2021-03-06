// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/advertising_channel_sub_type.proto

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

// Enum describing the different channel subtypes.
type AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType int32

const (
	// Not specified.
	AdvertisingChannelSubTypeEnum_UNSPECIFIED AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 0
	// Used as a return value only. Represents value unknown in this version.
	AdvertisingChannelSubTypeEnum_UNKNOWN AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 1
	// Mobile app campaigns for Search.
	AdvertisingChannelSubTypeEnum_SEARCH_MOBILE_APP AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 2
	// Mobile app campaigns for Display.
	AdvertisingChannelSubTypeEnum_DISPLAY_MOBILE_APP AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 3
	// AdWords express campaigns for search.
	AdvertisingChannelSubTypeEnum_SEARCH_EXPRESS AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 4
	// AdWords Express campaigns for display.
	AdvertisingChannelSubTypeEnum_DISPLAY_EXPRESS AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 5
	// Smart Shopping campaigns.
	AdvertisingChannelSubTypeEnum_SHOPPING_SMART_ADS AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 6
	// Gmail Ad campaigns.
	AdvertisingChannelSubTypeEnum_DISPLAY_GMAIL_AD AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 7
	// Smart display campaigns.
	AdvertisingChannelSubTypeEnum_DISPLAY_SMART_CAMPAIGN AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 8
	// Video Outstream campaigns.
	AdvertisingChannelSubTypeEnum_VIDEO_OUTSTREAM AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 9
	// Video TrueView for Action campaigns.
	AdvertisingChannelSubTypeEnum_VIDEO_ACTION AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 10
	// Video campaigns with non-skippable video ads.
	AdvertisingChannelSubTypeEnum_VIDEO_NON_SKIPPABLE AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 11
)

var AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "SEARCH_MOBILE_APP",
	3:  "DISPLAY_MOBILE_APP",
	4:  "SEARCH_EXPRESS",
	5:  "DISPLAY_EXPRESS",
	6:  "SHOPPING_SMART_ADS",
	7:  "DISPLAY_GMAIL_AD",
	8:  "DISPLAY_SMART_CAMPAIGN",
	9:  "VIDEO_OUTSTREAM",
	10: "VIDEO_ACTION",
	11: "VIDEO_NON_SKIPPABLE",
}
var AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType_value = map[string]int32{
	"UNSPECIFIED":            0,
	"UNKNOWN":                1,
	"SEARCH_MOBILE_APP":      2,
	"DISPLAY_MOBILE_APP":     3,
	"SEARCH_EXPRESS":         4,
	"DISPLAY_EXPRESS":        5,
	"SHOPPING_SMART_ADS":     6,
	"DISPLAY_GMAIL_AD":       7,
	"DISPLAY_SMART_CAMPAIGN": 8,
	"VIDEO_OUTSTREAM":        9,
	"VIDEO_ACTION":           10,
	"VIDEO_NON_SKIPPABLE":    11,
}

func (x AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType) String() string {
	return proto.EnumName(AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType_name, int32(x))
}
func (AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_advertising_channel_sub_type_2d2d5e9ae88f8a3a, []int{0, 0}
}

// An immutable specialization of an Advertising Channel.
type AdvertisingChannelSubTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdvertisingChannelSubTypeEnum) Reset()         { *m = AdvertisingChannelSubTypeEnum{} }
func (m *AdvertisingChannelSubTypeEnum) String() string { return proto.CompactTextString(m) }
func (*AdvertisingChannelSubTypeEnum) ProtoMessage()    {}
func (*AdvertisingChannelSubTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertising_channel_sub_type_2d2d5e9ae88f8a3a, []int{0}
}
func (m *AdvertisingChannelSubTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdvertisingChannelSubTypeEnum.Unmarshal(m, b)
}
func (m *AdvertisingChannelSubTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdvertisingChannelSubTypeEnum.Marshal(b, m, deterministic)
}
func (dst *AdvertisingChannelSubTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdvertisingChannelSubTypeEnum.Merge(dst, src)
}
func (m *AdvertisingChannelSubTypeEnum) XXX_Size() int {
	return xxx_messageInfo_AdvertisingChannelSubTypeEnum.Size(m)
}
func (m *AdvertisingChannelSubTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdvertisingChannelSubTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdvertisingChannelSubTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AdvertisingChannelSubTypeEnum)(nil), "google.ads.googleads.v1.enums.AdvertisingChannelSubTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType", AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType_name, AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/advertising_channel_sub_type.proto", fileDescriptor_advertising_channel_sub_type_2d2d5e9ae88f8a3a)
}

var fileDescriptor_advertising_channel_sub_type_2d2d5e9ae88f8a3a = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdf, 0x6e, 0xd3, 0x30,
	0x18, 0xc5, 0x69, 0x06, 0x1b, 0xb8, 0x88, 0x19, 0x0f, 0x86, 0x98, 0x28, 0xd2, 0xf6, 0x00, 0x89,
	0x22, 0xee, 0xc2, 0x0d, 0x4e, 0x62, 0x32, 0x6b, 0x4d, 0x62, 0xd5, 0x69, 0xf9, 0xa3, 0x4a, 0x56,
	0xba, 0x44, 0x21, 0x52, 0xeb, 0x44, 0x75, 0x5a, 0x69, 0xcf, 0xc2, 0x1d, 0x97, 0xbc, 0x03, 0x2f,
	0xc0, 0xa3, 0x70, 0xcb, 0x0b, 0xa0, 0xc4, 0x4d, 0xe1, 0xa6, 0xdc, 0x44, 0x47, 0xe7, 0xfb, 0xf9,
	0x7c, 0x91, 0x8f, 0xc1, 0xbb, 0xa2, 0xaa, 0x8a, 0x65, 0x6e, 0xa5, 0x99, 0xb2, 0xb4, 0x6c, 0xd5,
	0xd6, 0xb6, 0x72, 0xb9, 0x59, 0x29, 0x2b, 0xcd, 0xb6, 0xf9, 0xba, 0x29, 0x55, 0x29, 0x0b, 0x71,
	0xfb, 0x25, 0x95, 0x32, 0x5f, 0x0a, 0xb5, 0x59, 0x88, 0xe6, 0xae, 0xce, 0xcd, 0x7a, 0x5d, 0x35,
	0x15, 0x1a, 0xe9, 0x63, 0x66, 0x9a, 0x29, 0x73, 0x9f, 0x60, 0x6e, 0x6d, 0xb3, 0x4b, 0xb8, 0x78,
	0xd5, 0x2f, 0xa8, 0x4b, 0x2b, 0x95, 0xb2, 0x6a, 0xd2, 0xa6, 0xac, 0xa4, 0xd2, 0x87, 0xaf, 0x7e,
	0x18, 0x60, 0x84, 0xff, 0xee, 0xf0, 0xf4, 0x0a, 0xbe, 0x59, 0x24, 0x77, 0x75, 0x4e, 0xe4, 0x66,
	0x75, 0xf5, 0xd5, 0x00, 0x2f, 0x0f, 0x12, 0xe8, 0x14, 0x0c, 0xa7, 0x11, 0x67, 0xc4, 0xa3, 0xef,
	0x29, 0xf1, 0xe1, 0x3d, 0x34, 0x04, 0x27, 0xd3, 0xe8, 0x26, 0x8a, 0x3f, 0x44, 0x70, 0x80, 0x9e,
	0x83, 0xa7, 0x9c, 0xe0, 0x89, 0x77, 0x2d, 0xc2, 0xd8, 0xa5, 0x63, 0x22, 0x30, 0x63, 0xd0, 0x40,
	0xe7, 0x00, 0xf9, 0x94, 0xb3, 0x31, 0xfe, 0xf4, 0xaf, 0x7f, 0x84, 0x10, 0x78, 0xb2, 0xc3, 0xc9,
	0x47, 0x36, 0x21, 0x9c, 0xc3, 0xfb, 0xe8, 0x0c, 0x9c, 0xf6, 0x6c, 0x6f, 0x3e, 0x68, 0x03, 0xf8,
	0x75, 0xcc, 0x18, 0x8d, 0x02, 0xc1, 0x43, 0x3c, 0x49, 0x04, 0xf6, 0x39, 0x3c, 0x46, 0xcf, 0x00,
	0xec, 0xe1, 0x20, 0xc4, 0x74, 0x2c, 0xb0, 0x0f, 0x4f, 0xd0, 0x05, 0x38, 0xef, 0x5d, 0x0d, 0x7b,
	0x38, 0x64, 0x98, 0x06, 0x11, 0x7c, 0xd8, 0xc6, 0xcf, 0xa8, 0x4f, 0x62, 0x11, 0x4f, 0x13, 0x9e,
	0x4c, 0x08, 0x0e, 0xe1, 0x23, 0x04, 0xc1, 0x63, 0x6d, 0x62, 0x2f, 0xa1, 0x71, 0x04, 0x01, 0x7a,
	0x01, 0xce, 0xb4, 0x13, 0xc5, 0x91, 0xe0, 0x37, 0x94, 0x31, 0xec, 0x8e, 0x09, 0x1c, 0xba, 0xbf,
	0x07, 0xe0, 0xf2, 0xb6, 0x5a, 0x99, 0xff, 0xed, 0xc0, 0x7d, 0x7d, 0xf0, 0x02, 0x59, 0xdb, 0x02,
	0x1b, 0x7c, 0x76, 0x77, 0x01, 0x45, 0xb5, 0x4c, 0x65, 0x61, 0x56, 0xeb, 0xc2, 0x2a, 0x72, 0xd9,
	0x75, 0xd4, 0x3f, 0x8b, 0xba, 0x54, 0x07, 0x5e, 0xc9, 0xdb, 0xee, 0xfb, 0xcd, 0x38, 0x0a, 0x30,
	0xfe, 0x6e, 0x8c, 0x02, 0x1d, 0x85, 0x33, 0x65, 0x6a, 0xd9, 0xaa, 0x99, 0x6d, 0xb6, 0x75, 0xaa,
	0x9f, 0xfd, 0x7c, 0x8e, 0x33, 0x35, 0xdf, 0xcf, 0xe7, 0x33, 0x7b, 0xde, 0xcd, 0x7f, 0x19, 0x97,
	0xda, 0x74, 0x1c, 0x9c, 0x29, 0xc7, 0xd9, 0x13, 0x8e, 0x33, 0xb3, 0x1d, 0xa7, 0x63, 0x16, 0xc7,
	0xdd, 0x8f, 0xbd, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x3a, 0x45, 0xb1, 0xbd, 0x02, 0x00,
	0x00,
}
