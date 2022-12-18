package database

func (db *appdbimpl) LikePhoto(f User, g Photo) (Likes, error) {
	res, err := db.c.Exec(`INSERT INTO likes (id, likerid, photoid) VALUES (NULL, ?, ?)`,
		f.ID, g.ID)
	var like Likes
	like.
	if err != nil {
		return like, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return like, err
	}

	f.ID = uint64(lastInsertID)
	return like, nil
}
