package credentials

import (
	"database/sql"

	"github.com/EmiliodDev/gofeed/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db:db}
}

func (s *Store) CreateCredentials(cred *types.Credentials) error {
	_, err := s.db.Exec("INSERT INTO Credentials (employee_id, username, password) VALUES (?,?,?)", cred.EmployeeID, cred.Username, cred.Password)
	if err != nil {
		return err
	}

	return nil
}