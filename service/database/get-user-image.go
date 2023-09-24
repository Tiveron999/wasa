package database

import (
	"database/sql"
)

func (db *appdbimpl) GetUserImage(photoid uint64) ([]byte, error) {
	var image []byte

	// Plain simple SELECT query
	err := db.c.QueryRow(`SELECT image FROM photos WHERE pid=?`, photoid).Scan(&image)
	if err == sql.ErrNoRows {
		return nil, ErrPhotoDoesNotExist
	} else if err != nil {
		return nil, err
	}
	return image, nil
}
