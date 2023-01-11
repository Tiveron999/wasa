package database

func (db *appdbimpl) GetSearchedUserInfo(username string) (Profile, error) {
	var p Profile
	err := db.c.QueryRow("SELECT prid, username, numphotos FROM profiles WHERE username = ?", username).Scan(
		&p.PRID, &p.Username, &p.NumPhotos)

	if err != nil {
		return p, err
	}

	// Query the database for the user's photos
	rows, err := db.c.Query("SELECT filepath FROM photos WHERE userid = ?", p.PRID)

	if err != nil {
		return p, err
	}

	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var filepath string
		err = rows.Scan(&filepath)
		if err != nil {
			return p, err
		}
		p.Photos = append(p.Photos, filepath)
	}
	if err = rows.Err(); err != nil {
		return p, err
	}

	// Query the database for the user's followers
	rows, err = db.c.Query("SELECT username FROM followers JOIN users ON followers.followerid = users.id WHERE userid = ?", p.PRID)
	if err != nil {
		return p, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var followerUsername string
		err = rows.Scan(&followerUsername)
		if err != nil {
			return p, err
		}
		p.Followers = append(p.Followers, followerUsername)
	}
	if err = rows.Err(); err != nil {
		return p, err
	}

	// Query the database for the users the user is following
	rows, err = db.c.Query("SELECT username FROM following JOIN users ON following.followingid = users.id WHERE userid = ?", p.PRID)
	if err != nil {
		return p, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var followingUsername string
		err = rows.Scan(&followingUsername)
		if err != nil {
			return p, err
		}
		p.Following = append(p.Following, followingUsername)
	}
	if err = rows.Err(); err != nil {
		return p, err
	}

	return p, nil
}
