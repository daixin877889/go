// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/proto/userAgentChuangke/userAgentChuangke.proto

package userAgentChuangke

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

// 创客代理商请求数据
type ChuangkeAgentRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	UserId               int32    `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
	Nickname             string   `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`
	RealName             string   `protobuf:"bytes,5,opt,name=realName,proto3" json:"realName,omitempty"`
	Mobile               string   `protobuf:"bytes,6,opt,name=mobile,proto3" json:"mobile,omitempty"`
	VipType              int32    `protobuf:"varint,7,opt,name=vipType,proto3" json:"vipType,omitempty"`
	ReNickname           string   `protobuf:"bytes,8,opt,name=reNickname,proto3" json:"reNickname,omitempty"`
	PVipType             int32    `protobuf:"varint,9,opt,name=pVipType,proto3" json:"pVipType,omitempty"`
	PuserId              int32    `protobuf:"varint,10,opt,name=puserId,proto3" json:"puserId,omitempty"`
	ApplyStatus          int32    `protobuf:"varint,11,opt,name=applyStatus,proto3" json:"applyStatus,omitempty"`
	ShowUserId           string   `protobuf:"bytes,12,opt,name=showUserId,proto3" json:"showUserId,omitempty"`
	ShowpUserId          string   `protobuf:"bytes,13,opt,name=showpUserId,proto3" json:"showpUserId,omitempty"`
	Page                 int32    `protobuf:"varint,14,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int32    `protobuf:"varint,15,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChuangkeAgentRequest) Reset()         { *m = ChuangkeAgentRequest{} }
func (m *ChuangkeAgentRequest) String() string { return proto.CompactTextString(m) }
func (*ChuangkeAgentRequest) ProtoMessage()    {}
func (*ChuangkeAgentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_435bad2763cab632, []int{0}
}

func (m *ChuangkeAgentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChuangkeAgentRequest.Unmarshal(m, b)
}
func (m *ChuangkeAgentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChuangkeAgentRequest.Marshal(b, m, deterministic)
}
func (m *ChuangkeAgentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChuangkeAgentRequest.Merge(m, src)
}
func (m *ChuangkeAgentRequest) XXX_Size() int {
	return xxx_messageInfo_ChuangkeAgentRequest.Size(m)
}
func (m *ChuangkeAgentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChuangkeAgentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChuangkeAgentRequest proto.InternalMessageInfo

func (m *ChuangkeAgentRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ChuangkeAgentRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ChuangkeAgentRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *ChuangkeAgentRequest) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *ChuangkeAgentRequest) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

func (m *ChuangkeAgentRequest) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *ChuangkeAgentRequest) GetVipType() int32 {
	if m != nil {
		return m.VipType
	}
	return 0
}

func (m *ChuangkeAgentRequest) GetReNickname() string {
	if m != nil {
		return m.ReNickname
	}
	return ""
}

func (m *ChuangkeAgentRequest) GetPVipType() int32 {
	if m != nil {
		return m.PVipType
	}
	return 0
}

func (m *ChuangkeAgentRequest) GetPuserId() int32 {
	if m != nil {
		return m.PuserId
	}
	return 0
}

func (m *ChuangkeAgentRequest) GetApplyStatus() int32 {
	if m != nil {
		return m.ApplyStatus
	}
	return 0
}

func (m *ChuangkeAgentRequest) GetShowUserId() string {
	if m != nil {
		return m.ShowUserId
	}
	return ""
}

func (m *ChuangkeAgentRequest) GetShowpUserId() string {
	if m != nil {
		return m.ShowpUserId
	}
	return ""
}

