package database

func (db *appdbimpl) UnfollowUser(f Follower) error {
	res, err := db.c.Exec(`DELETE FROM followers WHERE userid=? AND followerid=?`, f.UserID, f.FollowerID)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the follower didn't exist
		return ErrFollowerDoesNotExist
	}
	return nil
}
