package httpclient

import (
	"bytes"
	"context"
	"ddd-boilerplate/pkg/logger"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http"
	"os"
	"time"

	"ddd-boilerplate/config"
	"go.elastic.co/apm/module/apmhttp/v2"
)

func CreateHttpClient(config *config.HttpClientConfig) *http.Client {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxIdleConns = config.MaxIdleConns
	transport.MaxConnsPerHost = config.MaxConnsPerHost
	transport.MaxIdleConnsPerHost = config.MaxIdleConnsPerHost
	return apmhttp.WrapClient(&http.Client{
		Timeout:   time.Duration(config.MaxTimeout),
		Transport: transport,
	})
}

func Prepare(ctx context.Context, method, path string, body interface{}) (*http.Request, error) {
	log := logger.Ctx(ctx)
	var requestBody io.Reader
	if body != nil {
		temp, err := json.Marshal(body)
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}
		requestBody = bytes.NewBuffer(temp)
	}

	req, err := http.NewRequest(method, path, requestBody)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return req, nil
}

func Exec(ctx context.Context, client *http.Client, req *http.Request, output interface{}) (int, []byte, error) {
	var result []byte
	log := logger.Ctx(ctx)

	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		log.Error(err.Error())
		return 0, nil, err
	}
	defer resp.Body.Close()

	if os.IsTimeout(err) {
		log.Error(err.Error())
		return fiber.StatusRequestTimeout, nil, err
	}

	result, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err.Error())
		return 0, nil, err
	}

	if resp.StatusCode == http.StatusOK {
		err = json.Unmarshal(result, &output)
		if err != nil {
			log.Error(err.Error())
			return 0, nil, err
		}
	}

	return resp.StatusCode, result, nil
}
