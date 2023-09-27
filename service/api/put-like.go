package api

import (
	"net/http"
	"strconv"

	"github.com/Tiveron999/Wasa/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Read the new content for the like from the request body.
	var like Likes

	// Get the user ID and photo ID from the path parameters
	posterID, err := strconv.ParseUint(ps.ByName("user_id"), 10, 64)
	if err != nil {
		// Handle the error, e.g., invalid photo ID
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the user ID and photo ID from the path parameters
	photoID, err := strconv.ParseUint(ps.ByName("photo_id"), 10, 64)
	if err != nil {
		// Handle the error, e.g., invalid photo ID
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	likerID, err := strconv.ParseUint(ps.ByName("liker_id"), 10, 64)
	if err != nil {
		// Handle the error, e.g., invalid user ID
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Build the like object with additional fields
	like.PosterID = posterID
	like.LikerID = likerID
	like.PhotoID = photoID

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

	// Request an unfiltered list of likes from the DB
	likes, err := rt.db.GetLikesList(photoID, posterID)

	likesNumber := uint64(len(likes))

	err2 := rt.db.UpdateLikes(photoID, posterID, likesNumber)
	if err2 != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't update the like number")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Here we can re-use `like` as FromDatabase is overwriting every variable in the structure.
	like.LikeFromDatabase(dblike)

	w.WriteHeader(http.StatusNoContent)
}
