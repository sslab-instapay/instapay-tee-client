// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client.proto

package client

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AgreeRequestsMessage struct {
	PaymentNumber        int64            `protobuf:"varint,1,opt,name=paymentNumber,proto3" json:"paymentNumber,omitempty"`
	ChannelPayments      *ChannelPayments `protobuf:"bytes,2,opt,name=channelPayments,proto3" json:"channelPayments,omitempty"`
	OriginalMessage      string           `protobuf:"bytes,3,opt,name=originalMessage,proto3" json:"originalMessage,omitempty"`
	Signature            string           `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *AgreeRequestsMessage) Reset()         { *m = AgreeRequestsMessage{} }
func (m *AgreeRequestsMessage) String() string { return proto.CompactTextString(m) }
func (*AgreeRequestsMessage) ProtoMessage()    {}
func (*AgreeRequestsMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{0}
}

func (m *AgreeRequestsMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgreeRequestsMessage.Unmarshal(m, b)
}
func (m *AgreeRequestsMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgreeRequestsMessage.Marshal(b, m, deterministic)
}
func (m *AgreeRequestsMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgreeRequestsMessage.Merge(m, src)
}
func (m *AgreeRequestsMessage) XXX_Size() int {
	return xxx_messageInfo_AgreeRequestsMessage.Size(m)
}
func (m *AgreeRequestsMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AgreeRequestsMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AgreeRequestsMessage proto.InternalMessageInfo

func (m *AgreeRequestsMessage) GetPaymentNumber() int64 {
	if m != nil {
		return m.PaymentNumber
	}
	return 0
}

func (m *AgreeRequestsMessage) GetChannelPayments() *ChannelPayments {
	if m != nil {
		return m.ChannelPayments
	}
	return nil
}

func (m *AgreeRequestsMessage) GetOriginalMessage() string {
	if m != nil {
		return m.OriginalMessage
	}
	return ""
}

func (m *AgreeRequestsMessage) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type UpdateRequestsMessage struct {
	PaymentNumber        int64            `protobuf:"varint,1,opt,name=paymentNumber,proto3" json:"paymentNumber,omitempty"`
	ChannelPayments      *ChannelPayments `protobuf:"bytes,2,opt,name=channelPayments,proto3" json:"channelPayments,omitempty"`
	OriginalMessage      string           `protobuf:"bytes,3,opt,name=originalMessage,proto3" json:"originalMessage,omitempty"`
	Signature            string           `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UpdateRequestsMessage) Reset()         { *m = UpdateRequestsMessage{} }
func (m *UpdateRequestsMessage) String() string { return proto.CompactTextString(m) }
func (*UpdateRequestsMessage) ProtoMessage()    {}
func (*UpdateRequestsMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{1}
}

