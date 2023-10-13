package database

func (db *appdbimpl) GetFollowers(id uint64, limit uint64, offset uint64) ([]string, error) {
	var usernames []string

	// Plain simple SELECT query
	rows, err := db.c.Query(`
		SELECT users.username FROM followers INNER JOIN users ON followers.followerid = users.usid WHERE followers.userid = ? LIMIT ? OFFSET ?`, id, limit, offset)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// Iterate through the rows to build the list of usernames
	for rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			return nil, err
		}

		usernames = append(usernames, username)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return usernames, nil
}
