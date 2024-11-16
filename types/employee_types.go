package types

type EmployeeStore interface {

}

type Employee struct {
	ID			int		`json:"id"`
	Name		string	`json:"name"`
	LastName	string	`json:"lastName"`
	Email		string	`json:"email"`
	Department	string	`json:"department"`
	Position	string	`json:"position"`
}

type RegisterEmployeePayload struct {
	Name		string	`json:"name" validate:"required"`
	LastName	string	`json:"lastName" validate:"required"`
	Email		string	`json:"email" validate:"required,email"`
	Department	string	`json:"department" validate:"required"`
	Position	string	`json:"position" validate:"required"`
	Username	string	`json:"username" validate:"required"`
	Password	string	`json:"password" validate:"required,min=3,max=30"`
}

type LoginEmployeePayload struct {
	Username	string	`json:"username" validate:"required"`
	Password	string	`json:"password" validate:"required"`
}
