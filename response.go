package publicapi

import "encoding/json"

// 공공데이터포털 API 응답.
type Response[T any] struct {
	// 응답 헤더.
	Header Header `json:"header" xml:"header"`
	// 응답 본문.
	Body Body[T] `json:"body" xml:"body"`
}

// 응답 헤더.
type Header struct {
	// 결과 코드.
	Code string `json:"resultCode" xml:"resultCode"`
	// 결과 메시지.
	Message string `json:"resultMsg" xml:"resultMsg"`
}

// 응답 본문.
// API 종류에 따라 Body가 없는 경우도 있다.
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

// 에러 응답 구조체.
type ErrorResponse struct {
	Header struct {
		ReasonCode ResultCode `xml:"returnReasonCode"`
	} `xml:"cmmMsgHeader"`
}

// 에러 원인 코드.
func (e *ErrorResponse) AsError() error {
	return e.Header.ReasonCode.AsError()
}
