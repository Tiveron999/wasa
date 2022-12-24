package database

func (db *appdbimpl) DeleteComment(id uint64) error {
	res, err := db.c.Exec(`DELETE FROM comments WHERE cid=?`, id)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the comment didn't exist
		return ErrCommentDoesNotExist
	}
	return nil
}
