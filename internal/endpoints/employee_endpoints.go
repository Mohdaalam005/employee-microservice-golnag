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
		CreateEmployee: makeEmployeeEndpoint(s),
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
		return res, nil
	}
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

		return res, nil
	}
}

func makeGetEmployeesEndpoint(s service.EmployeeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		students, err := s.GetEmployees(ctx)
		log.Println("makeGetEmployeesEndpoint()..>.>.>")
		return students, nil
	}
}

func makeEmployeeEndpoint(s service.EmployeeService) endpoint.Endpoint {
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

// swagger : model EmployeeRequest
type EmployeeRequest struct {
	ID     int         `json:"id"`
	Name   null.String `json:"name"`
	Dob    null.String `json:"dob"`
	Gender null.String `json:"gender"`
}

type EmployeeResponse struct {
	Id int `json:"id"`
}

type EmployeeId struct {
	ID int `json:"id"`
}
