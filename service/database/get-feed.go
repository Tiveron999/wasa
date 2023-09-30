package database

func (db *appdbimpl) GetMyStream(userid uint64, datet string, limit uint64, offset uint64) ([]Photo, error) {
	var ret []Photo

	// Plain simple SELECT query
	rows, err := db.c.Query(`SELECT pid, image, posterid, datetime, numlikes, numcomments FROM photos WHERE posterid IN 
	(SELECT userid FROM followers WHERE followerid = ?) AND datetime >= ? ORDER BY datetime DESC LIMIT ? OFFSET ?`,
		userid, datet, limit, offset)

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// Here we read the resultset and we build the list to be returned
	for rows.Next() {
		var photo Photo
		err = rows.Scan(&photo.PID, &photo.Image, &photo.PosterID, &photo.Datetime, &photo.NumLikes, &photo.NumComments)
		if err != nil {
			return nil, err
		}

		ret = append(ret, photo)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}
