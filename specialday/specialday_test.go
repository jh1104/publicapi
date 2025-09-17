package specialday_test

import (
	"net/url"
	"testing"

	"github.com/jh1104/publicapi"
	"github.com/jh1104/publicapi/specialday"
)

func TestNewSpecialDay(t *testing.T) {
	tests := []struct {
		name    string
		subtype specialday.Subtype
		params  specialday.Parameters
		want    specialday.SpecialDay
	}{
		{
			name:    "공휴일 API 생성",
			subtype: specialday.Holiday,
			params:  specialday.Parameters{Year: 2025, Month: 5, NumberOfRows: 10, PageNo: 1},
			want: specialday.SpecialDay{
				Subtype: specialday.Holiday,
				Params:  specialday.Parameters{Year: 2025, Month: 5, NumberOfRows: 10, PageNo: 1},
			},
		},
		{
			name:    "국경일 API 생성",
			subtype: specialday.NationalHoliday,
			params:  specialday.Parameters{Year: 2024, Month: 12, NumberOfRows: 20, PageNo: 2},
			want: specialday.SpecialDay{
				Subtype: specialday.NationalHoliday,
				Params:  specialday.Parameters{Year: 2024, Month: 12, NumberOfRows: 20, PageNo: 2},
			},
		},
		{
			name:    "기념일 API 생성",
			subtype: specialday.Anniversary,
			params:  specialday.Parameters{Year: 2023, Month: 1, NumberOfRows: 5, PageNo: 3},
			want: specialday.SpecialDay{
				Subtype: specialday.Anniversary,
				Params:  specialday.Parameters{Year: 2023, Month: 1, NumberOfRows: 5, PageNo: 3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := specialday.NewSpecialDay(tt.subtype, tt.params)

			if got.Subtype != tt.want.Subtype {
				t.Errorf("want %v, got %v", tt.want.Subtype, got.Subtype)
			}

			if got.Params != tt.want.Params {
				t.Errorf("want %v, got %v", tt.want.Params, got.Params)
			}
		})
	}
}

func TestSpecialDay_URL(t *testing.T) {
	tests := []struct {
		name       string
		api        publicapi.API
		serviceKey string
		wantParams map[string]string
	}{
		{
			name: "공휴일 API URL 생성",
			api: specialday.NewSpecialDay(
				specialday.Holiday,
				specialday.Parameters{2025, 05, 10, 1},
			),
			serviceKey: "holiday123",
			wantParams: map[string]string{
				"serviceKey": "holiday123",
				"solYear":    "2025",
				"solMonth":   "05",
				"numOfRows":  "10",
				"pageNo":     "1",
				"_type":      "json",
			},
		},
		{
			name: "국경일 URL 생성",
			api: specialday.NewSpecialDay(
				specialday.NationalHoliday,
				specialday.Parameters{2024, 12, 20, 2},
			),
			serviceKey: "national456",
			wantParams: map[string]string{
				"serviceKey": "national456",
				"solYear":    "2024",
				"solMonth":   "12",
				"numOfRows":  "20",
				"pageNo":     "2",
				"_type":      "json",
			},
		},
		{
			name: "기념일 URL 생성",
			api: specialday.NewSpecialDay(
				specialday.Anniversary,
				specialday.Parameters{2023, 1, 5, 3},
			),
			serviceKey: "anniversary789",
			wantParams: map[string]string{
				"serviceKey": "anniversary789",
				"solYear":    "2023",
				"solMonth":   "01",
				"numOfRows":  "5",
				"pageNo":     "3",
				"_type":      "json",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotURL := tt.api.URL(tt.serviceKey)
			urlValue, err := url.Parse(gotURL)
			if err != nil {
				t.Fatalf("failed to parse URL: %v", err)
			}

			// URL 쿼리 파라미터 확인
			for key, value := range tt.wantParams {
				if gotValue := urlValue.Query().Get(key); gotValue != value {
					t.Errorf("want %s = %s, got %s", key, value, gotValue)
				}
			}
		})
	}
}
