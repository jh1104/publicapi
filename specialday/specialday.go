// 특일 정보 조회 API
//
// https://www.data.go.kr/data/15012690/openapi.do
package specialday

import (
	"fmt"
	"io"
	"net/url"

	"github.com/jh1104/publicapi"
)

var _ publicapi.API = (*SpecialDay)(nil)

// 특일 정보 조회 API.
type SpecialDay struct {
	Subtype Subtype
	Params  Parameters
}

func NewSpecialDay(subtype Subtype, params Parameters) *SpecialDay {
	return &SpecialDay{
		Subtype: subtype,
		Params:  params,
	}
}

func (s *SpecialDay) Method() string {
	return "GET"
}

func (s *SpecialDay) URL(serviceKey string) string {
	values := url.Values{}
	values.Add("serviceKey", serviceKey)
	values.Add("solYear", fmt.Sprintf("%d", s.Params.Year))
	values.Add("solMonth", fmt.Sprintf("%02d", s.Params.Month))
	values.Add("numOfRows", fmt.Sprintf("%d", s.Params.NumberOfRows))
	values.Add("pageNo", fmt.Sprintf("%d", s.Params.PageNo))
	values.Add("_type", "json")

	return fmt.Sprintf("%s/%s?%s", publicapi.BaseURL, s.Subtype.Path(), values.Encode())
}

func (s *SpecialDay) Body() io.Reader {
	return nil
}
