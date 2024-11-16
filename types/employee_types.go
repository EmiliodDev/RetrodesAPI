package types

type EmployeeStore interface {
	GetEmployeeByEmail(email string) (*Employee, error)
	GetEmployeeByID(id int) (*Employee, error)
	CreateEmployee(Employee) error
}

type Employee struct {
	ID			int		`json:"id"`
	Name		string	`json:"name"`
	LastName	string	`json:"lastName"`
	Email		string	`json:"email"`
	Department	string	`json:"department"`
	Position	string	`json:"position"`
	Password	string	`json:"password"`
}

type RegisterEmployeePayload struct {
	Name		string	`json:"name" validate:"required"`
	LastName	string	`json:"lastName" validate:"required"`
	Email		string	`json:"email" validate:"required,email"`
	Department	string	`json:"department" validate:"required"`
	Position	string	`json:"position" validate:"required"`
	Password	string	`json:"password" validate:"required,min=3,max=30"`
}

type LoginEmployeePayload struct {
	Email	string	`json:"email" validate:"required,email"`
	Password	string	`json:"password" validate:"required"`
}
