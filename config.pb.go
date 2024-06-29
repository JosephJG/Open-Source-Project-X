// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: transport/internet/config.proto

package internet

import (
	serial "github.com/xtls/xray-core/common/serial"
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

type TransportProtocol int32

const (
	TransportProtocol_TCP          TransportProtocol = 0
	TransportProtocol_UDP          TransportProtocol = 1
	TransportProtocol_MKCP         TransportProtocol = 2
	TransportProtocol_WebSocket    TransportProtocol = 3
	TransportProtocol_HTTP         TransportProtocol = 4
	TransportProtocol_DomainSocket TransportProtocol = 5
	TransportProtocol_HTTPUpgrade  TransportProtocol = 6
	TransportProtocol_SplitHTTP    TransportProtocol = 7
)

// Enum value maps for TransportProtocol.
var (
	TransportProtocol_name = map[int32]string{
		0: "TCP",
		1: "UDP",
		2: "MKCP",
		3: "WebSocket",
		4: "HTTP",
		5: "DomainSocket",
		6: "HTTPUpgrade",
		7: "SplitHTTP",
	}
	TransportProtocol_value = map[string]int32{
		"TCP":          0,
		"UDP":          1,
		"MKCP":         2,
		"WebSocket":    3,
		"HTTP":         4,
		"DomainSocket": 5,
		"HTTPUpgrade":  6,
		"SplitHTTP":    7,
	}
)

func (x TransportProtocol) Enum() *TransportProtocol {
	p := new(TransportProtocol)
	*p = x
	return p
}

func (x TransportProtocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransportProtocol) Descriptor() protoreflect.EnumDescriptor {
	return file_transport_internet_config_proto_enumTypes[0].Descriptor()
}

func (TransportProtocol) Type() protoreflect.EnumType {
	return &file_transport_internet_config_proto_enumTypes[0]
}

func (x TransportProtocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransportProtocol.Descriptor instead.
func (TransportProtocol) EnumDescriptor() ([]byte, []int) {
	return file_transport_internet_config_proto_rawDescGZIP(), []int{0}
}

type DomainStrategy int32

const (
	DomainStrategy_AS_IS      DomainStrategy = 0
	DomainStrategy_USE_IP     DomainStrategy = 1
	DomainStrategy_USE_IP4    DomainStrategy = 2
	DomainStrategy_USE_IP6    DomainStrategy = 3
	DomainStrategy_USE_IP46   DomainStrategy = 4
	DomainStrategy_USE_IP64   DomainStrategy = 5
	DomainStrategy_FORCE_IP   DomainStrategy = 6
	DomainStrategy_FORCE_IP4  DomainStrategy = 7
	DomainStrategy_FORCE_IP6  DomainStrategy = 8
	DomainStrategy_FORCE_IP46 DomainStrategy = 9
	DomainStrategy_FORCE_IP64 DomainStrategy = 10
)

// Enum value maps for DomainStrategy.
var (
	DomainStrategy_name = map[int32]string{
		0:  "AS_IS",
		1:  "USE_IP",
		2:  "USE_IP4",
		3:  "USE_IP6",
		4:  "USE_IP46",
		5:  "USE_IP64",
		6:  "FORCE_IP",
		7:  "FORCE_IP4",
		8:  "FORCE_IP6",
		9:  "FORCE_IP46",
		10: "FORCE_IP64",
	}
	DomainStrategy_value = map[string]int32{
		"AS_IS":      0,
		"USE_IP":     1,
		"USE_IP4":    2,
		"USE_IP6":    3,
		"USE_IP46":   4,
		"USE_IP64":   5,
		"FORCE_IP":   6,
		"FORCE_IP4":  7,
		"FORCE_IP6":  8,
		"FORCE_IP46": 9,
		"FORCE_IP64": 10,
	}
)

func (x DomainStrategy) Enum() *DomainStrategy {
	p := new(DomainStrategy)
	*p = x
	return p
}

func (x DomainStrategy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DomainStrategy) Descriptor() protoreflect.EnumDescriptor {
	return file_transport_internet_config_proto_enumTypes[1].Descriptor()
}

func (DomainStrategy) Type() protoreflect.EnumType {
	return &file_transport_internet_config_proto_enumTypes[1]
}

func (x DomainStrategy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DomainStrategy.Descriptor instead.
func (DomainStrategy) EnumDescriptor() ([]byte, []int) {
	return file_transport_internet_config_proto_rawDescGZIP(), []int{1}
}

type SocketConfig_TProxyMode int32

const (
	// TProxy is off.
	SocketConfig_Off SocketConfig_TProxyMode = 0
	// TProxy mode.
	SocketConfig_TProxy SocketConfig_TProxyMode = 1
	// Redirect mode.
	SocketConfig_Redirect SocketConfig_TProxyMode = 2
)

// Enum value maps for SocketConfig_TProxyMode.
var (
	SocketConfig_TProxyMode_name = map[int32]string{
		0: "Off",
		1: "TProxy",
		2: "Redirect",
	}
	SocketConfig_TProxyMode_value = map[string]int32{
		"Off":      0,
		"TProxy":   1,
		"Redirect": 2,
	}
)

func (x SocketConfig_TProxyMode) Enum() *SocketConfig_TProxyMode {
	p := new(SocketConfig_TProxyMode)
	*p = x
	return p
}

func (x SocketConfig_TProxyMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SocketConfig_TProxyMode) Descriptor() protoreflect.EnumDescriptor {
	return file_transport_internet_config_proto_enumTypes[2].Descriptor()
}

func (SocketConfig_TProxyMode) Type() protoreflect.EnumType {
	return &file_transport_internet_config_proto_enumTypes[2]
}

func (x SocketConfig_TProxyMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SocketConfig_TProxyMode.Descriptor instead.
func (SocketConfig_TProxyMode) EnumDescriptor() ([]byte, []int) {
	return file_transport_internet_config_proto_rawDescGZIP(), []int{3, 0}
}

type TransportConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type of network that this settings supports.
	// Deprecated. Use the string form below.
	//
	// Deprecated: Marked as deprecated in transport/internet/config.proto.
	Protocol TransportProtocol `protobuf:"varint,1,opt,name=protocol,proto3,enum=xray.transport.internet.TransportProtocol" json:"protocol,omitempty"`
	// Type of network that this settings supports.
	ProtocolName string `protobuf:"bytes,3,opt,name=protocol_name,json=protocolName,proto3" json:"protocol_name,omitempty"`
	// Specific settings. Must be of the transports.
	Settings *serial.TypedMessage `protobuf:"bytes,2,opt,name=settings,proto3" json:"settings,omitempty"`
}

func (x *TransportConfig) Reset() {
	*x = TransportConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_internet_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransportConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransportConfig) ProtoMessage() {}

func (x *TransportConfig) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransportConfig.ProtoReflect.Descriptor instead.
func (*TransportConfig) Descriptor() ([]byte, []int) {
	return file_transport_internet_config_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Marked as deprecated in transport/internet/config.proto.
func (x *TransportConfig) GetProtocol() TransportProtocol {
	if x != nil {
		return x.Protocol
	}
	return TransportProtocol_TCP
}

func (x *TransportConfig) GetProtocolName() string {
	if x != nil {
		return x.ProtocolName
	}
	return ""
}

func (x *TransportConfig) GetSettings() *serial.TypedMessage {
	if x != nil {
		return x.Settings
	}
	return nil
}

type StreamConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Effective network. Deprecated. Use the string form below.
	//
	// Deprecated: Marked as deprecated in transport/internet/config.proto.
	Protocol TransportProtocol `protobuf:"varint,1,opt,name=protocol,proto3,enum=xray.transport.internet.TransportProtocol" json:"protocol,omitempty"`
	// Effective network.
	ProtocolName      string             `protobuf:"bytes,5,opt,name=protocol_name,json=protocolName,proto3" json:"protocol_name,omitempty"`
	TransportSettings []*TransportConfig `protobuf:"bytes,2,rep,name=transport_settings,json=transportSettings,proto3" json:"transport_settings,omitempty"`
	// Type of security. Must be a message name of the settings proto.
	SecurityType string `protobuf:"bytes,3,opt,name=security_type,json=securityType,proto3" json:"security_type,omitempty"`
	// Settings for transport security. For now the only choice is TLS.
	SecuritySettings []*serial.TypedMessage `protobuf:"bytes,4,rep,name=security_settings,json=securitySettings,proto3" json:"security_settings,omitempty"`
	SocketSettings   *SocketConfig          `protobuf:"bytes,6,opt,name=socket_settings,json=socketSettings,proto3" json:"socket_settings,omitempty"`
}

func (x *StreamConfig) Reset() {
	*x = StreamConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_internet_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamConfig) ProtoMessage() {}

func (x *StreamConfig) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamConfig.ProtoReflect.Descriptor instead.
func (*StreamConfig) Descriptor() ([]byte, []int) {
	return file_transport_internet_config_proto_rawDescGZIP(), []int{1}
}

// Deprecated: Marked as deprecated in transport/internet/config.proto.
func (x *StreamConfig) GetProtocol() TransportProtocol {
	if x != nil {
		return x.Protocol
	}
	return TransportProtocol_TCP
}

func (x *StreamConfig) GetProtocolName() string {
	if x != nil {
		return x.ProtocolName
	}
	return ""
}

func (x *StreamConfig) GetTransportSettings() []*TransportConfig {
	if x != nil {
		return x.TransportSettings
	}
	return nil
}

func (x *StreamConfig) GetSecurityType() string {
	if x != nil {
		return x.SecurityType
	}
	return ""
}

func (x *StreamConfig) GetSecuritySettings() []*serial.TypedMessage {
	if x != nil {
		return x.SecuritySettings
	}
	return nil
}

func (x *StreamConfig) GetSocketSettings() *SocketConfig {
	if x != nil {
		return x.SocketSettings
	}
	return nil
}

type ProxyConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag                 string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	TransportLayerProxy bool   `protobuf:"varint,2,opt,name=transportLayerProxy,proto3" json:"transportLayerProxy,omitempty"`
}

func (x *ProxyConfig) Reset() {
	*x = ProxyConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_internet_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyConfig) ProtoMessage() {}

func (x *ProxyConfig) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyConfig.ProtoReflect.Descriptor instead.
func (*ProxyConfig) Descriptor() ([]byte, []int) {
	return file_transport_internet_config_proto_rawDescGZIP(), []int{2}
}

func (x *ProxyConfig) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *ProxyConfig) GetTransportLayerProxy() bool {
	if x != nil {
		return x.TransportLayerProxy
	}
	return false
}

// SocketConfig is options to be applied on network sockets.
type SocketConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Mark of the connection. If non-zero, the value will be set to SO_MARK.
	Mark int32 `protobuf:"varint,1,opt,name=mark,proto3" json:"mark,omitempty"`
	// TFO is the state of TFO settings.
	Tfo int32 `protobuf:"varint,2,opt,name=tfo,proto3" json:"tfo,omitempty"`
	// TProxy is for enabling TProxy socket option.
	Tproxy SocketConfig_TProxyMode `protobuf:"varint,3,opt,name=tproxy,proto3,enum=xray.transport.internet.SocketConfig_TProxyMode" json:"tproxy,omitempty"`
	// ReceiveOriginalDestAddress is for enabling IP_RECVORIGDSTADDR socket
	// option. This option is for UDP only.
	ReceiveOriginalDestAddress bool           `protobuf:"varint,4,opt,name=receive_original_dest_address,json=receiveOriginalDestAddress,proto3" json:"receive_original_dest_address,omitempty"`
	BindAddress                []byte         `protobuf:"bytes,5,opt,name=bind_address,json=bindAddress,proto3" json:"bind_address,omitempty"`
	BindPort                   uint32         `protobuf:"varint,6,opt,name=bind_port,json=bindPort,proto3" json:"bind_port,omitempty"`
	AcceptProxyProtocol        bool           `protobuf:"varint,7,opt,name=accept_proxy_protocol,json=acceptProxyProtocol,proto3" json:"accept_proxy_protocol,omitempty"`
	DomainStrategy             DomainStrategy `protobuf:"varint,8,opt,name=domain_strategy,json=domainStrategy,proto3,enum=xray.transport.internet.DomainStrategy" json:"domain_strategy,omitempty"`
	DialerProxy                string         `protobuf:"bytes,9,opt,name=dialer_proxy,json=dialerProxy,proto3" json:"dialer_proxy,omitempty"`
	TcpKeepAliveInterval       int32          `protobuf:"varint,10,opt,name=tcp_keep_alive_interval,json=tcpKeepAliveInterval,proto3" json:"tcp_keep_alive_interval,omitempty"`
	TcpKeepAliveIdle           int32          `protobuf:"varint,11,opt,name=tcp_keep_alive_idle,json=tcpKeepAliveIdle,proto3" json:"tcp_keep_alive_idle,omitempty"`
	TcpCongestion              string         `protobuf:"bytes,12,opt,name=tcp_congestion,json=tcpCongestion,proto3" json:"tcp_congestion,omitempty"`
	Interface                  string         `protobuf:"bytes,13,opt,name=interface,proto3" json:"interface,omitempty"`
	V6Only                     bool           `protobuf:"varint,14,opt,name=v6only,proto3" json:"v6only,omitempty"`
	TcpWindowClamp             int32          `protobuf:"varint,15,opt,name=tcp_window_clamp,json=tcpWindowClamp,proto3" json:"tcp_window_clamp,omitempty"`
	TcpUserTimeout             int32          `protobuf:"varint,16,opt,name=tcp_user_timeout,json=tcpUserTimeout,proto3" json:"tcp_user_timeout,omitempty"`
	TcpMaxSeg                  int32          `protobuf:"varint,17,opt,name=tcp_max_seg,json=tcpMaxSeg,proto3" json:"tcp_max_seg,omitempty"`
	TcpNoDelay                 bool           `protobuf:"varint,18,opt,name=tcp_no_delay,json=tcpNoDelay,proto3" json:"tcp_no_delay,omitempty"`
	TcpMptcp                   bool           `protobuf:"varint,19,opt,name=tcp_mptcp,json=tcpMptcp,proto3" json:"tcp_mptcp,omitempty"`
}

