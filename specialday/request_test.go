package specialday_test

import (
	"context"
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
	specialday.SetDefaultClient(client)

	tests := []struct {
		name      string
		year      int
		month     int
		fn        func(ctx context.Context, year, month int) (*specialday.Response, error)
		wantItems int
	}{
		{
			name:      "2025-05 공휴일 조회",
			year:      2025,
			month:     5,
			fn:        specialday.ListHolidays,
			wantItems: 3,
		},
		{
			name:      "2025-05 국경일 조회",
			year:      2025,
			month:     5,
			fn:        specialday.ListNationalHolidays,
			wantItems: 3,
		},
		{
			name:      "2025-05 기념일 조회",
			year:      2025,
			month:     5,
			fn:        specialday.ListAnniversaries,
			wantItems: 10,
		},
		{
			name:      "2025-10 공휴일 조회",
			year:      2025,
			month:     10,
			fn:        specialday.ListHolidays,
			wantItems: 6,
		},
		{
			name:      "2025-10 국경일 조회",
			year:      2025,
			month:     10,
			fn:        specialday.ListNationalHolidays,
			wantItems: 6,
		},
		{
			name:      "2025-10 기념일 조회",
			year:      2025,
			month:     10,
			fn:        specialday.ListAnniversaries,
			wantItems: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := tt.fn(context.Background(), tt.year, tt.month)
			if err != nil {
				t.Fatalf("failed to request API: %v", err)
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
