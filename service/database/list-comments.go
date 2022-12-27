package database

func (db *appdbimpl) ListComments() ([]Comment, error) {
	var ret []Comment

	// Plain simple SELECT query
	rows, err := db.c.Query(`SELECT * FROM comments LIMIT 10`)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// Here we read the resultset and we build the list to be returned
	for rows.Next() {
		var f Comment
		err = rows.Scan(&f.CID, &f.CommenterID, &f.PhotoID, &f.Text, &f.Datetime)
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
