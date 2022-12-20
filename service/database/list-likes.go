package database

func (db *appdbimpl) ListLikes() ([]Likes, error) {
	var ret []Likes

	// Plain simple SELECT query
	rows, err := db.c.Query(`SELECT * FROM likes LIMIT 10`)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// Here we read the resultset and we build the list to be returned
	for rows.Next() {
		var f Likes
		err = rows.Scan(&f.ID, &f.LikerID, &f.PhotoID)
		if err != nil {
			return nil, err
		}

		ret = append(ret, f)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}