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
	GetEmployees(ctx context.Context) ([]models.Employee, error)
	GetEmployee(ctx context.Context, id int) (models.Employee, error)
	UpdateEmployee(ctx context.Context, id int, employee models.Employee) (*models.Employee, error)
	DeleteEmployee(ctx context.Context, id int) error
}

type employeeServiceImp struct {
	dao database.EmployeeDao
}

func (e employeeServiceImp) DeleteEmployee(ctx context.Context, id int) error {
	_, err := e.dao.GetEmployee(ctx, id)
	if err != nil {
		errors.New("id is not present")
		return nil
	}
	e.dao.DeleteEmployee(ctx, id)
	if err != nil {
		errors.New("failed to delete")
	}
	return nil
}

func (e employeeServiceImp) UpdateEmployee(ctx context.Context, id int, employeeUpdate models.Employee) (*models.Employee, error) {
	existingEmp, err := e.GetEmployee(ctx, id)
	if err != nil {
		errors.New("Service()..... unable to update")
	}
	employee := existingEmp
	employee.ID = employeeUpdate.ID
	employee.Name = employeeUpdate.Name
	employee.Dob = employeeUpdate.Dob
	employee.Gender = employeeUpdate.Gender

	empModel := employee.MakeDbModel()
	returnEmp, err := e.dao.UpdateEmployee(ctx, empModel)
	newEmp := models.MakeModelToDb(*returnEmp)
	return &newEmp, nil

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
func (e employeeServiceImp) GetEmployees(ctx context.Context) ([]models.Employee, error) {
	var employees []models.Employee
	emp, err := e.dao.GetEmployees(ctx)
	if err != nil {
		return nil, err
	}
	for _, employee := range emp {
		employee := models.Employee{
			ID:     employee.ID,
			Name:   employee.Name,
			Dob:    employee.Dob,
			Gender: employee.Gender,
		}
		employees = append(employees, employee)

	}

	return employees, nil

}

func NewEmployeeService(dao database.EmployeeDao) EmployeeService {
	return &employeeServiceImp{
		dao: dao,
	}
}
