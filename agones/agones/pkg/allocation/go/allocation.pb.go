// Copyright 2020 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was autogenerated. Do not edit directly.
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/allocation/allocation.proto

package allocation

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	_ "google.golang.org/genproto/googleapis/api/annotations"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AllocationRequest_SchedulingStrategy int32

const (
	AllocationRequest_Packed      AllocationRequest_SchedulingStrategy = 0
	AllocationRequest_Distributed AllocationRequest_SchedulingStrategy = 1
)

var AllocationRequest_SchedulingStrategy_name = map[int32]string{
	0: "Packed",
	1: "Distributed",
}
var AllocationRequest_SchedulingStrategy_value = map[string]int32{
	"Packed":      0,
	"Distributed": 1,
}

func (x AllocationRequest_SchedulingStrategy) String() string {
	return proto.EnumName(AllocationRequest_SchedulingStrategy_name, int32(x))
}
func (AllocationRequest_SchedulingStrategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_allocation_bb4a9ec25bdc6979, []int{0, 0}
}

type GameServerSelector_GameServerState int32

const (
	GameServerSelector_READY     GameServerSelector_GameServerState = 0
	GameServerSelector_ALLOCATED GameServerSelector_GameServerState = 1
)

var GameServerSelector_GameServerState_name = map[int32]string{
	0: "READY",
	1: "ALLOCATED",
}
var GameServerSelector_GameServerState_value = map[string]int32{
	"READY":     0,
	"ALLOCATED": 1,
}

func (x GameServerSelector_GameServerState) String() string {
	return proto.EnumName(GameServerSelector_GameServerState_name, int32(x))
}
func (GameServerSelector_GameServerState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_allocation_bb4a9ec25bdc6979, []int{5, 0}
}

