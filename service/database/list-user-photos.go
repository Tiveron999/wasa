package database

func (db *appdbimpl) GetUserPhotos(id uint64, datet string, limit uint64, offset uint64) ([]string, error) {
	var ret []string

	// Plain simple SELECT query
	rows, err := db.c.Query(`SELECT url FROM photos WHERE posterid=? AND datetime >= ? ORDER BY datetime DESC LIMIT ? OFFSET ?`,
		id, datet, limit, offset)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// Here we read the resultset and we build the list to be returned
	for rows.Next() {
		var url string
		err = rows.Scan(&url)
		if err != nil {
			return nil, err
		}

		ret = append(ret, url)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}
