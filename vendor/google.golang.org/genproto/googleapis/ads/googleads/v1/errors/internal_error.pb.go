// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/internal_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v1/errors"

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

// Enum describing possible internal errors.
type InternalErrorEnum_InternalError int32

const (
	// Enum unspecified.
	InternalErrorEnum_UNSPECIFIED InternalErrorEnum_InternalError = 0
	// The received error code is not known in this version.
	InternalErrorEnum_UNKNOWN InternalErrorEnum_InternalError = 1
	// Google Ads API encountered unexpected internal error.
	InternalErrorEnum_INTERNAL_ERROR InternalErrorEnum_InternalError = 2
	// The intended error code doesn't exist in any API version. This will be
	// fixed by adding a new error code as soon as possible.
	InternalErrorEnum_ERROR_CODE_NOT_PUBLISHED InternalErrorEnum_InternalError = 3
	// Google Ads API encountered an unexpected transient error. The user
	// should retry their request in these cases.
	InternalErrorEnum_TRANSIENT_ERROR InternalErrorEnum_InternalError = 4
)

var InternalErrorEnum_InternalError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "INTERNAL_ERROR",
	3: "ERROR_CODE_NOT_PUBLISHED",
	4: "TRANSIENT_ERROR",
}
var InternalErrorEnum_InternalError_value = map[string]int32{
	"UNSPECIFIED":              0,
	"UNKNOWN":                  1,
	"INTERNAL_ERROR":           2,
	"ERROR_CODE_NOT_PUBLISHED": 3,
	"TRANSIENT_ERROR":          4,
}

func (x InternalErrorEnum_InternalError) String() string {
	return proto.EnumName(InternalErrorEnum_InternalError_name, int32(x))
}
func (InternalErrorEnum_InternalError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_internal_error_67cb6d6df548b15b, []int{0, 0}
}

// Container for enum describing possible internal errors.
type InternalErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InternalErrorEnum) Reset()         { *m = InternalErrorEnum{} }
func (m *InternalErrorEnum) String() string { return proto.CompactTextString(m) }
func (*InternalErrorEnum) ProtoMessage()    {}
func (*InternalErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_internal_error_67cb6d6df548b15b, []int{0}
}
func (m *InternalErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InternalErrorEnum.Unmarshal(m, b)
}
func (m *InternalErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InternalErrorEnum.Marshal(b, m, deterministic)
}
func (dst *InternalErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalErrorEnum.Merge(dst, src)
}
func (m *InternalErrorEnum) XXX_Size() int {
	return xxx_messageInfo_InternalErrorEnum.Size(m)
}
func (m *InternalErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_InternalErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*InternalErrorEnum)(nil), "google.ads.googleads.v1.errors.InternalErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v1.errors.InternalErrorEnum_InternalError", InternalErrorEnum_InternalError_name, InternalErrorEnum_InternalError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/internal_error.proto", fileDescriptor_internal_error_67cb6d6df548b15b)
}

var fileDescriptor_internal_error_67cb6d6df548b15b = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x86, 0x6d, 0x27, 0x0a, 0x19, 0xba, 0x1a, 0x6f, 0x44, 0xc6, 0x2e, 0xfa, 0x00, 0x29, 0x65,
	0x77, 0xf1, 0x2a, 0x5b, 0xe3, 0x2c, 0x8e, 0xb4, 0x74, 0xdd, 0x04, 0x29, 0x94, 0x6a, 0x4b, 0x29,
	0x6c, 0xc9, 0x48, 0xea, 0xde, 0xc1, 0xd7, 0xf0, 0xd2, 0x47, 0xf1, 0x51, 0x04, 0xdf, 0x41, 0xda,
	0x6c, 0x85, 0x5d, 0xe8, 0x55, 0xff, 0x1e, 0xbe, 0xef, 0xe4, 0x9c, 0x03, 0xc6, 0xa5, 0x10, 0xe5,
	0xba, 0x70, 0xb2, 0x5c, 0x39, 0x3a, 0x36, 0x69, 0xe7, 0x3a, 0x85, 0x94, 0x42, 0x2a, 0xa7, 0xe2,
	0x75, 0x21, 0x79, 0xb6, 0x4e, 0xdb, 0x7f, 0xb4, 0x95, 0xa2, 0x16, 0x70, 0xa4, 0x49, 0x94, 0xe5,
	0x0a, 0x75, 0x12, 0xda, 0xb9, 0x48, 0x4b, 0xb7, 0xc3, 0x43, 0xd3, 0x6d, 0xe5, 0x64, 0x9c, 0x8b,
	0x3a, 0xab, 0x2b, 0xc1, 0x95, 0xb6, 0xed, 0x77, 0x03, 0x5c, 0xf9, 0xfb, 0xb6, 0xb4, 0x11, 0x28,
	0x7f, 0xdb, 0xd8, 0x35, 0xb8, 0x38, 0x2a, 0xc2, 0x01, 0xe8, 0x2f, 0xd9, 0x22, 0xa4, 0x53, 0xff,
	0xde, 0xa7, 0x9e, 0x75, 0x02, 0xfb, 0xe0, 0x7c, 0xc9, 0x1e, 0x59, 0xf0, 0xc4, 0x2c, 0x03, 0x42,
	0x70, 0xe9, 0xb3, 0x98, 0x46, 0x8c, 0xcc, 0x53, 0x1a, 0x45, 0x41, 0x64, 0x99, 0x70, 0x08, 0x6e,
	0xda, 0x98, 0x4e, 0x03, 0x8f, 0xa6, 0x2c, 0x88, 0xd3, 0x70, 0x39, 0x99, 0xfb, 0x8b, 0x07, 0xea,
	0x59, 0x3d, 0x78, 0x0d, 0x06, 0x71, 0x44, 0xd8, 0xc2, 0xa7, 0x2c, 0xde, 0x2b, 0xa7, 0x93, 0x1f,
	0x03, 0xd8, 0xaf, 0x62, 0x83, 0xfe, 0x5f, 0x68, 0x02, 0x8f, 0x46, 0x0b, 0x9b, 0x35, 0x42, 0xe3,
	0xd9, 0xdb, 0x5b, 0xa5, 0x58, 0x67, 0xbc, 0x44, 0x42, 0x96, 0x4e, 0x59, 0xf0, 0x76, 0xc9, 0xc3,
	0x2d, 0xb7, 0x95, 0xfa, 0xeb, 0xb4, 0x77, 0xfa, 0xf3, 0x61, 0xf6, 0x66, 0x84, 0x7c, 0x9a, 0xa3,
	0x99, 0x6e, 0x46, 0x72, 0x85, 0x74, 0x6c, 0xd2, 0xca, 0x45, 0xed, 0x93, 0xea, 0xeb, 0x00, 0x24,
	0x24, 0x57, 0x49, 0x07, 0x24, 0x2b, 0x37, 0xd1, 0xc0, 0xb7, 0x69, 0xeb, 0x2a, 0xc6, 0x24, 0x57,
	0x18, 0x77, 0x08, 0xc6, 0x2b, 0x17, 0x63, 0x0d, 0xbd, 0x9c, 0xb5, 0xd3, 0x8d, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xef, 0xf8, 0x2e, 0x2f, 0xf7, 0x01, 0x00, 0x00,
}
