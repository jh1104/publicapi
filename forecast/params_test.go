package forecast_test

import (
	"testing"
	"time"

	"github.com/jh1104/publicapi/forecast"
)

func TestNextPage(t *testing.T) {
	tests := []struct {
		name  string
		input forecast.Parameters
		want  forecast.Parameters
	}{
		{
			name:  "1 페이지에서 다음 페이지로",
			input: forecast.Parameters{NumberOfRows: 10, PageNo: 1},
			want:  forecast.Parameters{NumberOfRows: 10, PageNo: 2},
		},
		{
			name:  "5 페이지에서 다음 페이지로",
			input: forecast.Parameters{NumberOfRows: 20, PageNo: 5},
			want:  forecast.Parameters{NumberOfRows: 20, PageNo: 6},
		},
		{
			name:  "0 페이지에서 다음 페이지로",
			input: forecast.Parameters{NumberOfRows: 15, PageNo: 0},
			want:  forecast.Parameters{NumberOfRows: 15, PageNo: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.NextPage()

			if got.PageNo != tt.want.PageNo {
				t.Errorf("want %d, got %d", tt.want.PageNo, got.PageNo)
			}

			if got.NumberOfRows != tt.want.NumberOfRows {
				t.Errorf("want %d, got %d", tt.want.NumberOfRows, got.NumberOfRows)
			}

			// 원본 파라미터가 변경되지 않았는지 확인
			if tt.input.PageNo != tt.want.PageNo-1 {
				t.Errorf("want no change in original PageNo, got %d", tt.input.PageNo)
			}
		})
	}
}

func TestBaseForUltraShortTermForecast(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		wantDate string
		wantTime string
	}{
		{"2025-01-01 00:00", time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), "20241231", "2330"},
		{"2025-01-01 00:40", time.Date(2025, 1, 1, 0, 40, 0, 0, time.UTC), "20250101", "0030"},
		{"2025-01-01 23:30", time.Date(2025, 1, 1, 23, 30, 0, 0, time.UTC), "20250101", "2330"},
		{"2025-01-01 23:40", time.Date(2025, 1, 1, 23, 40, 0, 0, time.UTC), "20250101", "2330"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDate, gotTime := forecast.BaseForUltraShortTermForecast(tt.input)
			if gotDate != tt.wantDate {
				t.Errorf("want date %s, got %s", tt.wantDate, gotDate)
			}
			if gotTime != tt.wantTime {
				t.Errorf("want time %s, got %s", tt.wantTime, gotTime)
			}
		})
	}
}
