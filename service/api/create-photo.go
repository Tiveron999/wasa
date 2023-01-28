package api

import (
	"encoding/json"
	"github.com/Tiveron999/Wasa/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Read the new content for the photo from the request body.
	var photo Photo
	err := json.NewDecoder(r.Body).Decode(&photo)
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !photo.PhotoIsValid() {
		// Note: the PhotoIsValid() function skips the ID check (see below).
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Assume the image is already uploaded and the url is in photo.URL
	dbphoto, err := rt.db.UploadPhoto(photo.PhotoToDatabase())
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the photo
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't create the photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Send the output to the user.
	w.Header().Set("Content-Type", "image/png")
	http.ServeFile(w, r, dbphoto.URL)
}