func (m *UpdateRequestsMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequestsMessage.Unmarshal(m, b)
}
func (m *UpdateRequestsMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequestsMessage.Marshal(b, m, deterministic)
}
func (m *UpdateRequestsMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequestsMessage.Merge(m, src)
}
func (m *UpdateRequestsMessage) XXX_Size() int {
	return xxx_messageInfo_UpdateRequestsMessage.Size(m)
}
func (m *UpdateRequestsMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequestsMessage.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequestsMessage proto.InternalMessageInfo

func (m *UpdateRequestsMessage) GetPaymentNumber() int64 {
	if m != nil {
		return m.PaymentNumber
	}
	return 0
}

func (m *UpdateRequestsMessage) GetChannelPayments() *ChannelPayments {
	if m != nil {
		return m.ChannelPayments
	}
	return nil
}

func (m *UpdateRequestsMessage) GetOriginalMessage() string {
	if m != nil {
		return m.OriginalMessage
	}
	return ""
}

func (m *UpdateRequestsMessage) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type ConfirmRequestsMessage struct {
	PaymentNumber        int64    `protobuf:"varint,1,opt,name=paymentNumber,proto3" json:"paymentNumber,omitempty"`
	OriginalMessage      string   `protobuf:"bytes,2,opt,name=originalMessage,proto3" json:"originalMessage,omitempty"`
	Signature            string   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfirmRequestsMessage) Reset()         { *m = ConfirmRequestsMessage{} }
func (m *ConfirmRequestsMessage) String() string { return proto.CompactTextString(m) }
func (*ConfirmRequestsMessage) ProtoMessage()    {}
func (*ConfirmRequestsMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{2}
}

func (m *ConfirmRequestsMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfirmRequestsMessage.Unmarshal(m, b)
}
func (m *ConfirmRequestsMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfirmRequestsMessage.Marshal(b, m, deterministic)
}
func (m *ConfirmRequestsMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfirmRequestsMessage.Merge(m, src)
}
func (m *ConfirmRequestsMessage) XXX_Size() int {
	return xxx_messageInfo_ConfirmRequestsMessage.Size(m)
}
func (m *ConfirmRequestsMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfirmRequestsMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ConfirmRequestsMessage proto.InternalMessageInfo

func (m *ConfirmRequestsMessage) GetPaymentNumber() int64 {
	if m != nil {
		return m.PaymentNumber
	}
	return 0
}

func (m *ConfirmRequestsMessage) GetOriginalMessage() string {
	if m != nil {
		return m.OriginalMessage
	}
	return ""
}

func (m *ConfirmRequestsMessage) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type ChannelPayment struct {
	ChannelId            int64    `protobuf:"varint,1,opt,name=channelId,proto3" json:"channelId,omitempty"`
	Amount               int64    `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	OriginalMessage      string   `protobuf:"bytes,3,opt,name=originalMessage,proto3" json:"originalMessage,omitempty"`
	Signature            string   `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChannelPayment) Reset()         { *m = ChannelPayment{} }
func (m *ChannelPayment) String() string { return proto.CompactTextString(m) }
func (*ChannelPayment) ProtoMessage()    {}
func (*ChannelPayment) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{3}
}

func (m *ChannelPayment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelPayment.Unmarshal(m, b)
}
func (m *ChannelPayment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelPayment.Marshal(b, m, deterministic)
}
func (m *ChannelPayment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelPayment.Merge(m, src)
}
func (m *ChannelPayment) XXX_Size() int {
	return xxx_messageInfo_ChannelPayment.Size(m)
}
func (m *ChannelPayment) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelPayment.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelPayment proto.InternalMessageInfo

func (m *ChannelPayment) GetChannelId() int64 {
	if m != nil {
		return m.ChannelId
	}
	return 0
}

func (m *ChannelPayment) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ChannelPayment) GetOriginalMessage() string {
	if m != nil {
		return m.OriginalMessage
	}
	return ""
}

func (m *ChannelPayment) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type ChannelPayments struct {
	ChannelPayments      []*ChannelPayment `protobuf:"bytes,1,rep,name=channelPayments,proto3" json:"channelPayments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ChannelPayments) Reset()         { *m = ChannelPayments{} }
func (m *ChannelPayments) String() string { return proto.CompactTextString(m) }
func (*ChannelPayments) ProtoMessage()    {}
func (*ChannelPayments) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{4}
}

func (m *ChannelPayments) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelPayments.Unmarshal(m, b)
}
func (m *ChannelPayments) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelPayments.Marshal(b, m, deterministic)
}
func (m *ChannelPayments) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelPayments.Merge(m, src)
}
func (m *ChannelPayments) XXX_Size() int {
	return xxx_messageInfo_ChannelPayments.Size(m)
}
func (m *ChannelPayments) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelPayments.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelPayments proto.InternalMessageInfo

func (m *ChannelPayments) GetChannelPayments() []*ChannelPayment {
	if m != nil {
		return m.ChannelPayments
	}
	return nil
}

type AgreementResult struct {
	PaymentNumber        int64    `protobuf:"varint,1,opt,name=paymentNumber,proto3" json:"paymentNumber,omitempty"`
	Result               bool     `protobuf:"varint,2,opt,name=result,proto3" json:"result,omitempty"`
	OriginalMessage      string   `protobuf:"bytes,3,opt,name=originalMessage,proto3" json:"originalMessage,omitempty"`
	Signature            string   `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgreementResult) Reset()         { *m = AgreementResult{} }
