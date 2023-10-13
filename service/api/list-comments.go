package api

import (
	"encoding/json"
	// "fmt"
	"net/http"
	"strconv"

	"github.com/Tiveron999/Wasa/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getCommentsList(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	photoID, err := strconv.ParseUint(ps.ByName("photo_id"), 10, 64)
	if err != nil {
		// Handle the error, e.g., invalid photo ID
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Fetch comments from the database based on 'id'
	comments, err := rt.db.GetCommentsList(photoID)
	if err != nil {
		// Handle database error
		ctx.Logger.WithError(err).Error("can't list comments")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// fmt.Println(comments)

	// Convert comments to the desired frontend structure if needed
	var frontendComments []Comment
	for _, c := range comments {
		var co Comment
		co.CommentFromDatabase(c)
		frontendComments = append(frontendComments, co)
	}

	// Send the list of comments to the client as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(frontendComments); err != nil {
		ctx.Logger.WithError(err).Error("can't encode comments to JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
