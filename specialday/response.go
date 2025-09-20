package specialday

import (
	"encoding/json"

	"github.com/jh1104/publicapi"
)

// 특일 조회 API의 응답.
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

// 특일 조회 API의 응답 데이터.
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
