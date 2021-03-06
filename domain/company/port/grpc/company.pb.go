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
	Zip                  string   `protobuf:"bytes,6,opt,name=zip" json:"zip,omitempty"`
	Number               string   `protobuf:"bytes,7,opt,name=number" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCompanyRequest) Reset()         { *m = CreateCompanyRequest{} }
func (m *CreateCompanyRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCompanyRequest) ProtoMessage()    {}
func (*CreateCompanyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_company_8ec355be8bbbd6dc, []int{0}
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

func (m *CreateCompanyRequest) GetZip() string {
	if m != nil {
		return m.Zip
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
	return fileDescriptor_company_8ec355be8bbbd6dc, []int{1}
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
	proto.RegisterFile("domain/company/port/grpc/company.proto", fileDescriptor_company_8ec355be8bbbd6dc)
}

var fileDescriptor_company_8ec355be8bbbd6dc = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x51, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xad, 0xbb, 0xdb, 0xe2, 0x80, 0x22, 0xc3, 0x2a, 0x41, 0x50, 0xa4, 0x0f, 0x22, 0x08,
	0x5b, 0xd0, 0x23, 0xec, 0x05, 0x64, 0x3d, 0x41, 0x9a, 0x0e, 0x92, 0x87, 0x26, 0x31, 0x99, 0x08,
	0xf5, 0x54, 0x1e, 0x51, 0x92, 0xb4, 0x0f, 0x8a, 0xfb, 0xf6, 0xff, 0x5f, 0xfe, 0xcc, 0x0c, 0x3f,
	0x3c, 0x0c, 0x76, 0x94, 0xda, 0x74, 0xca, 0x8e, 0x4e, 0x9a, 0xa9, 0x73, 0xd6, 0x73, 0xf7, 0xee,
	0x9d, 0x5a, 0xc8, 0xce, 0x79, 0xcb, 0x16, 0x9b, 0xd9, 0xb6, 0xdf, 0x15, 0x6c, 0xf7, 0x9e, 0x24,
	0xd3, 0xbe, 0x90, 0x03, 0x7d, 0x44, 0x0a, 0x8c, 0x08, 0x6b, 0x23, 0x47, 0x12, 0xd5, 0x7d, 0xf5,
	0x78, 0x76, 0xc8, 0x1a, 0x05, 0x34, 0xca, 0x46, 0xc3, 0x7e, 0x12, 0xa7, 0x19, 0x2f, 0x16, 0xb7,
	0xb0, 0x09, 0x2c, 0x99, 0xc4, 0x2a, 0xf3, 0x62, 0xd2, 0x0c, 0xa5, 0x79, 0x12, 0xeb, 0x32, 0x23,
	0x69, 0xbc, 0x86, 0x3a, 0xb0, 0x27, 0x62, 0xb1, 0xc9, 0x74, 0x76, 0x78, 0x09, 0xab, 0x2f, 0xed,
	0x44, 0x9d, 0x61, 0x92, 0x29, 0x69, 0xe2, 0xd8, 0x93, 0x17, 0x4d, 0x49, 0x16, 0xd7, 0x3e, 0xc1,
	0xd5, 0x9f, 0x8b, 0x83, 0xb3, 0x26, 0xe4, 0x75, 0x31, 0xea, 0x61, 0x39, 0x39, 0xe9, 0xe7, 0x1e,
	0x2e, 0xe6, 0xd8, 0x1b, 0xf9, 0x4f, 0xad, 0x08, 0x5f, 0xe1, 0xfc, 0xd7, 0x77, 0xbc, 0xdd, 0x2d,
	0xdd, 0xfc, 0x57, 0xc4, 0xcd, 0xdd, 0xb1, 0xe7, 0xb2, 0xb5, 0x3d, 0xe9, 0xeb, 0xdc, 0xe9, 0xcb,
	0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xce, 0x01, 0xc6, 0x7d, 0x01, 0x00, 0x00,
}
