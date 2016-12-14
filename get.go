package feeds

import (
	"net/http"
	"google.golang.org/appengine"
	"encoding/json"
	"github.com/gorilla/mux"
	"google.golang.org/appengine/datastore"
)

type getFeedsHandler struct {
	fetcher feedsFetcher
}

func (h getFeedsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	feeds, err := h.fetcher.GetAll(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(feeds)
}

type getFeedHandler struct {
	fetcher feedsFetcher
}

func (h getFeedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ctx := appengine.NewContext(r)

	result, err := h.fetcher.Get(ctx, vars["feedId"])

	switch {
	case err == datastore.ErrNoSuchEntity: {
		http.Error(w, "value not found", http.StatusNotFound)
		return
	}
	case err != nil: {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	default: {
		json.NewEncoder(w).Encode(result)
	}
	}
}