package forecast

import (
	"fmt"
	"time"
)

type Parameters struct {
	// 20250101 포맷의 발표 일자.
	BaseDate string
	// 0930 포맷의 발표 시간.
	//
	// 초단기 예보는 매 시간 30분마다 발표한다.
	BaseTime string
	// 예보 지점의 X 좌표.
	NX int
	// 예보 지점의 Y 좌표.
	NY int

	// 한 페이지의 아이템 수.
	NumberOfRows int
	// 페이지 번호.
	PageNo int
}

func NewParameters(nx, ny int) Parameters {
	baseDate, baseTime := BaseForUltraShortTermForecast(time.Now())
	return Parameters{
		BaseDate:     baseDate,
		BaseTime:     baseTime,
		NX:           nx,
		NY:           ny,
		NumberOfRows: 10,
		PageNo:       1,
	}
}

// 현재 Parameters에서 페이지 번호를 1 증가시킨 새로운 Parameters를 반환합니다.
func (p Parameters) NextPage() Parameters {
	p.PageNo++
	return p
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
