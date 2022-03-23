package dyflexis

import (
	"fmt"
	"net/http"

	d_types "github.com/leapforce-libraries/go_dyflexis/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

type Worked struct {
	Id             go_types.Int64String    `json:"id"`
	UserId         go_types.Int64String    `json:"user_id"`
	Firstname      string                  `json:"firstname"`
	Infix          string                  `json:"infix"`
	Surname        string                  `json:"surname"`
	OfficeId       *go_types.Int64String   `json:"office_id"`
	OfficeName     *string                 `json:"office_name"`
	DepartmentId   *go_types.Int64String   `json:"department_id"`
	DepartmentName *string                 `json:"department_name"`
	CostplaceId    *go_types.Int64String   `json:"costplace_id"`
	CostplaceName  *string                 `json:"costplace_name"`
	CostplaceCode  *string                 `json:"costplace_code"`
	Kilometers     *go_types.Float64String `json:"kilometers"`
	StartDate      d_types.DateTimeString  `json:"start_date"`
	EndDate        d_types.DateTimeString  `json:"end_date"`
	Pauze          go_types.Float64String  `json:"pauze"`
	Duration       go_types.Float64String  `json:"duration"`
}

// GetWorkeds returns all employees
//
func (service *Service) GetWorkeds(year int, month int) (*[]Worked, *errortools.Error) {
	page := 0

	workeds := []Worked{}

	for {
		_workeds := []Worked{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(fmt.Sprintf("worked?year=%v&month=%v&page=%v", year, month, page)),
			ResponseModel: &_workeds,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		if len(_workeds) == 0 {
			break
		}

		workeds = append(workeds, _workeds...)

		page++
	}

	return &workeds, nil
}
