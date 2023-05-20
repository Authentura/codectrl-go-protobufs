// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.6
// source: auth.proto

package auth_service

import (
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

type TokenIntent int32

const (
	TokenIntent_LOGGER   TokenIntent = 0
	TokenIntent_FRONTEND TokenIntent = 1
)

// Enum value maps for TokenIntent.
var (
	TokenIntent_name = map[int32]string{
		0: "LOGGER",
		1: "FRONTEND",
	}
	TokenIntent_value = map[string]int32{
		"LOGGER":   0,
		"FRONTEND": 1,
	}
)

func (x TokenIntent) Enum() *TokenIntent {
	p := new(TokenIntent)
	*p = x
	return p
}

func (x TokenIntent) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TokenIntent) Descriptor() protoreflect.EnumDescriptor {
	return file_auth_proto_enumTypes[0].Descriptor()
}

func (TokenIntent) Type() protoreflect.EnumType {
	return &file_auth_proto_enumTypes[0]
}

func (x TokenIntent) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TokenIntent.Descriptor instead.
func (TokenIntent) EnumDescriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

type VerifyTokenRequestResultEnum int32

const (
	VerifyTokenRequestResultEnum_UNAUTHORISED VerifyTokenRequestResultEnum = 0
	VerifyTokenRequestResultEnum_NOTFOUND     VerifyTokenRequestResultEnum = 1
	VerifyTokenRequestResultEnum_AUTHORISED   VerifyTokenRequestResultEnum = 2
)

// Enum value maps for VerifyTokenRequestResultEnum.
var (
	VerifyTokenRequestResultEnum_name = map[int32]string{
		0: "UNAUTHORISED",
		1: "NOTFOUND",
		2: "AUTHORISED",
	}
	VerifyTokenRequestResultEnum_value = map[string]int32{
		"UNAUTHORISED": 0,
		"NOTFOUND":     1,
		"AUTHORISED":   2,
	}
)

func (x VerifyTokenRequestResultEnum) Enum() *VerifyTokenRequestResultEnum {
	p := new(VerifyTokenRequestResultEnum)
	*p = x
	return p
}

func (x VerifyTokenRequestResultEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VerifyTokenRequestResultEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_auth_proto_enumTypes[1].Descriptor()
}

func (VerifyTokenRequestResultEnum) Type() protoreflect.EnumType {
	return &file_auth_proto_enumTypes[1]
}

func (x VerifyTokenRequestResultEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VerifyTokenRequestResultEnum.Descriptor instead.
func (VerifyTokenRequestResultEnum) EnumDescriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1}
}

type GenerateTokenRequestResultEnum int32

const (
	GenerateTokenRequestResultEnum_NAME_ALREADY_EXISTS        GenerateTokenRequestResultEnum = 0 // If another token under this name is already present.
	GenerateTokenRequestResultEnum_TOKEN_ALREADY_EXISTS       GenerateTokenRequestResultEnum = 1 // If the exact token is already present.
	GenerateTokenRequestResultEnum_TOKEN_GENERATION_SUCCEEDED GenerateTokenRequestResultEnum = 2 // If the token was successfully added.
)

// Enum value maps for GenerateTokenRequestResultEnum.
var (
	GenerateTokenRequestResultEnum_name = map[int32]string{
		0: "NAME_ALREADY_EXISTS",
		1: "TOKEN_ALREADY_EXISTS",
		2: "TOKEN_GENERATION_SUCCEEDED",
	}
	GenerateTokenRequestResultEnum_value = map[string]int32{
		"NAME_ALREADY_EXISTS":        0,
		"TOKEN_ALREADY_EXISTS":       1,
		"TOKEN_GENERATION_SUCCEEDED": 2,
	}
)

func (x GenerateTokenRequestResultEnum) Enum() *GenerateTokenRequestResultEnum {
	p := new(GenerateTokenRequestResultEnum)
	*p = x
	return p
}

func (x GenerateTokenRequestResultEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GenerateTokenRequestResultEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_auth_proto_enumTypes[2].Descriptor()
}

