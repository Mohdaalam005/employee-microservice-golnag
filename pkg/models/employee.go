package models

import (
	"github.com/mohdaalam005/internal/dbmodels"
	"github.com/volatiletech/null/v8"
)

type Employee struct {
	// in body
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

type GetEmployeesResponse struct {
	Employees dbmodels.EmployeeSlice
}

func MakeModelToDb(employee dbmodels.Employee) Employee {
	return Employee{
		ID:     employee.ID,
		Name:   employee.Name,
		Dob:    employee.Dob,
		Gender: employee.Gender,
	}

}

func (e *Employee) MakeModels(slice dbmodels.EmployeeSlice) []Employee {
	employeeSlice := make([]Employee, len(slice))

	for i, portal := range employeeSlice {
		employee := Employee{
			ID:     portal.ID,
			Name:   portal.Name,
			Dob:    portal.Dob,
			Gender: portal.Gender,
			// Assign other fields accordingly
		}
		employeeSlice[i] = employee
	}
	return employeeSlice
}
