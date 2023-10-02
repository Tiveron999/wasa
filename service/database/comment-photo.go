package database

func (db *appdbimpl) CommentPhoto(f Comment) error {
	_, err := db.c.Exec(`INSERT INTO comments (cid, commenterid, photoid, text, datetime) VALUES (NULL, ?, ?, ?, ?)`,
		f.CommenterID, f.PhotoID, f.Text, f.Datetime)

	if err != nil {
		return err
	}

	return nil
}
