// 특일 정보 조회 API
//
// https://www.data.go.kr/data/15012690/openapi.do
package specialday

import (
	"encoding/json"

	"github.com/jh1104/publicapi"
)

// 특일 종류.
type Kind string

const (
	// 국경일 (어린이 날, 광복절, 개천절 등)
	KindNationalHoliday Kind = "01"
	// 기념일 (의병의 날, 정보보호의 날, 4·19 혁명 기념일 등)
	KindAnniversary Kind = "02"
	// 24절기 (청명, 경칩, 하지 등)
	KindSolarTerm Kind = "03"
	// 잡절 (단오, 한식 등)
	KindMiscellaneous Kind = "04"
)

type Response publicapi.Response[Item]

func (r *Response) UnmarshalJSON(data []byte) error {
	type alias Response
	raw := &struct {
		Response alias `json:"response"`
	}{}
	if err := json.Unmarshal(data, raw); err != nil {
		return err
	}
	*r = Response(raw.Response)
	return nil
}

type Item struct {
	// 20060102 포맷의 날짜.
	Date int `json:"locdate"`
	// 같은 날짜에 여러 항목이 있는 경우 구분하기 위한 시퀀스.
	Seq int `json:"seq"`
	// 특일 종류.
	DateKind Kind `json:"dateKind"`
	// 공공기관 휴일여부.
	IsHoliday string `json:"isHoliday"`
	// 특일 이름.
	Name string `json:"dateName"`
}
