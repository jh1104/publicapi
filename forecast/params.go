package forecast

import (
	"fmt"
	"slices"
	"time"
)

// 단기 예보 조회 API 파라미터.
type Parameters struct {
	// 20250101 포맷의 발표 일자.
	BaseDate string
	// 0930 포맷의 발표 시간.
	//
	// 초단기 예보는 매 시간 30분마다 발표하며 45분부터 조회할 수 있다.
	// 0030, 0130, 0230, ..., 2330 (1일 24회)
	//
	// 단기 예보는 매 3시간마다 발표하며 10분부터 조회할 수 있다.
	// 0200, 0500, 0800, 1100, 1400, 1700, 2000, 2300 (1일 8회)
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

// 현재 Parameters에서 페이지 번호를 1 증가시킨 새로운 Parameters를 반환한다.
func (p Parameters) NextPage() Parameters {
	p.PageNo++
	return p
}

// 주어진 시간에 가장 가까운 초단기 예보의 baseDate와 baseTime을 반환한다.
func BaseForUltraShortTermForecast(t time.Time) (baseDate string, baseTime string) {
	base := t

	// 초단기 예보는 매 시간 30분마다 발표하며 45분부터 조회할 수 있다.
	if base.Minute() < 45 {
		base = base.Add(-time.Hour)
	}

	baseDate = base.Format("20060102")
	baseTime = fmt.Sprintf("%02d30", base.Hour())

	return baseDate, baseTime
}

func BaseForShortTermForecast(t time.Time) (baseDate string, baseTime string) {
	base := t

	// 단기 예보는 매 3시간마다 발표하며 10분부터 조회할 수 있다.
	if base.Minute() < 10 {
		base = base.Add(-time.Hour)
	}
	hours := []int{2, 5, 8, 11, 14, 17, 20, 23}
	for !slices.Contains(hours, base.Hour()) {
		base = base.Add(-time.Hour)
	}

	baseDate = base.Format("20060102")
	baseTime = fmt.Sprintf("%02d00", base.Hour())

	return baseDate, baseTime
}
