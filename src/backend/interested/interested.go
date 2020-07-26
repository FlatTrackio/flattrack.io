/*
  interested
    manage interested data
*/

package interested

import (
	"database/sql"
	"errors"

	"gitlab.com/flattrack/flattrack.io/src/backend/common"
	"gitlab.com/flattrack/flattrack.io/src/backend/types"
)

// CheckIfEmailInDB ...
// returns bool if an email is in the database
func CheckIfEmailInDB(db *sql.DB, email string) (found bool, err error) {
	sqlStatement := `select email from interested where email = $1`
	rows, err := db.Query(sqlStatement, email)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	rows.Next()
	var emailFromDB string
	rows.Scan(&emailFromDB)
	err = rows.Err()
	return emailFromDB == email, err
}

// AddEmailToInterested ...
// inserts an email into the database
func AddEmailToInterested(db *sql.DB, interested types.InterestedSpec) (message string, err error) {
	found, err := CheckIfEmailInDB(db, interested.Email)
	if found == true || err != nil {
		var errString string
		if err != nil {
			errString = err.Error()
		}
		errorMsg := "Did not insert row. Email either exists or an error occured. " + errString
		return "Your email has been added to the notify list", errors.New(errorMsg)
	}
	valid := common.RegexMatchEmail(interested.Email)
	if valid != true {
		return "Unable to use that email, it doesn't appear that it's valid", errors.New("Email validation failed")
	}

	sqlStatement := `insert into interested (email)
                         values ($1)
                         returning *`
	rows, err := db.Query(sqlStatement, interested.Email)
	if err != nil {
		return "Failed to add email to notify list", errors.New("Failed to insert the email into the database")
	}
	rows.Next()
	if err != nil {
		return "Failed to add email to notify list", err
	}

	return "Your email has been added to the notify list", nil
}

// ResetAllEntries ...
// empties all emails sent
func ResetAllEntries (db *sql.DB) (err error) {
	sqlStatement := `delete from interested`
	_, err = db.Query(sqlStatement)
	return err
}