func (m *ChuangkeAgentRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ChuangkeAgentRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

// 创客代理商返回数据
type ChuangkeAgentReply struct {
	Id                    int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid                   string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	UserId                int32    `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
	Nickname              string   `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Mobile                string   `protobuf:"bytes,5,opt,name=mobile,proto3" json:"mobile,omitempty"`
	VipType               int32    `protobuf:"varint,6,opt,name=vipType,proto3" json:"vipType,omitempty"`
	VipGivenQuota         int32    `protobuf:"varint,7,opt,name=vipGivenQuota,proto3" json:"vipGivenQuota,omitempty"`
	VipRemainingQuota     int32    `protobuf:"varint,8,opt,name=vipRemainingQuota,proto3" json:"vipRemainingQuota,omitempty"`
	SvipGivenQuota        int32    `protobuf:"varint,9,opt,name=svipGivenQuota,proto3" json:"svipGivenQuota,omitempty"`
	SvipRemainingQuota    int32    `protobuf:"varint,10,opt,name=svipRemainingQuota,proto3" json:"svipRemainingQuota,omitempty"`
	Headimgurl            string   `protobuf:"bytes,11,opt,name=headimgurl,proto3" json:"headimgurl,omitempty"`
	ReNickname            string   `protobuf:"bytes,12,opt,name=reNickname,proto3" json:"reNickname,omitempty"`
	PVipType              int32    `protobuf:"varint,13,opt,name=pVipType,proto3" json:"pVipType,omitempty"`
	PuserId               int32    `protobuf:"varint,14,opt,name=puserId,proto3" json:"puserId,omitempty"`
	AaId                  int32    `protobuf:"varint,15,opt,name=aaId,proto3" json:"aaId,omitempty"`
	RealName              string   `protobuf:"bytes,16,opt,name=realName,proto3" json:"realName,omitempty"`
	ApplyStatus           int32    `protobuf:"varint,17,opt,name=applyStatus,proto3" json:"applyStatus,omitempty"`
	AuditAt               string   `protobuf:"bytes,18,opt,name=auditAt,proto3" json:"auditAt,omitempty"`
	Province              string   `protobuf:"bytes,19,opt,name=province,proto3" json:"province,omitempty"`
	City                  string   `protobuf:"bytes,20,opt,name=city,proto3" json:"city,omitempty"`
	Area                  string   `protobuf:"bytes,21,opt,name=area,proto3" json:"area,omitempty"`
	AgentGivenQuota       int32    `protobuf:"varint,22,opt,name=agentGivenQuota,proto3" json:"agentGivenQuota,omitempty"`
	AgentRemainingQuota   int32    `protobuf:"varint,23,opt,name=agentRemainingQuota,proto3" json:"agentRemainingQuota,omitempty"`
	PartnerGivenQuota     int32    `protobuf:"varint,24,opt,name=partnerGivenQuota,proto3" json:"partnerGivenQuota,omitempty"`
	PartnerRemainingQuota int32    `protobuf:"varint,25,opt,name=partnerRemainingQuota,proto3" json:"partnerRemainingQuota,omitempty"`
	BranchGivenQuota      int32    `protobuf:"varint,26,opt,name=branchGivenQuota,proto3" json:"branchGivenQuota,omitempty"`
	BranchRemainingQuota  int32    `protobuf:"varint,27,opt,name=branchRemainingQuota,proto3" json:"branchRemainingQuota,omitempty"`
	ShowUserId            string   `protobuf:"bytes,28,opt,name=showUserId,proto3" json:"showUserId,omitempty"`
	ShowpUserId           string   `protobuf:"bytes,29,opt,name=showpUserId,proto3" json:"showpUserId,omitempty"`
	CreatedAt             string   `protobuf:"bytes,30,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt             string   `protobuf:"bytes,31,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt             string   `protobuf:"bytes,32,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *ChuangkeAgentReply) Reset()         { *m = ChuangkeAgentReply{} }
func (m *ChuangkeAgentReply) String() string { return proto.CompactTextString(m) }
func (*ChuangkeAgentReply) ProtoMessage()    {}
func (*ChuangkeAgentReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_435bad2763cab632, []int{1}
}

func (m *ChuangkeAgentReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChuangkeAgentReply.Unmarshal(m, b)
}
func (m *ChuangkeAgentReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChuangkeAgentReply.Marshal(b, m, deterministic)
}
func (m *ChuangkeAgentReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChuangkeAgentReply.Merge(m, src)
}
func (m *ChuangkeAgentReply) XXX_Size() int {
	return xxx_messageInfo_ChuangkeAgentReply.Size(m)
}
func (m *ChuangkeAgentReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ChuangkeAgentReply.DiscardUnknown(m)
}

var xxx_messageInfo_ChuangkeAgentReply proto.InternalMessageInfo

func (m *ChuangkeAgentReply) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ChuangkeAgentReply) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ChuangkeAgentReply) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *ChuangkeAgentReply) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *ChuangkeAgentReply) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *ChuangkeAgentReply) GetVipType() int32 {
	if m != nil {
		return m.VipType
	}
	return 0
}

func (m *ChuangkeAgentReply) GetVipGivenQuota() int32 {
	if m != nil {
		return m.VipGivenQuota
	}
	return 0
}

