package feeds

import (
	"golang.org/x/net/context"

	"google.golang.org/appengine/urlfetch"

	"github.com/SlyMarbo/rss"
)

type fetchRss func(ctx context.Context, url string) (*Feed, error)

func fetchRssWithUrlFetch(ctx context.Context, url string) (*Feed, error) {
	client := urlfetch.Client(ctx)
	feed, err := rss.FetchByClient(url, client)
	if err != nil {
		return nil, err
	}

	return &Feed{
		Title:feed.Title,
		Url:feed.UpdateURL,
	}, nil
}
