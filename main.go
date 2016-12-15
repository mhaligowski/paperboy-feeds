package feeds

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Run() {
	dao := datastoreFeedsDao{}
	router := mux.NewRouter()

	router.Handle("/feeds", getFeedsHandler{dao}).
		Methods(http.MethodGet)

	router.Handle("/feeds/{feedId}", getFeedHandler{dao}).Methods(http.MethodGet)
	router.Handle("/feeds", handlePostFeed{dao}).Methods(http.MethodPost)

	http.Handle("/", router)
}