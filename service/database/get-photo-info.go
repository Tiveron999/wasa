package database

func (db *appdbimpl) GetPhotoInfo(f Photo) (Photo, error) {
	res, err := db.c.Exec(`SELECT * FROM photos WHERE pid=?`,
		f.PID)
	if err != nil {
		return f, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return f, err
	} else if affected == 0 {
		// If we didn't delete any row, then the photo didn't exist
		return f, ErrPhotoDoesNotExist
	}
	return f, nil
}
