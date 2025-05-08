package logging

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"

	logger "cinema/pkg/config"
)

var (
	_logger           = NewTmpLogger()
	_xRequestIDHeader = "x_request_id"
)

func NewLogger(msg *logger.Logger) (*zap.Logger, error) {
	var c zap.Config
	var opts []zap.Option
	if msg.GetPretty() {
		c = zap.NewDevelopmentConfig()
		opts = append(opts, zap.AddStacktrace(zap.ErrorLevel))
	} else {
		c = zap.NewProductionConfig()
	}

	level := zap.NewAtomicLevel()

	levelName := "INFO"
	if msg.Level != logger.Logger_UNSPECIFIED {
		levelName = msg.Level.String()
	}

	if err := level.UnmarshalText([]byte(levelName)); err != nil {
		return nil, fmt.Errorf("could not parse log level %s", msg.Level.String())
	}
	c.Level = level

	return c.Build(opts...)
}

func InitLogger(msg *logger.Logger) (err error) {
	_logger, err = NewLogger(msg)
	return err
}

func NewTmpLogger() *zap.Logger {
	c := zap.NewProductionConfig()
	c.DisableStacktrace = true
	l, err := c.Build()
	if err != nil {
		panic(err)
	}
	return l
}

// Logger Return new logger with context value
// ctx:  nillable
func Logger(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return _logger
	}
	logger := injectXRequestID(_logger, ctx)
	return logger
}

func SetXRequestIDHeader(headerName string) {
	_xRequestIDHeader = headerName
}

func injectXRequestID(logger *zap.Logger, ctx context.Context) *zap.Logger {
	if ctx == nil {
		return logger
	}
	requestID := getRequestID(ctx)
	if requestID == "" {
		return logger
	}
	return logger.With(zap.String(_xRequestIDHeader, requestID))
}

func getRequestID(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	requestIds := md.Get(_xRequestIDHeader)
	if len(requestIds) < 1 {
		return ""
	}
	return requestIds[0]
}
