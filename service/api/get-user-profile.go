package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Tiveron999/Wasa/service/api/reqcontext"
	"github.com/Tiveron999/Wasa/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse user ID from the URL
	userID, err := strconv.ParseUint(ps.ByName("user_id"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var profile Profile

	username, err := rt.db.GetUsername(userID)
	if err != nil {
		ctx.Logger.WithError(err).Error("an error occurred")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse query parameters
	queryValues := r.URL.Query()
	laterThan := queryValues.Get("laterthan")
	limit := queryValues.Get("limit")
	offset := queryValues.Get("offset")

	if laterThan == "" {
		laterThan = "2006-01-02 15:04:05" // Set a default value for laterThan
	}
	if limit == "" {
		limit = "10" // Set a default value for limit
	}
	if offset == "" {
		offset = "0" // Set a default value for offset
	}

	// Convert limit and offset to uints
	limitUint, err := strconv.ParseUint(limit, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	offsetUint, err := strconv.ParseUint(offset, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retrieve user photos from the database as BLOBs
	photos, err := rt.db.GetUserPhotos(userID, laterThan, limitUint, offsetUint)
	if err != nil {
		// Handle the error and return an appropriate response
		ctx.Logger.WithError(err).Error("can't retrieve photos")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var following []database.Follower
	const limitf = "10"

	limitfUint, err := strconv.ParseUint(limitf, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	const offsetf = "0"

	offsetfUint, err := strconv.ParseUint(offsetf, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Request an unfiltered list of following from the DB
	following, err = rt.db.GetFollowing(userID, limitfUint, offsetfUint)
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the photo
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't list followed users")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var followers []database.Follower

	const limitfs = "10"
	limitfsUint, err := strconv.ParseUint(limitfs, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	const offsetfs = "0"
	offsetfsUint, err := strconv.ParseUint(offsetfs, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Request an unfiltered list of followers from the DB
	followers, err = rt.db.GetFollowers(userID, limitfsUint, offsetfsUint)
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the photo
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't list followers")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create the profile object
	profile = Profile{
		Username:  username,
		NumPhotos: uint64(len(photos)),
		Followers: uint64(len(followers)),
		Following: uint64(len(following)),
	}

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(profile)
}
