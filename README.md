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
func main() {
	client := publicapi.NewClient("YOUR_SERVICE_KEY")
	specialday.SetDefaultClient(client)

	// 2025년 5월 공휴일 조회.
	resp, err := specialday.ListHolidays(context.Background(), 2025, 5)
	if err != nil {
		panic(err)
	}

	for _, item := range resp.Body.Data.Items {
		fmt.Printf("날짜: %d, 이름: %s\n", item.Date, item.Name)
	}
}
```
