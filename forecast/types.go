package forecast

// 단기 예보 조회 API의 세부 분류.
type Subtype int

const (
	// 초단기 예보 조회
	UltraShortTermForecast Subtype = iota + 1
)

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
