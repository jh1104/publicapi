# PublicAPI

공공데이터 포털(data.go.kr) API를 위한 Go 라이브러리.

## 기능

* [특일 정보 조회 API](https://www.data.go.kr/data/15012690/openapi.do)
* [단기 예보 조회 API](https://www.data.go.kr/data/15084084/openapi.do)

## 설치

```bash
go get github.com/jh1104/publicapi
```

## 사용법

```go
package main

import (
    "fmt"

    "github.com/jh1104/publicapi/specialday"
)

func main() {
    // 클라이언트 생성
    client := specialday.NewClient("YOUR_API_KEY")

	// 2025년 5월 조회
	params := specialday.Parameters{
		Year:         2025,
		Month:        05,
		NumberOfRows: 10,
		PageNo:       1,
	}

    // 공휴일 조회 API 호출
	resp, err := client.ListNationalHolidays(context.Background(), params)
	if err != nil {
		panic(err)
	}

	// 결과 출력
	for _, item := range resp.Body.Data.Items {
		fmt.Printf("날짜: %d, 이름: %s\n", item.Date, item.Name)
	}
}
```
