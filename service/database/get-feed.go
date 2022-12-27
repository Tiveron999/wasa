package database

func (db *appdbimpl) GetFeed() ([]Photo, error) {
	var ret []Photo

	// Plain simple SELECT query
	rows, err := db.c.Query(`SELECT * FROM photos WHERE posterid IN 
	(SELECT userid FROM followers WHERE followerid = ?) SORT BY datetime DESC LIMIT 10`) // , u.USID

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// Here we read the resultset and we build the list to be returned
	for rows.Next() {
		var f Photo
		err = rows.Scan(&f.PID, &f.URL, &f.Posterid, &f.Datetime, &f.Numlikes, &f.Numcomments)
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
