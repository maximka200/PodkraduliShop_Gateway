package product

import (
	"context"
	"log/slog"
	"time"

	grpclog "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	productv1 "github.com/maximka200/protobuff_product/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	Api productv1.ProductClient
	Log *slog.Logger
}

func NewClient(log *slog.Logger, addr string, timeout time.Duration, retryCount int) (*Client, error) {
	const op = "product.NewClient"

	retryOpts := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(retryCount)),
		grpcretry.WithPerRetryTimeout(timeout),
	}

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(grpcretry.UnaryClientInterceptor(retryOpts...)))
	if err != nil {
		log.Error("%s: %s", op, err)
	}

	return &Client{Api: productv1.NewProductClient(conn)}, nil // insert logg if need (use InterceptorLogger)
}

// InterceptorLogger adapts slog logger to interceptor logger.
func InterceptorLogger(l *slog.Logger) grpclog.Logger {
	return grpclog.LoggerFunc(func(ctx context.Context, lvl grpclog.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}
