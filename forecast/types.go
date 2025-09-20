package forecast

// 단기 예보 조회 API의 세부 분류.
type Subtype int

const (
	// 초단기 예보 조회
	UltraShortTermForecast Subtype = iota + 1
)

// Subtype에 해당하는 API 경로를 반환한다.
func (s Subtype) Path() string {
	switch s {
	case UltraShortTermForecast:
		return "1360000/VilageFcstInfoService_2.0/getUltraSrtFcst"
	default:
		return ""
	}
}

// 단기 예보 조회 시 사용되는 예보 구분 코드.
type ForecastCategory string

const (
	// 기온 (°C).
	CategoryTemperature ForecastCategory = "T1H"
	// 1시간 강수량 범주 (mm).
	CategoryRainfall ForecastCategory = "RN1"
	// 하늘상태 코드.
	CategorySky ForecastCategory = "SKY"
	// 동서바람성분 (m/s).
	CategoryEastWestWind ForecastCategory = "UUU"
	// 남북바람성분 (m/s).
	CategoryNorthSouthWind ForecastCategory = "VVV"
	// 습도 (%).
	CategoryHumidity ForecastCategory = "REH"
	// 강수형태 코드.
	CategoryPrecipitation ForecastCategory = "PTY"
	// 낙뢰 (kA).
	CategoryLightning ForecastCategory = "LGT"
	// 풍향 (deg).
	CategoryWindDirection ForecastCategory = "VEC"
	// 풍속 (m/s).
	CategoryWindSpeed ForecastCategory = "WSD"
)

func (f ForecastCategory) String() string {
	switch f {
	case CategoryTemperature:
		return "기온"
	case CategoryRainfall:
		return "1시간 강수량"
	case CategorySky:
		return "하늘 상태"
	case CategoryEastWestWind:
		return "동서바람성분"
	case CategoryNorthSouthWind:
		return "남북바람성분"
	case CategoryHumidity:
		return "습도"
	case CategoryPrecipitation:
		return "강수 형태"
	case CategoryLightning:
		return "낙뢰"
	case CategoryWindDirection:
		return "풍향"
	case CategoryWindSpeed:
		return "풍속"
	default:
		return string(f)
	}
}

// 강수형태 코드.
type PrecipitationCode string

const (
	// 없음
	PrecipitationCodeNone PrecipitationCode = "0"
	// 비
	PrecipitationCodeRain PrecipitationCode = "1"
	// 비/눈
	PrecipitationCodeRainSnow PrecipitationCode = "2"
	// 눈
	PrecipitationCodeSnow PrecipitationCode = "3"
	// 소나기
	PrecipitationCodeShower PrecipitationCode = "4"
	// 빗방울
	PrecipitationCodeDrizzle PrecipitationCode = "5"
	// 진눈깨비
	PrecipitationCodeSleet PrecipitationCode = "6"
	// 눈날림
	PrecipitationCodeSnowGrain PrecipitationCode = "7"
)

func (p PrecipitationCode) String() string {
	switch p {
	case PrecipitationCodeNone:
		return "없음"
	case PrecipitationCodeRain:
		return "비"
	case PrecipitationCodeRainSnow:
		return "비/눈"
	case PrecipitationCodeSnow:
		return "눈"
	case PrecipitationCodeShower:
		return "소나기"
	case PrecipitationCodeDrizzle:
		return "빗방울"
	case PrecipitationCodeSleet:
		return "진눈깨비"
	case PrecipitationCodeSnowGrain:
		return "눈날림"
	default:
		return string(p)
	}
}

// 하늘상태 코드.
type SkyCode string

const (
	// 맑음
	SkyCodeClear SkyCode = "1"
	// 구름 조금
	SkyCodeFewClouds SkyCode = "2"
	// 구름 많음
	SkyCodePartlyCloudy SkyCode = "3"
	// 흐림
	SkyCodeCloudy SkyCode = "4"
)

func (s SkyCode) String() string {
	switch s {
	case SkyCodeClear:
		return "맑음"
	case SkyCodeFewClouds:
		return "구름 조금"
	case SkyCodePartlyCloudy:
		return "구름 많음"
	case SkyCodeCloudy:
		return "흐림"
	default:
		return string(s)
	}
}
