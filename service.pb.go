// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package permissionproto

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

type RequestHeader struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestHeader) Reset()         { *m = RequestHeader{} }
func (m *RequestHeader) String() string { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()    {}
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *RequestHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestHeader.Unmarshal(m, b)
}
func (m *RequestHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestHeader.Marshal(b, m, deterministic)
}
func (m *RequestHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestHeader.Merge(m, src)
}
func (m *RequestHeader) XXX_Size() int {
	return xxx_messageInfo_RequestHeader.Size(m)
}
func (m *RequestHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestHeader.DiscardUnknown(m)
}

var xxx_messageInfo_RequestHeader proto.InternalMessageInfo

func (m *RequestHeader) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type RequestField struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	AddressStreet        string   `protobuf:"bytes,4,opt,name=addressStreet,proto3" json:"addressStreet,omitempty"`
	AddressCity          string   `protobuf:"bytes,5,opt,name=addressCity,proto3" json:"addressCity,omitempty"`
	AddressPostCode      string   `protobuf:"bytes,6,opt,name=addressPostCode,proto3" json:"addressPostCode,omitempty"`
	AddressState         string   `protobuf:"bytes,7,opt,name=addressState,proto3" json:"addressState,omitempty"`
	AddressCountry       string   `protobuf:"bytes,8,opt,name=addressCountry,proto3" json:"addressCountry,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestField) Reset()         { *m = RequestField{} }
func (m *RequestField) String() string { return proto.CompactTextString(m) }
func (*RequestField) ProtoMessage()    {}
func (*RequestField) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *RequestField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestField.Unmarshal(m, b)
}
func (m *RequestField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestField.Marshal(b, m, deterministic)
}
func (m *RequestField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestField.Merge(m, src)
}
func (m *RequestField) XXX_Size() int {
	return xxx_messageInfo_RequestField.Size(m)
}
func (m *RequestField) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestField.DiscardUnknown(m)
}

var xxx_messageInfo_RequestField proto.InternalMessageInfo

func (m *RequestField) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *RequestField) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RequestField) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RequestField) GetAddressStreet() string {
	if m != nil {
		return m.AddressStreet
	}
	return ""
}

func (m *RequestField) GetAddressCity() string {
	if m != nil {
		return m.AddressCity
	}
	return ""
}

func (m *RequestField) GetAddressPostCode() string {
	if m != nil {
		return m.AddressPostCode
	}
	return ""
}

func (m *RequestField) GetAddressState() string {
	if m != nil {
		return m.AddressState
	}
	return ""
}

func (m *RequestField) GetAddressCountry() string {
	if m != nil {
		return m.AddressCountry
	}
	return ""
}

type GetRequest struct {
	Header               *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Id                   int32          `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *GetRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CreateRequest struct {
	Header               *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Field                *RequestField  `protobuf:"bytes,2,opt,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CreateRequest) GetField() *RequestField {
	if m != nil {
		return m.Field
	}
	return nil
}

type UpdateRequest struct {
	Header               *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Id                   int32          `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Field                *RequestField  `protobuf:"bytes,3,opt,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *UpdateRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateRequest) GetField() *RequestField {
	if m != nil {
		return m.Field
	}
	return nil
}

type ActionRequest struct {
	Header               *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Id                   int32          `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string         `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ActionRequest) Reset()         { *m = ActionRequest{} }
func (m *ActionRequest) String() string { return proto.CompactTextString(m) }
func (*ActionRequest) ProtoMessage()    {}
func (*ActionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *ActionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionRequest.Unmarshal(m, b)
}
func (m *ActionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionRequest.Marshal(b, m, deterministic)
}
func (m *ActionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionRequest.Merge(m, src)
}
func (m *ActionRequest) XXX_Size() int {
	return xxx_messageInfo_ActionRequest.Size(m)
}
func (m *ActionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ActionRequest proto.InternalMessageInfo

func (m *ActionRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ActionRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ActionRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DataResponse struct {
	StatusCode           int32              `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Message              string             `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 *DataResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DataResponse) Reset()         { *m = DataResponse{} }
func (m *DataResponse) String() string { return proto.CompactTextString(m) }
func (*DataResponse) ProtoMessage()    {}
func (*DataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6}
}

func (m *DataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataResponse.Unmarshal(m, b)
}
func (m *DataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataResponse.Marshal(b, m, deterministic)
}
func (m *DataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataResponse.Merge(m, src)
}
func (m *DataResponse) XXX_Size() int {
	return xxx_messageInfo_DataResponse.Size(m)
}
func (m *DataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DataResponse proto.InternalMessageInfo

func (m *DataResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *DataResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DataResponse) GetData() *DataResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type DataResponse_Data struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Code                 string   `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	AddressStreet        string   `protobuf:"bytes,5,opt,name=addressStreet,proto3" json:"addressStreet,omitempty"`
	AddressCity          string   `protobuf:"bytes,6,opt,name=addressCity,proto3" json:"addressCity,omitempty"`
	AddressPostCode      string   `protobuf:"bytes,7,opt,name=addressPostCode,proto3" json:"addressPostCode,omitempty"`
	AddressState         string   `protobuf:"bytes,8,opt,name=addressState,proto3" json:"addressState,omitempty"`
	AddressCountry       string   `protobuf:"bytes,9,opt,name=addressCountry,proto3" json:"addressCountry,omitempty"`
	CreatedDate          string   `protobuf:"bytes,10,opt,name=createdDate,proto3" json:"createdDate,omitempty"`
	CreatedBy            string   `protobuf:"bytes,11,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	LastModifiedDate     string   `protobuf:"bytes,12,opt,name=lastModifiedDate,proto3" json:"lastModifiedDate,omitempty"`
	LastModifiedBy       string   `protobuf:"bytes,13,opt,name=lastModifiedBy,proto3" json:"lastModifiedBy,omitempty"`
	RecordState          string   `protobuf:"bytes,14,opt,name=recordState,proto3" json:"recordState,omitempty"`
	ReferencedBy         int64    `protobuf:"varint,36,opt,name=referencedBy,proto3" json:"referencedBy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataResponse_Data) Reset()         { *m = DataResponse_Data{} }
func (m *DataResponse_Data) String() string { return proto.CompactTextString(m) }
func (*DataResponse_Data) ProtoMessage()    {}
func (*DataResponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6, 0}
}

