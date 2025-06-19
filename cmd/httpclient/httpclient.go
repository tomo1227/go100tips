package main

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
)

type httpClient struct {
	*http.Client
}

var (
	client         *httpClient
	httpClientOnce sync.Once
)

// GetHTTPClient は http.Client のシングルトンインスタンスを返します。

func GetHTTPClient() *httpClient {
	httpClientOnce.Do(func() {
		client = &httpClient{Client: &http.Client{
			// TODO: timeoutの設定
			Timeout: 30 * time.Second,
		}}
	})
	return client
}

// DoRequest は http.Request を送信し、レスポンスを返します。

// レスポンスは resty.Response で返されます。 https://pkg.go.dev/github.com/go-resty/resty/v2 を参照してください。

func DoRequest(method Method, url string, headers map[string]string, data []byte) (*resty.Response, error) {
	req := resty.NewWithClient(GetHTTPClient().Client).R()

	for key, value := range headers {
		req.SetHeader(key, value)
	}

	if data != nil && (method == MethodPost || method == MethodPut || method == MethodPatch) {
		req.SetBody(data)
	}

	var resp *resty.Response
	var err error

	switch method {
	case MethodGet:
		resp, err = req.Get(url)
	case MethodPost:
		resp, err = req.Post(url)
	case MethodPut:
		resp, err = req.Put(url)
	case MethodDelete:
		resp, err = req.Delete(url)
	case MethodPatch:
		resp, err = req.Patch(url)
	case MethodHead:
		resp, err = req.Head(url)
	case MethodOptions:
		resp, err = req.Options(url)
	default:
		err = errors.New("invalid method")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to send %s request to %s: %w", method, url, err)
	}

	return resp, nil
}
