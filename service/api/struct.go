package api

import (
	"strings"

	"github.com/Tiveron999/Wasa/service/database"
)

// User struct represents a user in every data exchange with the external world via REST API. JSON tags have been
// added to the struct to conform to the OpenAPI specifications regarding JSON key names.
// Note: there is a similar struct in the database package. See User.FromDatabase (below) to understand why.
type User struct {
	USID     uint64 `json:"usid"`
	Username string `json:"username"`
}

type Photo struct {
	PID         uint64 `json:"pid"`
	URL         string `json:"url"`
	Posterid    uint64 `json:"posterid"`
	Datetime    string `json:"datetime"`
	Numlikes    uint64 `json:"numlikes"`
	Numcomments uint64 `json:"numcomments"`
}

type Likes struct {
	LID     uint64 `json:"lid"`
	LikerID uint64 `json:"likerid"`
	PhotoID uint64 `json:"photoid"`
}

type Comment struct {
	CID         uint64 `json:"cid"`
	CommenterID uint64 `json:"commenterid"`
	PhotoID     uint64 `json:"photoid"`
	Text        string `json:"text"`
	Datetime    string `json:"datetime"`
}

type Follower struct {
	UserID     uint64 `json:"userid"`
	FollowerID uint64 `json:"followerid"`
}

type Banned struct {
	UserID   uint64 `json:"userid"`
	BannedID uint64 `json:"bannedid"`
}

type Profile struct {
	UserID    uint64   `json:"userid"`
	Username  string   `json:"username"`
	Photos    []string `json:"photos"`
	NumPhotos uint64   `json:"numphotos"`
	Followers uint64   `json:"followers"`
	Following uint64   `json:"following"`
}

// FromDatabase populates the struct with data from the database, overwriting all values.
// You might think this is code duplication, which is correct. However, it's "good" code duplication because it allows
// us to uncouple the database and API packages.
// Suppose we were using the "database.User" struct inside the API package; in that case, we were forced to conform
// either the API specifications to the database package or the other way around. However, very often, the database
// structure is different from the structure of the REST API.
// Also, in this way the database package is freely usable by other packages without the assumption that structs from
// the database should somehow be JSON-serializable (or, in general, serializable).
func (f *User) UserFromDatabase(user database.User) {
	f.USID = user.USID
	f.Username = user.Username
}

// ToDatabase returns the user in a database-compatible representation
func (f *User) UserToDatabase() database.User {
	return database.User{
		USID:     f.USID,
		Username: f.Username,
	}
}

// IsValid checks the validity of the content. In particular, coordinates should be in their range of validity, and the
// status should be either UserStatusGood or UserStatusFaulty. Note that the ID is not checked, as users
// read from requests have zero IDs as the user won't send us the ID in that way.
func (f *User) UserIsValid() bool {
	if strings.TrimSpace(f.Username) != "" {
		return true
	} else {
		return false
	}
}

func (f *Photo) PhotoFromDatabase(photo database.Photo) {
	f.PID = photo.PID
	f.URL = photo.URL
	f.Posterid = photo.Posterid
	f.Datetime = photo.Datetime
	f.Numlikes = photo.Numlikes
	f.Numcomments = photo.Numcomments
}

// ToDatabase returns the photo in a database-compatible representation
func (f *Photo) PhotoToDatabase() database.Photo {
	return database.Photo{
		PID:         f.PID,
		URL:         f.URL,
		Posterid:    f.Posterid,
		Datetime:    f.Datetime,
		Numlikes:    f.Numlikes,
		Numcomments: f.Numcomments,
	}
}

// IsValid checks the validity of the content. In particular, coordinates should be in their range of validity, and the
// status should be either UserStatusGood or UserStatusFaulty. Note that the ID is not checked, as users
// read from requests have zero IDs as the user won't send us the ID in that way.
func (f *Photo) PhotoIsValid() bool {
	if strings.TrimSpace(f.URL) != "" {
		return true
	} else {
		return false
	}
}

func (f *Likes) LikeFromDatabase(like database.Likes) {
	f.LID = like.LID
	f.LikerID = like.LikerID
	f.PhotoID = like.PhotoID
}

// ToDatabase returns the like in a database-compatible representation
func (f *Likes) LikeToDatabase() database.Likes {
	return database.Likes{
		LID:     f.LID,
		LikerID: f.LikerID,
		PhotoID: f.PhotoID,
	}
}

func (f *Likes) LikeIsValid() bool {
	return true
}

func (f *Comment) CommentFromDatabase(comment database.Comment) {
	f.CID = comment.CID
	f.CommenterID = comment.CommenterID
	f.PhotoID = comment.PhotoID
	f.Datetime = comment.Datetime
}

// ToDatabase returns the comment in a database-compatible representation
func (f *Comment) CommentToDatabase() database.Comment {
	return database.Comment{
		CID:         f.CID,
		CommenterID: f.CommenterID,
		PhotoID:     f.PhotoID,
		Datetime:    f.Datetime,
	}
}

func (f *Comment) CommentIsValid() bool {
	return true
}

func (f *Follower) FollowerFromDatabase(follower database.Follower) {
	f.UserID = follower.UserID
	f.FollowerID = follower.FollowerID
}

// ToDatabase returns the follower in a database-compatible representation
func (f *Follower) FollowerToDatabase() database.Follower {
	return database.Follower{
		UserID:     f.UserID,
		FollowerID: f.FollowerID,
	}
}

func (f *Follower) FollowerIsValid() bool {
	return true
}

func (f *Banned) BannedFromDatabase(banned database.Banned) {
	f.UserID = banned.UserID
	f.BannedID = banned.BannedID
}

// ToDatabase returns the banned in a database-compatible representation
func (f *Banned) BannedToDatabase() database.Banned {
	return database.Banned{
		UserID:   f.UserID,
		BannedID: f.BannedID,
	}
}

func (f *Banned) BannedIsValid() bool {
	return true
}

/* func (f *Profile) ProfileFromDatabase(profile database.Profile) {
	f.PRID = profile.PRID
	f.Username = profile.Username
	f.NumPhotos = profile.NumPhotos
}

// ToDatabase returns the profile in a database-compatible representation
func (f *Profile) ProfileToDatabase() database.Profile {
	return database.Profile{
		PRID:      f.PRID,
		Username:  f.Username,
		NumPhotos: f.NumPhotos,
	}
} */

func (f *Profile) ProfileIsValid() bool {
	return true
}