type AllocationRequest struct {
	// The k8s namespace that is hosting the targeted fleet of gameservers to be allocated
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// If specified, multi-cluster policies are applied. Otherwise, allocation will happen locally.
	MultiClusterSetting *MultiClusterSetting `protobuf:"bytes,2,opt,name=multiClusterSetting,proto3" json:"multiClusterSetting,omitempty"`
	// The required allocation. Defaults to all GameServers.
	RequiredGameServerSelector *GameServerSelector `protobuf:"bytes,3,opt,name=requiredGameServerSelector,proto3" json:"requiredGameServerSelector,omitempty"`
	// The ordered list of preferred allocations out of the `required` set.
	// If the first selector is not matched, the selection attempts the second selector, and so on.
	PreferredGameServerSelectors []*GameServerSelector `protobuf:"bytes,4,rep,name=preferredGameServerSelectors,proto3" json:"preferredGameServerSelectors,omitempty"`
	// Scheduling strategy. Defaults to "Packed".
	Scheduling AllocationRequest_SchedulingStrategy `protobuf:"varint,5,opt,name=scheduling,proto3,enum=allocation.AllocationRequest_SchedulingStrategy" json:"scheduling,omitempty"`
	// Deprecated: Please use metadata instead. This field is ignored if the
	// metadata field is set
	MetaPatch *MetaPatch `protobuf:"bytes,6,opt,name=metaPatch,proto3" json:"metaPatch,omitempty"`
	// Metadata is optional custom metadata that is added to the game server at
	// allocation. You can use this to tell the server necessary session data
	Metadata             *MetaPatch `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AllocationRequest) Reset()         { *m = AllocationRequest{} }
func (m *AllocationRequest) String() string { return proto.CompactTextString(m) }
func (*AllocationRequest) ProtoMessage()    {}
func (*AllocationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_allocation_bb4a9ec25bdc6979, []int{0}
}
func (m *AllocationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllocationRequest.Unmarshal(m, b)
}
func (m *AllocationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllocationRequest.Marshal(b, m, deterministic)
}
func (dst *AllocationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllocationRequest.Merge(dst, src)
}
func (m *AllocationRequest) XXX_Size() int {
	return xxx_messageInfo_AllocationRequest.Size(m)
}
func (m *AllocationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AllocationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AllocationRequest proto.InternalMessageInfo

func (m *AllocationRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *AllocationRequest) GetMultiClusterSetting() *MultiClusterSetting {
	if m != nil {
		return m.MultiClusterSetting
	}
	return nil
}

func (m *AllocationRequest) GetRequiredGameServerSelector() *GameServerSelector {
	if m != nil {
		return m.RequiredGameServerSelector
	}
	return nil
}

func (m *AllocationRequest) GetPreferredGameServerSelectors() []*GameServerSelector {
	if m != nil {
		return m.PreferredGameServerSelectors
	}
	return nil
}

func (m *AllocationRequest) GetScheduling() AllocationRequest_SchedulingStrategy {
	if m != nil {
		return m.Scheduling
	}
	return AllocationRequest_Packed
}

func (m *AllocationRequest) GetMetaPatch() *MetaPatch {
	if m != nil {
		return m.MetaPatch
	}
	return nil
}

func (m *AllocationRequest) GetMetadata() *MetaPatch {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type AllocationResponse struct {
	GameServerName       string                                     `protobuf:"bytes,2,opt,name=gameServerName,proto3" json:"gameServerName,omitempty"`
	Ports                []*AllocationResponse_GameServerStatusPort `protobuf:"bytes,3,rep,name=ports,proto3" json:"ports,omitempty"`
	Address              string                                     `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	NodeName             string                                     `protobuf:"bytes,5,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *AllocationResponse) Reset()         { *m = AllocationResponse{} }
func (m *AllocationResponse) String() string { return proto.CompactTextString(m) }
func (*AllocationResponse) ProtoMessage()    {}
func (*AllocationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_allocation_bb4a9ec25bdc6979, []int{1}
}
func (m *AllocationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllocationResponse.Unmarshal(m, b)
}
func (m *AllocationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllocationResponse.Marshal(b, m, deterministic)
}
func (dst *AllocationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllocationResponse.Merge(dst, src)
}
func (m *AllocationResponse) XXX_Size() int {
	return xxx_messageInfo_AllocationResponse.Size(m)
}
func (m *AllocationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AllocationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AllocationResponse proto.InternalMessageInfo

func (m *AllocationResponse) GetGameServerName() string {
	if m != nil {
		return m.GameServerName
	}
	return ""
}

func (m *AllocationResponse) GetPorts() []*AllocationResponse_GameServerStatusPort {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *AllocationResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AllocationResponse) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

// The gameserver port info that is allocated.
type AllocationResponse_GameServerStatusPort struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllocationResponse_GameServerStatusPort) Reset() {
	*m = AllocationResponse_GameServerStatusPort{}
}
func (m *AllocationResponse_GameServerStatusPort) String() string { return proto.CompactTextString(m) }
func (*AllocationResponse_GameServerStatusPort) ProtoMessage()    {}
func (*AllocationResponse_GameServerStatusPort) Descriptor() ([]byte, []int) {
	return fileDescriptor_allocation_bb4a9ec25bdc6979, []int{1, 0}
}
func (m *AllocationResponse_GameServerStatusPort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllocationResponse_GameServerStatusPort.Unmarshal(m, b)
}
func (m *AllocationResponse_GameServerStatusPort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllocationResponse_GameServerStatusPort.Marshal(b, m, deterministic)
}
func (dst *AllocationResponse_GameServerStatusPort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllocationResponse_GameServerStatusPort.Merge(dst, src)
}
func (m *AllocationResponse_GameServerStatusPort) XXX_Size() int {
	return xxx_messageInfo_AllocationResponse_GameServerStatusPort.Size(m)
}
func (m *AllocationResponse_GameServerStatusPort) XXX_DiscardUnknown() {
	xxx_messageInfo_AllocationResponse_GameServerStatusPort.DiscardUnknown(m)
}

var xxx_messageInfo_AllocationResponse_GameServerStatusPort proto.InternalMessageInfo

func (m *AllocationResponse_GameServerStatusPort) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AllocationResponse_GameServerStatusPort) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

// Specifies settings for multi-cluster allocation.
type MultiClusterSetting struct {
	// If set to true, multi-cluster allocation is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Selects multi-cluster allocation policies to apply. If not specified, all multi-cluster allocation policies are to be applied.
	PolicySelector       *LabelSelector `protobuf:"bytes,2,opt,name=policySelector,proto3" json:"policySelector,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MultiClusterSetting) Reset()         { *m = MultiClusterSetting{} }
func (m *MultiClusterSetting) String() string { return proto.CompactTextString(m) }
func (*MultiClusterSetting) ProtoMessage()    {}
func (*MultiClusterSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_allocation_bb4a9ec25bdc6979, []int{2}
}
func (m *MultiClusterSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiClusterSetting.Unmarshal(m, b)
}
func (m *MultiClusterSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiClusterSetting.Marshal(b, m, deterministic)
}
func (dst *MultiClusterSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiClusterSetting.Merge(dst, src)
}
func (m *MultiClusterSetting) XXX_Size() int {
	return xxx_messageInfo_MultiClusterSetting.Size(m)
}
func (m *MultiClusterSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiClusterSetting.DiscardUnknown(m)
}

var xxx_messageInfo_MultiClusterSetting proto.InternalMessageInfo

func (m *MultiClusterSetting) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *MultiClusterSetting) GetPolicySelector() *LabelSelector {
	if m != nil {
		return m.PolicySelector
	}
	return nil
}

