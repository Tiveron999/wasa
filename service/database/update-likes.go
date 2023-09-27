package database

func (db *appdbimpl) UpdateLikes(f Photo) error {
	res, err := db.c.Exec(`UPDATE photos SET numlikes = ? WHERE pid = ? AND posterid = ?`, f.NumLikes, f.PID, f.PosterID)
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
