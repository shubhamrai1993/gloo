// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo-edge/projects/gloo/api/v1/options/grpc/grpc.proto

package grpc

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	transformation "github.com/solo-io/gloo-edge/projects/gloo/pkg/api/v1/options/transformation"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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

// Service spec describing GRPC upstreams. This will usually be filled
// automatically via function discovery (if the upstream supports reflection).
// If your upstream service is a GRPC service, use this service spec (an empty
// spec is fine), to make sure that traffic to it is routed with http2.
type ServiceSpec struct {
	// Descriptors that contain information of the services listed below.
	// this is a serialized google.protobuf.FileDescriptorSet
	Descriptors []byte `protobuf:"bytes,1,opt,name=descriptors,proto3" json:"descriptors,omitempty"`
	// List of services used by this upstream. For a grpc upstream where you don't
	// need to use Gloo's function routing, this can be an empty list. These
	// services must be present in the descriptors.
	GrpcServices         []*ServiceSpec_GrpcService `protobuf:"bytes,2,rep,name=grpc_services,json=grpcServices,proto3" json:"grpc_services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ServiceSpec) Reset()         { *m = ServiceSpec{} }
func (m *ServiceSpec) String() string { return proto.CompactTextString(m) }
func (*ServiceSpec) ProtoMessage()    {}
func (*ServiceSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bddd1d7957d358a, []int{0}
}
func (m *ServiceSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceSpec.Unmarshal(m, b)
}
func (m *ServiceSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceSpec.Marshal(b, m, deterministic)
}
func (m *ServiceSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceSpec.Merge(m, src)
}
func (m *ServiceSpec) XXX_Size() int {
	return xxx_messageInfo_ServiceSpec.Size(m)
}
func (m *ServiceSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceSpec proto.InternalMessageInfo

func (m *ServiceSpec) GetDescriptors() []byte {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

func (m *ServiceSpec) GetGrpcServices() []*ServiceSpec_GrpcService {
	if m != nil {
		return m.GrpcServices
	}
	return nil
}

// Describes a grpc service
type ServiceSpec_GrpcService struct {
	// The package of this service.
	PackageName string `protobuf:"bytes,1,opt,name=package_name,json=packageName,proto3" json:"package_name,omitempty"`
	// The service name of this service.
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// The functions available in this service.
	FunctionNames        []string `protobuf:"bytes,3,rep,name=function_names,json=functionNames,proto3" json:"function_names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceSpec_GrpcService) Reset()         { *m = ServiceSpec_GrpcService{} }
func (m *ServiceSpec_GrpcService) String() string { return proto.CompactTextString(m) }
func (*ServiceSpec_GrpcService) ProtoMessage()    {}
func (*ServiceSpec_GrpcService) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bddd1d7957d358a, []int{0, 0}
}
func (m *ServiceSpec_GrpcService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceSpec_GrpcService.Unmarshal(m, b)
}
func (m *ServiceSpec_GrpcService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceSpec_GrpcService.Marshal(b, m, deterministic)
}
func (m *ServiceSpec_GrpcService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceSpec_GrpcService.Merge(m, src)
}
func (m *ServiceSpec_GrpcService) XXX_Size() int {
	return xxx_messageInfo_ServiceSpec_GrpcService.Size(m)
}
func (m *ServiceSpec_GrpcService) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceSpec_GrpcService.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceSpec_GrpcService proto.InternalMessageInfo

func (m *ServiceSpec_GrpcService) GetPackageName() string {
	if m != nil {
		return m.PackageName
	}
	return ""
}

func (m *ServiceSpec_GrpcService) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *ServiceSpec_GrpcService) GetFunctionNames() []string {
	if m != nil {
		return m.FunctionNames
	}
	return nil
}

