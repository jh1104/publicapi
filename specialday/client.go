package specialday

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/jh1104/publicapi"
)

type apiPath = string

const (
	// 공휴일 API
	apiHoliday apiPath = "/B090041/openapi/service/SpcdeInfoService/getRestDeInfo"
	// 국경일 API
	apiNationalHoliday apiPath = "/B090041/openapi/service/SpcdeInfoService/getHoliDeInfo"
	// 기념일 API
	apiAnniversary apiPath = "/B090041/openapi/service/SpcdeInfoService/getAnniversaryInfo"
)

type Client struct {
	// 서비스 키.
	ServiceKey string

	// HTTP 클라이언트. nil인 경우 http.DefaultClient를 사용한다.
	HTTPClient *http.Client
}

func NewClient(serviceKey string) *Client {
	return &Client{
		ServiceKey: serviceKey,
		HTTPClient: publicapi.DefaultClient,
	}
}

// 공휴일 조회.
func (c *Client) ListHolidays(ctx context.Context, params Parameters) (*Response, error) {
	return c.requestAPI(ctx, "GET", apiHoliday, params)
}

// 국경일 조회.
func (c *Client) ListNationalHolidays(ctx context.Context, params Parameters) (*Response, error) {
	return c.requestAPI(ctx, "GET", apiNationalHoliday, params)
}

// 기념일 조회.
func (c *Client) ListAnniversaries(ctx context.Context, params Parameters) (*Response, error) {
	return c.requestAPI(ctx, "GET", apiAnniversary, params)
}

func (c *Client) buildURL(path string, params Parameters) string {
	values := url.Values{}
	values.Add("serviceKey", c.ServiceKey)
	values.Add("solYear", fmt.Sprintf("%d", params.Year))
	values.Add("solMonth", fmt.Sprintf("%02d", params.Month))
	values.Add("numOfRows", fmt.Sprintf("%d", params.NumberOfRows))
	values.Add("pageNo", fmt.Sprintf("%d", params.PageNo))
	values.Add("_type", "json")

	url.JoinPath(publicapi.BaseURL, path, values.Encode())
	return publicapi.BaseURL + path + "?" + values.Encode()
}

func (c *Client) requestAPI(
	ctx context.Context,
	method string,
	path string,
	params Parameters,
) (*Response, error) {
	url := c.buildURL(path, params)
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
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
		if len(data) > 0 {
			return nil, errors.New(string(data))
		} else {
			return nil, errors.New(resp.Status)
		}
	}

	result := &Response{}
	if err := json.Unmarshal(data, result); err != nil {
		return nil, errors.Join(errors.New("response body unmarshal error"), err)
	}

	return result, nil
}
