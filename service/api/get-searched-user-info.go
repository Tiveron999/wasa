package api

import (
	"encoding/json"
	// "fmt"
	"net/http"
	"strconv"

	"github.com/Tiveron999/Wasa/service/api/reqcontext"
	"github.com/Tiveron999/Wasa/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getSearchedUserInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the value of the "username" query parameter
	username := r.URL.Query().Get("username")

	// fmt.Fprintf(w, "Hello, %s!", username)

	// fmt.Printf("please")

	/*profile, err := rt.db.GetSearchedUserInfo(username)
		if err != nil {
			ctx.Logger.WithError(err).Error("can't list profile")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(profile)
	}*/

	var profile Profile

	userID, err := rt.db.GetUserID(username)
	if err != nil {
		ctx.Logger.WithError(err).Error("an error occurred")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var photoURLs []string
	const limit = "10"
	limitUint, err := strconv.ParseUint(limit, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	const offset = "0"
	offsetUint, err := strconv.ParseUint(offset, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Request an unfiltered list of photo URLs from the DB
	photoURLs, err = rt.db.GetUserPhotos(userID, "2006-01-02 15:04:05", limitUint, offsetUint)
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the photo
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't list photos")
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
	// fmt.Println(userID, limitfUint, offsetfUint)
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
		Photos:    photoURLs,
		NumPhotos: uint64(len(photoURLs)),
		Followers: uint64(len(followers)),
		Following: uint64(len(following)),
	}

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(profile)
}
