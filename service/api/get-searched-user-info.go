package api

import (
	"encoding/json"
	// "fmt"
	"net/http"

	"github.com/Tiveron999/Wasa/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getSearchedUserInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the value of the "username" query parameter
	username := r.URL.Query().Get("username")

	// fmt.Fprintf(w, "Hello, %s!", username)

	// fmt.Printf("please")

	profile, err := rt.db.GetSearchedUserInfo(username)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't list profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(profile)
}
