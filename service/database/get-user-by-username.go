package database

import (
	"database/sql"
)

func (db *appdbimpl) GetUserID(username string) (int, error) {
	var userID int
	err := db.c.QueryRow("SELECT usid FROM users WHERE username = ?", username).Scan(&userID)
	if err == sql.ErrNoRows {
		return 0, ErrUserDoesNotExist
	} else if err != nil {
		return 0, err
	}
	return userID, nil
}
