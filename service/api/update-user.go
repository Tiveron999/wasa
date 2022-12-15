package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Tiveron999/Wasa/service/api/reqcontext"
	"github.com/Tiveron999/Wasa/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) updateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// The User ID in the path is a 64-bit unsigned integer. Let's parse it.
	id, err := strconv.ParseUint(ps.ByName("id"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Read the new content for the user from the request body.
	var updatedUser User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !updatedUser.IsValid() {
		// Here we validated the user structure content (e.g., location coordinates in correct range, etc.), and we
		// discovered that the user data are not valid.
		// Note: the IsValid() function skips the ID check (see below).
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// The client is not supposed to send us the ID in the body, as the user ID is already specified in the path,
	// and it's immutable. So, here we overwrite the ID in the JSON with the `id` variable (that comes from the URL).
	updatedUser.ID = id

	// Update the user in the database.
	err = rt.db.UpdateUser(updatedUser.ToDatabase())
	if errors.Is(err, database.ErrUserDoesNotExist) {
		// The user (indicated by `id`) does not exist, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		// Note (2): we are adding the error and an additional field (`id`) to the log entry, so that we will receive
		// the identifier of the user that triggered the error.
		ctx.Logger.WithError(err).WithField("id", id).Error("can't update the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
