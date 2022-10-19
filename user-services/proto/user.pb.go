// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Address              *Address `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Account              *Account `protobuf:"bytes,5,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *User) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type Address struct {
	Street               string   `protobuf:"bytes,1,opt,name=street,proto3" json:"street,omitempty"`
	ZipCode              string   `protobuf:"bytes,2,opt,name=zipCode,proto3" json:"zipCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetStreet() string {
	if m != nil {
		return m.Street
	}
	return ""
}

func (m *Address) GetZipCode() string {
	if m != nil {
		return m.ZipCode
	}
	return ""
}

type Account struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Account) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Token struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type TokenValidation struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenValidation) Reset()         { *m = TokenValidation{} }
func (m *TokenValidation) String() string { return proto.CompactTextString(m) }
func (*TokenValidation) ProtoMessage()    {}
func (*TokenValidation) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *TokenValidation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenValidation.Unmarshal(m, b)
}
func (m *TokenValidation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenValidation.Marshal(b, m, deterministic)
}
func (m *TokenValidation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenValidation.Merge(m, src)
}
func (m *TokenValidation) XXX_Size() int {
	return xxx_messageInfo_TokenValidation.Size(m)
}
func (m *TokenValidation) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenValidation.DiscardUnknown(m)
}

var xxx_messageInfo_TokenValidation proto.InternalMessageInfo

func (m *TokenValidation) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *TokenValidation) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type UserId struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserId) Reset()         { *m = UserId{} }
func (m *UserId) String() string { return proto.CompactTextString(m) }
func (*UserId) ProtoMessage()    {}
func (*UserId) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *UserId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserId.Unmarshal(m, b)
}
func (m *UserId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserId.Marshal(b, m, deterministic)
}
func (m *UserId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserId.Merge(m, src)
}
func (m *UserId) XXX_Size() int {
	return xxx_messageInfo_UserId.Size(m)
}
func (m *UserId) XXX_DiscardUnknown() {
	xxx_messageInfo_UserId.DiscardUnknown(m)
}

var xxx_messageInfo_UserId proto.InternalMessageInfo

func (m *UserId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*Address)(nil), "user.Address")
	proto.RegisterType((*Account)(nil), "user.Account")
	proto.RegisterType((*Token)(nil), "user.Token")
	proto.RegisterType((*TokenValidation)(nil), "user.TokenValidation")
	proto.RegisterType((*UserId)(nil), "user.UserId")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0x4d, 0x4b, 0x3b, 0x31,
	0x10, 0xc6, 0xbb, 0xfd, 0x6f, 0xbb, 0xfd, 0x4f, 0xab, 0xc2, 0xa0, 0x12, 0x0a, 0x42, 0x09, 0x88,
	0x3d, 0xf5, 0x50, 0xbd, 0x88, 0x27, 0xe9, 0x41, 0x7a, 0x5d, 0x5f, 0xee, 0xb1, 0x19, 0x24, 0xda,
	0x6e, 0x96, 0x24, 0xab, 0xe0, 0xc7, 0xf0, 0x13, 0x4b, 0x5e, 0xb6, 0xab, 0xde, 0xe6, 0x37, 0xcf,
	0x3c, 0x9b, 0x27, 0xb3, 0x01, 0x68, 0x2c, 0x99, 0x45, 0x6d, 0xb4, 0xd3, 0x98, 0xfb, 0x9a, 0x7f,
	0x65, 0x90, 0x3f, 0x5a, 0x32, 0x78, 0x08, 0x7d, 0x25, 0x59, 0x36, 0xcb, 0xe6, 0xff, 0xcb, 0xbe,
	0x92, 0x88, 0x90, 0x57, 0x62, 0x47, 0xac, 0x1f, 0x3a, 0xa1, 0xc6, 0x63, 0x18, 0xd0, 0x4e, 0xa8,
	0x2d, 0xfb, 0x17, 0x9a, 0x11, 0xf0, 0x02, 0x0a, 0x21, 0xa5, 0x21, 0x6b, 0x59, 0x3e, 0xcb, 0xe6,
	0xe3, 0xe5, 0xc1, 0x22, 0x1c, 0x73, 0x1b, 0x9b, 0x65, 0xab, 0x86, 0xc1, 0xcd, 0x46, 0x37, 0x95,
	0x63, 0x83, 0x5f, 0x83, 0xb1, 0x59, 0xb6, 0x2a, 0xbf, 0x81, 0x22, 0x99, 0xf1, 0x14, 0x86, 0xd6,
	0x19, 0x22, 0x97, 0xa2, 0x25, 0x42, 0x06, 0xc5, 0xa7, 0xaa, 0x57, 0x5a, 0xb6, 0x09, 0x5b, 0xe4,
	0xd7, 0x50, 0xa4, 0x0f, 0xfa, 0x3b, 0xf8, 0x03, 0x92, 0x35, 0xd4, 0x38, 0x85, 0x51, 0x2d, 0xac,
	0xfd, 0xd0, 0x46, 0x26, 0xe7, 0x9e, 0xf9, 0x19, 0x0c, 0x1e, 0xf4, 0x1b, 0x55, 0xfe, 0xa2, 0xce,
	0x17, 0xc9, 0x19, 0x81, 0xaf, 0xe0, 0x28, 0xc8, 0x4f, 0x62, 0xab, 0xa4, 0x70, 0x4a, 0x57, 0x31,
	0x9e, 0x70, 0x8d, 0x0d, 0x93, 0xa3, 0x32, 0x91, 0x8f, 0xb7, 0x23, 0x6b, 0xc5, 0xcb, 0x3e, 0x5e,
	0x42, 0xce, 0x60, 0xe8, 0xf7, 0xbd, 0x96, 0x7f, 0x37, 0xbe, 0x7c, 0x85, 0xb1, 0x57, 0xee, 0xc9,
	0xbc, 0xab, 0x0d, 0xe1, 0x39, 0x14, 0x77, 0xe4, 0xc2, 0xbf, 0x99, 0xc4, 0x3d, 0x45, 0xdf, 0x14,
	0x3a, 0xe2, 0x3d, 0xbc, 0x82, 0xc9, 0xda, 0x76, 0xb1, 0x70, 0x1c, 0xd5, 0xd0, 0x99, 0x9e, 0xfc,
	0x80, 0x2e, 0x35, 0xef, 0x3d, 0x0f, 0xc3, 0x1b, 0xb8, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x1d,
	0x03, 0x5a, 0x23, 0x11, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUser(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*User, error)
	IsTokenValid(ctx context.Context, in *Token, opts ...grpc.CallOption) (*TokenValidation, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUser(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) IsTokenValid(ctx context.Context, in *Token, opts ...grpc.CallOption) (*TokenValidation, error) {
	out := new(TokenValidation)
	err := c.cc.Invoke(ctx, "/user.UserService/IsTokenValid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	GetUser(context.Context, *UserId) (*User, error)
	IsTokenValid(context.Context, *Token) (*TokenValidation, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) GetUser(ctx context.Context, req *UserId) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUserServiceServer) IsTokenValid(ctx context.Context, req *Token) (*TokenValidation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsTokenValid not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_IsTokenValid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).IsTokenValid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/IsTokenValid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).IsTokenValid(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "IsTokenValid",
			Handler:    _UserService_IsTokenValid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