func (m *DataResponse_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataResponse_Data.Unmarshal(m, b)
}
func (m *DataResponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataResponse_Data.Marshal(b, m, deterministic)
}
func (m *DataResponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataResponse_Data.Merge(m, src)
}
func (m *DataResponse_Data) XXX_Size() int {
	return xxx_messageInfo_DataResponse_Data.Size(m)
}
func (m *DataResponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_DataResponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_DataResponse_Data proto.InternalMessageInfo

func (m *DataResponse_Data) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DataResponse_Data) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DataResponse_Data) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *DataResponse_Data) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *DataResponse_Data) GetAddressStreet() string {
	if m != nil {
		return m.AddressStreet
	}
	return ""
}

func (m *DataResponse_Data) GetAddressCity() string {
	if m != nil {
		return m.AddressCity
	}
	return ""
}

func (m *DataResponse_Data) GetAddressPostCode() string {
	if m != nil {
		return m.AddressPostCode
	}
	return ""
}

func (m *DataResponse_Data) GetAddressState() string {
	if m != nil {
		return m.AddressState
	}
	return ""
}

func (m *DataResponse_Data) GetAddressCountry() string {
	if m != nil {
		return m.AddressCountry
	}
	return ""
}

func (m *DataResponse_Data) GetCreatedDate() string {
	if m != nil {
		return m.CreatedDate
	}
	return ""
}

