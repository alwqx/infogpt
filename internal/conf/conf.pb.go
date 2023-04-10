// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: conf/conf.proto

package conf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Bootstrap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server *Server `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Admin  *Admin  `protobuf:"bytes,2,opt,name=admin,proto3" json:"admin,omitempty"`
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Bootstrap) GetAdmin() *Admin {
	if x != nil {
		return x.Admin
	}
	return nil
}

type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Http *Server_HTTP `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	Grpc *Server_GRPC `protobuf:"bytes,2,opt,name=grpc,proto3" json:"grpc,omitempty"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{1}
}

func (x *Server) GetHttp() *Server_HTTP {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *Server) GetGrpc() *Server_GRPC {
	if x != nil {
		return x.Grpc
	}
	return nil
}

type Admin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpenaiApiKey string    `protobuf:"bytes,1,opt,name=openai_api_key,json=openaiApiKey,proto3" json:"openai_api_key,omitempty"`
	ProxyUrl     string    `protobuf:"bytes,2,opt,name=proxy_url,json=proxyUrl,proto3" json:"proxy_url,omitempty"`
	GinRatelimit string    `protobuf:"bytes,3,opt,name=gin_ratelimit,json=ginRatelimit,proto3" json:"gin_ratelimit,omitempty"`
	Telegram     *Telegram `protobuf:"bytes,4,opt,name=telegram,proto3" json:"telegram,omitempty"`
	Wechat       *WeChat   `protobuf:"bytes,5,opt,name=wechat,proto3" json:"wechat,omitempty"`
}

func (x *Admin) Reset() {
	*x = Admin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Admin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Admin) ProtoMessage() {}

func (x *Admin) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Admin.ProtoReflect.Descriptor instead.
func (*Admin) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{2}
}

func (x *Admin) GetOpenaiApiKey() string {
	if x != nil {
		return x.OpenaiApiKey
	}
	return ""
}

func (x *Admin) GetProxyUrl() string {
	if x != nil {
		return x.ProxyUrl
	}
	return ""
}

func (x *Admin) GetGinRatelimit() string {
	if x != nil {
		return x.GinRatelimit
	}
	return ""
}

func (x *Admin) GetTelegram() *Telegram {
	if x != nil {
		return x.Telegram
	}
	return nil
}

func (x *Admin) GetWechat() *WeChat {
	if x != nil {
		return x.Wechat
	}
	return nil
}

type Telegram struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token         string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Ratelimit     string   `protobuf:"bytes,2,opt,name=ratelimit,proto3" json:"ratelimit,omitempty"`
	UserRatelimit string   `protobuf:"bytes,3,opt,name=user_ratelimit,json=userRatelimit,proto3" json:"user_ratelimit,omitempty"`
	ExcludeKeys   []string `protobuf:"bytes,4,rep,name=exclude_keys,json=excludeKeys,proto3" json:"exclude_keys,omitempty"`
}

func (x *Telegram) Reset() {
	*x = Telegram{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Telegram) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Telegram) ProtoMessage() {}

func (x *Telegram) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Telegram.ProtoReflect.Descriptor instead.
func (*Telegram) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{3}
}

func (x *Telegram) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Telegram) GetRatelimit() string {
	if x != nil {
		return x.Ratelimit
	}
	return ""
}

func (x *Telegram) GetUserRatelimit() string {
	if x != nil {
		return x.UserRatelimit
	}
	return ""
}

func (x *Telegram) GetExcludeKeys() []string {
	if x != nil {
		return x.ExcludeKeys
	}
	return nil
}

type WeChat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId          string            `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	AppSecret      string            `protobuf:"bytes,2,opt,name=app_secret,json=appSecret,proto3" json:"app_secret,omitempty"`
	Token          string            `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	EncodingAesKey string            `protobuf:"bytes,4,opt,name=encoding_aes_key,json=encodingAesKey,proto3" json:"encoding_aes_key,omitempty"`
	Ratelimit      string            `protobuf:"bytes,5,opt,name=ratelimit,proto3" json:"ratelimit,omitempty"`
	UserRatelimit  string            `protobuf:"bytes,6,opt,name=user_ratelimit,json=userRatelimit,proto3" json:"user_ratelimit,omitempty"`
	ExcludeKeys    []string          `protobuf:"bytes,7,rep,name=exclude_keys,json=excludeKeys,proto3" json:"exclude_keys,omitempty"`
	AutoReplay     map[string]string `protobuf:"bytes,8,rep,name=auto_replay,json=autoReplay,proto3" json:"auto_replay,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *WeChat) Reset() {
	*x = WeChat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeChat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeChat) ProtoMessage() {}

