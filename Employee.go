package dyflexis

import (
	"fmt"
	"net/http"

	d_types "github.com/leapforce-libraries/go_dyflexis/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

type Employee struct {
	Id                 go_types.Int64String   `json:"id"`
	Firstname          string                 `json:"firstname"`
	Infix              string                 `json:"infix"`
	Surname            string                 `json:"surname"`
	ContractId         go_types.Int64String   `json:"contract_id"`
	ContractStart      d_types.DateString     `json:"contract_start"`
	ContractEnd        d_types.DateString     `json:"contract_end"`
	ContractHoursWeek  go_types.Float64String `json:"contract_hours_week"`
	ContractSalaryHour go_types.Float64String `json:"contract_salary_hour"`
	ContractTypeId     go_types.Int64String   `json:"contract_type_id"`
	ContractTypeName   string                 `json:"contract_type_name"`
}

// GetEmployees returns all employees
//
func (service *Service) GetEmployees() (*[]Employee, *errortools.Error) {
	page := 0

	employees := []Employee{}

	for {
		_employees := []Employee{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(fmt.Sprintf("employee?page=%v", page)),
			ResponseModel: &_employees,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		if len(_employees) == 0 {
			break
		}

		employees = append(employees, _employees...)

		page++
	}

	return &employees, nil
}
