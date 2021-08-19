// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/personcrud/personcrud.proto

package personcrud

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

type CreatePersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *CreatePersonRequest) Reset() {
	*x = CreatePersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_personcrud_personcrud_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePersonRequest) ProtoMessage() {}

func (x *CreatePersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_personcrud_personcrud_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePersonRequest.ProtoReflect.Descriptor instead.
func (*CreatePersonRequest) Descriptor() ([]byte, []int) {
	return file_api_personcrud_personcrud_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePersonRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreatePersonRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CreatePersonRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *CreatePersonRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type CreatePersonReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreatePersonReply) Reset() {
	*x = CreatePersonReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_personcrud_personcrud_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePersonReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePersonReply) ProtoMessage() {}

func (x *CreatePersonReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_personcrud_personcrud_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePersonReply.ProtoReflect.Descriptor instead.
func (*CreatePersonReply) Descriptor() ([]byte, []int) {
	return file_api_personcrud_personcrud_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePersonReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdatePersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *UpdatePersonRequest) Reset() {
	*x = UpdatePersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_personcrud_personcrud_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePersonRequest) ProtoMessage() {}

func (x *UpdatePersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_personcrud_personcrud_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePersonRequest.ProtoReflect.Descriptor instead.
func (*UpdatePersonRequest) Descriptor() ([]byte, []int) {
	return file_api_personcrud_personcrud_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePersonRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdatePersonRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UpdatePersonRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UpdatePersonRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type GetPersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPersonRequest) Reset() {
	*x = GetPersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_personcrud_personcrud_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonRequest) ProtoMessage() {}

func (x *GetPersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_personcrud_personcrud_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonRequest.ProtoReflect.Descriptor instead.
func (*GetPersonRequest) Descriptor() ([]byte, []int) {
	return file_api_personcrud_personcrud_proto_rawDescGZIP(), []int{3}
}

func (x *GetPersonRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PersonReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *PersonReply) Reset() {
	*x = PersonReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_personcrud_personcrud_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonReply) ProtoMessage() {}

func (x *PersonReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_personcrud_personcrud_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonReply.ProtoReflect.Descriptor instead.
func (*PersonReply) Descriptor() ([]byte, []int) {
	return file_api_personcrud_personcrud_proto_rawDescGZIP(), []int{4}
}

func (x *PersonReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PersonReply) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *PersonReply) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *PersonReply) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type ListPersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastName string `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
}

func (x *ListPersonRequest) Reset() {
	*x = ListPersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_personcrud_personcrud_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPersonRequest) ProtoMessage() {}

func (x *ListPersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_personcrud_personcrud_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPersonRequest.ProtoReflect.Descriptor instead.
func (*ListPersonRequest) Descriptor() ([]byte, []int) {
	return file_api_personcrud_personcrud_proto_rawDescGZIP(), []int{5}
}

func (x *ListPersonRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type ListPersonReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Persons []*PersonReply `protobuf:"bytes,1,rep,name=persons,proto3" json:"persons,omitempty"`
}

func (x *ListPersonReply) Reset() {
	*x = ListPersonReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_personcrud_personcrud_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPersonReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPersonReply) ProtoMessage() {}

func (x *ListPersonReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_personcrud_personcrud_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPersonReply.ProtoReflect.Descriptor instead.
func (*ListPersonReply) Descriptor() ([]byte, []int) {
	return file_api_personcrud_personcrud_proto_rawDescGZIP(), []int{6}
}

func (x *ListPersonReply) GetPersons() []*PersonReply {
	if x != nil {
		return x.Persons
	}
	return nil
}

var File_api_personcrud_personcrud_proto protoreflect.FileDescriptor

var file_api_personcrud_personcrud_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x63, 0x72, 0x75, 0x64,
	0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x63, 0x72, 0x75,
	0x64, 0x22, 0x75, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x23, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x75, 0x0a,
	0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6d, 0x0a, 0x0b, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x2f, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x35, 0x0a, 0x07, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x07, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x73, 0x32, 0xd5, 0x02, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x63, 0x72, 0x75,
	0x64, 0x12, 0x56, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x63, 0x72,
	0x75, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x50, 0x0a, 0x0c, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x63, 0x72, 0x75, 0x64, 0x2e,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4a, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x51, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x38, 0x0a, 0x0e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x63, 0x72, 0x75, 0x64, 0x50, 0x01, 0x5a, 0x24,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x63, 0x72, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x63, 0x72, 0x75, 0x64, 0x3b, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x63, 0x72, 0x75, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_personcrud_personcrud_proto_rawDescOnce sync.Once
	file_api_personcrud_personcrud_proto_rawDescData = file_api_personcrud_personcrud_proto_rawDesc
)

func file_api_personcrud_personcrud_proto_rawDescGZIP() []byte {
	file_api_personcrud_personcrud_proto_rawDescOnce.Do(func() {
		file_api_personcrud_personcrud_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_personcrud_personcrud_proto_rawDescData)
	})
	return file_api_personcrud_personcrud_proto_rawDescData
}

var file_api_personcrud_personcrud_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_personcrud_personcrud_proto_goTypes = []interface{}{
	(*CreatePersonRequest)(nil), // 0: api.personcrud.CreatePersonRequest
	(*CreatePersonReply)(nil),   // 1: api.personcrud.CreatePersonReply
	(*UpdatePersonRequest)(nil), // 2: api.personcrud.UpdatePersonRequest
	(*GetPersonRequest)(nil),    // 3: api.personcrud.GetPersonRequest
	(*PersonReply)(nil),         // 4: api.personcrud.PersonReply
	(*ListPersonRequest)(nil),   // 5: api.personcrud.ListPersonRequest
	(*ListPersonReply)(nil),     // 6: api.personcrud.ListPersonReply
}
var file_api_personcrud_personcrud_proto_depIdxs = []int32{
	4, // 0: api.personcrud.ListPersonReply.persons:type_name -> api.personcrud.PersonReply
	0, // 1: api.personcrud.Personcrud.CreatePerson:input_type -> api.personcrud.CreatePersonRequest
	2, // 2: api.personcrud.Personcrud.UpdatePerson:input_type -> api.personcrud.UpdatePersonRequest
	3, // 3: api.personcrud.Personcrud.GetPerson:input_type -> api.personcrud.GetPersonRequest
	5, // 4: api.personcrud.Personcrud.ListPersons:input_type -> api.personcrud.ListPersonRequest
	1, // 5: api.personcrud.Personcrud.CreatePerson:output_type -> api.personcrud.CreatePersonReply
	4, // 6: api.personcrud.Personcrud.UpdatePerson:output_type -> api.personcrud.PersonReply
	4, // 7: api.personcrud.Personcrud.GetPerson:output_type -> api.personcrud.PersonReply
	6, // 8: api.personcrud.Personcrud.ListPersons:output_type -> api.personcrud.ListPersonReply
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_personcrud_personcrud_proto_init() }
func file_api_personcrud_personcrud_proto_init() {
	if File_api_personcrud_personcrud_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_personcrud_personcrud_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePersonRequest); i {
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
		file_api_personcrud_personcrud_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePersonReply); i {
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
		file_api_personcrud_personcrud_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePersonRequest); i {
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
		file_api_personcrud_personcrud_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonRequest); i {
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
		file_api_personcrud_personcrud_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonReply); i {
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
		file_api_personcrud_personcrud_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPersonRequest); i {
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
		file_api_personcrud_personcrud_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPersonReply); i {
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
			RawDescriptor: file_api_personcrud_personcrud_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_personcrud_personcrud_proto_goTypes,
		DependencyIndexes: file_api_personcrud_personcrud_proto_depIdxs,
		MessageInfos:      file_api_personcrud_personcrud_proto_msgTypes,
	}.Build()
	File_api_personcrud_personcrud_proto = out.File
	file_api_personcrud_personcrud_proto_rawDesc = nil
	file_api_personcrud_personcrud_proto_goTypes = nil
	file_api_personcrud_personcrud_proto_depIdxs = nil
}
