// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order.proto

package order

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateOrderRequest struct {
	Uuid                 string               `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Type                 int32                `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	VisitUuid            string               `protobuf:"bytes,4,opt,name=visit_uuid,json=visitUuid,proto3" json:"visit_uuid,omitempty"`
	StreamCode           string               `protobuf:"bytes,5,opt,name=stream_code,json=streamCode,proto3" json:"stream_code,omitempty"`
	OfferId              int32                `protobuf:"varint,6,opt,name=offer_id,json=offerId,proto3" json:"offer_id,omitempty"`
	UserId               int32                `protobuf:"varint,7,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ClientUuid           string               `protobuf:"bytes,8,opt,name=client_uuid,json=clientUuid,proto3" json:"client_uuid,omitempty"`
	Status               int32                `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty"`
	Sum                  float32              `protobuf:"fixed32,10,opt,name=sum,proto3" json:"sum,omitempty"`
	Payment              float32              `protobuf:"fixed32,11,opt,name=payment,proto3" json:"payment,omitempty"`
	Currency             string               `protobuf:"bytes,12,opt,name=currency,proto3" json:"currency,omitempty"`
	GeoCodes             []string             `protobuf:"bytes,13,rep,name=geo_codes,json=geoCodes,proto3" json:"geo_codes,omitempty"`
	Sub1                 string               `protobuf:"bytes,14,opt,name=sub1,proto3" json:"sub1,omitempty"`
	Sub2                 string               `protobuf:"bytes,15,opt,name=sub2,proto3" json:"sub2,omitempty"`
	Sub3                 string               `protobuf:"bytes,16,opt,name=sub3,proto3" json:"sub3,omitempty"`
	Sub4                 string               `protobuf:"bytes,17,opt,name=sub4,proto3" json:"sub4,omitempty"`
	Sub5                 string               `protobuf:"bytes,18,opt,name=sub5,proto3" json:"sub5,omitempty"`
	ErrorReason          string               `protobuf:"bytes,19,opt,name=error_reason,json=errorReason,proto3" json:"error_reason,omitempty"`
	Pills                int32                `protobuf:"varint,20,opt,name=pills,proto3" json:"pills,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreateOrderRequest) Reset()         { *m = CreateOrderRequest{} }
func (m *CreateOrderRequest) String() string { return proto.CompactTextString(m) }
func (*CreateOrderRequest) ProtoMessage()    {}
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{0}
}

func (m *CreateOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderRequest.Unmarshal(m, b)
}
func (m *CreateOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderRequest.Marshal(b, m, deterministic)
}
func (m *CreateOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderRequest.Merge(m, src)
}
func (m *CreateOrderRequest) XXX_Size() int {
	return xxx_messageInfo_CreateOrderRequest.Size(m)
}
func (m *CreateOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderRequest proto.InternalMessageInfo

func (m *CreateOrderRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *CreateOrderRequest) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *CreateOrderRequest) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *CreateOrderRequest) GetVisitUuid() string {
	if m != nil {
		return m.VisitUuid
	}
	return ""
}

func (m *CreateOrderRequest) GetStreamCode() string {
	if m != nil {
		return m.StreamCode
	}
	return ""
}

func (m *CreateOrderRequest) GetOfferId() int32 {
	if m != nil {
		return m.OfferId
	}
	return 0
}

func (m *CreateOrderRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CreateOrderRequest) GetClientUuid() string {
	if m != nil {
		return m.ClientUuid
	}
	return ""
}

func (m *CreateOrderRequest) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *CreateOrderRequest) GetSum() float32 {
	if m != nil {
		return m.Sum
	}
	return 0
}

func (m *CreateOrderRequest) GetPayment() float32 {
	if m != nil {
		return m.Payment
	}
	return 0
}

func (m *CreateOrderRequest) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *CreateOrderRequest) GetGeoCodes() []string {
	if m != nil {
		return m.GeoCodes
	}
	return nil
}

func (m *CreateOrderRequest) GetSub1() string {
	if m != nil {
		return m.Sub1
	}
	return ""
}

func (m *CreateOrderRequest) GetSub2() string {
	if m != nil {
		return m.Sub2
	}
	return ""
}

func (m *CreateOrderRequest) GetSub3() string {
	if m != nil {
		return m.Sub3
	}
	return ""
}

func (m *CreateOrderRequest) GetSub4() string {
	if m != nil {
		return m.Sub4
	}
	return ""
}

func (m *CreateOrderRequest) GetSub5() string {
	if m != nil {
		return m.Sub5
	}
	return ""
}

func (m *CreateOrderRequest) GetErrorReason() string {
	if m != nil {
		return m.ErrorReason
	}
	return ""
}

func (m *CreateOrderRequest) GetPills() int32 {
	if m != nil {
		return m.Pills
	}
	return 0
}

type Order struct {
	Uuid                 string               `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Type                 int32                `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	VisitUuid            string               `protobuf:"bytes,4,opt,name=visit_uuid,json=visitUuid,proto3" json:"visit_uuid,omitempty"`
	StreamCode           string               `protobuf:"bytes,5,opt,name=stream_code,json=streamCode,proto3" json:"stream_code,omitempty"`
	OfferId              int32                `protobuf:"varint,6,opt,name=offer_id,json=offerId,proto3" json:"offer_id,omitempty"`
	UserId               int32                `protobuf:"varint,7,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ClientUuid           string               `protobuf:"bytes,8,opt,name=client_uuid,json=clientUuid,proto3" json:"client_uuid,omitempty"`
	Status               int32                `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty"`
	Sum                  float32              `protobuf:"fixed32,10,opt,name=sum,proto3" json:"sum,omitempty"`
	Payment              float32              `protobuf:"fixed32,11,opt,name=payment,proto3" json:"payment,omitempty"`
	Currency             string               `protobuf:"bytes,12,opt,name=currency,proto3" json:"currency,omitempty"`
	GeoCodes             []string             `protobuf:"bytes,13,rep,name=geo_codes,json=geoCodes,proto3" json:"geo_codes,omitempty"`
	Sub1                 string               `protobuf:"bytes,14,opt,name=sub1,proto3" json:"sub1,omitempty"`
	Sub2                 string               `protobuf:"bytes,15,opt,name=sub2,proto3" json:"sub2,omitempty"`
	Sub3                 string               `protobuf:"bytes,16,opt,name=sub3,proto3" json:"sub3,omitempty"`
	Sub4                 string               `protobuf:"bytes,17,opt,name=sub4,proto3" json:"sub4,omitempty"`
	Sub5                 string               `protobuf:"bytes,18,opt,name=sub5,proto3" json:"sub5,omitempty"`
	ErrorReason          string               `protobuf:"bytes,19,opt,name=error_reason,json=errorReason,proto3" json:"error_reason,omitempty"`
	Pills                int32                `protobuf:"varint,20,opt,name=pills,proto3" json:"pills,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{1}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Order) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Order) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Order) GetVisitUuid() string {
	if m != nil {
		return m.VisitUuid
	}
	return ""
}

func (m *Order) GetStreamCode() string {
	if m != nil {
		return m.StreamCode
	}
	return ""
}

func (m *Order) GetOfferId() int32 {
	if m != nil {
		return m.OfferId
	}
	return 0
}

func (m *Order) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Order) GetClientUuid() string {
	if m != nil {
		return m.ClientUuid
	}
	return ""
}

func (m *Order) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Order) GetSum() float32 {
	if m != nil {
		return m.Sum
	}
	return 0
}

func (m *Order) GetPayment() float32 {
	if m != nil {
		return m.Payment
	}
	return 0
}

func (m *Order) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *Order) GetGeoCodes() []string {
	if m != nil {
		return m.GeoCodes
	}
	return nil
}

func (m *Order) GetSub1() string {
	if m != nil {
		return m.Sub1
	}
	return ""
}

func (m *Order) GetSub2() string {
	if m != nil {
		return m.Sub2
	}
	return ""
}

func (m *Order) GetSub3() string {
	if m != nil {
		return m.Sub3
	}
	return ""
}

func (m *Order) GetSub4() string {
	if m != nil {
		return m.Sub4
	}
	return ""
}

func (m *Order) GetSub5() string {
	if m != nil {
		return m.Sub5
	}
	return ""
}

func (m *Order) GetErrorReason() string {
	if m != nil {
		return m.ErrorReason
	}
	return ""
}

func (m *Order) GetPills() int32 {
	if m != nil {
		return m.Pills
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateOrderRequest)(nil), "order.CreateOrderRequest")
	proto.RegisterType((*Order)(nil), "order.Order")
}

func init() { proto.RegisterFile("order.proto", fileDescriptor_cd01338c35d87077) }

var fileDescriptor_cd01338c35d87077 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x94, 0xdf, 0x8a, 0xd3, 0x50,
	0x10, 0xc6, 0xc9, 0xb6, 0x69, 0x9b, 0x49, 0xd5, 0x75, 0x5c, 0x74, 0xb6, 0x22, 0x1b, 0xf7, 0xaa,
	0x57, 0x59, 0x6c, 0x77, 0x05, 0x2f, 0x65, 0x41, 0xd8, 0x2b, 0x21, 0xea, 0x75, 0x48, 0x93, 0x69,
	0x09, 0x34, 0x3d, 0xf1, 0xfc, 0x59, 0xe8, 0xa3, 0xf8, 0x02, 0x3e, 0xa7, 0x64, 0x4e, 0x1a, 0x14,
	0x5f, 0x61, 0xef, 0xe6, 0xfb, 0xcd, 0x39, 0xf3, 0x71, 0x0e, 0x1f, 0x03, 0xb1, 0xd2, 0x15, 0xeb,
	0xb4, 0xd5, 0xca, 0x2a, 0x0c, 0x45, 0x2c, 0xae, 0x76, 0x4a, 0xed, 0xf6, 0x7c, 0x23, 0x70, 0xe3,
	0xb6, 0x37, 0xb6, 0x6e, 0xd8, 0xd8, 0xa2, 0x69, 0xfd, 0xb9, 0xeb, 0xdf, 0x63, 0xc0, 0x7b, 0xcd,
	0x85, 0xe5, 0xaf, 0xdd, 0x85, 0x8c, 0x7f, 0x3a, 0x36, 0x16, 0x11, 0xc6, 0xce, 0xd5, 0x15, 0x05,
	0x49, 0xb0, 0x8c, 0x32, 0xa9, 0xf1, 0x13, 0x40, 0x29, 0x27, 0xab, 0xbc, 0xb0, 0x74, 0x96, 0x04,
	0xcb, 0x78, 0xb5, 0x48, 0xbd, 0x41, 0x7a, 0x32, 0x48, 0xbf, 0x9f, 0x0c, 0xb2, 0xa8, 0x3f, 0xfd,
	0x59, 0xc6, 0xd9, 0x63, 0xcb, 0x34, 0x4a, 0x82, 0x65, 0x98, 0x49, 0x8d, 0xef, 0x00, 0x1e, 0x6b,
	0x53, 0xdb, 0x5c, 0x8c, 0xc6, 0x62, 0x14, 0x09, 0xf9, 0xd1, 0xb9, 0x5d, 0x41, 0x6c, 0xac, 0xe6,
	0xa2, 0xc9, 0x4b, 0x55, 0x31, 0x85, 0xd2, 0x07, 0x8f, 0xee, 0x55, 0xc5, 0x78, 0x09, 0x33, 0xb5,
	0xdd, 0xb2, 0xce, 0xeb, 0x8a, 0x26, 0x32, 0x77, 0x2a, 0xfa, 0xa1, 0xc2, 0x37, 0x30, 0x75, 0xc6,
	0x77, 0xa6, 0xd2, 0x99, 0x74, 0xf2, 0x41, 0x86, 0x96, 0xfb, 0x9a, 0x0f, 0xbd, 0xe9, 0xcc, 0x0f,
	0xf5, 0x48, 0x5c, 0x5f, 0xc3, 0xc4, 0xd8, 0xc2, 0x3a, 0x43, 0x91, 0xbf, 0xe8, 0x15, 0x9e, 0xc3,
	0xc8, 0xb8, 0x86, 0x20, 0x09, 0x96, 0x67, 0x59, 0x57, 0x22, 0xc1, 0xb4, 0x2d, 0x8e, 0x0d, 0x1f,
	0x2c, 0xc5, 0x42, 0x4f, 0x12, 0x17, 0x30, 0x2b, 0x9d, 0xd6, 0x7c, 0x28, 0x8f, 0x34, 0x17, 0x87,
	0x41, 0xe3, 0x5b, 0x88, 0x76, 0xac, 0xe4, 0x49, 0x86, 0x9e, 0x25, 0xa3, 0xae, 0xb9, 0x63, 0xd5,
	0x3d, 0xc8, 0x74, 0xbf, 0x64, 0xdc, 0xe6, 0x03, 0x3d, 0xf7, 0x9f, 0xde, 0xd5, 0x3d, 0x5b, 0xd1,
	0x8b, 0x81, 0xad, 0x7a, 0xb6, 0xa6, 0xf3, 0x81, 0xad, 0x7b, 0x76, 0x4b, 0x2f, 0x07, 0x76, 0xdb,
	0xb3, 0x3b, 0xc2, 0x81, 0xdd, 0xe1, 0x7b, 0x98, 0xb3, 0xd6, 0x4a, 0xe7, 0x9a, 0x0b, 0xa3, 0x0e,
	0xf4, 0x4a, 0x7a, 0xb1, 0xb0, 0x4c, 0x10, 0x5e, 0x40, 0xd8, 0xd6, 0xfb, 0xbd, 0xa1, 0x0b, 0xf9,
	0x02, 0x2f, 0xae, 0x7f, 0x8d, 0x21, 0x94, 0x88, 0x3c, 0x65, 0xe3, 0x29, 0x1b, 0xff, 0x66, 0x63,
	0xf5, 0x05, 0xe6, 0x12, 0x8d, 0x6f, 0xac, 0x1f, 0xeb, 0x92, 0xf1, 0x23, 0xc4, 0x7f, 0xed, 0x14,
	0xbc, 0x4c, 0xfd, 0x66, 0xfa, 0x7f, 0xcf, 0x2c, 0xe6, 0x7d, 0x4b, 0xe0, 0x66, 0x22, 0x49, 0x59,
	0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x81, 0x6a, 0xf0, 0xca, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderServiceClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*Order, error)
}

type orderServiceClient struct {
	cc *grpc.ClientConn
}

func NewOrderServiceClient(cc *grpc.ClientConn) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/order.OrderService/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
type OrderServiceServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*Order, error)
}

func RegisterOrderServiceServer(s *grpc.Server, srv OrderServiceServer) {
	s.RegisterService(&_OrderService_serviceDesc, srv)
}

func _OrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderService/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "order.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderService_CreateOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