// This is only for upstream with Grpc service spec.
type DestinationSpec struct {
	// The proto package of the function.
	Package string `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`
	// The name of the service of the function.
	Service string `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	// The name of the function.
	Function string `protobuf:"bytes,3,opt,name=function,proto3" json:"function,omitempty"`
	// Parameters describe how to extract the function parameters from the
	// request.
	Parameters           *transformation.Parameters `protobuf:"bytes,4,opt,name=parameters,proto3" json:"parameters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *DestinationSpec) Reset()         { *m = DestinationSpec{} }
func (m *DestinationSpec) String() string { return proto.CompactTextString(m) }
func (*DestinationSpec) ProtoMessage()    {}
func (*DestinationSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bddd1d7957d358a, []int{1}
}
func (m *DestinationSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestinationSpec.Unmarshal(m, b)
}
func (m *DestinationSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestinationSpec.Marshal(b, m, deterministic)
}
func (m *DestinationSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestinationSpec.Merge(m, src)
}
func (m *DestinationSpec) XXX_Size() int {
	return xxx_messageInfo_DestinationSpec.Size(m)
}
func (m *DestinationSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_DestinationSpec.DiscardUnknown(m)
}

var xxx_messageInfo_DestinationSpec proto.InternalMessageInfo

func (m *DestinationSpec) GetPackage() string {
	if m != nil {
		return m.Package
	}
	return ""
}

func (m *DestinationSpec) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *DestinationSpec) GetFunction() string {
	if m != nil {
		return m.Function
	}
	return ""
}

