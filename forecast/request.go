package forecast

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"sync/atomic"

	"github.com/jh1104/publicapi"
)

var defaultClient atomic.Pointer[publicapi.Client]

func DefaultClient() *publicapi.Client {
	return defaultClient.Load()
}

func SetDefaultClient(client *publicapi.Client) {
	defaultClient.Store(client)
}

// 주어진 Parameters를 사용하여 초단기예보를 조회한다.
func GetUltraShortTermForecast(ctx context.Context, params Parameters) (*Response, error) {
	api := &Forecast{
		Subtype: UltraShortTermForecast,
		Params:  params,
	}
	return request(ctx, api)
}

func GetShortTermForecast(ctx context.Context, params Parameters) (*Response, error) {
	api := &Forecast{
		Subtype: ShortTermForecast,
		Params:  params,
	}
	return request(ctx, api)
}

func request(ctx context.Context, api publicapi.API) (*Response, error) {
	client := DefaultClient()
	if client == nil {
		return nil, errors.New("default client is not initialized")
	}

	data, err := client.RequestAPI(ctx, api)
	if err != nil {
		return nil, err
	}

	if len(data) > 0 && data[0] == '<' {
		resp := &publicapi.ErrorResponse{}
		if err := xml.Unmarshal(data, resp); err == nil && resp.AsError() != nil {
			return nil, resp.AsError()
		}
		return nil, errors.New("received XML response instead of JSON")
	}

	resp := &Response{}
	if err := json.Unmarshal(data, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
