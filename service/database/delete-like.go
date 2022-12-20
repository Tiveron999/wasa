package database

func (db *appdbimpl) DeleteLike(id uint64) error {
	res, err := db.c.Exec(`DELETE FROM likes WHERE id=?`, id)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the photo didn't exist
		return ErrLikeDoesNotExist
	}
	return nil
}
