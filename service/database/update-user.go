package database

func (db *appdbimpl) UpdateUser(f User) error {
	res, err := db.c.Exec(`UPDATE users SET username=? WHERE id=?`,
		f.Username, f.ID)
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
