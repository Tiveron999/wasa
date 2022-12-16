package database

func (db *appdbimpl) GetPhotoInfo(f Photo) error {
	res, err := db.c.Exec(`SELECT * FROM photos WHERE id=?`,
		f.id)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the photo didn't exist
		return ErrPhotoDoesNotExist
	}
	return nil
}
