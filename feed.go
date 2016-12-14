package feeds

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"crypto/sha256"
	"encoding/hex"
)

type Feed struct {
	FeedId string
	Title  string
	Url    string
}

type feedsFetcher interface {
	Get(context.Context, string) (*Feed, error)
	GetAll(context.Context) ([]Feed, error)
}

type feedsPutter interface {
	Put(context.Context, *Feed) error
}

type datastoreFeedsDao struct {
}

func (d datastoreFeedsDao) Get(ctx context.Context, id string) (*Feed, error) {
	result := &Feed{}
	key := datastore.NewKey(ctx, "Feed", id, 0, nil)

	err := datastore.Get(ctx, key, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (d datastoreFeedsDao) GetAll(ctx context.Context) ([]Feed, error) {
	var result []Feed

	q := datastore.NewQuery("Feed")
	if _, err := q.GetAll(ctx, &result); err != nil {
		return nil, err
	}

	if result == nil {
		return make([]Feed, 0), nil
	} else {
		return result, nil
	}
}

func (d datastoreFeedsDao) Put(ctx context.Context, f *Feed) error {
	f.FeedId = id(f)
	k := datastore.NewKey(ctx, "Feed", f.FeedId, 0, nil)

	_, err := datastore.Put(ctx, k, f)
	return err
}

func id(f *Feed) string {
	b := sha256.New()
	b.Write([]byte(f.Url))

	return hex.EncodeToString(b.Sum(nil))
}
