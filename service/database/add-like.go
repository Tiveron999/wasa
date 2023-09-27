package database

func (db *appdbimpl) LikePhoto(posterID, likerID, photoID uint64) error {
	_, err := db.c.Exec(`INSERT INTO likes (lid, likerid, photoid, posterid) VALUES (NULL, ?, ?, ?)`,
		likerID, photoID, posterID)

	if err != nil {
		return err
	}

	return nil
}
