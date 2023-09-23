package database

func (db *appdbimpl) GetMyStream(userid uint64, datet string, limit uint64, offset uint64) ([]Photo, error) {
	var ret []Photo

	// Plain simple SELECT query
	rows, err := db.c.Query(`SELECT * FROM photos WHERE posterid IN 
	(SELECT userid FROM followers WHERE followerid = ?) AND datetime >= ? ORDER BY datetime DESC LIMIT ? OFFSET ?`,
		userid, datet, limit, offset)

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// Here we read the resultset and we build the list to be returned
	for rows.Next() {
		var f Photo
		err = rows.Scan(&f.PID, &f.Image, &f.PosterID, &f.Datetime, &f.NumLikes, &f.NumComments)
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
