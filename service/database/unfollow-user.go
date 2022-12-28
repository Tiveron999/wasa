package database

func (db *appdbimpl) UnfollowUser(userid uint64, followingid uint64) error {
	res, err := db.c.Exec(`DELETE FROM followers WHERE userid=? AND followerid=?`, followingid, userid)
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
