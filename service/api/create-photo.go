package api

import (
	"encoding/json"
	"net/http"

	"github.com/Tiveron999/Wasa/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) createPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Read the new content for the photo from the request body.
	var photo Photo
	err := json.NewDecoder(r.Body).Decode(&photo)
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !photo.PhotoIsValid() {
		// Here we validated the photo structure content (e.g., location coordinates in correct range, etc.), and we
		// discovered that the photo data are not valid.
		// Note: the IsValid() function skips the ID check (see below).
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create the photo in the database. Note that this function will return a new instance of the photo with the
	// same information, plus the ID.
	dbphoto, err := rt.db.CreatePhoto(photo.PhotoToDatabase())
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the photo
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't create the photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Here we can re-use `photo` as FromDatabase is overwriting every variabile in the structure.
	photo.PhotoFromDatabase(dbphoto)

	// Send the output to the photo.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(photo)
}
