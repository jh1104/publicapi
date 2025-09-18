package specialday_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/jh1104/publicapi/specialday"
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
			name:      "2025년 5월 공휴일",
			input:     `{ "response": { "header": { "resultCode": "00", "resultMsg": "NORMAL SERVICE." }, "body": { "items": { "item": [ { "dateKind": "01", "dateName": "어린이날", "isHoliday": "Y", "locdate": 20250505, "seq": 1 }, { "dateKind": "01", "dateName": "부처님오신날", "isHoliday": "Y", "locdate": 20250505, "seq": 2 }, { "dateKind": "01", "dateName": "대체공휴일", "isHoliday": "Y", "locdate": 20250506, "seq": 1 } ] }, "numOfRows": 10, "pageNo": 1, "totalCount": 3 } } }`,
			wantPage:  1,
			wantRows:  10,
			wantTotal: 3,
			wantItems: 3,
		},
		{
			name:      "2025년 9월 공휴일",
			input:     `{ "response": { "header": { "resultCode": "00", "resultMsg": "NORMAL SERVICE." }, "body": { "items": "", "numOfRows": 10, "pageNo": 1, "totalCount": 0 } } }`,
			wantPage:  1,
			wantRows:  10,
			wantTotal: 0,
			wantItems: 0,
		},
		{
			name:      "2025년 10월 공휴일",
			input:     `{ "response": { "header": { "resultCode": "00", "resultMsg": "NORMAL SERVICE." }, "body": { "items": { "item": [ { "dateKind": "01", "dateName": "개천절", "isHoliday": "Y", "locdate": 20251003, "seq": 1 }, { "dateKind": "01", "dateName": "추석", "isHoliday": "Y", "locdate": 20251005, "seq": 1 }, { "dateKind": "01", "dateName": "추석", "isHoliday": "Y", "locdate": 20251006, "seq": 1 }, { "dateKind": "01", "dateName": "추석", "isHoliday": "Y", "locdate": 20251007, "seq": 1 }, { "dateKind": "01", "dateName": "대체공휴일", "isHoliday": "Y", "locdate": 20251008, "seq": 1 }, { "dateKind": "01", "dateName": "한글날", "isHoliday": "Y", "locdate": 20251009, "seq": 1 } ] }, "numOfRows": 10, "pageNo": 1, "totalCount": 6 } } }`,
			wantPage:  1,
			wantRows:  10,
			wantTotal: 6,
			wantItems: 6,
		},
		{
			name:      "빈 응답",
			input:     `{ "response": { "header": { "resultCode": "00", "resultMsg": "NORMAL SERVICE." }, "body": { "items": "", "numOfRows": 10, "pageNo": 1, "totalCount": 0 } } }`,
			wantPage:  1,
			wantRows:  10,
			wantTotal: 0,
			wantItems: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &specialday.Response{}
			err := json.Unmarshal([]byte(tt.input), resp)
			if err != nil {
				t.Fatalf("failed to unmarshal %v", err)
			}

			// 헤더 검증
			if resp.Header.Code == "" {
				t.Errorf("want Header.Code non-empty, got empty")
			}
			if resp.Header.Message == "" {
				t.Errorf("want Header.Message non-empty, got empty")
			}

			// 바디 검증
			if resp.Body.Page != tt.wantPage {
				t.Errorf("want Body.Page %d, got %d", tt.wantPage, resp.Body.Page)
			}
			if resp.Body.Rows != tt.wantRows {
				t.Errorf("want Body.Rows %d, got %d", tt.wantRows, resp.Body.Rows)
			}
			if resp.Body.Total != tt.wantTotal {
				t.Errorf("want Body.Total %d, got %d", tt.wantTotal, resp.Body.Total)
			}

			if len(resp.Body.Data.Items) != tt.wantItems {
				t.Errorf("want %d items, got %d", tt.wantItems, len(resp.Body.Data.Items))
			}

			for _, item := range resp.Body.Data.Items {
				if _, err := time.Parse("20060102", fmt.Sprintf("%d", item.Date)); err != nil {
					t.Errorf("want item.Date 20060102 format, got %d %v", item.Date, err)
				}
				if item.Seq < 1 {
					t.Errorf("want item.Seq > 0, got %d", item.Seq)
				}
				if item.Name == "" {
					t.Errorf("want item.Name non-empty, got empty")
				}
				if item.IsHoliday != "Y" && item.IsHoliday != "N" {
					t.Errorf("want item.IsHoliday Y or N, got %s", item.IsHoliday)
				}
				if item.DateKind == "" {
					t.Errorf("want item.DateKind non-empty, got empty")
				}
			}
		})
	}
}
