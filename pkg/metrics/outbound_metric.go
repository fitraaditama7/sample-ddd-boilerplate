package metrics

import (
	"context"
	"ddd-boilerplate/pkg/logger"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"net/url"
)

func SetupOutboundMetric() *prometheus.CounterVec {
	outboundCall := prometheus.CounterOpts{
		Name: "http_request_outbound_callbacks",
		Help: "Get Response information from outbound http request",
	}
	outboundCallLabelName := []string{
		"host",
		"path",
		"response_code",
	}

	metricOutbound := prometheus.NewCounterVec(outboundCall, outboundCallLabelName)
	return metricOutbound
}

func BuildOutboundPrometheusMetrics(ctx context.Context, c *prometheus.CounterVec, callbackURL string, statusCode int) {
	log := logger.Ctx(ctx)
	parsedURL, err := url.Parse(callbackURL)
	if err != nil {
		log.Error(err.Error())
		return
	}

	c.With(prometheus.Labels{
		"host":          parsedURL.Host,
		"path":          parsedURL.Path,
		"response_code": fmt.Sprintf("%d", statusCode),
	}).Inc()
}
