package employee

import (
	"database/sql"
	"fmt"

	"github.com/EmiliodDev/gofeed/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db:db}
}

func (s *Store) CreateEmployee(employee types.Employee) error {
	_, err := s.db.Exec("INSERT INTO Employees (name, lastname, email, department, position, password) VALUES (?,?,?,?,?,?)", employee.Name, employee.LastName, employee.Email, employee.Department, employee.Position, employee.Password)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetEmployeeByEmail(email string) (*types.Employee, error) {
	rows, err := s.db.Query("SELECT 1 FROM Employees WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	e := new(types.Employee)
	for rows.Next() {
		e, err = scanRowsIntoEmployee(rows)
		if err != nil {
			return nil, err
		}
	}

	if e.ID == 0 {
		return nil, fmt.Errorf("employee not found")
	}

	return e, nil
}

func (s *Store) GetEmployeeByID(id int) (*types.Employee, error) {
	rows, err := s.db.Query("SELECT 1 FROM Employees WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	e := new(types.Employee)
	for rows.Next() {
		e, err = scanRowsIntoEmployee(rows)
		if err != nil {
			return nil, err
		}
	}
	if e.ID == 0 {
		return nil, fmt.Errorf("employee not found")
	}
	return e, nil
}

func scanRowsIntoEmployee(rows *sql.Rows) (*types.Employee, error) {
	e := new(types.Employee)

	err := rows.Scan(
		&e.ID,
		&e.Name,
		&e.LastName,
		&e.Email,
		&e.Department,
		&e.Position,
	)
	if err != nil {
		return nil, err
	}
	return e, nil
}