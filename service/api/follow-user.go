package api

import (
	"net/http"
	"strconv"

	"github.com/Tiveron999/Wasa/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the user's ID from the URL
	toBeFollowedID, err := strconv.ParseUint(ps.ByName(("following_id")), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the Authorization header value
	authorizationHeader := r.Header.Get("Authorization")

	// Get the user ID from the Authorization header
	toBeFollowerID, isValid := idFromAuthorizationHeader(authorizationHeader)
	if !isValid {
		// If the Authorization header is invalid or not present, handle the error
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Any user is not allowed to follow themself
	if toBeFollowedID == toBeFollowerID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create the follower object
	following := Follower{
		UserID:     toBeFollowedID,
		FollowerID: toBeFollowerID,
	}

	// Create the following in the database.
	err1 := rt.db.FollowUser(following.FollowerToDatabase())
	if err1 != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't create the following")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 204 http status
	w.WriteHeader(http.StatusNoContent)
}
