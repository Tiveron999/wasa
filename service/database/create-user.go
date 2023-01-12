package database

import "fmt"

func (db *appdbimpl) CreateUser(username string) (uint64, error) {
	/*res, err := db.c.Exec(`INSERT INTO users (usid, username) VALUES (NULL, ?)`,
			username)
		if err != nil {
			return User{}, err
		}

		lastInsertID, err := res.LastInsertId()
		if err != nil {
			return User{}, err
		}

		user := User{USID: uint64(lastInsertID), Username: username}
		return user, nil
	}*/
	fmt.Print("tomatoes")
	res, err := db.c.Exec(`INSERT INTO users (usid, username) VALUES (NULL, ?)`, username)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}