func (m *DataResponse_Data) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *DataResponse_Data) GetLastModifiedDate() string {
	if m != nil {
		return m.LastModifiedDate
	}
	return ""
}

func (m *DataResponse_Data) GetLastModifiedBy() string {
	if m != nil {
		return m.LastModifiedBy
	}
	return ""
}

func (m *DataResponse_Data) GetRecordState() string {
	if m != nil {
		return m.RecordState
	}
	return ""
}

func (m *DataResponse_Data) GetReferencedBy() int64 {
	if m != nil {
		return m.ReferencedBy
	}
	return 0
}

func init() {
	proto.RegisterType((*RequestHeader)(nil), "permissionproto.RequestHeader")
	proto.RegisterType((*RequestField)(nil), "permissionproto.RequestField")
	proto.RegisterType((*GetRequest)(nil), "permissionproto.GetRequest")
	proto.RegisterType((*CreateRequest)(nil), "permissionproto.CreateRequest")
	proto.RegisterType((*UpdateRequest)(nil), "permissionproto.UpdateRequest")
	proto.RegisterType((*ActionRequest)(nil), "permissionproto.ActionRequest")
	proto.RegisterType((*DataResponse)(nil), "permissionproto.DataResponse")
	proto.RegisterType((*DataResponse_Data)(nil), "permissionproto.DataResponse.Data")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 576 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x55, 0xe2, 0x8f, 0x34, 0x93, 0x38, 0x45, 0x2b, 0x0e, 0x56, 0xa1, 0x55, 0x64, 0x15, 0x14,
	0x71, 0xc8, 0x21, 0x95, 0x7a, 0x6f, 0x53, 0x11, 0x2e, 0x48, 0xc8, 0x80, 0x38, 0x2f, 0xde, 0x09,
	0x5d, 0xb5, 0xf1, 0x86, 0xdd, 0x0d, 0x52, 0x24, 0x7e, 0x02, 0x67, 0xfe, 0x03, 0xff, 0x10, 0x89,
	0x0b, 0xda, 0x5d, 0x9b, 0xd8, 0x49, 0xa0, 0x46, 0xed, 0xcd, 0xf3, 0xfc, 0xfc, 0xfc, 0x76, 0x76,
	0xde, 0x40, 0xa4, 0x50, 0x7e, 0xe1, 0x19, 0x8e, 0x97, 0x52, 0x68, 0x41, 0x0e, 0x97, 0x28, 0x17,
	0x5c, 0x29, 0x2e, 0x72, 0x0b, 0x24, 0xcf, 0x20, 0x4a, 0xf1, 0xf3, 0x0a, 0x95, 0x7e, 0x85, 0x94,
	0xa1, 0x24, 0x8f, 0x21, 0xd0, 0xe2, 0x06, 0xf3, 0xb8, 0x35, 0x6c, 0x8d, 0xba, 0xa9, 0x2b, 0x92,
	0xef, 0x6d, 0xe8, 0x17, 0xbc, 0x97, 0x1c, 0x6f, 0x19, 0x21, 0xe0, 0x67, 0x82, 0x61, 0xc1, 0xb2,
	0xcf, 0x06, 0xcb, 0xe9, 0x02, 0xe3, 0xb6, 0xc3, 0xcc, 0x33, 0x19, 0x42, 0x8f, 0xa1, 0xca, 0x24,
	0x5f, 0x6a, 0x2e, 0xf2, 0xd8, 0xb3, 0xaf, 0xaa, 0x10, 0x39, 0x85, 0x88, 0x32, 0x26, 0x51, 0xa9,
	0xb7, 0x5a, 0x22, 0xea, 0xd8, 0xb7, 0x9c, 0x3a, 0x68, 0x74, 0x0a, 0x60, 0xca, 0xf5, 0x3a, 0x0e,
	0x9c, 0x4e, 0x05, 0x22, 0x23, 0x38, 0x2c, 0xca, 0x37, 0x42, 0xe9, 0xa9, 0x31, 0x17, 0x5a, 0xd6,
	0x36, 0x4c, 0x12, 0xe8, 0xff, 0x11, 0xa7, 0x1a, 0xe3, 0x8e, 0xa5, 0xd5, 0x30, 0xf2, 0x1c, 0x06,
	0xa5, 0xb8, 0x58, 0xe5, 0x5a, 0xae, 0xe3, 0x03, 0xcb, 0xda, 0x42, 0x93, 0x77, 0x00, 0x33, 0xd4,
	0x45, 0x6b, 0xc8, 0x39, 0x84, 0xd7, 0xb6, 0x8d, 0xb6, 0x2f, 0xbd, 0xc9, 0xc9, 0x78, 0xab, 0xdf,
	0xe3, 0x5a, 0xb3, 0xd3, 0x82, 0x4d, 0x06, 0xd0, 0xe6, 0xcc, 0xf6, 0x2d, 0x48, 0xdb, 0x9c, 0x25,
	0x5f, 0x21, 0x9a, 0x4a, 0xa4, 0x1a, 0xef, 0x2b, 0x7c, 0x06, 0xc1, 0xdc, 0xdc, 0x97, 0xd5, 0xee,
	0x4d, 0x8e, 0xff, 0xf6, 0x99, 0xbd, 0xd4, 0xd4, 0x71, 0x93, 0x6f, 0x2d, 0x88, 0xde, 0x2f, 0xd9,
	0x03, 0xfc, 0x7e, 0xeb, 0x5c, 0x1b, 0x3b, 0xde, 0x7f, 0xd8, 0xb9, 0x81, 0xe8, 0x22, 0x33, 0xa3,
	0xf2, 0xd0, 0x6e, 0xca, 0x79, 0xf5, 0x36, 0xf3, 0x9a, 0xfc, 0xf2, 0xa1, 0x7f, 0x45, 0x35, 0x4d,
	0x51, 0x2d, 0x45, 0xae, 0x90, 0x9c, 0x00, 0x28, 0x4d, 0xf5, 0x4a, 0x4d, 0xcb, 0x71, 0x0f, 0xd2,
	0x0a, 0x42, 0x62, 0xe8, 0x2c, 0x50, 0x29, 0xfa, 0xa9, 0x9c, 0xfb, 0xb2, 0x24, 0xe7, 0xe0, 0x33,
	0xaa, 0x69, 0x71, 0xd6, 0x64, 0xc7, 0x64, 0xf5, 0x37, 0xae, 0xb0, 0xfc, 0xa3, 0x9f, 0x1e, 0xf8,
	0xa6, 0x2c, 0xfc, 0xba, 0x84, 0x55, 0xfd, 0x56, 0xf3, 0x55, 0xe6, 0xd0, 0xab, 0xe4, 0x70, 0x2b,
	0x73, 0x7e, 0x83, 0xcc, 0x05, 0x0d, 0x32, 0x17, 0x36, 0xca, 0x5c, 0xa7, 0x59, 0xe6, 0x0e, 0x1a,
	0x65, 0xae, 0xbb, 0x2f, 0x73, 0xc6, 0x57, 0x66, 0xd3, 0xc1, 0xae, 0x8c, 0x14, 0x38, 0x5f, 0x15,
	0x88, 0x3c, 0x85, 0x6e, 0x51, 0x5e, 0xae, 0xe3, 0x9e, 0x7d, 0xbf, 0x01, 0xc8, 0x0b, 0x78, 0x74,
	0x4b, 0x95, 0x7e, 0x2d, 0x18, 0x9f, 0xf3, 0x42, 0xa4, 0x6f, 0x49, 0x3b, 0xb8, 0xf1, 0x54, 0xc5,
	0x2e, 0xd7, 0x71, 0xe4, 0x3c, 0xd5, 0x51, 0xe3, 0x49, 0x62, 0x26, 0x24, 0x73, 0xc7, 0x1b, 0x38,
	0x4f, 0x15, 0xc8, 0x74, 0x40, 0xe2, 0x1c, 0x25, 0xe6, 0x99, 0xd5, 0x39, 0x1d, 0xb6, 0x46, 0x5e,
	0x5a, 0xc3, 0x26, 0x3f, 0xda, 0xd0, 0xfd, 0x40, 0x25, 0x5e, 0x8b, 0x95, 0x42, 0x32, 0x83, 0xd0,
	0x6d, 0x01, 0xb2, 0x3b, 0xe1, 0xb5, 0xf5, 0x70, 0x74, 0xfc, 0xcf, 0xe1, 0x22, 0x17, 0xe0, 0xcd,
	0x50, 0x93, 0x27, 0x3b, 0xac, 0xcd, 0xea, 0xba, 0x4b, 0x62, 0x06, 0xa1, 0x5b, 0x09, 0x7b, 0xbc,
	0xd4, 0x76, 0x45, 0x03, 0x21, 0x97, 0xe6, 0x3d, 0x42, 0xb5, 0x98, 0xdf, 0x21, 0xf4, 0x31, 0xb4,
	0xd8, 0xd9, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x46, 0x06, 0x9f, 0xd9, 0xe2, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WarehouseClient is the client API for Warehouse service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WarehouseClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*DataResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*DataResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*DataResponse, error)
	Action(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*DataResponse, error)
}

