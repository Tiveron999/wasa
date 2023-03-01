package api

import (
	"encoding/json"
	"net/http"

	"github.com/Tiveron999/Wasa/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Read the new content for the like from the request body.
	var like Likes
	err := json.NewDecoder(r.Body).Decode(&like)
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !like.LikeIsValid() {
		// Here we validated the like structure content (e.g., location coordinates in correct range, etc.), and we
		// discovered that the like data are not valid.
		// Note: the IsValid() function skips the ID check (see below).
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create the like in the database. Note that this function will return a new instance of the like with the
	// same information, plus the ID.
	dblike, err := rt.db.LikePhoto(like.LikeToDatabase())
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the like
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't create the like")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Here we can re-use `like` as FromDatabase is overwriting every variabile in the structure.
	like.LikeFromDatabase(dblike)

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(like)
}
