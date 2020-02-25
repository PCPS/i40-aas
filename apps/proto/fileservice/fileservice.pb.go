// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fileservice.proto

package fileservice

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type UploadStatusCode int32

const (
	UploadStatusCode_Unknown UploadStatusCode = 0
	UploadStatusCode_Ok      UploadStatusCode = 1
	UploadStatusCode_Failed  UploadStatusCode = 2
)

var UploadStatusCode_name = map[int32]string{
	0: "Unknown",
	1: "Ok",
	2: "Failed",
}

var UploadStatusCode_value = map[string]int32{
	"Unknown": 0,
	"Ok":      1,
	"Failed":  2,
}

func (x UploadStatusCode) String() string {
	return proto.EnumName(UploadStatusCode_name, int32(x))
}

func (UploadStatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_27a5d7178d16d994, []int{0}
}

type DownloadRequest struct {
	FileName             string   `protobuf:"bytes,1,opt,name=FileName,proto3" json:"FileName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadRequest) Reset()         { *m = DownloadRequest{} }
func (m *DownloadRequest) String() string { return proto.CompactTextString(m) }
func (*DownloadRequest) ProtoMessage()    {}
func (*DownloadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_27a5d7178d16d994, []int{0}
}

func (m *DownloadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadRequest.Unmarshal(m, b)
}
func (m *DownloadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadRequest.Marshal(b, m, deterministic)
}
func (m *DownloadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadRequest.Merge(m, src)
}
func (m *DownloadRequest) XXX_Size() int {
	return xxx_messageInfo_DownloadRequest.Size(m)
}
func (m *DownloadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadRequest proto.InternalMessageInfo

func (m *DownloadRequest) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

type Chunk struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
	FileName             string   `protobuf:"bytes,2,opt,name=FileName,proto3" json:"FileName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chunk) Reset()         { *m = Chunk{} }
func (m *Chunk) String() string { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()    {}
func (*Chunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_27a5d7178d16d994, []int{1}
}

func (m *Chunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chunk.Unmarshal(m, b)
}
func (m *Chunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chunk.Marshal(b, m, deterministic)
}
func (m *Chunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chunk.Merge(m, src)
}
func (m *Chunk) XXX_Size() int {
	return xxx_messageInfo_Chunk.Size(m)
}
func (m *Chunk) XXX_DiscardUnknown() {
	xxx_messageInfo_Chunk.DiscardUnknown(m)
}

var xxx_messageInfo_Chunk proto.InternalMessageInfo

func (m *Chunk) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Chunk) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

type UploadStatus struct {
	Message              string           `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Code                 UploadStatusCode `protobuf:"varint,2,opt,name=Code,proto3,enum=fileservice.UploadStatusCode" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UploadStatus) Reset()         { *m = UploadStatus{} }
func (m *UploadStatus) String() string { return proto.CompactTextString(m) }
func (*UploadStatus) ProtoMessage()    {}
func (*UploadStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_27a5d7178d16d994, []int{2}
}

func (m *UploadStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadStatus.Unmarshal(m, b)
}
func (m *UploadStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadStatus.Marshal(b, m, deterministic)
}
func (m *UploadStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadStatus.Merge(m, src)
}
func (m *UploadStatus) XXX_Size() int {
	return xxx_messageInfo_UploadStatus.Size(m)
}
func (m *UploadStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadStatus.DiscardUnknown(m)
}

var xxx_messageInfo_UploadStatus proto.InternalMessageInfo

func (m *UploadStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UploadStatus) GetCode() UploadStatusCode {
	if m != nil {
		return m.Code
	}
	return UploadStatusCode_Unknown
}

func init() {
	proto.RegisterEnum("fileservice.UploadStatusCode", UploadStatusCode_name, UploadStatusCode_value)
	proto.RegisterType((*DownloadRequest)(nil), "fileservice.DownloadRequest")
	proto.RegisterType((*Chunk)(nil), "fileservice.Chunk")
	proto.RegisterType((*UploadStatus)(nil), "fileservice.UploadStatus")
}

func init() { proto.RegisterFile("fileservice.proto", fileDescriptor_27a5d7178d16d994) }

var fileDescriptor_27a5d7178d16d994 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xb3, 0x41, 0xd3, 0x3a, 0x2d, 0x35, 0xce, 0xa9, 0x06, 0x05, 0xc9, 0xa9, 0x08, 0x16,
	0x6d, 0x8f, 0x22, 0x58, 0x22, 0xbd, 0xa9, 0x90, 0xd0, 0x93, 0xa7, 0xd5, 0x8c, 0x1a, 0x12, 0x77,
	0x6b, 0x77, 0x63, 0x3e, 0x86, 0x5f, 0x59, 0xb2, 0x61, 0x25, 0x09, 0x7a, 0x7c, 0xbb, 0x6f, 0x7e,
	0x6f, 0xfe, 0xc0, 0xd1, 0x6b, 0x56, 0x90, 0xa2, 0xdd, 0x57, 0xf6, 0x42, 0xf3, 0xed, 0x4e, 0x6a,
	0x89, 0xa3, 0xd6, 0x53, 0x78, 0x01, 0x87, 0x77, 0xb2, 0x12, 0x85, 0xe4, 0x69, 0x4c, 0x9f, 0x25,
	0x29, 0x8d, 0x01, 0x0c, 0xd7, 0x59, 0x41, 0x0f, 0xfc, 0x83, 0xa6, 0xec, 0x8c, 0xcd, 0x0e, 0xe2,
	0x5f, 0x1d, 0xde, 0xc0, 0x7e, 0xf4, 0x5e, 0x8a, 0x1c, 0xa7, 0x30, 0x88, 0xa4, 0xd0, 0x24, 0xb4,
	0xf1, 0x8c, 0x63, 0x2b, 0x3b, 0xe5, 0x6e, 0xaf, 0xfc, 0x09, 0xc6, 0x9b, 0x6d, 0x9d, 0x95, 0x68,
	0xae, 0x4b, 0x55, 0x53, 0xee, 0x49, 0x29, 0xfe, 0x66, 0x93, 0xac, 0xc4, 0x2b, 0xd8, 0x8b, 0x64,
	0xda, 0x10, 0x26, 0x8b, 0xd3, 0x79, 0x7b, 0x8c, 0x36, 0xa2, 0x36, 0xc5, 0xc6, 0x7a, 0xbe, 0x04,
	0xbf, 0xff, 0x83, 0x23, 0x18, 0x6c, 0x44, 0x2e, 0x64, 0x25, 0x7c, 0x07, 0x3d, 0x70, 0x1f, 0x73,
	0x9f, 0x21, 0x80, 0xb7, 0xe6, 0x59, 0x41, 0xa9, 0xef, 0x2e, 0xbe, 0x19, 0x4c, 0x56, 0xab, 0xa4,
	0xee, 0x30, 0x69, 0xf0, 0x78, 0x0d, 0x5e, 0xc3, 0x41, 0xec, 0xc4, 0x9a, 0xc1, 0x83, 0xe3, 0x7f,
	0x5b, 0x09, 0x9d, 0x19, 0xc3, 0x5b, 0x18, 0xda, 0x7d, 0xe2, 0x49, 0xc7, 0xda, 0x5b, 0x73, 0xf0,
	0x07, 0x3c, 0x74, 0x2e, 0xd9, 0xb3, 0x67, 0xae, 0xb4, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0xd6,
	0xc8, 0x0f, 0xc4, 0xba, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AASFileServiceClient is the client API for AASFileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AASFileServiceClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (AASFileService_UploadClient, error)
	Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (AASFileService_DownloadClient, error)
}

