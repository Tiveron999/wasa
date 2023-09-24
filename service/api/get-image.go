package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/Tiveron999/Wasa/service/api/reqcontext"

	// "github.com/Tiveron999/Wasa/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getImageFromPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse photo ID from the URL
	imageID, err := strconv.ParseUint(ps.ByName("image_id"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var image []byte
	// Retrieve user photos from the database as BLOBs
	image, err = rt.db.GetUserImage(imageID)
	if err != nil {
		// Handle the error and return an appropriate response
		ctx.Logger.WithError(err).Error("can't retrieve photos")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Save the image blob to disk
	if err := ioutil.WriteFile("images.jpeg", image, os.ModePerm); err != nil {
		http.Error(w, "Failed to save image", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(image)

}
