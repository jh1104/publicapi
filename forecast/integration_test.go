package forecast_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/jh1104/publicapi/forecast"
)

func TestRequest(t *testing.T) {
	if os.Getenv("INTEGRATION_TEST") != "on" {
		t.Skip("skipping integration test")
	}

	key, ok := os.LookupEnv("PUBLIC_API_SERVICE_KEY")
	if !ok {
		t.Fatal("PUBLIC_API_SERVICE_KEY environment variable is required")
	}

	// 현재 시간 기준으로 baseDate와 baseTime 설정한다.
	baseDate, baseTime := forecast.BaseForUltraShortTermForecast(time.Now())

	client := forecast.NewClient(key)

	tests := []struct {
		name      string
		params    forecast.Parameters
		fn        func(context.Context, forecast.Parameters) (*forecast.Response, error)
		wantItems int
	}{
		{
			name:      "서울시청 좌표 초단기예보 조회",
			params:    forecast.Parameters{baseDate, baseTime, 60, 127, 5, 1},
			fn:        client.GetUltraShortTermForecast,
			wantItems: 5,
		},
		{
			name:      "부산시청 좌표 초단기예보 조회",
			params:    forecast.Parameters{baseDate, baseTime, 98, 76, 10, 1},
			fn:        client.GetUltraShortTermForecast,
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
			if resp.Body == nil {
				t.Fatal("want Body not nil, got nil")
			}
			if len(resp.Body.Data.Items) != tt.wantItems {
				t.Errorf("want %d items, got %d", tt.wantItems, len(resp.Body.Data.Items))
			}
		})
	}
}
