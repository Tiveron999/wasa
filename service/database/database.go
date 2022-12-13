package database

import (
	"database/sql"
	"errors"
	"fmt"
)

var ErrUserDoesNotExist = errors.New("User does not exist")

// Fountain struct represent a fountain in every API call between this package and the outside world.
// Note that the internal representation of fountain in the database might be different.
type User struct {
	ID       uint64
	username string
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	// ListFountains returns the list of fountains with their status
	// ListFountains() ([]Fountain, error)

	// CreateFountain creates a new fountain in the database. It returns an updated Fountain object (with the ID)
	CreateUser(User) (User, error)

	// UpdateFountain updates the fountain, replacing every value with those provided in the argument
	UpdateUser(User) error

	// DeleteFountain removes the fountain with the given ID
	// DeleteFountain(id uint64) error

	GetUserInfo(User) error

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

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var users string
	err := db.QueryRow(`SELECT username FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&users)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY,
    username TEXT NOT NULL
);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	var photos string
	_, _ = db.Exec("PRAGMA foreign_keys = ON;")

	err1 := db.QueryRow(`SELECT url FROM sqlite_master WHERE type='table' AND name='photos';`).Scan(&photos)
	if errors.Is(err1, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE photos (
    id INTEGER NOT NULL PRIMARY KEY,
    url TEXT NOT NULL,
	poster INTEGER NOT NULL foreign key,
	datetime TEXT NOT NULL,
	numlikes INTEGER NOT NULL,
	numcomments INTEGER NOT NULL
);`
		_, err1 = db.Exec(sqlStmt)
		if err1 != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err1)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
