package api

import (
	"encoding/json"
	"net/http"

	"github.com/Tiveron999/Wasa/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Read the new content for the following from the request body.
	var following Follower
	err := json.NewDecoder(r.Body).Decode(&following)
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !following.FollowerIsValid() {
		// Here we validated the following structure content (e.g., location coordinates in correct range, etc.), and we
		// discovered that the following data are not valid.
		// Note: the IsValid() function skips the ID check (see below).
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create the following in the database. Note that this function will return a new instance of the following with the
	// same information, plus the ID.
	dbfollowing, err := rt.db.FollowUser(following.FollowerToDatabase())
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the following
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't create the following")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Here we can re-use `following` as FromDatabase is overwriting every variabile in the structure.
	following.FollowerFromDatabase(dbfollowing)

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(following)
}
