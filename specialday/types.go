package specialday

// 특일 정보 조회 API의 세부 분류.
type Subtype int

const (
	// 공휴일 조회 API.
	Holiday Subtype = iota + 1
	// 국경일 조회 API.
	NationalHoliday
	// 기념일 조회 API.
	Anniversary
)

// Subtype에 해당하는 API 경로를 한다.
func (a Subtype) Path() string {
	switch a {
	case Holiday:
		return "B090041/openapi/service/SpcdeInfoService/getRestDeInfo"
	case NationalHoliday:
		return "B090041/openapi/service/SpcdeInfoService/getHoliDeInfo"
	case Anniversary:
		return "B090041/openapi/service/SpcdeInfoService/getAnniversaryInfo"
	default:
		return ""
	}
}

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
