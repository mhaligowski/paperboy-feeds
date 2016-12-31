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

	corsHandler := cors.New(cors.Options{
		AllowedHeaders:{"Location"},
	})
	router.Handle("/feeds", corsHandler(getFeedsHandler{dao})).
		Methods(http.MethodGet)

	getHandler := corsHandler(getFeedHandler{dao})
	router.Handle("/feeds/{feedId}", getHandler).
		Methods(http.MethodGet)

	postHandler := corsHandler(postHttpHandler{appengine.NewContext, fetchRssWithUrlFetch, dao})
	router.Handle("/feeds", postHandler).
		Methods(http.MethodPost).
		Headers("Content-Type", "application/json")

	http.Handle("/", corsHandler(router))
}