package api

import (
	"net/http"
	"strconv"

	"github.com/Tiveron999/Wasa/service/api/reqcontext"
	// "github.com/Tiveron999/Wasa/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserPhotos(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse user ID from the URL
	userID, err := strconv.ParseUint(ps.ByName("user_id"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Parse query parameters
	queryValues := r.URL.Query()
	laterThan := queryValues.Get("laterthan")
	limit := queryValues.Get("limit")
	offset := queryValues.Get("offset")

	if laterThan == "" {
		laterThan = "2006-01-02 15:04:05" // Set a default value for laterThan
	}
	if limit == "" {
		limit = "10" // Set a default value for limit
	}
	if offset == "" {
		offset = "0" // Set a default value for offset
	}

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

	// Retrieve user photos from the database as BLOBs
	photoBlobs, err := rt.db.GetUserPhotos(userID, laterThan, limitUint, offsetUint)
	if err != nil {
		// Handle the error and return an appropriate response
		ctx.Logger.WithError(err).Error("can't retrieve photos")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")

	// Loop through the BLOBs and send them as binary data in the response
	for _, photoBlob := range photoBlobs {
		_, err := w.Write([]byte{photoBlob})
		if err != nil {
			// Handle any write errors, log them, and return an appropriate response
			ctx.Logger.WithError(err).Error("error writing photo data")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

/*
	// Parse the query string part. To do that, we need to check whether the latitude, longitude and range exists.
	// If latitude and longitude are specified, we parse them, and we filter results for them. If range is specified,
	// the value will be parsed and used as a filter. If it's not specified, 10 will be used as default (as specified in
	// the OpenAPI file).
	// If one of latitude or longitude is not specified (or both), no filter will be applied.

	var err error
	var photoURLs []string

	// page limit sort, has
	userID, err := strconv.ParseUint(ps.ByName("user_id"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	queryValues := r.URL.Query()
	laterThan := queryValues.Get("laterthan")
	limit := queryValues.Get("limit")
	offset := queryValues.Get("offset")

	if laterThan == "" {
		laterThan = "2006-01-02 15:04:05" // Set a default value for laterThan
	}
	if limit == "" {
		limit = "10" // Set a default value for limit
	}
	if offset == "" {
		offset = "0" // Set a default value for offset
	}

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

	// Request an unfiltered list of photo Blobs from the DB
	photoURLs, err = rt.db.GetUserPhotos(userID, laterThan, limitUint, offsetUint)

	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the photo
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't list photos")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the list of photo Blobss to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(photoURLs)
}
*/
