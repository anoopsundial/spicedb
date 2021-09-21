// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: servok/api/v1/v1.proto

package apiv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to RequestTypeOneof:
	//	*WatchRequest_Srv
	RequestTypeOneof isWatchRequest_RequestTypeOneof `protobuf_oneof:"request_type_oneof"`
}

func (x *WatchRequest) Reset() {
	*x = WatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servok_api_v1_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchRequest) ProtoMessage() {}

func (x *WatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_servok_api_v1_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchRequest.ProtoReflect.Descriptor instead.
func (*WatchRequest) Descriptor() ([]byte, []int) {
	return file_servok_api_v1_v1_proto_rawDescGZIP(), []int{0}
}

func (m *WatchRequest) GetRequestTypeOneof() isWatchRequest_RequestTypeOneof {
	if m != nil {
		return m.RequestTypeOneof
	}
	return nil
}

func (x *WatchRequest) GetSrv() *WatchRequest_SRVRequest {
	if x, ok := x.GetRequestTypeOneof().(*WatchRequest_Srv); ok {
		return x.Srv
	}
	return nil
}

type isWatchRequest_RequestTypeOneof interface {
	isWatchRequest_RequestTypeOneof()
}

type WatchRequest_Srv struct {
	Srv *WatchRequest_SRVRequest `protobuf:"bytes,1,opt,name=srv,proto3,oneof"`
}

func (*WatchRequest_Srv) isWatchRequest_RequestTypeOneof() {}

type WatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoints []*Endpoint `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
}

func (x *WatchResponse) Reset() {
	*x = WatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servok_api_v1_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchResponse) ProtoMessage() {}

func (x *WatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_servok_api_v1_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchResponse.ProtoReflect.Descriptor instead.
func (*WatchResponse) Descriptor() ([]byte, []int) {
	return file_servok_api_v1_v1_proto_rawDescGZIP(), []int{1}
}

func (x *WatchResponse) GetEndpoints() []*Endpoint {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

type Endpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Port     uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Weight   uint32 `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *Endpoint) Reset() {
	*x = Endpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servok_api_v1_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endpoint) ProtoMessage() {}

func (x *Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_servok_api_v1_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endpoint.ProtoReflect.Descriptor instead.
func (*Endpoint) Descriptor() ([]byte, []int) {
	return file_servok_api_v1_v1_proto_rawDescGZIP(), []int{2}
}

func (x *Endpoint) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *Endpoint) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Endpoint) GetWeight() uint32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

