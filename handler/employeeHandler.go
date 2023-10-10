package handler

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"github.com/Tango992/moduletasty/entity"
)

func GetEmployeeInput() (entity.Employee, error) {
	scanner := bufio.NewScanner(os.Stdin)
	employee := entity.Employee{}

	fmt.Print("Enter first name: ")
	scanner.Scan()
	if employee.First_name = scanner.Text(); len(employee.First_name) < 1 {
		return entity.Employee{}, errors.New("empty input")
	}

	fmt.Print("Enter last name: ")
	scanner.Scan()
	if employee.Last_name = scanner.Text(); len(employee.Last_name) < 1 {
		return entity.Employee{}, errors.New("empty input")
	}

	fmt.Print("Enter position: ")
	scanner.Scan()
	if employee.Position = scanner.Text(); len(employee.Position) < 1 {
		return entity.Employee{}, errors.New("empty input")
	}
	fmt.Println("Employee added successfully")
	return employee, nil
}

func AddEmployee(db *sql.DB, employee entity.Employee) error {
	_, err := db.Exec("INSERT INTO Employees (first_name, last_name, position) VALUES (?,?,?)", employee.First_name, employee.Last_name, employee.Position)
	return err
}