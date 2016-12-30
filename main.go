package feeds

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"google.golang.org/appengine"
)

func Run() {
	dao := datastoreFeedsDao{}
	router := mux.NewRouter()

	router.Handle("/feeds", cors.Default().Handler(getFeedsHandler{dao})).
		Methods(http.MethodGet)

	getHandler := cors.Default().Handler(getFeedHandler{dao})
	router.Handle("/feeds/{feedId}", getHandler).
		Methods(http.MethodGet)

	postHandler := cors.Default().Handler(postHttpHandler{appengine.NewContext, fetchRssWithUrlFetch, dao})
	router.Handle("/feeds", postHandler).
		Methods(http.MethodPost).
		Headers("Content-Type", "application/json")

	http.Handle("/", cors.Default().Handler(router))
}