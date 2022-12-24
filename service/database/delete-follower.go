package database

func (db *appdbimpl) DeleteFollower(userid uint64, followerid uint64) error {
	res, err := db.c.Exec(`DELETE FROM followers WHERE userid=? AND followerid=?`, userid, followerid)
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
