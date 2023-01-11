package api

import (
	"encoding/json"
	"github.com/Tiveron999/Wasa/service/api/reqcontext"
	"github.com/Tiveron999/Wasa/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) getFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Parse the query string part. To do that, we need to check whether the latitude, longitude and range exists.
	// If latitude and longitude are specified, we parse them, and we filter results for them. If range is specified,
	// the value will be parsed and used as a filter. If it's not specified, 10 will be used as default (as specified in
	// the OpenAPI file).
	// If one of latitude or longitude is not specified (or both), no filter will be applied.

	var err error
	var followers []database.Follower

	// Extract the user ID from the path
	userID, err := strconv.ParseUint(ps.ByName("user_id"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	queryValues := r.URL.Query()
	limit := queryValues.Get("limit")
	offset := queryValues.Get("offset")

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

	// Request an unfiltered list of followers from the DB
	followers, err = rt.db.GetFollowers(userID, limitUint, offsetUint)

	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the photo
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't list followers")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Before sending the list to the client we need to convert it to the "frontend" type
	var frontendFollower = make([]Follower, len(followers))
	for idx := range followers {
		frontendFollower[idx].FollowerFromDatabase(followers[idx])
	}

	// Send the list to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(frontendFollower)
}
