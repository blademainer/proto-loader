package loader

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Client grpc proto client
type Client interface {
	FindType(ctx context.Context, req *FindTypeRequest) (*FindTypeResponse, error)
}

// FindTypeRequest request of find type
type FindTypeRequest struct {
	TypeName protoreflect.FullName
}

// FindTypeResponse response of find type
type FindTypeResponse struct {
	TypeName protoreflect.FullName
	TypeUrl  string
}

// RemoteClient remote client
type RemoteClient struct {
}

// FindType 查找远程类型
//  @receiver r
//  @param ctx
//  @param req
//  @return *FindTypeResponse
//  @return error
//
func (r *RemoteClient) FindType(ctx context.Context, req *FindTypeRequest) (*FindTypeResponse, error) {
	// TODO 增加实现
	return nil, nil
}

// NewRemoteProtoServiceClient 创建client
//  @param ctx
//  @param conn
//  @return Client
//  @return error
//
func NewRemoteProtoServiceClient(ctx context.Context, conn *grpc.ClientConn) (Client, error) {
	c := &RemoteClient{}
	return c, nil
}
