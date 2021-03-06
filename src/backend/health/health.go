package health

import (
	"database/sql"
	"gitlab.com/flattrack/flattrack.io/src/backend/database"
)

// Healthy ...
// returns if the instance is healthy
func Healthy(db *sql.DB) (err error) {
	err = database.Ping(db)
	return err
}
