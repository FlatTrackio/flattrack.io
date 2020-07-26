/*
  database
    handle connections to the database
*/

package database

import (
	"database/sql"
	"fmt"

	// include pg
	_ "github.com/lib/pq"
)

// DB ...
// given database credentials, return a database connection
func DB(username string, password string, hostname string, database string, sslmode string) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=%v", username, password, hostname, database, sslmode)
	return sql.Open("postgres", connStr)
}