type warehouseClient struct {
	cc *grpc.ClientConn
}

func NewWarehouseClient(cc *grpc.ClientConn) WarehouseClient {
	return &warehouseClient{cc}
}

func (c *warehouseClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*DataResponse, error) {
	out := new(DataResponse)
	err := c.cc.Invoke(ctx, "/permissionproto.Warehouse/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*DataResponse, error) {
	out := new(DataResponse)
	err := c.cc.Invoke(ctx, "/permissionproto.Warehouse/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*DataResponse, error) {
	out := new(DataResponse)
	err := c.cc.Invoke(ctx, "/permissionproto.Warehouse/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseClient) Action(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*DataResponse, error) {
	out := new(DataResponse)
	err := c.cc.Invoke(ctx, "/permissionproto.Warehouse/Action", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WarehouseServer is the server API for Warehouse service.
type WarehouseServer interface {
	Create(context.Context, *CreateRequest) (*DataResponse, error)
	Get(context.Context, *GetRequest) (*DataResponse, error)
	Update(context.Context, *UpdateRequest) (*DataResponse, error)
	Action(context.Context, *ActionRequest) (*DataResponse, error)
}

// UnimplementedWarehouseServer can be embedded to have forward compatible implementations.
type UnimplementedWarehouseServer struct {
}

func (*UnimplementedWarehouseServer) Create(ctx context.Context, req *CreateRequest) (*DataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedWarehouseServer) Get(ctx context.Context, req *GetRequest) (*DataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedWarehouseServer) Update(ctx context.Context, req *UpdateRequest) (*DataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedWarehouseServer) Action(ctx context.Context, req *ActionRequest) (*DataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Action not implemented")
}

func RegisterWarehouseServer(s *grpc.Server, srv WarehouseServer) {
	s.RegisterService(&_Warehouse_serviceDesc, srv)
}

func _Warehouse_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/permissionproto.Warehouse/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Warehouse_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/permissionproto.Warehouse/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Warehouse_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/permissionproto.Warehouse/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Warehouse_Action_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServer).Action(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/permissionproto.Warehouse/Action",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServer).Action(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Warehouse_serviceDesc = grpc.ServiceDesc{
	ServiceName: "permissionproto.Warehouse",
	HandlerType: (*WarehouseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Warehouse_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Warehouse_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Warehouse_Update_Handler,
		},
		{
			MethodName: "Action",
			Handler:    _Warehouse_Action_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
