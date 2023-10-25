package sample

import (
	"github.com/prometheus/client_golang/prometheus"
	"go.elastic.co/apm/module/apmhttp/v2"
	"net/http"
)

type SampleOutbound struct {
	client *http.Client
	metric *prometheus.CounterVec
}

func NewOutbound(metric *prometheus.CounterVec) *SampleOutbound {
	client := apmhttp.WrapClient(http.DefaultClient)

	return &SampleOutbound{
		client: client,
		metric: metric,
	}
}
