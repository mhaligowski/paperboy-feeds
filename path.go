package feeds

import (
	"net/url"
)

var (
	parentRef, _ = url.Parse("/")
	feedsRef, _ = url.Parse("feeds/")
)

func getFeedPath(u *url.URL, id string) string {

	root := u.ResolveReference(parentRef)

	idRef, _ := url.Parse(id)
	return root.ResolveReference(feedsRef).
		ResolveReference(idRef).
		String()
}

