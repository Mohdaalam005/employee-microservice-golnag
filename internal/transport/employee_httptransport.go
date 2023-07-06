package transport

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/mohdaalam005/internal/dbmodels"
	"github.com/mohdaalam005/internal/endpoints"
	"log"
	"net/http"
	"strconv"
)

func CreateEmployeeHttpHandler(ctx context.Context, endpoint endpoints.EmployeeEndpoints) http.Handler {
	route := mux.NewRouter()
	route.Use(middle)

	route.Methods("GET").Path("/employees").Handler(httptransport.NewServer(
		endpoint.GetEmployees,
		decodeGetEmployees,
		encodeResponse,
	))
	route.Methods("POST").Path("/employees").Handler(httptransport.NewServer(
		endpoint.CreateEmployee,
		decodeCreateEmployee,
		encodeResponse,
	))

	route.Methods("GET").Path("/employees/{employeeId}").Handler(httptransport.NewServer(
		endpoint.GetEmployee,
		decodeGetEmployee,
		encodeResponse,
	))
	route.Methods("PUT").Path("/employees/{employeeId}").Handler(httptransport.NewServer(
		endpoint.UpdateEmployee,
		decodeUpdateEmployee,
		encodeResponse,
	))
	route.Methods("DELETE").Path("/employees/{employeeId}").Handler(httptransport.NewServer(
		endpoint.DeleteEmployee,
		decodeDeleteEmployee,
		encodeResponse,
	))

	return route
}

func decodeDeleteEmployee(ctx context.Context, request2 *http.Request) (request interface{}, err error) {
	vars := mux.Vars(request2)
	id, err := strconv.Atoi(vars["employeeId"])
	return endpoints.EmployeeId{
		ID: id,
	}, err
}

func decodeUpdateEmployee(ctx context.Context, request2 *http.Request) (request interface{}, err error) {
	var employee endpoints.EmployeeRequest
	vars := mux.Vars(request2)

	id, err := strconv.Atoi(vars["employeeId"])

	json.NewDecoder(request2.Body).Decode(&employee)

	if err != nil {
		return nil, err
	}
	return endpoints.EmployeeRequest{
		ID:     id,
		Name:   employee.Name,
		Dob:    employee.Dob,
		Gender: employee.Gender,
	}, nil
}

func decodeGetEmployee(ctx context.Context, request2 *http.Request) (request interface{}, err error) {
	vars := mux.Vars(request2)

	id, err := strconv.Atoi(vars["employeeId"])

	if err != nil {
		return nil, err
	}
	log.Println(id, "decode()...")
	req := endpoints.EmployeeId{
		ID: id,
	}
	return req, nil

}

func decodeCreateEmployee(ctx context.Context, request2 *http.Request) (request interface{}, err error) {
	var employee endpoints.EmployeeRequest
	json.NewDecoder(request2.Body).Decode(&employee)
	log.Println("decode()", employee)
	return employee, err
}

func encodeResponse(ctx context.Context, writer http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(writer).Encode(response)
}

func decodeGetEmployees(ctx context.Context, r *http.Request) (request1 interface{}, err error) {
	var employee dbmodels.Employee
	return employee, nil
}

func middle(handle http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/josn")
		handle.ServeHTTP(writer, request)
	})
}
