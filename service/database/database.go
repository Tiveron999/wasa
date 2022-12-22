package database

import (
	"database/sql"
	"errors"
	"fmt"
)

var ErrUserDoesNotExist = errors.New("user does not exist")

var ErrPhotoDoesNotExist = errors.New("photo does not exist")

var ErrLikeDoesNotExist = errors.New("like does not exist")

// User struct represents a user in every API call between this package and the outside world.
// Note that the internal representation of a user in the database might be different.
type User struct {
	USID     uint64
	Username string
}

type Photo struct {
	PID         uint64
	URL         string
	Posterid    uint64
	Datetime    string
	Numlikes    uint64
	Numcomments uint64
}

type Likes struct {
	LID     uint64
	LikerID uint64
	PhotoID uint64
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	// ListUsers returns the list of users with their status
	// ListUsers() ([]User, error)

	// CreateUser creates a new user in the database. It returns an updated User object (with the ID)
	CreateUser(User) (User, error)

	// UpdateUser updates the user, replacing every value with those provided in the argument
	UpdateUser(User) error

	// DeleteUser removes the user with the given ID
	// DeleteUser(id uint64) error

	GetUserInfo(User) (User, error)

	ListUserPhotos() ([]Photo, error)

	GetPhotoInfo(Photo) (Photo, error)

	CreatePhoto(Photo) (Photo, error)

	DeletePhoto(uint64) error

	ListLikes() ([]Likes, error)

	LikePhoto(Likes) (Likes, error)

	DeleteLike(uint64) error

	// Ping checks whether the database is available or not (in that case, an error will be returned)
	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building an AppDatabase")
	}
	//fmt.Print("something")
	// Check if table exists. If not, the database is empty, and we need to create the structure
	var users string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&users)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE users (
    usid INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL
);`
		fmt.Print(sqlStmt)
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	var photos string
	_, _ = db.Exec("PRAGMA foreign_keys = ON;")

	err1 := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='photos';`).Scan(&photos)
	if errors.Is(err1, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE photos (
    pid INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    url TEXT NOT NULL,
	posterid INTEGER NOT NULL,
	datetime TEXT NOT NULL,
	numlikes INTEGER NOT NULL,
	numcomments INTEGER NOT NULL,
	FOREIGN KEY(posterid) REFERENCES users(usid) ON UPDATE CASCADE ON DELETE CASCADE
);`

		_, err1 = db.Exec(sqlStmt)
		if err1 != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err1)
		}
	}

	var likes string
	_, _ = db.Exec("PRAGMA foreign_keys = ON;")

	err2 := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='likes';`).Scan(&likes)
	if errors.Is(err2, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE likes (
	lid INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	likerid INTEGER NOT NULL,
	photoid INTEGER NOT NULL,
	FOREIGN KEY(likerid) REFERENCES users(usid) ON UPDATE CASCADE ON DELETE CASCADE,
	FOREIGN KEY(photoid) REFERENCES photos(pid) ON UPDATE CASCADE ON DELETE CASCADE
);`
		// PRIMARY KEY (likerid, photoid)
		_, err2 = db.Exec(sqlStmt)
		if err2 != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err2)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
