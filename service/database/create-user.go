package database

func (db *appdbimpl) CreateUser(f User) (User, error) {
	res, err := db.c.Exec(`INSERT INTO users (id, username) VALUES (?, ?)`,
		f.ID, f.Username)
	if err != nil {
		return f, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return f, err
	}

	f.ID = uint64(lastInsertID)
	return f, nil
}
