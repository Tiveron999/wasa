package api

import (
	"encoding/json"
	// "fmt"
	"net/http"

	"github.com/Tiveron999/Wasa/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) login(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse the username from the json object
	type loginData struct {
		Username string `json:"username"`
	}

	// fmt.Printf("please")
	var t loginData
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get or create the user in the database
	userID, err := rt.db.CheckUserExists(t.Username)
	if err != nil {
		ctx.Logger.WithError(err).Error("an error occurred")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// _, _ = w.Write([]byte("Success!"))
	// Return the user ID to the user.
	w.Header().Set("Content-Type", "application/json")
	// w.writeHeader
	_ = json.NewEncoder(w).Encode(map[string]int{"userID": int(userID)})
}
