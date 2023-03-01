package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Tiveron999/Wasa/service/api/reqcontext"
	"github.com/Tiveron999/Wasa/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// The user ID in the path is a 64-bit unsigned integer. Let's parse it.
	uid, err := strconv.ParseUint(ps.ByName(("user_id")), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// The followed user ID in the path is a 64-bit unsigned integer. Let's parse it.
	fid, err := strconv.ParseUint(ps.ByName(("follower_id")), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rt.db.UnfollowUser(uid, fid)
	if errors.Is(err, database.ErrFollowerDoesNotExist) {
		// The followed user (indicated by `(uid, fid)`) does not exist, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		// Note (2): we are adding the error and an additional field (`id`) to the log entry, so that we will receive
		// the identifier of the followed user that triggered the error.

		s := fmt.Sprintf("%s + %s", strconv.FormatUint(fid, 10), strconv.FormatUint(uid, 10)) // to create one string that keeps track of both IDs
		ctx.Logger.WithError(err).WithField("(fid, uid)", s).Error("can't delete the followed user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
