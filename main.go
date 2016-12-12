package feeds

import (
	"github.com/gorilla/mux"
	"net/http"
)

func init() {
	router := mux.NewRouter()

	router.HandleFunc("/feeds", getFeeds).
		Methods(http.MethodGet)

	http.Handle("/", router)
}