package middleware

import (
	"bytes"
	"ddd-boilerplate/pkg/logger"
	"ddd-boilerplate/pkg/responses"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.elastic.co/apm/v2"
	"go.uber.org/zap"
	"strconv"
	"time"
)

func Logger() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		startAt := time.Now()
		preHandleRequest(c)

		defer func() {
			postHandleRequest(c, startAt)
		}()

		return c.Next()
	}
}

func preHandleRequest(c *fiber.Ctx) {
	log := logger.Logger
	traceID := apm.TransactionFromContext(c.Context()).TraceContext().Trace
	header := make(map[string]string)
	c.Request().Header.VisitAll(func(key, val []byte) {
		k := bytes.NewBuffer(key).String()
		header[k] = bytes.NewBuffer(val).String()
	})

	headerByte, _ := json.Marshal(header)

	loggerField := []zap.Field{
		zap.String("path", c.Path()),
		zap.String("trace_id", traceID.String()),
		zap.Any("header", json.RawMessage(headerByte)),
		zap.String("method", c.Method()),
		zap.String("protocol", c.Protocol()),
		zap.String("remote_ip", c.IP()),
	}

	if len(c.Body()) != 0 {
		zap.Any("body", json.RawMessage(c.Body()))
	}

	if (len(c.Request().URI().QueryString())) != 0 {
		zap.Any("query_param", c.Request().URI().QueryString())
	}

	msg := fmt.Sprintf("[REQUEST] %v %v", c.Method(), c.Path())
	log.Info(msg, loggerField...)
}

func postHandleRequest(c *fiber.Ctx, startAt time.Time) {
	log := logger.Logger

	rvr := recover()
	if rvr != nil {
		var ok bool
		err, ok := rvr.(error)
		if !ok {
			err = fmt.Errorf("%v", rvr)
		}

		resp := responses.Error(c, err)
		c.Status(fiber.StatusInternalServerError).JSON(resp)
	}

	traceID := apm.TransactionFromContext(c.Context()).TraceContext().Trace
	fullURL := fmt.Sprintf("%s://%s%s", c.Protocol(), c.Hostname(), c.Path())

	loggerField := []zap.Field{
		zap.String("path", c.Path()),
		zap.String("trace_id", traceID.String()),
		zap.String("method", c.Method()),
		zap.String("protocol", c.Protocol()),
		zap.String("remote_ip", c.IP()),
		zap.Any("status_code", c.Response().StatusCode()),
		zap.Any("response", json.RawMessage(c.Response().Body())),
		zap.Float64("latency", time.Since(startAt).Seconds()),
	}

	msg := fmt.Sprintf("[RESPONSE] %d %s %s", c.Response().StatusCode(), c.Method(), fullURL)
	switch strconv.Itoa(c.Response().StatusCode())[0] {
	case '1', '2', '3':
		log.Info(msg, loggerField...)
	case '4', '5':
		log.Error(msg, loggerField...)
	default:
		log.Panic(msg, loggerField...)
	}
}
