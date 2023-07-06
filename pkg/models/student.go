package models

import (
	"github.com/mohdaalam005/internal/dbmodels"
	"github.com/volatiletech/null/v8"
)

type Employee struct {
	ID     int         `json:"id"`
	Name   null.String `json:"name"`
	Dob    null.String `json:"dob"`
	Gender null.String `json:"gender"`
}

func (e *Employee) MakeDbModel() dbmodels.Employee {
	model := dbmodels.Employee{
		ID:     e.ID,
		Name:   e.Name,
		Dob:    e.Dob,
		Gender: e.Gender,
	}
	return model
}

type GetEmployeeResponse struct {
	Employees dbmodels.EmployeeSlice
}
