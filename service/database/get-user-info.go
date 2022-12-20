package database

func (db *appdbimpl) GetUserInfo(f User) (User, error) {
	res, err := db.c.Exec(`SELECT * FROM users WHERE id=?`, f.ID)
	if err != nil {
		return f, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return f, err
	} else if affected == 0 {
		// If we didn't delete any row, then the user didn't exist
		return f, ErrUserDoesNotExist
	}
	return f, nil
}
