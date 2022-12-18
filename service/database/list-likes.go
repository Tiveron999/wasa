package database

func (db *appdbimpl) ListLikes(f Photo) error {
	res, err := db.c.Exec(`SELECT * FROM likes WHERE posterid=? LIMIT 10`,
		f.Posterid)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the user didn't exist
		return ErrPhotoDoesNotExist
	}
	return nil
}