func (GenerateTokenRequestResultEnum) Type() protoreflect.EnumType {
	return &file_auth_proto_enumTypes[2]
}

func (x GenerateTokenRequestResultEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GenerateTokenRequestResultEnum.Descriptor instead.
func (GenerateTokenRequestResultEnum) EnumDescriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{2}
}

type RevokeTokenRequestResultEnum int32

const (
	RevokeTokenRequestResultEnum_TOKEN_NOT_FOUND            RevokeTokenRequestResultEnum = 0 // If the token is not found.
	RevokeTokenRequestResultEnum_TOKEN_REVOKATION_SUCCEEDED RevokeTokenRequestResultEnum = 1 // If the token was successfully added.
)

// Enum value maps for RevokeTokenRequestResultEnum.
var (
	RevokeTokenRequestResultEnum_name = map[int32]string{
		0: "TOKEN_NOT_FOUND",
		1: "TOKEN_REVOKATION_SUCCEEDED",
	}
	RevokeTokenRequestResultEnum_value = map[string]int32{
		"TOKEN_NOT_FOUND":            0,
		"TOKEN_REVOKATION_SUCCEEDED": 1,
	}
)

func (x RevokeTokenRequestResultEnum) Enum() *RevokeTokenRequestResultEnum {
	p := new(RevokeTokenRequestResultEnum)
	*p = x
	return p
}

func (x RevokeTokenRequestResultEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RevokeTokenRequestResultEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_auth_proto_enumTypes[3].Descriptor()
}

func (RevokeTokenRequestResultEnum) Type() protoreflect.EnumType {
	return &file_auth_proto_enumTypes[3]
}

func (x RevokeTokenRequestResultEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RevokeTokenRequestResultEnum.Descriptor instead.
func (RevokeTokenRequestResultEnum) EnumDescriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{3}
}

type Name struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inner string `protobuf:"bytes,1,opt,name=inner,proto3" json:"inner,omitempty"`
}

func (x *Name) Reset() {
	*x = Name{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Name) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Name) ProtoMessage() {}

func (x *Name) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Name.ProtoReflect.Descriptor instead.
func (*Name) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

func (x *Name) GetInner() string {
	if x != nil {
		return x.Inner
	}
	return ""
}

type TokenPermissions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Read  bool `protobuf:"varint,1,opt,name=read,proto3" json:"read,omitempty"`
	Write bool `protobuf:"varint,2,opt,name=write,proto3" json:"write,omitempty"`
}

func (x *TokenPermissions) Reset() {
	*x = TokenPermissions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenPermissions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenPermissions) ProtoMessage() {}

func (x *TokenPermissions) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenPermissions.ProtoReflect.Descriptor instead.
func (*TokenPermissions) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1}
}

func (x *TokenPermissions) GetRead() bool {
	if x != nil {
		return x.Read
	}
	return false
}

func (x *TokenPermissions) GetWrite() bool {
	if x != nil {
		return x.Write
	}
	return false
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        *Name             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Permissions *TokenPermissions `protobuf:"bytes,2,opt,name=permissions,proto3" json:"permissions,omitempty"`
	Intent      TokenIntent       `protobuf:"varint,3,opt,name=intent,proto3,enum=codectrl.auth_service.TokenIntent" json:"intent,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{2}
}

func (x *Token) GetName() *Name {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *Token) GetPermissions() *TokenPermissions {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *Token) GetIntent() TokenIntent {
	if x != nil {
		return x.Intent
	}
	return TokenIntent_LOGGER
}

type GenerateTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   *Name       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Intent TokenIntent `protobuf:"varint,2,opt,name=intent,proto3,enum=codectrl.auth_service.TokenIntent" json:"intent,omitempty"`
}

func (x *GenerateTokenRequest) Reset() {
	*x = GenerateTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenRequest) ProtoMessage() {}

func (x *GenerateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenRequest.ProtoReflect.Descriptor instead.
func (*GenerateTokenRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{3}
}

func (x *GenerateTokenRequest) GetName() *Name {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *GenerateTokenRequest) GetIntent() TokenIntent {
	if x != nil {
		return x.Intent
	}
	return TokenIntent_LOGGER
}

type VerifyTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  *Token      `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Intent TokenIntent `protobuf:"varint,2,opt,name=intent,proto3,enum=codectrl.auth_service.TokenIntent" json:"intent,omitempty"`
}

