package database

func (db *appdbimpl) GetCommentInfo(f Comment) (Comment, error) {
	res, err := db.c.Exec(`SELECT * FROM comments WHERE pid=?`,
		f.CID)
	if err != nil {
		return f, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return f, err
	} else if affected == 0 {
		// If we didn't delete any row, then the comment didn't exist
		return f, ErrCommentDoesNotExist
	}
	return f, nil
}
