package forecast_test

import (
	"testing"
	"time"

	"github.com/jh1104/publicapi"
	"github.com/jh1104/publicapi/forecast"
)

func TestNewForecast(t *testing.T) {
	baseDate, baseTime := forecast.BaseForUltraShortTermForecast(time.Now())

	tests := []struct {
		name    string
		subtype forecast.Subtype
		params  forecast.Parameters
		want    forecast.Forecast
	}{
		{
			name:    "초단기 예보 API 생성",
			subtype: forecast.UltraShortTermForecast,
			params:  forecast.Parameters{baseDate, baseTime, 123, 456, 10, 1},
			want: forecast.Forecast{
				Subtype: forecast.UltraShortTermForecast,
				Params:  forecast.Parameters{baseDate, baseTime, 123, 456, 10, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := forecast.NewForecast(tt.subtype, tt.params)

			if got.Subtype != tt.want.Subtype {
				t.Errorf("want %v, got %v", tt.want.Subtype, got.Subtype)
			}

			if got.Params != tt.want.Params {
				t.Errorf("want %v, got %v", tt.want.Params, got.Params)
			}
		})
	}
}

func TestForecast_URL(t *testing.T) {
	baseDate, baseTime := forecast.BaseForUltraShortTermForecast(time.Now())

	tests := []struct {
		name       string
		api        publicapi.API
		serviceKey string
		wantParams map[string]string
	}{
		{
			name: "초단기 예보 조회 API URL 생성",
			api: forecast.NewForecast(
				forecast.UltraShortTermForecast,
				forecast.Parameters{baseDate, baseTime, 123, 456, 10, 1},
			),
			serviceKey: "ultarshortterm123",
			wantParams: map[string]string{
				"serviceKey": "ultarshortterm123",
				"baseDate":   baseDate,
				"baseTime":   baseTime,
				"nx":         "123",
				"ny":         "456",
				"numOfRows":  "10",
				"pageNo":     "1",
				"dataType":   "JSON",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
