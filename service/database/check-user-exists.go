package database

import (
	"database/sql"
)

func (db *appdbimpl) CheckUserExists(u string) (int, error) {
	// Check if the user exists in the database
	var userID int
	err := db.c.QueryRow("SELECT usid FROM users WHERE username = ? ", u).Scan(&userID)
	if err == sql.ErrNoRows {
		// User does not exist, register the user and return the new user ID
		userID, err := db.GetUserID(u)
		if err != nil {
			return 0, err
		}
		return userID, nil
	} else if err != nil {
		return 0, err
	}

	// User exists, return the user ID
	return userID, nil
}
