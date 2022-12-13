package database

func (db *appdbimpl) GetUserInfo(f User) error {
	res, err := db.c.Exec(`SELECT * FROM users WHERE id=?`,
		f.ID)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the user didn't exist
		return ErrUserDoesNotExist
	}
	return nil
}
