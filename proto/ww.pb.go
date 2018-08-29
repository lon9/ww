// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/ww.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
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

type Camp int32

const (
	Camp_GOOD Camp = 0
	Camp_EVIL Camp = 1
)

var Camp_name = map[int32]string{
	0: "GOOD",
	1: "EVIL",
}
var Camp_value = map[string]int32{
	"GOOD": 0,
	"EVIL": 1,
}

func (x Camp) String() string {
	return proto.EnumName(Camp_name, int32(x))
}
func (Camp) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ww_bf24c6214e2dda08, []int{0}
}

type Kind int32

const (
	Kind_CITIZEN  Kind = 0
	Kind_WAREWOLF Kind = 1
	Kind_TELLER   Kind = 2
	Kind_KNIGHT   Kind = 3
)

var Kind_name = map[int32]string{
	0: "CITIZEN",
	1: "WAREWOLF",
	2: "TELLER",
	3: "KNIGHT",
}
var Kind_value = map[string]int32{
	"CITIZEN":  0,
	"WAREWOLF": 1,
	"TELLER":   2,
	"KNIGHT":   3,
}

func (x Kind) String() string {
	return proto.EnumName(Kind_name, int32(x))
}
func (Kind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ww_bf24c6214e2dda08, []int{1}
}

type State int32

const (
	State_BEFORE  State = 0
	State_MORNING State = 1
	State_NIGHT   State = 2
	State_AFTER   State = 3
)

var State_name = map[int32]string{
	0: "BEFORE",
	1: "MORNING",
	2: "NIGHT",
	3: "AFTER",
}
var State_value = map[string]int32{
	"BEFORE":  0,
	"MORNING": 1,
	"NIGHT":   2,
	"AFTER":   3,
}

