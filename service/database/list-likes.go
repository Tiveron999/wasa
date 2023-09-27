package database

// GetLikesList retrieves the number of likes for a specific photo.
func (db *appdbimpl) GetLikesList(photoID, posterID uint64) ([]uint64, error) {
	var userIDs []uint64

	// Execute a query to retrieve user IDs who liked the given photo and poster.
	rows, err := db.c.Query("SELECT likerid FROM likes WHERE photoid = ? AND posterid = ?", photoID, posterID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the rows to build the list of user IDs.
	for rows.Next() {
		var likerID uint64
		if err := rows.Scan(&likerID); err != nil {
			return nil, err
		}

		userIDs = append(userIDs, likerID)
	}

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return userIDs, nil
}
