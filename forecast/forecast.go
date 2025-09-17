// 단기 예보 조회
//
// https://www.data.go.kr/data/15084084/openapi.do
package forecast

import (
	"fmt"
	"io"
	"net/url"
	"strconv"

	"github.com/jh1104/publicapi"
)

var _ publicapi.API = (*Forecast)(nil)

// 단기 예보 조회 API.
type Forecast struct {
	Subtype Subtype
	Params  Parameters
}

func NewForecast(subtype Subtype, params Parameters) *Forecast {
	return &Forecast{
		Subtype: subtype,
		Params:  params,
	}
}

func (f *Forecast) Method() string {
	return "GET"
}

func (f *Forecast) URL(serviceKey string) string {
	values := url.Values{}
	values.Add("serviceKey", serviceKey)
	values.Add("pageNo", strconv.Itoa(f.Params.PageNo))
	values.Add("numOfRows", strconv.Itoa(f.Params.NumberOfRows))
	values.Add("base_date", f.Params.BaseDate)
	values.Add("base_time", f.Params.BaseTime)
	values.Add("nx", strconv.Itoa(f.Params.NX))
	values.Add("ny", strconv.Itoa(f.Params.NY))
	values.Add("dataType", "JSON")

	return fmt.Sprintf("%s/%s?%s", publicapi.BaseURL, f.Subtype.Path(), values.Encode())
}

func (f *Forecast) Body() io.Reader {
	return nil
}
