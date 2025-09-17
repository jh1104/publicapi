package specialday_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/jh1104/publicapi"
	"github.com/jh1104/publicapi/specialday"
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

	tests := []struct {
		name      string
		api       publicapi.API
		wantItems int
	}{
		{
			name:      "2025-05 공휴일 조회",
			api:       specialday.NewSpecialDay(specialday.Holiday, specialday.Parameters{2025, 05, 10, 1}),
			wantItems: 3,
		},
		{
			name:      "2025-05 국경일 조회",
			api:       specialday.NewSpecialDay(specialday.NationalHoliday, specialday.Parameters{2025, 05, 10, 1}),
			wantItems: 3,
		},
		{
			name:      "2025-05 기념일 조회",
			api:       specialday.NewSpecialDay(specialday.Anniversary, specialday.Parameters{2025, 05, 20, 1}),
			wantItems: 15,
		},
		{
			name:      "2025-10 공휴일 조회",
			api:       specialday.NewSpecialDay(specialday.Holiday, specialday.Parameters{2025, 10, 10, 1}),
			wantItems: 6,
		},
		{
			name:      "2025-10 국경일 조회",
			api:       specialday.NewSpecialDay(specialday.NationalHoliday, specialday.Parameters{2025, 10, 10, 1}),
			wantItems: 6,
		},
		{
			name:      "2025-10 기념일 조회",
			api:       specialday.NewSpecialDay(specialday.Anniversary, specialday.Parameters{2025, 10, 20, 1}),
			wantItems: 14,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := client.RequestAPI(context.Background(), tt.api)
			if err != nil {
				t.Fatalf("failed to request API: %v", err)
			}

			resp := &specialday.Response{}
			if err := json.Unmarshal(data, resp); err != nil {
				t.Fatalf("failed to unmarshal response: %v", err)
			}

			if resp.Header.Code != "00" {
				t.Fatalf("want code 00, got code=%s message=%s", resp.Header.Code, resp.Header.Message)
			}

			if resp.Body == nil {
				t.Fatal("want body not nil, got nil")
			}

			if len(resp.Body.Data.Items) != tt.wantItems {
				t.Fatalf("want items %d, got %d", tt.wantItems, len(resp.Body.Data.Items))
			}
		})
	}
}