type aASFileServiceClient struct {
	cc *grpc.ClientConn
}

func NewAASFileServiceClient(cc *grpc.ClientConn) AASFileServiceClient {
	return &aASFileServiceClient{cc}
}

func (c *aASFileServiceClient) Upload(ctx context.Context, opts ...grpc.CallOption) (AASFileService_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AASFileService_serviceDesc.Streams[0], "/fileservice.AASFileService/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &aASFileServiceUploadClient{stream}
	return x, nil
}

type AASFileService_UploadClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*UploadStatus, error)
	grpc.ClientStream
}

type aASFileServiceUploadClient struct {
	grpc.ClientStream
}

func (x *aASFileServiceUploadClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aASFileServiceUploadClient) CloseAndRecv() (*UploadStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aASFileServiceClient) Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (AASFileService_DownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AASFileService_serviceDesc.Streams[1], "/fileservice.AASFileService/Download", opts...)
	if err != nil {
		return nil, err
	}
	x := &aASFileServiceDownloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AASFileService_DownloadClient interface {
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type aASFileServiceDownloadClient struct {
	grpc.ClientStream
}

func (x *aASFileServiceDownloadClient) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AASFileServiceServer is the server API for AASFileService service.
type AASFileServiceServer interface {
	Upload(AASFileService_UploadServer) error
	Download(*DownloadRequest, AASFileService_DownloadServer) error
}

// UnimplementedAASFileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAASFileServiceServer struct {
}

func (*UnimplementedAASFileServiceServer) Upload(srv AASFileService_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (*UnimplementedAASFileServiceServer) Download(req *DownloadRequest, srv AASFileService_DownloadServer) error {
	return status.Errorf(codes.Unimplemented, "method Download not implemented")
}

func RegisterAASFileServiceServer(s *grpc.Server, srv AASFileServiceServer) {
	s.RegisterService(&_AASFileService_serviceDesc, srv)
}

func _AASFileService_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AASFileServiceServer).Upload(&aASFileServiceUploadServer{stream})
}

type AASFileService_UploadServer interface {
	SendAndClose(*UploadStatus) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type aASFileServiceUploadServer struct {
	grpc.ServerStream
}

func (x *aASFileServiceUploadServer) SendAndClose(m *UploadStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aASFileServiceUploadServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AASFileService_Download_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AASFileServiceServer).Download(m, &aASFileServiceDownloadServer{stream})
}

type AASFileService_DownloadServer interface {
	Send(*Chunk) error
	grpc.ServerStream
}

type aASFileServiceDownloadServer struct {
	grpc.ServerStream
}

func (x *aASFileServiceDownloadServer) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

var _AASFileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fileservice.AASFileService",
	HandlerType: (*AASFileServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _AASFileService_Upload_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Download",
			Handler:       _AASFileService_Download_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "fileservice.proto",
}