package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Tiveron999/Wasa/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	toBeBannedID, err := strconv.ParseUint(ps.ByName(("banned_id")), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	BannerID, err := strconv.ParseUint(ps.ByName(("user_id")), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create banned object
	banned := Banned{
		UserID:   BannerID,
		BannedID: toBeBannedID,
	}

	// Create the banned in the database. Note that this function will return a new instance of the banned with the
	// same information, plus the ID.
	dbbanned, err := rt.db.BanUser(banned.BannedToDatabase())
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the banned
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't create the banned user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Here we can re-use `banned` as FromDatabase is overwriting every variabile in the structure.
	banned.BannedFromDatabase(dbbanned)

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(banned)
}
