package database

import (
	"errors"
	// "fmt"
)

func (db *appdbimpl) CheckUserExists(u string) (uint64, error) {
	// Check if the user exists in the database
	var userID uint64

	userID, err := db.GetUserID(u)
	if errors.Is(err, ErrUserDoesNotExist) {
		// fmt.Print("hii")
		// User does not exist, register the user and return the new user ID
		userID, err := db.CreateUser(u)
		// fmt.Print(err)
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