func (x *SocketConfig) Reset() {
	*x = SocketConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_internet_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocketConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocketConfig) ProtoMessage() {}

func (x *SocketConfig) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocketConfig.ProtoReflect.Descriptor instead.
func (*SocketConfig) Descriptor() ([]byte, []int) {
	return file_transport_internet_config_proto_rawDescGZIP(), []int{3}
}

func (x *SocketConfig) GetMark() int32 {
	if x != nil {
		return x.Mark
	}
	return 0
}

func (x *SocketConfig) GetTfo() int32 {
	if x != nil {
		return x.Tfo
	}
	return 0
}

func (x *SocketConfig) GetTproxy() SocketConfig_TProxyMode {
	if x != nil {
		return x.Tproxy
	}
	return SocketConfig_Off
}

func (x *SocketConfig) GetReceiveOriginalDestAddress() bool {
	if x != nil {
		return x.ReceiveOriginalDestAddress
	}
	return false
}

func (x *SocketConfig) GetBindAddress() []byte {
	if x != nil {
		return x.BindAddress
	}
	return nil
}

func (x *SocketConfig) GetBindPort() uint32 {
	if x != nil {
		return x.BindPort
	}
	return 0
}

func (x *SocketConfig) GetAcceptProxyProtocol() bool {
	if x != nil {
		return x.AcceptProxyProtocol
	}
	return false
}

