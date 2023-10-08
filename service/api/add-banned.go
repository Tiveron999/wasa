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
	_, err = rt.db.BanUser(banned.BannedToDatabase())
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the banned
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't create the banned user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Here we can re-use `banned` as FromDatabase is overwriting every variabile in the structure.
	// banned.BannedFromDatabase(dbbanned)

	// stop following user you ban (may not be already following)
	err = rt.db.RemoveFollower(BannerID, toBeBannedID)
	if errors.Is(err, database.ErrFollowerDoesNotExist) {
		// The follower (indicated by `(uid, fid)`) does not exist, reject the action indicating an error on the client side.
		s := fmt.Sprintf("%s + %s", strconv.FormatUint(BannerID, 10), strconv.FormatUint(toBeBannedID, 10)) // to create one string that keeps track of both IDs
		ctx.Logger.WithError(err).WithField("(uid, fid)", s).Error("banner wasn't following the banned user")
		return
	} else if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		// Note (2): we are adding the error and an additional field (`id`) to the log entry, so that we will receive
		// the identifier of the follower that triggered the error.

		s := fmt.Sprintf("%s + %s", strconv.FormatUint(BannerID, 10), strconv.FormatUint(toBeBannedID, 10)) // to create one string that keeps track of both IDs
		ctx.Logger.WithError(err).WithField("(uid, fid)", s).Error("can't delete the follower")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// banned user can't see the banner's posts
	err = rt.db.RemoveFollower(toBeBannedID, BannerID)
	if errors.Is(err, database.ErrFollowerDoesNotExist) {
		// The follower (indicated by `(uid, fid)`) does not exist, reject the action indicating an error on the client side.
		s := fmt.Sprintf("%s + %s", strconv.FormatUint(toBeBannedID, 10), strconv.FormatUint(BannerID, 10)) // to create one string that keeps track of both IDs
		ctx.Logger.WithError(err).WithField("(uid, fid)", s).Error("banned user wasn't following the banner")
		return
	} else if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		// Note (2): we are adding the error and an additional field (`id`) to the log entry, so that we will receive
		// the identifier of the follower that triggered the error.

		s := fmt.Sprintf("%s + %s", strconv.FormatUint(toBeBannedID, 10), strconv.FormatUint(BannerID, 10)) // to create one string that keeps track of both IDs
		ctx.Logger.WithError(err).WithField("(uid, fid)", s).Error("can't delete the follower")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 204 http status
	w.WriteHeader(http.StatusNoContent)
}
