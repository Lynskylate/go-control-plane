// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/ratelimit/v3alpha/rls.proto

package envoy_service_ratelimit_v3alpha

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3alpha1 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3alpha"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/extensions/common/ratelimit/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RateLimitResponse_Code int32

const (
	RateLimitResponse_UNKNOWN    RateLimitResponse_Code = 0
	RateLimitResponse_OK         RateLimitResponse_Code = 1
	RateLimitResponse_OVER_LIMIT RateLimitResponse_Code = 2
)

var RateLimitResponse_Code_name = map[int32]string{
	0: "UNKNOWN",
	1: "OK",
	2: "OVER_LIMIT",
}

var RateLimitResponse_Code_value = map[string]int32{
	"UNKNOWN":    0,
	"OK":         1,
	"OVER_LIMIT": 2,
}

func (x RateLimitResponse_Code) String() string {
	return proto.EnumName(RateLimitResponse_Code_name, int32(x))
}

func (RateLimitResponse_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5929a49ed6c3ed07, []int{1, 0}
}

type RateLimitResponse_RateLimit_Unit int32

const (
	RateLimitResponse_RateLimit_UNKNOWN RateLimitResponse_RateLimit_Unit = 0
	RateLimitResponse_RateLimit_SECOND  RateLimitResponse_RateLimit_Unit = 1
	RateLimitResponse_RateLimit_MINUTE  RateLimitResponse_RateLimit_Unit = 2
	RateLimitResponse_RateLimit_HOUR    RateLimitResponse_RateLimit_Unit = 3
	RateLimitResponse_RateLimit_DAY     RateLimitResponse_RateLimit_Unit = 4
)

var RateLimitResponse_RateLimit_Unit_name = map[int32]string{
	0: "UNKNOWN",
	1: "SECOND",
	2: "MINUTE",
	3: "HOUR",
	4: "DAY",
}

var RateLimitResponse_RateLimit_Unit_value = map[string]int32{
	"UNKNOWN": 0,
	"SECOND":  1,
	"MINUTE":  2,
	"HOUR":    3,
	"DAY":     4,
}

func (x RateLimitResponse_RateLimit_Unit) String() string {
	return proto.EnumName(RateLimitResponse_RateLimit_Unit_name, int32(x))
}

func (RateLimitResponse_RateLimit_Unit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5929a49ed6c3ed07, []int{1, 0, 0}
}