func (x *SocketConfig) GetDomainStrategy() DomainStrategy {
	if x != nil {
		return x.DomainStrategy
	}
	return DomainStrategy_AS_IS
}

func (x *SocketConfig) GetDialerProxy() string {
	if x != nil {
		return x.DialerProxy
	}
	return ""
}

func (x *SocketConfig) GetTcpKeepAliveInterval() int32 {
	if x != nil {
		return x.TcpKeepAliveInterval
	}
	return 0
}

func (x *SocketConfig) GetTcpKeepAliveIdle() int32 {
	if x != nil {
		return x.TcpKeepAliveIdle
	}
	return 0
}

func (x *SocketConfig) GetTcpCongestion() string {
	if x != nil {
		return x.TcpCongestion
	}
	return ""
}

func (x *SocketConfig) GetInterface() string {
	if x != nil {
		return x.Interface
	}
	return ""
}

func (x *SocketConfig) GetV6Only() bool {
	if x != nil {
		return x.V6Only
	}
	return false
}

func (x *SocketConfig) GetTcpWindowClamp() int32 {
	if x != nil {
		return x.TcpWindowClamp
	}
	return 0
}

func (x *SocketConfig) GetTcpUserTimeout() int32 {
	if x != nil {
		return x.TcpUserTimeout
	}
	return 0
}

