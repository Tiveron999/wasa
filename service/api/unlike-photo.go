package api

import (
	"encoding/json"
	"errors"

	// "fmt"
	"net/http"
	"strconv"

	"github.com/Tiveron999/Wasa/service/api/reqcontext"
	"github.com/Tiveron999/Wasa/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Get the poster ID, photo ID and liker ID from the path parameters
	posterID, err := strconv.ParseUint(ps.ByName("user_id"), 10, 64)
	if err != nil {
		// Handle the error, e.g., invalid photo ID
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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

	// User is unliking their own photo, save time by returning early
	if posterID == likerID {
		ctx.Logger.WithError(err).Error("you cannot unlike your own photo")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// likes1, err := rt.db.GetLikesList(photoID, posterID)
	// fmt.Println("likes 1", likes1)

	err = rt.db.UnlikePhoto(likerID, photoID, posterID)
	if errors.Is(err, database.ErrLikeDoesNotExist) {
		// The like (indicated by `id`) does not exist, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// Handle other errors.
		ctx.Logger.WithError(err).Error("error while unliking the photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Request an unfiltered list of likes from the DB
	likes, err := rt.db.GetLikesList(photoID, posterID)
	// fmt.Println("likes 2", likes)

	likesNumber := uint64(len(likes))

	err2 := rt.db.UpdateLikes(photoID, posterID, likesNumber)
	if err2 != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't update the like number")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to indicate JSON response
	w.Header().Set("Content-Type", "application/json")

	// Write the response with the updated like count
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]int{"likes": int(likesNumber)})
}
