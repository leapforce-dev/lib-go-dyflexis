package dyflexis

import (
	"fmt"
	"github.com/leapforce-libraries/go_dyflexis/types"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type EmployeeDataResponse struct {
	Page         string         `json:"page"`
	TotalPages   int            `json:"totalPages"`
	EmployeeData []EmployeeData `json:"employeeData"`
}

type EmployeeData struct {
	DyflexisId            int64             `json:"dyflexisId"`
	PersonnelNumber       string            `json:"personnelNumber"`
	FirstName             string            `json:"firstName"`
	LastNamePrefix        *string           `json:"lastNamePrefix"`
	LastName              string            `json:"lastName"`
	Email                 string            `json:"email"`
	PhoneNumber           string            `json:"phoneNumber"`
	PhoneNumber2          string            `json:"phoneNumber2"`
	PartnerLastNamePrefix *string           `json:"partnerLastNamePrefix"`
	PartnerLastName       *string           `json:"partnerLastName"`
	NameFormat            string            `json:"nameFormat"`
	EmploymentStart       *types.DateString `json:"employmentStart"`
	EmploymentEnd         *types.DateString `json:"employmentEnd"`
	CardNumbers           []string          `json:"cardNumbers"`
}

// GetEmployeeData returns all employee data
func (service *Service) GetEmployeeData() (*[]EmployeeData, *errortools.Error) {
	page := 1

	var employeeData []EmployeeData

	for {
		var employeeDataResponse EmployeeDataResponse

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.urlV2(fmt.Sprintf("employee-data?page=%v", page)),
			ResponseModel: &employeeDataResponse,
		}

		_, _, e := service.httpRequestV2(&requestConfig)
		if e != nil {
			return nil, e
		}

		employeeData = append(employeeData, employeeDataResponse.EmployeeData...)

		if page <= employeeDataResponse.TotalPages {
			break
		}

		page++
	}

	return &employeeData, nil
}
