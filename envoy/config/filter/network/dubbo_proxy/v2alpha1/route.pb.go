// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/dubbo_proxy/v2alpha1/route.proto

package envoy_config_filter_network_dubbo_proxy_v2alpha1

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	_type "github.com/envoyproxy/go-control-plane/envoy/type"
	matcher "github.com/envoyproxy/go-control-plane/envoy/type/matcher"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type RouteConfiguration struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Interface            string   `protobuf:"bytes,2,opt,name=interface,proto3" json:"interface,omitempty"`
	Group                string   `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`
	Version              string   `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	Routes               []*Route `protobuf:"bytes,5,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteConfiguration) Reset()         { *m = RouteConfiguration{} }
func (m *RouteConfiguration) String() string { return proto.CompactTextString(m) }
func (*RouteConfiguration) ProtoMessage()    {}
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_74a572433a3292e0, []int{0}
}

func (m *RouteConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteConfiguration.Unmarshal(m, b)
}
func (m *RouteConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteConfiguration.Marshal(b, m, deterministic)
}
func (m *RouteConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfiguration.Merge(m, src)
}
func (m *RouteConfiguration) XXX_Size() int {
	return xxx_messageInfo_RouteConfiguration.Size(m)
}
func (m *RouteConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfiguration proto.InternalMessageInfo

func (m *RouteConfiguration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteConfiguration) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *RouteConfiguration) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *RouteConfiguration) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RouteConfiguration) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type Route struct {
	Match                *RouteMatch  `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	Route                *RouteAction `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_74a572433a3292e0, []int{1}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetMatch() *RouteMatch {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *Route) GetRoute() *RouteAction {
	if m != nil {
		return m.Route
	}
	return nil
}

type RouteMatch struct {
	Method               *MethodMatch           `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Headers              []*route.HeaderMatcher `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *RouteMatch) Reset()         { *m = RouteMatch{} }
func (m *RouteMatch) String() string { return proto.CompactTextString(m) }
func (*RouteMatch) ProtoMessage()    {}
func (*RouteMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_74a572433a3292e0, []int{2}
}

func (m *RouteMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteMatch.Unmarshal(m, b)
}
func (m *RouteMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteMatch.Marshal(b, m, deterministic)
}
func (m *RouteMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteMatch.Merge(m, src)
}
func (m *RouteMatch) XXX_Size() int {
	return xxx_messageInfo_RouteMatch.Size(m)
}
func (m *RouteMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteMatch.DiscardUnknown(m)
}

var xxx_messageInfo_RouteMatch proto.InternalMessageInfo

func (m *RouteMatch) GetMethod() *MethodMatch {
	if m != nil {
		return m.Method
	}
	return nil
}

func (m *RouteMatch) GetHeaders() []*route.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

type RouteAction struct {
	// Types that are valid to be assigned to ClusterSpecifier:
	//	*RouteAction_Cluster
	//	*RouteAction_WeightedClusters
	ClusterSpecifier     isRouteAction_ClusterSpecifier `protobuf_oneof:"cluster_specifier"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *RouteAction) Reset()         { *m = RouteAction{} }
func (m *RouteAction) String() string { return proto.CompactTextString(m) }
func (*RouteAction) ProtoMessage()    {}
func (*RouteAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_74a572433a3292e0, []int{3}
}

func (m *RouteAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteAction.Unmarshal(m, b)
}
func (m *RouteAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteAction.Marshal(b, m, deterministic)
}
func (m *RouteAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteAction.Merge(m, src)
}
func (m *RouteAction) XXX_Size() int {
	return xxx_messageInfo_RouteAction.Size(m)
}
func (m *RouteAction) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteAction.DiscardUnknown(m)
}

var xxx_messageInfo_RouteAction proto.InternalMessageInfo

type isRouteAction_ClusterSpecifier interface {
	isRouteAction_ClusterSpecifier()
}

type RouteAction_Cluster struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3,oneof"`
}

type RouteAction_WeightedClusters struct {
	WeightedClusters *route.WeightedCluster `protobuf:"bytes,2,opt,name=weighted_clusters,json=weightedClusters,proto3,oneof"`
}

func (*RouteAction_Cluster) isRouteAction_ClusterSpecifier() {}

func (*RouteAction_WeightedClusters) isRouteAction_ClusterSpecifier() {}

func (m *RouteAction) GetClusterSpecifier() isRouteAction_ClusterSpecifier {
	if m != nil {
		return m.ClusterSpecifier
	}
	return nil
}

func (m *RouteAction) GetCluster() string {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *RouteAction) GetWeightedClusters() *route.WeightedCluster {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_WeightedClusters); ok {
		return x.WeightedClusters
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RouteAction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RouteAction_Cluster)(nil),
		(*RouteAction_WeightedClusters)(nil),
	}
}

