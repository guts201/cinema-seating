package server

import (
	"context"
	"fmt"
	"net"
	"os"

	config "cinema/pkg/carbon"
	mykit "cinema/pkg/genkit/pkg/api"
	"cinema/pkg/logging"

	"github.com/DataDog/datadog-go/statsd"
	grpcctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthv1 "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
	grpctrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc"

	conf "cinema/pkg/config"
)

func Run(f *config.Flags) {
	cfg := loadConfig(f)

	Serve(cfg)
}

func newService(cfg *conf.Config, opts ...mykit.Option) mykit.Service {
	var (
		statsdClient *statsd.Client
		err          error
	)

	err = logging.InitLogger(cfg.Logger)
	if err != nil {
		logging.NewTmpLogger().Error("fail to init logger", zap.Error(err))
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Listener.GetTcp().Address, cfg.Listener.GetTcp().Port))
	if err != nil {
		logging.NewTmpLogger().Fatal("failed to new listener", zap.Error(err))
	}

	logger := logging.Logger(context.Background())

	healthServer := health.NewServer()
	healthServer.SetServingStatus("", healthv1.HealthCheckResponse_SERVING)

	defaultOpts := []mykit.Option{
		mykit.Stats(statsdClient),
		mykit.Logger(logger),
		mykit.Listener(listener),
		mykit.ServerOptions(
			grpc.ChainUnaryInterceptor(
				grpcctxtags.UnaryServerInterceptor(grpcctxtags.WithFieldExtractor(grpcctxtags.CodeGenRequestFieldExtractor)),
				grpctrace.UnaryServerInterceptor(grpctrace.WithUntracedMethods("/grpc.health.v1.Health/Check"), grpctrace.WithServiceName(os.Getenv("DD_SERVICE"))),
			),
		),
		mykit.HealthServer(healthServer),
	}

	svc := mykit.NewService(append(defaultOpts, opts...)...)
	return svc
}

func GetRequestID(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	mdUserID := md.Get("x-request-id")
	if len(mdUserID) < 1 {
		return ""
	}
	return mdUserID[0]
}

func loadConfig(f *config.Flags) *conf.Config {
	// Use a temporary logger to parse the configuration and output.
	tmpLogger := logging.NewTmpLogger().With(zap.String("filename", f.ConfigPath))

	var cfg conf.Config
	if err := config.ParseFile(f.ConfigPath, &cfg, f.Template); err != nil {
		tmpLogger.Fatal("parsing configuration failed", zap.Error(err))
	}

	if err := cfg.Validate(); err != nil {
		tmpLogger.Fatal("validating configuration failed", zap.Error(err))
	}

	if f.Validate {
		tmpLogger.Info("configuration validation was successful")
		os.Exit(0)
	}

	return &cfg
}
