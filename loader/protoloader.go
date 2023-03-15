package loader

import (
	"context"
	"fmt"

	// "my/remote/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

// LoadRemoteProtoAny 在这个示例中，我们使用了 google.golang.org/protobuf/types/known/anypb 包中的 Any 类型来表示任意类型的 proto 消息。
// 首先，我们使用 any.Descriptor() 方法获取消息类型的描述符，然后通过 gRPC 调用查找远程服务上的消息类型。最后，我们将消息反序列化为具体的消息类型。
// 请注意，为了使此示例代码工作，您需要替换 my/remote/proto 为您自己的远程 proto 包路径。此外，您还需要实现 proto.RemoteProtoService 服务来提供查找消息类型的功能。
func LoadRemoteProtoAny(ctx context.Context, conn *grpc.ClientConn, any *anypb.Any) (proto.Message, error) {
	client, err := NewRemoteProtoServiceClient(ctx, conn)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to get descriptor for any: %v", err))
	}

	// 创建一个消息类型的实例，它可以解析任意消息类型。
	var msg proto.Message

	// 获取消息类型的描述符。
	descriptor := any.ProtoReflect().Descriptor()
	// if err != nil {
	// 	return nil, status.Error(codes.Internal, fmt.Sprintf("failed to get descriptor for any: %v", err))
	// }

	// 通过描述符在远程服务上查找消息类型。
	findReq := &FindTypeRequest{TypeName: descriptor.FullName()}
	findResp, err := client.FindType(ctx, findReq)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to find type for any: %v", err))
	}

	// 从远程服务上获取到消息类型之后，创建一个对应的实例。
	if findResp.TypeUrl != "" {
		msgType := MessageType(findResp.TypeUrl)
		if msgType != nil {
			// msg = msgType.New().Interface().(proto.Message)
			msg = msgType.New().Interface()
		}
	}

	if msg == nil {
		return nil, status.Error(codes.Internal, "failed to create message for any")
	}

	// 将任意消息类型反序列化为具体的消息类型。
	if err := any.UnmarshalTo(msg); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to unmarshal any to message: %v", err))
	}

	return msg, nil
}
