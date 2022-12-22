package database

func (db *appdbimpl) LikePhoto(f Likes) (Likes, error) {
	res, err := db.c.Exec(`INSERT INTO likes (id, likerid, photoid) VALUES (NULL, ?, ?)`,
		f.LikerID, f.PhotoID)
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
		return f, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return f, err
	}

	f.ID = uint64(lastInsertID)
	return f, nil
}
