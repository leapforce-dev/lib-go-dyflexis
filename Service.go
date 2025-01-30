package dyflexis

import (
	"encoding/json"
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
)

const (
	apiName       string = "Dyflexis"
	apiDomain     string = "https://app.planning.nu"
	apiBasePath   string = "api/v0"
	apiBasePathV2 string = "api2"
)

// type
type Service struct {
	clientName  string
	authToken   string
	authTokenV2 string
	httpService *go_http.Service
}

type ServiceConfig struct {
	ClientName  string
	AuthToken   string
	AuthTokenV2 string
}

func NewService(serviceConfig *ServiceConfig) (*Service, *errortools.Error) {
	if serviceConfig == nil {
		return nil, errortools.ErrorMessage("ServiceConfig must not be a nil pointer")
	}

	if serviceConfig.ClientName == "" {
		return nil, errortools.ErrorMessage("Service ClientName not provided")
	}

	httpService, e := go_http.NewService(&go_http.ServiceConfig{})
	if e != nil {
		return nil, e
	}

	return &Service{
		clientName:  serviceConfig.ClientName,
		authToken:   serviceConfig.AuthToken,
		authTokenV2: serviceConfig.AuthTokenV2,
		httpService: httpService,
	}, nil
}

func (service *Service) httpRequest(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	if service.authToken == "" {
		return nil, nil, errortools.ErrorMessage("authToken not provided")
	}

	// add authentication header
	header := http.Header{}
	header.Set("X-Dyflexis-AuthToken", service.authToken)
	(*requestConfig).NonDefaultHeaders = &header

	responseModel := requestConfig.ResponseModel
	_response := Response{}
	requestConfig.ResponseModel = &_response

	request, response, e := service.httpService.HttpRequest(requestConfig)
	if e != nil {
		if _response.Error != "" {
			e.SetMessage(_response.Error)
		}
	} else {
		err := json.Unmarshal(_response.Response.Data, responseModel)
		if err != nil {
			return request, response, errortools.ErrorMessage(err)
		}
	}

	return request, response, e
}

func (service *Service) httpRequestV2(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	if service.authTokenV2 == "" {
		return nil, nil, errortools.ErrorMessage("authTokenV2 not provided")
	}

	// add authentication header
	header := http.Header{}
	header.Set("Authorization", fmt.Sprintf("Token %s", service.authTokenV2))
	(*requestConfig).NonDefaultHeaders = &header

	request, response, e := service.httpService.HttpRequest(requestConfig)
	if e != nil {
		return nil, nil, e
	}

	return request, response, e
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s/%s/%s", apiDomain, service.clientName, apiBasePath, path)
}

func (service *Service) urlV2(path string) string {
	return fmt.Sprintf("%s/%s/%s/%s", apiDomain, service.clientName, apiBasePathV2, path)
}

func (service *Service) ApiName() string {
	return apiName
}

func (service *Service) ApiKey() string {
	return service.authToken
}

func (service *Service) ApiCallCount() int64 {
	return service.httpService.RequestCount()
}

func (service *Service) ApiReset() {
	service.httpService.ResetRequestCount()
}