func (x *SocketConfig) GetTcpMaxSeg() int32 {
	if x != nil {
		return x.TcpMaxSeg
	}
	return 0
}

func (x *SocketConfig) GetTcpNoDelay() bool {
	if x != nil {
		return x.TcpNoDelay
	}
	return false
}

func (x *SocketConfig) GetTcpMptcp() bool {
	if x != nil {
		return x.TcpMptcp
	}
	return false
}

var File_transport_internet_config_proto protoreflect.FileDescriptor

var file_transport_internet_config_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x17, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x1a, 0x21, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x64, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x01,
	0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x4a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x42,
	0x02, 0x18, 0x01, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x23, 0x0a,
	0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x22, 0x9c, 0x03, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x4a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x42,
	0x02, 0x18, 0x01, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x23, 0x0a,
	0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x57, 0x0a, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x4d, 0x0a, 0x11, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x78, 0x72,
	0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x10, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x4e, 0x0a, 0x0f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x2e, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x0e, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22,
	0x51, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x10,
	0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67,
	0x12, 0x30, 0x0a, 0x13, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x61, 0x79,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x22, 0xd1, 0x06, 0x0a, 0x0c, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x66, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x66, 0x6f, 0x12, 0x48, 0x0a, 0x06, 0x74, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x78, 0x72, 0x61, 0x79,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x2e, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x54, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x74, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x12, 0x41, 0x0a, 0x1d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x5f, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1a, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x44, 0x65, 0x73, 0x74, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x69, 0x6e, 0x64, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x62, 0x69,
	0x6e, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x69, 0x6e,
	0x64, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x62, 0x69,
	0x6e, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x50, 0x0a, 0x0f, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x0e, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x21, 0x0a, 0x0c,
	0x64, 0x69, 0x61, 0x6c, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x61, 0x6c, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12,
	0x35, 0x0a, 0x17, 0x74, 0x63, 0x70, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x61, 0x6c, 0x69, 0x76,
	0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x14, 0x74, 0x63, 0x70, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x2d, 0x0a, 0x13, 0x74, 0x63, 0x70, 0x5f, 0x6b, 0x65,
	0x65, 0x70, 0x5f, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x64, 0x6c, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x10, 0x74, 0x63, 0x70, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76,
	0x65, 0x49, 0x64, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x63, 0x70, 0x5f, 0x63, 0x6f, 0x6e,
	0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74,
	0x63, 0x70, 0x43, 0x6f, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x36,
	0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x76, 0x36, 0x6f, 0x6e,
	0x6c, 0x79, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x63, 0x70, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77,
	0x5f, 0x63, 0x6c, 0x61, 0x6d, 0x70, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x74, 0x63,
	0x70, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x43, 0x6c, 0x61, 0x6d, 0x70, 0x12, 0x28, 0x0a, 0x10,
	0x74, 0x63, 0x70, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x74, 0x63, 0x70, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x1e, 0x0a, 0x0b, 0x74, 0x63, 0x70, 0x5f, 0x6d, 0x61,
	0x78, 0x5f, 0x73, 0x65, 0x67, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x63, 0x70,
	0x4d, 0x61, 0x78, 0x53, 0x65, 0x67, 0x12, 0x20, 0x0a, 0x0c, 0x74, 0x63, 0x70, 0x5f, 0x6e, 0x6f,
	0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x74, 0x63,
	0x70, 0x4e, 0x6f, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x63, 0x70, 0x5f,
	0x6d, 0x70, 0x74, 0x63, 0x70, 0x18, 0x13, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x74, 0x63, 0x70,
	0x4d, 0x70, 0x74, 0x63, 0x70, 0x22, 0x2f, 0x0a, 0x0a, 0x54, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4d,
	0x6f, 0x64, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x66, 0x66, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x54, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x65, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x10, 0x02, 0x2a, 0x7a, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x07, 0x0a, 0x03, 0x54,
	0x43, 0x50, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x55, 0x44, 0x50, 0x10, 0x01, 0x12, 0x08, 0x0a,
	0x04, 0x4d, 0x4b, 0x43, 0x50, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x57, 0x65, 0x62, 0x53, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x54, 0x54, 0x50, 0x10, 0x04,
	0x12, 0x10, 0x0a, 0x0c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74,
	0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x48, 0x54, 0x54, 0x50, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64,
	0x65, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x48, 0x54, 0x54, 0x50,
	0x10, 0x07, 0x2a, 0xa9, 0x01, 0x0a, 0x0e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x53, 0x5f, 0x49, 0x53, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x55, 0x53, 0x45, 0x5f, 0x49, 0x50, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x53, 0x45, 0x5f, 0x49, 0x50, 0x34, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x53, 0x45,
	0x5f, 0x49, 0x50, 0x36, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x53, 0x45, 0x5f, 0x49, 0x50,
	0x34, 0x36, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x53, 0x45, 0x5f, 0x49, 0x50, 0x36, 0x34,
	0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x5f, 0x49, 0x50, 0x10, 0x06,
	0x12, 0x0d, 0x0a, 0x09, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x5f, 0x49, 0x50, 0x34, 0x10, 0x07, 0x12,
	0x0d, 0x0a, 0x09, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x5f, 0x49, 0x50, 0x36, 0x10, 0x08, 0x12, 0x0e,
	0x0a, 0x0a, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x5f, 0x49, 0x50, 0x34, 0x36, 0x10, 0x09, 0x12, 0x0e,
	0x0a, 0x0a, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x5f, 0x49, 0x50, 0x36, 0x34, 0x10, 0x0a, 0x42, 0x67,
	0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x50, 0x01, 0x5a,
	0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x74, 0x6c, 0x73,
	0x2f, 0x78, 0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0xaa, 0x02, 0x17,
	0x58, 0x72, 0x61, 0x79, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transport_internet_config_proto_rawDescOnce sync.Once
	file_transport_internet_config_proto_rawDescData = file_transport_internet_config_proto_rawDesc
)