func (m *ChuangkeAgentReply) GetVipRemainingQuota() int32 {
	if m != nil {
		return m.VipRemainingQuota
	}
	return 0
}

func (m *ChuangkeAgentReply) GetSvipGivenQuota() int32 {
	if m != nil {
		return m.SvipGivenQuota
	}
	return 0
}

func (m *ChuangkeAgentReply) GetSvipRemainingQuota() int32 {
	if m != nil {
		return m.SvipRemainingQuota
	}
	return 0
}

func (m *ChuangkeAgentReply) GetHeadimgurl() string {
	if m != nil {
		return m.Headimgurl
	}
	return ""
}

func (m *ChuangkeAgentReply) GetReNickname() string {
	if m != nil {
		return m.ReNickname
	}
	return ""
}

func (m *ChuangkeAgentReply) GetPVipType() int32 {
	if m != nil {
		return m.PVipType
	}
	return 0
}

func (m *ChuangkeAgentReply) GetPuserId() int32 {
	if m != nil {
		return m.PuserId
	}
	return 0
}

func (m *ChuangkeAgentReply) GetAaId() int32 {
	if m != nil {
		return m.AaId
	}
	return 0
}

func (m *ChuangkeAgentReply) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

func (m *ChuangkeAgentReply) GetApplyStatus() int32 {
	if m != nil {
		return m.ApplyStatus
	}
	return 0
}

func (m *ChuangkeAgentReply) GetAuditAt() string {
	if m != nil {
		return m.AuditAt
	}
	return ""
}

func (m *ChuangkeAgentReply) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *ChuangkeAgentReply) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *ChuangkeAgentReply) GetArea() string {
	if m != nil {
		return m.Area
	}
	return ""
}

func (m *ChuangkeAgentReply) GetAgentGivenQuota() int32 {
	if m != nil {
		return m.AgentGivenQuota
	}
	return 0
}

func (m *ChuangkeAgentReply) GetAgentRemainingQuota() int32 {
	if m != nil {
		return m.AgentRemainingQuota
	}
	return 0
}

func (m *ChuangkeAgentReply) GetPartnerGivenQuota() int32 {
	if m != nil {
		return m.PartnerGivenQuota
	}
	return 0
}

func (m *ChuangkeAgentReply) GetPartnerRemainingQuota() int32 {
	if m != nil {
		return m.PartnerRemainingQuota
	}
	return 0
}

func (m *ChuangkeAgentReply) GetBranchGivenQuota() int32 {
	if m != nil {
		return m.BranchGivenQuota
	}
	return 0
}

func (m *ChuangkeAgentReply) GetBranchRemainingQuota() int32 {
	if m != nil {
		return m.BranchRemainingQuota
	}
	return 0
}

func (m *ChuangkeAgentReply) GetShowUserId() string {
	if m != nil {
		return m.ShowUserId
	}
	return ""
}

func (m *ChuangkeAgentReply) GetShowpUserId() string {
	if m != nil {
		return m.ShowpUserId
	}
	return ""
}

func (m *ChuangkeAgentReply) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *ChuangkeAgentReply) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *ChuangkeAgentReply) GetDeletedAt() string {
	if m != nil {
		return m.DeletedAt
	}
	return ""
}

