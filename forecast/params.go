package forecast

type Parameters struct {
	// 20250101 포맷의 발표 일자.
	BaseDate string `json:"baseDate"`
	// 0930 포맷의 발표 시간.
	//
	// 초단기 예보는 매 시간 30분마다 발표한다.
	BaseTime string `json:"baseTime"`
	// 예보 지점의 X 좌표.
	NX int
	// 예보 지점의 Y 좌표.
	NY int

	// 한 페이지 결과 수.
	NumberOfRows int
	// 페이지 번호.
	PageNo int
}