func file_transport_internet_config_proto_rawDescGZIP() []byte {
	file_transport_internet_config_proto_rawDescOnce.Do(func() {
		file_transport_internet_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_transport_internet_config_proto_rawDescData)
	})
	return file_transport_internet_config_proto_rawDescData
}

var file_transport_internet_config_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_transport_internet_config_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_transport_internet_config_proto_goTypes = []interface{}{
	(TransportProtocol)(0),       // 0: xray.transport.internet.TransportProtocol
	(DomainStrategy)(0),          // 1: xray.transport.internet.DomainStrategy
	(SocketConfig_TProxyMode)(0), // 2: xray.transport.internet.SocketConfig.TProxyMode
	(*TransportConfig)(nil),      // 3: xray.transport.internet.TransportConfig
	(*StreamConfig)(nil),         // 4: xray.transport.internet.StreamConfig
	(*ProxyConfig)(nil),          // 5: xray.transport.internet.ProxyConfig
	(*SocketConfig)(nil),         // 6: xray.transport.internet.SocketConfig
	(*serial.TypedMessage)(nil),  // 7: xray.common.serial.TypedMessage
}
var file_transport_internet_config_proto_depIdxs = []int32{
	0, // 0: xray.transport.internet.TransportConfig.protocol:type_name -> xray.transport.internet.TransportProtocol
	7, // 1: xray.transport.internet.TransportConfig.settings:type_name -> xray.common.serial.TypedMessage
	0, // 2: xray.transport.internet.StreamConfig.protocol:type_name -> xray.transport.internet.TransportProtocol
	3, // 3: xray.transport.internet.StreamConfig.transport_settings:type_name -> xray.transport.internet.TransportConfig
	7, // 4: xray.transport.internet.StreamConfig.security_settings:type_name -> xray.common.serial.TypedMessage
	6, // 5: xray.transport.internet.StreamConfig.socket_settings:type_name -> xray.transport.internet.SocketConfig
	2, // 6: xray.transport.internet.SocketConfig.tproxy:type_name -> xray.transport.internet.SocketConfig.TProxyMode
	1, // 7: xray.transport.internet.SocketConfig.domain_strategy:type_name -> xray.transport.internet.DomainStrategy
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_transport_internet_config_proto_init() }
func file_transport_internet_config_proto_init() {
	if File_transport_internet_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transport_internet_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransportConfig); i {
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
		file_transport_internet_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamConfig); i {
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
		file_transport_internet_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyConfig); i {
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
		file_transport_internet_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocketConfig); i {
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
			RawDescriptor: file_transport_internet_config_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transport_internet_config_proto_goTypes,
		DependencyIndexes: file_transport_internet_config_proto_depIdxs,
		EnumInfos:         file_transport_internet_config_proto_enumTypes,
		MessageInfos:      file_transport_internet_config_proto_msgTypes,
	}.Build()
	File_transport_internet_config_proto = out.File
	file_transport_internet_config_proto_rawDesc = nil
	file_transport_internet_config_proto_goTypes = nil
	file_transport_internet_config_proto_depIdxs = nil
}