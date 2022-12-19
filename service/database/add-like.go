package database

func (db *appdbimpl) LikePhoto(f User, g Photo) (Likes, error) {
	res, err := db.c.Exec(`INSERT INTO likes (id, likerid, photoid) VALUES (NULL, ?, ?)`,
		f.ID, g.ID)
	var like Likes
	/*rows, err := db.Query(`INSERT INTO likes (id, likerid, photoid) VALUES (NULL, ?, ?)`,
	f.ID, g.ID) // db *sql.DB in func
	// A like slice to hold data from returned rows.
    var likes []Likes

    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var like Likes
        if err := rows.Scan(&like.ID, &like.LikerID, &like.PhotoID); err != nil {
            return likes, err
        }
        likes = append(likes, like)
    }
    if err = rows.Err(); err != nil {
        return likes, err
    }
    return likes, nil
	
*/
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
