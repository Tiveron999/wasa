package database

func (db *appdbimpl) CreateComment(f Comment) (Comment, error) {
	res, err := db.c.Exec(`INSERT INTO comments (cid, commenterid, f.photoid, f.datetime) VALUES (NULL, ?, ?, ?)`,
		f.CID, f.CommenterID, f.PhotoID, f.Datetime)
	if err != nil {
		return f, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return f, err
	}

	f.CID = uint64(lastInsertID)
	return f, nil
}
