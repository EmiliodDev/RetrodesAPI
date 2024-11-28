package complaints

import (
	"database/sql"
	"fmt"

	"github.com/EmiliodDev/gofeed/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateComplaint(payload types.CreateComplaintPayload) error {
	_, err := s.db.Exec(
		"INSERT INTO Complaints (employee_id, type, content) VALUES (?, ?, ?)",
		payload.EmployeeID, payload.Type, payload.Content,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetAllComplaints() ([]types.Complaint, error) {
	rows, err := s.db.Query("SELECT * FROM Complaints ORDER BY id desc");
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	complaints := []types.Complaint{}
	for rows.Next() {
		c, err := scanRowsIntoComplaint(rows)
		if err != nil {
			return nil, err
		}
		complaints = append(complaints, *c)
	} 

	return complaints, nil
}

func (s *Store) GetComplaintByID(id int) (*types.Complaint, error) {
	rows, err := s.db.Query("SELECT * FROM Complaints WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	c := new(types.Complaint)
	if rows.Next() {
		c, err = scanRowsIntoComplaint(rows)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("complaint not found")
	}

	return c, nil
}

func (s *Store) GetComplaintsByEmployeeID(employeeID int) ([]types.Complaint, error) {
	rows, err := s.db.Query("SELECT * FROM Complaints WHERE employee_id = ?", employeeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	complaints := []types.Complaint{}
	for rows.Next() {
		c, err := scanRowsIntoComplaint(rows)
		if err != nil {
			return nil, err
		}
		complaints = append(complaints, *c)
	}

	return complaints, nil
}

func (s *Store) UpdateComplaint(payload types.UpdateComplaintPayload) error {
	_, err := s.db.Exec(
		"UPDATE Complaints SET type = ?, content = ? WHERE id = ?",
		payload.Type, payload.Content, payload.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) DeleteComplaint(id int) error {
	_, err := s.db.Exec("DELETE FROM Complaints WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func scanRowsIntoComplaint(rows *sql.Rows) (*types.Complaint, error) {
	c := new(types.Complaint)

	err := rows.Scan(
		&c.ID,
		&c.EmployeeID,
		&c.Type,
		&c.Content,
		&c.Date,
	)
	if err != nil {
		return nil, err
	}
	return c, nil
}