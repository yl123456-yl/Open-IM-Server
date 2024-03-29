// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat/chat.proto

package pbChat // import "./chat"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import sdk_ws "Open_IM/pkg/proto/sdk_ws"

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

type MsgDataToMQ struct {
	Token                string          `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	OperationID          string          `protobuf:"bytes,2,opt,name=operationID" json:"operationID,omitempty"`
	MsgData              *sdk_ws.MsgData `protobuf:"bytes,3,opt,name=msgData" json:"msgData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MsgDataToMQ) Reset()         { *m = MsgDataToMQ{} }
func (m *MsgDataToMQ) String() string { return proto.CompactTextString(m) }
func (*MsgDataToMQ) ProtoMessage()    {}
func (*MsgDataToMQ) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_a5e95d84ecbd21a3, []int{0}
}
func (m *MsgDataToMQ) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgDataToMQ.Unmarshal(m, b)
}
func (m *MsgDataToMQ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgDataToMQ.Marshal(b, m, deterministic)
}
func (dst *MsgDataToMQ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDataToMQ.Merge(dst, src)
}
func (m *MsgDataToMQ) XXX_Size() int {
	return xxx_messageInfo_MsgDataToMQ.Size(m)
}
func (m *MsgDataToMQ) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDataToMQ.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDataToMQ proto.InternalMessageInfo

func (m *MsgDataToMQ) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *MsgDataToMQ) GetOperationID() string {
	if m != nil {
		return m.OperationID
	}
	return ""
}

func (m *MsgDataToMQ) GetMsgData() *sdk_ws.MsgData {
	if m != nil {
		return m.MsgData
	}
	return nil
}

type MsgDataToDB struct {
	MsgData              *sdk_ws.MsgData `protobuf:"bytes,1,opt,name=msgData" json:"msgData,omitempty"`
	OperationID          string          `protobuf:"bytes,2,opt,name=operationID" json:"operationID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MsgDataToDB) Reset()         { *m = MsgDataToDB{} }
func (m *MsgDataToDB) String() string { return proto.CompactTextString(m) }
func (*MsgDataToDB) ProtoMessage()    {}
func (*MsgDataToDB) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_a5e95d84ecbd21a3, []int{1}
}
func (m *MsgDataToDB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgDataToDB.Unmarshal(m, b)
}
func (m *MsgDataToDB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgDataToDB.Marshal(b, m, deterministic)
}
func (dst *MsgDataToDB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDataToDB.Merge(dst, src)
}
func (m *MsgDataToDB) XXX_Size() int {
	return xxx_messageInfo_MsgDataToDB.Size(m)
}
func (m *MsgDataToDB) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDataToDB.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDataToDB proto.InternalMessageInfo

func (m *MsgDataToDB) GetMsgData() *sdk_ws.MsgData {
	if m != nil {
		return m.MsgData
	}
	return nil
}

func (m *MsgDataToDB) GetOperationID() string {
	if m != nil {
		return m.OperationID
	}
	return ""
}

type PushMsgDataToMQ struct {
	OperationID          string          `protobuf:"bytes,1,opt,name=OperationID" json:"OperationID,omitempty"`
	MsgData              *sdk_ws.MsgData `protobuf:"bytes,2,opt,name=msgData" json:"msgData,omitempty"`
	PushToUserID         string          `protobuf:"bytes,3,opt,name=pushToUserID" json:"pushToUserID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *PushMsgDataToMQ) Reset()         { *m = PushMsgDataToMQ{} }
func (m *PushMsgDataToMQ) String() string { return proto.CompactTextString(m) }
func (*PushMsgDataToMQ) ProtoMessage()    {}
func (*PushMsgDataToMQ) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_a5e95d84ecbd21a3, []int{2}
}
func (m *PushMsgDataToMQ) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushMsgDataToMQ.Unmarshal(m, b)
}
func (m *PushMsgDataToMQ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushMsgDataToMQ.Marshal(b, m, deterministic)
}
func (dst *PushMsgDataToMQ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushMsgDataToMQ.Merge(dst, src)
}
func (m *PushMsgDataToMQ) XXX_Size() int {
	return xxx_messageInfo_PushMsgDataToMQ.Size(m)
}
func (m *PushMsgDataToMQ) XXX_DiscardUnknown() {
	xxx_messageInfo_PushMsgDataToMQ.DiscardUnknown(m)
}

var xxx_messageInfo_PushMsgDataToMQ proto.InternalMessageInfo

