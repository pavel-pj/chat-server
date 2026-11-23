package interceptor

import (
	"context"

	"github.com/pavel-pj/go-chat/internal/authclient"
	"google.golang.org/grpc"
)

func Auth(client *authclient.Client) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		if err := client.Check(ctx, info.FullMethod); err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}
