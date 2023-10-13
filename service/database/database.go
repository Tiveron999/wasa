package database

import (
	"database/sql"
	"errors"
	"fmt"
)

var ErrUserDoesNotExist = errors.New("user does not exist")

var ErrPhotoDoesNotExist = errors.New("photo does not exist")

var ErrLikeDoesNotExist = errors.New("like does not exist")

var ErrCommentDoesNotExist = errors.New("comment does not exist")

var ErrFollowerDoesNotExist = errors.New("follower does not exist")

var ErrBannedDoesNotExist = errors.New("banned user does not exist")

// User struct represents a user in every API call between this package and the outside world.
// Note that the internal representation of a user in the database might be different.
type User struct {
	USID     uint64
	Username string
}

type Photo struct {
	PID         uint64
	Image       []byte
	PosterID    uint64
	Datetime    string
	NumLikes    uint64
	NumComments uint64
}

type Likes struct {
	LID      uint64
	LikerID  uint64
	PhotoID  uint64
	PosterID uint64
}

type Comment struct {
	CID         uint64
	CommenterID uint64
	PhotoID     uint64
	Text        string
	Datetime    string
}

type Follower struct {
	UserID     uint64
	FollowerID uint64
}

type Banned struct {
	UserID   uint64
	BannedID uint64
}

/* type Profile struct {
	PRID      uint64
	Username  string
	Photos    []string
	NumPhotos uint64
	Followers []string
	Following []string
}*/

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	// ListUsers returns the list of users with their status
	// ListUsers() ([]User, error)

	// CreateUser creates a new user in the database. It returns an updated User object (with the ID)
	CreateUser(username string) (uint64, error)

	// SetMyUserName updates the user, replacing every value with those provided in the argument
	SetMyUserName(User) error

	// DeleteUser removes the user with the given ID
	// DeleteUser(id uint64) error

	GetUserInfo(User) (User, error)

	GetUserPhotos(uint64, string, uint64, uint64) ([]Photo, error)

	GetPhotoInfo(Photo) (Photo, error)

	UploadPhoto(Photo) (uint64, error)

	DeletePhoto(uint64) error

	GetLikesList(uint64, uint64) ([]string, error)

	LikePhoto(uint64, uint64, uint64) error

	UnlikePhoto(uint64, uint64, uint64) error

	CommentPhoto(Comment) error

	GetCommentsList(uint64) ([]Comment, error)

	GetCommentInfo(uint64) (Comment, error)

	UncommentPhoto(uint64) error

	GetFollowers(uint64, uint64, uint64) ([]string, error)

	GetBannedList() ([]Banned, error)

	UnbanUser(uint64, uint64) error

	BanUser(Banned) (Banned, error)

	GetFollowing(uint64, uint64, uint64) ([]Follower, error)

	FollowUser(Follower) error

	GetMyStream(uint64, string, uint64, uint64) ([]Photo, error)

	UnfollowUser(Follower) error

	// GetUserProfile(uint64) (Profile, error)

	// GetSearchedUserInfo(username string) (Profile, error)

	// GetToken(uint64) (Token, error)

	GetUserID(string) (uint64, error)

	CheckUserExists(string) (uint64, error)

	GetUserImage(uint64) ([]byte, error)

	GetUsername(uint64) (string, error)

	UpdateLikes(uint64, uint64, uint64) error

	UpdateComments(uint64, uint64) error

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
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&users)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE users (
    usid INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT UNIQUE NOT NULL
);`
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
    image BLOB NOT NULL,
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
	posterid INTEGER NOT NULL,
	FOREIGN KEY(likerid) REFERENCES users(usid) ON UPDATE CASCADE ON DELETE CASCADE,
	FOREIGN KEY(photoid) REFERENCES photos(pid) ON UPDATE CASCADE ON DELETE CASCADE,
	FOREIGN KEY(posterid) REFERENCES users(usid) ON UPDATE CASCADE ON DELETE CASCADE
);`
		_, err2 = db.Exec(sqlStmt)
		if err2 != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err2)
		}
	}

	var comments string
	_, _ = db.Exec("PRAGMA foreign_keys = ON;")

	err3 := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='comments';`).Scan(&comments)
	if errors.Is(err3, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE comments (
	cid INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	commenterid INTEGER NOT NULL,
	photoid INTEGER NOT NULL,
	text TEXT NOT NULL,
	datetime TEXT NOT NULL,
	FOREIGN KEY(commenterid) REFERENCES users(usid) ON UPDATE CASCADE ON DELETE CASCADE,
	FOREIGN KEY(photoid) REFERENCES photos(pid) ON UPDATE CASCADE ON DELETE CASCADE
);`
		_, err3 = db.Exec(sqlStmt)
		if err3 != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err3)
		}
	}

	var followers string
	_, _ = db.Exec("PRAGMA foreign_keys = ON;")

	err4 := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='followers';`).Scan(&followers)
	if errors.Is(err4, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE followers (
	userid INTEGER NOT NULL,
	followerid INTEGER NOT NULL,
	FOREIGN KEY(userid) REFERENCES users(usid) ON UPDATE CASCADE ON DELETE CASCADE,
	FOREIGN KEY(followerid) REFERENCES users(usid) ON UPDATE CASCADE ON DELETE CASCADE,
	PRIMARY KEY (userid, followerid)
);`
		_, err4 = db.Exec(sqlStmt)
		if err4 != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err4)
		}
	}

	var banned string
	_, _ = db.Exec("PRAGMA foreign_keys = ON;")

	err5 := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='banned';`).Scan(&banned)
	if errors.Is(err5, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE banned (
	userid INTEGER NOT NULL,
	bannedid INTEGER NOT NULL,
	FOREIGN KEY(userid) REFERENCES users(usid) ON UPDATE CASCADE ON DELETE CASCADE,
	FOREIGN KEY(bannedid) REFERENCES users(usid) ON UPDATE CASCADE ON DELETE CASCADE,
	PRIMARY KEY (userid, bannedid)
);`
		_, err5 = db.Exec(sqlStmt)
		if err5 != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err5)
		}
	}

	/* var profiles string
		_, _ = db.Exec("PRAGMA foreign_keys = ON;")

		err6 := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='profiles';`).Scan(&profiles)
		if errors.Is(err6, sql.ErrNoRows) {
			sqlStmt := `CREATE TABLE profiles (
		prid INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		photos TEXT NOT NULL,
		numphotos INTEGER NOT NULL,
		followers TEXT NOT NULL,
		following TEXT NOT NULL,
		FOREIGN KEY(prid) REFERENCES users(usid) ON UPDATE CASCADE ON DELETE CASCADE
	);`
			_, err6 = db.Exec(sqlStmt)
			if err6 != nil {
				return nil, fmt.Errorf("error creating database structure: %w", err6)
			}
		} */

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