type RateLimitRequest struct {
	Domain               string                         `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Descriptors          []*v3alpha.RateLimitDescriptor `protobuf:"bytes,2,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
	HitsAddend           uint32                         `protobuf:"varint,3,opt,name=hits_addend,json=hitsAddend,proto3" json:"hits_addend,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *RateLimitRequest) Reset()         { *m = RateLimitRequest{} }
func (m *RateLimitRequest) String() string { return proto.CompactTextString(m) }
func (*RateLimitRequest) ProtoMessage()    {}
func (*RateLimitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5929a49ed6c3ed07, []int{0}
}

func (m *RateLimitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitRequest.Unmarshal(m, b)
}
func (m *RateLimitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitRequest.Marshal(b, m, deterministic)
}
func (m *RateLimitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitRequest.Merge(m, src)
}
func (m *RateLimitRequest) XXX_Size() int {
	return xxx_messageInfo_RateLimitRequest.Size(m)
}
func (m *RateLimitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitRequest proto.InternalMessageInfo

func (m *RateLimitRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimitRequest) GetDescriptors() []*v3alpha.RateLimitDescriptor {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

func (m *RateLimitRequest) GetHitsAddend() uint32 {
	if m != nil {
		return m.HitsAddend
	}
	return 0
}

type RateLimitResponse struct {
	OverallCode          RateLimitResponse_Code                `protobuf:"varint,1,opt,name=overall_code,json=overallCode,proto3,enum=envoy.service.ratelimit.v3alpha.RateLimitResponse_Code" json:"overall_code,omitempty"`
	Statuses             []*RateLimitResponse_DescriptorStatus `protobuf:"bytes,2,rep,name=statuses,proto3" json:"statuses,omitempty"`
	ResponseHeadersToAdd []*v3alpha1.HeaderValue               `protobuf:"bytes,3,rep,name=response_headers_to_add,json=responseHeadersToAdd,proto3" json:"response_headers_to_add,omitempty"`
	RequestHeadersToAdd  []*v3alpha1.HeaderValue               `protobuf:"bytes,4,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *RateLimitResponse) Reset()         { *m = RateLimitResponse{} }
func (m *RateLimitResponse) String() string { return proto.CompactTextString(m) }
func (*RateLimitResponse) ProtoMessage()    {}
func (*RateLimitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5929a49ed6c3ed07, []int{1}
}

func (m *RateLimitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitResponse.Unmarshal(m, b)
}
func (m *RateLimitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitResponse.Marshal(b, m, deterministic)
}
func (m *RateLimitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitResponse.Merge(m, src)
}
func (m *RateLimitResponse) XXX_Size() int {
	return xxx_messageInfo_RateLimitResponse.Size(m)
}
func (m *RateLimitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitResponse proto.InternalMessageInfo

func (m *RateLimitResponse) GetOverallCode() RateLimitResponse_Code {
	if m != nil {
		return m.OverallCode
	}
	return RateLimitResponse_UNKNOWN
}

func (m *RateLimitResponse) GetStatuses() []*RateLimitResponse_DescriptorStatus {
	if m != nil {
		return m.Statuses
	}
	return nil
}

func (m *RateLimitResponse) GetResponseHeadersToAdd() []*v3alpha1.HeaderValue {
	if m != nil {
		return m.ResponseHeadersToAdd
	}
	return nil
}

func (m *RateLimitResponse) GetRequestHeadersToAdd() []*v3alpha1.HeaderValue {
	if m != nil {
		return m.RequestHeadersToAdd
	}
	return nil
}

type RateLimitResponse_RateLimit struct {
	RequestsPerUnit      uint32                           `protobuf:"varint,1,opt,name=requests_per_unit,json=requestsPerUnit,proto3" json:"requests_per_unit,omitempty"`
	Unit                 RateLimitResponse_RateLimit_Unit `protobuf:"varint,2,opt,name=unit,proto3,enum=envoy.service.ratelimit.v3alpha.RateLimitResponse_RateLimit_Unit" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *RateLimitResponse_RateLimit) Reset()         { *m = RateLimitResponse_RateLimit{} }
func (m *RateLimitResponse_RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimitResponse_RateLimit) ProtoMessage()    {}
func (*RateLimitResponse_RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_5929a49ed6c3ed07, []int{1, 0}
}

func (m *RateLimitResponse_RateLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitResponse_RateLimit.Unmarshal(m, b)
}
func (m *RateLimitResponse_RateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitResponse_RateLimit.Marshal(b, m, deterministic)
}
func (m *RateLimitResponse_RateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitResponse_RateLimit.Merge(m, src)
}
func (m *RateLimitResponse_RateLimit) XXX_Size() int {
	return xxx_messageInfo_RateLimitResponse_RateLimit.Size(m)
}
func (m *RateLimitResponse_RateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitResponse_RateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitResponse_RateLimit proto.InternalMessageInfo

func (m *RateLimitResponse_RateLimit) GetRequestsPerUnit() uint32 {
	if m != nil {
		return m.RequestsPerUnit
	}
	return 0
}

func (m *RateLimitResponse_RateLimit) GetUnit() RateLimitResponse_RateLimit_Unit {
	if m != nil {
		return m.Unit
	}
	return RateLimitResponse_RateLimit_UNKNOWN
}

type RateLimitResponse_DescriptorStatus struct {
	Code                 RateLimitResponse_Code       `protobuf:"varint,1,opt,name=code,proto3,enum=envoy.service.ratelimit.v3alpha.RateLimitResponse_Code" json:"code,omitempty"`
	CurrentLimit         *RateLimitResponse_RateLimit `protobuf:"bytes,2,opt,name=current_limit,json=currentLimit,proto3" json:"current_limit,omitempty"`
	LimitRemaining       uint32                       `protobuf:"varint,3,opt,name=limit_remaining,json=limitRemaining,proto3" json:"limit_remaining,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *RateLimitResponse_DescriptorStatus) Reset()         { *m = RateLimitResponse_DescriptorStatus{} }
func (m *RateLimitResponse_DescriptorStatus) String() string { return proto.CompactTextString(m) }
func (*RateLimitResponse_DescriptorStatus) ProtoMessage()    {}
func (*RateLimitResponse_DescriptorStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_5929a49ed6c3ed07, []int{1, 1}
}

func (m *RateLimitResponse_DescriptorStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitResponse_DescriptorStatus.Unmarshal(m, b)
}
func (m *RateLimitResponse_DescriptorStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitResponse_DescriptorStatus.Marshal(b, m, deterministic)
}
func (m *RateLimitResponse_DescriptorStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitResponse_DescriptorStatus.Merge(m, src)
}
func (m *RateLimitResponse_DescriptorStatus) XXX_Size() int {
	return xxx_messageInfo_RateLimitResponse_DescriptorStatus.Size(m)
}
func (m *RateLimitResponse_DescriptorStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitResponse_DescriptorStatus.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitResponse_DescriptorStatus proto.InternalMessageInfo

func (m *RateLimitResponse_DescriptorStatus) GetCode() RateLimitResponse_Code {
	if m != nil {
		return m.Code
	}
	return RateLimitResponse_UNKNOWN
}

func (m *RateLimitResponse_DescriptorStatus) GetCurrentLimit() *RateLimitResponse_RateLimit {
	if m != nil {
		return m.CurrentLimit
	}
	return nil
}

func (m *RateLimitResponse_DescriptorStatus) GetLimitRemaining() uint32 {
	if m != nil {
		return m.LimitRemaining
	}
	return 0
}

func init() {
	proto.RegisterEnum("envoy.service.ratelimit.v3alpha.RateLimitResponse_Code", RateLimitResponse_Code_name, RateLimitResponse_Code_value)
	proto.RegisterEnum("envoy.service.ratelimit.v3alpha.RateLimitResponse_RateLimit_Unit", RateLimitResponse_RateLimit_Unit_name, RateLimitResponse_RateLimit_Unit_value)
	proto.RegisterType((*RateLimitRequest)(nil), "envoy.service.ratelimit.v3alpha.RateLimitRequest")
	proto.RegisterType((*RateLimitResponse)(nil), "envoy.service.ratelimit.v3alpha.RateLimitResponse")
	proto.RegisterType((*RateLimitResponse_RateLimit)(nil), "envoy.service.ratelimit.v3alpha.RateLimitResponse.RateLimit")
	proto.RegisterType((*RateLimitResponse_DescriptorStatus)(nil), "envoy.service.ratelimit.v3alpha.RateLimitResponse.DescriptorStatus")
}

func init() {
	proto.RegisterFile("envoy/service/ratelimit/v3alpha/rls.proto", fileDescriptor_5929a49ed6c3ed07)
}

var fileDescriptor_5929a49ed6c3ed07 = []byte{
	// 709 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xdd, 0x4e, 0x13, 0x41,
	0x14, 0xc7, 0xd9, 0xb6, 0x29, 0xe5, 0x94, 0x8f, 0x65, 0x34, 0xd0, 0xf4, 0x42, 0xb0, 0x31, 0x8a,
	0x22, 0xbb, 0xb1, 0x24, 0x7e, 0x34, 0x82, 0x29, 0x94, 0x04, 0x02, 0xb4, 0x64, 0x4b, 0x31, 0x6a,
	0xcc, 0x3a, 0x74, 0x47, 0xba, 0xc9, 0x76, 0xa6, 0xce, 0x4c, 0x1b, 0xf0, 0x09, 0xb8, 0xd1, 0x07,
	0xf0, 0x2d, 0xf4, 0x1d, 0x7c, 0x10, 0xdf, 0xc4, 0xec, 0xcc, 0x76, 0xa9, 0x45, 0x83, 0x45, 0xef,
	0x76, 0xcf, 0x9c, 0xf3, 0xfb, 0xcf, 0xfc, 0xcf, 0x99, 0x81, 0xfb, 0x84, 0xf6, 0xd8, 0x99, 0x2d,
	0x08, 0xef, 0xf9, 0x4d, 0x62, 0x73, 0x2c, 0x49, 0xe0, 0xb7, 0x7d, 0x69, 0xf7, 0x56, 0x71, 0xd0,
	0x69, 0x61, 0x9b, 0x07, 0xc2, 0xea, 0x70, 0x26, 0x19, 0x5a, 0x50, 0xa9, 0x56, 0x94, 0x6a, 0xc5,
	0xa9, 0x56, 0x94, 0x9a, 0xbf, 0xa3, 0x59, 0x4d, 0x46, 0xdf, 0xfb, 0x27, 0x76, 0x93, 0x71, 0x12,
	0x53, 0x8e, 0xb1, 0x20, 0x1a, 0x93, 0x7f, 0xa6, 0xb3, 0xc8, 0xa9, 0x24, 0x54, 0xf8, 0x8c, 0x0a,
	0xbb, 0xc9, 0xda, 0x6d, 0x46, 0x7f, 0xa7, 0x1d, 0x4b, 0xe8, 0xd2, 0xdb, 0x5d, 0xaf, 0x83, 0x6d,
	0x4c, 0x29, 0x93, 0x58, 0xaa, 0xd2, 0x1e, 0xe1, 0x21, 0xc3, 0xa7, 0x27, 0x51, 0xca, 0x7c, 0x0f,
	0x07, 0xbe, 0x87, 0x25, 0xb1, 0xfb, 0x1f, 0x7a, 0xa1, 0xf0, 0xc3, 0x00, 0xd3, 0xc1, 0x92, 0xec,
	0x85, 0x3c, 0x87, 0x7c, 0xe8, 0x12, 0x21, 0xd1, 0x1c, 0xa4, 0x3d, 0xd6, 0xc6, 0x3e, 0xcd, 0x19,
	0x8b, 0xc6, 0xd2, 0x84, 0x13, 0xfd, 0xa1, 0x77, 0x90, 0xf5, 0x88, 0x68, 0x72, 0xbf, 0x23, 0x19,
	0x17, 0xb9, 0xc4, 0x62, 0x72, 0x29, 0x5b, 0x5c, 0xb7, 0xb4, 0x01, 0x17, 0x3b, 0xb7, 0xf4, 0xce,
	0x2f, 0x5b, 0x61, 0xc5, 0x4a, 0x95, 0x18, 0xe3, 0x0c, 0x22, 0xd1, 0x02, 0x64, 0x5b, 0xbe, 0x14,
	0x2e, 0xf6, 0x3c, 0x42, 0xbd, 0x5c, 0x72, 0xd1, 0x58, 0x9a, 0x72, 0x20, 0x0c, 0x95, 0x55, 0xa4,
	0x54, 0xfc, 0xf2, 0xfd, 0xfc, 0xd6, 0x0a, 0x2c, 0xff, 0xd1, 0xf4, 0xa2, 0x35, 0x7c, 0x9c, 0xc2,
	0xb7, 0x0c, 0xcc, 0x0e, 0x04, 0x45, 0x87, 0x51, 0x41, 0xd0, 0x6b, 0x98, 0x64, 0x3d, 0xc2, 0x71,
	0x10, 0xb8, 0x4d, 0xe6, 0x11, 0x75, 0xd4, 0xe9, 0xe2, 0x13, 0xeb, 0x8a, 0x76, 0x5a, 0x97, 0x48,
	0xd6, 0x26, 0xf3, 0x88, 0x93, 0x8d, 0x60, 0xe1, 0x0f, 0x72, 0x21, 0x23, 0x24, 0x96, 0x5d, 0x41,
	0xfa, 0x2e, 0x6d, 0x5e, 0x83, 0x7b, 0x61, 0x52, 0x5d, 0xc1, 0x9c, 0x18, 0x8a, 0xde, 0xc2, 0x3c,
	0x8f, 0xd2, 0xdc, 0x16, 0xc1, 0x1e, 0xe1, 0xc2, 0x95, 0x2c, 0xb4, 0x2d, 0x97, 0x54, 0x7a, 0x77,
	0x23, 0x3d, 0x3d, 0x75, 0x56, 0x38, 0x75, 0xb1, 0xd2, 0xb6, 0x2a, 0x38, 0xc2, 0x41, 0x97, 0x38,
	0x37, 0xfb, 0x18, 0x1d, 0x14, 0x87, 0xac, 0xec, 0x79, 0xe8, 0x0d, 0xcc, 0x71, 0x6d, 0xde, 0x30,
	0x3d, 0x35, 0x12, 0xfd, 0x46, 0x44, 0x19, 0x84, 0xe7, 0x3f, 0x25, 0x60, 0x22, 0x3e, 0x2c, 0x7a,
	0x00, 0xb3, 0x51, 0x92, 0x70, 0x3b, 0x84, 0xbb, 0x5d, 0xea, 0x4b, 0xd5, 0x8b, 0x29, 0x67, 0xa6,
	0xbf, 0x70, 0x40, 0x78, 0x83, 0xfa, 0x12, 0x35, 0x20, 0xa5, 0x96, 0x13, 0xaa, 0x55, 0xe5, 0x6b,
	0x58, 0x1a, 0x47, 0xac, 0x10, 0xe8, 0x28, 0x5c, 0x61, 0x1d, 0x52, 0x0a, 0x9f, 0x85, 0xf1, 0x46,
	0x75, 0xb7, 0x5a, 0x7b, 0x59, 0x35, 0xc7, 0x10, 0x40, 0xba, 0xbe, 0xb5, 0x59, 0xab, 0x56, 0x4c,
	0x23, 0xfc, 0xde, 0xdf, 0xa9, 0x36, 0x0e, 0xb7, 0xcc, 0x04, 0xca, 0x40, 0x6a, 0xbb, 0xd6, 0x70,
	0xcc, 0x24, 0x1a, 0x87, 0x64, 0xa5, 0xfc, 0xca, 0x4c, 0x95, 0xd6, 0xc2, 0x99, 0x7c, 0x0a, 0x8f,
	0xff, 0x6e, 0x26, 0x87, 0x77, 0x92, 0xff, 0x9a, 0x00, 0x73, 0xb8, 0xd5, 0x68, 0x17, 0x52, 0xff,
	0x63, 0x2a, 0x15, 0x04, 0x61, 0x98, 0x6a, 0x76, 0x39, 0x27, 0x54, 0xba, 0xaa, 0x4a, 0x19, 0x98,
	0x2d, 0x3e, 0xff, 0x17, 0x03, 0x9d, 0xc9, 0x08, 0xa9, 0xdb, 0x78, 0x0f, 0x66, 0x54, 0xa9, 0xcb,
	0x49, 0xf8, 0x54, 0xf8, 0xf4, 0x24, 0xba, 0xbc, 0xd3, 0x81, 0xae, 0x8f, 0xa2, 0xa5, 0x4a, 0x68,
	0xd6, 0x0b, 0x58, 0x1b, 0xc9, 0xac, 0x61, 0x7b, 0x0a, 0xcb, 0x90, 0x52, 0x17, 0xed, 0x97, 0x96,
	0xa5, 0x21, 0x51, 0xdb, 0x35, 0x0d, 0x34, 0x0d, 0x50, 0x3b, 0xda, 0x72, 0xdc, 0xbd, 0x9d, 0xfd,
	0x9d, 0x43, 0x33, 0x51, 0x5a, 0x0d, 0x25, 0x2d, 0x78, 0x38, 0x8a, 0x64, 0xf1, 0xf3, 0xe0, 0xc3,
	0x58, 0xd7, 0x35, 0xe8, 0x23, 0xcc, 0xd4, 0x5b, 0xac, 0x1b, 0x78, 0x17, 0xf3, 0xfb, 0x68, 0x14,
	0x13, 0xd5, 0x38, 0xe7, 0x8b, 0xa3, 0xfb, 0x5e, 0x18, 0xdb, 0x58, 0x87, 0x15, 0x9f, 0xe9, 0xca,
	0x0e, 0x67, 0xa7, 0x67, 0x57, 0x41, 0x36, 0x32, 0x4e, 0x20, 0x0e, 0xc2, 0x47, 0xfe, 0xc0, 0x38,
	0x37, 0x8c, 0xe3, 0xb4, 0x7a, 0xf0, 0x57, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xa2, 0x1e,
	0x2d, 0xdb, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RateLimitServiceClient is the client API for RateLimitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RateLimitServiceClient interface {
	ShouldRateLimit(ctx context.Context, in *RateLimitRequest, opts ...grpc.CallOption) (*RateLimitResponse, error)
}

type rateLimitServiceClient struct {
	cc *grpc.ClientConn
}

func NewRateLimitServiceClient(cc *grpc.ClientConn) RateLimitServiceClient {
	return &rateLimitServiceClient{cc}
}

func (c *rateLimitServiceClient) ShouldRateLimit(ctx context.Context, in *RateLimitRequest, opts ...grpc.CallOption) (*RateLimitResponse, error) {
	out := new(RateLimitResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.ratelimit.v3alpha.RateLimitService/ShouldRateLimit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RateLimitServiceServer is the server API for RateLimitService service.
type RateLimitServiceServer interface {
	ShouldRateLimit(context.Context, *RateLimitRequest) (*RateLimitResponse, error)
}

// UnimplementedRateLimitServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRateLimitServiceServer struct {
}

func (*UnimplementedRateLimitServiceServer) ShouldRateLimit(ctx context.Context, req *RateLimitRequest) (*RateLimitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShouldRateLimit not implemented")
}

func RegisterRateLimitServiceServer(s *grpc.Server, srv RateLimitServiceServer) {
	s.RegisterService(&_RateLimitService_serviceDesc, srv)
}

func _RateLimitService_ShouldRateLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateLimitServiceServer).ShouldRateLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.ratelimit.v3alpha.RateLimitService/ShouldRateLimit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateLimitServiceServer).ShouldRateLimit(ctx, req.(*RateLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RateLimitService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.ratelimit.v3alpha.RateLimitService",
	HandlerType: (*RateLimitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShouldRateLimit",
			Handler:    _RateLimitService_ShouldRateLimit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "envoy/service/ratelimit/v3alpha/rls.proto",
}
