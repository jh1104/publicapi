package forecast_test

import (
	"context"
	"os"
	"testing"

	"github.com/jh1104/publicapi"
	"github.com/jh1104/publicapi/forecast"
)

func TestRequestAPI(t *testing.T) {
	if os.Getenv("INTEGRATION_TEST") != "on" {
		t.Skip("skipping integration test")
	}

	key, ok := os.LookupEnv("PUBLIC_API_SERVICE_KEY")
	if !ok {
		t.Fatal("PUBLIC_API_SERVICE_KEY environment variable is required")
	}

	client := publicapi.NewClient(key)
	forecast.SetDefaultClient(client)

	tests := []struct {
		name      string
		params    forecast.Parameters
		fn        func(context.Context, forecast.Parameters) (*forecast.Response, error)
		wantItems int
	}{
		{
			name:      "서울시청 좌표 초단기예보 조회",
			params:    forecast.NewParameters(60, 127),
			fn:        forecast.GetUltraShortTermForecast,
			wantItems: 10,
		},
		{
			name:      "부산시청 좌표 초단기예보 조회",
			params:    forecast.NewParameters(98, 76),
			fn:        forecast.GetUltraShortTermForecast,
			wantItems: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := tt.fn(context.Background(), tt.params)
			if err != nil {
				t.Fatalf("failed to request API: %v", err)
			}

			if resp.Header.Code != "00" {
				t.Errorf("want Header.Code '00', got %s message=%s", resp.Header.Code, resp.Header.Message)
			}
			if len(resp.Body.Data.Items) != tt.wantItems {
				t.Errorf("want %d items, got %d", tt.wantItems, len(resp.Body.Data.Items))
			}
		})
	}
}
