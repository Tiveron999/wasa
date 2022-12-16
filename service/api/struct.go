package api

import (
	"strings"

	"github.com/Tiveron999/Wasa/service/database"
)

// User struct represents a user in every data exchange with the external world via REST API. JSON tags have been
// added to the struct to conform to the OpenAPI specifications regarding JSON key names.
// Note: there is a similar struct in the database package. See User.FromDatabase (below) to understand why.
type User struct {
	ID       uint64 `json:"id"`
	username string `json:"username"`
}

type Photo struct {
	id          uint64
	url         string
	posterid    uint64
	datetime    string
	numlikes    uint64
	numcomments uint64
}

// FromDatabase populates the struct with data from the database, overwriting all values.
// You might think this is code duplication, which is correct. However, it's "good" code duplication because it allows
// us to uncouple the database and API packages.
// Suppose we were using the "database.User" struct inside the API package; in that case, we were forced to conform
// either the API specifications to the database package or the other way around. However, very often, the database
// structure is different from the structure of the REST API.
// Also, in this way the database package is freely usable by other packages without the assumption that structs from
// the database should somehow be JSON-serializable (or, in general, serializable).
func (f *User) FromDatabase(user database.User) {
	f.ID = user.ID
	f.username = user.username
}

// ToDatabase returns the user in a database-compatible representation
func (f *User) ToDatabase() database.User {
	return database.User{
		ID:       f.ID,
		username: f.username,
	}
}

// IsValid checks the validity of the content. In particular, coordinates should be in their range of validity, and the
// status should be either UserStatusGood or UserStatusFaulty. Note that the ID is not checked, as users
// read from requests have zero IDs as the user won't send us the ID in that way.
func (f *User) IsValid() bool {
	if strings.TrimSpace(f.username) == "" {
		return true
	} else {
		return false
	}
}
