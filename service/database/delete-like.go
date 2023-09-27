package database

func (db *appdbimpl) UnlikePhoto(likerid, photoid, posterid uint64) error {
	res, err := db.c.Exec(`DELETE FROM likes WHERE likerid=? AND photoid=? AND posterid=?`, likerid, photoid, posterid)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the like didn't exist
		return ErrLikeDoesNotExist
	}
	return nil
}
