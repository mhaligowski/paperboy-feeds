package feeds

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Run() {
	dao := datastoreFeedsDao{}
	router := mux.NewRouter()

	router.Handle("/feeds", cors.Default().Handler(getFeedsHandler{dao})).
		Methods(http.MethodGet)

	router.Handle("/feeds/{feedId}", cors.Default().Handler(getFeedHandler{dao})).
		Methods(http.MethodGet)

	router.Handle("/feeds", cors.Default().Handler(handlePostJsonFeed{dao})).
		Methods(http.MethodPost).
		Headers("Content-Type", "application/json")
	router.Handle("/feeds", cors.Default().Handler(handlePostFormFeed{dao, dao})).
		Methods(http.MethodPost)

	http.Handle("/", cors.Default().Handler(router))
}