package database

func (db *appdbimpl) GetBannedList() ([]Banned, error) {
	var ret []Banned

	// Plain simple SELECT query
	rows, err := db.c.Query(`SELECT userid, bannedid FROM banned`)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// Here we read the resultset and we build the list to be returned
	for rows.Next() {
		var f Banned
		err = rows.Scan(&f.UserID, &f.BannedID)
		if err != nil {
			return nil, err
		}

		ret = append(ret, f)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}
