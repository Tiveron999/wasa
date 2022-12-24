package database

func (db *appdbimpl) ListFollowers() ([]Follower, error) {
	var ret []Follower

	// Plain simple SELECT query
	rows, err := db.c.Query(`SELECT followerid FROM followers WHERE userid=? LIMIT 10`)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// Here we read the resultset and we build the list to be returned
	for rows.Next() {
		var f Follower
		err = rows.Scan(&f.userid, &f.followerid)
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
