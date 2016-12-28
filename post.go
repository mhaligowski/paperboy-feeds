package feeds

import (
	"net/http"
	"encoding/json"

	"github.com/SlyMarbo/rss"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"google.golang.org/appengine/datastore"
)

type handlePostJsonFeed struct {
	putter feedsPutter
}

func (h handlePostJsonFeed) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

type handlePostFormFeed struct {
	putter feedsPutter
	fetcher feedsFetcher
}

func (h handlePostFormFeed) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	url := r.FormValue("url")

	if url == "" {
		http.Error(w, "Invalid url", http.StatusBadRequest)
		return
	}

	// TODO sanitize url

	client := urlfetch.Client(ctx)
	feed, err := rss.FetchByClient(url, client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newFeed := &Feed{
		Title:feed.Title,
		Url:feed.UpdateURL,
	}

	feedId := setId(newFeed)
	_, err = h.fetcher.Get(ctx, feedId)
	if err == nil {
		w.WriteHeader(http.StatusConflict)
		return
	} else if err != datastore.ErrNoSuchEntity {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.putter.Put(ctx, newFeed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	// TODO: Add Location header

}