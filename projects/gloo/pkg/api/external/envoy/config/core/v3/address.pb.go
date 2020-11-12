// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo-edge/projects/gloo/api/external/envoy/config/core/v3/address.proto

package v3

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

type SocketAddress_Protocol int32

const (
	SocketAddress_TCP SocketAddress_Protocol = 0
	SocketAddress_UDP SocketAddress_Protocol = 1
)

var SocketAddress_Protocol_name = map[int32]string{
	0: "TCP",
	1: "UDP",
}

var SocketAddress_Protocol_value = map[string]int32{
	"TCP": 0,
	"UDP": 1,
}

func (x SocketAddress_Protocol) String() string {
	return proto.EnumName(SocketAddress_Protocol_name, int32(x))
}

func (SocketAddress_Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fb356bd27ab3b7eb, []int{1, 0}
}

type Pipe struct {
	// Unix Domain Socket path. On Linux, paths starting with '@' will use the
	// abstract namespace. The starting '@' is replaced by a null byte by Envoy.
	// Paths starting with '@' will result in an error in environments other than
	// Linux.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// The mode for the Pipe. Not applicable for abstract sockets.
	Mode                 uint32   `protobuf:"varint,2,opt,name=mode,proto3" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pipe) Reset()         { *m = Pipe{} }
func (m *Pipe) String() string { return proto.CompactTextString(m) }
func (*Pipe) ProtoMessage()    {}
func (*Pipe) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb356bd27ab3b7eb, []int{0}
}
func (m *Pipe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pipe.Unmarshal(m, b)
}
func (m *Pipe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pipe.Marshal(b, m, deterministic)
}
func (m *Pipe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pipe.Merge(m, src)
}
func (m *Pipe) XXX_Size() int {
	return xxx_messageInfo_Pipe.Size(m)
}
func (m *Pipe) XXX_DiscardUnknown() {
	xxx_messageInfo_Pipe.DiscardUnknown(m)
}

var xxx_messageInfo_Pipe proto.InternalMessageInfo

func (m *Pipe) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Pipe) GetMode() uint32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

// [#next-free-field: 7]
type SocketAddress struct {
	Protocol SocketAddress_Protocol `protobuf:"varint,1,opt,name=protocol,proto3,enum=envoy.config.core.v3.SocketAddress_Protocol" json:"protocol,omitempty"`
	// The address for this socket. :ref:`Listeners <config_listeners>` will bind
	// to the address. An empty address is not allowed. Specify ``0.0.0.0`` or ``::``
	// to bind to any address. [#comment:TODO(zuercher) reinstate when implemented:
	// It is possible to distinguish a Listener address via the prefix/suffix matching
	// in :ref:`FilterChainMatch <envoy_api_msg_config.listener.v3.FilterChainMatch>`.] When used
	// within an upstream :ref:`BindConfig <envoy_api_msg_config.core.v3.BindConfig>`, the address
	// controls the source address of outbound connections. For :ref:`clusters
	// <envoy_api_msg_config.cluster.v3.Cluster>`, the cluster type determines whether the
	// address must be an IP (*STATIC* or *EDS* clusters) or a hostname resolved by DNS
	// (*STRICT_DNS* or *LOGICAL_DNS* clusters). Address resolution can be customized
	// via :ref:`resolver_name <envoy_api_field_config.core.v3.SocketAddress.resolver_name>`.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// Types that are valid to be assigned to PortSpecifier:
	//	*SocketAddress_PortValue
	//	*SocketAddress_NamedPort
	PortSpecifier isSocketAddress_PortSpecifier `protobuf_oneof:"port_specifier"`
	// The name of the custom resolver. This must have been registered with Envoy. If
	// this is empty, a context dependent default applies. If the address is a concrete
	// IP address, no resolution will occur. If address is a hostname this
	// should be set for resolution other than DNS. Specifying a custom resolver with
	// *STRICT_DNS* or *LOGICAL_DNS* will generate an error at runtime.
	ResolverName string `protobuf:"bytes,5,opt,name=resolver_name,json=resolverName,proto3" json:"resolver_name,omitempty"`
	// When binding to an IPv6 address above, this enables `IPv4 compatibility
	// <https://tools.ietf.org/html/rfc3493#page-11>`_. Binding to ``::`` will
	// allow both IPv4 and IPv6 connections, with peer IPv4 addresses mapped into
	// IPv6 space as ``::FFFF:<IPv4-address>``.
	Ipv4Compat           bool     `protobuf:"varint,6,opt,name=ipv4_compat,json=ipv4Compat,proto3" json:"ipv4_compat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SocketAddress) Reset()         { *m = SocketAddress{} }
func (m *SocketAddress) String() string { return proto.CompactTextString(m) }
func (*SocketAddress) ProtoMessage()    {}
func (*SocketAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb356bd27ab3b7eb, []int{1}
}
func (m *SocketAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocketAddress.Unmarshal(m, b)
}
func (m *SocketAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocketAddress.Marshal(b, m, deterministic)
}
func (m *SocketAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocketAddress.Merge(m, src)
}
func (m *SocketAddress) XXX_Size() int {
	return xxx_messageInfo_SocketAddress.Size(m)
}
func (m *SocketAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_SocketAddress.DiscardUnknown(m)
}

var xxx_messageInfo_SocketAddress proto.InternalMessageInfo

type isSocketAddress_PortSpecifier interface {
	isSocketAddress_PortSpecifier()
	Equal(interface{}) bool
}

type SocketAddress_PortValue struct {
	PortValue uint32 `protobuf:"varint,3,opt,name=port_value,json=portValue,proto3,oneof" json:"port_value,omitempty"`
}
type SocketAddress_NamedPort struct {
	NamedPort string `protobuf:"bytes,4,opt,name=named_port,json=namedPort,proto3,oneof" json:"named_port,omitempty"`
}

func (*SocketAddress_PortValue) isSocketAddress_PortSpecifier() {}
func (*SocketAddress_NamedPort) isSocketAddress_PortSpecifier() {}

func (m *SocketAddress) GetPortSpecifier() isSocketAddress_PortSpecifier {
	if m != nil {
		return m.PortSpecifier
	}
	return nil
}

func (m *SocketAddress) GetProtocol() SocketAddress_Protocol {
	if m != nil {
		return m.Protocol
	}
	return SocketAddress_TCP
}

func (m *SocketAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *SocketAddress) GetPortValue() uint32 {
	if x, ok := m.GetPortSpecifier().(*SocketAddress_PortValue); ok {
		return x.PortValue
	}
	return 0
}

func (m *SocketAddress) GetNamedPort() string {
	if x, ok := m.GetPortSpecifier().(*SocketAddress_NamedPort); ok {
		return x.NamedPort
	}
	return ""
}

func (m *SocketAddress) GetResolverName() string {
	if m != nil {
		return m.ResolverName
	}
	return ""
}

func (m *SocketAddress) GetIpv4Compat() bool {
	if m != nil {
		return m.Ipv4Compat
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SocketAddress) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SocketAddress_PortValue)(nil),
		(*SocketAddress_NamedPort)(nil),
	}
}

type TcpKeepalive struct {
	// Maximum number of keepalive probes to send without response before deciding
	// the connection is dead. Default is to use the OS level configuration (unless
	// overridden, Linux defaults to 9.)
	KeepaliveProbes *types.UInt32Value `protobuf:"bytes,1,opt,name=keepalive_probes,json=keepaliveProbes,proto3" json:"keepalive_probes,omitempty"`
	// The number of seconds a connection needs to be idle before keep-alive probes
	// start being sent. Default is to use the OS level configuration (unless
	// overridden, Linux defaults to 7200s (i.e., 2 hours.)
	KeepaliveTime *types.UInt32Value `protobuf:"bytes,2,opt,name=keepalive_time,json=keepaliveTime,proto3" json:"keepalive_time,omitempty"`
	// The number of seconds between keep-alive probes. Default is to use the OS
	// level configuration (unless overridden, Linux defaults to 75s.)
	KeepaliveInterval    *types.UInt32Value `protobuf:"bytes,3,opt,name=keepalive_interval,json=keepaliveInterval,proto3" json:"keepalive_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TcpKeepalive) Reset()         { *m = TcpKeepalive{} }
func (m *TcpKeepalive) String() string { return proto.CompactTextString(m) }
func (*TcpKeepalive) ProtoMessage()    {}
func (*TcpKeepalive) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb356bd27ab3b7eb, []int{2}
}
func (m *TcpKeepalive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpKeepalive.Unmarshal(m, b)
}
func (m *TcpKeepalive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpKeepalive.Marshal(b, m, deterministic)
}
func (m *TcpKeepalive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpKeepalive.Merge(m, src)
}
func (m *TcpKeepalive) XXX_Size() int {
	return xxx_messageInfo_TcpKeepalive.Size(m)
}
func (m *TcpKeepalive) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpKeepalive.DiscardUnknown(m)
}

