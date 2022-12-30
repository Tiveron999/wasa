package database

func (db *appdbimpl) GetSearchedUserInfo(f Profile) (Profile, error) {
	res, err := db.c.Exec(`SELECT * FROM profiles WHERE prid=?`, f.PRID)
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
