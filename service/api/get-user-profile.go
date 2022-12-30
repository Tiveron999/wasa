package api

import (
	"encoding/json"
	"net/http"

	"github.com/Tiveron999/Wasa/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Read the new content for the profile from the request body.
	var profile Profile
	err := json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !profile.ProfileIsValid() {
		// Here we validated the profile structure content (e.g., location coordinates in correct range, etc.), and we
		// discovered that the profile data are not valid.
		// Note: the IsValid() function skips the ID check (see below).
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create the profile in the database. Note that this function will return a new instance of the profile with the
	// same information, plus the ID.
	dbprofile, err := rt.db.GetUserProfile(profile.ProfileToDatabase())
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the profile
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't create the profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Here we can re-use `profile` as FromDatabase is overwriting every variabile in the structure.
	profile.ProfileFromDatabase(dbprofile)

	// Send the output to the profile.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(profile)
}
