package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"go.elastic.co/apm/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"strconv"
	"time"
)

var Logger *zap.Logger

func InitializeLogger() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	Logger, _ = config.Build()
}

func Ctx(ctx context.Context) *zap.Logger {
	log := Logger
	traceID := apm.TransactionFromContext(ctx).TraceContext().Trace
	return log.With(zap.String("trace_id", traceID.String()))
}

func LogOutboundRequest(ctx context.Context, req *http.Request, request interface{}, res []byte, code int, startAt time.Time) {
	log := Ctx(ctx)
	traceID := apm.TransactionFromContext(ctx).TraceContext().Trace
	fullURL := fmt.Sprintf("%s://%s%s", req.URL.Scheme, req.Host, req.URL.Path)

	loggerField := []zap.Field{
		zap.String("protocol", req.URL.Scheme),
		zap.String("host", req.URL.Host),
		zap.String("path", req.URL.Path),
		zap.String("trace_id", traceID.String()),
		zap.String("method", req.Method),
		zap.String("remote_ip", req.RemoteAddr),
	}

	if request != nil {
		b, err := json.Marshal(request)
		if err != nil {
			log.Error(err.Error())
		}
		loggerField = append(loggerField, zap.Any("request", json.RawMessage(b)))
	}

	if len(req.URL.Query()) != 0 {
		loggerField = append(loggerField, zap.String("query_param", req.URL.RawQuery))
	}

	loggerField = append(loggerField,
		zap.Int("status_code", code),
		zap.Float64("latency", time.Since(startAt).Seconds()),
	)

	if len(res) != 0 {
		loggerField = append(loggerField, zap.Any("response", json.RawMessage(res)))
	}

	msg := fmt.Sprintf("%d %s %s", code, req.Method, fullURL)
	switch strconv.Itoa(code)[0] {
	case '1', '2', '3':
		log.Info(msg, loggerField...)
	case '4', '5':
		log.Error(msg, loggerField...)
	default:
		log.Panic(msg, loggerField...)
	}
}
