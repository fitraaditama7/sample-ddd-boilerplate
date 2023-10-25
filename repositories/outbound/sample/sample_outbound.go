package sample

import (
	"context"
	httpclient "ddd-boilerplate/pkg/http-client"
	"ddd-boilerplate/pkg/metrics"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"

	"ddd-boilerplate/pkg/logger"
)

func (a *SampleOutbound) FindSampleAPI(ctx context.Context, id int64) (*SampleAPIResponse, error) {
	log := logger.Ctx(ctx)
	var sampleResponse SampleAPIResponse
	var startAt = time.Now()
	var path = fmt.Sprintf("https://dummyjson.com/products/1")

	req, err := httpclient.Prepare(ctx, fiber.MethodGet, path, nil)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	q := req.URL.Query()
	q.Add("test", "test")
	req.URL.RawQuery = q.Encode()

	code, resp, err := httpclient.Exec(ctx, a.client, req, &sampleResponse)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	defer func() {
		logger.LogOutboundRequest(ctx, req, nil, resp, code, startAt)
		metrics.BuildOutboundPrometheusMetrics(ctx, a.metric, path, code)
	}()

	return &sampleResponse, nil
}

func (a *SampleOutbound) PostSampleAPI(ctx context.Context, request SamplePostAPIRequest) (*SamplePostAPIResponse, error) {
	log := logger.Ctx(ctx)
	var sampleResponse SamplePostAPIResponse
	var startAt = time.Now()
	var path = fmt.Sprintf("https://dummyjson.com/products/add")

	req, err := httpclient.Prepare(ctx, fiber.MethodPost, path, request)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	req.Header.Add("test", "test")

	code, resp, err := httpclient.Exec(ctx, a.client, req, &sampleResponse)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	defer func() {
		logger.LogOutboundRequest(ctx, req, request, resp, code, startAt)
		metrics.BuildOutboundPrometheusMetrics(ctx, a.metric, path, code)
	}()

	return &sampleResponse, nil
}
