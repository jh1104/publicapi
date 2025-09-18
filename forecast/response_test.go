package forecast_test

import (
	"encoding/json"
	"testing"

	"github.com/jh1104/publicapi/forecast"
)

func TestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantPage  int
		wantRows  int
		wantTotal int
		wantItems int
	}{
		{
			name:      "초단기 예보 응답",
			input:     `{ "response": { "header": { "resultCode": "00", "resultMsg": "NORMAL_SERVICE" }, "body": { "dataType": "JSON", "items": { "item": [ { "baseDate": "20250915", "baseTime": "1730", "category": "LGT", "fcstDate": "20250915", "fcstTime": "1800", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "LGT", "fcstDate": "20250915", "fcstTime": "1900", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "LGT", "fcstDate": "20250915", "fcstTime": "2000", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "LGT", "fcstDate": "20250915", "fcstTime": "2100", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "LGT", "fcstDate": "20250915", "fcstTime": "2200", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "LGT", "fcstDate": "20250915", "fcstTime": "2300", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "PTY", "fcstDate": "20250915", "fcstTime": "1800", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "PTY", "fcstDate": "20250915", "fcstTime": "1900", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "PTY", "fcstDate": "20250915", "fcstTime": "2000", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "PTY", "fcstDate": "20250915", "fcstTime": "2100", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "PTY", "fcstDate": "20250915", "fcstTime": "2200", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "PTY", "fcstDate": "20250915", "fcstTime": "2300", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "RN1", "fcstDate": "20250915", "fcstTime": "1800", "fcstValue": "강수없음", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "RN1", "fcstDate": "20250915", "fcstTime": "1900", "fcstValue": "강수없음", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "RN1", "fcstDate": "20250915", "fcstTime": "2000", "fcstValue": "강수없음", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "RN1", "fcstDate": "20250915", "fcstTime": "2100", "fcstValue": "강수없음", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "RN1", "fcstDate": "20250915", "fcstTime": "2200", "fcstValue": "강수없음", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "RN1", "fcstDate": "20250915", "fcstTime": "2300", "fcstValue": "강수없음", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "SKY", "fcstDate": "20250915", "fcstTime": "1800", "fcstValue": "3", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "SKY", "fcstDate": "20250915", "fcstTime": "1900", "fcstValue": "1", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "SKY", "fcstDate": "20250915", "fcstTime": "2000", "fcstValue": "3", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "SKY", "fcstDate": "20250915", "fcstTime": "2100", "fcstValue": "3", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "SKY", "fcstDate": "20250915", "fcstTime": "2200", "fcstValue": "4", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "SKY", "fcstDate": "20250915", "fcstTime": "2300", "fcstValue": "1", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "T1H", "fcstDate": "20250915", "fcstTime": "1800", "fcstValue": "26", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "T1H", "fcstDate": "20250915", "fcstTime": "1900", "fcstValue": "25", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "T1H", "fcstDate": "20250915", "fcstTime": "2000", "fcstValue": "25", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "T1H", "fcstDate": "20250915", "fcstTime": "2100", "fcstValue": "25", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "T1H", "fcstDate": "20250915", "fcstTime": "2200", "fcstValue": "25", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "T1H", "fcstDate": "20250915", "fcstTime": "2300", "fcstValue": "25", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "REH", "fcstDate": "20250915", "fcstTime": "1800", "fcstValue": "75", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "REH", "fcstDate": "20250915", "fcstTime": "1900", "fcstValue": "85", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "REH", "fcstDate": "20250915", "fcstTime": "2000", "fcstValue": "90", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "REH", "fcstDate": "20250915", "fcstTime": "2100", "fcstValue": "90", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "REH", "fcstDate": "20250915", "fcstTime": "2200", "fcstValue": "90", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "REH", "fcstDate": "20250915", "fcstTime": "2300", "fcstValue": "90", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "UUU", "fcstDate": "20250915", "fcstTime": "1800", "fcstValue": "-0.1", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "UUU", "fcstDate": "20250915", "fcstTime": "1900", "fcstValue": "-0.5", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "UUU", "fcstDate": "20250915", "fcstTime": "2000", "fcstValue": "-0.7", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "UUU", "fcstDate": "20250915", "fcstTime": "2100", "fcstValue": "-0.1", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "UUU", "fcstDate": "20250915", "fcstTime": "2200", "fcstValue": "0.2", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "UUU", "fcstDate": "20250915", "fcstTime": "2300", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "VVV", "fcstDate": "20250915", "fcstTime": "1800", "fcstValue": "0.1", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "VVV", "fcstDate": "20250915", "fcstTime": "1900", "fcstValue": "0.7", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "VVV", "fcstDate": "20250915", "fcstTime": "2000", "fcstValue": "0.4", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "VVV", "fcstDate": "20250915", "fcstTime": "2100", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "VVV", "fcstDate": "20250915", "fcstTime": "2200", "fcstValue": "-0.2", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "VVV", "fcstDate": "20250915", "fcstTime": "2300", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "VEC", "fcstDate": "20250915", "fcstTime": "1800", "fcstValue": "112", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "VEC", "fcstDate": "20250915", "fcstTime": "1900", "fcstValue": "136", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "VEC", "fcstDate": "20250915", "fcstTime": "2000", "fcstValue": "118", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "VEC", "fcstDate": "20250915", "fcstTime": "2100", "fcstValue": "102", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "VEC", "fcstDate": "20250915", "fcstTime": "2200", "fcstValue": "318", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "VEC", "fcstDate": "20250915", "fcstTime": "2300", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "WSD", "fcstDate": "20250915", "fcstTime": "1800", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "WSD", "fcstDate": "20250915", "fcstTime": "1900", "fcstValue": "1", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "WSD", "fcstDate": "20250915", "fcstTime": "2000", "fcstValue": "1", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "WSD", "fcstDate": "20250915", "fcstTime": "2100", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "WSD", "fcstDate": "20250915", "fcstTime": "2200", "fcstValue": "0", "nx": 60, "ny": 123 }, { "baseDate": "20250915", "baseTime": "1730", "category": "WSD", "fcstDate": "20250915", "fcstTime": "2300", "fcstValue": "0", "nx": 60, "ny": 123 } ] }, "pageNo": 1, "numOfRows": 60, "totalCount": 60 } } }`,
			wantPage:  1,
			wantRows:  60,
			wantTotal: 60,
			wantItems: 60,
		},
		{
			name:      "빈 응답",
			input:     `{ "response": { "header": { "resultCode": "03", "resultMsg": "NO_DATA" } } }`,
			wantPage:  0,
			wantRows:  0,
			wantTotal: 0,
			wantItems: 0,
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

			if resp.Body.Page != tt.wantPage {
				t.Errorf("want Body.Page %d, got %d", tt.wantPage, resp.Body.Page)
			}
			if resp.Body.Rows != tt.wantRows {
				t.Errorf("want Body.Rows %d, got %d", tt.wantRows, resp.Body.Rows)
			}
			if resp.Body.Total != tt.wantTotal {
				t.Errorf("want Body.Total %d, got %d", tt.wantTotal, resp.Body.Total)
			}

			if tt.wantItems != len(resp.Body.Data.Items) {
				t.Errorf("want %d items, got %d", tt.wantItems, len(resp.Body.Data.Items))
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
