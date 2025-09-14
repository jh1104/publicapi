package specialday

// 특일 조회 API 파라미터.
type Parameters struct {
	// 조회 연도.
	Year int
	// 조회 월.
	Month int

	// 한 페이지 결과 수.
	NumberOfRows int
	// 페이지 번호.
	PageNo int
}

func (p Parameters) NextPage() Parameters {
	p.PageNo++
	return p
}