func (m *AgreementResult) String() string { return proto.CompactTextString(m) }
func (*AgreementResult) ProtoMessage()    {}
func (*AgreementResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{5}
}

func (m *AgreementResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgreementResult.Unmarshal(m, b)
}
func (m *AgreementResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgreementResult.Marshal(b, m, deterministic)
}
func (m *AgreementResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgreementResult.Merge(m, src)
}
func (m *AgreementResult) XXX_Size() int {
	return xxx_messageInfo_AgreementResult.Size(m)
}
func (m *AgreementResult) XXX_DiscardUnknown() {
	xxx_messageInfo_AgreementResult.DiscardUnknown(m)
}

var xxx_messageInfo_AgreementResult proto.InternalMessageInfo

func (m *AgreementResult) GetPaymentNumber() int64 {
	if m != nil {
		return m.PaymentNumber
	}
	return 0
}

func (m *AgreementResult) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *AgreementResult) GetOriginalMessage() string {
	if m != nil {
		return m.OriginalMessage
	}
	return ""
}

func (m *AgreementResult) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type UpdateResult struct {
	PaymentNumber        int64    `protobuf:"varint,1,opt,name=paymentNumber,proto3" json:"paymentNumber,omitempty"`
	Result               bool     `protobuf:"varint,2,opt,name=result,proto3" json:"result,omitempty"`
	OriginalMessage      string   `protobuf:"bytes,3,opt,name=originalMessage,proto3" json:"originalMessage,omitempty"`
	Signature            string   `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResult) Reset()         { *m = UpdateResult{} }
func (m *UpdateResult) String() string { return proto.CompactTextString(m) }
func (*UpdateResult) ProtoMessage()    {}
func (*UpdateResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{6}
}

func (m *UpdateResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResult.Unmarshal(m, b)
}
func (m *UpdateResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResult.Marshal(b, m, deterministic)
}
func (m *UpdateResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResult.Merge(m, src)
}
func (m *UpdateResult) XXX_Size() int {
	return xxx_messageInfo_UpdateResult.Size(m)
}
func (m *UpdateResult) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResult.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResult proto.InternalMessageInfo

func (m *UpdateResult) GetPaymentNumber() int64 {
	if m != nil {
		return m.PaymentNumber
	}
	return 0
}

func (m *UpdateResult) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *UpdateResult) GetOriginalMessage() string {
	if m != nil {
		return m.OriginalMessage
	}
	return ""
}

func (m *UpdateResult) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type Result struct {
	PaymentNumber        int64    `protobuf:"varint,1,opt,name=paymentNumber,proto3" json:"paymentNumber,omitempty"`
	Result               bool     `protobuf:"varint,2,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{7}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetPaymentNumber() int64 {
	if m != nil {
		return m.PaymentNumber
	}
	return 0
}

