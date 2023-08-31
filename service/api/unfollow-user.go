package api

import (
	"net/http"
	"strconv"

	"github.com/Tiveron999/Wasa/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// ID of the person who we want to stop following
	toBeDeletedFollowedID, err := strconv.ParseUint(ps.ByName(("following_id")), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the Authorization header
	authorizationHeader := r.Header.Get("Authorization")

	// Get the user ID from the Authorization header
	clickerID, isValid := idFromAuthorizationHeader(authorizationHeader)
	if !isValid {
		// If the Authorization header is invalid or not present, handle the error
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Users cannot unfollow themselves, return no content
	if clickerID == toBeDeletedFollowedID {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Create the follower object
	deleted := Follower{
		UserID:     toBeDeletedFollowedID,
		FollowerID: clickerID,
	}

	// Remove the follower in the db via db function
	err = rt.db.UnfollowUser(deleted.FollowerToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("Error unfollowing user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
