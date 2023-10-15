package database

func (db *appdbimpl) CreateUser(username string) (uint64, error) {
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