var xxx_messageInfo_TcpKeepalive proto.InternalMessageInfo

func (m *TcpKeepalive) GetKeepaliveProbes() *types.UInt32Value {
	if m != nil {
		return m.KeepaliveProbes
	}
	return nil
}

func (m *TcpKeepalive) GetKeepaliveTime() *types.UInt32Value {
	if m != nil {
		return m.KeepaliveTime
	}
	return nil
}

func (m *TcpKeepalive) GetKeepaliveInterval() *types.UInt32Value {
	if m != nil {
		return m.KeepaliveInterval
	}
	return nil
}

type BindConfig struct {
	// The address to bind to when creating a socket.
	SourceAddress *SocketAddress `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	// Whether to set the *IP_FREEBIND* option when creating the socket. When this
	// flag is set to true, allows the :ref:`source_address
	// <envoy_api_field_config.cluster.v3.UpstreamBindConfig.source_address>` to be an IP address
	// that is not configured on the system running Envoy. When this flag is set
	// to false, the option *IP_FREEBIND* is disabled on the socket. When this
	// flag is not set (default), the socket is not modified, i.e. the option is
	// neither enabled nor disabled.
	Freebind *types.BoolValue `protobuf:"bytes,2,opt,name=freebind,proto3" json:"freebind,omitempty"`
	// Additional socket options that may not be present in Envoy source code or
	// precompiled binaries.
	SocketOptions        []*SocketOption `protobuf:"bytes,3,rep,name=socket_options,json=socketOptions,proto3" json:"socket_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BindConfig) Reset()         { *m = BindConfig{} }
func (m *BindConfig) String() string { return proto.CompactTextString(m) }
func (*BindConfig) ProtoMessage()    {}
func (*BindConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb356bd27ab3b7eb, []int{3}
}
func (m *BindConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindConfig.Unmarshal(m, b)
}
func (m *BindConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindConfig.Marshal(b, m, deterministic)
}
func (m *BindConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindConfig.Merge(m, src)
}
func (m *BindConfig) XXX_Size() int {
	return xxx_messageInfo_BindConfig.Size(m)
}
func (m *BindConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_BindConfig.DiscardUnknown(m)
}

var xxx_messageInfo_BindConfig proto.InternalMessageInfo

func (m *BindConfig) GetSourceAddress() *SocketAddress {
	if m != nil {
		return m.SourceAddress
	}
	return nil
}

func (m *BindConfig) GetFreebind() *types.BoolValue {
	if m != nil {
		return m.Freebind
	}
	return nil
}

func (m *BindConfig) GetSocketOptions() []*SocketOption {
	if m != nil {
		return m.SocketOptions
	}
	return nil
}

// Addresses specify either a logical or physical address and port, which are
// used to tell Envoy where to bind/listen, connect to upstream and find
// management servers.
type Address struct {
	// Types that are valid to be assigned to Address:
	//	*Address_SocketAddress
	//	*Address_Pipe
	Address              isAddress_Address `protobuf_oneof:"address"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb356bd27ab3b7eb, []int{4}
}
func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

