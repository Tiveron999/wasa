package database

func (db *appdbimpl) CreatePhoto(f Photo) (Photo, error) {
	res, err := db.c.Exec(`INSERT INTO photos (id, url, posterid, datetime, numlikes, numcomments) VALUES (?, ?, ?, ?, ?, ?)`,
		f.id, f.url, f.posterid, f.datetime, f.numlikes, f.numcomments)
	if err != nil {
		return f, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return f, err
	}

	f.id = uint64(lastInsertID)
	return f, nil
}
