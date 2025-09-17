package specialday_test

import (
	"testing"

	"github.com/jh1104/publicapi/specialday"
)

func TestNewParameters(t *testing.T) {
	params := specialday.NewParameters(2025, 10)
	if params.Year != 2025 {
		t.Errorf("want Year 2025, got %d", params.Year)
	}
	if params.Month != 10 {
		t.Errorf("want Month 10, got %d", params.Month)
	}
}

func TestNextPage(t *testing.T) {
	tests := []struct {
		name  string
		input specialday.Parameters
		want  specialday.Parameters
	}{
		{
			name:  "1 페이지에서 다음 페이지로",
			input: specialday.Parameters{NumberOfRows: 10, PageNo: 1},
			want:  specialday.Parameters{NumberOfRows: 10, PageNo: 2},
		},
		{
			name:  "5 페이지에서 다음 페이지로",
			input: specialday.Parameters{NumberOfRows: 20, PageNo: 5},
			want:  specialday.Parameters{NumberOfRows: 20, PageNo: 6},
		},
		{
			name:  "0 페이지에서 다음 페이지로",
			input: specialday.Parameters{NumberOfRows: 15, PageNo: 0},
			want:  specialday.Parameters{NumberOfRows: 15, PageNo: 1},
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
				t.Errorf("want no change in original, got %d", tt.input.PageNo)
			}
		})
	}
}
