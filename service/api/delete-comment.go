package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Tiveron999/Wasa/service/api/reqcontext"
	"github.com/Tiveron999/Wasa/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// The Comment ID in the path is a 64-bit unsigned integer. Let's parse it.
	commentID, err := strconv.ParseUint(ps.ByName("comment_id"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rt.db.UncommentPhoto(commentID)
	if errors.Is(err, database.ErrCommentDoesNotExist) {
		// The comment (indicated by `id`) does not exist, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" ins ide the "ctx" (context) because the scope of this issue is the request.
		// Note (2): we are adding the error and an additional field (`id`) to the log entry, so that we will receive
		// the identifier of the comment that triggered the error.
		ctx.Logger.WithError(err).WithField("cid", commentID).Error("can't delete the comment")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	photoID, err := strconv.ParseUint(ps.ByName("photo_id"), 10, 64)
	if err != nil {
		// Handle the error, e.g., invalid photo ID
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Request an unfiltered list of comments from the DB
	comments, err := rt.db.GetCommentsList(photoID)
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the photo
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't list comments")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	commentsNumber := uint64(len(comments))

	err = rt.db.UpdateComments(photoID, commentsNumber)
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't update the comment number")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to indicate JSON response
	w.Header().Set("Content-Type", "application/json")

	// Write the response with the updated comment count
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]int{"comments": int(commentsNumber)})
}
