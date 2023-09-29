// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: event/v1/event.proto

package eventv1

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

type Flight struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Pilot      string  `protobuf:"bytes,2,opt,name=pilot,proto3" json:"pilot,omitempty"`
	Squadron   string  `protobuf:"bytes,3,opt,name=squadron,proto3" json:"squadron,omitempty"`
	Ac         string  `protobuf:"bytes,4,opt,name=ac,proto3" json:"ac,omitempty"`
	FuelStatus float32 `protobuf:"fixed32,5,opt,name=fuel_status,json=fuelStatus,proto3" json:"fuel_status,omitempty"`
	LastUpdate string  `protobuf:"bytes,6,opt,name=last_update,json=lastUpdate,proto3" json:"last_update,omitempty"`
}

func (x *Flight) Reset() {
	*x = Flight{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_v1_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Flight) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Flight) ProtoMessage() {}

func (x *Flight) ProtoReflect() protoreflect.Message {
	mi := &file_event_v1_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Flight.ProtoReflect.Descriptor instead.
func (*Flight) Descriptor() ([]byte, []int) {
	return file_event_v1_event_proto_rawDescGZIP(), []int{0}
}

func (x *Flight) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Flight) GetPilot() string {
	if x != nil {
		return x.Pilot
	}
	return ""
}

func (x *Flight) GetSquadron() string {
	if x != nil {
		return x.Squadron
	}
	return ""
}

func (x *Flight) GetAc() string {
	if x != nil {
		return x.Ac
	}
	return ""
}

func (x *Flight) GetFuelStatus() float32 {
	if x != nil {
		return x.FuelStatus
	}
	return 0
}

func (x *Flight) GetLastUpdate() string {
	if x != nil {
		return x.LastUpdate
	}
	return ""
}

type InsertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Pilot      string  `protobuf:"bytes,2,opt,name=pilot,proto3" json:"pilot,omitempty"`
	Squadron   string  `protobuf:"bytes,3,opt,name=squadron,proto3" json:"squadron,omitempty"`
	Ac         string  `protobuf:"bytes,4,opt,name=ac,proto3" json:"ac,omitempty"`
	FuelStatus float32 `protobuf:"fixed32,5,opt,name=fuel_status,json=fuelStatus,proto3" json:"fuel_status,omitempty"`
	LastUpdate string  `protobuf:"bytes,6,opt,name=last_update,json=lastUpdate,proto3" json:"last_update,omitempty"`
}

func (x *InsertRequest) Reset() {
	*x = InsertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_v1_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertRequest) ProtoMessage() {}

