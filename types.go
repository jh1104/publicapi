package publicapi

import (
	"crypto/tls"
	"errors"
	"net/http"
	"time"
)

// 공공 데이터 포털 API 기본 URL.
const BaseURL = "https://apis.data.go.kr"

// 공공 데이터 포털 API 호출시 사용할 기본 HTTP 클라이언트.
// 특정 TLS 버전과 암호화 스위트를 사용해야 문제없이 호출할 수 있다.
var DefaultHTTPClient = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
			MaxVersion: tls.VersionTLS13,
			CipherSuites: []uint16{
				tls.TLS_RSA_WITH_AES_128_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			},
		},
	},
	Timeout: 30 * time.Second,
}

var (
	// 어플리케이션 에러.
	ErrApplicationError = errors.New("application error")
	// 데이터베이스 에러.
	ErrDBError = errors.New("database error")
	// 데이터 없음 에러.
	ErrNoData = errors.New("no data")
	// HTTP 에러.
	ErrHTTPError = errors.New("http error")
	// 서비스 연결실패 에러.
	ErrServiceTimeout = errors.New("service timeout")
	// 잘못된 파라미터 요청.
	ErrInvalidParams = errors.New("invalid request parameter")
	// 필수 요청 파라미터가 없음.
	ErrNoMandatoryParams = errors.New("no mandatory request parameters")
	// 해당 API 서비스가 없거나 폐기.
	ErrNoAPIService = errors.New("no api service")
	// 서비스 접근거부.
	ErrAccessDenied = errors.New("service access denied")
	// 일시적으로 사용할 수 없는 서비스 키.
	ErrDisabledKey = errors.New("temporarily disabled service key")
	// 서비스 요청제한횟수 초과.
	ErrLimitExceeded = errors.New("request limit exceeded")
	// 등록되지 않은 서비스키.
	ErrInvalidKey = errors.New("invalid service key")
	// 기한이 만료된 서비스키.
	ErrExpiredKey = errors.New("expired service key")
	// 등록되지 않은 IP주소.
	ErrUnregisteredIP = errors.New("unregistered ip")
	// 서명되지 않은 호출.
	ErrUnsignedCall = errors.New("unsigned call")
	// 기타 에러.
	ErrUnknown = errors.New("unknown error")
)

// API 결과 코드.
type ResultCode string

const (
	CodeNormal            ResultCode = "00"
	CodeApplicationError  ResultCode = "01"
	CodeDBError           ResultCode = "02"
	CodeNoData            ResultCode = "03"
	CodeHTTPError         ResultCode = "04"
	CodeServiceTimeout    ResultCode = "05"
	CodeInvalidParams     ResultCode = "10"
	CodeNoMandatoryParams ResultCode = "11"
	CodeNoAPIService      ResultCode = "12"
	CodeAccessDenied      ResultCode = "20"
	CodeDisabledKey       ResultCode = "21"
	CodeLimitExceeded     ResultCode = "22"
	CodeInvalidKey        ResultCode = "30"
	CodeExpiredKey        ResultCode = "31"
	CodeUnregisteredIP    ResultCode = "32"
	CodeUnsignedCall      ResultCode = "33"
	CodeUnknownError      ResultCode = "99"
)

// 결과 코드를 error 타입으로 변환한다.
func (r ResultCode) AsError() error {
	switch r {
	case CodeNormal:
		return nil
	case CodeApplicationError:
		return ErrApplicationError
	case CodeDBError:
		return ErrDBError
	case CodeNoData:
		return ErrNoData
	case CodeHTTPError:
		return ErrHTTPError
	case CodeServiceTimeout:
		return ErrServiceTimeout
	case CodeInvalidParams:
		return ErrInvalidParams
	case CodeNoMandatoryParams:
		return ErrNoMandatoryParams
	case CodeNoAPIService:
		return ErrNoAPIService
	case CodeAccessDenied:
		return ErrAccessDenied
	case CodeDisabledKey:
		return ErrDisabledKey
	case CodeLimitExceeded:
		return ErrLimitExceeded
	case CodeInvalidKey:
		return ErrInvalidKey
	case CodeExpiredKey:
		return ErrExpiredKey
	case CodeUnregisteredIP:
		return ErrUnregisteredIP
	case CodeUnsignedCall:
		return ErrUnsignedCall
	case CodeUnknownError:
		return ErrUnknown
	default:
		if string(r) == "" {
			return errors.New("empty code")
		}
		return errors.New("unknown result code: " + string(r))
	}
}
