package database

func (db *appdbimpl) UploadPhoto(f Photo) (uint64, error) {
	res, err := db.c.Exec(`INSERT INTO photos (pid, image, posterid, datetime, numlikes, numcomments) VALUES (NULL, ?, ?, ?, ?, ?)`,
		f.Image, f.PosterID, f.Datetime, f.NumLikes, f.NumComments)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	f.PID = uint64(lastInsertID)
	return f.PID, nil
}
