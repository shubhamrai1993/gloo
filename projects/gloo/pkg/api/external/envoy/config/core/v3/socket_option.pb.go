// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo-edge/projects/gloo/api/external/envoy/config/core/v3/socket_option.proto

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

type SocketOption_SocketState int32

const (
	// Socket options are applied after socket creation but before binding the socket to a port
	SocketOption_STATE_PREBIND SocketOption_SocketState = 0
	// Socket options are applied after binding the socket to a port but before calling listen()
	SocketOption_STATE_BOUND SocketOption_SocketState = 1
	// Socket options are applied after calling listen()
	SocketOption_STATE_LISTENING SocketOption_SocketState = 2
)

var SocketOption_SocketState_name = map[int32]string{
	0: "STATE_PREBIND",
	1: "STATE_BOUND",
	2: "STATE_LISTENING",
}

var SocketOption_SocketState_value = map[string]int32{
	"STATE_PREBIND":   0,
	"STATE_BOUND":     1,
	"STATE_LISTENING": 2,
}

func (x SocketOption_SocketState) String() string {
	return proto.EnumName(SocketOption_SocketState_name, int32(x))
}

func (SocketOption_SocketState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ff80f4c8bc804c53, []int{0, 0}
}

// Generic socket option message. This would be used to set socket options that
// might not exist in upstream kernels or precompiled Envoy binaries.
// [#next-free-field: 7]
type SocketOption struct {
	// An optional name to give this socket option for debugging, etc.
	// Uniqueness is not required and no special meaning is assumed.
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	// Corresponding to the level value passed to setsockopt, such as IPPROTO_TCP
	Level int64 `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	// The numeric name as passed to setsockopt
	Name int64 `protobuf:"varint,3,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*SocketOption_IntValue
	//	*SocketOption_BufValue
	Value isSocketOption_Value `protobuf_oneof:"value"`
	// The state in which the option will be applied. When used in BindConfig
	// STATE_PREBIND is currently the only valid value.
	State                SocketOption_SocketState `protobuf:"varint,6,opt,name=state,proto3,enum=envoy.config.core.v3.SocketOption_SocketState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *SocketOption) Reset()         { *m = SocketOption{} }
func (m *SocketOption) String() string { return proto.CompactTextString(m) }
func (*SocketOption) ProtoMessage()    {}
func (*SocketOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff80f4c8bc804c53, []int{0}
}
func (m *SocketOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocketOption.Unmarshal(m, b)
}
func (m *SocketOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocketOption.Marshal(b, m, deterministic)
}
func (m *SocketOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocketOption.Merge(m, src)
}
func (m *SocketOption) XXX_Size() int {
	return xxx_messageInfo_SocketOption.Size(m)
}
func (m *SocketOption) XXX_DiscardUnknown() {
	xxx_messageInfo_SocketOption.DiscardUnknown(m)
}

var xxx_messageInfo_SocketOption proto.InternalMessageInfo

type isSocketOption_Value interface {
	isSocketOption_Value()
	Equal(interface{}) bool
}

type SocketOption_IntValue struct {
	IntValue int64 `protobuf:"varint,4,opt,name=int_value,json=intValue,proto3,oneof" json:"int_value,omitempty"`
}
type SocketOption_BufValue struct {
	BufValue []byte `protobuf:"bytes,5,opt,name=buf_value,json=bufValue,proto3,oneof" json:"buf_value,omitempty"`
}

func (*SocketOption_IntValue) isSocketOption_Value() {}
func (*SocketOption_BufValue) isSocketOption_Value() {}

func (m *SocketOption) GetValue() isSocketOption_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *SocketOption) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SocketOption) GetLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *SocketOption) GetName() int64 {
	if m != nil {
		return m.Name
	}
	return 0
}

func (m *SocketOption) GetIntValue() int64 {
	if x, ok := m.GetValue().(*SocketOption_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (m *SocketOption) GetBufValue() []byte {
	if x, ok := m.GetValue().(*SocketOption_BufValue); ok {
		return x.BufValue
	}
	return nil
}

func (m *SocketOption) GetState() SocketOption_SocketState {
	if m != nil {
		return m.State
	}
	return SocketOption_STATE_PREBIND
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SocketOption) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SocketOption_IntValue)(nil),
		(*SocketOption_BufValue)(nil),
	}
}

func init() {
	proto.RegisterEnum("envoy.config.core.v3.SocketOption_SocketState", SocketOption_SocketState_name, SocketOption_SocketState_value)
	proto.RegisterType((*SocketOption)(nil), "envoy.config.core.v3.SocketOption")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo-edge/projects/gloo/api/external/envoy/config/core/v3/socket_option.proto", fileDescriptor_ff80f4c8bc804c53)
}

var fileDescriptor_ff80f4c8bc804c53 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x77, 0xd2, 0x6d, 0xd9, 0x9d, 0x56, 0x37, 0x3b, 0x16, 0x0c, 0x0b, 0x5b, 0x62, 0x41,
	0xe8, 0xc5, 0x0c, 0x6c, 0x6f, 0xde, 0x0c, 0x5b, 0xb5, 0x28, 0xd9, 0x92, 0x56, 0x0f, 0x7a, 0x28,
	0xd3, 0x74, 0x1a, 0xc7, 0xcd, 0xce, 0x0b, 0xc9, 0x24, 0xec, 0xde, 0xc4, 0x93, 0x67, 0x8f, 0x7e,
	0x02, 0x3f, 0x83, 0x77, 0xc1, 0xab, 0x5f, 0x41, 0xf0, 0x3b, 0x88, 0x27, 0x99, 0x99, 0x0a, 0x01,
	0x17, 0xdc, 0xdb, 0x7b, 0xff, 0xdf, 0x7f, 0x1e, 0xff, 0x97, 0x3c, 0xfc, 0x3a, 0x15, 0xea, 0x4d,
	0xb5, 0x0a, 0x12, 0xb8, 0xa0, 0x25, 0x64, 0xf0, 0x40, 0x00, 0x4d, 0x33, 0x00, 0x9a, 0x17, 0xf0,
	0x96, 0x27, 0xaa, 0xb4, 0x1d, 0xcb, 0x05, 0xe5, 0x97, 0x8a, 0x17, 0x92, 0x65, 0x94, 0xcb, 0x1a,
	0xae, 0x68, 0x02, 0x72, 0x23, 0x52, 0x9a, 0x40, 0xc1, 0x69, 0x3d, 0xa6, 0x25, 0x24, 0xe7, 0x5c,
	0x2d, 0x21, 0x57, 0x02, 0x64, 0x90, 0x17, 0xa0, 0x80, 0xf4, 0x8d, 0x33, 0xb0, 0xce, 0x40, 0x3b,
	0x83, 0x7a, 0x7c, 0x74, 0x5c, 0xad, 0x73, 0x46, 0x99, 0x94, 0xa0, 0x98, 0x36, 0x97, 0xb4, 0x54,
	0x4c, 0x55, 0xa5, 0x7d, 0x74, 0x74, 0xef, 0x1f, 0x5c, 0xf3, 0xa2, 0x14, 0x20, 0x85, 0x4c, 0xb7,
	0x96, 0xbb, 0x35, 0xcb, 0xc4, 0x9a, 0x29, 0x4e, 0xff, 0x16, 0x5b, 0xd0, 0x4f, 0x21, 0x05, 0x53,
	0x52, 0x5d, 0x59, 0x75, 0xf8, 0xd3, 0xc1, 0xbd, 0xb9, 0x89, 0x77, 0x66, 0xd2, 0x11, 0x1f, 0x77,
	0xd7, 0xbc, 0x4c, 0x0a, 0x61, 0x5a, 0x0f, 0xf9, 0x68, 0xb4, 0x1f, 0x37, 0x25, 0xd2, 0xc7, 0xed,
	0x8c, 0xd7, 0x3c, 0xf3, 0x1c, 0x1f, 0x8d, 0x5a, 0xb1, 0x6d, 0x08, 0xc1, 0xbb, 0x92, 0x5d, 0x70,
	0xaf, 0x65, 0x44, 0x53, 0x93, 0x63, 0xbc, 0x2f, 0xa4, 0x5a, 0xd6, 0x2c, 0xab, 0xb8, 0xb7, 0xab,
	0xc1, 0xd3, 0x9d, 0x78, 0x4f, 0x48, 0xf5, 0x52, 0x2b, 0x1a, 0xaf, 0xaa, 0xcd, 0x16, 0xb7, 0x7d,
	0x34, 0xea, 0x69, 0xbc, 0xaa, 0x36, 0x16, 0x47, 0xb8, 0xad, 0x97, 0xe7, 0x5e, 0xc7, 0x47, 0xa3,
	0xdb, 0x27, 0x41, 0x70, 0xdd, 0x17, 0x0b, 0x9a, 0xe1, 0xb7, 0xcd, 0x5c, 0xbf, 0x0a, 0xf7, 0x7e,
	0x87, 0xed, 0xf7, 0xc8, 0x71, 0x51, 0x6c, 0xc7, 0x0c, 0x1f, 0xe3, 0x6e, 0x83, 0x93, 0x43, 0x7c,
	0x6b, 0xbe, 0x78, 0xb4, 0x98, 0x2c, 0x67, 0xf1, 0x24, 0x9c, 0x46, 0xa7, 0xee, 0x0e, 0x39, 0xc0,
	0x5d, 0x2b, 0x85, 0x67, 0x2f, 0xa2, 0x53, 0x17, 0x91, 0x3b, 0xf8, 0xc0, 0x0a, 0xcf, 0xa7, 0xf3,
	0xc5, 0x24, 0x9a, 0x46, 0x4f, 0x5c, 0xe7, 0xe1, 0xfd, 0x4f, 0x5f, 0x3f, 0x0c, 0x7c, 0x3c, 0xb0,
	0x71, 0x58, 0x2e, 0x82, 0xfa, 0xc4, 0xc6, 0x69, 0x66, 0x09, 0x7b, 0xb8, 0x6d, 0x36, 0x23, 0xad,
	0x5f, 0x21, 0x0a, 0x3f, 0xa2, 0xcf, 0x3f, 0x06, 0xe8, 0xcb, 0xbb, 0x6f, 0xdf, 0x3b, 0x8e, 0xeb,
	0xe0, 0xa1, 0x00, 0xbb, 0x4e, 0x5e, 0xc0, 0xe5, 0xd5, 0xb5, 0x9b, 0x85, 0x87, 0xcd, 0x71, 0x33,
	0xfd, 0xb7, 0x66, 0xe8, 0xd5, 0xb3, 0x9b, 0xdd, 0x64, 0x7e, 0x9e, 0xfe, 0xff, 0x2e, 0x57, 0x1d,
	0x73, 0x03, 0xe3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x46, 0xbf, 0x54, 0x98, 0xe9, 0x02, 0x00,
	0x00,
}

func (this *SocketOption) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SocketOption)
	if !ok {
		that2, ok := that.(SocketOption)
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
	if this.Description != that1.Description {
		return false
	}
	if this.Level != that1.Level {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if that1.Value == nil {
		if this.Value != nil {
			return false
		}
	} else if this.Value == nil {
		return false
	} else if !this.Value.Equal(that1.Value) {
		return false
	}
	if this.State != that1.State {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *SocketOption_IntValue) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SocketOption_IntValue)
	if !ok {
		that2, ok := that.(SocketOption_IntValue)
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
	if this.IntValue != that1.IntValue {
		return false
	}
	return true
}
func (this *SocketOption_BufValue) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SocketOption_BufValue)
	if !ok {
		that2, ok := that.(SocketOption_BufValue)
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
	if !bytes.Equal(this.BufValue, that1.BufValue) {
		return false
	}
	return true
}
