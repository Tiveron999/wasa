package database

func (db *appdbimpl) ListFollowing() ([]Follower, error) {
	var ret []Follower

	// Plain simple SELECT query
	rows, err := db.c.Query(`SELECT userid FROM followers WHERE followerid=? LIMIT 10`)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// Here we read the resultset and we build the list to be returned
	for rows.Next() {
		var f Follower
		err = rows.Scan(&f.UserID, &f.FollowerID)
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
