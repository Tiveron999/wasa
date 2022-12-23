package database

func (db *appdbimpl) CreatePhoto(f Photo) (Photo, error) {
	res, err := db.c.Exec(`INSERT INTO photos (pid, url, posterid, datetime, numlikes, numcomments) VALUES (NULL, ?, ?, ?, ?, ?)`,
		f.PID, f.URL, f.Posterid, f.Datetime, f.Numlikes, f.Numcomments)
	if err != nil {
		return f, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return f, err
	}

	f.PID = uint64(lastInsertID)
	return f, nil
}
