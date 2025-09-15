package forecast_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/jh1104/publicapi/forecast"
)

func TestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "초단기 예보 응답",
			input: `{ "response": { "header": { "resultCode": "00", "resultMsg": "NORMAL_SERVICE" }, "body": { "dataType": "JSON", "items": { "item": [ { "baseDate": "20250915", "baseTime": "1730", "category": "LGT", "fcstDate": "20250915", "fcstTime": "1800", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "LGT", "fcstDate": "20250915", "fcstTime": "1900", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "LGT", "fcstDate": "20250915", "fcstTime": "2000", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "LGT", "fcstDate": "20250915", "fcstTime": "2100", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "LGT", "fcstDate": "20250915", "fcstTime": "2200", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "LGT", "fcstDate": "20250915", "fcstTime": "2300", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "PTY", "fcstDate": "20250915", "fcstTime": "1800", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "PTY", "fcstDate": "20250915", "fcstTime": "1900", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "PTY", "fcstDate": "20250915", "fcstTime": "2000", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "PTY", "fcstDate": "20250915", "fcstTime": "2100", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "PTY", "fcstDate": "20250915", "fcstTime": "2200", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "PTY", "fcstDate": "20250915", "fcstTime": "2300", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "RN1", "fcstDate": "20250915", "fcstTime": "1800", "fcstValue": "강수없음", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "RN1", "fcstDate": "20250915", "fcstTime": "1900", "fcstValue": "강수없음", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "RN1", "fcstDate": "20250915", "fcstTime": "2000", "fcstValue": "강수없음", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "RN1", "fcstDate": "20250915", "fcstTime": "2100", "fcstValue": "강수없음", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "RN1", "fcstDate": "20250915", "fcstTime": "2200", "fcstValue": "강수없음", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "RN1", "fcstDate": "20250915", "fcstTime": "2300", "fcstValue": "강수없음", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "SKY", "fcstDate": "20250915", "fcstTime": "1800", "fcstValue": "3", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "SKY", "fcstDate": "20250915", "fcstTime": "1900", "fcstValue": "1", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "SKY", "fcstDate": "20250915", "fcstTime": "2000", "fcstValue": "3", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "SKY", "fcstDate": "20250915", "fcstTime": "2100", "fcstValue": "3", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "SKY", "fcstDate": "20250915", "fcstTime": "2200", "fcstValue": "4", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "SKY", "fcstDate": "20250915", "fcstTime": "2300", "fcstValue": "1", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "T1H", "fcstDate": "20250915", "fcstTime": "1800", "fcstValue": "26", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "T1H", "fcstDate": "20250915", "fcstTime": "1900", "fcstValue": "25", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "T1H", "fcstDate": "20250915", "fcstTime": "2000", "fcstValue": "25", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "T1H", "fcstDate": "20250915", "fcstTime": "2100", "fcstValue": "25", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "T1H", "fcstDate": "20250915", "fcstTime": "2200", "fcstValue": "25", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "T1H", "fcstDate": "20250915", "fcstTime": "2300", "fcstValue": "25", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "REH", "fcstDate": "20250915", "fcstTime": "1800", "fcstValue": "75", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "REH", "fcstDate": "20250915", "fcstTime": "1900", "fcstValue": "85", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "REH", "fcstDate": "20250915", "fcstTime": "2000", "fcstValue": "90", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "REH", "fcstDate": "20250915", "fcstTime": "2100", "fcstValue": "90", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "REH", "fcstDate": "20250915", "fcstTime": "2200", "fcstValue": "90", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "REH", "fcstDate": "20250915", "fcstTime": "2300", "fcstValue": "90", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "UUU", "fcstDate": "20250915", "fcstTime": "1800", "fcstValue": "-0.1", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "UUU", "fcstDate": "20250915", "fcstTime": "1900", "fcstValue": "-0.5", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "UUU", "fcstDate": "20250915", "fcstTime": "2000", "fcstValue": "-0.7", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "UUU", "fcstDate": "20250915", "fcstTime": "2100", "fcstValue": "-0.1", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "UUU", "fcstDate": "20250915", "fcstTime": "2200", "fcstValue": "0.2", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "UUU", "fcstDate": "20250915", "fcstTime": "2300", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "VVV", "fcstDate": "20250915", "fcstTime": "1800", "fcstValue": "0.1", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "VVV", "fcstDate": "20250915", "fcstTime": "1900", "fcstValue": "0.7", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "VVV", "fcstDate": "20250915", "fcstTime": "2000", "fcstValue": "0.4", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "VVV", "fcstDate": "20250915", "fcstTime": "2100", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "VVV", "fcstDate": "20250915", "fcstTime": "2200", "fcstValue": "-0.2", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "VVV", "fcstDate": "20250915", "fcstTime": "2300", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "VEC", "fcstDate": "20250915", "fcstTime": "1800", "fcstValue": "112", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "VEC", "fcstDate": "20250915", "fcstTime": "1900", "fcstValue": "136", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "VEC", "fcstDate": "20250915", "fcstTime": "2000", "fcstValue": "118", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "VEC", "fcstDate": "20250915", "fcstTime": "2100", "fcstValue": "102", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "VEC", "fcstDate": "20250915", "fcstTime": "2200", "fcstValue": "318", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "VEC", "fcstDate": "20250915", "fcstTime": "2300", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "WSD", "fcstDate": "20250915", "fcstTime": "1800", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "WSD", "fcstDate": "20250915", "fcstTime": "1900", "fcstValue": "1", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "WSD", "fcstDate": "20250915", "fcstTime": "2000", "fcstValue": "1", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "WSD", "fcstDate": "20250915", "fcstTime": "2100", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "WSD", "fcstDate": "20250915", "fcstTime": "2200", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "WSD", "fcstDate": "20250915", "fcstTime": "2300", "fcstValue": "0", "nx": 60, "ny": 123 } ] }, "pageNo": 1, "numOfRows": 60, "totalCount": 60 } } }`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &forecast.Response{}
			if err := json.Unmarshal([]byte(tt.input), resp); err != nil {
				t.Fatalf("failed to unmarshal %v", err)
			}

			if resp.Header.Code == "" {
				t.Errorf("want Header.Code non-empty, got empty")
			}
			if resp.Header.Message == "" {
				t.Errorf("want Header.Message non-empty, got empty")
			}

			if resp.Body == nil {
				t.Fatalf("want Body not nil, got nil")
			}
			if resp.Body.Page <= 0 {
				t.Errorf("want Body.Page > 0, got %d", resp.Body.Page)
			}
			if resp.Body.Rows <= 0 {
				t.Errorf("want Body.Rows > 0, got %d", resp.Body.Rows)
			}
			if resp.Body.Total < 0 {
				t.Errorf("want Body.Total >= 0, got %d", resp.Body.Total)
			}

			for _, item := range resp.Body.Data.Items {
				if item.BaseDate == "" {
					t.Errorf("want Item.BaseDate non-empty, got empty")
				}
				if item.BaseTime == "" {
					t.Errorf("want Item.BaseTime non-empty, got empty")
				}
				if item.Date == "" {
					t.Errorf("want Item.Date non-empty, got empty")
				}
				if item.Time == "" {
					t.Errorf("want Item.Time non-empty, got empty")
				}
				if item.Category == "" {
					t.Errorf("want Item.Category non-empty, got empty")
				}
				if item.Value == "" {
					t.Errorf("want Item.Value non-empty, got empty")
				}
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
