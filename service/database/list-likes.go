package database

func (db *appdbimpl) GetLikesList(photoID, posterID uint64) ([]string, error) {
	var usernames []string

	// Execute a query to retrieve usernames of users who liked the given photo and poster.
	rows, err := db.c.Query(`
		SELECT users.username FROM likes INNER JOIN users ON likes.likerid = users.usid WHERE likes.photoid = ? AND likes.posterid = ?`, photoID, posterID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the rows to build the list of usernames.
	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			return nil, err
		}

		usernames = append(usernames, username)
	}

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return usernames, nil
}
