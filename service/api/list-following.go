package api

import (
	"encoding/json"
	"net/http"

	"github.com/Tiveron999/Wasa/service/api/reqcontext"
	"github.com/Tiveron999/Wasa/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getFollowing(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Parse the query string part. To do that, we need to check whether the latitude, longitude and range exists.
	// If latitude and longitude are specified, we parse them, and we filter results for them. If range is specified,
	// the value will be parsed and used as a filter. If it's not specified, 10 will be used as default (as specified in
	// the OpenAPI file).
	// If one of latitude or longitude is not specified (or both), no filter will be applied.

	var err error
	var followers []database.Follower

	// Request an unfiltered list of followers from the DB
	followers, err = rt.db.GetFollowing()

	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the photo
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't list followed users")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Before sending the list to the client we need to convert it to the "frontend" type
	var frontendFollowing = make([]Follower, len(followers))
	for idx := range followers {
		frontendFollowing[idx].FollowerFromDatabase(followers[idx])
	}

	// Send the list to the photo.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(frontendFollowing)
}
