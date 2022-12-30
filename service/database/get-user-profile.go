package database

func (db *appdbimpl) GetUserProfile(f Profile) (Profile, error) {
	res, err := db.c.Exec(`INSERT INTO profiles (prid, username, photos, numphotos, followers, following) VALUES (NULL, ?, ?, ?, ?, ?)`,
		f.PRID)
	if err != nil {
		return f, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return f, err
	}

	f.PRID = uint64(lastInsertID)
	return f, nil
}

/*var p Profile
res, err := db.c.Exec(`SELECT * FROM profiles WHERE prid=?`, f.USID)
if err != nil {
	return p, err
}

affected, err := res.RowsAffected()
if err != nil {
	return f, err
} else if affected == 0 {
	// If we didn't delete any row, then the user didn't exist
	return f, ErrUserDoesNotExist
}
return f, nil*/

/*var p Profile
	var ret []Profile

	// Plain simple SELECT query
	profs, err := db.c.Query(`SELECT * FROM profiles WHERE prid=?`, f.USID)
	if err != nil {
		return p, err
	}
	defer func() { _ = profs.Close() }()

	// Here we read the resultset and we build the list to be returned
	for profs.Next() {
		var f Profile
		err = profs.Scan(&f.PRID, &f.Username, &f.NumPhotos) // &f.Photos, &f.Followers, &f.Following
		if err != nil {
			return f, err
		}

		ret = append(ret, f)
	}
	if err = profs.Err(); err != nil {
		return p, err
	}


	// Plain simple SELECT query
	rows, err := db.c.Query(`SELECT bannedid FROM banned WHERE userid=? LIMIT 10`)
	if err != nil {
		return p, err
	}
	defer func() { _ = rows.Close() }()

	// Here we read the resultset and we build the list to be returned
	for rows.Next() {
		var f Profile
		err = rows.Scan(&f.PRID, &f.Username, &f.NumPhotos)
		if err != nil {
			return p, err
		}

		ret = append(ret, f)
	}
	if err = rows.Err(); err != nil {
		return p, err
	}

	return err, nil
}*/
