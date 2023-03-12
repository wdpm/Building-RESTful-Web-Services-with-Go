// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.3.0
// source: transaction.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MoneyTransaction_MakeTransaction_FullMethodName = "/datafiles.MoneyTransaction/MakeTransaction"
)

// MoneyTransactionClient is the client API for MoneyTransaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MoneyTransactionClient interface {
	MakeTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (MoneyTransaction_MakeTransactionClient, error)
}

type moneyTransactionClient struct {
	cc grpc.ClientConnInterface
}

func NewMoneyTransactionClient(cc grpc.ClientConnInterface) MoneyTransactionClient {
	return &moneyTransactionClient{cc}
}

func (c *moneyTransactionClient) MakeTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (MoneyTransaction_MakeTransactionClient, error) {
	stream, err := c.cc.NewStream(ctx, &MoneyTransaction_ServiceDesc.Streams[0], MoneyTransaction_MakeTransaction_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &moneyTransactionMakeTransactionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MoneyTransaction_MakeTransactionClient interface {
	Recv() (*TransactionResponse, error)
	grpc.ClientStream
}

type moneyTransactionMakeTransactionClient struct {
	grpc.ClientStream
}

func (x *moneyTransactionMakeTransactionClient) Recv() (*TransactionResponse, error) {
	m := new(TransactionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MoneyTransactionServer is the server API for MoneyTransaction service.
// All implementations should embed UnimplementedMoneyTransactionServer
// for forward compatibility
type MoneyTransactionServer interface {
	MakeTransaction(*TransactionRequest, MoneyTransaction_MakeTransactionServer) error
}

// UnimplementedMoneyTransactionServer should be embedded to have forward compatible implementations.
type UnimplementedMoneyTransactionServer struct {
}

func (UnimplementedMoneyTransactionServer) MakeTransaction(*TransactionRequest, MoneyTransaction_MakeTransactionServer) error {
	return status.Errorf(codes.Unimplemented, "method MakeTransaction not implemented")
}

// UnsafeMoneyTransactionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MoneyTransactionServer will
// result in compilation errors.
type UnsafeMoneyTransactionServer interface {
	mustEmbedUnimplementedMoneyTransactionServer()
}

func RegisterMoneyTransactionServer(s grpc.ServiceRegistrar, srv MoneyTransactionServer) {
	s.RegisterService(&MoneyTransaction_ServiceDesc, srv)
}

func _MoneyTransaction_MakeTransaction_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TransactionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MoneyTransactionServer).MakeTransaction(m, &moneyTransactionMakeTransactionServer{stream})
}

type MoneyTransaction_MakeTransactionServer interface {
	Send(*TransactionResponse) error
	grpc.ServerStream
}

type moneyTransactionMakeTransactionServer struct {
	grpc.ServerStream
}

func (x *moneyTransactionMakeTransactionServer) Send(m *TransactionResponse) error {
	return x.ServerStream.SendMsg(m)
}

// MoneyTransaction_ServiceDesc is the grpc.ServiceDesc for MoneyTransaction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MoneyTransaction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "datafiles.MoneyTransaction",
	HandlerType: (*MoneyTransactionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MakeTransaction",
			Handler:       _MoneyTransaction_MakeTransaction_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "transaction.proto",
}
