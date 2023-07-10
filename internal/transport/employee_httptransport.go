package transport

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/mohdaalam005/internal/dbmodels"
	"github.com/mohdaalam005/internal/endpoints"
	"io"
	"log"
	"net/http"
	"strconv"
)

func CreateEmployeeHttpHandler(ctx context.Context, endpoint endpoints.EmployeeEndpoints, route *mux.Router) http.Handler {

	route.Methods("GET").Path("/employees").Handler(httptransport.NewServer(
		endpoint.GetEmployees,
		decodeGetEmployees,
		encodeResponse,
	))
	// swagger:operation POST /employees employee createEmployee
	// ---
	// summary: Returns the employee ID
	// parameters:
	// - name: employees
	//   in: body
	//   description: create the employee
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/EmployeeRequest"
	// responses:
	//   "200":
	//     "$ref": "#/responses/EmployeeResponse"
	route.Methods("POST").Path("/employees").Handler(httptransport.NewServer(
		endpoint.CreateEmployee,
		decodeCreateEmployee,
		encodeResponse,
	))
	// swagger:operation GET /employees/{employeeId} employee getEmployee
	// ---
	// summary: Returns The Employee By ID
	// parameters:
	// - name: employeeId
	//   in: path
	//   description: Get employee by ID
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/getEmployeeResponse"
	route.Methods("GET").Path("/employees/{employeeId}").Handler(httptransport.NewServer(
		endpoint.GetEmployee,
		decodeGetEmployee,
		encodeResponse,
	))
	// swagger:operation PUT /employees/{employeeId} employee updateEmployee
	// ---
	// summary: Returns The Employee By ID
	// parameters:
	// - name: employeeId
	//   in: path
	//   description: Get employee by ID
	//   type: string
	//   required: true
	// - name: employee
	//   in: body
	//   description: create the employee
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/EmployeeRequest"
	// responses:
	//   "200":
	//     "$ref": "#/responses/getUpdateEmployeeResponse"
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
	body, err := io.ReadAll(request2.Body)

	//json.NewDecoder(request2.Body).Decode(&employee)
	err = json.Unmarshal(body, &employee)
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
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(writer).Encode(response)
}

func decodeGetEmployees(ctx context.Context, r *http.Request) (request1 interface{}, err error) {
	var employee dbmodels.Employee
	return employee, nil
}
