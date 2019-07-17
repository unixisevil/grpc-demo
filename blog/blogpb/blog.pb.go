// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blog/blogpb/blog.proto

package blogpb

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

type Blog struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorId             string   `protobuf:"bytes,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blog) Reset()         { *m = Blog{} }
func (m *Blog) String() string { return proto.CompactTextString(m) }
func (*Blog) ProtoMessage()    {}
func (*Blog) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b0406114889fe6, []int{0}
}

func (m *Blog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blog.Unmarshal(m, b)
}
func (m *Blog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blog.Marshal(b, m, deterministic)
}
func (m *Blog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blog.Merge(m, src)
}
func (m *Blog) XXX_Size() int {
	return xxx_messageInfo_Blog.Size(m)
}
func (m *Blog) XXX_DiscardUnknown() {
	xxx_messageInfo_Blog.DiscardUnknown(m)
}

var xxx_messageInfo_Blog proto.InternalMessageInfo

func (m *Blog) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Blog) GetAuthorId() string {
	if m != nil {
		return m.AuthorId
	}
	return ""
}

func (m *Blog) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Blog) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type CreateBlogReq struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBlogReq) Reset()         { *m = CreateBlogReq{} }
func (m *CreateBlogReq) String() string { return proto.CompactTextString(m) }
func (*CreateBlogReq) ProtoMessage()    {}
func (*CreateBlogReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b0406114889fe6, []int{1}
}

func (m *CreateBlogReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBlogReq.Unmarshal(m, b)
}
func (m *CreateBlogReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBlogReq.Marshal(b, m, deterministic)
}
func (m *CreateBlogReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBlogReq.Merge(m, src)
}
func (m *CreateBlogReq) XXX_Size() int {
	return xxx_messageInfo_CreateBlogReq.Size(m)
}
func (m *CreateBlogReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBlogReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBlogReq proto.InternalMessageInfo

func (m *CreateBlogReq) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type CreateBlogRes struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBlogRes) Reset()         { *m = CreateBlogRes{} }
func (m *CreateBlogRes) String() string { return proto.CompactTextString(m) }
func (*CreateBlogRes) ProtoMessage()    {}
func (*CreateBlogRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b0406114889fe6, []int{2}
}

func (m *CreateBlogRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBlogRes.Unmarshal(m, b)
}
func (m *CreateBlogRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBlogRes.Marshal(b, m, deterministic)
}
func (m *CreateBlogRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBlogRes.Merge(m, src)
}
func (m *CreateBlogRes) XXX_Size() int {
	return xxx_messageInfo_CreateBlogRes.Size(m)
}
func (m *CreateBlogRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBlogRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBlogRes proto.InternalMessageInfo

func (m *CreateBlogRes) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadBlogReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadBlogReq) Reset()         { *m = ReadBlogReq{} }
func (m *ReadBlogReq) String() string { return proto.CompactTextString(m) }
func (*ReadBlogReq) ProtoMessage()    {}
func (*ReadBlogReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b0406114889fe6, []int{3}
}

func (m *ReadBlogReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadBlogReq.Unmarshal(m, b)
}
func (m *ReadBlogReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadBlogReq.Marshal(b, m, deterministic)
}
func (m *ReadBlogReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadBlogReq.Merge(m, src)
}
func (m *ReadBlogReq) XXX_Size() int {
	return xxx_messageInfo_ReadBlogReq.Size(m)
}
func (m *ReadBlogReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadBlogReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReadBlogReq proto.InternalMessageInfo

func (m *ReadBlogReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadBlogRes struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadBlogRes) Reset()         { *m = ReadBlogRes{} }
func (m *ReadBlogRes) String() string { return proto.CompactTextString(m) }
func (*ReadBlogRes) ProtoMessage()    {}
func (*ReadBlogRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b0406114889fe6, []int{4}
}

func (m *ReadBlogRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadBlogRes.Unmarshal(m, b)
}
func (m *ReadBlogRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadBlogRes.Marshal(b, m, deterministic)
}
func (m *ReadBlogRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadBlogRes.Merge(m, src)
}
func (m *ReadBlogRes) XXX_Size() int {
	return xxx_messageInfo_ReadBlogRes.Size(m)
}
func (m *ReadBlogRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadBlogRes.DiscardUnknown(m)
}

var xxx_messageInfo_ReadBlogRes proto.InternalMessageInfo

func (m *ReadBlogRes) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type UpdateBlogReq struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBlogReq) Reset()         { *m = UpdateBlogReq{} }
func (m *UpdateBlogReq) String() string { return proto.CompactTextString(m) }
func (*UpdateBlogReq) ProtoMessage()    {}
func (*UpdateBlogReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b0406114889fe6, []int{5}
}

func (m *UpdateBlogReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBlogReq.Unmarshal(m, b)
}
func (m *UpdateBlogReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBlogReq.Marshal(b, m, deterministic)
}
func (m *UpdateBlogReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBlogReq.Merge(m, src)
}
func (m *UpdateBlogReq) XXX_Size() int {
	return xxx_messageInfo_UpdateBlogReq.Size(m)
}
func (m *UpdateBlogReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBlogReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBlogReq proto.InternalMessageInfo

func (m *UpdateBlogReq) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type UpdateBlogRes struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBlogRes) Reset()         { *m = UpdateBlogRes{} }
func (m *UpdateBlogRes) String() string { return proto.CompactTextString(m) }
func (*UpdateBlogRes) ProtoMessage()    {}
func (*UpdateBlogRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b0406114889fe6, []int{6}
}

func (m *UpdateBlogRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBlogRes.Unmarshal(m, b)
}
func (m *UpdateBlogRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBlogRes.Marshal(b, m, deterministic)
}
func (m *UpdateBlogRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBlogRes.Merge(m, src)
}
func (m *UpdateBlogRes) XXX_Size() int {
	return xxx_messageInfo_UpdateBlogRes.Size(m)
}
func (m *UpdateBlogRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBlogRes.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBlogRes proto.InternalMessageInfo

func (m *UpdateBlogRes) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type DelBlogReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelBlogReq) Reset()         { *m = DelBlogReq{} }
func (m *DelBlogReq) String() string { return proto.CompactTextString(m) }
func (*DelBlogReq) ProtoMessage()    {}
func (*DelBlogReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b0406114889fe6, []int{7}
}

func (m *DelBlogReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelBlogReq.Unmarshal(m, b)
}
func (m *DelBlogReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelBlogReq.Marshal(b, m, deterministic)
}
func (m *DelBlogReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelBlogReq.Merge(m, src)
}
func (m *DelBlogReq) XXX_Size() int {
	return xxx_messageInfo_DelBlogReq.Size(m)
}
func (m *DelBlogReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DelBlogReq.DiscardUnknown(m)
}

var xxx_messageInfo_DelBlogReq proto.InternalMessageInfo

func (m *DelBlogReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DelBlogRes struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelBlogRes) Reset()         { *m = DelBlogRes{} }
func (m *DelBlogRes) String() string { return proto.CompactTextString(m) }
func (*DelBlogRes) ProtoMessage()    {}
func (*DelBlogRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b0406114889fe6, []int{8}
}

func (m *DelBlogRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelBlogRes.Unmarshal(m, b)
}
func (m *DelBlogRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelBlogRes.Marshal(b, m, deterministic)
}
func (m *DelBlogRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelBlogRes.Merge(m, src)
}
func (m *DelBlogRes) XXX_Size() int {
	return xxx_messageInfo_DelBlogRes.Size(m)
}
func (m *DelBlogRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DelBlogRes.DiscardUnknown(m)
}

var xxx_messageInfo_DelBlogRes proto.InternalMessageInfo

func (m *DelBlogRes) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListBlogReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBlogReq) Reset()         { *m = ListBlogReq{} }
func (m *ListBlogReq) String() string { return proto.CompactTextString(m) }
func (*ListBlogReq) ProtoMessage()    {}
func (*ListBlogReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b0406114889fe6, []int{9}
}

func (m *ListBlogReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBlogReq.Unmarshal(m, b)
}
func (m *ListBlogReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBlogReq.Marshal(b, m, deterministic)
}
func (m *ListBlogReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBlogReq.Merge(m, src)
}
func (m *ListBlogReq) XXX_Size() int {
	return xxx_messageInfo_ListBlogReq.Size(m)
}
func (m *ListBlogReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBlogReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListBlogReq proto.InternalMessageInfo

type ListBlogRes struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBlogRes) Reset()         { *m = ListBlogRes{} }
func (m *ListBlogRes) String() string { return proto.CompactTextString(m) }
func (*ListBlogRes) ProtoMessage()    {}
func (*ListBlogRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b0406114889fe6, []int{10}
}

func (m *ListBlogRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBlogRes.Unmarshal(m, b)
}
func (m *ListBlogRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBlogRes.Marshal(b, m, deterministic)
}
func (m *ListBlogRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBlogRes.Merge(m, src)
}
func (m *ListBlogRes) XXX_Size() int {
	return xxx_messageInfo_ListBlogRes.Size(m)
}
func (m *ListBlogRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBlogRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListBlogRes proto.InternalMessageInfo

func (m *ListBlogRes) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

func init() {
	proto.RegisterType((*Blog)(nil), "blog.Blog")
	proto.RegisterType((*CreateBlogReq)(nil), "blog.CreateBlogReq")
	proto.RegisterType((*CreateBlogRes)(nil), "blog.CreateBlogRes")
	proto.RegisterType((*ReadBlogReq)(nil), "blog.ReadBlogReq")
	proto.RegisterType((*ReadBlogRes)(nil), "blog.ReadBlogRes")
	proto.RegisterType((*UpdateBlogReq)(nil), "blog.UpdateBlogReq")
	proto.RegisterType((*UpdateBlogRes)(nil), "blog.UpdateBlogRes")
	proto.RegisterType((*DelBlogReq)(nil), "blog.DelBlogReq")
	proto.RegisterType((*DelBlogRes)(nil), "blog.DelBlogRes")
	proto.RegisterType((*ListBlogReq)(nil), "blog.ListBlogReq")
	proto.RegisterType((*ListBlogRes)(nil), "blog.ListBlogRes")
}

func init() { proto.RegisterFile("blog/blogpb/blog.proto", fileDescriptor_a4b0406114889fe6) }

var fileDescriptor_a4b0406114889fe6 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x4f, 0xb3, 0x40,
	0x18, 0x6c, 0x79, 0xfb, 0xb6, 0xf4, 0x21, 0x18, 0x7d, 0x34, 0x86, 0xe0, 0x67, 0x38, 0x79, 0xb1,
	0x18, 0xf4, 0xd0, 0x73, 0xf5, 0x62, 0xe2, 0x09, 0xe3, 0xc5, 0x8b, 0x01, 0x76, 0x53, 0x37, 0x21,
	0x5d, 0x84, 0xd5, 0x5f, 0xe3, 0x8f, 0x35, 0xbb, 0xdb, 0x0d, 0x5f, 0x35, 0xc4, 0x0b, 0xf0, 0xcc,
	0xec, 0x30, 0xb3, 0xb3, 0x59, 0x38, 0x4e, 0x73, 0xbe, 0x0e, 0xe5, 0xa3, 0x48, 0xd5, 0x6b, 0x51,
	0x94, 0x5c, 0x70, 0x9c, 0xc8, 0xef, 0x20, 0x83, 0xc9, 0x2a, 0xe7, 0x6b, 0xdc, 0x03, 0x8b, 0x11,
	0x6f, 0x7c, 0x39, 0xbe, 0x9a, 0xc7, 0x16, 0x23, 0x78, 0x02, 0xf3, 0xe4, 0x53, 0xbc, 0xf3, 0xf2,
	0x8d, 0x11, 0xcf, 0x52, 0xb0, 0xad, 0x81, 0x47, 0x82, 0x47, 0xf0, 0x5f, 0x30, 0x91, 0x53, 0xef,
	0x9f, 0x22, 0xf4, 0x80, 0x1e, 0xcc, 0x32, 0xbe, 0x11, 0x74, 0x23, 0xbc, 0x89, 0xc2, 0xcd, 0x18,
	0x84, 0xe0, 0xde, 0x97, 0x34, 0x11, 0x54, 0x5a, 0xc5, 0xf4, 0x03, 0xcf, 0x41, 0xb9, 0x2b, 0x3f,
	0x27, 0x82, 0x85, 0x8a, 0xa5, 0x48, 0x9d, 0xea, 0xa2, 0x2d, 0xa8, 0xba, 0xf1, 0x82, 0x33, 0x70,
	0x62, 0x9a, 0x10, 0xf3, 0xbf, 0x2e, 0x7d, 0xdd, 0xa4, 0xab, 0x41, 0xbb, 0x10, 0xdc, 0x97, 0x82,
	0xfc, 0x21, 0x5f, 0x47, 0x30, 0xec, 0x70, 0x0a, 0xf0, 0x40, 0xf3, 0xdf, 0xe2, 0x36, 0xd9, 0xfe,
	0x5e, 0x5d, 0x70, 0x9e, 0x58, 0x25, 0xb6, 0x62, 0xb9, 0xb7, 0x7a, 0x1c, 0x74, 0x8e, 0xbe, 0x2d,
	0x70, 0xe4, 0xf8, 0x4c, 0xcb, 0x2f, 0x96, 0x51, 0x5c, 0x02, 0xd4, 0xd5, 0xe2, 0xa1, 0x5e, 0xdf,
	0x3a, 0x1d, 0x7f, 0x07, 0x58, 0x05, 0x23, 0x8c, 0xc0, 0x36, 0xa5, 0xe2, 0x81, 0x5e, 0xd2, 0x38,
	0x03, 0xbf, 0x07, 0x49, 0xcd, 0x12, 0xa0, 0x2e, 0xca, 0xb8, 0xb5, 0xba, 0xf6, 0x77, 0x80, 0x52,
	0x19, 0xc2, 0x6c, 0xdb, 0x09, 0xee, 0xeb, 0x15, 0x75, 0x81, 0x7e, 0x17, 0x91, 0x82, 0x3b, 0xb0,
	0x4d, 0x2f, 0x26, 0x5e, 0xa3, 0x36, 0xbf, 0x07, 0x55, 0xc1, 0xe8, 0x66, 0xbc, 0xb2, 0x5f, 0xa7,
	0xfa, 0x6a, 0xa4, 0x53, 0x75, 0x2d, 0x6e, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x72, 0xb9,
	0x1c, 0x30, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlogServiceClient is the client API for BlogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlogServiceClient interface {
	CreateBlog(ctx context.Context, in *CreateBlogReq, opts ...grpc.CallOption) (*CreateBlogRes, error)
	ReadBlog(ctx context.Context, in *ReadBlogReq, opts ...grpc.CallOption) (*ReadBlogRes, error)
	UpdateBlog(ctx context.Context, in *UpdateBlogReq, opts ...grpc.CallOption) (*UpdateBlogRes, error)
	DelBlog(ctx context.Context, in *DelBlogReq, opts ...grpc.CallOption) (*DelBlogRes, error)
	ListBlog(ctx context.Context, in *ListBlogReq, opts ...grpc.CallOption) (BlogService_ListBlogClient, error)
}

type blogServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlogServiceClient(cc *grpc.ClientConn) BlogServiceClient {
	return &blogServiceClient{cc}
}

func (c *blogServiceClient) CreateBlog(ctx context.Context, in *CreateBlogReq, opts ...grpc.CallOption) (*CreateBlogRes, error) {
	out := new(CreateBlogRes)
	err := c.cc.Invoke(ctx, "/blog.BlogService/CreateBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) ReadBlog(ctx context.Context, in *ReadBlogReq, opts ...grpc.CallOption) (*ReadBlogRes, error) {
	out := new(ReadBlogRes)
	err := c.cc.Invoke(ctx, "/blog.BlogService/ReadBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) UpdateBlog(ctx context.Context, in *UpdateBlogReq, opts ...grpc.CallOption) (*UpdateBlogRes, error) {
	out := new(UpdateBlogRes)
	err := c.cc.Invoke(ctx, "/blog.BlogService/UpdateBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) DelBlog(ctx context.Context, in *DelBlogReq, opts ...grpc.CallOption) (*DelBlogRes, error) {
	out := new(DelBlogRes)
	err := c.cc.Invoke(ctx, "/blog.BlogService/DelBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) ListBlog(ctx context.Context, in *ListBlogReq, opts ...grpc.CallOption) (BlogService_ListBlogClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BlogService_serviceDesc.Streams[0], "/blog.BlogService/ListBlog", opts...)
	if err != nil {
		return nil, err
	}
	x := &blogServiceListBlogClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlogService_ListBlogClient interface {
	Recv() (*ListBlogRes, error)
	grpc.ClientStream
}

type blogServiceListBlogClient struct {
	grpc.ClientStream
}

func (x *blogServiceListBlogClient) Recv() (*ListBlogRes, error) {
	m := new(ListBlogRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BlogServiceServer is the server API for BlogService service.
type BlogServiceServer interface {
	CreateBlog(context.Context, *CreateBlogReq) (*CreateBlogRes, error)
	ReadBlog(context.Context, *ReadBlogReq) (*ReadBlogRes, error)
	UpdateBlog(context.Context, *UpdateBlogReq) (*UpdateBlogRes, error)
	DelBlog(context.Context, *DelBlogReq) (*DelBlogRes, error)
	ListBlog(*ListBlogReq, BlogService_ListBlogServer) error
}

// UnimplementedBlogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBlogServiceServer struct {
}

func (*UnimplementedBlogServiceServer) CreateBlog(ctx context.Context, req *CreateBlogReq) (*CreateBlogRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlog not implemented")
}
func (*UnimplementedBlogServiceServer) ReadBlog(ctx context.Context, req *ReadBlogReq) (*ReadBlogRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadBlog not implemented")
}
func (*UnimplementedBlogServiceServer) UpdateBlog(ctx context.Context, req *UpdateBlogReq) (*UpdateBlogRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBlog not implemented")
}
func (*UnimplementedBlogServiceServer) DelBlog(ctx context.Context, req *DelBlogReq) (*DelBlogRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelBlog not implemented")
}
func (*UnimplementedBlogServiceServer) ListBlog(req *ListBlogReq, srv BlogService_ListBlogServer) error {
	return status.Errorf(codes.Unimplemented, "method ListBlog not implemented")
}

func RegisterBlogServiceServer(s *grpc.Server, srv BlogServiceServer) {
	s.RegisterService(&_BlogService_serviceDesc, srv)
}

func _BlogService_CreateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).CreateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/CreateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).CreateBlog(ctx, req.(*CreateBlogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_ReadBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadBlogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).ReadBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/ReadBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).ReadBlog(ctx, req.(*ReadBlogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_UpdateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBlogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).UpdateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/UpdateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).UpdateBlog(ctx, req.(*UpdateBlogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_DelBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelBlogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).DelBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/DelBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).DelBlog(ctx, req.(*DelBlogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_ListBlog_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListBlogReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlogServiceServer).ListBlog(m, &blogServiceListBlogServer{stream})
}

type BlogService_ListBlogServer interface {
	Send(*ListBlogRes) error
	grpc.ServerStream
}

type blogServiceListBlogServer struct {
	grpc.ServerStream
}

func (x *blogServiceListBlogServer) Send(m *ListBlogRes) error {
	return x.ServerStream.SendMsg(m)
}

var _BlogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blog.BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBlog",
			Handler:    _BlogService_CreateBlog_Handler,
		},
		{
			MethodName: "ReadBlog",
			Handler:    _BlogService_ReadBlog_Handler,
		},
		{
			MethodName: "UpdateBlog",
			Handler:    _BlogService_UpdateBlog_Handler,
		},
		{
			MethodName: "DelBlog",
			Handler:    _BlogService_DelBlog_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListBlog",
			Handler:       _BlogService_ListBlog_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "blog/blogpb/blog.proto",
}
