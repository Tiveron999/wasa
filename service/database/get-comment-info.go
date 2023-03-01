package database

func (db *appdbimpl) GetCommentInfo(id uint64) (Comment, error) {
	var c Comment
	err := db.c.QueryRow(`SELECT * FROM comments WHERE cid=?`,
		id).Scan(&c.CID, &c.CommenterID, &c.PhotoID, &c.Text, &c.Datetime)

	if err != nil {
		return c, err
	}

	/*
		// If we didn't delete any row, then the comment didn't exist
		return f, ErrCommentDoesNotExist*/

	return c, nil
}
