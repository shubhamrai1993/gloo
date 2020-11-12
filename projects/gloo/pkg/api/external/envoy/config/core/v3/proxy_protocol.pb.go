// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo-edge/projects/gloo/api/external/envoy/config/core/v3/proxy_protocol.proto

package v3

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/gloo-edge/projects/gloo/pkg/api/external/udpa/annotations"
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

type ProxyProtocolConfig_Version int32

const (
	// PROXY protocol version 1. Human readable format.
	ProxyProtocolConfig_V1 ProxyProtocolConfig_Version = 0
	// PROXY protocol version 2. Binary format.
	ProxyProtocolConfig_V2 ProxyProtocolConfig_Version = 1
)

var ProxyProtocolConfig_Version_name = map[int32]string{
	0: "V1",
	1: "V2",
}

var ProxyProtocolConfig_Version_value = map[string]int32{
	"V1": 0,
	"V2": 1,
}

func (x ProxyProtocolConfig_Version) String() string {
	return proto.EnumName(ProxyProtocolConfig_Version_name, int32(x))
}

func (ProxyProtocolConfig_Version) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_990e30a2fd2cd065, []int{0, 0}
}

type ProxyProtocolConfig struct {
	// The PROXY protocol version to use. See https://www.haproxy.org/download/2.1/doc/proxy-protocol.txt for details
	Version              ProxyProtocolConfig_Version `protobuf:"varint,1,opt,name=version,proto3,enum=envoy.config.core.v3.ProxyProtocolConfig_Version" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ProxyProtocolConfig) Reset()         { *m = ProxyProtocolConfig{} }
func (m *ProxyProtocolConfig) String() string { return proto.CompactTextString(m) }
func (*ProxyProtocolConfig) ProtoMessage()    {}
func (*ProxyProtocolConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_990e30a2fd2cd065, []int{0}
}
func (m *ProxyProtocolConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyProtocolConfig.Unmarshal(m, b)
}
func (m *ProxyProtocolConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyProtocolConfig.Marshal(b, m, deterministic)
}
func (m *ProxyProtocolConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyProtocolConfig.Merge(m, src)
}
func (m *ProxyProtocolConfig) XXX_Size() int {
	return xxx_messageInfo_ProxyProtocolConfig.Size(m)
}
func (m *ProxyProtocolConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyProtocolConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyProtocolConfig proto.InternalMessageInfo

func (m *ProxyProtocolConfig) GetVersion() ProxyProtocolConfig_Version {
	if m != nil {
		return m.Version
	}
	return ProxyProtocolConfig_V1
}

func init() {
	proto.RegisterEnum("envoy.config.core.v3.ProxyProtocolConfig_Version", ProxyProtocolConfig_Version_name, ProxyProtocolConfig_Version_value)
	proto.RegisterType((*ProxyProtocolConfig)(nil), "envoy.config.core.v3.ProxyProtocolConfig")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo-edge/projects/gloo/api/external/envoy/config/core/v3/proxy_protocol.proto", fileDescriptor_990e30a2fd2cd065)
}

var fileDescriptor_990e30a2fd2cd065 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4a, 0xc4, 0x30,
	0x14, 0x86, 0x4d, 0x17, 0x1d, 0xc8, 0x42, 0x4a, 0x1d, 0x50, 0x07, 0x14, 0xe9, 0xca, 0x8d, 0x79,
	0xcc, 0xf4, 0x06, 0xe3, 0x72, 0x36, 0xc5, 0xc5, 0x2c, 0x44, 0x90, 0x4c, 0x27, 0xc6, 0x68, 0xcd,
	0x2b, 0x4d, 0x1a, 0x66, 0x16, 0x82, 0x77, 0xf0, 0x12, 0x9e, 0xc1, 0x13, 0xb8, 0xf5, 0x0a, 0x9e,
	0x44, 0x9a, 0x74, 0x16, 0x42, 0xc1, 0x59, 0xbd, 0x3f, 0x2f, 0xff, 0x9f, 0x2f, 0xbc, 0x47, 0xef,
	0xa4, 0xb2, 0x8f, 0xed, 0x8a, 0x95, 0xf8, 0x02, 0x06, 0x2b, 0xbc, 0x52, 0x08, 0xb2, 0x42, 0x84,
	0xba, 0xc1, 0x27, 0x51, 0x5a, 0x13, 0x4e, 0xbc, 0x56, 0x20, 0x36, 0x56, 0x34, 0x9a, 0x57, 0x20,
	0xb4, 0xc3, 0x2d, 0x94, 0xa8, 0x1f, 0x94, 0x84, 0x12, 0x1b, 0x01, 0x2e, 0xef, 0xfc, 0x9b, 0xed,
	0x7d, 0xdd, 0xa0, 0xc5, 0x12, 0x2b, 0xe6, 0x45, 0x3a, 0xf6, 0x56, 0x16, 0xac, 0xac, 0xb3, 0x32,
	0x97, 0x4f, 0xce, 0xda, 0x75, 0xcd, 0x81, 0x6b, 0x8d, 0x96, 0x5b, 0x85, 0xda, 0x80, 0xb1, 0xdc,
	0xb6, 0x26, 0x84, 0x26, 0xc7, 0x8e, 0x57, 0x6a, 0xcd, 0xad, 0x80, 0x9d, 0xe8, 0x2f, 0xc6, 0x12,
	0x25, 0x7a, 0x09, 0x9d, 0x0a, 0xdd, 0xec, 0x95, 0x1e, 0x15, 0x1d, 0xbb, 0xe8, 0xd1, 0xd7, 0x1e,
	0x96, 0x2e, 0xe8, 0xc8, 0x89, 0xc6, 0x28, 0xd4, 0x27, 0xe4, 0x82, 0x5c, 0x1e, 0xce, 0xa6, 0x6c,
	0xe8, 0x33, 0x6c, 0x20, 0xcb, 0x96, 0x21, 0x78, 0xb3, 0x7b, 0x21, 0x3b, 0xa5, 0xa3, 0xbe, 0x97,
	0xc6, 0x34, 0x5a, 0x4e, 0x93, 0x03, 0x5f, 0x67, 0x09, 0x99, 0xbf, 0x93, 0x8f, 0x9f, 0x73, 0xf2,
	0xf9, 0xf6, 0xf5, 0x1d, 0x47, 0x49, 0x44, 0x33, 0x85, 0x81, 0xe3, 0xe7, 0x31, 0x88, 0x9c, 0xa7,
	0x7f, 0x98, 0xbe, 0x16, 0xe4, 0x76, 0xb1, 0xdf, 0x26, 0xea, 0x67, 0xf9, 0xff, 0x36, 0x56, 0xb1,
	0x9f, 0x4d, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x38, 0xf9, 0xdd, 0xdf, 0x01, 0x00, 0x00,
}

func (this *ProxyProtocolConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProxyProtocolConfig)
	if !ok {
		that2, ok := that.(ProxyProtocolConfig)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
