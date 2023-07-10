package endpoints

import (
	"context"
	"github.com/friendsofgo/errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/mohdaalam005/internal/service"
	"github.com/mohdaalam005/pkg/models"
	"github.com/volatiletech/null/v8"
	"log"
)

type EmployeeEndpoints struct {
	CreateEmployee endpoint.Endpoint
	GetEmployees   endpoint.Endpoint
	GetEmployee    endpoint.Endpoint
	UpdateEmployee endpoint.Endpoint
	DeleteEmployee endpoint.Endpoint
}

func MakeEmployeeEndpoints(s service.EmployeeService) EmployeeEndpoints {
	return EmployeeEndpoints{
		CreateEmployee: makeCreateEmployeeEndpoint(s),
		GetEmployees:   makeGetEmployeesEndpoint(s),
		GetEmployee:    makeGetEmployeeEndpoint(s),
		UpdateEmployee: makeUpdateEmployeeEndpoint(s),
		DeleteEmployee: makeDeleteEmployeeEndpoint(s),
	}
}

func makeDeleteEmployeeEndpoint(s service.EmployeeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(EmployeeId)
		err = s.DeleteEmployee(ctx, req.ID)
		if err != nil {
			errors.New("not able to delete MakeDeleteEndpoint().......")
		}
		return req, nil

	}
}

func makeUpdateEmployeeEndpoint(s service.EmployeeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(EmployeeRequest)
		res, err := s.UpdateEmployee(ctx, req.ID, models.Employee{
			ID:     req.ID,
			Name:   req.Name,
			Gender: req.Gender,
			Dob:    req.Dob,
		})
		return getUpdateEmployeeResponse{
			Employee: res,
		}, nil
	}
}

// swagger:response getUpdateEmployeeResponse
type getUpdateEmployeeResponse struct {
	// in : body
	Employee *models.Employee
}

func makeGetEmployeeEndpoint(s service.EmployeeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(EmployeeId)
		if !ok {
			return nil, errors.New("id is not present for that employee")
		}

		res, err := s.GetEmployee(ctx, req.ID)
		if err != nil {
			return nil, err
		}

		return getEmployeeResponse{
			Employee: res}, nil
	}
}

type EmployeeId struct {
	ID int `json:"id"`
}

// swagger:response getEmployeeResponse
type getEmployeeResponse struct {
	// in : body
	Employee models.Employee
}

func makeGetEmployeesEndpoint(s service.EmployeeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		students, err := s.GetEmployees(ctx)
		log.Println("makeGetEmployeesEndpoint()..>.>.>")
		return students, nil
	}
}

// swagger:response getEmployeesResponse
type getEmployeesResponse struct {
	// in : body
	employees []models.Employee
}

func makeCreateEmployeeEndpoint(s service.EmployeeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(EmployeeRequest)
		if !ok {
			return nil, err
		}
		res, err := s.CreateEmployee(ctx, models.Employee{
			ID:     req.ID,
			Name:   req.Name,
			Gender: req.Gender,
			Dob:    req.Dob,
		})
		return res, nil

	}
}

// swagger:model EmployeeRequest
type EmployeeRequest struct {
	ID     int         `json:"id"`
	Name   null.String `json:"name"`
	Dob    null.String `json:"dob"`
	Gender null.String `json:"gender"`
}

// swagger:response EmployeeResponse
type EmployeeResponse struct {
	// in body
	Id int `json:"id"`
}