// MetaPatch is the metadata used to patch the GameServer metadata on allocation
type MetaPatch struct {
	Labels               map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Annotations          map[string]string `protobuf:"bytes,2,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MetaPatch) Reset()         { *m = MetaPatch{} }
func (m *MetaPatch) String() string { return proto.CompactTextString(m) }
func (*MetaPatch) ProtoMessage()    {}
func (*MetaPatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_allocation_bb4a9ec25bdc6979, []int{3}
}
func (m *MetaPatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaPatch.Unmarshal(m, b)
}
func (m *MetaPatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaPatch.Marshal(b, m, deterministic)
}
func (dst *MetaPatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaPatch.Merge(dst, src)
}
func (m *MetaPatch) XXX_Size() int {
	return xxx_messageInfo_MetaPatch.Size(m)
}
func (m *MetaPatch) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaPatch.DiscardUnknown(m)
}

var xxx_messageInfo_MetaPatch proto.InternalMessageInfo

func (m *MetaPatch) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *MetaPatch) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

// LabelSelector used for finding a GameServer with matching labels.
type LabelSelector struct {
	// Labels to match.
	MatchLabels          map[string]string `protobuf:"bytes,1,rep,name=matchLabels,proto3" json:"matchLabels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *LabelSelector) Reset()         { *m = LabelSelector{} }
func (m *LabelSelector) String() string { return proto.CompactTextString(m) }
func (*LabelSelector) ProtoMessage()    {}
func (*LabelSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_allocation_bb4a9ec25bdc6979, []int{4}
}
func (m *LabelSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabelSelector.Unmarshal(m, b)
}
func (m *LabelSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabelSelector.Marshal(b, m, deterministic)
}
func (dst *LabelSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelSelector.Merge(dst, src)
}
func (m *LabelSelector) XXX_Size() int {
	return xxx_messageInfo_LabelSelector.Size(m)
}
func (m *LabelSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelSelector.DiscardUnknown(m)
}

var xxx_messageInfo_LabelSelector proto.InternalMessageInfo

func (m *LabelSelector) GetMatchLabels() map[string]string {
	if m != nil {
		return m.MatchLabels
	}
	return nil
}

// GameServerSelector used for finding a GameServer with matching filters.
type GameServerSelector struct {
	// Labels to match.
	MatchLabels          map[string]string                  `protobuf:"bytes,1,rep,name=matchLabels,proto3" json:"matchLabels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	GameServerState      GameServerSelector_GameServerState `protobuf:"varint,2,opt,name=gameServerState,proto3,enum=allocation.GameServerSelector_GameServerState" json:"gameServerState,omitempty"`
	Players              *PlayerSelector                    `protobuf:"bytes,3,opt,name=players,proto3" json:"players,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *GameServerSelector) Reset()         { *m = GameServerSelector{} }
func (m *GameServerSelector) String() string { return proto.CompactTextString(m) }
func (*GameServerSelector) ProtoMessage()    {}
func (*GameServerSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_allocation_bb4a9ec25bdc6979, []int{5}
}
func (m *GameServerSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameServerSelector.Unmarshal(m, b)
}
func (m *GameServerSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameServerSelector.Marshal(b, m, deterministic)
}
func (dst *GameServerSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameServerSelector.Merge(dst, src)
}
func (m *GameServerSelector) XXX_Size() int {
	return xxx_messageInfo_GameServerSelector.Size(m)
}
func (m *GameServerSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_GameServerSelector.DiscardUnknown(m)
}

var xxx_messageInfo_GameServerSelector proto.InternalMessageInfo

func (m *GameServerSelector) GetMatchLabels() map[string]string {
	if m != nil {
		return m.MatchLabels
	}
	return nil
}

func (m *GameServerSelector) GetGameServerState() GameServerSelector_GameServerState {
	if m != nil {
		return m.GameServerState
	}
	return GameServerSelector_READY
}

func (m *GameServerSelector) GetPlayers() *PlayerSelector {
	if m != nil {
		return m.Players
	}
	return nil
}

// PlayerSelector is filter for player capacity values.
// minAvailable should always be less or equal to maxAvailable.
type PlayerSelector struct {
	MinAvailable         uint64   `protobuf:"varint,1,opt,name=minAvailable,proto3" json:"minAvailable,omitempty"`
	MaxAvailable         uint64   `protobuf:"varint,2,opt,name=maxAvailable,proto3" json:"maxAvailable,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerSelector) Reset()         { *m = PlayerSelector{} }
func (m *PlayerSelector) String() string { return proto.CompactTextString(m) }
func (*PlayerSelector) ProtoMessage()    {}
func (*PlayerSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_allocation_bb4a9ec25bdc6979, []int{6}
}
func (m *PlayerSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerSelector.Unmarshal(m, b)
}
func (m *PlayerSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerSelector.Marshal(b, m, deterministic)
}
func (dst *PlayerSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerSelector.Merge(dst, src)
}
func (m *PlayerSelector) XXX_Size() int {
	return xxx_messageInfo_PlayerSelector.Size(m)
}
func (m *PlayerSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerSelector.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerSelector proto.InternalMessageInfo

func (m *PlayerSelector) GetMinAvailable() uint64 {
	if m != nil {
		return m.MinAvailable
	}
	return 0
}

func (m *PlayerSelector) GetMaxAvailable() uint64 {
	if m != nil {
		return m.MaxAvailable
	}
	return 0
}

func init() {
	proto.RegisterType((*AllocationRequest)(nil), "allocation.AllocationRequest")
	proto.RegisterType((*AllocationResponse)(nil), "allocation.AllocationResponse")
	proto.RegisterType((*AllocationResponse_GameServerStatusPort)(nil), "allocation.AllocationResponse.GameServerStatusPort")
	proto.RegisterType((*MultiClusterSetting)(nil), "allocation.MultiClusterSetting")
	proto.RegisterType((*MetaPatch)(nil), "allocation.MetaPatch")
	proto.RegisterMapType((map[string]string)(nil), "allocation.MetaPatch.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "allocation.MetaPatch.LabelsEntry")
	proto.RegisterType((*LabelSelector)(nil), "allocation.LabelSelector")
	proto.RegisterMapType((map[string]string)(nil), "allocation.LabelSelector.MatchLabelsEntry")
	proto.RegisterType((*GameServerSelector)(nil), "allocation.GameServerSelector")
	proto.RegisterMapType((map[string]string)(nil), "allocation.GameServerSelector.MatchLabelsEntry")
	proto.RegisterType((*PlayerSelector)(nil), "allocation.PlayerSelector")
	proto.RegisterEnum("allocation.AllocationRequest_SchedulingStrategy", AllocationRequest_SchedulingStrategy_name, AllocationRequest_SchedulingStrategy_value)
	proto.RegisterEnum("allocation.GameServerSelector_GameServerState", GameServerSelector_GameServerState_name, GameServerSelector_GameServerState_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AllocationServiceClient is the client API for AllocationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AllocationServiceClient interface {
	Allocate(ctx context.Context, in *AllocationRequest, opts ...grpc.CallOption) (*AllocationResponse, error)
}

type allocationServiceClient struct {
	cc *grpc.ClientConn
}

func NewAllocationServiceClient(cc *grpc.ClientConn) AllocationServiceClient {
	return &allocationServiceClient{cc}
}

func (c *allocationServiceClient) Allocate(ctx context.Context, in *AllocationRequest, opts ...grpc.CallOption) (*AllocationResponse, error) {
	out := new(AllocationResponse)
	err := c.cc.Invoke(ctx, "/allocation.AllocationService/Allocate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AllocationServiceServer is the server API for AllocationService service.
type AllocationServiceServer interface {
	Allocate(context.Context, *AllocationRequest) (*AllocationResponse, error)
}

func RegisterAllocationServiceServer(s *grpc.Server, srv AllocationServiceServer) {
	s.RegisterService(&_AllocationService_serviceDesc, srv)
}

func _AllocationService_Allocate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllocationServiceServer).Allocate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/allocation.AllocationService/Allocate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllocationServiceServer).Allocate(ctx, req.(*AllocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AllocationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "allocation.AllocationService",
	HandlerType: (*AllocationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Allocate",
			Handler:    _AllocationService_Allocate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/allocation/allocation.proto",
}

func init() {
	proto.RegisterFile("proto/allocation/allocation.proto", fileDescriptor_allocation_bb4a9ec25bdc6979)
}

var fileDescriptor_allocation_bb4a9ec25bdc6979 = []byte{
	// 760 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x6e, 0xf3, 0x44,
	0x14, 0xfd, 0x9c, 0xbf, 0x26, 0x37, 0x7c, 0x69, 0xb8, 0x6d, 0x25, 0x63, 0x95, 0x92, 0x7a, 0x51,
	0x55, 0x45, 0x4a, 0x68, 0xca, 0x82, 0x76, 0x51, 0x29, 0x6a, 0x2b, 0x40, 0x4a, 0x21, 0x75, 0x58,
	0x94, 0x0d, 0xd2, 0xc4, 0x1e, 0x52, 0xab, 0x8e, 0xed, 0xce, 0x8c, 0x2b, 0xb2, 0x43, 0x6c, 0x59,
	0xb2, 0xe6, 0x31, 0x78, 0x12, 0x5e, 0xa1, 0xaf, 0x81, 0x84, 0x66, 0x9c, 0x38, 0x93, 0xc4, 0x0d,
	0x54, 0xdf, 0x6e, 0xe6, 0xde, 0x73, 0xcf, 0x9c, 0x39, 0x73, 0x7d, 0x0d, 0x87, 0x31, 0x8b, 0x44,
	0xd4, 0x21, 0x41, 0x10, 0xb9, 0x44, 0xf8, 0x51, 0xa8, 0x2d, 0xdb, 0x2a, 0x87, 0xb0, 0x88, 0x58,
	0xfb, 0xe3, 0x28, 0x1a, 0x07, 0xb4, 0x43, 0x62, 0xbf, 0x43, 0xc2, 0x30, 0x12, 0x2a, 0xcc, 0x53,
	0xa4, 0xfd, 0x57, 0x09, 0x3e, 0xee, 0x65, 0x60, 0x87, 0x3e, 0x25, 0x94, 0x0b, 0xdc, 0x87, 0x5a,
	0x48, 0x26, 0x94, 0xc7, 0xc4, 0xa5, 0xa6, 0xd1, 0x32, 0x8e, 0x6b, 0xce, 0x22, 0x80, 0x77, 0xb0,
	0x33, 0x49, 0x02, 0xe1, 0x5f, 0x05, 0x09, 0x17, 0x94, 0x0d, 0xa9, 0x10, 0x7e, 0x38, 0x36, 0x0b,
	0x2d, 0xe3, 0xb8, 0xde, 0xfd, 0xac, 0xad, 0xa9, 0xb9, 0x5d, 0x87, 0x39, 0x79, 0xb5, 0xf8, 0x13,
	0x58, 0x8c, 0x3e, 0x25, 0x3e, 0xa3, 0xde, 0xd7, 0x64, 0x42, 0x87, 0x94, 0x3d, 0xcb, 0x64, 0x40,
	0x5d, 0x11, 0x31, 0xb3, 0xa8, 0x98, 0x0f, 0x74, 0xe6, 0x75, 0x94, 0xb3, 0x81, 0x01, 0x47, 0xb0,
	0x1f, 0x33, 0xfa, 0x33, 0x65, 0xb9, 0x69, 0x6e, 0x96, 0x5a, 0xc5, 0xff, 0x71, 0xc2, 0x46, 0x0e,
	0x1c, 0x00, 0x70, 0xf7, 0x81, 0x7a, 0x49, 0x20, 0xdd, 0x28, 0xb7, 0x8c, 0xe3, 0x46, 0xf7, 0x0b,
	0x9d, 0x71, 0xcd, 0xe7, 0xf6, 0x30, 0xc3, 0x0f, 0x05, 0x23, 0x82, 0x8e, 0xa7, 0x8e, 0xc6, 0x81,
	0x67, 0x50, 0x9b, 0x50, 0x41, 0x06, 0x44, 0xb8, 0x0f, 0x66, 0x45, 0x99, 0xb0, 0xb7, 0x64, 0xef,
	0x3c, 0xe9, 0x2c, 0x70, 0x78, 0x0a, 0x55, 0xb9, 0xf1, 0x88, 0x20, 0xe6, 0xd6, 0xa6, 0x9a, 0x0c,
	0x66, 0x9f, 0x02, 0xae, 0x2b, 0x41, 0x80, 0xca, 0x80, 0xb8, 0x8f, 0xd4, 0x6b, 0xbe, 0xc3, 0x6d,
	0xa8, 0x5f, 0xfb, 0x5c, 0x30, 0x7f, 0x94, 0x08, 0xea, 0x35, 0x0d, 0xfb, 0x1f, 0x03, 0x50, 0xbf,
	0x0f, 0x8f, 0xa3, 0x90, 0x53, 0x3c, 0x82, 0xc6, 0x38, 0xb3, 0xe6, 0x3b, 0x32, 0xa1, 0xaa, 0x2b,
	0x6a, 0xce, 0x4a, 0x14, 0xbf, 0x85, 0x72, 0x1c, 0x31, 0xc1, 0xcd, 0xa2, 0x32, 0xfe, 0xec, 0x35,
	0x9b, 0x52, 0x5a, 0xfd, 0x2d, 0x04, 0x11, 0x09, 0x1f, 0x44, 0x4c, 0x38, 0x29, 0x03, 0x9a, 0xb0,
	0x45, 0x3c, 0x8f, 0x51, 0x2e, 0x5f, 0x51, 0x9e, 0x35, 0xdf, 0xa2, 0x05, 0xd5, 0x30, 0xf2, 0xa8,
	0x92, 0x51, 0x56, 0xa9, 0x6c, 0x6f, 0x5d, 0xc2, 0x6e, 0x1e, 0x29, 0x22, 0x94, 0x64, 0xa3, 0xcf,
	0x9a, 0x5e, 0xad, 0x65, 0x4c, 0x1e, 0xa5, 0xae, 0x52, 0x76, 0xd4, 0xda, 0x66, 0xb0, 0x93, 0xd3,
	0xdc, 0x52, 0x0c, 0x0d, 0xc9, 0x28, 0xa0, 0x9e, 0x62, 0xa8, 0x3a, 0xf3, 0x2d, 0xf6, 0xa0, 0x11,
	0x47, 0x81, 0xef, 0x4e, 0xb3, 0xae, 0x4e, 0xbf, 0x97, 0x4f, 0xf4, 0xab, 0xf7, 0xc9, 0x88, 0x06,
	0x59, 0xbb, 0xad, 0x14, 0xd8, 0xbf, 0x17, 0xa0, 0x96, 0x3d, 0x1f, 0x9e, 0x43, 0x25, 0x90, 0x70,
	0x6e, 0x1a, 0xca, 0xc3, 0xc3, 0xdc, 0x57, 0x4e, 0x29, 0xf9, 0x4d, 0x28, 0xd8, 0xd4, 0x99, 0x15,
	0xe0, 0x37, 0x50, 0xd7, 0x26, 0x81, 0x59, 0x50, 0xf5, 0x47, 0xf9, 0xf5, 0xbd, 0x05, 0x30, 0x25,
	0xd1, 0x4b, 0xad, 0x73, 0xa8, 0x6b, 0x07, 0x60, 0x13, 0x8a, 0x8f, 0x74, 0x3a, 0x33, 0x4f, 0x2e,
	0x71, 0x17, 0xca, 0xcf, 0x24, 0x48, 0xe6, 0x7d, 0x90, 0x6e, 0x2e, 0x0a, 0x5f, 0x19, 0xd6, 0x25,
	0x34, 0x57, 0xb9, 0xdf, 0x52, 0x6f, 0xff, 0x69, 0xc0, 0xfb, 0x25, 0xbf, 0xb0, 0x0f, 0xf5, 0x89,
	0xd4, 0xdc, 0xd7, 0x6d, 0x39, 0x79, 0xd5, 0xdf, 0xf6, 0xed, 0x02, 0x3c, 0xbb, 0x9a, 0x56, 0x2e,
	0xf5, 0xad, 0x02, 0xde, 0xa4, 0xef, 0xa5, 0x00, 0x98, 0x33, 0x89, 0xee, 0xf2, 0x44, 0x76, 0x36,
	0x0f, 0x9e, 0xcd, 0x4a, 0xf1, 0x1e, 0xb6, 0xc7, 0x4b, 0xbd, 0x9c, 0xaa, 0x69, 0x74, 0xdb, 0xff,
	0x41, 0xbb, 0xfc, 0x05, 0x50, 0x67, 0x95, 0x06, 0xbf, 0x84, 0xad, 0x38, 0x20, 0x53, 0xca, 0xf8,
	0x6c, 0x06, 0x5b, 0x3a, 0xe3, 0x40, 0xa5, 0xb2, 0x76, 0x9d, 0x43, 0x3f, 0xd8, 0xb9, 0xcf, 0x61,
	0x7b, 0x45, 0x19, 0xd6, 0xa0, 0xec, 0xdc, 0xf4, 0xae, 0x7f, 0x6c, 0xbe, 0xc3, 0xf7, 0x50, 0xeb,
	0xf5, 0xfb, 0xdf, 0x5f, 0xf5, 0x7e, 0xb8, 0xb9, 0x6e, 0x1a, 0xf6, 0x3d, 0x34, 0x96, 0x75, 0xa0,
	0x0d, 0x1f, 0x4d, 0xfc, 0xb0, 0xf7, 0x4c, 0xfc, 0x40, 0x7e, 0x7a, 0xea, 0xcc, 0x92, 0xb3, 0x14,
	0x53, 0x18, 0xf2, 0xcb, 0x02, 0x53, 0x98, 0x61, 0xb4, 0x58, 0xf7, 0x57, 0x43, 0xff, 0x35, 0x4a,
	0x35, 0xbe, 0x4b, 0xf1, 0x11, 0xaa, 0xb3, 0x20, 0xc5, 0x4f, 0x37, 0x4e, 0x77, 0xeb, 0x60, 0xf3,
	0x54, 0xb3, 0x5b, 0xbf, 0xfd, 0xfd, 0xf2, 0x47, 0xc1, 0xb2, 0xf7, 0x3a, 0xd2, 0x77, 0xae, 0xae,
	0xbb, 0xa8, 0xb8, 0x30, 0x4e, 0x46, 0x15, 0xf5, 0x93, 0x3e, 0xfb, 0x37, 0x00, 0x00, 0xff, 0xff,
	0x0d, 0x34, 0x24, 0x46, 0xf3, 0x07, 0x00, 0x00,
}