func (x *InsertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_v1_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertRequest.ProtoReflect.Descriptor instead.
func (*InsertRequest) Descriptor() ([]byte, []int) {
	return file_event_v1_event_proto_rawDescGZIP(), []int{1}
}

func (x *InsertRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InsertRequest) GetPilot() string {
	if x != nil {
		return x.Pilot
	}
	return ""
}

func (x *InsertRequest) GetSquadron() string {
	if x != nil {
		return x.Squadron
	}
	return ""
}

func (x *InsertRequest) GetAc() string {
	if x != nil {
		return x.Ac
	}
	return ""
}

func (x *InsertRequest) GetFuelStatus() float32 {
	if x != nil {
		return x.FuelStatus
	}
	return 0
}

func (x *InsertRequest) GetLastUpdate() string {
	if x != nil {
		return x.LastUpdate
	}
	return ""
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Pilot      string  `protobuf:"bytes,2,opt,name=pilot,proto3" json:"pilot,omitempty"`
	Squadron   string  `protobuf:"bytes,3,opt,name=squadron,proto3" json:"squadron,omitempty"`
	Ac         string  `protobuf:"bytes,4,opt,name=ac,proto3" json:"ac,omitempty"`
	FuelStatus float32 `protobuf:"fixed32,5,opt,name=fuel_status,json=fuelStatus,proto3" json:"fuel_status,omitempty"`
	LastUpdate string  `protobuf:"bytes,6,opt,name=last_update,json=lastUpdate,proto3" json:"last_update,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_v1_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_v1_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_event_v1_event_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRequest) GetPilot() string {
	if x != nil {
		return x.Pilot
	}
	return ""
}

func (x *UpdateRequest) GetSquadron() string {
	if x != nil {
		return x.Squadron
	}
	return ""
}

func (x *UpdateRequest) GetAc() string {
	if x != nil {
		return x.Ac
	}
	return ""
}

func (x *UpdateRequest) GetFuelStatus() float32 {
	if x != nil {
		return x.FuelStatus
	}
	return 0
}

func (x *UpdateRequest) GetLastUpdate() string {
	if x != nil {
		return x.LastUpdate
	}
	return ""
}

type RemoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Pilot      string  `protobuf:"bytes,2,opt,name=pilot,proto3" json:"pilot,omitempty"`
	Squadron   string  `protobuf:"bytes,3,opt,name=squadron,proto3" json:"squadron,omitempty"`
	Ac         string  `protobuf:"bytes,4,opt,name=ac,proto3" json:"ac,omitempty"`
	FuelStatus float32 `protobuf:"fixed32,5,opt,name=fuel_status,json=fuelStatus,proto3" json:"fuel_status,omitempty"`
	LastUpdate string  `protobuf:"bytes,6,opt,name=last_update,json=lastUpdate,proto3" json:"last_update,omitempty"`
}

func (x *RemoveRequest) Reset() {
	*x = RemoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_v1_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRequest) ProtoMessage() {}

func (x *RemoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_v1_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRequest.ProtoReflect.Descriptor instead.
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return file_event_v1_event_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RemoveRequest) GetPilot() string {
	if x != nil {
		return x.Pilot
	}
	return ""
}

func (x *RemoveRequest) GetSquadron() string {
	if x != nil {
		return x.Squadron
	}
	return ""
}

func (x *RemoveRequest) GetAc() string {
	if x != nil {
		return x.Ac
	}
	return ""
}

func (x *RemoveRequest) GetFuelStatus() float32 {
	if x != nil {
		return x.FuelStatus
	}
	return 0
}

func (x *RemoveRequest) GetLastUpdate() string {
	if x != nil {
		return x.LastUpdate
	}
	return ""
}

type RetrieveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response map[string]string `protobuf:"bytes,1,rep,name=response,proto3" json:"response,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RetrieveRequest) Reset() {
	*x = RetrieveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_v1_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveRequest) ProtoMessage() {}

func (x *RetrieveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_v1_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveRequest.ProtoReflect.Descriptor instead.
func (*RetrieveRequest) Descriptor() ([]byte, []int) {
	return file_event_v1_event_proto_rawDescGZIP(), []int{4}
}

func (x *RetrieveRequest) GetResponse() map[string]string {
	if x != nil {
		return x.Response
	}
	return nil
}

type InsertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response map[string]string `protobuf:"bytes,1,rep,name=response,proto3" json:"response,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *InsertResponse) Reset() {
	*x = InsertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_v1_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertResponse) ProtoMessage() {}

func (x *InsertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_event_v1_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertResponse.ProtoReflect.Descriptor instead.
func (*InsertResponse) Descriptor() ([]byte, []int) {
	return file_event_v1_event_proto_rawDescGZIP(), []int{5}
}

func (x *InsertResponse) GetResponse() map[string]string {
	if x != nil {
		return x.Response
	}
	return nil
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response map[string]string `protobuf:"bytes,1,rep,name=response,proto3" json:"response,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_v1_event_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_event_v1_event_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_event_v1_event_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateResponse) GetResponse() map[string]string {
	if x != nil {
		return x.Response
	}
	return nil
}

type RemoveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response map[string]string `protobuf:"bytes,1,rep,name=response,proto3" json:"response,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RemoveResponse) Reset() {
	*x = RemoveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_v1_event_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveResponse) ProtoMessage() {}

func (x *RemoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_event_v1_event_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveResponse.ProtoReflect.Descriptor instead.
func (*RemoveResponse) Descriptor() ([]byte, []int) {
	return file_event_v1_event_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveResponse) GetResponse() map[string]string {
	if x != nil {
		return x.Response
	}
	return nil
}

type RetrieveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flights []*Flight `protobuf:"bytes,1,rep,name=flights,proto3" json:"flights,omitempty"`
}

func (x *RetrieveResponse) Reset() {
	*x = RetrieveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_v1_event_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveResponse) ProtoMessage() {}

func (x *RetrieveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_event_v1_event_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveResponse.ProtoReflect.Descriptor instead.
func (*RetrieveResponse) Descriptor() ([]byte, []int) {
	return file_event_v1_event_proto_rawDescGZIP(), []int{8}
}

func (x *RetrieveResponse) GetFlights() []*Flight {
	if x != nil {
		return x.Flights
	}
	return nil
}

var File_event_v1_event_proto protoreflect.FileDescriptor

var file_event_v1_event_proto_rawDesc = []byte{
	0x0a, 0x14, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x22, 0x9c, 0x01, 0x0a, 0x06, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x69, 0x6c, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x69, 0x6c, 0x6f,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x71, 0x75, 0x61, 0x64, 0x72, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x71, 0x75, 0x61, 0x64, 0x72, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x61, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x61, 0x63, 0x12, 0x1f, 0x0a,
	0x0b, 0x66, 0x75, 0x65, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0a, 0x66, 0x75, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22,
	0xa3, 0x01, 0x0a, 0x0d, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x69, 0x6c, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x69, 0x6c, 0x6f, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x71, 0x75, 0x61, 0x64,
	0x72, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x71, 0x75, 0x61, 0x64,
	0x72, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x61, 0x63, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x75, 0x65, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x66, 0x75, 0x65, 0x6c, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0xa3, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x69, 0x6c, 0x6f, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x69, 0x6c, 0x6f, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x71, 0x75, 0x61, 0x64, 0x72, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x71, 0x75, 0x61, 0x64, 0x72, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x63, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x61, 0x63, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x75, 0x65,
	0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a,
	0x66, 0x75, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0xa3, 0x01, 0x0a, 0x0d,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x69, 0x6c, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x69,
	0x6c, 0x6f, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x71, 0x75, 0x61, 0x64, 0x72, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x71, 0x75, 0x61, 0x64, 0x72, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x61, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x61, 0x63, 0x12,
	0x1f, 0x0a, 0x0b, 0x66, 0x75, 0x65, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x66, 0x75, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x22, 0x93, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x3b, 0x0a, 0x0d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x91, 0x01, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x3b,
	0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x91, 0x01, 0x0a, 0x0e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x1a, 0x3b, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x91, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x42, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x3b, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x3e, 0x0a, 0x10, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x66, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x07, 0x66, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x73, 0x32, 0x90, 0x02, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x17,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3d, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x17, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x43, 0x0a, 0x08, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x12, 0x19, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_v1_event_proto_rawDescOnce sync.Once
	file_event_v1_event_proto_rawDescData = file_event_v1_event_proto_rawDesc
)

func file_event_v1_event_proto_rawDescGZIP() []byte {
	file_event_v1_event_proto_rawDescOnce.Do(func() {
		file_event_v1_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_v1_event_proto_rawDescData)
	})
	return file_event_v1_event_proto_rawDescData
}

var file_event_v1_event_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_event_v1_event_proto_goTypes = []interface{}{
	(*Flight)(nil),           // 0: event.v1.Flight
	(*InsertRequest)(nil),    // 1: event.v1.InsertRequest
	(*UpdateRequest)(nil),    // 2: event.v1.UpdateRequest
	(*RemoveRequest)(nil),    // 3: event.v1.RemoveRequest
	(*RetrieveRequest)(nil),  // 4: event.v1.RetrieveRequest
	(*InsertResponse)(nil),   // 5: event.v1.InsertResponse
	(*UpdateResponse)(nil),   // 6: event.v1.UpdateResponse
	(*RemoveResponse)(nil),   // 7: event.v1.RemoveResponse
	(*RetrieveResponse)(nil), // 8: event.v1.RetrieveResponse
	nil,                      // 9: event.v1.RetrieveRequest.ResponseEntry
	nil,                      // 10: event.v1.InsertResponse.ResponseEntry
	nil,                      // 11: event.v1.UpdateResponse.ResponseEntry
	nil,                      // 12: event.v1.RemoveResponse.ResponseEntry
}
var file_event_v1_event_proto_depIdxs = []int32{
	9,  // 0: event.v1.RetrieveRequest.response:type_name -> event.v1.RetrieveRequest.ResponseEntry
	10, // 1: event.v1.InsertResponse.response:type_name -> event.v1.InsertResponse.ResponseEntry
	11, // 2: event.v1.UpdateResponse.response:type_name -> event.v1.UpdateResponse.ResponseEntry
	12, // 3: event.v1.RemoveResponse.response:type_name -> event.v1.RemoveResponse.ResponseEntry
	0,  // 4: event.v1.RetrieveResponse.flights:type_name -> event.v1.Flight
	1,  // 5: event.v1.EventService.Insert:input_type -> event.v1.InsertRequest
	2,  // 6: event.v1.EventService.Update:input_type -> event.v1.UpdateRequest
	3,  // 7: event.v1.EventService.Remove:input_type -> event.v1.RemoveRequest
	4,  // 8: event.v1.EventService.Retrieve:input_type -> event.v1.RetrieveRequest
	5,  // 9: event.v1.EventService.Insert:output_type -> event.v1.InsertResponse
	6,  // 10: event.v1.EventService.Update:output_type -> event.v1.UpdateResponse
	7,  // 11: event.v1.EventService.Remove:output_type -> event.v1.RemoveResponse
	8,  // 12: event.v1.EventService.Retrieve:output_type -> event.v1.RetrieveResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_event_v1_event_proto_init() }
func file_event_v1_event_proto_init() {
	if File_event_v1_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_event_v1_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Flight); i {
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
		file_event_v1_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertRequest); i {
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
		file_event_v1_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_event_v1_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRequest); i {
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
		file_event_v1_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveRequest); i {
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
		file_event_v1_event_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertResponse); i {
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
		file_event_v1_event_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
		file_event_v1_event_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveResponse); i {
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
		file_event_v1_event_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveResponse); i {
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
			RawDescriptor: file_event_v1_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_event_v1_event_proto_goTypes,
		DependencyIndexes: file_event_v1_event_proto_depIdxs,
		MessageInfos:      file_event_v1_event_proto_msgTypes,
	}.Build()
	File_event_v1_event_proto = out.File
	file_event_v1_event_proto_rawDesc = nil
	file_event_v1_event_proto_goTypes = nil
	file_event_v1_event_proto_depIdxs = nil
}
