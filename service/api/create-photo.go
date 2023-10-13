package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/Tiveron999/Wasa/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the user's ID from the URL
	uploaderID, err := strconv.ParseUint(ps.ByName("user_id"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Parse the multipart form data from the request
	err = r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the uploaded file from the form data
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Read the image data into a []byte
	imageData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	/*// Save the image blob to disk
	if err := ioutil.WriteFile("images.png", imageData, os.ModePerm); err != nil {
		http.Error(w, "Failed to save image", http.StatusInternalServerError)
		return
	}*/

	// Create the photo object to pass to the db
	photo := Photo{
		PID:         0,
		Image:       imageData,
		PosterID:    uploaderID,
		Datetime:    time.Now().Format("2006-01-02 15:04:05"),
		NumLikes:    0,
		NumComments: 0,
	}

	// Save the file data to the database and get the photo ID
	photoID, err := rt.db.UploadPhoto(photo.PhotoToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the photo ID and return with a 201 status code
	photo.PID = photoID
	w.Header().Set("Content-Type", "image/jpeg")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(photo)
}
