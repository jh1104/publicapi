package publicapi

import "encoding/json"

type Response[T any] struct {
	// 응답 헤더.
	Header Header `json:"header" xml:"header"`
	// 응답 본문.
	Body *Body[T] `json:"body,omitempty" xml:"body"`
}

type Header struct {
	// 결과 코드.
	Code string `json:"resultCode" xml:"resultCode"`
	// 결과 메시지.
	Message string `json:"resultMsg" xml:"resultMsg"`
}

type Body[T any] struct {
	// 페이지 번호.
	Page int `json:"pageNo" xml:"pageNo"`
	// 페이지 당 데이터 수.
	Rows int `json:"numOfRows" xml:"numOfRows"`
	// 전체 데이터 수.
	Total int `json:"totalCount" xml:"totalCount"`
	// 데이터 목록.
	Data BodyData[T] `json:"items" xml:"items"`
}

type BodyData[T any] struct {
	Items []T `json:"item" xml:"item"`
}

func (b *BodyData[T]) UnmarshalJSON(data []byte) error {
	if len(data) <= 2 && string(data) == `""` {
		return nil
	}
	type alias BodyData[T]
	raw := &alias{}
	if err := json.Unmarshal(data, raw); err != nil {
		return err
	}
	b.Items = raw.Items
	return nil
}
