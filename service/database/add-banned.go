package database

func (db *appdbimpl) BanUser(f Banned) (Banned, error) {
	res, err := db.c.Exec(`INSERT INTO banned (userid, bannedid) VALUES (?, ?)`, f.UserID, f.BannedID)
	if err != nil {
		return f, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return f, err
	}

	f.BannedID = uint64(lastInsertID)
	return f, nil
}
