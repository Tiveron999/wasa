package database

func (db *appdbimpl) LikePhoto(f Likes) (Likes, error) {
	res, err := db.c.Exec(`INSERT INTO likes (lid, likerid, photoid) VALUES (NULL, ?, ?)`,
		f.LikerID, f.PhotoID)
	
	if err != nil {
		return f, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return f, err
	}

	f.LID = uint64(lastInsertID)
	return f, nil
}