func (x *VerifyTokenRequest) Reset() {
	*x = VerifyTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyTokenRequest) ProtoMessage() {}

func (x *VerifyTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyTokenRequest.ProtoReflect.Descriptor instead.
func (*VerifyTokenRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{4}
}

func (x *VerifyTokenRequest) GetToken() *Token {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *VerifyTokenRequest) GetIntent() TokenIntent {
	if x != nil {
		return x.Intent
	}
	return TokenIntent_LOGGER
}

type VerifyTokenRequestResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string                       `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status  VerifyTokenRequestResultEnum `protobuf:"varint,2,opt,name=status,proto3,enum=codectrl.auth_service.VerifyTokenRequestResultEnum" json:"status,omitempty"`
}

func (x *VerifyTokenRequestResult) Reset() {
	*x = VerifyTokenRequestResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyTokenRequestResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyTokenRequestResult) ProtoMessage() {}

func (x *VerifyTokenRequestResult) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyTokenRequestResult.ProtoReflect.Descriptor instead.
func (*VerifyTokenRequestResult) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{5}
}

func (x *VerifyTokenRequestResult) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *VerifyTokenRequestResult) GetStatus() VerifyTokenRequestResultEnum {
	if x != nil {
		return x.Status
	}
	return VerifyTokenRequestResultEnum_UNAUTHORISED
}

type RevokeTokenRequestResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string                       `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status  RevokeTokenRequestResultEnum `protobuf:"varint,2,opt,name=status,proto3,enum=codectrl.auth_service.RevokeTokenRequestResultEnum" json:"status,omitempty"`
}

func (x *RevokeTokenRequestResult) Reset() {
	*x = RevokeTokenRequestResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevokeTokenRequestResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeTokenRequestResult) ProtoMessage() {}

func (x *RevokeTokenRequestResult) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeTokenRequestResult.ProtoReflect.Descriptor instead.
func (*RevokeTokenRequestResult) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{6}
}

func (x *RevokeTokenRequestResult) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RevokeTokenRequestResult) GetStatus() RevokeTokenRequestResultEnum {
	if x != nil {
		return x.Status
	}
	return RevokeTokenRequestResultEnum_TOKEN_NOT_FOUND
}

type GenerateTokenRequestResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status GenerateTokenRequestResultEnum `protobuf:"varint,1,opt,name=status,proto3,enum=codectrl.auth_service.GenerateTokenRequestResultEnum" json:"status,omitempty"`
	Token  *Token                         `protobuf:"bytes,2,opt,name=token,proto3,oneof" json:"token,omitempty"`
}

func (x *GenerateTokenRequestResult) Reset() {
	*x = GenerateTokenRequestResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateTokenRequestResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenRequestResult) ProtoMessage() {}

func (x *GenerateTokenRequestResult) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenRequestResult.ProtoReflect.Descriptor instead.
func (*GenerateTokenRequestResult) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{7}
}

func (x *GenerateTokenRequestResult) GetStatus() GenerateTokenRequestResultEnum {
	if x != nil {
		return x.Status
	}
	return GenerateTokenRequestResultEnum_NAME_ALREADY_EXISTS
}

func (x *GenerateTokenRequestResult) GetToken() *Token {
	if x != nil {
		return x.Token
	}
	return nil
}

type LoginUrl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *LoginUrl) Reset() {
	*x = LoginUrl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginUrl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginUrl) ProtoMessage() {}

func (x *LoginUrl) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginUrl.ProtoReflect.Descriptor instead.
func (*LoginUrl) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{8}
}

func (x *LoginUrl) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_auth_proto protoreflect.FileDescriptor

