package specialday_test

import (
	"context"
	"os"
	"testing"

	"github.com/jh1104/publicapi/specialday"
)

func TestRequest(t *testing.T) {
	if os.Getenv("INTEGRATION_TEST") != "on" {
		t.Skip("skipping integration test")
	}

	key, ok := os.LookupEnv("PUBLIC_API_SERVICE_KEY")
	if !ok {
		t.Fatal("PUBLIC_API_SERVICE_KEY environment variable is required")
	}

	client := specialday.NewClient(key)

	tests := []struct {
		name      string
		params    specialday.Parameters
		f         func(context.Context, specialday.Parameters) (*specialday.Response, error)
		wantItems int
	}{
		{
			name:      "2025-05 공휴일 조회",
			params:    specialday.Parameters{2025, 05, 10, 1},
			f:         client.ListHolidays,
			wantItems: 3,
		},
		{
			name:      "2025-05 국경일 조회",
			params:    specialday.Parameters{2025, 05, 10, 1},
			f:         client.ListNationalHolidays,
			wantItems: 3,
		},
		{
			name:      "2025-05 기념일 조회",
			params:    specialday.Parameters{2025, 05, 20, 1},
			f:         client.ListAnniversaries,
			wantItems: 15,
		},
		{
			name:      "2025-10 공휴일 조회",
			params:    specialday.Parameters{2025, 10, 10, 1},
			f:         client.ListHolidays,
			wantItems: 6,
		},
		{
			name:      "2025-10 국경일 조회",
			params:    specialday.Parameters{2025, 10, 10, 1},
			f:         client.ListNationalHolidays,
			wantItems: 6,
		},
		{
			name:      "2025-10 기념일 조회",
			params:    specialday.Parameters{2025, 10, 20, 1},
			f:         client.ListAnniversaries,
			wantItems: 14,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := tt.f(context.Background(), tt.params)
			if err != nil {
				t.Fatalf("")
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
