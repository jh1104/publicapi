package specialday

import (
	"context"
	"encoding/json"
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

func ListHolidays(ctx context.Context, params Parameters) (*Response, error) {
	api := &SpecialDay{
		Subtype: Holiday,
		Params:  params,
	}
	return request(ctx, api)
}

func ListNationalHolidays(ctx context.Context, params Parameters) (*Response, error) {
	api := &SpecialDay{
		Subtype: NationalHoliday,
		Params:  params,
	}
	return request(ctx, api)
}

func ListAnniversaries(ctx context.Context, params Parameters) (*Response, error) {
	api := &SpecialDay{
		Subtype: Anniversary,
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

	resp := &Response{}
	if err := json.Unmarshal(data, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
