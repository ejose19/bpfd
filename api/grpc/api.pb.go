// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package grpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	math "math"
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

type CreateRuleRequest struct {
	Rule                 *Rule    `protobuf:"bytes,2,opt,name=rule,proto3" json:"rule,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRuleRequest) Reset()         { *m = CreateRuleRequest{} }
func (m *CreateRuleRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRuleRequest) ProtoMessage()    {}
func (*CreateRuleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *CreateRuleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRuleRequest.Unmarshal(m, b)
}
func (m *CreateRuleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRuleRequest.Marshal(b, m, deterministic)
}
func (m *CreateRuleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRuleRequest.Merge(m, src)
}
func (m *CreateRuleRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRuleRequest.Size(m)
}
func (m *CreateRuleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRuleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRuleRequest proto.InternalMessageInfo

func (m *CreateRuleRequest) GetRule() *Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

type CreateRuleResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRuleResponse) Reset()         { *m = CreateRuleResponse{} }
func (m *CreateRuleResponse) String() string { return proto.CompactTextString(m) }
func (*CreateRuleResponse) ProtoMessage()    {}
func (*CreateRuleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *CreateRuleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRuleResponse.Unmarshal(m, b)
}
func (m *CreateRuleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRuleResponse.Marshal(b, m, deterministic)
}
func (m *CreateRuleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRuleResponse.Merge(m, src)
}
func (m *CreateRuleResponse) XXX_Size() int {
	return xxx_messageInfo_CreateRuleResponse.Size(m)
}
func (m *CreateRuleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRuleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRuleResponse proto.InternalMessageInfo

type RemoveRuleRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tracer               string   `protobuf:"bytes,2,opt,name=tracer,proto3" json:"tracer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveRuleRequest) Reset()         { *m = RemoveRuleRequest{} }
func (m *RemoveRuleRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveRuleRequest) ProtoMessage()    {}
func (*RemoveRuleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *RemoveRuleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRuleRequest.Unmarshal(m, b)
}
func (m *RemoveRuleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRuleRequest.Marshal(b, m, deterministic)
}
func (m *RemoveRuleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRuleRequest.Merge(m, src)
}
func (m *RemoveRuleRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveRuleRequest.Size(m)
}
func (m *RemoveRuleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRuleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRuleRequest proto.InternalMessageInfo

func (m *RemoveRuleRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RemoveRuleRequest) GetTracer() string {
	if m != nil {
		return m.Tracer
	}
	return ""
}

type RemoveRuleResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveRuleResponse) Reset()         { *m = RemoveRuleResponse{} }
func (m *RemoveRuleResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveRuleResponse) ProtoMessage()    {}
func (*RemoveRuleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *RemoveRuleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRuleResponse.Unmarshal(m, b)
}
func (m *RemoveRuleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRuleResponse.Marshal(b, m, deterministic)
}
func (m *RemoveRuleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRuleResponse.Merge(m, src)
}
func (m *RemoveRuleResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveRuleResponse.Size(m)
}
func (m *RemoveRuleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRuleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRuleResponse proto.InternalMessageInfo

type ListRulesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRulesRequest) Reset()         { *m = ListRulesRequest{} }
func (m *ListRulesRequest) String() string { return proto.CompactTextString(m) }
func (*ListRulesRequest) ProtoMessage()    {}
func (*ListRulesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *ListRulesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRulesRequest.Unmarshal(m, b)
}
func (m *ListRulesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRulesRequest.Marshal(b, m, deterministic)
}
func (m *ListRulesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRulesRequest.Merge(m, src)
}
func (m *ListRulesRequest) XXX_Size() int {
	return xxx_messageInfo_ListRulesRequest.Size(m)
}
func (m *ListRulesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRulesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRulesRequest proto.InternalMessageInfo

type Rule struct {
	Name                 string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tracer               string             `protobuf:"bytes,2,opt,name=tracer,proto3" json:"tracer,omitempty"`
	FilterEvents         map[string]*Filter `protobuf:"bytes,3,rep,name=filterEvents,proto3" json:"filterEvents,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ContainerRuntimes    []string           `protobuf:"bytes,4,rep,name=containerRuntimes,proto3" json:"containerRuntimes,omitempty"`
	Actions              []string           `protobuf:"bytes,5,rep,name=actions,proto3" json:"actions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Rule) Reset()         { *m = Rule{} }
func (m *Rule) String() string { return proto.CompactTextString(m) }
func (*Rule) ProtoMessage()    {}
func (*Rule) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *Rule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rule.Unmarshal(m, b)
}
func (m *Rule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rule.Marshal(b, m, deterministic)
}
func (m *Rule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rule.Merge(m, src)
}
func (m *Rule) XXX_Size() int {
	return xxx_messageInfo_Rule.Size(m)
}
func (m *Rule) XXX_DiscardUnknown() {
	xxx_messageInfo_Rule.DiscardUnknown(m)
}

var xxx_messageInfo_Rule proto.InternalMessageInfo

func (m *Rule) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Rule) GetTracer() string {
	if m != nil {
		return m.Tracer
	}
	return ""
}

func (m *Rule) GetFilterEvents() map[string]*Filter {
	if m != nil {
		return m.FilterEvents
	}
	return nil
}

func (m *Rule) GetContainerRuntimes() []string {
	if m != nil {
		return m.ContainerRuntimes
	}
	return nil
}

func (m *Rule) GetActions() []string {
	if m != nil {
		return m.Actions
	}
	return nil
}

type Filter struct {
	Values               []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{6}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type ListRulesResponse struct {
	Rules                []*Rule  `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRulesResponse) Reset()         { *m = ListRulesResponse{} }
func (m *ListRulesResponse) String() string { return proto.CompactTextString(m) }
func (*ListRulesResponse) ProtoMessage()    {}
func (*ListRulesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{7}
}

func (m *ListRulesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRulesResponse.Unmarshal(m, b)
}
func (m *ListRulesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRulesResponse.Marshal(b, m, deterministic)
}
func (m *ListRulesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRulesResponse.Merge(m, src)
}
func (m *ListRulesResponse) XXX_Size() int {
	return xxx_messageInfo_ListRulesResponse.Size(m)
}
func (m *ListRulesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRulesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListRulesResponse proto.InternalMessageInfo

func (m *ListRulesResponse) GetRules() []*Rule {
	if m != nil {
		return m.Rules
	}
	return nil
}

type Event struct {
	PID                  uint32            `protobuf:"varint,1,opt,name=PID,proto3" json:"PID,omitempty"`
	TGID                 uint32            `protobuf:"varint,2,opt,name=TGID,proto3" json:"TGID,omitempty"`
	UID                  uint32            `protobuf:"varint,3,opt,name=UID,proto3" json:"UID,omitempty"`
	GID                  uint32            `protobuf:"varint,4,opt,name=GID,proto3" json:"GID,omitempty"`
	Data                 map[string]string `protobuf:"bytes,5,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ContainerID          string            `protobuf:"bytes,6,opt,name=containerID,proto3" json:"containerID,omitempty"`
	ContainerRuntime     string            `protobuf:"bytes,7,opt,name=containerRuntime,proto3" json:"containerRuntime,omitempty"`
	Tracer               string            `protobuf:"bytes,8,opt,name=tracer,proto3" json:"tracer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{8}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetPID() uint32 {
	if m != nil {
		return m.PID
	}
	return 0
}

func (m *Event) GetTGID() uint32 {
	if m != nil {
		return m.TGID
	}
	return 0
}

func (m *Event) GetUID() uint32 {
	if m != nil {
		return m.UID
	}
	return 0
}

func (m *Event) GetGID() uint32 {
	if m != nil {
		return m.GID
	}
	return 0
}

func (m *Event) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Event) GetContainerID() string {
	if m != nil {
		return m.ContainerID
	}
	return ""
}

func (m *Event) GetContainerRuntime() string {
	if m != nil {
		return m.ContainerRuntime
	}
	return ""
}

func (m *Event) GetTracer() string {
	if m != nil {
		return m.Tracer
	}
	return ""
}

type LiveTraceRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LiveTraceRequest) Reset()         { *m = LiveTraceRequest{} }
func (m *LiveTraceRequest) String() string { return proto.CompactTextString(m) }
func (*LiveTraceRequest) ProtoMessage()    {}
func (*LiveTraceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{9}
}

func (m *LiveTraceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LiveTraceRequest.Unmarshal(m, b)
}
func (m *LiveTraceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LiveTraceRequest.Marshal(b, m, deterministic)
}
func (m *LiveTraceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiveTraceRequest.Merge(m, src)
}
func (m *LiveTraceRequest) XXX_Size() int {
	return xxx_messageInfo_LiveTraceRequest.Size(m)
}
func (m *LiveTraceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LiveTraceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LiveTraceRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateRuleRequest)(nil), "grpc.CreateRuleRequest")
	proto.RegisterType((*CreateRuleResponse)(nil), "grpc.CreateRuleResponse")
	proto.RegisterType((*RemoveRuleRequest)(nil), "grpc.RemoveRuleRequest")
	proto.RegisterType((*RemoveRuleResponse)(nil), "grpc.RemoveRuleResponse")
	proto.RegisterType((*ListRulesRequest)(nil), "grpc.ListRulesRequest")
	proto.RegisterType((*Rule)(nil), "grpc.Rule")
	proto.RegisterMapType((map[string]*Filter)(nil), "grpc.Rule.FilterEventsEntry")
	proto.RegisterType((*Filter)(nil), "grpc.Filter")
	proto.RegisterType((*ListRulesResponse)(nil), "grpc.ListRulesResponse")
	proto.RegisterType((*Event)(nil), "grpc.Event")
	proto.RegisterMapType((map[string]string)(nil), "grpc.Event.DataEntry")
	proto.RegisterType((*LiveTraceRequest)(nil), "grpc.LiveTraceRequest")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcf, 0x6b, 0xd4, 0x40,
	0x14, 0xde, 0xfc, 0xd8, 0xd4, 0xbc, 0xb4, 0xb0, 0x19, 0x6a, 0x3b, 0x04, 0x91, 0x90, 0xd3, 0x2a,
	0x12, 0x64, 0xab, 0x28, 0x1e, 0xd4, 0x6a, 0xaa, 0x04, 0x14, 0xca, 0xd0, 0x5e, 0xbc, 0x8d, 0xeb,
	0x28, 0xc1, 0xdd, 0x64, 0xcd, 0x4c, 0x16, 0x7a, 0xf4, 0xec, 0x9f, 0xe0, 0x3f, 0x2b, 0xf3, 0x26,
	0xbb, 0x3b, 0xdb, 0xf4, 0xe2, 0x6d, 0xf2, 0xbd, 0xef, 0x7d, 0xef, 0x77, 0x20, 0xe4, 0xab, 0x2a,
	0x5f, 0xb5, 0x8d, 0x6a, 0x88, 0xff, 0xa3, 0x5d, 0xcd, 0xb3, 0x33, 0x88, 0xdf, 0xb7, 0x82, 0x2b,
	0xc1, 0xba, 0x85, 0x60, 0xe2, 0x57, 0x27, 0xa4, 0x22, 0x0f, 0xc1, 0x6f, 0xbb, 0x85, 0xa0, 0x6e,
	0xea, 0x4c, 0xa3, 0x19, 0xe4, 0x9a, 0x99, 0x23, 0x01, 0xf1, 0xec, 0x18, 0x88, 0xed, 0x24, 0x57,
	0x4d, 0x2d, 0x45, 0xf6, 0x06, 0x62, 0x26, 0x96, 0xcd, 0x7a, 0x4f, 0x8a, 0x80, 0x5f, 0xf3, 0xa5,
	0xa0, 0x4e, 0xea, 0x4c, 0x43, 0x86, 0x6f, 0x72, 0x02, 0x81, 0x6a, 0xf9, 0x5c, 0xb4, 0x18, 0x20,
	0x64, 0xfd, 0x97, 0x96, 0xb5, 0x05, 0x7a, 0x59, 0x02, 0x93, 0x4f, 0x95, 0x54, 0x1a, 0x93, 0xbd,
	0x6a, 0xf6, 0xc7, 0x05, 0x5f, 0x03, 0xff, 0x23, 0x4f, 0xde, 0xc2, 0xe1, 0xf7, 0x6a, 0xa1, 0x44,
	0x7b, 0xb1, 0x16, 0xb5, 0x92, 0xd4, 0x4b, 0xbd, 0x69, 0x34, 0x7b, 0xb0, 0xab, 0x2e, 0xff, 0x60,
	0x99, 0x2f, 0x6a, 0xd5, 0xde, 0xb0, 0x3d, 0x0f, 0xf2, 0x04, 0xe2, 0x79, 0x53, 0x2b, 0x5e, 0xd5,
	0xa2, 0x65, 0x5d, 0xad, 0xaa, 0xa5, 0x90, 0xd4, 0x4f, 0xbd, 0x69, 0xc8, 0x86, 0x06, 0x42, 0xe1,
	0x80, 0xcf, 0x55, 0xd5, 0xd4, 0x92, 0x8e, 0x91, 0xb3, 0xf9, 0x4c, 0x3e, 0x43, 0x3c, 0x08, 0x45,
	0x26, 0xe0, 0xfd, 0x14, 0x37, 0x7d, 0x25, 0xfa, 0x49, 0x32, 0x18, 0xaf, 0xf9, 0xa2, 0xdb, 0xcc,
	0xe1, 0xd0, 0x64, 0x6a, 0x3c, 0x99, 0x31, 0xbd, 0x72, 0x5f, 0x3a, 0x59, 0x0a, 0x81, 0x01, 0x75,
	0xe9, 0x08, 0x4b, 0xea, 0x60, 0xc4, 0xfe, 0x2b, 0x7b, 0x0e, 0xb1, 0xd5, 0x43, 0xd3, 0x58, 0x92,
	0xc2, 0x58, 0x4f, 0xd3, 0x70, 0xf7, 0xc7, 0x6c, 0x0c, 0xd9, 0x5f, 0x17, 0xc6, 0x98, 0xa2, 0x4e,
	0xee, 0xb2, 0x2c, 0x30, 0xb9, 0x23, 0xa6, 0x9f, 0xba, 0xf3, 0x57, 0x1f, 0xcb, 0x02, 0x73, 0x3b,
	0x62, 0xf8, 0xd6, 0xac, 0xeb, 0xb2, 0xa0, 0x9e, 0x61, 0x5d, 0x1b, 0x44, 0x93, 0x7c, 0x83, 0x68,
	0xce, 0x23, 0xf0, 0xbf, 0x71, 0xc5, 0xb1, 0x25, 0xd1, 0xec, 0xbe, 0x09, 0x8a, 0x41, 0xf2, 0x82,
	0x2b, 0x6e, 0xda, 0x8e, 0x14, 0x92, 0x42, 0xb4, 0xed, 0x6a, 0x59, 0xd0, 0x00, 0x3b, 0x63, 0x43,
	0xe4, 0x31, 0x4c, 0x6e, 0xf7, 0x9d, 0x1e, 0x20, 0x6d, 0x80, 0x5b, 0x6b, 0x71, 0xcf, 0x5e, 0x8b,
	0xe4, 0x05, 0x84, 0xdb, 0xc0, 0x77, 0x0c, 0xe1, 0xd8, 0x1e, 0x42, 0x68, 0xb7, 0x1d, 0x17, 0x73,
	0x2d, 0xae, 0xb4, 0x4c, 0xbf, 0x98, 0xb3, 0xdf, 0x2e, 0x78, 0xe7, 0x97, 0x25, 0x39, 0x07, 0xd8,
	0x5d, 0x08, 0x39, 0x35, 0x55, 0x0e, 0x0e, 0x2d, 0xa1, 0x43, 0x43, 0xbf, 0xf5, 0x23, 0x2d, 0xb1,
	0xbb, 0x86, 0x8d, 0xc4, 0xe0, 0xc0, 0x36, 0x12, 0x77, 0x1c, 0xce, 0x88, 0xbc, 0x86, 0x70, 0x3b,
	0x76, 0x72, 0x62, 0x88, 0xb7, 0x6f, 0x29, 0x39, 0x1d, 0xe0, 0x5b, 0xff, 0x67, 0xda, 0xbf, 0xaf,
	0x70, 0xe7, 0xbf, 0x5f, 0x72, 0x12, 0x59, 0x23, 0xcc, 0x46, 0x4f, 0x9d, 0x77, 0xc1, 0x17, 0xfc,
	0xb5, 0x7c, 0x0d, 0xf0, 0x3f, 0x73, 0xf6, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xae, 0xa6, 0x2d, 0x2f,
	0x74, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type APIClient interface {
	CreateRule(ctx context.Context, in *CreateRuleRequest, opts ...grpc.CallOption) (*CreateRuleResponse, error)
	RemoveRule(ctx context.Context, in *RemoveRuleRequest, opts ...grpc.CallOption) (*RemoveRuleResponse, error)
	ListRules(ctx context.Context, in *ListRulesRequest, opts ...grpc.CallOption) (*ListRulesResponse, error)
	LiveTrace(ctx context.Context, in *LiveTraceRequest, opts ...grpc.CallOption) (API_LiveTraceClient, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) CreateRule(ctx context.Context, in *CreateRuleRequest, opts ...grpc.CallOption) (*CreateRuleResponse, error) {
	out := new(CreateRuleResponse)
	err := c.cc.Invoke(ctx, "/grpc.API/CreateRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) RemoveRule(ctx context.Context, in *RemoveRuleRequest, opts ...grpc.CallOption) (*RemoveRuleResponse, error) {
	out := new(RemoveRuleResponse)
	err := c.cc.Invoke(ctx, "/grpc.API/RemoveRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListRules(ctx context.Context, in *ListRulesRequest, opts ...grpc.CallOption) (*ListRulesResponse, error) {
	out := new(ListRulesResponse)
	err := c.cc.Invoke(ctx, "/grpc.API/ListRules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) LiveTrace(ctx context.Context, in *LiveTraceRequest, opts ...grpc.CallOption) (API_LiveTraceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_API_serviceDesc.Streams[0], "/grpc.API/LiveTrace", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPILiveTraceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type API_LiveTraceClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type aPILiveTraceClient struct {
	grpc.ClientStream
}

func (x *aPILiveTraceClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// APIServer is the server API for API service.
type APIServer interface {
	CreateRule(context.Context, *CreateRuleRequest) (*CreateRuleResponse, error)
	RemoveRule(context.Context, *RemoveRuleRequest) (*RemoveRuleResponse, error)
	ListRules(context.Context, *ListRulesRequest) (*ListRulesResponse, error)
	LiveTrace(*LiveTraceRequest, API_LiveTraceServer) error
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_CreateRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CreateRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.API/CreateRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CreateRule(ctx, req.(*CreateRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_RemoveRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).RemoveRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.API/RemoveRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).RemoveRule(ctx, req.(*RemoveRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ListRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.API/ListRules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ListRules(ctx, req.(*ListRulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_LiveTrace_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LiveTraceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(APIServer).LiveTrace(m, &aPILiveTraceServer{stream})
}

type API_LiveTraceServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type aPILiveTraceServer struct {
	grpc.ServerStream
}

func (x *aPILiveTraceServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRule",
			Handler:    _API_CreateRule_Handler,
		},
		{
			MethodName: "RemoveRule",
			Handler:    _API_RemoveRule_Handler,
		},
		{
			MethodName: "ListRules",
			Handler:    _API_ListRules_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LiveTrace",
			Handler:       _API_LiveTrace_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}
