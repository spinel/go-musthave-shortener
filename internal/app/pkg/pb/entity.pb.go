// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/app/pkg/pb/entity.proto

package pb

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

type GetURLRequest struct {
	UrlCode              string   `protobuf:"bytes,1,opt,name=url_code,json=urlCode,proto3" json:"url_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetURLRequest) Reset()         { *m = GetURLRequest{} }
func (m *GetURLRequest) String() string { return proto.CompactTextString(m) }
func (*GetURLRequest) ProtoMessage()    {}
func (*GetURLRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f423ff2260e2d107, []int{0}
}

func (m *GetURLRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetURLRequest.Unmarshal(m, b)
}
func (m *GetURLRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetURLRequest.Marshal(b, m, deterministic)
}
func (m *GetURLRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetURLRequest.Merge(m, src)
}
func (m *GetURLRequest) XXX_Size() int {
	return xxx_messageInfo_GetURLRequest.Size(m)
}
func (m *GetURLRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetURLRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetURLRequest proto.InternalMessageInfo

func (m *GetURLRequest) GetUrlCode() string {
	if m != nil {
		return m.UrlCode
	}
	return ""
}

type GetURLResponse struct {
	EntityId             int64    `protobuf:"varint,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetURLResponse) Reset()         { *m = GetURLResponse{} }
func (m *GetURLResponse) String() string { return proto.CompactTextString(m) }
func (*GetURLResponse) ProtoMessage()    {}
func (*GetURLResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f423ff2260e2d107, []int{1}
}

func (m *GetURLResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetURLResponse.Unmarshal(m, b)
}
func (m *GetURLResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetURLResponse.Marshal(b, m, deterministic)
}
func (m *GetURLResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetURLResponse.Merge(m, src)
}
func (m *GetURLResponse) XXX_Size() int {
	return xxx_messageInfo_GetURLResponse.Size(m)
}
func (m *GetURLResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetURLResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetURLResponse proto.InternalMessageInfo

func (m *GetURLResponse) GetEntityId() int64 {
	if m != nil {
		return m.EntityId
	}
	return 0
}

func (m *GetURLResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type CreateURLRequest struct {
	UserUuid             string   `protobuf:"bytes,1,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Url                  string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateURLRequest) Reset()         { *m = CreateURLRequest{} }
func (m *CreateURLRequest) String() string { return proto.CompactTextString(m) }
func (*CreateURLRequest) ProtoMessage()    {}
func (*CreateURLRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f423ff2260e2d107, []int{2}
}

func (m *CreateURLRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateURLRequest.Unmarshal(m, b)
}
func (m *CreateURLRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateURLRequest.Marshal(b, m, deterministic)
}
func (m *CreateURLRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateURLRequest.Merge(m, src)
}
func (m *CreateURLRequest) XXX_Size() int {
	return xxx_messageInfo_CreateURLRequest.Size(m)
}
func (m *CreateURLRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateURLRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateURLRequest proto.InternalMessageInfo

func (m *CreateURLRequest) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *CreateURLRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *CreateURLRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type CreateURLResponse struct {
	EntityId             int64    `protobuf:"varint,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateURLResponse) Reset()         { *m = CreateURLResponse{} }
func (m *CreateURLResponse) String() string { return proto.CompactTextString(m) }
func (*CreateURLResponse) ProtoMessage()    {}
func (*CreateURLResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f423ff2260e2d107, []int{3}
}

func (m *CreateURLResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateURLResponse.Unmarshal(m, b)
}
func (m *CreateURLResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateURLResponse.Marshal(b, m, deterministic)
}
func (m *CreateURLResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateURLResponse.Merge(m, src)
}
func (m *CreateURLResponse) XXX_Size() int {
	return xxx_messageInfo_CreateURLResponse.Size(m)
}
func (m *CreateURLResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateURLResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateURLResponse proto.InternalMessageInfo

func (m *CreateURLResponse) GetEntityId() int64 {
	if m != nil {
		return m.EntityId
	}
	return 0
}

func (m *CreateURLResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type EnqueueDeleteRequest struct {
	Codes                []string `protobuf:"bytes,1,rep,name=codes,proto3" json:"codes,omitempty"`
	UserUuid             string   `protobuf:"bytes,2,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnqueueDeleteRequest) Reset()         { *m = EnqueueDeleteRequest{} }
func (m *EnqueueDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*EnqueueDeleteRequest) ProtoMessage()    {}
func (*EnqueueDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f423ff2260e2d107, []int{4}
}

func (m *EnqueueDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnqueueDeleteRequest.Unmarshal(m, b)
}
func (m *EnqueueDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnqueueDeleteRequest.Marshal(b, m, deterministic)
}
func (m *EnqueueDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnqueueDeleteRequest.Merge(m, src)
}
func (m *EnqueueDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_EnqueueDeleteRequest.Size(m)
}
func (m *EnqueueDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EnqueueDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EnqueueDeleteRequest proto.InternalMessageInfo

func (m *EnqueueDeleteRequest) GetCodes() []string {
	if m != nil {
		return m.Codes
	}
	return nil
}

func (m *EnqueueDeleteRequest) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

type EnqueueDeleteResponse struct {
	Error                string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnqueueDeleteResponse) Reset()         { *m = EnqueueDeleteResponse{} }
func (m *EnqueueDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*EnqueueDeleteResponse) ProtoMessage()    {}
func (*EnqueueDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f423ff2260e2d107, []int{5}
}

func (m *EnqueueDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnqueueDeleteResponse.Unmarshal(m, b)
}
func (m *EnqueueDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnqueueDeleteResponse.Marshal(b, m, deterministic)
}
func (m *EnqueueDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnqueueDeleteResponse.Merge(m, src)
}
func (m *EnqueueDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_EnqueueDeleteResponse.Size(m)
}
func (m *EnqueueDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EnqueueDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EnqueueDeleteResponse proto.InternalMessageInfo

func (m *EnqueueDeleteResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*GetURLRequest)(nil), "entity.GetURLRequest")
	proto.RegisterType((*GetURLResponse)(nil), "entity.GetURLResponse")
	proto.RegisterType((*CreateURLRequest)(nil), "entity.CreateURLRequest")
	proto.RegisterType((*CreateURLResponse)(nil), "entity.CreateURLResponse")
	proto.RegisterType((*EnqueueDeleteRequest)(nil), "entity.EnqueueDeleteRequest")
	proto.RegisterType((*EnqueueDeleteResponse)(nil), "entity.EnqueueDeleteResponse")
}

func init() { proto.RegisterFile("internal/app/pkg/pb/entity.proto", fileDescriptor_f423ff2260e2d107) }

var fileDescriptor_f423ff2260e2d107 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4f, 0x6b, 0xc2, 0x40,
	0x10, 0xc5, 0x1b, 0x53, 0xad, 0x19, 0x50, 0xec, 0xa2, 0x6d, 0xd4, 0x16, 0x24, 0xa7, 0x52, 0xa8,
	0x81, 0xf6, 0xd4, 0xab, 0xf6, 0x0f, 0x42, 0xe9, 0x21, 0xe2, 0xa5, 0x17, 0x89, 0x66, 0x90, 0xd0,
	0x90, 0x5d, 0x27, 0xbb, 0x85, 0x7e, 0xda, 0x7e, 0x95, 0x92, 0x5d, 0x23, 0x46, 0xd2, 0x4b, 0x6f,
	0x3b, 0xb3, 0x8f, 0xdf, 0x7b, 0x3b, 0x3b, 0x30, 0x8a, 0x53, 0x89, 0x94, 0x86, 0x89, 0x1f, 0x0a,
	0xe1, 0x8b, 0xcf, 0x8d, 0x2f, 0x56, 0x3e, 0xa6, 0x32, 0x96, 0xdf, 0x63, 0x41, 0x5c, 0x72, 0xd6,
	0x30, 0x95, 0x77, 0x0b, 0xad, 0x57, 0x94, 0x8b, 0xe0, 0x2d, 0xc0, 0xad, 0xc2, 0x4c, 0xb2, 0x3e,
	0x34, 0x15, 0x25, 0xcb, 0x35, 0x8f, 0xd0, 0xb5, 0x46, 0xd6, 0x8d, 0x13, 0x9c, 0x29, 0x4a, 0xa6,
	0x3c, 0x42, 0x6f, 0x0a, 0xed, 0x42, 0x9b, 0x09, 0x9e, 0x66, 0xc8, 0x86, 0xe0, 0x18, 0xce, 0x32,
	0x8e, 0xb4, 0xda, 0x0e, 0x9a, 0xa6, 0x31, 0x8b, 0x58, 0x17, 0xea, 0x48, 0xc4, 0xc9, 0xad, 0x69,
	0x8c, 0x29, 0xbc, 0x05, 0x74, 0xa6, 0x84, 0xa1, 0xc4, 0x03, 0xcf, 0x21, 0x38, 0x2a, 0x43, 0x5a,
	0x2a, 0xb5, 0xc3, 0x38, 0x41, 0x33, 0x6f, 0x2c, 0x54, 0x1c, 0x31, 0x06, 0xa7, 0x3a, 0x8c, 0xa1,
	0xe8, 0x33, 0xeb, 0x80, 0xad, 0x28, 0x71, 0x6d, 0xdd, 0xca, 0x8f, 0xde, 0x0b, 0x9c, 0x1f, 0x60,
	0xff, 0x1f, 0x6f, 0x06, 0xdd, 0xe7, 0x74, 0xab, 0x50, 0xe1, 0x13, 0x26, 0x28, 0xb1, 0x88, 0xd8,
	0x85, 0x7a, 0xee, 0x9c, 0xb9, 0xd6, 0xc8, 0xce, 0xd5, 0xba, 0x28, 0x07, 0xaf, 0x95, 0x83, 0x7b,
	0x77, 0xd0, 0x3b, 0x42, 0xed, 0x62, 0xed, 0x9d, 0xad, 0x03, 0xe7, 0xfb, 0x1f, 0x0b, 0xda, 0x73,
	0xc9, 0x29, 0xdc, 0xe0, 0x1c, 0xe9, 0x2b, 0x5e, 0x23, 0x7b, 0x84, 0x86, 0x19, 0x38, 0xeb, 0x8d,
	0x77, 0xbf, 0x57, 0xfa, 0xac, 0xc1, 0xc5, 0x71, 0xdb, 0x38, 0x78, 0x27, 0x6c, 0x02, 0xce, 0x7e,
	0x1e, 0xcc, 0x2d, 0x64, 0xc7, 0x93, 0x1f, 0xf4, 0x2b, 0x6e, 0xf6, 0x8c, 0x77, 0x68, 0x95, 0x1e,
	0xc0, 0xae, 0x0a, 0x75, 0xd5, 0x88, 0x06, 0xd7, 0x7f, 0xdc, 0x16, 0xbc, 0xc9, 0xe5, 0x47, 0x6f,
	0xec, 0x57, 0x6c, 0xe6, 0xaa, 0xa1, 0x77, 0xf2, 0xe1, 0x37, 0x00, 0x00, 0xff, 0xff, 0x76, 0xe7,
	0x62, 0xb6, 0xb7, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StorageServiceClient is the client API for StorageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StorageServiceClient interface {
	GetURL(ctx context.Context, in *GetURLRequest, opts ...grpc.CallOption) (*GetURLResponse, error)
	CreateURL(ctx context.Context, in *CreateURLRequest, opts ...grpc.CallOption) (*CreateURLResponse, error)
	EnqueueDelete(ctx context.Context, in *EnqueueDeleteRequest, opts ...grpc.CallOption) (*EnqueueDeleteResponse, error)
}

type storageServiceClient struct {
	cc *grpc.ClientConn
}

func NewStorageServiceClient(cc *grpc.ClientConn) StorageServiceClient {
	return &storageServiceClient{cc}
}

func (c *storageServiceClient) GetURL(ctx context.Context, in *GetURLRequest, opts ...grpc.CallOption) (*GetURLResponse, error) {
	out := new(GetURLResponse)
	err := c.cc.Invoke(ctx, "/entity.StorageService/GetURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) CreateURL(ctx context.Context, in *CreateURLRequest, opts ...grpc.CallOption) (*CreateURLResponse, error) {
	out := new(CreateURLResponse)
	err := c.cc.Invoke(ctx, "/entity.StorageService/CreateURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) EnqueueDelete(ctx context.Context, in *EnqueueDeleteRequest, opts ...grpc.CallOption) (*EnqueueDeleteResponse, error) {
	out := new(EnqueueDeleteResponse)
	err := c.cc.Invoke(ctx, "/entity.StorageService/EnqueueDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServiceServer is the server API for StorageService service.
type StorageServiceServer interface {
	GetURL(context.Context, *GetURLRequest) (*GetURLResponse, error)
	CreateURL(context.Context, *CreateURLRequest) (*CreateURLResponse, error)
	EnqueueDelete(context.Context, *EnqueueDeleteRequest) (*EnqueueDeleteResponse, error)
}

func RegisterStorageServiceServer(s *grpc.Server, srv StorageServiceServer) {
	s.RegisterService(&_StorageService_serviceDesc, srv)
}

func _StorageService_GetURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).GetURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.StorageService/GetURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).GetURL(ctx, req.(*GetURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_CreateURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).CreateURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.StorageService/CreateURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).CreateURL(ctx, req.(*CreateURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_EnqueueDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnqueueDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).EnqueueDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.StorageService/EnqueueDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).EnqueueDelete(ctx, req.(*EnqueueDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StorageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "entity.StorageService",
	HandlerType: (*StorageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetURL",
			Handler:    _StorageService_GetURL_Handler,
		},
		{
			MethodName: "CreateURL",
			Handler:    _StorageService_CreateURL_Handler,
		},
		{
			MethodName: "EnqueueDelete",
			Handler:    _StorageService_EnqueueDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/app/pkg/pb/entity.proto",
}
