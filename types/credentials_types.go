package types

type CredentialsStore interface {
	
}

type Credentials struct {
	ID			int		`json:"id"`
	EmployeeID	int		`json:"employeeID"`
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}