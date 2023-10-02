package database

// UpdateComments updates the number of comments for a specific photo.
func (db *appdbimpl) UpdateComments(photoID, numComments uint64) error {
	res, err := db.c.Exec(`UPDATE photos SET numcomments = ? WHERE pid = ?`, numComments, photoID)
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
