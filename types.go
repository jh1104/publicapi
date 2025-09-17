package publicapi

import (
	"crypto/tls"
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
