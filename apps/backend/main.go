package main

import (
	"log"
	"net/http"

	"github.com/213-team/recbot_backend/backend"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/{user_id}/channels", backend.GetChannelsHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