func (m *DestinationSpec) GetParameters() *transformation.Parameters {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func init() {
	proto.RegisterType((*ServiceSpec)(nil), "grpc.options.gloo.solo.io.ServiceSpec")
	proto.RegisterType((*ServiceSpec_GrpcService)(nil), "grpc.options.gloo.solo.io.ServiceSpec.GrpcService")
	proto.RegisterType((*DestinationSpec)(nil), "grpc.options.gloo.solo.io.DestinationSpec")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo-edge/projects/gloo/api/v1/options/grpc/grpc.proto", fileDescriptor_3bddd1d7957d358a)
}

var fileDescriptor_3bddd1d7957d358a = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x55, 0x9a, 0x0a, 0xa8, 0xd3, 0x82, 0x14, 0x71, 0x08, 0x39, 0xa0, 0x50, 0x09, 0x29, 0x17,
	0x6c, 0x51, 0xce, 0x1c, 0x40, 0x95, 0x7a, 0x03, 0x94, 0x1e, 0x90, 0xb8, 0x54, 0xae, 0x71, 0x8d,
	0x69, 0x93, 0xb1, 0x6c, 0xb7, 0xaa, 0x38, 0xf3, 0x31, 0x9c, 0x39, 0xed, 0xf7, 0xec, 0x3f, 0xec,
	0x7d, 0x65, 0xc7, 0x69, 0xb3, 0x52, 0x57, 0xbb, 0x97, 0x68, 0xde, 0x9b, 0x37, 0xf3, 0x26, 0xe3,
	0x41, 0x73, 0x21, 0xed, 0xaf, 0xfd, 0x1a, 0x33, 0xa8, 0x89, 0x81, 0x1d, 0xbc, 0x93, 0x40, 0xc4,
	0x0e, 0x80, 0x28, 0x0d, 0xbf, 0x39, 0xb3, 0xa6, 0x45, 0x54, 0x49, 0x72, 0x78, 0x4f, 0x40, 0x59,
	0x09, 0x8d, 0x21, 0x42, 0x2b, 0xe6, 0x3f, 0x58, 0x69, 0xb0, 0x90, 0xbe, 0xf2, 0x71, 0xc8, 0x62,
	0x57, 0x81, 0x5d, 0x33, 0x2c, 0x21, 0x7f, 0x29, 0x40, 0x80, 0x57, 0x11, 0x17, 0xb5, 0x05, 0x79,
	0xca, 0x8f, 0xb6, 0x25, 0xf9, 0xd1, 0x06, 0xee, 0xd3, 0xc3, 0xbe, 0x56, 0xd3, 0xc6, 0x6c, 0x40,
	0xd7, 0xd4, 0x61, 0xa2, 0xa8, 0xa6, 0x35, 0xb7, 0x5c, 0x9b, 0xb6, 0xc5, 0xf4, 0xef, 0x00, 0x25,
	0x4b, 0xae, 0x0f, 0x92, 0xf1, 0xa5, 0xe2, 0x2c, 0x2d, 0x50, 0xf2, 0x93, 0x1b, 0xa6, 0xa5, 0xb2,
	0xa0, 0x4d, 0x16, 0x15, 0x51, 0x39, 0xae, 0xfa, 0x54, 0xfa, 0x1d, 0x4d, 0xdc, 0xec, 0x2b, 0xd3,
	0x56, 0x99, 0x6c, 0x50, 0xc4, 0x65, 0x32, 0x9b, 0xe1, 0x7b, 0xff, 0x08, 0xf7, 0x0c, 0xf0, 0x42,
	0x2b, 0x16, 0x70, 0x35, 0x16, 0x67, 0x60, 0xf2, 0x3f, 0x28, 0xe9, 0x25, 0xd3, 0x37, 0x68, 0xac,
	0x28, 0xdb, 0x52, 0xc1, 0x57, 0x0d, 0xad, 0xb9, 0x1f, 0x65, 0x54, 0x25, 0x81, 0xfb, 0x42, 0x6b,
	0x2f, 0x09, 0x53, 0xb4, 0x92, 0x41, 0x2b, 0x09, 0x9c, 0x97, 0xbc, 0x45, 0xcf, 0x37, 0xfb, 0x86,
	0xb9, 0xa1, 0xbc, 0xc6, 0x64, 0x71, 0x11, 0x97, 0xa3, 0x6a, 0xd2, 0xb1, 0x4e, 0x65, 0xa6, 0xff,
	0x23, 0xf4, 0x62, 0xce, 0x8d, 0x95, 0x8d, 0xdf, 0x93, 0x5f, 0x45, 0x86, 0x9e, 0x06, 0xb3, 0xe0,
	0xdd, 0x41, 0x97, 0x09, 0x1e, 0xc1, 0xb2, 0x83, 0x69, 0x8e, 0x9e, 0x75, 0x8d, 0xb3, 0xd8, 0xa7,
	0x4e, 0x38, 0xfd, 0x8a, 0xd0, 0x79, 0xfd, 0xd9, 0xb0, 0x88, 0xca, 0x64, 0x46, 0xf0, 0xdd, 0x07,
	0xba, 0xbc, 0xbf, 0x6f, 0xa7, 0xb2, 0xaa, 0xd7, 0xe2, 0xf3, 0xe2, 0xea, 0x66, 0x18, 0xfd, 0xbb,
	0x7e, 0x1d, 0xfd, 0xf8, 0xf8, 0xb8, 0x9b, 0x54, 0x5b, 0x71, 0xe9, 0x2e, 0xd7, 0x4f, 0xfc, 0x2d,
	0x7c, 0xb8, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x57, 0xeb, 0xa3, 0xdb, 0x02, 0x00, 0x00,
}

func (this *ServiceSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSpec)
	if !ok {
		that2, ok := that.(ServiceSpec)
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
	if !bytes.Equal(this.Descriptors, that1.Descriptors) {
		return false
	}
	if len(this.GrpcServices) != len(that1.GrpcServices) {
		return false
	}
	for i := range this.GrpcServices {
		if !this.GrpcServices[i].Equal(that1.GrpcServices[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ServiceSpec_GrpcService) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSpec_GrpcService)
	if !ok {
		that2, ok := that.(ServiceSpec_GrpcService)
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
	if this.PackageName != that1.PackageName {
		return false
	}
	if this.ServiceName != that1.ServiceName {
		return false
	}
	if len(this.FunctionNames) != len(that1.FunctionNames) {
		return false
	}
	for i := range this.FunctionNames {
		if this.FunctionNames[i] != that1.FunctionNames[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *DestinationSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DestinationSpec)
	if !ok {
		that2, ok := that.(DestinationSpec)
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
	if this.Package != that1.Package {
		return false
	}
	if this.Service != that1.Service {
		return false
	}
	if this.Function != that1.Function {
		return false
	}
	if !this.Parameters.Equal(that1.Parameters) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
