// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: bank/v1/client.proto

package bank_v1

import (
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

type Client struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Age       uint32 `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *Client) Reset() {
	*x = Client{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_v1_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client) ProtoMessage() {}

func (x *Client) ProtoReflect() protoreflect.Message {
	mi := &file_bank_v1_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client.ProtoReflect.Descriptor instead.
func (*Client) Descriptor() ([]byte, []int) {
	return file_bank_v1_client_proto_rawDescGZIP(), []int{0}
}

func (x *Client) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Client) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Client) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Client) GetAge() uint32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type CreateClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Client *Client `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
}

func (x *CreateClientRequest) Reset() {
	*x = CreateClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_v1_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClientRequest) ProtoMessage() {}

func (x *CreateClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_v1_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClientRequest.ProtoReflect.Descriptor instead.
func (*CreateClientRequest) Descriptor() ([]byte, []int) {
	return file_bank_v1_client_proto_rawDescGZIP(), []int{1}
}

func (x *CreateClientRequest) GetClient() *Client {
	if x != nil {
		return x.Client
	}
	return nil
}

type CreateClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Client *Client `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
}

func (x *CreateClientResponse) Reset() {
	*x = CreateClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_v1_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClientResponse) ProtoMessage() {}

func (x *CreateClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bank_v1_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClientResponse.ProtoReflect.Descriptor instead.
func (*CreateClientResponse) Descriptor() ([]byte, []int) {
	return file_bank_v1_client_proto_rawDescGZIP(), []int{2}
}

func (x *CreateClientResponse) GetClient() *Client {
	if x != nil {
		return x.Client
	}
	return nil
}

type GetAllClientsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllClientsRequest) Reset() {
	*x = GetAllClientsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_v1_client_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllClientsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllClientsRequest) ProtoMessage() {}

func (x *GetAllClientsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_v1_client_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllClientsRequest.ProtoReflect.Descriptor instead.
func (*GetAllClientsRequest) Descriptor() ([]byte, []int) {
	return file_bank_v1_client_proto_rawDescGZIP(), []int{3}
}

type GetAllClientsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clients []*Client `protobuf:"bytes,1,rep,name=clients,proto3" json:"clients,omitempty"`
}

func (x *GetAllClientsResponse) Reset() {
	*x = GetAllClientsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_v1_client_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllClientsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllClientsResponse) ProtoMessage() {}

func (x *GetAllClientsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bank_v1_client_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllClientsResponse.ProtoReflect.Descriptor instead.
func (*GetAllClientsResponse) Descriptor() ([]byte, []int) {
	return file_bank_v1_client_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllClientsResponse) GetClients() []*Client {
	if x != nil {
		return x.Clients
	}
	return nil
}

type GetClientByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetClientByIdRequest) Reset() {
	*x = GetClientByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_v1_client_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientByIdRequest) ProtoMessage() {}

func (x *GetClientByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_v1_client_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClientByIdRequest.ProtoReflect.Descriptor instead.
func (*GetClientByIdRequest) Descriptor() ([]byte, []int) {
	return file_bank_v1_client_proto_rawDescGZIP(), []int{5}
}

func (x *GetClientByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetClientByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Client *Client `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
}

func (x *GetClientByIdResponse) Reset() {
	*x = GetClientByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_v1_client_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientByIdResponse) ProtoMessage() {}

func (x *GetClientByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bank_v1_client_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClientByIdResponse.ProtoReflect.Descriptor instead.
func (*GetClientByIdResponse) Descriptor() ([]byte, []int) {
	return file_bank_v1_client_proto_rawDescGZIP(), []int{6}
}

func (x *GetClientByIdResponse) GetClient() *Client {
	if x != nil {
		return x.Client
	}
	return nil
}

var File_bank_v1_client_proto protoreflect.FileDescriptor

var file_bank_v1_client_proto_rawDesc = []byte{
	0x0a, 0x14, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x64, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x61, 0x67, 0x65, 0x22, 0x46, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x47, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a,
	0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x16,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4a, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x31, 0x0a, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x26, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x48, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6a, 0x6d, 0x69, 0x63, 0x68, 0x69, 0x65, 0x6c, 0x73, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x74, 0x65, 0x73,
	0x74, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x61,
	0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bank_v1_client_proto_rawDescOnce sync.Once
	file_bank_v1_client_proto_rawDescData = file_bank_v1_client_proto_rawDesc
)

func file_bank_v1_client_proto_rawDescGZIP() []byte {
	file_bank_v1_client_proto_rawDescOnce.Do(func() {
		file_bank_v1_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_bank_v1_client_proto_rawDescData)
	})
	return file_bank_v1_client_proto_rawDescData
}

var file_bank_v1_client_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_bank_v1_client_proto_goTypes = []interface{}{
	(*Client)(nil),                // 0: bank.service.v1.Client
	(*CreateClientRequest)(nil),   // 1: bank.service.v1.CreateClientRequest
	(*CreateClientResponse)(nil),  // 2: bank.service.v1.CreateClientResponse
	(*GetAllClientsRequest)(nil),  // 3: bank.service.v1.GetAllClientsRequest
	(*GetAllClientsResponse)(nil), // 4: bank.service.v1.GetAllClientsResponse
	(*GetClientByIdRequest)(nil),  // 5: bank.service.v1.GetClientByIdRequest
	(*GetClientByIdResponse)(nil), // 6: bank.service.v1.GetClientByIdResponse
}
var file_bank_v1_client_proto_depIdxs = []int32{
	0, // 0: bank.service.v1.CreateClientRequest.client:type_name -> bank.service.v1.Client
	0, // 1: bank.service.v1.CreateClientResponse.client:type_name -> bank.service.v1.Client
	0, // 2: bank.service.v1.GetAllClientsResponse.clients:type_name -> bank.service.v1.Client
	0, // 3: bank.service.v1.GetClientByIdResponse.client:type_name -> bank.service.v1.Client
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_bank_v1_client_proto_init() }
func file_bank_v1_client_proto_init() {
	if File_bank_v1_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bank_v1_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Client); i {
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
		file_bank_v1_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClientRequest); i {
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
		file_bank_v1_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClientResponse); i {
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
		file_bank_v1_client_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllClientsRequest); i {
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
		file_bank_v1_client_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllClientsResponse); i {
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
		file_bank_v1_client_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClientByIdRequest); i {
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
		file_bank_v1_client_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClientByIdResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bank_v1_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bank_v1_client_proto_goTypes,
		DependencyIndexes: file_bank_v1_client_proto_depIdxs,
		MessageInfos:      file_bank_v1_client_proto_msgTypes,
	}.Build()
	File_bank_v1_client_proto = out.File
	file_bank_v1_client_proto_rawDesc = nil
	file_bank_v1_client_proto_goTypes = nil
	file_bank_v1_client_proto_depIdxs = nil
}