func (x State) String() string {
	return proto.EnumName(State_name, int32(x))
}
func (State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ww_bf24c6214e2dda08, []int{2}
}

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ww_bf24c6214e2dda08, []int{0}
}
func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (dst *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(dst, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string   `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Kind                 Kind     `protobuf:"varint,4,opt,name=kind,proto3,enum=ww.Kind" json:"kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ww_bf24c6214e2dda08, []int{1}
}
func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (dst *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(dst, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *HelloResponse) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *HelloResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HelloResponse) GetKind() Kind {
	if m != nil {
		return m.Kind
	}
	return Kind_CITIZEN
}

type StateRequest struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StateRequest) Reset()         { *m = StateRequest{} }
func (m *StateRequest) String() string { return proto.CompactTextString(m) }
func (*StateRequest) ProtoMessage()    {}
func (*StateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ww_bf24c6214e2dda08, []int{2}
}
func (m *StateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateRequest.Unmarshal(m, b)
}
func (m *StateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateRequest.Marshal(b, m, deterministic)
}
func (dst *StateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateRequest.Merge(dst, src)
}
func (m *StateRequest) XXX_Size() int {
	return xxx_messageInfo_StateRequest.Size(m)
}
func (m *StateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StateRequest proto.InternalMessageInfo

func (m *StateRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type StateResponse struct {
	State                State     `protobuf:"varint,1,opt,name=state,proto3,enum=ww.State" json:"state,omitempty"`
	Players              []*Player `protobuf:"bytes,2,rep,name=players,proto3" json:"players,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *StateResponse) Reset()         { *m = StateResponse{} }
func (m *StateResponse) String() string { return proto.CompactTextString(m) }
func (*StateResponse) ProtoMessage()    {}
func (*StateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ww_bf24c6214e2dda08, []int{3}
}
func (m *StateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateResponse.Unmarshal(m, b)
}
func (m *StateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateResponse.Marshal(b, m, deterministic)
}
func (dst *StateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateResponse.Merge(dst, src)
}
func (m *StateResponse) XXX_Size() int {
	return xxx_messageInfo_StateResponse.Size(m)
}
func (m *StateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StateResponse proto.InternalMessageInfo

func (m *StateResponse) GetState() State {
	if m != nil {
		return m.State
	}
	return State_BEFORE
}

func (m *StateResponse) GetPlayers() []*Player {
	if m != nil {
		return m.Players
	}
	return nil
}

type VoteRequest struct {
	SrcUuid              string   `protobuf:"bytes,1,opt,name=src_uuid,json=srcUuid,proto3" json:"src_uuid,omitempty"`
	DstId                int32    `protobuf:"varint,2,opt,name=dst_id,json=dstId,proto3" json:"dst_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VoteRequest) Reset()         { *m = VoteRequest{} }
func (m *VoteRequest) String() string { return proto.CompactTextString(m) }
func (*VoteRequest) ProtoMessage()    {}
func (*VoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ww_bf24c6214e2dda08, []int{4}
}
func (m *VoteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoteRequest.Unmarshal(m, b)
}
func (m *VoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoteRequest.Marshal(b, m, deterministic)
}
func (dst *VoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteRequest.Merge(dst, src)
}
func (m *VoteRequest) XXX_Size() int {
	return xxx_messageInfo_VoteRequest.Size(m)
}
func (m *VoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VoteRequest proto.InternalMessageInfo

func (m *VoteRequest) GetSrcUuid() string {
	if m != nil {
		return m.SrcUuid
	}
	return ""
}

func (m *VoteRequest) GetDstId() int32 {
	if m != nil {
		return m.DstId
	}
	return 0
}

type VoteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VoteResponse) Reset()         { *m = VoteResponse{} }
func (m *VoteResponse) String() string { return proto.CompactTextString(m) }
func (*VoteResponse) ProtoMessage()    {}
func (*VoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ww_bf24c6214e2dda08, []int{5}
}
func (m *VoteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoteResponse.Unmarshal(m, b)
}
func (m *VoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoteResponse.Marshal(b, m, deterministic)
}
func (dst *VoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteResponse.Merge(dst, src)
}
func (m *VoteResponse) XXX_Size() int {
	return xxx_messageInfo_VoteResponse.Size(m)
}
func (m *VoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VoteResponse proto.InternalMessageInfo

type BiteRequest struct {
	SrcUuid              string   `protobuf:"bytes,1,opt,name=src_uuid,json=srcUuid,proto3" json:"src_uuid,omitempty"`
	DstId                int32    `protobuf:"varint,2,opt,name=dst_id,json=dstId,proto3" json:"dst_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BiteRequest) Reset()         { *m = BiteRequest{} }
func (m *BiteRequest) String() string { return proto.CompactTextString(m) }
func (*BiteRequest) ProtoMessage()    {}
func (*BiteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ww_bf24c6214e2dda08, []int{6}
}
func (m *BiteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BiteRequest.Unmarshal(m, b)
}
func (m *BiteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BiteRequest.Marshal(b, m, deterministic)
}
func (dst *BiteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BiteRequest.Merge(dst, src)
}
func (m *BiteRequest) XXX_Size() int {
	return xxx_messageInfo_BiteRequest.Size(m)
}
func (m *BiteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BiteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BiteRequest proto.InternalMessageInfo

func (m *BiteRequest) GetSrcUuid() string {
	if m != nil {
		return m.SrcUuid
	}
	return ""
}

func (m *BiteRequest) GetDstId() int32 {
	if m != nil {
		return m.DstId
	}
	return 0
}

type BiteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BiteResponse) Reset()         { *m = BiteResponse{} }
func (m *BiteResponse) String() string { return proto.CompactTextString(m) }
func (*BiteResponse) ProtoMessage()    {}
func (*BiteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ww_bf24c6214e2dda08, []int{7}
}
func (m *BiteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BiteResponse.Unmarshal(m, b)
}
func (m *BiteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BiteResponse.Marshal(b, m, deterministic)
}
func (dst *BiteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BiteResponse.Merge(dst, src)
}
func (m *BiteResponse) XXX_Size() int {
	return xxx_messageInfo_BiteResponse.Size(m)
}
func (m *BiteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BiteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BiteResponse proto.InternalMessageInfo

type ProtectRequest struct {
	SrcUuid              string   `protobuf:"bytes,1,opt,name=src_uuid,json=srcUuid,proto3" json:"src_uuid,omitempty"`
	DstId                int32    `protobuf:"varint,2,opt,name=dst_id,json=dstId,proto3" json:"dst_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProtectRequest) Reset()         { *m = ProtectRequest{} }
func (m *ProtectRequest) String() string { return proto.CompactTextString(m) }
func (*ProtectRequest) ProtoMessage()    {}
func (*ProtectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ww_bf24c6214e2dda08, []int{8}
}
func (m *ProtectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtectRequest.Unmarshal(m, b)
}
func (m *ProtectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtectRequest.Marshal(b, m, deterministic)
}
func (dst *ProtectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtectRequest.Merge(dst, src)
}
func (m *ProtectRequest) XXX_Size() int {
	return xxx_messageInfo_ProtectRequest.Size(m)
}
func (m *ProtectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProtectRequest proto.InternalMessageInfo

func (m *ProtectRequest) GetSrcUuid() string {
	if m != nil {
		return m.SrcUuid
	}
	return ""
}

func (m *ProtectRequest) GetDstId() int32 {
	if m != nil {
		return m.DstId
	}
	return 0
}

type ProtectResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProtectResponse) Reset()         { *m = ProtectResponse{} }
func (m *ProtectResponse) String() string { return proto.CompactTextString(m) }
func (*ProtectResponse) ProtoMessage()    {}
func (*ProtectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ww_bf24c6214e2dda08, []int{9}
}
func (m *ProtectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtectResponse.Unmarshal(m, b)
}
func (m *ProtectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtectResponse.Marshal(b, m, deterministic)
}
func (dst *ProtectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtectResponse.Merge(dst, src)
}
func (m *ProtectResponse) XXX_Size() int {
	return xxx_messageInfo_ProtectResponse.Size(m)
}
func (m *ProtectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProtectResponse proto.InternalMessageInfo

type TellRequest struct {
	SrcUuid              string   `protobuf:"bytes,1,opt,name=src_uuid,json=srcUuid,proto3" json:"src_uuid,omitempty"`
	DstId                int32    `protobuf:"varint,2,opt,name=dst_id,json=dstId,proto3" json:"dst_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TellRequest) Reset()         { *m = TellRequest{} }
func (m *TellRequest) String() string { return proto.CompactTextString(m) }
func (*TellRequest) ProtoMessage()    {}
func (*TellRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ww_bf24c6214e2dda08, []int{10}
}
func (m *TellRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TellRequest.Unmarshal(m, b)
}
func (m *TellRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TellRequest.Marshal(b, m, deterministic)
}
func (dst *TellRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TellRequest.Merge(dst, src)
}
func (m *TellRequest) XXX_Size() int {
	return xxx_messageInfo_TellRequest.Size(m)
}
func (m *TellRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TellRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TellRequest proto.InternalMessageInfo

func (m *TellRequest) GetSrcUuid() string {
	if m != nil {
		return m.SrcUuid
	}
	return ""
}

func (m *TellRequest) GetDstId() int32 {
	if m != nil {
		return m.DstId
	}
	return 0
}

type TellResponse struct {
	Camp                 Camp     `protobuf:"varint,1,opt,name=camp,proto3,enum=ww.Camp" json:"camp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TellResponse) Reset()         { *m = TellResponse{} }
func (m *TellResponse) String() string { return proto.CompactTextString(m) }
func (*TellResponse) ProtoMessage()    {}
func (*TellResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ww_bf24c6214e2dda08, []int{11}
}
func (m *TellResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TellResponse.Unmarshal(m, b)
}
func (m *TellResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TellResponse.Marshal(b, m, deterministic)
}
func (dst *TellResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TellResponse.Merge(dst, src)
}
func (m *TellResponse) XXX_Size() int {
	return xxx_messageInfo_TellResponse.Size(m)
}
func (m *TellResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TellResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TellResponse proto.InternalMessageInfo

func (m *TellResponse) GetCamp() Camp {
	if m != nil {
		return m.Camp
	}
	return Camp_GOOD
}

type SleepRequest struct {
	SrcUuid              string   `protobuf:"bytes,1,opt,name=src_uuid,json=srcUuid,proto3" json:"src_uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SleepRequest) Reset()         { *m = SleepRequest{} }
func (m *SleepRequest) String() string { return proto.CompactTextString(m) }
func (*SleepRequest) ProtoMessage()    {}
func (*SleepRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ww_bf24c6214e2dda08, []int{12}
}
func (m *SleepRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SleepRequest.Unmarshal(m, b)
}
func (m *SleepRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SleepRequest.Marshal(b, m, deterministic)
}
func (dst *SleepRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SleepRequest.Merge(dst, src)
}
func (m *SleepRequest) XXX_Size() int {
	return xxx_messageInfo_SleepRequest.Size(m)
}
func (m *SleepRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SleepRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SleepRequest proto.InternalMessageInfo

func (m *SleepRequest) GetSrcUuid() string {
	if m != nil {
		return m.SrcUuid
	}
	return ""
}

type SleepResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SleepResponse) Reset()         { *m = SleepResponse{} }
func (m *SleepResponse) String() string { return proto.CompactTextString(m) }
func (*SleepResponse) ProtoMessage()    {}
func (*SleepResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ww_bf24c6214e2dda08, []int{13}
}
func (m *SleepResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SleepResponse.Unmarshal(m, b)
}
func (m *SleepResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SleepResponse.Marshal(b, m, deterministic)
}
func (dst *SleepResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SleepResponse.Merge(dst, src)
}
func (m *SleepResponse) XXX_Size() int {
	return xxx_messageInfo_SleepResponse.Size(m)
}
func (m *SleepResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SleepResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SleepResponse proto.InternalMessageInfo

type Player struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Kind                 Kind     `protobuf:"varint,2,opt,name=kind,proto3,enum=ww.Kind" json:"kind,omitempty"`
	Camp                 Camp     `protobuf:"varint,3,opt,name=camp,proto3,enum=ww.Camp" json:"camp,omitempty"`
	IsDead               bool     `protobuf:"varint,4,opt,name=is_dead,json=isDead,proto3" json:"is_dead,omitempty"`
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_ww_bf24c6214e2dda08, []int{14}
}
func (m *Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player.Unmarshal(m, b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player.Marshal(b, m, deterministic)
}
func (dst *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(dst, src)
}
func (m *Player) XXX_Size() int {
	return xxx_messageInfo_Player.Size(m)
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Player) GetKind() Kind {
	if m != nil {
		return m.Kind
	}
	return Kind_CITIZEN
}

func (m *Player) GetCamp() Camp {
	if m != nil {
		return m.Camp
	}
	return Camp_GOOD
}

func (m *Player) GetIsDead() bool {
	if m != nil {
		return m.IsDead
	}
	return false
}

func (m *Player) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "ww.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "ww.HelloResponse")
	proto.RegisterType((*StateRequest)(nil), "ww.StateRequest")
	proto.RegisterType((*StateResponse)(nil), "ww.StateResponse")
	proto.RegisterType((*VoteRequest)(nil), "ww.VoteRequest")
	proto.RegisterType((*VoteResponse)(nil), "ww.VoteResponse")
	proto.RegisterType((*BiteRequest)(nil), "ww.BiteRequest")
	proto.RegisterType((*BiteResponse)(nil), "ww.BiteResponse")
	proto.RegisterType((*ProtectRequest)(nil), "ww.ProtectRequest")
	proto.RegisterType((*ProtectResponse)(nil), "ww.ProtectResponse")
	proto.RegisterType((*TellRequest)(nil), "ww.TellRequest")
	proto.RegisterType((*TellResponse)(nil), "ww.TellResponse")
	proto.RegisterType((*SleepRequest)(nil), "ww.SleepRequest")
	proto.RegisterType((*SleepResponse)(nil), "ww.SleepResponse")
	proto.RegisterType((*Player)(nil), "ww.Player")
	proto.RegisterEnum("ww.Camp", Camp_name, Camp_value)
	proto.RegisterEnum("ww.Kind", Kind_name, Kind_value)
	proto.RegisterEnum("ww.State", State_name, State_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WWClient is the client API for WW service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WWClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	State(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (WW_StateClient, error)
	Vote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteResponse, error)
	Bite(ctx context.Context, in *BiteRequest, opts ...grpc.CallOption) (*BiteResponse, error)
	Protect(ctx context.Context, in *ProtectRequest, opts ...grpc.CallOption) (*ProtectResponse, error)
	Tell(ctx context.Context, in *TellRequest, opts ...grpc.CallOption) (*TellResponse, error)
	Sleep(ctx context.Context, in *SleepRequest, opts ...grpc.CallOption) (*SleepResponse, error)
}

type wWClient struct {
	cc *grpc.ClientConn
}

func NewWWClient(cc *grpc.ClientConn) WWClient {
	return &wWClient{cc}
}

func (c *wWClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/ww.WW/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wWClient) State(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (WW_StateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_WW_serviceDesc.Streams[0], "/ww.WW/State", opts...)
	if err != nil {
		return nil, err
	}
	x := &wWStateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WW_StateClient interface {
	Recv() (*StateResponse, error)
	grpc.ClientStream
}

type wWStateClient struct {
	grpc.ClientStream
}

func (x *wWStateClient) Recv() (*StateResponse, error) {
	m := new(StateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *wWClient) Vote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteResponse, error) {
	out := new(VoteResponse)
	err := c.cc.Invoke(ctx, "/ww.WW/Vote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wWClient) Bite(ctx context.Context, in *BiteRequest, opts ...grpc.CallOption) (*BiteResponse, error) {
	out := new(BiteResponse)
	err := c.cc.Invoke(ctx, "/ww.WW/Bite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wWClient) Protect(ctx context.Context, in *ProtectRequest, opts ...grpc.CallOption) (*ProtectResponse, error) {
	out := new(ProtectResponse)
	err := c.cc.Invoke(ctx, "/ww.WW/Protect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wWClient) Tell(ctx context.Context, in *TellRequest, opts ...grpc.CallOption) (*TellResponse, error) {
	out := new(TellResponse)
	err := c.cc.Invoke(ctx, "/ww.WW/Tell", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wWClient) Sleep(ctx context.Context, in *SleepRequest, opts ...grpc.CallOption) (*SleepResponse, error) {
	out := new(SleepResponse)
	err := c.cc.Invoke(ctx, "/ww.WW/Sleep", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WWServer is the server API for WW service.
type WWServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	State(*StateRequest, WW_StateServer) error
	Vote(context.Context, *VoteRequest) (*VoteResponse, error)
	Bite(context.Context, *BiteRequest) (*BiteResponse, error)
	Protect(context.Context, *ProtectRequest) (*ProtectResponse, error)
	Tell(context.Context, *TellRequest) (*TellResponse, error)
	Sleep(context.Context, *SleepRequest) (*SleepResponse, error)
}

func RegisterWWServer(s *grpc.Server, srv WWServer) {
	s.RegisterService(&_WW_serviceDesc, srv)
}

func _WW_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WWServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ww.WW/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WWServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WW_State_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StateRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WWServer).State(m, &wWStateServer{stream})
}

type WW_StateServer interface {
	Send(*StateResponse) error
	grpc.ServerStream
}

type wWStateServer struct {
	grpc.ServerStream
}

func (x *wWStateServer) Send(m *StateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _WW_Vote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WWServer).Vote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ww.WW/Vote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WWServer).Vote(ctx, req.(*VoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WW_Bite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BiteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WWServer).Bite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ww.WW/Bite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WWServer).Bite(ctx, req.(*BiteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WW_Protect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProtectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WWServer).Protect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ww.WW/Protect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WWServer).Protect(ctx, req.(*ProtectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WW_Tell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TellRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WWServer).Tell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ww.WW/Tell",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WWServer).Tell(ctx, req.(*TellRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WW_Sleep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SleepRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WWServer).Sleep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ww.WW/Sleep",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WWServer).Sleep(ctx, req.(*SleepRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WW_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ww.WW",
	HandlerType: (*WWServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _WW_Hello_Handler,
		},
		{
			MethodName: "Vote",
			Handler:    _WW_Vote_Handler,
		},
		{
			MethodName: "Bite",
			Handler:    _WW_Bite_Handler,
		},
		{
			MethodName: "Protect",
			Handler:    _WW_Protect_Handler,
		},
		{
			MethodName: "Tell",
			Handler:    _WW_Tell_Handler,
		},
		{
			MethodName: "Sleep",
			Handler:    _WW_Sleep_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "State",
			Handler:       _WW_State_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/ww.proto",
}

func init() { proto.RegisterFile("proto/ww.proto", fileDescriptor_ww_bf24c6214e2dda08) }

var fileDescriptor_ww_bf24c6214e2dda08 = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcf, 0x8b, 0xda, 0x50,
	0x10, 0xc7, 0x4d, 0x4c, 0x8c, 0x8e, 0x1a, 0xdf, 0xbe, 0x52, 0x6a, 0x65, 0xa1, 0x12, 0x7a, 0x50,
	0x11, 0x5b, 0x2c, 0x14, 0x7a, 0x2a, 0xeb, 0x6e, 0x74, 0xc3, 0x5a, 0x5d, 0xde, 0x5a, 0x85, 0xbd,
	0x48, 0x6a, 0xde, 0x21, 0x34, 0x9a, 0xd4, 0x44, 0x42, 0x8f, 0xfd, 0xe3, 0xfa, 0x7f, 0x95, 0xf7,
	0x43, 0x13, 0xe9, 0x42, 0x61, 0x7b, 0x72, 0x98, 0x7c, 0xdf, 0x77, 0x66, 0x98, 0xcf, 0x08, 0x66,
	0xb4, 0x0f, 0x93, 0xf0, 0x5d, 0x9a, 0x0e, 0x78, 0x80, 0xd5, 0x34, 0xb5, 0x2c, 0xa8, 0xdd, 0xd2,
	0x20, 0x08, 0x09, 0xfd, 0x71, 0xa0, 0x71, 0x82, 0x31, 0x68, 0x3b, 0x77, 0x4b, 0x9b, 0x4a, 0x5b,
	0xe9, 0x54, 0x08, 0x8f, 0x2d, 0x0a, 0x75, 0xa9, 0x89, 0xa3, 0x70, 0x17, 0x53, 0x6c, 0x82, 0xea,
	0x7b, 0x5c, 0xa2, 0x13, 0xd5, 0xf7, 0xd8, 0xa3, 0xc3, 0xc1, 0xf7, 0x9a, 0xaa, 0x78, 0xc4, 0xe2,
	0x93, 0x51, 0x31, 0x33, 0xc2, 0x97, 0xa0, 0x7d, 0xf7, 0x77, 0x5e, 0x53, 0x6b, 0x2b, 0x1d, 0x73,
	0x58, 0x1e, 0xa4, 0xe9, 0xe0, 0xce, 0xdf, 0x79, 0x84, 0x67, 0x59, 0x2b, 0x0f, 0x89, 0x9b, 0xd0,
	0x5c, 0x2b, 0xdc, 0x55, 0xc9, 0x5c, 0xad, 0x25, 0xd4, 0xa5, 0x46, 0xb6, 0xf2, 0x06, 0xf4, 0x98,
	0x25, 0xb8, 0xca, 0x1c, 0x56, 0x98, 0xa7, 0x50, 0x88, 0x3c, 0x7e, 0x0b, 0x46, 0x14, 0xb8, 0x3f,
	0xe9, 0x3e, 0x6e, 0xaa, 0xed, 0x62, 0xa7, 0x3a, 0x04, 0x26, 0xb9, 0xe7, 0x29, 0x72, 0xfc, 0x64,
	0x7d, 0x86, 0xea, 0x32, 0xcc, 0x4a, 0xbf, 0x86, 0x72, 0xbc, 0xdf, 0xac, 0x73, 0xe5, 0x8d, 0x78,
	0xbf, 0xf9, 0xca, 0xe6, 0x7a, 0x09, 0x25, 0x2f, 0x4e, 0xd6, 0x72, 0x5a, 0x9d, 0xe8, 0x5e, 0x9c,
	0x38, 0x9e, 0x65, 0x42, 0x4d, 0x18, 0x88, 0xbe, 0x98, 0xe1, 0xc8, 0xff, 0x4f, 0x43, 0x61, 0x20,
	0x0d, 0x47, 0x60, 0xde, 0xef, 0xc3, 0x84, 0x6e, 0x92, 0xe7, 0x7b, 0x5e, 0x40, 0xe3, 0xe4, 0x91,
	0xf5, 0xb9, 0xa0, 0x41, 0xf0, 0x7c, 0xcf, 0x3e, 0xd4, 0x84, 0x81, 0x5c, 0xc8, 0x25, 0x68, 0x1b,
	0x77, 0x1b, 0xc9, 0x7d, 0xf0, 0x1d, 0x5f, 0xbb, 0xdb, 0x88, 0xf0, 0xac, 0xd5, 0x85, 0xda, 0x43,
	0x40, 0x69, 0xf4, 0xef, 0x7a, 0x56, 0x03, 0xea, 0x52, 0x2a, 0x5b, 0xfd, 0xa5, 0x40, 0x49, 0xec,
	0xed, 0x2f, 0x00, 0x8f, 0x60, 0xa9, 0x4f, 0x81, 0x75, 0x6a, 0xa9, 0xf8, 0x54, 0x4b, 0xf8, 0x15,
	0x18, 0x7e, 0xbc, 0xf6, 0xa8, 0x2b, 0xb8, 0x2c, 0x93, 0x92, 0x1f, 0xdf, 0x50, 0x37, 0x23, 0x58,
	0xcf, 0x08, 0xee, 0xb5, 0x40, 0x63, 0x4f, 0x71, 0x19, 0xb4, 0xc9, 0x7c, 0x7e, 0x83, 0x0a, 0x2c,
	0xb2, 0x97, 0xce, 0x14, 0x29, 0xbd, 0x4f, 0xa0, 0xb1, 0xa2, 0xb8, 0x0a, 0xc6, 0xb5, 0xb3, 0x70,
	0x1e, 0xed, 0x19, 0x2a, 0xe0, 0x1a, 0x94, 0x57, 0x57, 0xc4, 0x5e, 0xcd, 0xa7, 0x63, 0xa4, 0x60,
	0x80, 0xd2, 0xc2, 0x9e, 0x4e, 0x6d, 0x82, 0x54, 0x16, 0xdf, 0xcd, 0x9c, 0xc9, 0xed, 0x02, 0x15,
	0x7b, 0x1f, 0x41, 0xe7, 0xd0, 0xb2, 0xe4, 0xc8, 0x1e, 0xcf, 0x89, 0x8d, 0x0a, 0xcc, 0xe7, 0xcb,
	0x9c, 0xcc, 0x9c, 0xd9, 0x04, 0x29, 0xb8, 0x02, 0xba, 0x10, 0xab, 0x2c, 0xbc, 0x1a, 0x2f, 0x6c,
	0x82, 0x8a, 0xc3, 0xdf, 0x2a, 0xa8, 0xab, 0x15, 0xee, 0x83, 0xce, 0x0f, 0x14, 0x23, 0x36, 0x5b,
	0xfe, 0x9e, 0x5b, 0x17, 0xb9, 0x8c, 0xdc, 0xd0, 0xe0, 0x58, 0x0c, 0x65, 0xc7, 0x92, 0x57, 0x9f,
	0x1d, 0xd8, 0x7b, 0x05, 0x77, 0x41, 0x63, 0x68, 0xe3, 0x06, 0xfb, 0x98, 0xbb, 0x92, 0x16, 0xca,
	0x12, 0xd2, 0xba, 0x0b, 0x1a, 0x83, 0x56, 0x48, 0x73, 0xfc, 0x0b, 0x69, 0x9e, 0x67, 0x3c, 0x04,
	0x43, 0xb2, 0x88, 0x31, 0xbf, 0xc8, 0x33, 0xb8, 0x5b, 0x2f, 0xce, 0x72, 0x99, 0x3d, 0x63, 0x4d,
	0xd8, 0xe7, 0xb0, 0x15, 0xf6, 0x67, 0x18, 0xf6, 0x41, 0xe7, 0xf4, 0xc8, 0x21, 0x73, 0xcc, 0xc9,
	0x21, 0xf3, 0x68, 0x8d, 0x8c, 0x47, 0x9d, 0xff, 0x25, 0x7e, 0x2b, 0xf1, 0x9f, 0x0f, 0x7f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xcf, 0x15, 0xca, 0x2e, 0x2b, 0x05, 0x00, 0x00,
}
