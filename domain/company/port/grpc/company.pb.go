// Code generated by protoc-gen-go. DO NOT EDIT.
// source: domain/company/port/grpc/company.proto

package company

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

type CreateCompanyRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Country              string   `protobuf:"bytes,2,opt,name=country" json:"country,omitempty"`
	State                string   `protobuf:"bytes,3,opt,name=state" json:"state,omitempty"`
	City                 string   `protobuf:"bytes,4,opt,name=city" json:"city,omitempty"`
	Street               string   `protobuf:"bytes,5,opt,name=street" json:"street,omitempty"`
	Number               string   `protobuf:"bytes,6,opt,name=number" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCompanyRequest) Reset()         { *m = CreateCompanyRequest{} }
func (m *CreateCompanyRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCompanyRequest) ProtoMessage()    {}
func (*CreateCompanyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_company_b95c3e0c02594e2e, []int{0}
}
func (m *CreateCompanyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCompanyRequest.Unmarshal(m, b)
}
func (m *CreateCompanyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCompanyRequest.Marshal(b, m, deterministic)
}
func (dst *CreateCompanyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCompanyRequest.Merge(dst, src)
}
func (m *CreateCompanyRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCompanyRequest.Size(m)
}
func (m *CreateCompanyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCompanyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCompanyRequest proto.InternalMessageInfo

func (m *CreateCompanyRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateCompanyRequest) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *CreateCompanyRequest) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *CreateCompanyRequest) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *CreateCompanyRequest) GetStreet() string {
	if m != nil {
		return m.Street
	}
	return ""
}

func (m *CreateCompanyRequest) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

type CreateCompanyResponse struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCompanyResponse) Reset()         { *m = CreateCompanyResponse{} }
func (m *CreateCompanyResponse) String() string { return proto.CompactTextString(m) }
func (*CreateCompanyResponse) ProtoMessage()    {}
func (*CreateCompanyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_company_b95c3e0c02594e2e, []int{1}
}
func (m *CreateCompanyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCompanyResponse.Unmarshal(m, b)
}
func (m *CreateCompanyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCompanyResponse.Marshal(b, m, deterministic)
}
func (dst *CreateCompanyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCompanyResponse.Merge(dst, src)
}
func (m *CreateCompanyResponse) XXX_Size() int {
	return xxx_messageInfo_CreateCompanyResponse.Size(m)
}
func (m *CreateCompanyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCompanyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCompanyResponse proto.InternalMessageInfo

func (m *CreateCompanyResponse) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateCompanyRequest)(nil), "company.CreateCompanyRequest")
	proto.RegisterType((*CreateCompanyResponse)(nil), "company.CreateCompanyResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CompanyService service

type CompanyServiceClient interface {
	CreateCompany(ctx context.Context, in *CreateCompanyRequest, opts ...grpc.CallOption) (*CreateCompanyResponse, error)
}

type companyServiceClient struct {
	cc *grpc.ClientConn
}

func NewCompanyServiceClient(cc *grpc.ClientConn) CompanyServiceClient {
	return &companyServiceClient{cc}
}

func (c *companyServiceClient) CreateCompany(ctx context.Context, in *CreateCompanyRequest, opts ...grpc.CallOption) (*CreateCompanyResponse, error) {
	out := new(CreateCompanyResponse)
	err := grpc.Invoke(ctx, "/company.CompanyService/CreateCompany", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CompanyService service

type CompanyServiceServer interface {
	CreateCompany(context.Context, *CreateCompanyRequest) (*CreateCompanyResponse, error)
}

func RegisterCompanyServiceServer(s *grpc.Server, srv CompanyServiceServer) {
	s.RegisterService(&_CompanyService_serviceDesc, srv)
}

func _CompanyService_CreateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).CreateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/company.CompanyService/CreateCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).CreateCompany(ctx, req.(*CreateCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CompanyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "company.CompanyService",
	HandlerType: (*CompanyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCompany",
			Handler:    _CompanyService_CreateCompany_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "domain/company/port/grpc/company.proto",
}

func init() {
	proto.RegisterFile("domain/company/port/grpc/company.proto", fileDescriptor_company_b95c3e0c02594e2e)
}

var fileDescriptor_company_b95c3e0c02594e2e = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x51, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xad, 0xee, 0x76, 0x71, 0x40, 0x1f, 0x86, 0x55, 0x82, 0xa0, 0x48, 0x1f, 0x44, 0x10,
	0xb6, 0xa0, 0x47, 0xd8, 0x0b, 0xc8, 0x7a, 0x82, 0x34, 0x3b, 0x48, 0x1f, 0x9a, 0xc4, 0xc9, 0x44,
	0xe8, 0x65, 0x3c, 0xab, 0x24, 0x69, 0x1e, 0x14, 0x7d, 0xfb, 0xbf, 0xaf, 0xcd, 0x9f, 0xc9, 0xc0,
	0xc3, 0xd1, 0x4d, 0x7a, 0xb4, 0xbd, 0x71, 0x93, 0xd7, 0x76, 0xee, 0xbd, 0x63, 0xe9, 0xdf, 0xd9,
	0x9b, 0x6a, 0x76, 0x9e, 0x9d, 0x38, 0xdc, 0x2c, 0xd8, 0x7d, 0x35, 0xb0, 0xdd, 0x33, 0x69, 0xa1,
	0x7d, 0x31, 0x07, 0xfa, 0x88, 0x14, 0x04, 0x11, 0x56, 0x56, 0x4f, 0xa4, 0x9a, 0xfb, 0xe6, 0xf1,
	0xfc, 0x90, 0x33, 0x2a, 0xd8, 0x18, 0x17, 0xad, 0xf0, 0xac, 0x4e, 0xb3, 0xae, 0x88, 0x5b, 0x58,
	0x07, 0xd1, 0x42, 0xea, 0x2c, 0xfb, 0x02, 0xa9, 0xc3, 0x8c, 0x32, 0xab, 0x55, 0xe9, 0x48, 0x19,
	0xaf, 0xa1, 0x0d, 0xc2, 0x44, 0xa2, 0xd6, 0xd9, 0x2e, 0x94, 0xbc, 0x8d, 0xd3, 0x40, 0xac, 0xda,
	0xe2, 0x0b, 0x75, 0x4f, 0x70, 0xf5, 0x6b, 0xbe, 0xe0, 0x9d, 0x0d, 0xb9, 0x3c, 0xc6, 0xf1, 0x58,
	0x07, 0x4c, 0xf9, 0x79, 0x80, 0xcb, 0xe5, 0xb7, 0x37, 0xe2, 0xcf, 0xd1, 0x10, 0xbe, 0xc2, 0xc5,
	0x8f, 0xe3, 0x78, 0xbb, 0xab, 0x9b, 0xf8, 0xeb, 0xd9, 0x37, 0x77, 0xff, 0x7d, 0x2e, 0xb7, 0x76,
	0x27, 0x43, 0x9b, 0x37, 0xf8, 0xf2, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x61, 0xa8, 0x5a, 0x6b,
	0x01, 0x00, 0x00,
}
