package database

func (db *appdbimpl) GetFollowing(id uint64, limit uint64, offset uint64) ([]string, error) {
	var ret []string

	// Plain simple SELECT query
	rows, err := db.c.Query(
		`SELECT users.username FROM followers INNER JOIN users ON followers.userid = users.usid WHERE followers.followerid = ? LIMIT ? OFFSET ?`, id, limit, offset)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// Here we read the resultset and we build the list to be returned
	for rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			return nil, err
		}

		ret = append(ret, username)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}