// 创客代理商返回列表返回数据
type ChuangkeAgentListReply struct {
	AgentList            []*ChuangkeAgentReply `protobuf:"bytes,1,rep,name=agentList,proto3" json:"agentList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ChuangkeAgentListReply) Reset()         { *m = ChuangkeAgentListReply{} }
func (m *ChuangkeAgentListReply) String() string { return proto.CompactTextString(m) }
func (*ChuangkeAgentListReply) ProtoMessage()    {}
func (*ChuangkeAgentListReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_435bad2763cab632, []int{2}
}

func (m *ChuangkeAgentListReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChuangkeAgentListReply.Unmarshal(m, b)
}
func (m *ChuangkeAgentListReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChuangkeAgentListReply.Marshal(b, m, deterministic)
}
func (m *ChuangkeAgentListReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChuangkeAgentListReply.Merge(m, src)
}
func (m *ChuangkeAgentListReply) XXX_Size() int {
	return xxx_messageInfo_ChuangkeAgentListReply.Size(m)
}
func (m *ChuangkeAgentListReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ChuangkeAgentListReply.DiscardUnknown(m)
}

var xxx_messageInfo_ChuangkeAgentListReply proto.InternalMessageInfo

func (m *ChuangkeAgentListReply) GetAgentList() []*ChuangkeAgentReply {
	if m != nil {
		return m.AgentList
	}
	return nil
}

func init() {
	proto.RegisterType((*ChuangkeAgentRequest)(nil), "userAgentChuangke.ChuangkeAgentRequest")
	proto.RegisterType((*ChuangkeAgentReply)(nil), "userAgentChuangke.ChuangkeAgentReply")
	proto.RegisterType((*ChuangkeAgentListReply)(nil), "userAgentChuangke.ChuangkeAgentListReply")
}

func init() {
	proto.RegisterFile("grpc/proto/userAgentChuangke/userAgentChuangke.proto", fileDescriptor_435bad2763cab632)
}

var fileDescriptor_435bad2763cab632 = []byte{
	// 676 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcb, 0x4e, 0x1b, 0x3d,
	0x14, 0x26, 0x04, 0x42, 0x72, 0x80, 0x10, 0xcc, 0xe5, 0xf7, 0x4f, 0x29, 0x8d, 0xa2, 0x5e, 0xd2,
	0xaa, 0x82, 0x2a, 0xe5, 0x05, 0x28, 0x0b, 0x84, 0x54, 0x21, 0x35, 0xbd, 0xec, 0xba, 0x30, 0x33,
	0xd6, 0xc4, 0x62, 0x32, 0xe3, 0x7a, 0x3c, 0xa9, 0x52, 0x55, 0x5d, 0xf7, 0x21, 0xfa, 0x04, 0x7d,
	0xca, 0xca, 0xc7, 0x33, 0xc9, 0xdc, 0x20, 0x2c, 0xda, 0xdd, 0x39, 0xdf, 0x77, 0x6e, 0xf2, 0x77,
	0x6c, 0xc3, 0xa9, 0xa7, 0xa4, 0x73, 0x22, 0x55, 0xa8, 0xc3, 0x93, 0x38, 0xe2, 0xea, 0xcc, 0xe3,
	0x81, 0x3e, 0x1f, 0xc5, 0x2c, 0xf0, 0x6e, 0x78, 0x19, 0x39, 0xc6, 0x48, 0xb2, 0x5d, 0x22, 0x7a,
	0xbf, 0xea, 0xb0, 0x9b, 0x3a, 0xc8, 0x0c, 0xf9, 0x97, 0x98, 0x47, 0x9a, 0xb4, 0x61, 0x59, 0xb8,
	0xb4, 0xd6, 0xad, 0xf5, 0x57, 0x87, 0xcb, 0xc2, 0x25, 0x1d, 0xa8, 0xc7, 0xc2, 0xa5, 0xcb, 0xdd,
	0x5a, 0xbf, 0x35, 0x34, 0x26, 0xd9, 0x87, 0x86, 0xa9, 0x77, 0xe9, 0xd2, 0x3a, 0x46, 0x25, 0x1e,
	0x39, 0x80, 0x66, 0x20, 0x9c, 0x9b, 0x80, 0x8d, 0x39, 0x5d, 0xc1, 0xf0, 0x99, 0x6f, 0x38, 0xc5,
	0x99, 0x7f, 0x65, 0xb8, 0x55, 0xcb, 0xa5, 0xbe, 0xa9, 0x37, 0x0e, 0xaf, 0x85, 0xcf, 0x69, 0x03,
	0x99, 0xc4, 0x23, 0x14, 0xd6, 0x26, 0x42, 0x7e, 0x98, 0x4a, 0x4e, 0xd7, 0xb0, 0x51, 0xea, 0x92,
	0x23, 0x00, 0xc5, 0xaf, 0xd2, 0x5e, 0x4d, 0xcc, 0xca, 0x20, 0xa6, 0x9b, 0xfc, 0x94, 0xa4, 0xb6,
	0x30, 0x75, 0xe6, 0x9b, 0xaa, 0x32, 0x19, 0x1f, 0x6c, 0xd5, 0xc4, 0x25, 0x5d, 0x58, 0x67, 0x52,
	0xfa, 0xd3, 0xf7, 0x9a, 0xe9, 0x38, 0xa2, 0xeb, 0xc8, 0x66, 0x21, 0xd3, 0x37, 0x1a, 0x85, 0x5f,
	0x3f, 0xda, 0xf4, 0x0d, 0xdb, 0x77, 0x8e, 0x98, 0x0a, 0xc6, 0x93, 0x49, 0xc0, 0x26, 0x06, 0x64,
	0x21, 0x42, 0x60, 0x45, 0x32, 0x8f, 0xd3, 0x36, 0x16, 0x47, 0xdb, 0x60, 0x91, 0xf8, 0xc6, 0xe9,
	0x96, 0xc5, 0x8c, 0xdd, 0xfb, 0xdd, 0x04, 0x52, 0x90, 0x47, 0xfa, 0xd3, 0x7f, 0x24, 0xce, 0x5c,
	0x80, 0xd5, 0xdb, 0x04, 0x68, 0xe4, 0x05, 0x78, 0x0c, 0x9b, 0x13, 0x21, 0x2f, 0xc4, 0x84, 0x07,
	0xef, 0xe2, 0x50, 0xb3, 0x44, 0xa0, 0x3c, 0x48, 0x5e, 0xc2, 0xf6, 0x44, 0xc8, 0x21, 0x1f, 0x33,
	0x11, 0x88, 0xc0, 0xb3, 0x91, 0x4d, 0x8c, 0x2c, 0x13, 0xe4, 0x29, 0xb4, 0xa3, 0x7c, 0x51, 0x2b,
	0x5d, 0x01, 0x25, 0xc7, 0x40, 0xa2, 0x72, 0x59, 0xab, 0x65, 0x05, 0x63, 0x44, 0x1b, 0x71, 0xe6,
	0x8a, 0xb1, 0x17, 0x2b, 0x1f, 0x55, 0x6d, 0x0d, 0x33, 0x48, 0x61, 0x99, 0x36, 0xee, 0x5c, 0xa6,
	0xcd, 0xdb, 0x97, 0xa9, 0x9d, 0x5f, 0x26, 0x02, 0x2b, 0x8c, 0x5d, 0xba, 0xa9, 0xa8, 0xc6, 0xce,
	0x5d, 0x82, 0x4e, 0xe1, 0x12, 0x14, 0x96, 0x6f, 0xbb, 0xbc, 0x7c, 0x14, 0xd6, 0x58, 0xec, 0x0a,
	0x7d, 0xa6, 0x29, 0xc1, 0xe4, 0xd4, 0xc5, 0x09, 0x55, 0x38, 0x11, 0x81, 0xc3, 0xe9, 0x8e, 0xad,
	0x9b, 0xfa, 0x66, 0x0e, 0x47, 0xe8, 0x29, 0xdd, 0x45, 0x1c, 0x6d, 0x9c, 0x4d, 0x71, 0x46, 0xf7,
	0x2c, 0x66, 0x6c, 0xd2, 0x87, 0x2d, 0x66, 0xf6, 0x2c, 0x73, 0xfc, 0xfb, 0x38, 0x43, 0x11, 0x26,
	0xaf, 0x60, 0x87, 0xd9, 0x8d, 0xcc, 0x09, 0xf0, 0x1f, 0x46, 0x57, 0x51, 0x66, 0x0f, 0x24, 0x53,
	0x3a, 0xe0, 0x2a, 0x53, 0x9d, 0xda, 0x3d, 0x28, 0x11, 0xe4, 0x14, 0xf6, 0x12, 0xb0, 0xd0, 0xe1,
	0x7f, 0xcc, 0xa8, 0x26, 0xc9, 0x0b, 0xe8, 0x5c, 0x2b, 0x16, 0x38, 0xa3, 0x4c, 0x8b, 0x03, 0x4c,
	0x28, 0xe1, 0x64, 0x00, 0xbb, 0x16, 0x2b, 0x34, 0x78, 0x80, 0xf1, 0x95, 0x5c, 0xe1, 0xea, 0x1f,
	0x2e, 0xba, 0xfa, 0x0f, 0xcb, 0x57, 0xff, 0x10, 0x5a, 0x8e, 0xe2, 0x4c, 0x73, 0xf7, 0x4c, 0xd3,
	0x23, 0xe4, 0xe7, 0x80, 0x61, 0x63, 0xe9, 0x26, 0xec, 0x23, 0xcb, 0xce, 0x00, 0xc3, 0xba, 0xdc,
	0xe7, 0x96, 0xed, 0x5a, 0x76, 0x06, 0xf4, 0x3e, 0xc3, 0x7e, 0xee, 0xad, 0x78, 0x2b, 0xa2, 0xe4,
	0xbd, 0x38, 0x87, 0x16, 0x4b, 0x11, 0x5a, 0xeb, 0xd6, 0xfb, 0xeb, 0x83, 0x27, 0xc7, 0xe5, 0x5f,
	0xa2, 0xfc, 0xd2, 0x0c, 0xe7, 0x79, 0x83, 0x1f, 0xb0, 0x73, 0x8e, 0x73, 0xe6, 0xc2, 0x88, 0x57,
	0x0d, 0x3f, 0x5b, 0x5c, 0x1f, 0x3f, 0x9a, 0x83, 0xfb, 0x0d, 0xd2, 0x5b, 0x1a, 0x7c, 0x87, 0xce,
	0x05, 0xd7, 0xf9, 0x2e, 0xa3, 0x0a, 0xec, 0xde, 0x9d, 0x9f, 0x2f, 0x0a, 0x9c, 0x1d, 0x60, 0x6f,
	0x69, 0xf0, 0xb3, 0x06, 0xb4, 0xd8, 0xea, 0xcd, 0x34, 0xd1, 0xd4, 0xbf, 0x83, 0xfb, 0xeb, 0x07,
	0x71, 0xdd, 0xc0, 0xdf, 0xfc, 0xf5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x40, 0x47, 0xe1,
	0x05, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CreateChuangkeAgentClient is the client API for CreateChuangkeAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CreateChuangkeAgentClient interface {
	CreateChuangkeAgent(ctx context.Context, in *ChuangkeAgentRequest, opts ...grpc.CallOption) (*ChuangkeAgentReply, error)
}

type createChuangkeAgentClient struct {
	cc *grpc.ClientConn
}

func NewCreateChuangkeAgentClient(cc *grpc.ClientConn) CreateChuangkeAgentClient {
	return &createChuangkeAgentClient{cc}
}

func (c *createChuangkeAgentClient) CreateChuangkeAgent(ctx context.Context, in *ChuangkeAgentRequest, opts ...grpc.CallOption) (*ChuangkeAgentReply, error) {
	out := new(ChuangkeAgentReply)
	err := c.cc.Invoke(ctx, "/userAgentChuangke.CreateChuangkeAgent/CreateChuangkeAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateChuangkeAgentServer is the server API for CreateChuangkeAgent service.
type CreateChuangkeAgentServer interface {
	CreateChuangkeAgent(context.Context, *ChuangkeAgentRequest) (*ChuangkeAgentReply, error)
}

// UnimplementedCreateChuangkeAgentServer can be embedded to have forward compatible implementations.
type UnimplementedCreateChuangkeAgentServer struct {
}

func (*UnimplementedCreateChuangkeAgentServer) CreateChuangkeAgent(ctx context.Context, req *ChuangkeAgentRequest) (*ChuangkeAgentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChuangkeAgent not implemented")
}

func RegisterCreateChuangkeAgentServer(s *grpc.Server, srv CreateChuangkeAgentServer) {
	s.RegisterService(&_CreateChuangkeAgent_serviceDesc, srv)
}

func _CreateChuangkeAgent_CreateChuangkeAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChuangkeAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateChuangkeAgentServer).CreateChuangkeAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userAgentChuangke.CreateChuangkeAgent/CreateChuangkeAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateChuangkeAgentServer).CreateChuangkeAgent(ctx, req.(*ChuangkeAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CreateChuangkeAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "userAgentChuangke.CreateChuangkeAgent",
	HandlerType: (*CreateChuangkeAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChuangkeAgent",
			Handler:    _CreateChuangkeAgent_CreateChuangkeAgent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/proto/userAgentChuangke/userAgentChuangke.proto",
}

// GetChuangkeAgentClient is the client API for GetChuangkeAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GetChuangkeAgentClient interface {
	GetChuangkeAgent(ctx context.Context, in *ChuangkeAgentRequest, opts ...grpc.CallOption) (*ChuangkeAgentListReply, error)
}

type getChuangkeAgentClient struct {
	cc *grpc.ClientConn
}

func NewGetChuangkeAgentClient(cc *grpc.ClientConn) GetChuangkeAgentClient {
	return &getChuangkeAgentClient{cc}
}

func (c *getChuangkeAgentClient) GetChuangkeAgent(ctx context.Context, in *ChuangkeAgentRequest, opts ...grpc.CallOption) (*ChuangkeAgentListReply, error) {
	out := new(ChuangkeAgentListReply)
	err := c.cc.Invoke(ctx, "/userAgentChuangke.GetChuangkeAgent/GetChuangkeAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetChuangkeAgentServer is the server API for GetChuangkeAgent service.
type GetChuangkeAgentServer interface {
	GetChuangkeAgent(context.Context, *ChuangkeAgentRequest) (*ChuangkeAgentListReply, error)
}

// UnimplementedGetChuangkeAgentServer can be embedded to have forward compatible implementations.
type UnimplementedGetChuangkeAgentServer struct {
}

func (*UnimplementedGetChuangkeAgentServer) GetChuangkeAgent(ctx context.Context, req *ChuangkeAgentRequest) (*ChuangkeAgentListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChuangkeAgent not implemented")
}

func RegisterGetChuangkeAgentServer(s *grpc.Server, srv GetChuangkeAgentServer) {
	s.RegisterService(&_GetChuangkeAgent_serviceDesc, srv)
}

func _GetChuangkeAgent_GetChuangkeAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChuangkeAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetChuangkeAgentServer).GetChuangkeAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userAgentChuangke.GetChuangkeAgent/GetChuangkeAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetChuangkeAgentServer).GetChuangkeAgent(ctx, req.(*ChuangkeAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GetChuangkeAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "userAgentChuangke.GetChuangkeAgent",
	HandlerType: (*GetChuangkeAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChuangkeAgent",
			Handler:    _GetChuangkeAgent_GetChuangkeAgent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/proto/userAgentChuangke/userAgentChuangke.proto",
}

// GetChuangkeAgentByUserIdClient is the client API for GetChuangkeAgentByUserId service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GetChuangkeAgentByUserIdClient interface {
	GetChuangkeAgentByUserId(ctx context.Context, in *ChuangkeAgentRequest, opts ...grpc.CallOption) (*ChuangkeAgentReply, error)
}

type getChuangkeAgentByUserIdClient struct {
	cc *grpc.ClientConn
}

func NewGetChuangkeAgentByUserIdClient(cc *grpc.ClientConn) GetChuangkeAgentByUserIdClient {
	return &getChuangkeAgentByUserIdClient{cc}
}

func (c *getChuangkeAgentByUserIdClient) GetChuangkeAgentByUserId(ctx context.Context, in *ChuangkeAgentRequest, opts ...grpc.CallOption) (*ChuangkeAgentReply, error) {
	out := new(ChuangkeAgentReply)
	err := c.cc.Invoke(ctx, "/userAgentChuangke.GetChuangkeAgentByUserId/GetChuangkeAgentByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetChuangkeAgentByUserIdServer is the server API for GetChuangkeAgentByUserId service.
type GetChuangkeAgentByUserIdServer interface {
	GetChuangkeAgentByUserId(context.Context, *ChuangkeAgentRequest) (*ChuangkeAgentReply, error)
}

// UnimplementedGetChuangkeAgentByUserIdServer can be embedded to have forward compatible implementations.
type UnimplementedGetChuangkeAgentByUserIdServer struct {
}

func (*UnimplementedGetChuangkeAgentByUserIdServer) GetChuangkeAgentByUserId(ctx context.Context, req *ChuangkeAgentRequest) (*ChuangkeAgentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChuangkeAgentByUserId not implemented")
}

func RegisterGetChuangkeAgentByUserIdServer(s *grpc.Server, srv GetChuangkeAgentByUserIdServer) {
	s.RegisterService(&_GetChuangkeAgentByUserId_serviceDesc, srv)
}

func _GetChuangkeAgentByUserId_GetChuangkeAgentByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChuangkeAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetChuangkeAgentByUserIdServer).GetChuangkeAgentByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userAgentChuangke.GetChuangkeAgentByUserId/GetChuangkeAgentByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetChuangkeAgentByUserIdServer).GetChuangkeAgentByUserId(ctx, req.(*ChuangkeAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GetChuangkeAgentByUserId_serviceDesc = grpc.ServiceDesc{
	ServiceName: "userAgentChuangke.GetChuangkeAgentByUserId",
	HandlerType: (*GetChuangkeAgentByUserIdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChuangkeAgentByUserId",
			Handler:    _GetChuangkeAgentByUserId_GetChuangkeAgentByUserId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/proto/userAgentChuangke/userAgentChuangke.proto",
}