package database

func (db *appdbimpl) FollowUser(f Follower) error {
	res, err := db.c.Exec(`INSERT INTO followers (userid, followerid) VALUES (NULL, ?)`,
		f.UserID, f.FollowerID)
	if err != nil {
		return err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	f.UserID = uint64(lastInsertID)
	return nil
}
