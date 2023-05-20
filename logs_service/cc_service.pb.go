// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.6
// source: cc_service.proto

package logs_service

import (
	auth_service "github.com/Authentura/codectrl-go-protobufs/auth_service"
	log "github.com/Authentura/codectrl-go-protobufs/data/log"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Status codes for whether or not a particular request has succeeded.
type RequestStatus int32

const (
	RequestStatus_CONFIRMED RequestStatus = 0
	RequestStatus_ERROR     RequestStatus = 1
)

// Enum value maps for RequestStatus.
var (
	RequestStatus_name = map[int32]string{
		0: "CONFIRMED",
		1: "ERROR",
	}
	RequestStatus_value = map[string]int32{
		"CONFIRMED": 0,
		"ERROR":     1,
	}
)

func (x RequestStatus) Enum() *RequestStatus {
	p := new(RequestStatus)
	*p = x
	return p
}

func (x RequestStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_cc_service_proto_enumTypes[0].Descriptor()
}

func (RequestStatus) Type() protoreflect.EnumType {
	return &file_cc_service_proto_enumTypes[0]
}

func (x RequestStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RequestStatus.Descriptor instead.
func (RequestStatus) EnumDescriptor() ([]byte, []int) {
	return file_cc_service_proto_rawDescGZIP(), []int{0}
}

// Describes the connection between the interface and a given server. Each
// client is supplied with a uuid that is saved to disk or to localStorage. The
// server uses this information to determine which logs should be sent to each
// client and to skip duplicate identified by the `uuid` of the log.
type Connection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid  string              `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Token *auth_service.Token `protobuf:"bytes,2,opt,name=token,proto3,oneof" json:"token,omitempty"`
}

func (x *Connection) Reset() {
	*x = Connection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Connection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connection) ProtoMessage() {}

func (x *Connection) ProtoReflect() protoreflect.Message {
	mi := &file_cc_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connection.ProtoReflect.Descriptor instead.
func (*Connection) Descriptor() ([]byte, []int) {
	return file_cc_service_proto_rawDescGZIP(), []int{0}
}

func (x *Connection) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Connection) GetToken() *auth_service.Token {
	if x != nil {
		return x.Token
	}
	return nil
}

// Returned by the procedures to describe the result of a request.
type RequestResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message    string                                   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status     RequestStatus                            `protobuf:"varint,2,opt,name=status,proto3,enum=codectrl.logs_service.RequestStatus" json:"status,omitempty"`
	AuthStatus *auth_service.GenerateTokenRequestResult `protobuf:"bytes,3,opt,name=authStatus,proto3,oneof" json:"authStatus,omitempty"`
}

func (x *RequestResult) Reset() {
	*x = RequestResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestResult) ProtoMessage() {}

func (x *RequestResult) ProtoReflect() protoreflect.Message {
	mi := &file_cc_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestResult.ProtoReflect.Descriptor instead.
func (*RequestResult) Descriptor() ([]byte, []int) {
	return file_cc_service_proto_rawDescGZIP(), []int{1}
}

func (x *RequestResult) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RequestResult) GetStatus() RequestStatus {
	if x != nil {
		return x.Status
	}
	return RequestStatus_CONFIRMED
}

func (x *RequestResult) GetAuthStatus() *auth_service.GenerateTokenRequestResult {
	if x != nil {
		return x.AuthStatus
	}
	return nil
}

// Server details about the current gRPC server.
type ServerDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host                   string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port                   uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Uptime                 uint64 `protobuf:"varint,3,opt,name=uptime,proto3" json:"uptime,omitempty"`
	RequiresAuthentication bool   `protobuf:"varint,4,opt,name=requiresAuthentication,proto3" json:"requiresAuthentication,omitempty"`
}

func (x *ServerDetails) Reset() {
	*x = ServerDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerDetails) ProtoMessage() {}

func (x *ServerDetails) ProtoReflect() protoreflect.Message {
	mi := &file_cc_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerDetails.ProtoReflect.Descriptor instead.
func (*ServerDetails) Descriptor() ([]byte, []int) {
	return file_cc_service_proto_rawDescGZIP(), []int{2}
}

func (x *ServerDetails) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ServerDetails) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ServerDetails) GetUptime() uint64 {
	if x != nil {
		return x.Uptime
	}
	return 0
}

func (x *ServerDetails) GetRequiresAuthentication() bool {
	if x != nil {
		return x.RequiresAuthentication
	}
	return false
}

var File_cc_service_proto protoreflect.FileDescriptor

var file_cc_service_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x15, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x74, 0x72, 0x6c, 0x2e, 0x6c, 0x6f, 0x67,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x09, 0x6c, 0x6f, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a,
	0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12,
	0x37, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x74, 0x72, 0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x48, 0x00, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0xce, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3c,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x74, 0x72, 0x6c, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x56, 0x0a, 0x0a,
	0x61, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x74, 0x72, 0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x87, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75,
	0x70, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x16, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x73, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x73, 0x41,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x29, 0x0a,
	0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0d,
	0x0a, 0x09, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x32, 0xa4, 0x03, 0x0a, 0x09, 0x4c, 0x6f, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67,
	0x12, 0x21, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x74, 0x72, 0x6c, 0x2e, 0x6c, 0x6f, 0x67, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x74, 0x72, 0x6c, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x22, 0x00, 0x12, 0x48, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63,
	0x74, 0x72, 0x6c, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x63, 0x74, 0x72, 0x6c, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c, 0x6f, 0x67, 0x2e,
	0x4c, 0x6f, 0x67, 0x22, 0x00, 0x30, 0x01, 0x12, 0x52, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x74, 0x72, 0x6c, 0x2e, 0x6c,
	0x6f, 0x67, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x21, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x74, 0x72, 0x6c,
	0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x63, 0x0a, 0x16, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x74, 0x72, 0x6c, 0x2e,
	0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x74,
	0x72, 0x6c, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x32,
	0xa4, 0x01, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x49, 0x0a,
	0x07, 0x53, 0x65, 0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x12, 0x16, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63,
	0x74, 0x72, 0x6c, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67,
	0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x74, 0x72, 0x6c, 0x2e, 0x6c, 0x6f, 0x67, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64,
	0x4c, 0x6f, 0x67, 0x73, 0x12, 0x16, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x74, 0x72, 0x6c, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x1a, 0x24, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x63, 0x74, 0x72, 0x6c, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x00, 0x28, 0x01, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x61, 0x2f,
	0x63, 0x6f, 0x64, 0x65, 0x63, 0x74, 0x72, 0x6c, 0x2d, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cc_service_proto_rawDescOnce sync.Once
	file_cc_service_proto_rawDescData = file_cc_service_proto_rawDesc
)

func file_cc_service_proto_rawDescGZIP() []byte {
	file_cc_service_proto_rawDescOnce.Do(func() {
		file_cc_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_cc_service_proto_rawDescData)
	})
	return file_cc_service_proto_rawDescData
}

var file_cc_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cc_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cc_service_proto_goTypes = []interface{}{
	(RequestStatus)(0),                              // 0: codectrl.logs_service.RequestStatus
	(*Connection)(nil),                              // 1: codectrl.logs_service.Connection
	(*RequestResult)(nil),                           // 2: codectrl.logs_service.RequestResult
	(*ServerDetails)(nil),                           // 3: codectrl.logs_service.ServerDetails
	(*auth_service.Token)(nil),                      // 4: codectrl.auth_service.Token
	(*auth_service.GenerateTokenRequestResult)(nil), // 5: codectrl.auth_service.GenerateTokenRequestResult
	(*emptypb.Empty)(nil),                           // 6: google.protobuf.Empty
	(*log.Log)(nil),                                 // 7: codectrl.data.log.Log
}
var file_cc_service_proto_depIdxs = []int32{
	4,  // 0: codectrl.logs_service.Connection.token:type_name -> codectrl.auth_service.Token
	0,  // 1: codectrl.logs_service.RequestResult.status:type_name -> codectrl.logs_service.RequestStatus
	5,  // 2: codectrl.logs_service.RequestResult.authStatus:type_name -> codectrl.auth_service.GenerateTokenRequestResult
	1,  // 3: codectrl.logs_service.LogServer.GetLog:input_type -> codectrl.logs_service.Connection
	1,  // 4: codectrl.logs_service.LogServer.GetLogs:input_type -> codectrl.logs_service.Connection
	6,  // 5: codectrl.logs_service.LogServer.GetServerDetails:input_type -> google.protobuf.Empty
	6,  // 6: codectrl.logs_service.LogServer.RegisterClient:input_type -> google.protobuf.Empty
	1,  // 7: codectrl.logs_service.LogServer.RegisterExistingClient:input_type -> codectrl.logs_service.Connection
	7,  // 8: codectrl.logs_service.LogClient.SendLog:input_type -> codectrl.data.log.Log
	7,  // 9: codectrl.logs_service.LogClient.SendLogs:input_type -> codectrl.data.log.Log
	7,  // 10: codectrl.logs_service.LogServer.GetLog:output_type -> codectrl.data.log.Log
	7,  // 11: codectrl.logs_service.LogServer.GetLogs:output_type -> codectrl.data.log.Log
	3,  // 12: codectrl.logs_service.LogServer.GetServerDetails:output_type -> codectrl.logs_service.ServerDetails
	1,  // 13: codectrl.logs_service.LogServer.RegisterClient:output_type -> codectrl.logs_service.Connection
	2,  // 14: codectrl.logs_service.LogServer.RegisterExistingClient:output_type -> codectrl.logs_service.RequestResult
	2,  // 15: codectrl.logs_service.LogClient.SendLog:output_type -> codectrl.logs_service.RequestResult
	2,  // 16: codectrl.logs_service.LogClient.SendLogs:output_type -> codectrl.logs_service.RequestResult
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_cc_service_proto_init() }
func file_cc_service_proto_init() {
	if File_cc_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cc_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Connection); i {
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
		file_cc_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestResult); i {
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
		file_cc_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerDetails); i {
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
	file_cc_service_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_cc_service_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cc_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_cc_service_proto_goTypes,
		DependencyIndexes: file_cc_service_proto_depIdxs,
		EnumInfos:         file_cc_service_proto_enumTypes,
		MessageInfos:      file_cc_service_proto_msgTypes,
	}.Build()
	File_cc_service_proto = out.File
	file_cc_service_proto_rawDesc = nil
	file_cc_service_proto_goTypes = nil
	file_cc_service_proto_depIdxs = nil
}
