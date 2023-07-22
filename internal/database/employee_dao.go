package database

import (
	"context"
	"github.com/friendsofgo/errors"
	database2 "github.com/mohdaalam005/go-common/database"
	"github.com/mohdaalam005/internal/dbmodels"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"log"
)

type EmployeeDao interface {
	CreateEmployee(ctx context.Context, employee dbmodels.Employee) (*dbmodels.Employee, error)
	UpdateEmployee(ctx context.Context, employee dbmodels.Employee) (*dbmodels.Employee, error)
	GetEmployee(ctx context.Context, id int) (dbmodels.Employee, error)
	GetEmployees(ctx context.Context) (dbmodels.EmployeeSlice, error)
	DeleteEmployee(ctx context.Context, id int) error
}

type employeeImp struct {
	DB database2.DbConnection
}

func (e employeeImp) CreateEmployee(ctx context.Context, employee dbmodels.Employee) (*dbmodels.Employee, error) {
	log.Println("dao() created employee")
	err := employee.Insert(ctx, e.DB.Conn, boil.Infer())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("dao()", employee)
	return &employee, nil
}

func (e employeeImp) UpdateEmployee(ctx context.Context, employee dbmodels.Employee) (*dbmodels.Employee, error) {
	log.Println("Dao() update employee")
	emp, err := employee.Update(ctx, e.DB.Conn, boil.Infer())
	if err != nil {
		errors.New("Dao() failed to updated")
	}
	log.Println("updated the ", emp)
	return &employee, err

}

func (e employeeImp) GetEmployee(ctx context.Context, id int) (dbmodels.Employee, error) {
	log.Println("GetEmployee() dao.............")
	employee, err := dbmodels.FindEmployee(ctx, e.DB.Conn, id)

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
	employees, err := dbmodels.Employees().All(ctx, e.DB.Conn)
	if err != nil {
		return nil, err
	}
	return employees, nil

}

func (e employeeImp) DeleteEmployee(ctx context.Context, id int) error {
	emp, err := dbmodels.FindEmployee(ctx, e.DB.Conn, id)
	if err != nil {
		return err
	}
	emp.Delete(ctx, e.DB.Conn)
	return nil
}

func NewEmployeeDao(db database2.DbConnection) EmployeeDao {
	return &employeeImp{
		DB: db,
	}
}

type (
	// EmployeeSlice is an alias for a slice of pointers to Employee.
	// This should almost always be used instead of []Employee.
	EmployeeSlice []*dbmodels.Employee
)
