package forecast

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/jh1104/publicapi"
)

var defaultClient *publicapi.Client

func SetDefaultClient(client *publicapi.Client) {
	defaultClient = client
}

func GetUltraShortTermForecast(ctx context.Context, nx, ny int) (*Response, error) {
	api := &Forecast{
		Subtype: UltraShortTermForecast,
		Params:  *NewParameters(nx, ny),
	}
	return request(ctx, api)
}

func request(ctx context.Context, api publicapi.API) (*Response, error) {
	if defaultClient == nil {
		return nil, errors.New("default client is not initialized")
	}

	data, err := defaultClient.RequestAPI(ctx, api)
	if err != nil {
		return nil, err
	}

	resp := &Response{}
	if err := json.Unmarshal(data, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
