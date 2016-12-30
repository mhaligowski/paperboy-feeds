package feeds

import (
	"net/http"
	"encoding/json"
)

type postHttpHandler struct {
	getContext getContext
	fetchRss   fetchRss

	feedsDao   feedsDao
}

func (h postHttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f := &Feed{}
	err := json.NewDecoder(r.Body).Decode(f)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO sanitize url
	ctx := h.getContext(r)
	newFeed, err := h.fetchRss(ctx, f.Url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	feedId := setId(newFeed)
	oldFeed, err := h.feedsDao.Get(ctx, feedId)
	if oldFeed != nil && err == nil {
		w.Header().Add("Location", getFeedPath(r.URL, oldFeed.FeedId))
		w.WriteHeader(http.StatusSeeOther)
		return
	} else if newFeed == nil && err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.feedsDao.Put(ctx, newFeed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Location", getFeedPath(r.URL, newFeed.FeedId))
	w.WriteHeader(http.StatusSeeOther)
}
