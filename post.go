package feeds

import (
	"net/http"
	"encoding/json"
	"google.golang.org/appengine"
)

type handlePostFeed struct {
	putter feedsPutter
}

func (h handlePostFeed) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f := &Feed{}
	err := json.NewDecoder(r.Body).Decode(f)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := appengine.NewContext(r)
	err = h.putter.Put(ctx, f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	// TODO: Add Location header
}