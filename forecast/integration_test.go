package forecast_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

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

	// 현재 시간 기준으로 baseDate와 baseTime 설정한다.
	baseDate, baseTime := forecast.BaseForUltraShortTermForecast(time.Now())

	client := publicapi.NewClient(key)

	tests := []struct {
		name      string
		api       publicapi.API
		wantItems int
	}{
		{
			name:      "서울시청 좌표 초단기예보 조회",
			api:       forecast.NewForecast(forecast.UltraShortTermForecast, forecast.Parameters{baseDate, baseTime, 60, 127, 5, 1}),
			wantItems: 5,
		},
		{
			name:      "부산시청 좌표 초단기예보 조회",
			api:       forecast.NewForecast(forecast.UltraShortTermForecast, forecast.Parameters{baseDate, baseTime, 98, 76, 10, 1}),
			wantItems: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := client.RequestAPI(context.Background(), tt.api)
			if err != nil {
				t.Fatalf("failed to request API: %v", err)
			}

			resp := &forecast.Response{}
			if err := json.Unmarshal(data, resp); err != nil {
				t.Fatalf("failed to unmarshal response: %v", err)
			}

			if resp.Header.Code != "00" {
				t.Errorf("want Header.Code '00', got %s message=%s", resp.Header.Code, resp.Header.Message)
			}
			if resp.Body == nil {
				t.Fatal("want Body not nil, got nil")
			}
			if len(resp.Body.Data.Items) != tt.wantItems {
				t.Errorf("want %d items, got %d", tt.wantItems, len(resp.Body.Data.Items))
			}
		})
	}
}