func (x *WeChat) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeChat.ProtoReflect.Descriptor instead.
func (*WeChat) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{4}
}

func (x *WeChat) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *WeChat) GetAppSecret() string {
	if x != nil {
		return x.AppSecret
	}
	return ""
}

func (x *WeChat) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *WeChat) GetEncodingAesKey() string {
	if x != nil {
		return x.EncodingAesKey
	}
	return ""
}

func (x *WeChat) GetRatelimit() string {
	if x != nil {
		return x.Ratelimit
	}
	return ""
}

func (x *WeChat) GetUserRatelimit() string {
	if x != nil {
		return x.UserRatelimit
	}
	return ""
}

func (x *WeChat) GetExcludeKeys() []string {
	if x != nil {
		return x.ExcludeKeys
	}
	return nil
}

func (x *WeChat) GetAutoReplay() map[string]string {
	if x != nil {
		return x.AutoReplay
	}
	return nil
}

type Server_HTTP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network      string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr         string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout      *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	ProxyUrl     string               `protobuf:"bytes,4,opt,name=proxy_url,json=proxyUrl,proto3" json:"proxy_url,omitempty"`
	ProxyTimeout *durationpb.Duration `protobuf:"bytes,5,opt,name=proxy_timeout,json=proxyTimeout,proto3" json:"proxy_timeout,omitempty"`
}

func (x *Server_HTTP) Reset() {
	*x = Server_HTTP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_HTTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_HTTP) ProtoMessage() {}

func (x *Server_HTTP) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_HTTP.ProtoReflect.Descriptor instead.
func (*Server_HTTP) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Server_HTTP) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_HTTP) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_HTTP) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *Server_HTTP) GetProxyUrl() string {
	if x != nil {
		return x.ProxyUrl
	}
	return ""
}

func (x *Server_HTTP) GetProxyTimeout() *durationpb.Duration {
	if x != nil {
		return x.ProxyTimeout
	}
	return nil
}

type Server_GRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr    string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *Server_GRPC) Reset() {
	*x = Server_GRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_GRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_GRPC) ProtoMessage() {}

func (x *Server_GRPC) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_GRPC.ProtoReflect.Descriptor instead.
func (*Server_GRPC) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Server_GRPC) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_GRPC) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_GRPC) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

var File_conf_conf_proto protoreflect.FileDescriptor

