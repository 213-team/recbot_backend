package backend

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//GetChannelsHandler returns a list of channels followed by a user
func GetChannelsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]

	message := Message{
		Message: fmt.Sprintf("User %s doesn't have followed channels yet", userID),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(message); err != nil {
		panic(err)
	}
}
