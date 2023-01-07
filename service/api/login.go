package api

import (
	"fmt"
	"github.com/Tiveron999/Wasa/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) login(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse the username from the request
	username := r.FormValue("username")

	// Get or create the user in the database
	userID, err := rt.db.CheckUserExists(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the user ID to the client
	fmt.Fprint(w, userID)
}