var file_conf_conf_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a,
	0x09, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x22,
	0x96, 0x03, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x04, 0x68, 0x74,
	0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48, 0x54, 0x54,
	0x50, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12, 0x2b, 0x0a, 0x04, 0x67, 0x72, 0x70, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x04,
	0x67, 0x72, 0x70, 0x63, 0x1a, 0xc6, 0x01, 0x0a, 0x04, 0x48, 0x54, 0x54, 0x50, 0x12, 0x18, 0x0a,
	0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x33, 0x0a, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x3e, 0x0a,
	0x0d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0c, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x1a, 0x69, 0x0a,
	0x04, 0x47, 0x52, 0x50, 0x43, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12,
	0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61,
	0x64, 0x64, 0x72, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0xcd, 0x01, 0x0a, 0x05, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x69, 0x5f, 0x61, 0x70, 0x69,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x69, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x67, 0x69, 0x6e, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x67, 0x69,
	0x6e, 0x52, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x30, 0x0a, 0x08, 0x74, 0x65,
	0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x67, 0x72,
	0x61, 0x6d, 0x52, 0x08, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x2a, 0x0a, 0x06,
	0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x65, 0x43, 0x68, 0x61, 0x74,
	0x52, 0x06, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x22, 0x88, 0x01, 0x0a, 0x08, 0x54, 0x65, 0x6c,
	0x65, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x52, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x4b,
	0x65, 0x79, 0x73, 0x22, 0xea, 0x02, 0x0a, 0x06, 0x57, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12, 0x15,
	0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x5f, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x6e,
	0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x65, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x65,
	0x73, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x72,
	0x52, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0b, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x43, 0x0a, 0x0b,
	0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x57,
	0x65, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x61,
	0x79, 0x1a, 0x3d, 0x0a, 0x0f, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x1c, 0x5a, 0x1a, 0x69, 0x6e, 0x66, 0x6f, 0x67, 0x70, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conf_conf_proto_rawDescOnce sync.Once
	file_conf_conf_proto_rawDescData = file_conf_conf_proto_rawDesc
)

func file_conf_conf_proto_rawDescGZIP() []byte {
	file_conf_conf_proto_rawDescOnce.Do(func() {
		file_conf_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_conf_proto_rawDescData)
	})
	return file_conf_conf_proto_rawDescData
}

var file_conf_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_conf_conf_proto_goTypes = []interface{}{
	(*Bootstrap)(nil),           // 0: kratos.api.Bootstrap
	(*Server)(nil),              // 1: kratos.api.Server
	(*Admin)(nil),               // 2: kratos.api.Admin
	(*Telegram)(nil),            // 3: kratos.api.Telegram
	(*WeChat)(nil),              // 4: kratos.api.WeChat
	(*Server_HTTP)(nil),         // 5: kratos.api.Server.HTTP
	(*Server_GRPC)(nil),         // 6: kratos.api.Server.GRPC
	nil,                         // 7: kratos.api.WeChat.AutoReplayEntry
	(*durationpb.Duration)(nil), // 8: google.protobuf.Duration
}
var file_conf_conf_proto_depIdxs = []int32{
	1,  // 0: kratos.api.Bootstrap.server:type_name -> kratos.api.Server
	2,  // 1: kratos.api.Bootstrap.admin:type_name -> kratos.api.Admin
	5,  // 2: kratos.api.Server.http:type_name -> kratos.api.Server.HTTP
	6,  // 3: kratos.api.Server.grpc:type_name -> kratos.api.Server.GRPC
	3,  // 4: kratos.api.Admin.telegram:type_name -> kratos.api.Telegram
	4,  // 5: kratos.api.Admin.wechat:type_name -> kratos.api.WeChat
	7,  // 6: kratos.api.WeChat.auto_replay:type_name -> kratos.api.WeChat.AutoReplayEntry
	8,  // 7: kratos.api.Server.HTTP.timeout:type_name -> google.protobuf.Duration
	8,  // 8: kratos.api.Server.HTTP.proxy_timeout:type_name -> google.protobuf.Duration
	8,  // 9: kratos.api.Server.GRPC.timeout:type_name -> google.protobuf.Duration
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_conf_conf_proto_init() }
func file_conf_conf_proto_init() {
	if File_conf_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_conf_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bootstrap); i {
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
		file_conf_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
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
		file_conf_conf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Admin); i {
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
		file_conf_conf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Telegram); i {
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
		file_conf_conf_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeChat); i {
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
		file_conf_conf_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_HTTP); i {
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
		file_conf_conf_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_GRPC); i {
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
			RawDescriptor: file_conf_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_conf_proto_goTypes,
		DependencyIndexes: file_conf_conf_proto_depIdxs,
		MessageInfos:      file_conf_conf_proto_msgTypes,
	}.Build()
	File_conf_conf_proto = out.File
	file_conf_conf_proto_rawDesc = nil
	file_conf_conf_proto_goTypes = nil
	file_conf_conf_proto_depIdxs = nil
}
