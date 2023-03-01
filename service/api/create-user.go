package api

import (
	"encoding/json"
	"net/http"

	"github.com/Tiveron999/Wasa/service/api/reqcontext"
	"github.com/Tiveron999/Wasa/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) createUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Read the new content for the user from the request body.
	type loginData struct {
		Username string `json:"username"`
		ID       string `json:"id"`
	}

	var t loginData
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var user User
	// user := User{USID: uint64(lastInsertID), Username: username}
	// Create the user in the database. Note that this function will return a new instance of the user with the
	// same information, plus the ID.
	dbuser, err := rt.db.CreateUser(user.Username)
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't create the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Here we can re-use `user` as FromDatabase is overwriting every variabile in the structure.
	user.UserFromDatabase(database.User{})

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)
	_ = json.NewEncoder(w).Encode(dbuser)
}
