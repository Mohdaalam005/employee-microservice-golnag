package service

import (
	"context"
	"github.com/friendsofgo/errors"
	"github.com/mohdaalam005/internal/database"
	"github.com/mohdaalam005/pkg/models"
	"log"
)

type EmployeeService interface {
	CreateEmployee(ctx context.Context, employee models.Employee) (*models.Employee, error)
	GetEmployees(ctx context.Context) (models.GetEmployeeResponse, error)
	GetEmployee(ctx context.Context, id int) (models.Employee, error)
	UpdateEmployee(ctx context.Context, id int, employee models.Employee) (*models.Employee, error)
	DeleteEmployee(ctx context.Context, id int) error
}

type employeeServiceImp struct {
	dao database.EmployeeDao
}

func (e employeeServiceImp) DeleteEmployee(ctx context.Context, id int) error {
	err := e.dao.DeleteEmployee(ctx, id)
	if err != nil {
		errors.New("failed to delete")
	}
	return nil
}

func (e employeeServiceImp) UpdateEmployee(ctx context.Context, id int, employee models.Employee) (*models.Employee, error) {
	emp, err := e.dao.UpdateEmployee(ctx, id, employee.MakeDbModel())
	if err != nil {
		errors.New("Service()..... unable to update")
	}
	return &models.Employee{
		ID:     emp.ID,
		Name:   emp.Name,
		Dob:    emp.Dob,
		Gender: emp.Gender,
	}, nil

}

func (e employeeServiceImp) GetEmployee(ctx context.Context, id int) (models.Employee, error) {
	employee, err := e.dao.GetEmployee(ctx, id)
	log.Println("GetEmloyee(id)", employee, id)
	if err != nil {
		return models.Employee{}, errors.New("id is not present")
	}
	return models.Employee{
		ID:     employee.ID,
		Name:   employee.Name,
		Dob:    employee.Dob,
		Gender: employee.Gender,
	}, err

}

func (e employeeServiceImp) CreateEmployee(ctx context.Context, employee models.Employee) (*models.Employee, error) {
	log.Println("service() created employee")
	emp, err := e.dao.CreateEmployee(ctx, employee.MakeDbModel())
	if err != nil {
		return &models.Employee{}, err
	}
	log.Println(emp, "service()")
	return &models.Employee{
		ID:     emp.ID,
		Name:   emp.Name,
		Dob:    emp.Dob,
		Gender: emp.Gender,
	}, nil

}

func (e employeeServiceImp) GetEmployees(ctx context.Context) (models.GetEmployeeResponse, error) {
	//TODO implement me
	student, err := e.dao.GetEmployees(ctx)
	if err != nil {
		return models.GetEmployeeResponse{}, err
	}
	return models.GetEmployeeResponse{
		Employees: student,
	}, nil

}

func NewEmployeeService(dao database.EmployeeDao) EmployeeService {
	return &employeeServiceImp{
		dao: dao,
	}
}
