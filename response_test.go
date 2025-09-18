package publicapi_test

import (
	"encoding/json"
	"testing"

	"github.com/jh1104/publicapi"
)

func TestUnmarshalJSON(t *testing.T) {
	type TestItem struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	tests := []struct {
		name    string
		input   string
		want    *publicapi.Body[TestItem]
		wantErr bool
	}{
		{
			name:  "정상적인 Body",
			input: `{ "pageNo": 1, "numOfRows": 10, "totalCount": 100, "items": { "item": [ {"id": 1, "name": "test1"}, {"id": 2, "name": "test2"} ] } }`,
			want: &publicapi.Body[TestItem]{
				Page:  1,
				Rows:  10,
				Total: 100,
				Data: publicapi.BodyData[TestItem]{
					Items: []TestItem{
						{ID: 1, Name: "test1"},
						{ID: 2, Name: "test2"},
					},
				},
			},
			wantErr: false,
		},
		{
			name:  "items 필드가 없는 Body",
			input: `{ "pageNo": 1, "numOfRows": 0, "totalCount": 0 }`,
			want: &publicapi.Body[TestItem]{
				Page:  1,
				Rows:  0,
				Total: 0,
				Data:  publicapi.BodyData[TestItem]{},
			},
			wantErr: false,
		},
		{
			name:  "items 필드가 빈 배열인 Body",
			input: `{ "pageNo": 1, "numOfRows": 0, "totalCount": 0, "items": { "item": [] } }`,
			want: &publicapi.Body[TestItem]{
				Page:  1,
				Rows:  0,
				Total: 0,
				Data:  publicapi.BodyData[TestItem]{},
			},
			wantErr: false,
		},
		{
			name:  "items 필드가 빈 문자열인 Body",
			input: `{ "pageNo": 1, "numOfRows": 0, "totalCount": 0, "items": "" }`,
			want: &publicapi.Body[TestItem]{
				Page:  1,
				Rows:  0,
				Total: 0,
				Data:  publicapi.BodyData[TestItem]{},
			},
			wantErr: false,
		},
		{
			name:    "items 필드가 빈 객체인 Body",
			input:   `{ "pageNo": 1, "numOfRows": 0, "totalCount": 0, "items": { "item": {} } }`,
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body publicapi.Body[TestItem]
			err := json.Unmarshal([]byte(tt.input), &body)

			if tt.wantErr {
				if err == nil {
					t.Errorf("want error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("want no error, got %v", err)
			}

			// 각 필드 비교
			if body.Page != tt.want.Page {
				t.Errorf("want pages %d, got %d", tt.want.Page, body.Page)
			}
			if body.Rows != tt.want.Rows {
				t.Errorf("want rows %d, got %d", tt.want.Rows, body.Rows)
			}
			if body.Total != tt.want.Total {
				t.Errorf("want total %d, got %d", tt.want.Total, body.Total)
			}

			// Items 비교
			if len(body.Data.Items) != len(tt.want.Data.Items) {
				t.Fatalf("want items %d, got %d", len(tt.want.Data.Items), len(body.Data.Items))
			}

			for i, item := range body.Data.Items {
				if item != tt.want.Data.Items[i] {
					t.Errorf("want item %v, got %v", tt.want.Data.Items[i], item)
				}
			}
		})
	}
}
