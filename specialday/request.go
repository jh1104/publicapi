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

func ListHolidays(ctx context.Context, year, month int) (*Response, error) {
	api := &SpecialDay{
		Subtype: Holiday,
		Params:  NewParameters(year, month),
	}
	return request(ctx, api)
}

func ListNationalHolidays(ctx context.Context, year, month int) (*Response, error) {
	api := &SpecialDay{
		Subtype: NationalHoliday,
		Params:  NewParameters(year, month),
	}
	return request(ctx, api)
}

func ListAnniversaries(ctx context.Context, year, month int) (*Response, error) {
	api := &SpecialDay{
		Subtype: Anniversary,
		Params:  NewParameters(year, month),
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
