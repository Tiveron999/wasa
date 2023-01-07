package database

func (db *appdbimpl) GetToken(id uint64) (Token, error) {
	/*
		// Prepare the SELECT statement
		stmt, err := db.Prepare("SELECT token, expiration FROM tokens WHERE userid = ?")
		if err != nil {
			return Token{}, err
		}
		defer stmt.Close()

		// Execute the SELECT statement and scan the result into a Token struct
		var token Token
		err = stmt.QueryRow(id).Scan(&token.Token, &token.ExpiresAt)
		if err != nil {
			return Token{}, err
		}

		return token, nil
	}
	*/
	var t Token
	_, err := db.c.Query(`SELECT * FROM photos WHERE posterid IN 
	(SELECT userid FROM followers WHERE followerid = ?) SORT BY datetime DESC LIMIT 10`) // , u.USID
	return t, err
}
