package product

import (
	"context"
	"log/slog"
	"time"

	grpclog "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	authv1 "github.com/maximka200/buffpr/gen/go/sso"
	productv1 "github.com/maximka200/protobuff_product/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductClient struct {
	Api productv1.ProductClient
	Log *slog.Logger
}

type AuthClient struct {
	Api authv1.AuthClient
	Log *slog.Logger
}

type GRPCMethods struct {
	AuthClient
	ProductClient
}

func NewGRPCMethods(auth AuthClient, product ProductClient) *GRPCMethods {
	return &GRPCMethods{
		auth, product,
	}
}

func NewClientProduct(log *slog.Logger, addr string,
	timeout time.Duration, retryCount int) (*ProductClient, error) {
	const op = "product.NewClientProduct"

	retryOpts := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(retryCount)),
		grpcretry.WithPerRetryTimeout(timeout),
	}

	conn, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(grpcretry.UnaryClientInterceptor(retryOpts...)))
	if err != nil {
		log.Error("%s: %s", op, err)
	}

	return &ProductClient{Api: productv1.NewProductClient(conn)}, nil
	// insert logg if need (use InterceptorLogger)
}

func NewClientAuth(log *slog.Logger, addr string,
	timeout time.Duration, retryCount int) (*AuthClient, error) {
	const op = "product.NewClientAuth"

	retryOpts := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(retryCount)),
		grpcretry.WithPerRetryTimeout(timeout),
	}

	conn, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(grpcretry.UnaryClientInterceptor(retryOpts...)))
	if err != nil {
		log.Error("%s: %s", op, err)
	}

	return &AuthClient{Api: authv1.NewAuthClient(conn)}, nil // insert logg if need (use InterceptorLogger)
}

// InterceptorLogger adapts slog logger to interceptor logger.
func InterceptorLogger(l *slog.Logger) grpclog.Logger {
	return grpclog.LoggerFunc(func(ctx context.Context,
		lvl grpclog.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}
