package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Tiveron999/Wasa/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the user's ID from the URL
	uploader_id, er := strconv.ParseUint(ps.ByName(("user_id")), 10, 64)
	if er != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Parse the multipart form data from the request
	err := r.ParseMultipartForm(10 << 20) // max memory 10MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the uploaded file from the form data
	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Check if the "webui/public/images" directory exists
	if _, err := os.Stat("./webui/src/assets/images"); os.IsNotExist(err) {
		// The "webui/public/images" directory does not exist, create it
		err := os.MkdirAll("./webui/src/assets/images", os.ModePerm)
		if err != nil {
			// Handle error when creating the directory
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Create a new file in the "webui/public/images" directory with the same name as the uploaded file
	out, err := os.Create("./webui/src/assets/images/" + header.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer out.Close()

	// Copy the file data to the new file
	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create the photo object to pass to the db
	photo := Photo{
		URL:         "/webui/src/assets/images/" + header.Filename,
		Posterid:    uploader_id,
		Datetime:    time.Now().Format("2006-01-02 15:04:05"),
		Numlikes:    0,
		Numcomments: 0,
	}

	// Save the file data to the database and get the photo ID
	photoId, err := rt.db.UploadPhoto(photo.PhotoToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the photo ID in the response
	response := map[string]interface{}{
		"photoId": photoId,
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

// Read the new content for the photo from the request body.
/*userId := r.Header.Get("userId")
imageType := r.Header.Get("imageType")
photoId := ps.ByName("photoId")

url := fmt.Sprintf("/images/users/%s/%s.%s", userId, photoId, imageType)
w.Header().Set("Content-Type", "image/"+imageType)
http.ServeFile(w, r, url)*/

/*var photo Photo
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
}*/
/*
	// Send the output to the user.
	w.Header().Set("Content-Type", "image/png")
	http.ServeFile(w, r, dbphoto.URL)*/
// Create a JSON object with the photo ID
/*response := map[string]interface{}{
		"photoId": dbphoto.ID,
	}

	// Encode the JSON object and write it to the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}*/

/*
	const fileInput = document.getElementById("photo");

	if (fileInput.files.length > 0) {
		const file = fileInput.files[0];
		const reader = new FileReader();

		reader.onloadend = function() {
		// Make a request to the server to save the photo URL to the database
		const photoUrl = reader.result;
		const photoId = savePhotoToDatabase(photoUrl);

		console.log("Photo ID: " + photoId);
		}

		reader.readAsDataURL(file);
	}
	}
*/