func (m *Result) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type DirectPaymentResult struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	ReplyMessage         string   `protobuf:"bytes,2,opt,name=replyMessage,proto3" json:"replyMessage,omitempty"`
	ReplySignature       string   `protobuf:"bytes,3,opt,name=replySignature,proto3" json:"replySignature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DirectPaymentResult) Reset()         { *m = DirectPaymentResult{} }
func (m *DirectPaymentResult) String() string { return proto.CompactTextString(m) }
func (*DirectPaymentResult) ProtoMessage()    {}
func (*DirectPaymentResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{8}
}

func (m *DirectPaymentResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DirectPaymentResult.Unmarshal(m, b)
}
func (m *DirectPaymentResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DirectPaymentResult.Marshal(b, m, deterministic)
}
func (m *DirectPaymentResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DirectPaymentResult.Merge(m, src)
}
func (m *DirectPaymentResult) XXX_Size() int {
	return xxx_messageInfo_DirectPaymentResult.Size(m)
}
func (m *DirectPaymentResult) XXX_DiscardUnknown() {
	xxx_messageInfo_DirectPaymentResult.DiscardUnknown(m)
}

var xxx_messageInfo_DirectPaymentResult proto.InternalMessageInfo

func (m *DirectPaymentResult) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *DirectPaymentResult) GetReplyMessage() string {
	if m != nil {
		return m.ReplyMessage
	}
	return ""
}

func (m *DirectPaymentResult) GetReplySignature() string {
	if m != nil {
		return m.ReplySignature
	}
	return ""
}

func init() {
	proto.RegisterType((*AgreeRequestsMessage)(nil), "AgreeRequestsMessage")
	proto.RegisterType((*UpdateRequestsMessage)(nil), "UpdateRequestsMessage")
	proto.RegisterType((*ConfirmRequestsMessage)(nil), "ConfirmRequestsMessage")
	proto.RegisterType((*ChannelPayment)(nil), "ChannelPayment")
	proto.RegisterType((*ChannelPayments)(nil), "ChannelPayments")
	proto.RegisterType((*AgreementResult)(nil), "AgreementResult")
	proto.RegisterType((*UpdateResult)(nil), "UpdateResult")
	proto.RegisterType((*Result)(nil), "Result")
	proto.RegisterType((*DirectPaymentResult)(nil), "DirectPaymentResult")
}

func init() { proto.RegisterFile("client.proto", fileDescriptor_014de31d7ac8c57c) }

var fileDescriptor_014de31d7ac8c57c = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x54, 0xcd, 0xca, 0xd3, 0x40,
	0x14, 0xed, 0x34, 0x12, 0xed, 0xed, 0x4f, 0xca, 0xd8, 0xc6, 0x10, 0x5c, 0x84, 0x41, 0x24, 0xab,
	0x59, 0x54, 0x17, 0x2a, 0x88, 0x48, 0x45, 0x10, 0x54, 0x24, 0xe2, 0x03, 0x4c, 0xd3, 0x31, 0x06,
	0x92, 0x49, 0x9c, 0x4c, 0x16, 0x7d, 0x01, 0x97, 0xae, 0x04, 0x5f, 0x47, 0xdf, 0xca, 0xa5, 0x38,
	0x4d, 0x4d, 0x33, 0xcd, 0xa2, 0x7c, 0x74, 0xf1, 0x7d, 0xcb, 0x39, 0xb9, 0x93, 0x39, 0xe7, 0x9e,
	0x73, 0x2f, 0x4c, 0xe2, 0x2c, 0xe5, 0x42, 0xd1, 0x52, 0x16, 0xaa, 0x20, 0xbf, 0x10, 0x2c, 0x5e,
	0x26, 0x92, 0xf3, 0x88, 0x7f, 0xad, 0x79, 0xa5, 0xaa, 0x77, 0xbc, 0xaa, 0x58, 0xc2, 0xf1, 0x03,
	0x98, 0x96, 0x6c, 0x97, 0x73, 0xa1, 0xde, 0xd7, 0xf9, 0x86, 0x4b, 0x0f, 0x05, 0x28, 0xb4, 0xa2,
	0x2e, 0x88, 0x9f, 0x81, 0x13, 0x7f, 0x61, 0x42, 0xf0, 0xec, 0xc3, 0x1e, 0xaf, 0xbc, 0x61, 0x80,
	0xc2, 0xf1, 0x6a, 0x4e, 0xd7, 0x5d, 0x3c, 0x32, 0x0b, 0x71, 0x08, 0x4e, 0x21, 0xd3, 0x24, 0x15,
	0x2c, 0x6b, 0x1e, 0xf5, 0xac, 0x00, 0x85, 0xa3, 0xc8, 0x84, 0xf1, 0x7d, 0x18, 0x55, 0x69, 0x22,
	0x98, 0xaa, 0x25, 0xf7, 0x6e, 0xe9, 0x9a, 0x16, 0x20, 0xbf, 0x11, 0x2c, 0x3f, 0x95, 0x5b, 0xa6,
	0x6e, 0xb0, 0x86, 0x6f, 0x08, 0xdc, 0x75, 0x21, 0x3e, 0xa7, 0x32, 0xbf, 0x9a, 0x88, 0x1e, 0x22,
	0xc3, 0x33, 0x88, 0x58, 0x26, 0x91, 0xef, 0x08, 0x66, 0x5d, 0xd5, 0xff, 0x2e, 0x34, 0xb2, 0xdf,
	0x6c, 0x9b, 0xc7, 0x5b, 0x00, 0xbb, 0x60, 0xb3, 0xbc, 0xa8, 0x85, 0xd2, 0xef, 0x59, 0x51, 0x73,
	0xba, 0x58, 0x67, 0xde, 0x82, 0x63, 0xb8, 0x80, 0x9f, 0x9e, 0x1a, 0x86, 0x02, 0x2b, 0x1c, 0xaf,
	0x1c, 0xc3, 0xb0, 0x13, 0xbf, 0xc8, 0x4f, 0x04, 0x8e, 0x8e, 0xbb, 0xfe, 0xcc, 0xab, 0x3a, 0x53,
	0x67, 0x36, 0xd8, 0x05, 0x5b, 0xea, 0x7a, 0xad, 0xf3, 0x4e, 0xd4, 0x9c, 0x2e, 0xa6, 0xf3, 0x07,
	0x82, 0xc9, 0x21, 0xc5, 0xd7, 0x88, 0xd6, 0x6b, 0xb0, 0x2f, 0xc1, 0x87, 0xec, 0xe0, 0xee, 0xab,
	0x54, 0xf2, 0x58, 0x1d, 0xac, 0xd9, 0xff, 0xb4, 0x2d, 0x47, 0x1d, 0xfa, 0x04, 0x26, 0x92, 0x97,
	0xd9, 0xae, 0x9b, 0xe5, 0x0e, 0x86, 0x1f, 0xc2, 0x4c, 0x9f, 0x3f, 0x1a, 0x69, 0x36, 0xd0, 0xd5,
	0x1f, 0x04, 0xf6, 0x5a, 0xef, 0x3c, 0xfc, 0x1c, 0xe6, 0xac, 0x75, 0x5f, 0xcf, 0x19, 0x5e, 0xd2,
	0xbe, 0xfd, 0xe7, 0xcf, 0xa9, 0x91, 0x13, 0x32, 0xc0, 0x4f, 0x60, 0x5a, 0x1f, 0x2f, 0x1a, 0xec,
	0xd2, 0xde, 0xc5, 0xe3, 0x4f, 0xe9, 0xb1, 0x95, 0x64, 0x80, 0x1f, 0xc3, 0x2c, 0xde, 0x8f, 0xf7,
	0x61, 0xaa, 0xee, 0xd1, 0xfe, 0x79, 0xf7, 0x6f, 0xd3, 0xff, 0xb7, 0x5e, 0xc0, 0x62, 0xab, 0x9b,
	0x66, 0x4c, 0xa4, 0x99, 0x73, 0x7f, 0x41, 0x7b, 0x9a, 0x4b, 0x06, 0x1b, 0x5b, 0x2f, 0xf9, 0x47,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x7c, 0xb0, 0x15, 0x2b, 0xf4, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClientClient is the client API for Client service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientClient interface {
	AgreementRequest(ctx context.Context, in *AgreeRequestsMessage, opts ...grpc.CallOption) (*AgreementResult, error)
	UpdateRequest(ctx context.Context, in *UpdateRequestsMessage, opts ...grpc.CallOption) (*UpdateResult, error)
	ConfirmPayment(ctx context.Context, in *ConfirmRequestsMessage, opts ...grpc.CallOption) (*Result, error)
	DirectChannelPayment(ctx context.Context, in *ChannelPayment, opts ...grpc.CallOption) (*DirectPaymentResult, error)
}

type clientClient struct {
	cc *grpc.ClientConn
}

func NewClientClient(cc *grpc.ClientConn) ClientClient {
	return &clientClient{cc}
}

func (c *clientClient) AgreementRequest(ctx context.Context, in *AgreeRequestsMessage, opts ...grpc.CallOption) (*AgreementResult, error) {
	out := new(AgreementResult)
	err := c.cc.Invoke(ctx, "/Client/agreementRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) UpdateRequest(ctx context.Context, in *UpdateRequestsMessage, opts ...grpc.CallOption) (*UpdateResult, error) {
	out := new(UpdateResult)
	err := c.cc.Invoke(ctx, "/Client/updateRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) ConfirmPayment(ctx context.Context, in *ConfirmRequestsMessage, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/Client/confirmPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) DirectChannelPayment(ctx context.Context, in *ChannelPayment, opts ...grpc.CallOption) (*DirectPaymentResult, error) {
	out := new(DirectPaymentResult)
	err := c.cc.Invoke(ctx, "/Client/directChannelPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServer is the server API for Client service.
type ClientServer interface {
	AgreementRequest(context.Context, *AgreeRequestsMessage) (*AgreementResult, error)
	UpdateRequest(context.Context, *UpdateRequestsMessage) (*UpdateResult, error)
	ConfirmPayment(context.Context, *ConfirmRequestsMessage) (*Result, error)
	DirectChannelPayment(context.Context, *ChannelPayment) (*DirectPaymentResult, error)
}

// UnimplementedClientServer can be embedded to have forward compatible implementations.
type UnimplementedClientServer struct {
}

func (*UnimplementedClientServer) AgreementRequest(ctx context.Context, req *AgreeRequestsMessage) (*AgreementResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AgreementRequest not implemented")
}
func (*UnimplementedClientServer) UpdateRequest(ctx context.Context, req *UpdateRequestsMessage) (*UpdateResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRequest not implemented")
}
func (*UnimplementedClientServer) ConfirmPayment(ctx context.Context, req *ConfirmRequestsMessage) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmPayment not implemented")
}
func (*UnimplementedClientServer) DirectChannelPayment(ctx context.Context, req *ChannelPayment) (*DirectPaymentResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DirectChannelPayment not implemented")
}

func RegisterClientServer(s *grpc.Server, srv ClientServer) {
	s.RegisterService(&_Client_serviceDesc, srv)
}

func _Client_AgreementRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgreeRequestsMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).AgreementRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Client/AgreementRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).AgreementRequest(ctx, req.(*AgreeRequestsMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_UpdateRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequestsMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).UpdateRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Client/UpdateRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).UpdateRequest(ctx, req.(*UpdateRequestsMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_ConfirmPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmRequestsMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).ConfirmPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Client/ConfirmPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).ConfirmPayment(ctx, req.(*ConfirmRequestsMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_DirectChannelPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChannelPayment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).DirectChannelPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Client/DirectChannelPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).DirectChannelPayment(ctx, req.(*ChannelPayment))
	}
	return interceptor(ctx, in, info, handler)
}

var _Client_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Client",
	HandlerType: (*ClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "agreementRequest",
			Handler:    _Client_AgreementRequest_Handler,
		},
		{
			MethodName: "updateRequest",
			Handler:    _Client_UpdateRequest_Handler,
		},
		{
			MethodName: "confirmPayment",
			Handler:    _Client_ConfirmPayment_Handler,
		},
		{
			MethodName: "directChannelPayment",
			Handler:    _Client_DirectChannelPayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client.proto",
}
