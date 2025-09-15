// 단기 예보 조회
//
// https://www.data.go.kr/data/15084084/openapi.do
package forecast

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/jh1104/publicapi"
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
	// 20250101 포맷의 발표 일자.
	BaseDate string `json:"baseDate"`
	// 0930 포맷의 발표 시간.
	//
	// 초단기 예보는 매 시간 30분마다 발표한다.
	BaseTime string `json:"baseTime"`
	// 20250101 포맷의 예측 일자.
	Date string `json:"fcstDate"`
	// 0930 포맷의 예측 시간.
	Time string `json:"fcstTime"`
	// 예보 값 분류.
	Category ForecastCategory `json:"category"`
	// 예보 값. Category에 따라 값이 달라진다.
	Value string `json:"fcstValue"`
	// 예보 지점의 X 좌표.
	NX int `json:"nx"`
	// 예보 지점의 Y 좌표.
	NY int `json:"ny"`
}

// 주어진 시간에 가장 가까운 초단기 예보의 baseDate와 baseTime을 반환한다.
func BaseForUltraShortTermForecast(t time.Time) (baseDate string, baseTime string) {
	base := t

	// 초단기예보는 매시간 30분에 발표되므로,
	// 현재 분이 30분 미만이면 이전 시간의 30분, 30분 이상이면 현재 시간의 30분을 사용한다.
	if base.Minute() < 30 {
		base = base.Add(-time.Hour)
	}

	baseDate = base.Format("20060102")
	baseTime = fmt.Sprintf("%02d30", base.Hour())

	return baseDate, baseTime
}