var file_auth_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x6f,
	0x64, 0x65, 0x63, 0x74, 0x72, 0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x1c, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x6e, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x3c,
	0x0a, 0x10, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x04, 0x72, 0x65, 0x61, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x77, 0x72, 0x69, 0x74, 0x65, 0x22, 0xbf, 0x01, 0x0a,
	0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2f, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x74, 0x72, 0x6c, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x49, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x63, 0x74, 0x72, 0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x3a, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x74, 0x72, 0x6c, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x83,
	0x01, 0x0a, 0x14, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x74, 0x72, 0x6c,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x61,
	0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63,
	0x74, 0x72, 0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x69, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x22, 0x84, 0x01, 0x0a, 0x12, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x63, 0x74, 0x72, 0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x3a, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x22, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x74, 0x72, 0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x81, 0x01, 0x0a, 0x18,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x74, 0x72, 0x6c, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x81, 0x01, 0x0a, 0x18, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x74, 0x72,
	0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52,
	0x65, 0x76, 0x6f, 0x6b, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0xae, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x4d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x35, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x74, 0x72, 0x6c, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x37, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x74, 0x72, 0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x48, 0x00,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x1c, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x72, 0x6c,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x2a, 0x27, 0x0a, 0x0b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x4f, 0x47, 0x47, 0x45, 0x52, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x46, 0x52, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x44, 0x10, 0x01, 0x2a, 0x4e, 0x0a, 0x1c, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x10, 0x0a, 0x0c, 0x55,
	0x4e, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x53, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x4e, 0x4f, 0x54, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x41,
	0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x53, 0x45, 0x44, 0x10, 0x02, 0x2a, 0x73, 0x0a, 0x1e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x17, 0x0a,
	0x13, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58,
	0x49, 0x53, 0x54, 0x53, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f,
	0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x01,
	0x12, 0x1e, 0x0a, 0x1a, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x02,
	0x2a, 0x53, 0x0a, 0x1c, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x6e, 0x75, 0x6d,
	0x12, 0x13, 0x0a, 0x0f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f,
	0x55, 0x4e, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x52,
	0x45, 0x56, 0x4f, 0x4b, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45,
	0x44, 0x45, 0x44, 0x10, 0x01, 0x32, 0xe8, 0x03, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6b, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x29, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x74,
	0x72, 0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x74, 0x72, 0x6c, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x71, 0x0a, 0x0d, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2b, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x74, 0x72,
	0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x74, 0x72, 0x6c, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x0b, 0x52, 0x65, 0x76, 0x6f,
	0x6b, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x74,
	0x72, 0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x2f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x74, 0x72, 0x6c,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65,
	0x76, 0x6f, 0x6b, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63,
	0x74, 0x72, 0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x1c, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x74, 0x72,
	0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0b, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x63, 0x74, 0x72, 0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x72, 0x6c, 0x22, 0x00,
	0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x61, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x74,
	0x72, 0x6c, 0x2d, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_proto_rawDescOnce sync.Once
	file_auth_proto_rawDescData = file_auth_proto_rawDesc
)

func file_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_proto_rawDescData)
	})
	return file_auth_proto_rawDescData
}

