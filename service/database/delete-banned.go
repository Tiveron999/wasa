package database

func (db *appdbimpl) UnbanUser(userid uint64, bannedid uint64) error {
	res, err := db.c.Exec(`DELETE FROM banned WHERE userid=? AND bannedid=?`, userid, bannedid)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the banned didn't exist
		return ErrBannedDoesNotExist
	}
	return nil
}