type MethodMatch struct {
	Name                 *matcher.StringMatcher                          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ParamsMatch          map[uint32]*MethodMatch_ParameterMatchSpecifier `protobuf:"bytes,2,rep,name=params_match,json=paramsMatch,proto3" json:"params_match,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                                        `json:"-"`
	XXX_unrecognized     []byte                                          `json:"-"`
	XXX_sizecache        int32                                           `json:"-"`
}

func (m *MethodMatch) Reset()         { *m = MethodMatch{} }
func (m *MethodMatch) String() string { return proto.CompactTextString(m) }
func (*MethodMatch) ProtoMessage()    {}
func (*MethodMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_74a572433a3292e0, []int{4}
}

func (m *MethodMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MethodMatch.Unmarshal(m, b)
}
func (m *MethodMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MethodMatch.Marshal(b, m, deterministic)
}
func (m *MethodMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MethodMatch.Merge(m, src)
}
func (m *MethodMatch) XXX_Size() int {
	return xxx_messageInfo_MethodMatch.Size(m)
}
func (m *MethodMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_MethodMatch.DiscardUnknown(m)
}

var xxx_messageInfo_MethodMatch proto.InternalMessageInfo

func (m *MethodMatch) GetName() *matcher.StringMatcher {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *MethodMatch) GetParamsMatch() map[uint32]*MethodMatch_ParameterMatchSpecifier {
	if m != nil {
		return m.ParamsMatch
	}
	return nil
}

type MethodMatch_ParameterMatchSpecifier struct {
	// Types that are valid to be assigned to ParameterMatchSpecifier:
	//	*MethodMatch_ParameterMatchSpecifier_ExactMatch
	//	*MethodMatch_ParameterMatchSpecifier_RangeMatch
	ParameterMatchSpecifier isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier `protobuf_oneof:"parameter_match_specifier"`
	XXX_NoUnkeyedLiteral    struct{}                                                      `json:"-"`
	XXX_unrecognized        []byte                                                        `json:"-"`
	XXX_sizecache           int32                                                         `json:"-"`
}

func (m *MethodMatch_ParameterMatchSpecifier) Reset()         { *m = MethodMatch_ParameterMatchSpecifier{} }
func (m *MethodMatch_ParameterMatchSpecifier) String() string { return proto.CompactTextString(m) }
func (*MethodMatch_ParameterMatchSpecifier) ProtoMessage()    {}
func (*MethodMatch_ParameterMatchSpecifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_74a572433a3292e0, []int{4, 0}
}

func (m *MethodMatch_ParameterMatchSpecifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.Unmarshal(m, b)
}
func (m *MethodMatch_ParameterMatchSpecifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.Marshal(b, m, deterministic)
}
func (m *MethodMatch_ParameterMatchSpecifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.Merge(m, src)
}
func (m *MethodMatch_ParameterMatchSpecifier) XXX_Size() int {
	return xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.Size(m)
}
func (m *MethodMatch_ParameterMatchSpecifier) XXX_DiscardUnknown() {
	xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.DiscardUnknown(m)
}

var xxx_messageInfo_MethodMatch_ParameterMatchSpecifier proto.InternalMessageInfo

type isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier interface {
	isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier()
}

type MethodMatch_ParameterMatchSpecifier_ExactMatch struct {
	ExactMatch string `protobuf:"bytes,3,opt,name=exact_match,json=exactMatch,proto3,oneof"`
}

type MethodMatch_ParameterMatchSpecifier_RangeMatch struct {
	RangeMatch *_type.Int64Range `protobuf:"bytes,4,opt,name=range_match,json=rangeMatch,proto3,oneof"`
}

func (*MethodMatch_ParameterMatchSpecifier_ExactMatch) isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier() {
}

func (*MethodMatch_ParameterMatchSpecifier_RangeMatch) isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier() {
}

func (m *MethodMatch_ParameterMatchSpecifier) GetParameterMatchSpecifier() isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier {
	if m != nil {
		return m.ParameterMatchSpecifier
	}
	return nil
}

func (m *MethodMatch_ParameterMatchSpecifier) GetExactMatch() string {
	if x, ok := m.GetParameterMatchSpecifier().(*MethodMatch_ParameterMatchSpecifier_ExactMatch); ok {
		return x.ExactMatch
	}
	return ""
}

func (m *MethodMatch_ParameterMatchSpecifier) GetRangeMatch() *_type.Int64Range {
	if x, ok := m.GetParameterMatchSpecifier().(*MethodMatch_ParameterMatchSpecifier_RangeMatch); ok {
		return x.RangeMatch
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MethodMatch_ParameterMatchSpecifier) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MethodMatch_ParameterMatchSpecifier_ExactMatch)(nil),
		(*MethodMatch_ParameterMatchSpecifier_RangeMatch)(nil),
	}
}

func init() {
	proto.RegisterType((*RouteConfiguration)(nil), "envoy.config.filter.network.dubbo_proxy.v2alpha1.RouteConfiguration")
	proto.RegisterType((*Route)(nil), "envoy.config.filter.network.dubbo_proxy.v2alpha1.Route")
	proto.RegisterType((*RouteMatch)(nil), "envoy.config.filter.network.dubbo_proxy.v2alpha1.RouteMatch")
	proto.RegisterType((*RouteAction)(nil), "envoy.config.filter.network.dubbo_proxy.v2alpha1.RouteAction")
	proto.RegisterType((*MethodMatch)(nil), "envoy.config.filter.network.dubbo_proxy.v2alpha1.MethodMatch")
	proto.RegisterMapType((map[uint32]*MethodMatch_ParameterMatchSpecifier)(nil), "envoy.config.filter.network.dubbo_proxy.v2alpha1.MethodMatch.ParamsMatchEntry")
	proto.RegisterType((*MethodMatch_ParameterMatchSpecifier)(nil), "envoy.config.filter.network.dubbo_proxy.v2alpha1.MethodMatch.ParameterMatchSpecifier")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/dubbo_proxy/v2alpha1/route.proto", fileDescriptor_74a572433a3292e0)
}

var fileDescriptor_74a572433a3292e0 = []byte{
	// 698 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0xae, 0x93, 0x6e, 0x4b, 0x67, 0x41, 0x4a, 0x2d, 0xd4, 0x2e, 0x01, 0x41, 0x1b, 0x2e, 0xe5,
	0xb2, 0x0b, 0x69, 0x29, 0x7f, 0x05, 0x89, 0x54, 0x48, 0xe5, 0x50, 0xa8, 0xb6, 0xaa, 0xb8, 0x80,
	0x22, 0x77, 0xe3, 0x24, 0x56, 0x13, 0x7b, 0xf1, 0x3a, 0x69, 0xf3, 0x0a, 0xbd, 0x70, 0x01, 0xc1,
	0x91, 0x87, 0x41, 0xe2, 0x4d, 0x78, 0x00, 0x4e, 0x88, 0x03, 0x42, 0x3b, 0xf6, 0x92, 0x00, 0xbd,
	0xb4, 0x70, 0x59, 0x79, 0xc7, 0xf3, 0x7d, 0xdf, 0xcc, 0xa7, 0x19, 0xc3, 0x06, 0x97, 0x43, 0x35,
	0x8a, 0x12, 0x25, 0xdb, 0xa2, 0x13, 0xb5, 0x45, 0xcf, 0x70, 0x1d, 0x49, 0x6e, 0x0e, 0x95, 0x3e,
	0x88, 0x5a, 0x83, 0xfd, 0x7d, 0xd5, 0x4c, 0xb5, 0x3a, 0x1a, 0x45, 0xc3, 0x3a, 0xeb, 0xa5, 0x5d,
	0x76, 0x2b, 0xd2, 0x6a, 0x60, 0x78, 0x98, 0x6a, 0x65, 0x14, 0xbd, 0x89, 0xe8, 0xd0, 0xa2, 0x43,
	0x8b, 0x0e, 0x1d, 0x3a, 0x9c, 0x40, 0x87, 0x05, 0xba, 0x7a, 0xc3, 0xea, 0xb1, 0x54, 0x44, 0xc3,
	0xba, 0xe5, 0xb2, 0xdf, 0x66, 0xa2, 0xfa, 0xa9, 0x92, 0x5c, 0x9a, 0xcc, 0x92, 0x57, 0xaf, 0xd9,
	0x54, 0x33, 0x4a, 0x79, 0xd4, 0x67, 0x26, 0xe9, 0x72, 0x1d, 0x65, 0x46, 0x0b, 0xd9, 0x71, 0x09,
	0x0b, 0x13, 0x09, 0x9a, 0xc9, 0x8e, 0xab, 0xaa, 0x7a, 0x75, 0xd0, 0x4a, 0x59, 0xc4, 0xa4, 0x54,
	0x86, 0x19, 0xa1, 0x64, 0x16, 0xf5, 0x45, 0x47, 0xb3, 0xa2, 0xea, 0xea, 0xe2, 0x90, 0xf5, 0x44,
	0x8b, 0x19, 0x1e, 0x15, 0x07, 0x7b, 0x51, 0xfb, 0x4c, 0x80, 0xc6, 0x79, 0x31, 0x9b, 0xd8, 0xd0,
	0x40, 0x23, 0x9c, 0x52, 0x98, 0x96, 0xac, 0xcf, 0x03, 0xb2, 0x44, 0x56, 0xe6, 0x62, 0x3c, 0xd3,
	0x2b, 0x30, 0x27, 0xa4, 0xe1, 0xba, 0xcd, 0x12, 0x1e, 0x94, 0xf0, 0x62, 0x1c, 0xa0, 0x17, 0xc1,
	0xeb, 0x68, 0x35, 0x48, 0x83, 0x32, 0xde, 0xd8, 0x1f, 0x1a, 0xc0, 0xec, 0x90, 0xeb, 0x4c, 0x28,
	0x19, 0x4c, 0x63, 0xbc, 0xf8, 0xa5, 0xcf, 0x61, 0x06, 0x4d, 0xc8, 0x02, 0x6f, 0xa9, 0xbc, 0xe2,
	0xd7, 0xef, 0x84, 0xa7, 0x35, 0x36, 0xc4, 0xba, 0x63, 0x47, 0x53, 0xfb, 0x44, 0xc0, 0xc3, 0x08,
	0x7d, 0x09, 0x1e, 0x9a, 0x87, 0xd5, 0xfb, 0xf5, 0x8d, 0x33, 0x32, 0x6f, 0xe7, 0x1c, 0x8d, 0x73,
	0xdf, 0x1b, 0xde, 0x31, 0x29, 0x55, 0x48, 0x6c, 0x49, 0xe9, 0x2b, 0xf0, 0x50, 0x11, 0x2d, 0xf0,
	0xeb, 0x0f, 0xcf, 0xc8, 0xfe, 0x38, 0xc9, 0x8d, 0x9e, 0xa4, 0x47, 0xd6, 0xda, 0x47, 0x02, 0x30,
	0x96, 0xa7, 0x7b, 0x30, 0xd3, 0xe7, 0xa6, 0xab, 0x5a, 0xae, 0x99, 0x33, 0xc8, 0x6d, 0x23, 0x1e,
	0xe9, 0x62, 0x47, 0x46, 0x1f, 0xc0, 0x6c, 0x97, 0xb3, 0x16, 0xd7, 0x59, 0x50, 0x42, 0xfb, 0x97,
	0x1d, 0x2f, 0x4b, 0x45, 0x38, 0xac, 0x87, 0x76, 0xe2, 0xb7, 0x30, 0x65, 0xdb, 0x0e, 0x62, 0x5c,
	0x20, 0x6a, 0xef, 0x09, 0xf8, 0x13, 0x3d, 0xd0, 0x2a, 0xcc, 0x26, 0xbd, 0x41, 0x66, 0xb8, 0xb6,
	0xf3, 0xb2, 0x35, 0x15, 0x17, 0x01, 0x1a, 0xc3, 0xfc, 0x21, 0x17, 0x9d, 0xae, 0xe1, 0xad, 0xa6,
	0x8b, 0x65, 0xce, 0xb9, 0xeb, 0x27, 0x49, 0xbe, 0x70, 0xc9, 0x9b, 0x36, 0x77, 0x6b, 0x2a, 0xae,
	0x1c, 0xfe, 0x1e, 0xca, 0x1a, 0x01, 0xcc, 0x3b, 0xaa, 0x66, 0x96, 0xf2, 0x44, 0xb4, 0x05, 0xd7,
	0xb4, 0xfc, 0xad, 0x41, 0x6a, 0x5f, 0xca, 0xe0, 0x4f, 0xb4, 0x4b, 0x6f, 0x4f, 0x8c, 0xf1, 0xb8,
	0xc7, 0x7c, 0x7b, 0x42, 0xb7, 0x5e, 0xe1, 0x2e, 0xae, 0x57, 0xd1, 0xa3, 0x9d, 0xf4, 0xd7, 0x70,
	0x3e, 0x65, 0x9a, 0xf5, 0xb3, 0xa6, 0x9d, 0x23, 0x6b, 0xd1, 0xb3, 0x7f, 0xb2, 0x3e, 0xdc, 0x41,
	0x46, 0x3c, 0x3f, 0x91, 0x46, 0x8f, 0x62, 0x3f, 0x1d, 0x47, 0xaa, 0x6f, 0x09, 0x2c, 0x62, 0x06,
	0x37, 0xce, 0xf1, 0xdd, 0x5f, 0xad, 0x2d, 0x83, 0xcf, 0x8f, 0x58, 0x62, 0x5c, 0x35, 0x65, 0xe7,
	0x31, 0x60, 0xd0, 0x36, 0x7a, 0x0f, 0x7c, 0x7c, 0x0e, 0x5c, 0xca, 0x34, 0xf6, 0xbb, 0x30, 0xd9,
	0xef, 0x53, 0x69, 0xd6, 0xd7, 0xe2, 0x3c, 0x27, 0x87, 0x62, 0xb2, 0x1d, 0xf0, 0xcb, 0x70, 0x29,
	0x2d, 0x84, 0x2d, 0x7c, 0xec, 0x6a, 0xf5, 0x1d, 0x81, 0xca, 0x9f, 0x85, 0xd3, 0x0a, 0x94, 0x0f,
	0xf8, 0x08, 0x4d, 0xbd, 0x10, 0xe7, 0x47, 0x7a, 0x00, 0xde, 0x90, 0xf5, 0x06, 0xc5, 0x4e, 0xec,
	0xfd, 0x07, 0xa7, 0xfe, 0xf6, 0x21, 0xb6, 0x1a, 0xf7, 0x4b, 0x77, 0x49, 0xe3, 0x98, 0x7c, 0xfd,
	0xf0, 0xe3, 0x8d, 0xb7, 0x4e, 0xd7, 0xac, 0x12, 0x3f, 0x32, 0x5c, 0xe6, 0xef, 0x4a, 0xe6, 0xd4,
	0xb2, 0x93, 0xe5, 0x56, 0x51, 0x0e, 0x1e, 0x09, 0x65, 0x4b, 0xb4, 0xe1, 0xd3, 0x56, 0xdb, 0xb0,
	0x1b, 0xba, 0x93, 0xbf, 0xa0, 0x3b, 0x64, 0x7f, 0x06, 0x9f, 0xd2, 0xd5, 0x9f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x90, 0x95, 0x27, 0x2b, 0x59, 0x06, 0x00, 0x00,
}
