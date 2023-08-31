package database

func (db *appdbimpl) FollowUser(f Follower) error {
	res, err := db.c.Exec(`INSERT INTO followers (userid, followerid) VALUES (?, ?)`,
		f.UserID, f.FollowerID)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return nil
	}

	return nil
}
