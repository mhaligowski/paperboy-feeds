package feeds

import (
	"net/http"
	"fmt"
)

func getFeeds(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}