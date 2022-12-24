package api

import (
	"encoding/json"
	"net/http"

	"github.com/Tiveron999/Wasa/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Read the new content for the comment from the request body.
	var comment Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !comment.CommentIsValid() {
		// Here we validated the comment structure content (e.g., location coordinates in correct range, etc.), and we
		// discovered that the comment data are not valid.
		// Note: the IsValid() function skips the ID check (see below).
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create the comment in the database. Note that this function will return a new instance of the comment with the
	// same information, plus the ID.
	dbcomment, err := rt.db.CommentPhoto(comment.CommentToDatabase())
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the comment
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't create the comment")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Here we can re-use `comment` as FromDatabase is overwriting every variabile in the structure.
	comment.CommentFromDatabase(dbcomment)

	// Send the output to the comment.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(comment)
}
