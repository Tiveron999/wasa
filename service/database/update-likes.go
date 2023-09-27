package database

// UpdateLikes updates the number of likes for a specific photo.
func (db *appdbimpl) UpdateLikes(photoID, posterID, numLikes uint64) error {
	res, err := db.c.Exec(`UPDATE photos SET numlikes = ? WHERE pid = ? AND posterid = ?`, numLikes, photoID, posterID)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't update any row, then the photo didn't exist
		return ErrPhotoDoesNotExist
	}

	return nil
}