type WatchRequest_SRVRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service  string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Protocol string `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	DnsName  string `protobuf:"bytes,3,opt,name=dns_name,json=dnsName,proto3" json:"dns_name,omitempty"`
}

func (x *WatchRequest_SRVRequest) Reset() {
	*x = WatchRequest_SRVRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servok_api_v1_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchRequest_SRVRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchRequest_SRVRequest) ProtoMessage() {}

func (x *WatchRequest_SRVRequest) ProtoReflect() protoreflect.Message {
	mi := &file_servok_api_v1_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchRequest_SRVRequest.ProtoReflect.Descriptor instead.
func (*WatchRequest_SRVRequest) Descriptor() ([]byte, []int) {
	return file_servok_api_v1_v1_proto_rawDescGZIP(), []int{0, 0}
}

func (x *WatchRequest_SRVRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *WatchRequest_SRVRequest) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *WatchRequest_SRVRequest) GetDnsName() string {
	if x != nil {
		return x.DnsName
	}
	return ""
}

var File_servok_api_v1_v1_proto protoreflect.FileDescriptor

var file_servok_api_v1_v1_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x65, 0x72, 0x76, 0x6f, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x6f, 0x6b,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xcd, 0x02, 0x0a, 0x0c, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x44, 0x0a, 0x03, 0x73, 0x72, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x6f, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x52, 0x56, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01,
	0x48, 0x00, 0x52, 0x03, 0x73, 0x72, 0x76, 0x1a, 0xdb, 0x01, 0x0a, 0x0a, 0x53, 0x52, 0x56, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x31, 0xfa, 0x42, 0x2e, 0x72, 0x2c, 0x28, 0xfd,
	0x01, 0x32, 0x27, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x28, 0x5b, 0x61, 0x2d,
	0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x5c, 0x2e, 0x5d, 0x7b, 0x30, 0x2c, 0x32, 0x35, 0x31, 0x7d, 0x5b,
	0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x3f, 0x24, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x16, 0xfa, 0x42, 0x13, 0x72, 0x11, 0x32, 0x0f, 0x5e, 0x28,
	0x28, 0x74, 0x63, 0x70, 0x29, 0x7c, 0x28, 0x75, 0x64, 0x70, 0x29, 0x29, 0x24, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x4c, 0x0a, 0x08, 0x64, 0x6e, 0x73, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x31, 0xfa, 0x42, 0x2e, 0x72, 0x2c,
	0x28, 0xfd, 0x01, 0x32, 0x27, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x28, 0x5b,
	0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x5c, 0x2e, 0x5d, 0x7b, 0x30, 0x2c, 0x32, 0x35, 0x31,
	0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x3f, 0x24, 0x52, 0x07, 0x64, 0x6e,
	0x73, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x19, 0x0a, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x12, 0x03, 0xf8, 0x42, 0x01,
	0x22, 0x46, 0x0a, 0x0d, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x35, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x6f, 0x6b, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x09, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x52, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x32, 0x59, 0x0a, 0x0f,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x46, 0x0a, 0x05, 0x57, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x6f,
	0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x6f, 0x6b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0xb1, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x6f, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x07, 0x56,
	0x31, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2f, 0x73, 0x70, 0x69,
	0x63, 0x65, 0x64, 0x62, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x6f, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x3b, 0x61, 0x70, 0x69, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x41, 0x58, 0xaa, 0x02, 0x0d,
	0x53, 0x65, 0x72, 0x76, 0x6f, 0x6b, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d,
	0x53, 0x65, 0x72, 0x76, 0x6f, 0x6b, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19,
	0x53, 0x65, 0x72, 0x76, 0x6f, 0x6b, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x53, 0x65, 0x72, 0x76,
	0x6f, 0x6b, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_servok_api_v1_v1_proto_rawDescOnce sync.Once
	file_servok_api_v1_v1_proto_rawDescData = file_servok_api_v1_v1_proto_rawDesc
)

func file_servok_api_v1_v1_proto_rawDescGZIP() []byte {
	file_servok_api_v1_v1_proto_rawDescOnce.Do(func() {
		file_servok_api_v1_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_servok_api_v1_v1_proto_rawDescData)
	})
	return file_servok_api_v1_v1_proto_rawDescData
}

var file_servok_api_v1_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_servok_api_v1_v1_proto_goTypes = []interface{}{
	(*WatchRequest)(nil),            // 0: servok.api.v1.WatchRequest
	(*WatchResponse)(nil),           // 1: servok.api.v1.WatchResponse
	(*Endpoint)(nil),                // 2: servok.api.v1.Endpoint
	(*WatchRequest_SRVRequest)(nil), // 3: servok.api.v1.WatchRequest.SRVRequest
}
var file_servok_api_v1_v1_proto_depIdxs = []int32{
	3, // 0: servok.api.v1.WatchRequest.srv:type_name -> servok.api.v1.WatchRequest.SRVRequest
	2, // 1: servok.api.v1.WatchResponse.endpoints:type_name -> servok.api.v1.Endpoint
	0, // 2: servok.api.v1.EndpointService.Watch:input_type -> servok.api.v1.WatchRequest
	1, // 3: servok.api.v1.EndpointService.Watch:output_type -> servok.api.v1.WatchResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_servok_api_v1_v1_proto_init() }
func file_servok_api_v1_v1_proto_init() {
	if File_servok_api_v1_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_servok_api_v1_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_servok_api_v1_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_servok_api_v1_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Endpoint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_servok_api_v1_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchRequest_SRVRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_servok_api_v1_v1_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*WatchRequest_Srv)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_servok_api_v1_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_servok_api_v1_v1_proto_goTypes,
		DependencyIndexes: file_servok_api_v1_v1_proto_depIdxs,
		MessageInfos:      file_servok_api_v1_v1_proto_msgTypes,
	}.Build()
	File_servok_api_v1_v1_proto = out.File
	file_servok_api_v1_v1_proto_rawDesc = nil
	file_servok_api_v1_v1_proto_goTypes = nil
	file_servok_api_v1_v1_proto_depIdxs = nil
}