var file_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_auth_proto_goTypes = []interface{}{
	(TokenIntent)(0),                    // 0: codectrl.auth_service.TokenIntent
	(VerifyTokenRequestResultEnum)(0),   // 1: codectrl.auth_service.VerifyTokenRequestResultEnum
	(GenerateTokenRequestResultEnum)(0), // 2: codectrl.auth_service.GenerateTokenRequestResultEnum
	(RevokeTokenRequestResultEnum)(0),   // 3: codectrl.auth_service.RevokeTokenRequestResultEnum
	(*Name)(nil),                        // 4: codectrl.auth_service.Name
	(*TokenPermissions)(nil),            // 5: codectrl.auth_service.TokenPermissions
	(*Token)(nil),                       // 6: codectrl.auth_service.Token
	(*GenerateTokenRequest)(nil),        // 7: codectrl.auth_service.GenerateTokenRequest
	(*VerifyTokenRequest)(nil),          // 8: codectrl.auth_service.VerifyTokenRequest
	(*VerifyTokenRequestResult)(nil),    // 9: codectrl.auth_service.VerifyTokenRequestResult
	(*RevokeTokenRequestResult)(nil),    // 10: codectrl.auth_service.RevokeTokenRequestResult
	(*GenerateTokenRequestResult)(nil),  // 11: codectrl.auth_service.GenerateTokenRequestResult
	(*LoginUrl)(nil),                    // 12: codectrl.auth_service.LoginUrl
	(*emptypb.Empty)(nil),               // 13: google.protobuf.Empty
}
var file_auth_proto_depIdxs = []int32{
	4,  // 0: codectrl.auth_service.Token.name:type_name -> codectrl.auth_service.Name
	5,  // 1: codectrl.auth_service.Token.permissions:type_name -> codectrl.auth_service.TokenPermissions
	0,  // 2: codectrl.auth_service.Token.intent:type_name -> codectrl.auth_service.TokenIntent
	4,  // 3: codectrl.auth_service.GenerateTokenRequest.name:type_name -> codectrl.auth_service.Name
	0,  // 4: codectrl.auth_service.GenerateTokenRequest.intent:type_name -> codectrl.auth_service.TokenIntent
	6,  // 5: codectrl.auth_service.VerifyTokenRequest.token:type_name -> codectrl.auth_service.Token
	0,  // 6: codectrl.auth_service.VerifyTokenRequest.intent:type_name -> codectrl.auth_service.TokenIntent
	1,  // 7: codectrl.auth_service.VerifyTokenRequestResult.status:type_name -> codectrl.auth_service.VerifyTokenRequestResultEnum
	3,  // 8: codectrl.auth_service.RevokeTokenRequestResult.status:type_name -> codectrl.auth_service.RevokeTokenRequestResultEnum
	2,  // 9: codectrl.auth_service.GenerateTokenRequestResult.status:type_name -> codectrl.auth_service.GenerateTokenRequestResultEnum
	6,  // 10: codectrl.auth_service.GenerateTokenRequestResult.token:type_name -> codectrl.auth_service.Token
	8,  // 11: codectrl.auth_service.Authentication.VerifyToken:input_type -> codectrl.auth_service.VerifyTokenRequest
	7,  // 12: codectrl.auth_service.Authentication.GenerateToken:input_type -> codectrl.auth_service.GenerateTokenRequest
	6,  // 13: codectrl.auth_service.Authentication.RevokeToken:input_type -> codectrl.auth_service.Token
	6,  // 14: codectrl.auth_service.Authentication.RefreshToken:input_type -> codectrl.auth_service.Token
	13, // 15: codectrl.auth_service.Authentication.GithubLogin:input_type -> google.protobuf.Empty
	9,  // 16: codectrl.auth_service.Authentication.VerifyToken:output_type -> codectrl.auth_service.VerifyTokenRequestResult
	11, // 17: codectrl.auth_service.Authentication.GenerateToken:output_type -> codectrl.auth_service.GenerateTokenRequestResult
	10, // 18: codectrl.auth_service.Authentication.RevokeToken:output_type -> codectrl.auth_service.RevokeTokenRequestResult
	6,  // 19: codectrl.auth_service.Authentication.RefreshToken:output_type -> codectrl.auth_service.Token
	12, // 20: codectrl.auth_service.Authentication.GithubLogin:output_type -> codectrl.auth_service.LoginUrl
	16, // [16:21] is the sub-list for method output_type
	11, // [11:16] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_auth_proto_init() }
func file_auth_proto_init() {
	if File_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Name); i {
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
		file_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenPermissions); i {
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
		file_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateTokenRequest); i {
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
		file_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyTokenRequest); i {
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
		file_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyTokenRequestResult); i {
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
		file_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevokeTokenRequestResult); i {
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
		file_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateTokenRequestResult); i {
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
		file_auth_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginUrl); i {
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
	file_auth_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_depIdxs,
		EnumInfos:         file_auth_proto_enumTypes,
		MessageInfos:      file_auth_proto_msgTypes,
	}.Build()
	File_auth_proto = out.File
	file_auth_proto_rawDesc = nil
	file_auth_proto_goTypes = nil
	file_auth_proto_depIdxs = nil
}