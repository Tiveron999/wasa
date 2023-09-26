package database

import (
	"database/sql"
)

func (db *appdbimpl) GetUsername(userid uint64) (string, error) {
	var username string
	err := db.c.QueryRow("SELECT username FROM users WHERE usid = ?", userid).Scan(&username)
	if err == sql.ErrNoRows {
		return "", ErrUserDoesNotExist
	} else if err != nil {
		return "", err
	}
	return username, nil
}
