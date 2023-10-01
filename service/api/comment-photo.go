package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Tiveron999/Wasa/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the poster ID and photo ID from the path parameters

	/*posterID, err := strconv.ParseUint(ps.ByName("user_id"), 10, 64)                     useful for the ban
	if err != nil {
		// Handle the error, e.g., invalid photo ID
		w.WriteHeader(http.StatusBadRequest)
		return
	}*/

	photoID, err := strconv.ParseUint(ps.ByName("photo_id"), 10, 64)
	if err != nil {
		// Handle the error, e.g., invalid photo ID
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the Authorization header value
	authorizationHeader := r.Header.Get("Authorization")

	// Get the user ID from the Authorization header
	commenterID, isValid := idFromAuthorizationHeader(authorizationHeader)
	if !isValid {
		// If the Authorization header is invalid or not present, handle the error
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get the comment text
	var comment Comment

	// Access the Text field
	commentText := comment.Text

	err = json.NewDecoder(r.Body).Decode(&commentText)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	comment = Comment{
		CID:         0,
		CommenterID: commenterID,
		PhotoID:     photoID,
		Text:        commentText,
		Datetime:    time.Now().Format("2006-01-02 15:04:05"),
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

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(comment)
}
