//
// Copyright 2017 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: mpls/mpls.proto

package mpls

import (
	_ "github.com/openconfig/gnoi/types"
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

type ClearLSPRequest_Mode int32

const (
	ClearLSPRequest_DEFAULT ClearLSPRequest_Mode = 0 // Same as NONAGGRESSIVE.
	// Reoptimize the LSP using the current bandwidth.
	ClearLSPRequest_NONAGGRESSIVE ClearLSPRequest_Mode = 0
	// Reoptimize the LSP using the current bandwidth. Only use IGP metric in
	// calcuation.
	ClearLSPRequest_AGGRESSIVE ClearLSPRequest_Mode = 1
	// Reset and restart all LSPs that originated from this routing device.
	ClearLSPRequest_RESET ClearLSPRequest_Mode = 2
	// Apply the highest bandwidth collected on a tunnel without waiting for
	// the current application period to end. Only use IGP metric in
	// calcuation.
	ClearLSPRequest_AUTOBW_AGGRESSIVE ClearLSPRequest_Mode = 3
	// Apply the highest bandwidth collected on a tunnel without waiting for
	// the current application period to end.
	ClearLSPRequest_AUTOBW_NONAGGRESSIVE ClearLSPRequest_Mode = 4
)

// Enum value maps for ClearLSPRequest_Mode.
var (
	ClearLSPRequest_Mode_name = map[int32]string{
		0: "DEFAULT",
		// Duplicate value: 0: "NONAGGRESSIVE",
		1: "AGGRESSIVE",
		2: "RESET",
		3: "AUTOBW_AGGRESSIVE",
		4: "AUTOBW_NONAGGRESSIVE",
	}
	ClearLSPRequest_Mode_value = map[string]int32{
		"DEFAULT":              0,
		"NONAGGRESSIVE":        0,
		"AGGRESSIVE":           1,
		"RESET":                2,
		"AUTOBW_AGGRESSIVE":    3,
		"AUTOBW_NONAGGRESSIVE": 4,
	}
)

func (x ClearLSPRequest_Mode) Enum() *ClearLSPRequest_Mode {
	p := new(ClearLSPRequest_Mode)
	*p = x
	return p
}

func (x ClearLSPRequest_Mode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClearLSPRequest_Mode) Descriptor() protoreflect.EnumDescriptor {
	return file_mpls_mpls_proto_enumTypes[0].Descriptor()
}

func (ClearLSPRequest_Mode) Type() protoreflect.EnumType {
	return &file_mpls_mpls_proto_enumTypes[0]
}

func (x ClearLSPRequest_Mode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClearLSPRequest_Mode.Descriptor instead.
func (ClearLSPRequest_Mode) EnumDescriptor() ([]byte, []int) {
	return file_mpls_mpls_proto_rawDescGZIP(), []int{0, 0}
}

type MPLSPingRequest_ReplyMode int32

const (
	// Reply via an IPv4 packet to this system.
	MPLSPingRequest_IPV4 MPLSPingRequest_ReplyMode = 0
	// Reply with a labeled packet with the router alert bit set.
	MPLSPingRequest_ROUTER_ALERT MPLSPingRequest_ReplyMode = 1
)

// Enum value maps for MPLSPingRequest_ReplyMode.
var (
	MPLSPingRequest_ReplyMode_name = map[int32]string{
		0: "IPV4",
		1: "ROUTER_ALERT",
	}
	MPLSPingRequest_ReplyMode_value = map[string]int32{
		"IPV4":         0,
		"ROUTER_ALERT": 1,
	}
)

func (x MPLSPingRequest_ReplyMode) Enum() *MPLSPingRequest_ReplyMode {
	p := new(MPLSPingRequest_ReplyMode)
	*p = x
	return p
}

func (x MPLSPingRequest_ReplyMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MPLSPingRequest_ReplyMode) Descriptor() protoreflect.EnumDescriptor {
	return file_mpls_mpls_proto_enumTypes[1].Descriptor()
}

func (MPLSPingRequest_ReplyMode) Type() protoreflect.EnumType {
	return &file_mpls_mpls_proto_enumTypes[1]
}

func (x MPLSPingRequest_ReplyMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MPLSPingRequest_ReplyMode.Descriptor instead.
func (MPLSPingRequest_ReplyMode) EnumDescriptor() ([]byte, []int) {
	return file_mpls_mpls_proto_rawDescGZIP(), []int{6, 0}
}

type MPLSPingResponse_EchoResponseCode int32

const (
	// A successful echo response was received.
	MPLSPingResponse_SUCCESS MPLSPingResponse_EchoResponseCode = 0
	// The MPLS ping packet was not sent, for an unknown reason.
	MPLSPingResponse_NOT_SENT MPLSPingResponse_EchoResponseCode = 1
	// The local system timed out waiting for an LSP ping response.
	MPLSPingResponse_TIMEOUT MPLSPingResponse_EchoResponseCode = 2 // TODO(robjs): Add additional error codes.
)

// Enum value maps for MPLSPingResponse_EchoResponseCode.
var (
	MPLSPingResponse_EchoResponseCode_name = map[int32]string{
		0: "SUCCESS",
		1: "NOT_SENT",
		2: "TIMEOUT",
	}
	MPLSPingResponse_EchoResponseCode_value = map[string]int32{
		"SUCCESS":  0,
		"NOT_SENT": 1,
		"TIMEOUT":  2,
	}
)

func (x MPLSPingResponse_EchoResponseCode) Enum() *MPLSPingResponse_EchoResponseCode {
	p := new(MPLSPingResponse_EchoResponseCode)
	*p = x
	return p
}

func (x MPLSPingResponse_EchoResponseCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MPLSPingResponse_EchoResponseCode) Descriptor() protoreflect.EnumDescriptor {
	return file_mpls_mpls_proto_enumTypes[2].Descriptor()
}

func (MPLSPingResponse_EchoResponseCode) Type() protoreflect.EnumType {
	return &file_mpls_mpls_proto_enumTypes[2]
}

func (x MPLSPingResponse_EchoResponseCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MPLSPingResponse_EchoResponseCode.Descriptor instead.
func (MPLSPingResponse_EchoResponseCode) EnumDescriptor() ([]byte, []int) {
	return file_mpls_mpls_proto_rawDescGZIP(), []int{7, 0}
}

// Request to clear a single tunnel on a target device.
type ClearLSPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                      // Name of the tunnel to clear.
	Mode ClearLSPRequest_Mode `protobuf:"varint,3,opt,name=mode,proto3,enum=gnoi.mpls.ClearLSPRequest_Mode" json:"mode,omitempty"` // Tunnel clearing mode.
}

func (x *ClearLSPRequest) Reset() {
	*x = ClearLSPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpls_mpls_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearLSPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearLSPRequest) ProtoMessage() {}

func (x *ClearLSPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mpls_mpls_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearLSPRequest.ProtoReflect.Descriptor instead.
func (*ClearLSPRequest) Descriptor() ([]byte, []int) {
	return file_mpls_mpls_proto_rawDescGZIP(), []int{0}
}

func (x *ClearLSPRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClearLSPRequest) GetMode() ClearLSPRequest_Mode {
	if x != nil {
		return x.Mode
	}
	return ClearLSPRequest_DEFAULT
}

type ClearLSPResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClearLSPResponse) Reset() {
	*x = ClearLSPResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpls_mpls_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearLSPResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearLSPResponse) ProtoMessage() {}

func (x *ClearLSPResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mpls_mpls_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearLSPResponse.ProtoReflect.Descriptor instead.
func (*ClearLSPResponse) Descriptor() ([]byte, []int) {
	return file_mpls_mpls_proto_rawDescGZIP(), []int{1}
}

// Request to clear a single tunnel counters on a target device.
type ClearLSPCountersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // Name of the tunnel to clear.
}

func (x *ClearLSPCountersRequest) Reset() {
	*x = ClearLSPCountersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpls_mpls_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearLSPCountersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearLSPCountersRequest) ProtoMessage() {}

func (x *ClearLSPCountersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mpls_mpls_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearLSPCountersRequest.ProtoReflect.Descriptor instead.
func (*ClearLSPCountersRequest) Descriptor() ([]byte, []int) {
	return file_mpls_mpls_proto_rawDescGZIP(), []int{2}
}

func (x *ClearLSPCountersRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ClearLSPCountersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClearLSPCountersResponse) Reset() {
	*x = ClearLSPCountersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpls_mpls_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearLSPCountersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearLSPCountersResponse) ProtoMessage() {}

func (x *ClearLSPCountersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mpls_mpls_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearLSPCountersResponse.ProtoReflect.Descriptor instead.
func (*ClearLSPCountersResponse) Descriptor() ([]byte, []int) {
	return file_mpls_mpls_proto_rawDescGZIP(), []int{3}
}

type MPLSPingPWEDestination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The address of the egress LER that the MPLS ping should be sent on when
	// destined to a PWE service.
	Eler string `protobuf:"bytes,1,opt,name=eler,proto3" json:"eler,omitempty"`
	// The virtual circuit ID for the PWE via which the ping should be sent.
	Vcid uint32 `protobuf:"varint,2,opt,name=vcid,proto3" json:"vcid,omitempty"`
}

func (x *MPLSPingPWEDestination) Reset() {
	*x = MPLSPingPWEDestination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpls_mpls_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MPLSPingPWEDestination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MPLSPingPWEDestination) ProtoMessage() {}

func (x *MPLSPingPWEDestination) ProtoReflect() protoreflect.Message {
	mi := &file_mpls_mpls_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MPLSPingPWEDestination.ProtoReflect.Descriptor instead.
func (*MPLSPingPWEDestination) Descriptor() ([]byte, []int) {
	return file_mpls_mpls_proto_rawDescGZIP(), []int{4}
}

func (x *MPLSPingPWEDestination) GetEler() string {
	if x != nil {
		return x.Eler
	}
	return ""
}

func (x *MPLSPingPWEDestination) GetVcid() uint32 {
	if x != nil {
		return x.Vcid
	}
	return 0
}

// MPLSPingRSVPTEDestination specifies the destination for an MPLS Ping in
// terms of an absolute specification of an RSVP-TE LSP. It can be used to
// identify an individual RSVP-TE session via which a ping should be sent.
type MPLSPingRSVPTEDestination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The IPv4 or IPv6 address used by the system initiating (acting as the
	// head-end) of the RSVP-TE LSP.
	Src string `protobuf:"bytes,1,opt,name=src,proto3" json:"src,omitempty"`
	// The IPv4 or IPv6 address used by the system terminating (acting as the
	// tail-end) of the RSVP-TE LSP.
	Dst string `protobuf:"bytes,2,opt,name=dst,proto3" json:"dst,omitempty"`
	// The extended tunnel ID of the RSVP-TE LSP, expressed as an unsigned, 32b
	// integer.
	ExtendedTunnelId uint32 `protobuf:"varint,3,opt,name=extended_tunnel_id,json=extendedTunnelId,proto3" json:"extended_tunnel_id,omitempty"`
}

func (x *MPLSPingRSVPTEDestination) Reset() {
	*x = MPLSPingRSVPTEDestination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpls_mpls_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MPLSPingRSVPTEDestination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MPLSPingRSVPTEDestination) ProtoMessage() {}

func (x *MPLSPingRSVPTEDestination) ProtoReflect() protoreflect.Message {
	mi := &file_mpls_mpls_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MPLSPingRSVPTEDestination.ProtoReflect.Descriptor instead.
func (*MPLSPingRSVPTEDestination) Descriptor() ([]byte, []int) {
	return file_mpls_mpls_proto_rawDescGZIP(), []int{5}
}

func (x *MPLSPingRSVPTEDestination) GetSrc() string {
	if x != nil {
		return x.Src
	}
	return ""
}

func (x *MPLSPingRSVPTEDestination) GetDst() string {
	if x != nil {
		return x.Dst
	}
	return ""
}

func (x *MPLSPingRSVPTEDestination) GetExtendedTunnelId() uint32 {
	if x != nil {
		return x.ExtendedTunnelId
	}
	return 0
}

// MPLSPingRequest specifies the parameters that should be used as input from
// the client, to a system that is initiating an RFC4379 MPLS ping request.
type MPLSPingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// One field within the destination field should be specified to determine
	// the destination of the LSP ping.
	//
	// Types that are assignable to Destination:
	//	*MPLSPingRequest_LdpFec
	//	*MPLSPingRequest_Fec129Pwe
	//	*MPLSPingRequest_RsvpteLspName
	//	*MPLSPingRequest_RsvpteLsp
	Destination isMPLSPingRequest_Destination `protobuf_oneof:"destination"`
	// How the target LER should respond to the LSP ping.
	ReplyMode MPLSPingRequest_ReplyMode `protobuf:"varint,6,opt,name=reply_mode,json=replyMode,proto3,enum=gnoi.mpls.MPLSPingRequest_ReplyMode" json:"reply_mode,omitempty"`
	// The number of MPLS echo request packets to send.
	Count uint32 `protobuf:"varint,7,opt,name=count,proto3" json:"count,omitempty"` // default=3
	// The size (in bytes) of each MPLS echo request packet.
	Size uint32 `protobuf:"varint,8,opt,name=size,proto3" json:"size,omitempty"` // default=64
	// The source IPv4 address that should be used in the request packet.
	SourceAddress string `protobuf:"bytes,9,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	// The MPLS TTL that should be set in the packets sent.
	MplsTtl uint32 `protobuf:"varint,10,opt,name=mpls_ttl,json=mplsTtl,proto3" json:"mpls_ttl,omitempty"`
	// The value of the traffic class (TC, formerly known as EXP) bits that
	// should be set in the MPLS ping packets.
	TrafficClass uint32 `protobuf:"varint,11,opt,name=traffic_class,json=trafficClass,proto3" json:"traffic_class,omitempty"`
}

func (x *MPLSPingRequest) Reset() {
	*x = MPLSPingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpls_mpls_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MPLSPingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MPLSPingRequest) ProtoMessage() {}

func (x *MPLSPingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mpls_mpls_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MPLSPingRequest.ProtoReflect.Descriptor instead.
func (*MPLSPingRequest) Descriptor() ([]byte, []int) {
	return file_mpls_mpls_proto_rawDescGZIP(), []int{6}
}

func (m *MPLSPingRequest) GetDestination() isMPLSPingRequest_Destination {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (x *MPLSPingRequest) GetLdpFec() string {
	if x, ok := x.GetDestination().(*MPLSPingRequest_LdpFec); ok {
		return x.LdpFec
	}
	return ""
}

func (x *MPLSPingRequest) GetFec129Pwe() *MPLSPingPWEDestination {
	if x, ok := x.GetDestination().(*MPLSPingRequest_Fec129Pwe); ok {
		return x.Fec129Pwe
	}
	return nil
}

func (x *MPLSPingRequest) GetRsvpteLspName() string {
	if x, ok := x.GetDestination().(*MPLSPingRequest_RsvpteLspName); ok {
		return x.RsvpteLspName
	}
	return ""
}

func (x *MPLSPingRequest) GetRsvpteLsp() *MPLSPingRSVPTEDestination {
	if x, ok := x.GetDestination().(*MPLSPingRequest_RsvpteLsp); ok {
		return x.RsvpteLsp
	}
	return nil
}

func (x *MPLSPingRequest) GetReplyMode() MPLSPingRequest_ReplyMode {
	if x != nil {
		return x.ReplyMode
	}
	return MPLSPingRequest_IPV4
}

func (x *MPLSPingRequest) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *MPLSPingRequest) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *MPLSPingRequest) GetSourceAddress() string {
	if x != nil {
		return x.SourceAddress
	}
	return ""
}

func (x *MPLSPingRequest) GetMplsTtl() uint32 {
	if x != nil {
		return x.MplsTtl
	}
	return 0
}

func (x *MPLSPingRequest) GetTrafficClass() uint32 {
	if x != nil {
		return x.TrafficClass
	}
	return 0
}

type isMPLSPingRequest_Destination interface {
	isMPLSPingRequest_Destination()
}

type MPLSPingRequest_LdpFec struct {
	// The LDP forwarding equivalence class that the ping should be sent to
	// expressed as an IPv4 or IPv6 prefix.
	LdpFec string `protobuf:"bytes,1,opt,name=ldp_fec,json=ldpFec,proto3,oneof"`
}

type MPLSPingRequest_Fec129Pwe struct {
	// The FEC129 PWE to which the LSP ping should be sent.
	Fec129Pwe *MPLSPingPWEDestination `protobuf:"bytes,2,opt,name=fec129_pwe,json=fec129Pwe,proto3,oneof"`
}

type MPLSPingRequest_RsvpteLspName struct {
	// The name of an RSVP-TE LSP via which the ping should be sent. The system
	// should locally resolve the name to a particular RSVP-TE session.
	RsvpteLspName string `protobuf:"bytes,4,opt,name=rsvpte_lsp_name,json=rsvpteLspName,proto3,oneof"`
}

type MPLSPingRequest_RsvpteLsp struct {
	// An exact specification of an RSVP-TE LSP to which the system should send
	// an MPLS ping message.
	RsvpteLsp *MPLSPingRSVPTEDestination `protobuf:"bytes,5,opt,name=rsvpte_lsp,json=rsvpteLsp,proto3,oneof"` // TODO(robjs): L3VPN, BGP-LU destination types. See RFC4379.
}

func (*MPLSPingRequest_LdpFec) isMPLSPingRequest_Destination() {}

func (*MPLSPingRequest_Fec129Pwe) isMPLSPingRequest_Destination() {}

func (*MPLSPingRequest_RsvpteLspName) isMPLSPingRequest_Destination() {}

func (*MPLSPingRequest_RsvpteLsp) isMPLSPingRequest_Destination() {}

// MPLSPingResponse (MPLSPong?) is sent from the target to the client based on
// each MPLS Echo Response packet it receives associated with an input MPLSPing
// RPC.
type MPLSPingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seq uint32 `protobuf:"varint,1,opt,name=seq,proto3" json:"seq,omitempty"` // The sequence number of the MPLS Echo Reply.
	// The response that was received from the egress LER.
	Response MPLSPingResponse_EchoResponseCode `protobuf:"varint,2,opt,name=response,proto3,enum=gnoi.mpls.MPLSPingResponse_EchoResponseCode" json:"response,omitempty"`
	// The time (in nanoseconds) between transmission of the Echo Response,
	// and the echo reply.
	ResponseTime uint64 `protobuf:"varint,3,opt,name=response_time,json=responseTime,proto3" json:"response_time,omitempty"`
}

func (x *MPLSPingResponse) Reset() {
	*x = MPLSPingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpls_mpls_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MPLSPingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MPLSPingResponse) ProtoMessage() {}

func (x *MPLSPingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mpls_mpls_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MPLSPingResponse.ProtoReflect.Descriptor instead.
func (*MPLSPingResponse) Descriptor() ([]byte, []int) {
	return file_mpls_mpls_proto_rawDescGZIP(), []int{7}
}

func (x *MPLSPingResponse) GetSeq() uint32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *MPLSPingResponse) GetResponse() MPLSPingResponse_EchoResponseCode {
	if x != nil {
		return x.Response
	}
	return MPLSPingResponse_SUCCESS
}

func (x *MPLSPingResponse) GetResponseTime() uint64 {
	if x != nil {
		return x.ResponseTime
	}
	return 0
}

var File_mpls_mpls_proto protoreflect.FileDescriptor

var file_mpls_mpls_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x70, 0x6c, 0x73, 0x2f, 0x6d, 0x70, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x6d, 0x70, 0x6c, 0x73, 0x1a, 0x2c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2f, 0x67, 0x6e, 0x6f, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x01, 0x0a, 0x0f, 0x43,
	0x6c, 0x65, 0x61, 0x72, 0x4c, 0x53, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1f, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x6d, 0x70, 0x6c, 0x73, 0x2e, 0x43, 0x6c, 0x65,
	0x61, 0x72, 0x4c, 0x53, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x76, 0x0a, 0x04, 0x4d, 0x6f, 0x64, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d,
	0x4e, 0x4f, 0x4e, 0x41, 0x47, 0x47, 0x52, 0x45, 0x53, 0x53, 0x49, 0x56, 0x45, 0x10, 0x00, 0x12,
	0x0e, 0x0a, 0x0a, 0x41, 0x47, 0x47, 0x52, 0x45, 0x53, 0x53, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12,
	0x09, 0x0a, 0x05, 0x52, 0x45, 0x53, 0x45, 0x54, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x55,
	0x54, 0x4f, 0x42, 0x57, 0x5f, 0x41, 0x47, 0x47, 0x52, 0x45, 0x53, 0x53, 0x49, 0x56, 0x45, 0x10,
	0x03, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x55, 0x54, 0x4f, 0x42, 0x57, 0x5f, 0x4e, 0x4f, 0x4e, 0x41,
	0x47, 0x47, 0x52, 0x45, 0x53, 0x53, 0x49, 0x56, 0x45, 0x10, 0x04, 0x1a, 0x02, 0x10, 0x01, 0x22,
	0x12, 0x0a, 0x10, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x4c, 0x53, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x2d, 0x0a, 0x17, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x4c, 0x53, 0x50, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x1a, 0x0a, 0x18, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x4c, 0x53, 0x50, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x40,
	0x0a, 0x16, 0x4d, 0x50, 0x4c, 0x53, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x57, 0x45, 0x44, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x6c, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x6c, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x76, 0x63, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x76, 0x63, 0x69, 0x64,
	0x22, 0x6d, 0x0a, 0x19, 0x4d, 0x50, 0x4c, 0x53, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x53, 0x56, 0x50,
	0x54, 0x45, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a,
	0x03, 0x73, 0x72, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x72, 0x63, 0x12,
	0x10, 0x0a, 0x03, 0x64, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x73,
	0x74, 0x12, 0x2c, 0x0a, 0x12, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x74, 0x75,
	0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x22,
	0xef, 0x03, 0x0a, 0x0f, 0x4d, 0x50, 0x4c, 0x53, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x07, 0x6c, 0x64, 0x70, 0x5f, 0x66, 0x65, 0x63, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x6c, 0x64, 0x70, 0x46, 0x65, 0x63, 0x12, 0x42,
	0x0a, 0x0a, 0x66, 0x65, 0x63, 0x31, 0x32, 0x39, 0x5f, 0x70, 0x77, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x6d, 0x70, 0x6c, 0x73, 0x2e, 0x4d,
	0x50, 0x4c, 0x53, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x57, 0x45, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x09, 0x66, 0x65, 0x63, 0x31, 0x32, 0x39, 0x50,
	0x77, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x73, 0x76, 0x70, 0x74, 0x65, 0x5f, 0x6c, 0x73, 0x70,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x72,
	0x73, 0x76, 0x70, 0x74, 0x65, 0x4c, 0x73, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x0a,
	0x72, 0x73, 0x76, 0x70, 0x74, 0x65, 0x5f, 0x6c, 0x73, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x6d, 0x70, 0x6c, 0x73, 0x2e, 0x4d, 0x50, 0x4c,
	0x53, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x53, 0x56, 0x50, 0x54, 0x45, 0x44, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x09, 0x72, 0x73, 0x76, 0x70, 0x74, 0x65,
	0x4c, 0x73, 0x70, 0x12, 0x43, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x6d,
	0x70, 0x6c, 0x73, 0x2e, 0x4d, 0x50, 0x4c, 0x53, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x72,
	0x65, 0x70, 0x6c, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x70, 0x6c,
	0x73, 0x5f, 0x74, 0x74, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x70, 0x6c,
	0x73, 0x54, 0x74, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x74, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x22, 0x27, 0x0a, 0x09, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x50, 0x56, 0x34, 0x10, 0x00,
	0x12, 0x10, 0x0a, 0x0c, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x52, 0x5f, 0x41, 0x4c, 0x45, 0x52, 0x54,
	0x10, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0xcf, 0x01, 0x0a, 0x10, 0x4d, 0x50, 0x4c, 0x53, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x73, 0x65, 0x71, 0x12, 0x48, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x67, 0x6e, 0x6f,
	0x69, 0x2e, 0x6d, 0x70, 0x6c, 0x73, 0x2e, 0x4d, 0x50, 0x4c, 0x53, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x3a, 0x0a, 0x10, 0x45, 0x63, 0x68, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53,
	0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x54, 0x5f,
	0x53, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55,
	0x54, 0x10, 0x02, 0x32, 0xf5, 0x01, 0x0a, 0x04, 0x4d, 0x50, 0x4c, 0x53, 0x12, 0x45, 0x0a, 0x08,
	0x43, 0x6c, 0x65, 0x61, 0x72, 0x4c, 0x53, 0x50, 0x12, 0x1a, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e,
	0x6d, 0x70, 0x6c, 0x73, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x4c, 0x53, 0x50, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x6d, 0x70, 0x6c, 0x73,
	0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x4c, 0x53, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x10, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x4c, 0x53, 0x50, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x12, 0x22, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x6d,
	0x70, 0x6c, 0x73, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x4c, 0x53, 0x50, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x67, 0x6e,
	0x6f, 0x69, 0x2e, 0x6d, 0x70, 0x6c, 0x73, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x4c, 0x53, 0x50,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x47, 0x0a, 0x08, 0x4d, 0x50, 0x4c, 0x53, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x1a,
	0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x6d, 0x70, 0x6c, 0x73, 0x2e, 0x4d, 0x50, 0x4c, 0x53, 0x50,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x6e, 0x6f,
	0x69, 0x2e, 0x6d, 0x70, 0x6c, 0x73, 0x2e, 0x4d, 0x50, 0x4c, 0x53, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x29, 0x5a, 0x1f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x67, 0x6e, 0x6f, 0x69, 0x2f, 0x6d, 0x70, 0x6c, 0x73, 0xd2, 0x3e,
	0x05, 0x30, 0x2e, 0x31, 0x2e, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mpls_mpls_proto_rawDescOnce sync.Once
	file_mpls_mpls_proto_rawDescData = file_mpls_mpls_proto_rawDesc
)

func file_mpls_mpls_proto_rawDescGZIP() []byte {
	file_mpls_mpls_proto_rawDescOnce.Do(func() {
		file_mpls_mpls_proto_rawDescData = protoimpl.X.CompressGZIP(file_mpls_mpls_proto_rawDescData)
	})
	return file_mpls_mpls_proto_rawDescData
}

var file_mpls_mpls_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_mpls_mpls_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_mpls_mpls_proto_goTypes = []interface{}{
	(ClearLSPRequest_Mode)(0),              // 0: gnoi.mpls.ClearLSPRequest.Mode
	(MPLSPingRequest_ReplyMode)(0),         // 1: gnoi.mpls.MPLSPingRequest.ReplyMode
	(MPLSPingResponse_EchoResponseCode)(0), // 2: gnoi.mpls.MPLSPingResponse.EchoResponseCode
	(*ClearLSPRequest)(nil),                // 3: gnoi.mpls.ClearLSPRequest
	(*ClearLSPResponse)(nil),               // 4: gnoi.mpls.ClearLSPResponse
	(*ClearLSPCountersRequest)(nil),        // 5: gnoi.mpls.ClearLSPCountersRequest
	(*ClearLSPCountersResponse)(nil),       // 6: gnoi.mpls.ClearLSPCountersResponse
	(*MPLSPingPWEDestination)(nil),         // 7: gnoi.mpls.MPLSPingPWEDestination
	(*MPLSPingRSVPTEDestination)(nil),      // 8: gnoi.mpls.MPLSPingRSVPTEDestination
	(*MPLSPingRequest)(nil),                // 9: gnoi.mpls.MPLSPingRequest
	(*MPLSPingResponse)(nil),               // 10: gnoi.mpls.MPLSPingResponse
}
var file_mpls_mpls_proto_depIdxs = []int32{
	0,  // 0: gnoi.mpls.ClearLSPRequest.mode:type_name -> gnoi.mpls.ClearLSPRequest.Mode
	7,  // 1: gnoi.mpls.MPLSPingRequest.fec129_pwe:type_name -> gnoi.mpls.MPLSPingPWEDestination
	8,  // 2: gnoi.mpls.MPLSPingRequest.rsvpte_lsp:type_name -> gnoi.mpls.MPLSPingRSVPTEDestination
	1,  // 3: gnoi.mpls.MPLSPingRequest.reply_mode:type_name -> gnoi.mpls.MPLSPingRequest.ReplyMode
	2,  // 4: gnoi.mpls.MPLSPingResponse.response:type_name -> gnoi.mpls.MPLSPingResponse.EchoResponseCode
	3,  // 5: gnoi.mpls.MPLS.ClearLSP:input_type -> gnoi.mpls.ClearLSPRequest
	5,  // 6: gnoi.mpls.MPLS.ClearLSPCounters:input_type -> gnoi.mpls.ClearLSPCountersRequest
	9,  // 7: gnoi.mpls.MPLS.MPLSPing:input_type -> gnoi.mpls.MPLSPingRequest
	4,  // 8: gnoi.mpls.MPLS.ClearLSP:output_type -> gnoi.mpls.ClearLSPResponse
	6,  // 9: gnoi.mpls.MPLS.ClearLSPCounters:output_type -> gnoi.mpls.ClearLSPCountersResponse
	10, // 10: gnoi.mpls.MPLS.MPLSPing:output_type -> gnoi.mpls.MPLSPingResponse
	8,  // [8:11] is the sub-list for method output_type
	5,  // [5:8] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_mpls_mpls_proto_init() }
func file_mpls_mpls_proto_init() {
	if File_mpls_mpls_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mpls_mpls_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearLSPRequest); i {
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
		file_mpls_mpls_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearLSPResponse); i {
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
		file_mpls_mpls_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearLSPCountersRequest); i {
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
		file_mpls_mpls_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearLSPCountersResponse); i {
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
		file_mpls_mpls_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MPLSPingPWEDestination); i {
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
		file_mpls_mpls_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MPLSPingRSVPTEDestination); i {
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
		file_mpls_mpls_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MPLSPingRequest); i {
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
		file_mpls_mpls_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MPLSPingResponse); i {
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
	file_mpls_mpls_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*MPLSPingRequest_LdpFec)(nil),
		(*MPLSPingRequest_Fec129Pwe)(nil),
		(*MPLSPingRequest_RsvpteLspName)(nil),
		(*MPLSPingRequest_RsvpteLsp)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mpls_mpls_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mpls_mpls_proto_goTypes,
		DependencyIndexes: file_mpls_mpls_proto_depIdxs,
		EnumInfos:         file_mpls_mpls_proto_enumTypes,
		MessageInfos:      file_mpls_mpls_proto_msgTypes,
	}.Build()
	File_mpls_mpls_proto = out.File
	file_mpls_mpls_proto_rawDesc = nil
	file_mpls_mpls_proto_goTypes = nil
	file_mpls_mpls_proto_depIdxs = nil
}
