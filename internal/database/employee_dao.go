package database

import (
	"context"
	"database/sql"
	"github.com/friendsofgo/errors"
	"github.com/mohdaalam005/internal/dbmodels"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"log"
)

type EmployeeDao interface {
	CreateEmployee(ctx context.Context, employee dbmodels.Employee) (*dbmodels.Employee, error)
	UpdateEmployee(ctx context.Context, id int, employee dbmodels.Employee) (*dbmodels.Employee, error)
	GetEmployee(ctx context.Context, id int) (dbmodels.Employee, error)
	GetEmployees(ctx context.Context) (dbmodels.EmployeeSlice, error)
	DeleteEmployee(ctx context.Context, id int) error
}

type employeeImp struct {
	DB sql.DB
}

func (e employeeImp) CreateEmployee(ctx context.Context, employee dbmodels.Employee) (*dbmodels.Employee, error) {
	log.Println("dao() created employee")
	err := employee.Insert(ctx, &e.DB, boil.Infer())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("dao()", employee)
	return &employee, nil
}

func (e employeeImp) UpdateEmployee(ctx context.Context, id int, employee dbmodels.Employee) (*dbmodels.Employee, error) {
	empById, err := dbmodels.FindEmployee(ctx, &e.DB, id)
	if err != nil {
		errors.New("employee is not present")
	}
	empById.ID = id
	empById.Dob = employee.Dob
	empById.Name = employee.Name
	empById.Gender = employee.Gender
	_, err = empById.Update(ctx, &e.DB, boil.Infer())

	if err != nil {
		errors.New("failed to updated")
	}
	return &dbmodels.Employee{
		ID:     empById.ID,
		Name:   empById.Name,
		Dob:    empById.Dob,
		Gender: empById.Gender,
	}, nil

}

func (e employeeImp) GetEmployee(ctx context.Context, id int) (dbmodels.Employee, error) {
	log.Println("GetEmployee() dao.............")
	employee, err := dbmodels.FindEmployee(ctx, &e.DB, id)

	log.Println(employee, " record", id)
	if err != nil {
		errors.New("id is not present")

	}
	if employee == nil {
		return dbmodels.Employee{}, errors.New("id is not present")
	}

	log.Println(employee, "Getting employee")
	return dbmodels.Employee{
		ID:     employee.ID,
		Name:   employee.Name,
		Dob:    employee.Dob,
		Gender: employee.Gender,
	}, nil
}

func (e employeeImp) GetEmployees(ctx context.Context) (dbmodels.EmployeeSlice, error) {
	employees, err := dbmodels.Employees().All(ctx, &e.DB)
	if err != nil {
		return nil, err
	}
	return employees, nil

}

func (e employeeImp) DeleteEmployee(ctx context.Context, id int) error {
	emp, err := dbmodels.FindEmployee(ctx, &e.DB, id)
	if err != nil {
		return err
	}
	emp.Delete(ctx, &e.DB)
	return nil
}

func NewEmployeeDao(db sql.DB) EmployeeDao {
	return &employeeImp{
		DB: db,
	}
}

type (
	// PortalSlice is an alias for a slice of pointers to Portal.
	// This should almost always be used instead of []Portal.
	PortalSlice []*dbmodels.Employee
)
