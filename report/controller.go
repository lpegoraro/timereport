package report

import (
	"net/http"
)

type controller struct{}

var reportService reportapi.ReportServiceClient

type ReportController interface {
	GetReport(response http.ResponseWriter, request *http.Request)
	AddReport(response http.ResponseWriter, request *http.Request)
}

type NewReportController(service service.ReportService) ReportController {
	reportService = service
	return &controller{}
}

func (*controller) GetReport(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var filter reportApi.ReportFilter
	err := json.NewDecoder(request.Body).Decode(&report)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json..NewEncoder(response).Encode(errors.ServiceError{Message: "Request body format invalid"})
	}
}
