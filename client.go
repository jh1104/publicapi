package publicapi

import (
	"context"
	"errors"
	"io"
	"net/http"
)

type API interface {
	// HTTP 메서드.
	Method() string
	// HTTP 요청 URL.
	URL(serviceKey string) string
	// HTTP 요청 바디. GET 메서드인 경우 nil을 반환한다.
	Body() io.Reader
}

type Client struct {
	// 공공데이터포털 서비스 키.
	ServiceKey string

	// HTTP 클라이언트.
	HTTPClient *http.Client
}

func NewClient(serviceKey string) *Client {
	return &Client{
		ServiceKey: serviceKey,
		HTTPClient: DefaultHTTPClient,
	}
}

// 주어진 API를 수행하고, 응답 바디를 바이트 슬라이스로 반환합니다.
func (c *Client) RequestAPI(ctx context.Context, api API) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, api.Method(), api.URL(c.ServiceKey), api.Body())
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Join(errors.New("response body read error"), err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, errors.New(resp.Status)
	}

	return data, nil
}