type isAddress_Address interface {
	isAddress_Address()
	Equal(interface{}) bool
}

type Address_SocketAddress struct {
	SocketAddress *SocketAddress `protobuf:"bytes,1,opt,name=socket_address,json=socketAddress,proto3,oneof" json:"socket_address,omitempty"`
}
type Address_Pipe struct {
	Pipe *Pipe `protobuf:"bytes,2,opt,name=pipe,proto3,oneof" json:"pipe,omitempty"`
}

func (*Address_SocketAddress) isAddress_Address() {}
func (*Address_Pipe) isAddress_Address()          {}

func (m *Address) GetAddress() isAddress_Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Address) GetSocketAddress() *SocketAddress {
	if x, ok := m.GetAddress().(*Address_SocketAddress); ok {
		return x.SocketAddress
	}
	return nil
}

func (m *Address) GetPipe() *Pipe {
	if x, ok := m.GetAddress().(*Address_Pipe); ok {
		return x.Pipe
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Address) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Address_SocketAddress)(nil),
		(*Address_Pipe)(nil),
	}
}

// CidrRange specifies an IP Address and a prefix length to construct
// the subnet mask for a `CIDR <https://tools.ietf.org/html/rfc4632>`_ range.
type CidrRange struct {
	// IPv4 or IPv6 address, e.g. ``192.0.0.0`` or ``2001:db8::``.
	AddressPrefix string `protobuf:"bytes,1,opt,name=address_prefix,json=addressPrefix,proto3" json:"address_prefix,omitempty"`
	// Length of prefix, e.g. 0, 32.
	PrefixLen            *types.UInt32Value `protobuf:"bytes,2,opt,name=prefix_len,json=prefixLen,proto3" json:"prefix_len,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CidrRange) Reset()         { *m = CidrRange{} }
func (m *CidrRange) String() string { return proto.CompactTextString(m) }
func (*CidrRange) ProtoMessage()    {}
func (*CidrRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb356bd27ab3b7eb, []int{5}
}
func (m *CidrRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CidrRange.Unmarshal(m, b)
}
func (m *CidrRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CidrRange.Marshal(b, m, deterministic)
}
func (m *CidrRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CidrRange.Merge(m, src)
}
func (m *CidrRange) XXX_Size() int {
	return xxx_messageInfo_CidrRange.Size(m)
}
func (m *CidrRange) XXX_DiscardUnknown() {
	xxx_messageInfo_CidrRange.DiscardUnknown(m)
}

var xxx_messageInfo_CidrRange proto.InternalMessageInfo

func (m *CidrRange) GetAddressPrefix() string {
	if m != nil {
		return m.AddressPrefix
	}
	return ""
}

func (m *CidrRange) GetPrefixLen() *types.UInt32Value {
	if m != nil {
		return m.PrefixLen
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.config.core.v3.SocketAddress_Protocol", SocketAddress_Protocol_name, SocketAddress_Protocol_value)
	proto.RegisterType((*Pipe)(nil), "envoy.config.core.v3.Pipe")
	proto.RegisterType((*SocketAddress)(nil), "envoy.config.core.v3.SocketAddress")
	proto.RegisterType((*TcpKeepalive)(nil), "envoy.config.core.v3.TcpKeepalive")
	proto.RegisterType((*BindConfig)(nil), "envoy.config.core.v3.BindConfig")
	proto.RegisterType((*Address)(nil), "envoy.config.core.v3.Address")
	proto.RegisterType((*CidrRange)(nil), "envoy.config.core.v3.CidrRange")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo-edge/projects/gloo/api/external/envoy/config/core/v3/address.proto", fileDescriptor_fb356bd27ab3b7eb)
}

var fileDescriptor_fb356bd27ab3b7eb = []byte{
	// 864 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x4d, 0x6f, 0xe3, 0x44,
	0x18, 0xae, 0x93, 0x6c, 0x9b, 0xbc, 0x6d, 0x42, 0x18, 0x2d, 0x60, 0xfa, 0x91, 0x4d, 0x5d, 0x81,
	0xa2, 0x0a, 0x6c, 0x94, 0x22, 0x0e, 0xb9, 0xad, 0x03, 0xa2, 0x55, 0x57, 0x10, 0x99, 0x96, 0x03,
	0x17, 0x33, 0xb1, 0x27, 0xde, 0xa1, 0x8e, 0x67, 0x34, 0x9e, 0x98, 0xee, 0x6d, 0xc5, 0x09, 0xad,
	0x84, 0xc4, 0x99, 0x1f, 0x80, 0xf8, 0x0d, 0x70, 0x46, 0xe2, 0xca, 0x5f, 0xe0, 0x47, 0x20, 0xd4,
	0x4b, 0xd1, 0x8c, 0x3f, 0xda, 0x25, 0x11, 0xed, 0xde, 0x66, 0xde, 0x8f, 0x67, 0x9e, 0xf7, 0xf1,
	0xe3, 0x19, 0x38, 0x8f, 0xa8, 0x7c, 0xba, 0x98, 0xda, 0x01, 0x9b, 0x3b, 0x29, 0x8b, 0xd9, 0xfb,
	0x94, 0x39, 0x51, 0xcc, 0x98, 0xc3, 0x05, 0xfb, 0x86, 0x04, 0x32, 0xcd, 0x77, 0x98, 0x53, 0x87,
	0x5c, 0x4a, 0x22, 0x12, 0x1c, 0x3b, 0x24, 0xc9, 0xd8, 0x33, 0x27, 0x60, 0xc9, 0x8c, 0x46, 0x4e,
	0xc0, 0x04, 0x71, 0xb2, 0x23, 0x07, 0x87, 0xa1, 0x20, 0x69, 0x6a, 0x73, 0xc1, 0x24, 0x43, 0x0f,
	0x75, 0x8d, 0x9d, 0xd7, 0xd8, 0xaa, 0xc6, 0xce, 0x8e, 0xb6, 0x07, 0x2b, 0x3b, 0x53, 0x16, 0x5c,
	0x10, 0xe9, 0x33, 0x2e, 0x29, 0x4b, 0xf2, 0xfe, 0xed, 0x5e, 0xc4, 0x58, 0x14, 0x13, 0x47, 0xef,
	0xa6, 0x8b, 0x99, 0xf3, 0xad, 0xc0, 0x9c, 0x13, 0x51, 0xe0, 0x6f, 0xef, 0x2d, 0x42, 0x8e, 0x1d,
	0x9c, 0x24, 0x4c, 0x62, 0xd5, 0x96, 0x3a, 0xa9, 0xc4, 0x72, 0x51, 0xa6, 0xf7, 0x97, 0xd2, 0x19,
	0x11, 0x29, 0x65, 0x09, 0x4d, 0xa2, 0xa2, 0xe4, 0xad, 0x0c, 0xc7, 0x34, 0xc4, 0x92, 0x38, 0xe5,
	0xa2, 0x48, 0x3c, 0x8c, 0x58, 0xc4, 0xf4, 0xd2, 0x51, 0xab, 0x3c, 0x6a, 0x7d, 0x0d, 0x8d, 0x09,
	0xe5, 0x04, 0xed, 0x40, 0x83, 0x63, 0xf9, 0xd4, 0x34, 0xfa, 0xc6, 0xa0, 0xe5, 0x6e, 0x5c, 0xb9,
	0x0d, 0x51, 0xeb, 0x1b, 0x9e, 0x0e, 0xa2, 0x5d, 0x68, 0xcc, 0x59, 0x48, 0xcc, 0x5a, 0xdf, 0x18,
	0xb4, 0xdd, 0xe6, 0x95, 0xfb, 0xe0, 0xb0, 0x6e, 0x5e, 0xd7, 0x3d, 0x1d, 0x1d, 0xed, 0xfd, 0xf4,
	0xfb, 0xf7, 0x3d, 0x13, 0xde, 0xcc, 0xa5, 0xc1, 0x9c, 0xda, 0xd9, 0x30, 0x97, 0x46, 0x21, 0x5b,
	0x7f, 0xd7, 0xa0, 0xfd, 0x85, 0x96, 0xe2, 0x71, 0x2e, 0x25, 0xf2, 0xa0, 0xa9, 0x0f, 0x0f, 0x58,
	0xac, 0xcf, 0xeb, 0x0c, 0xdf, 0xb3, 0x57, 0xe9, 0x6a, 0xbf, 0xd4, 0x66, 0x4f, 0x8a, 0x1e, 0x4d,
	0xe0, 0x3b, 0xa3, 0xd6, 0x35, 0xbc, 0x0a, 0x07, 0xed, 0xc3, 0x46, 0xf1, 0xa5, 0x34, 0xcb, 0x5b,
	0x23, 0x94, 0x71, 0x74, 0x08, 0xc0, 0x99, 0x90, 0x7e, 0x86, 0xe3, 0x05, 0x31, 0xeb, 0x7a, 0x96,
	0xd6, 0x95, 0xbb, 0x7e, 0xd8, 0x30, 0xaf, 0xaf, 0xeb, 0xc7, 0x6b, 0x5e, 0x4b, 0xa5, 0xbf, 0x54,
	0x59, 0xf4, 0x08, 0x20, 0xc1, 0x73, 0x12, 0xfa, 0x2a, 0x64, 0x36, 0x14, 0xa2, 0x2a, 0xd0, 0xb1,
	0x09, 0x13, 0x12, 0x1d, 0x40, 0x5b, 0x90, 0x94, 0xc5, 0x19, 0x11, 0xbe, 0x8a, 0x9a, 0x0f, 0x54,
	0x8d, 0xb7, 0x55, 0x06, 0x3f, 0xc3, 0x73, 0x85, 0xb2, 0x49, 0x79, 0xf6, 0xa1, 0x1f, 0xb0, 0x39,
	0xc7, 0xd2, 0x5c, 0xef, 0x1b, 0x83, 0xa6, 0x07, 0x2a, 0x34, 0xd6, 0x11, 0x6b, 0x17, 0x9a, 0xe5,
	0x54, 0x68, 0x03, 0xea, 0x67, 0xe3, 0x49, 0x77, 0x4d, 0x2d, 0xce, 0x3f, 0x9e, 0x74, 0x8d, 0xd1,
	0xbb, 0x4a, 0xd8, 0x7d, 0x78, 0xb4, 0x2c, 0xec, 0x4b, 0xc2, 0xb8, 0x6f, 0x40, 0x47, 0x0f, 0x96,
	0x72, 0x12, 0xd0, 0x19, 0x25, 0x02, 0xd5, 0xff, 0x71, 0x0d, 0xeb, 0xc7, 0x1a, 0x6c, 0x9d, 0x05,
	0xfc, 0x94, 0x10, 0x8e, 0x63, 0x9a, 0x11, 0xf4, 0x29, 0x74, 0x2f, 0xca, 0x8d, 0xcf, 0x05, 0x9b,
	0x92, 0x54, 0xeb, 0xbf, 0x39, 0xdc, 0xb5, 0x73, 0x5f, 0xda, 0xa5, 0x2f, 0xed, 0xf3, 0x93, 0x44,
	0x1e, 0x0d, 0xb5, 0x18, 0xde, 0x6b, 0x55, 0xd7, 0x44, 0x37, 0xa1, 0x31, 0x74, 0x6e, 0x80, 0x24,
	0x9d, 0xe7, 0xce, 0xb8, 0x0b, 0xa6, 0x5d, 0xf5, 0x9c, 0xd1, 0x39, 0x41, 0xa7, 0x80, 0x6e, 0x40,
	0x68, 0x22, 0x89, 0xc8, 0x70, 0xac, 0x3f, 0xcb, 0x5d, 0x40, 0xaf, 0x57, 0x7d, 0x27, 0x45, 0xdb,
	0xe8, 0x1d, 0x25, 0x55, 0x1f, 0x7a, 0xcb, 0x52, 0xdd, 0x56, 0xc0, 0x7a, 0x51, 0x03, 0x70, 0x69,
	0x12, 0x8e, 0xb5, 0xcf, 0xd0, 0x19, 0x74, 0x52, 0xb6, 0x10, 0x01, 0xf1, 0x4b, 0xef, 0xe4, 0x72,
	0x1c, 0xdc, 0xc3, 0x8e, 0xda, 0x85, 0x2f, 0xb4, 0x0b, 0xdb, 0x39, 0x48, 0x69, 0xef, 0x8f, 0xa0,
	0x39, 0x13, 0x84, 0x4c, 0x69, 0x12, 0x16, 0xba, 0x6c, 0x2f, 0x8d, 0xe3, 0x32, 0x16, 0xe7, 0xc3,
	0x54, 0xb5, 0xe8, 0x44, 0xb1, 0xb9, 0x75, 0x65, 0xa4, 0x66, 0xbd, 0x5f, 0x1f, 0x6c, 0x0e, 0xad,
	0xff, 0x63, 0xf3, 0xb9, 0x2e, 0x55, 0x14, 0x6e, 0x76, 0xe9, 0xe8, 0x40, 0xc9, 0xd1, 0x83, 0xdd,
	0x65, 0x39, 0x6e, 0xa6, 0xb7, 0x7e, 0x33, 0x60, 0xa3, 0xe4, 0xfc, 0xa4, 0x3a, 0xfb, 0xd5, 0x95,
	0x38, 0x5e, 0x2b, 0x8f, 0x2f, 0xd1, 0x3e, 0x80, 0x06, 0xa7, 0x9c, 0x54, 0xd3, 0xaf, 0xc4, 0x50,
	0x97, 0xc3, 0xf1, 0x9a, 0xa7, 0x2b, 0x47, 0x7d, 0x45, 0x78, 0x07, 0xde, 0x5e, 0x26, 0x5c, 0xca,
	0xdd, 0xa9, 0x7e, 0xf0, 0xdc, 0xdd, 0x3f, 0x1b, 0xd0, 0x1a, 0xd3, 0x50, 0x78, 0x38, 0x89, 0x08,
	0xb2, 0xa1, 0x53, 0x64, 0x7d, 0x2e, 0xc8, 0x8c, 0x5e, 0xfe, 0xf7, 0x22, 0x6b, 0x17, 0xe9, 0x89,
	0xce, 0xa2, 0x4f, 0x00, 0xf2, 0x3a, 0x3f, 0x26, 0xc9, 0x7d, 0xdc, 0x5b, 0xde, 0x7a, 0xcf, 0x0d,
	0xaf, 0x95, 0x77, 0x3e, 0x21, 0xc9, 0xc8, 0x52, 0xb4, 0xf7, 0x60, 0x67, 0x99, 0x76, 0x45, 0xcd,
	0xfd, 0xc1, 0xf8, 0xe5, 0xaf, 0x9e, 0xf1, 0xeb, 0xf3, 0x3f, 0xfe, 0x5c, 0xaf, 0x75, 0x6b, 0x60,
	0x51, 0x96, 0xeb, 0xc1, 0x05, 0xbb, 0x7c, 0xb6, 0x52, 0x1a, 0x77, 0xeb, 0x71, 0x49, 0x96, 0x49,
	0x36, 0x31, 0xbe, 0x3a, 0xbd, 0xdf, 0x63, 0xc6, 0x2f, 0xa2, 0xbb, 0x1f, 0xb4, 0xe9, 0xba, 0x1e,
	0xef, 0xe8, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd3, 0xfb, 0xa0, 0x48, 0x22, 0x07, 0x00, 0x00,
}

func (this *Pipe) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Pipe)
	if !ok {
		that2, ok := that.(Pipe)
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
	if this.Path != that1.Path {
		return false
	}
	if this.Mode != that1.Mode {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *SocketAddress) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SocketAddress)
	if !ok {
		that2, ok := that.(SocketAddress)
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
	if this.Protocol != that1.Protocol {
		return false
	}
	if this.Address != that1.Address {
		return false
	}
	if that1.PortSpecifier == nil {
		if this.PortSpecifier != nil {
			return false
		}
	} else if this.PortSpecifier == nil {
		return false
	} else if !this.PortSpecifier.Equal(that1.PortSpecifier) {
		return false
	}
	if this.ResolverName != that1.ResolverName {
		return false
	}
	if this.Ipv4Compat != that1.Ipv4Compat {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *SocketAddress_PortValue) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SocketAddress_PortValue)
	if !ok {
		that2, ok := that.(SocketAddress_PortValue)
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
	if this.PortValue != that1.PortValue {
		return false
	}
	return true
}
func (this *SocketAddress_NamedPort) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SocketAddress_NamedPort)
	if !ok {
		that2, ok := that.(SocketAddress_NamedPort)
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
	if this.NamedPort != that1.NamedPort {
		return false
	}
	return true
}
func (this *TcpKeepalive) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TcpKeepalive)
	if !ok {
		that2, ok := that.(TcpKeepalive)
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
	if !this.KeepaliveProbes.Equal(that1.KeepaliveProbes) {
		return false
	}
	if !this.KeepaliveTime.Equal(that1.KeepaliveTime) {
		return false
	}
	if !this.KeepaliveInterval.Equal(that1.KeepaliveInterval) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *BindConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*BindConfig)
	if !ok {
		that2, ok := that.(BindConfig)
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
	if !this.SourceAddress.Equal(that1.SourceAddress) {
		return false
	}
	if !this.Freebind.Equal(that1.Freebind) {
		return false
	}
	if len(this.SocketOptions) != len(that1.SocketOptions) {
		return false
	}
	for i := range this.SocketOptions {
		if !this.SocketOptions[i].Equal(that1.SocketOptions[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Address) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Address)
	if !ok {
		that2, ok := that.(Address)
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
	if that1.Address == nil {
		if this.Address != nil {
			return false
		}
	} else if this.Address == nil {
		return false
	} else if !this.Address.Equal(that1.Address) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Address_SocketAddress) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Address_SocketAddress)
	if !ok {
		that2, ok := that.(Address_SocketAddress)
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
	if !this.SocketAddress.Equal(that1.SocketAddress) {
		return false
	}
	return true
}
func (this *Address_Pipe) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Address_Pipe)
	if !ok {
		that2, ok := that.(Address_Pipe)
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
	if !this.Pipe.Equal(that1.Pipe) {
		return false
	}
	return true
}
func (this *CidrRange) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CidrRange)
	if !ok {
		that2, ok := that.(CidrRange)
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
	if this.AddressPrefix != that1.AddressPrefix {
		return false
	}
	if !this.PrefixLen.Equal(that1.PrefixLen) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
