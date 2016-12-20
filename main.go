package feeds

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/rs/cors"
)

func Run() {
	dao := datastoreFeedsDao{}
	router := mux.NewRouter()

	router.Handle("/feeds", getFeedsHandler{dao}).
		Methods(http.MethodGet)

	router.Handle("/feeds/{feedId}", getFeedHandler{dao}).Methods(http.MethodGet)
	router.Handle("/feeds", handlePostJsonFeed{dao}).
		Methods(http.MethodPost).
		Headers("Content-Type", "application/json")
	router.Handle("/feeds", handlePostFormFeed{dao, dao}).
		Methods(http.MethodPost)

	http.Handle("/", cors.Default().Handler(router))
}