func (m *PushMsgDataToMQ) GetOperationID() string {
	if m != nil {
		return m.OperationID
	}
	return ""
}

func (m *PushMsgDataToMQ) GetMsgData() *sdk_ws.MsgData {
	if m != nil {
		return m.MsgData
	}
	return nil
}

func (m *PushMsgDataToMQ) GetPushToUserID() string {
	if m != nil {
		return m.PushToUserID
	}
	return ""
}

// message PullMessageReq {
//  string UserID = 1;
//  int64 SeqBegin = 2;
//  int64 SeqEnd = 3;
//  string OperationID = 4;
// }
//
// message PullMessageResp {
//  int32 ErrCode = 1;
//  string ErrMsg = 2;
//  int64 MaxSeq = 3;
//  int64 MinSeq = 4;
//  repeated GatherFormat SingleUserMsg = 5;
//  repeated GatherFormat GroupUserMsg = 6;
// }
// message PullMessageBySeqListReq{
//  string UserID = 1;
//  string OperationID = 2;
//  repeated int64 seqList =3;
// }
type GetMaxAndMinSeqReq struct {
	UserID               string   `protobuf:"bytes,1,opt,name=UserID" json:"UserID,omitempty"`
	OperationID          string   `protobuf:"bytes,2,opt,name=OperationID" json:"OperationID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMaxAndMinSeqReq) Reset()         { *m = GetMaxAndMinSeqReq{} }
func (m *GetMaxAndMinSeqReq) String() string { return proto.CompactTextString(m) }
func (*GetMaxAndMinSeqReq) ProtoMessage()    {}
func (*GetMaxAndMinSeqReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_a5e95d84ecbd21a3, []int{3}
}
func (m *GetMaxAndMinSeqReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMaxAndMinSeqReq.Unmarshal(m, b)
}
func (m *GetMaxAndMinSeqReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMaxAndMinSeqReq.Marshal(b, m, deterministic)
}
func (dst *GetMaxAndMinSeqReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMaxAndMinSeqReq.Merge(dst, src)
}
func (m *GetMaxAndMinSeqReq) XXX_Size() int {
	return xxx_messageInfo_GetMaxAndMinSeqReq.Size(m)
}
func (m *GetMaxAndMinSeqReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMaxAndMinSeqReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetMaxAndMinSeqReq proto.InternalMessageInfo

func (m *GetMaxAndMinSeqReq) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *GetMaxAndMinSeqReq) GetOperationID() string {
	if m != nil {
		return m.OperationID
	}
	return ""
}

type GetMaxAndMinSeqResp struct {
	ErrCode              int32    `protobuf:"varint,1,opt,name=ErrCode" json:"ErrCode,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=ErrMsg" json:"ErrMsg,omitempty"`
	MaxSeq               uint32   `protobuf:"varint,3,opt,name=MaxSeq" json:"MaxSeq,omitempty"`
	MinSeq               uint32   `protobuf:"varint,4,opt,name=MinSeq" json:"MinSeq,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMaxAndMinSeqResp) Reset()         { *m = GetMaxAndMinSeqResp{} }
func (m *GetMaxAndMinSeqResp) String() string { return proto.CompactTextString(m) }
func (*GetMaxAndMinSeqResp) ProtoMessage()    {}
func (*GetMaxAndMinSeqResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_a5e95d84ecbd21a3, []int{4}
}
func (m *GetMaxAndMinSeqResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMaxAndMinSeqResp.Unmarshal(m, b)
}
func (m *GetMaxAndMinSeqResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMaxAndMinSeqResp.Marshal(b, m, deterministic)
}
func (dst *GetMaxAndMinSeqResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMaxAndMinSeqResp.Merge(dst, src)
}
func (m *GetMaxAndMinSeqResp) XXX_Size() int {
	return xxx_messageInfo_GetMaxAndMinSeqResp.Size(m)
}
func (m *GetMaxAndMinSeqResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMaxAndMinSeqResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetMaxAndMinSeqResp proto.InternalMessageInfo

func (m *GetMaxAndMinSeqResp) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *GetMaxAndMinSeqResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *GetMaxAndMinSeqResp) GetMaxSeq() uint32 {
	if m != nil {
		return m.MaxSeq
	}
	return 0
}

func (m *GetMaxAndMinSeqResp) GetMinSeq() uint32 {
	if m != nil {
		return m.MinSeq
	}
	return 0
}

type SendMsgReq struct {
	Token                string          `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	OperationID          string          `protobuf:"bytes,2,opt,name=operationID" json:"operationID,omitempty"`
	MsgData              *sdk_ws.MsgData `protobuf:"bytes,3,opt,name=msgData" json:"msgData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SendMsgReq) Reset()         { *m = SendMsgReq{} }
func (m *SendMsgReq) String() string { return proto.CompactTextString(m) }
func (*SendMsgReq) ProtoMessage()    {}
func (*SendMsgReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_a5e95d84ecbd21a3, []int{5}
}
func (m *SendMsgReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMsgReq.Unmarshal(m, b)
}
func (m *SendMsgReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMsgReq.Marshal(b, m, deterministic)
}
func (dst *SendMsgReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMsgReq.Merge(dst, src)
}
func (m *SendMsgReq) XXX_Size() int {
	return xxx_messageInfo_SendMsgReq.Size(m)
}
func (m *SendMsgReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMsgReq.DiscardUnknown(m)
}

var xxx_messageInfo_SendMsgReq proto.InternalMessageInfo

func (m *SendMsgReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *SendMsgReq) GetOperationID() string {
	if m != nil {
		return m.OperationID
	}
	return ""
}

func (m *SendMsgReq) GetMsgData() *sdk_ws.MsgData {
	if m != nil {
		return m.MsgData
	}
	return nil
}

type SendMsgResp struct {
	ErrCode              int32    `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=errMsg" json:"errMsg,omitempty"`
	ServerMsgID          string   `protobuf:"bytes,4,opt,name=serverMsgID" json:"serverMsgID,omitempty"`
	ClientMsgID          string   `protobuf:"bytes,5,opt,name=clientMsgID" json:"clientMsgID,omitempty"`
	SendTime             int64    `protobuf:"varint,6,opt,name=sendTime" json:"sendTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMsgResp) Reset()         { *m = SendMsgResp{} }
func (m *SendMsgResp) String() string { return proto.CompactTextString(m) }
func (*SendMsgResp) ProtoMessage()    {}
func (*SendMsgResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_a5e95d84ecbd21a3, []int{6}
}
func (m *SendMsgResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMsgResp.Unmarshal(m, b)
}
func (m *SendMsgResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMsgResp.Marshal(b, m, deterministic)
}
func (dst *SendMsgResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMsgResp.Merge(dst, src)
}
func (m *SendMsgResp) XXX_Size() int {
	return xxx_messageInfo_SendMsgResp.Size(m)
}
func (m *SendMsgResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMsgResp.DiscardUnknown(m)
}

var xxx_messageInfo_SendMsgResp proto.InternalMessageInfo

func (m *SendMsgResp) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *SendMsgResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *SendMsgResp) GetServerMsgID() string {
	if m != nil {
		return m.ServerMsgID
	}
	return ""
}

func (m *SendMsgResp) GetClientMsgID() string {
	if m != nil {
		return m.ClientMsgID
	}
	return ""
}

func (m *SendMsgResp) GetSendTime() int64 {
	if m != nil {
		return m.SendTime
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgDataToMQ)(nil), "pbChat.MsgDataToMQ")
	proto.RegisterType((*MsgDataToDB)(nil), "pbChat.MsgDataToDB")
	proto.RegisterType((*PushMsgDataToMQ)(nil), "pbChat.PushMsgDataToMQ")
	proto.RegisterType((*GetMaxAndMinSeqReq)(nil), "pbChat.GetMaxAndMinSeqReq")
	proto.RegisterType((*GetMaxAndMinSeqResp)(nil), "pbChat.GetMaxAndMinSeqResp")
	proto.RegisterType((*SendMsgReq)(nil), "pbChat.SendMsgReq")
	proto.RegisterType((*SendMsgResp)(nil), "pbChat.SendMsgResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Chat service

type ChatClient interface {
	GetMaxAndMinSeq(ctx context.Context, in *GetMaxAndMinSeqReq, opts ...grpc.CallOption) (*GetMaxAndMinSeqResp, error)
	PullMessageBySeqList(ctx context.Context, in *sdk_ws.PullMessageBySeqListReq, opts ...grpc.CallOption) (*sdk_ws.PullMessageBySeqListResp, error)
	SendMsg(ctx context.Context, in *SendMsgReq, opts ...grpc.CallOption) (*SendMsgResp, error)
}

type chatClient struct {
	cc *grpc.ClientConn
}

func NewChatClient(cc *grpc.ClientConn) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) GetMaxAndMinSeq(ctx context.Context, in *GetMaxAndMinSeqReq, opts ...grpc.CallOption) (*GetMaxAndMinSeqResp, error) {
	out := new(GetMaxAndMinSeqResp)
	err := grpc.Invoke(ctx, "/pbChat.Chat/GetMaxAndMinSeq", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) PullMessageBySeqList(ctx context.Context, in *sdk_ws.PullMessageBySeqListReq, opts ...grpc.CallOption) (*sdk_ws.PullMessageBySeqListResp, error) {
	out := new(sdk_ws.PullMessageBySeqListResp)
	err := grpc.Invoke(ctx, "/pbChat.Chat/PullMessageBySeqList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) SendMsg(ctx context.Context, in *SendMsgReq, opts ...grpc.CallOption) (*SendMsgResp, error) {
	out := new(SendMsgResp)
	err := grpc.Invoke(ctx, "/pbChat.Chat/SendMsg", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Chat service

type ChatServer interface {
	GetMaxAndMinSeq(context.Context, *GetMaxAndMinSeqReq) (*GetMaxAndMinSeqResp, error)
	PullMessageBySeqList(context.Context, *sdk_ws.PullMessageBySeqListReq) (*sdk_ws.PullMessageBySeqListResp, error)
	SendMsg(context.Context, *SendMsgReq) (*SendMsgResp, error)
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_GetMaxAndMinSeq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMaxAndMinSeqReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).GetMaxAndMinSeq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbChat.Chat/GetMaxAndMinSeq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).GetMaxAndMinSeq(ctx, req.(*GetMaxAndMinSeqReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_PullMessageBySeqList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(sdk_ws.PullMessageBySeqListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).PullMessageBySeqList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbChat.Chat/PullMessageBySeqList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).PullMessageBySeqList(ctx, req.(*sdk_ws.PullMessageBySeqListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_SendMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).SendMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbChat.Chat/SendMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).SendMsg(ctx, req.(*SendMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbChat.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMaxAndMinSeq",
			Handler:    _Chat_GetMaxAndMinSeq_Handler,
		},
		{
			MethodName: "PullMessageBySeqList",
			Handler:    _Chat_PullMessageBySeqList_Handler,
		},
		{
			MethodName: "SendMsg",
			Handler:    _Chat_SendMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat/chat.proto",
}

func init() { proto.RegisterFile("chat/chat.proto", fileDescriptor_chat_a5e95d84ecbd21a3) }

var fileDescriptor_chat_a5e95d84ecbd21a3 = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0x95, 0xd3, 0x26, 0xf9, 0x65, 0xf2, 0xab, 0x22, 0x6d, 0x2b, 0x64, 0x99, 0x8b, 0xf1, 0x29,
	0x02, 0xc9, 0x96, 0x02, 0x37, 0x4e, 0xa4, 0xae, 0x50, 0x10, 0x4b, 0x8b, 0x13, 0x2e, 0x5c, 0xa2,
	0x6d, 0x3d, 0x72, 0xac, 0x24, 0xf6, 0x66, 0xc7, 0x21, 0x05, 0x3e, 0x03, 0x9f, 0x81, 0x8f, 0xc9,
	0x15, 0x79, 0xd7, 0x69, 0x9d, 0xa6, 0x88, 0x9c, 0xb8, 0x58, 0x7a, 0x6f, 0xc6, 0x6f, 0xde, 0xdb,
	0x7f, 0xd0, 0xbb, 0x99, 0x89, 0x22, 0x28, 0x3f, 0xbe, 0x54, 0x79, 0x91, 0xb3, 0x96, 0xbc, 0x3e,
	0x9f, 0x89, 0xc2, 0x79, 0x76, 0x29, 0x31, 0x9b, 0x8e, 0x78, 0x20, 0xe7, 0x49, 0xa0, 0x4b, 0x01,
	0xc5, 0xf3, 0xe9, 0x86, 0x82, 0x0d, 0x99, 0x56, 0xef, 0x3b, 0x74, 0x39, 0x25, 0xa1, 0x28, 0xc4,
	0x24, 0xe7, 0x1f, 0xd9, 0x19, 0x34, 0x8b, 0x7c, 0x8e, 0x99, 0x6d, 0xb9, 0x56, 0xbf, 0x13, 0x19,
	0xc0, 0x5c, 0xe8, 0xe6, 0x12, 0x95, 0x28, 0xd2, 0x3c, 0x1b, 0x85, 0x76, 0x43, 0xd7, 0xea, 0x14,
	0x7b, 0x05, 0xed, 0xa5, 0x91, 0xb1, 0x8f, 0x5c, 0xab, 0xdf, 0x1d, 0x38, 0x3e, 0xa1, 0xfa, 0x82,
	0x6a, 0x2a, 0x64, 0x3a, 0x95, 0x42, 0x89, 0x25, 0xf9, 0xd5, 0xa0, 0x68, 0xdb, 0xea, 0x61, 0x6d,
	0x78, 0x38, 0xac, 0x8b, 0x58, 0x07, 0x8b, 0xfc, 0xdd, 0x9c, 0xf7, 0xc3, 0x82, 0xde, 0xd5, 0x9a,
	0x66, 0xf5, 0xa0, 0x2e, 0x74, 0x2f, 0x6b, 0x7f, 0x99, 0xb8, 0x75, 0xaa, 0xee, 0xa6, 0x71, 0xb8,
	0x1b, 0x0f, 0xfe, 0x97, 0x6b, 0x9a, 0x4d, 0xf2, 0x4f, 0x84, 0x6a, 0x14, 0xea, 0xd5, 0xe8, 0x44,
	0x3b, 0x9c, 0xf7, 0x01, 0xd8, 0x5b, 0x2c, 0xb8, 0xb8, 0x7d, 0x93, 0xc5, 0x3c, 0xcd, 0xc6, 0xb8,
	0x8a, 0x70, 0xc5, 0x9e, 0x40, 0xab, 0xfa, 0xc7, 0x98, 0xa9, 0xd0, 0x43, 0xa7, 0x8d, 0x3d, 0xa7,
	0xde, 0x06, 0x4e, 0xf7, 0xf4, 0x48, 0x32, 0x1b, 0xda, 0x17, 0x4a, 0x9d, 0xe7, 0x31, 0x6a, 0xc5,
	0x66, 0xb4, 0x85, 0xe5, 0xa8, 0x0b, 0xa5, 0x38, 0x25, 0x95, 0x5a, 0x85, 0x4a, 0x9e, 0x8b, 0xdb,
	0x31, 0xae, 0xb4, 0xed, 0x93, 0xa8, 0x42, 0x9a, 0xd7, 0xba, 0xf6, 0x71, 0xc5, 0x6b, 0xe4, 0x7d,
	0x03, 0x18, 0x63, 0x16, 0x73, 0x4a, 0xca, 0x00, 0xff, 0xf6, 0xec, 0xfc, 0xb4, 0xa0, 0x7b, 0x37,
	0xdc, 0xa4, 0xc5, 0xdd, 0xb4, 0x78, 0x9f, 0x16, 0x77, 0xd2, 0x1a, 0x54, 0x3a, 0x33, 0x73, 0x38,
	0x25, 0xa3, 0x50, 0x47, 0xeb, 0x44, 0x75, 0xaa, 0xec, 0xb8, 0x59, 0xa4, 0x98, 0x15, 0xa6, 0xa3,
	0x69, 0x3a, 0x6a, 0x14, 0x73, 0xe0, 0x3f, 0xc2, 0x2c, 0x9e, 0xa4, 0x4b, 0xb4, 0x5b, 0xae, 0xd5,
	0x3f, 0x8a, 0xee, 0xf0, 0xe0, 0x97, 0x05, 0xc7, 0xe5, 0x35, 0x64, 0xef, 0xa0, 0xf7, 0x60, 0x7f,
	0x98, 0xe3, 0x9b, 0x2b, 0xea, 0xef, 0x1f, 0x04, 0xe7, 0xe9, 0x1f, 0x6b, 0x24, 0x59, 0x0e, 0x67,
	0x57, 0xeb, 0xc5, 0x82, 0x23, 0x91, 0x48, 0x70, 0xf8, 0x75, 0x8c, 0xab, 0xf7, 0x29, 0x15, 0xec,
	0xf9, 0x23, 0x6b, 0xf6, 0x58, 0x63, 0x39, 0xe0, 0xc5, 0xc1, 0xbd, 0x24, 0xd9, 0x00, 0xda, 0xd5,
	0x32, 0x33, 0xb6, 0x35, 0x76, 0xbf, 0xe9, 0xce, 0xe9, 0x1e, 0x47, 0x72, 0xd8, 0xfb, 0x7c, 0xe2,
	0xeb, 0xf7, 0xe8, 0xb5, 0x29, 0x5e, 0xb7, 0xf4, 0x63, 0xf3, 0xf2, 0x77, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xe8, 0xa6, 0x3e, 0x67, 0xaa, 0x04, 0x00, 0x00,
}
