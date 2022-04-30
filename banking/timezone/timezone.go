package timezone

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	//router.HandleFunc("/api/time/{tz:[a-z]+}", getTime).Methods(http.MethodGet)
	router.HandleFunc("/api/time", getTime).Methods(http.MethodGet)

	// Starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
