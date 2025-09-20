package specialday

// 특일 조회 API 파라미터.
type Parameters struct {
	// 조회 연도.
	Year int
	// 조회 월.
	Month int

	// 한 페이지의 아이템 수.
	NumberOfRows int
	// 페이지 번호.
	PageNo int
}

func NewParameters(year int, month int) Parameters {
	return Parameters{
		Year:         year,
		Month:        month,
		NumberOfRows: 10,
		PageNo:       1,
	}
}

// 현재 Parameters에서 페이지 번호를 1 증가시킨 새로운 Parameters를 반환한다.
func (p Parameters) NextPage() Parameters {
	p.PageNo++
	return p
}
