package types

import "time"

type Complaint struct {
	ID           int64     `json:"id"`
	EmployeeID  int       `json:"employeeId"`
	Type         string    `json:"type"`
	Content      string    `json:"content"`
	Date         time.Time `json:"date"`
	EmployeeName string    `json:"employee_name,omitempty"`
	Department   string    `json:"department,omitempty"`
	Position     string    `json:"position,omitempty"`
}

type CreateComplaintPayload struct {
	EmployeeID  int    `json:"employeeId"`
	Type        string `json:"type" validate:"required,oneof=anonymous 'not anonymous'"`
	Content     string `json:"content" validate:"required"`
}

type UpdateComplaintPayload struct {
	ID          int    `json:"id" validate:"required"`
	Type        string `json:"type" validate:"omitempty,oneof=anonymous 'not anonymous'"`
	Content     string `json:"content" validate:"omitempty"`
}

type ComplaintStore interface {
	GetAllComplaints() ([]Complaint, error)
	GetComplaintByID(id int) (*Complaint, error)
	GetComplaintsByEmployeeID(employeeID int) ([]Complaint, error)
	CreateComplaint(payload CreateComplaintPayload) error
	UpdateComplaint(payload UpdateComplaintPayload) error
	DeleteComplaint(id int) error
}