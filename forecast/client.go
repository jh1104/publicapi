package forecast

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/jh1104/publicapi"
)

type apiPath = string

const (
	apiUltraShortTermForecast apiPath = "/1360000/VilageFcstInfoService_2.0/getUltraSrtFcst"
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

// 초단기 예보 조회.
func (c *Client) GetUltraShortTermForecast(ctx context.Context, params Parameters) (*Response, error) {
	return c.requestAPI(ctx, "GET", apiUltraShortTermForecast, params)
}

func (c *Client) buildURL(path string, params Parameters) string {
	values := url.Values{}
	values.Add("serviceKey", c.ServiceKey)
	values.Add("dataType", "JSON")
	values.Add("pageNo", strconv.Itoa(params.PageNo))
	values.Add("numOfRows", strconv.Itoa(params.NumberOfRows))
	values.Add("base_date", params.BaseDate)
	values.Add("base_time", params.BaseTime)
	values.Add("nx", strconv.Itoa(params.NX))
	values.Add("ny", strconv.Itoa(params.NY))
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
