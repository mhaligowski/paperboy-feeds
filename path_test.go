package feeds

import (
	"testing"
	"net/url"
)

var tests = []struct {
	url string
	id string
	expected string
}{
	{"http://example.com/feeds/", "123", "http://example.com/feeds/123"},
	{"http://example.com/feeds", "123", "http://example.com/feeds/123"},
	{"http://example.com/feeds", "otherId", "http://example.com/feeds/otherId"},
	{"http://otherexample.com/feeds", "otherId", "http://otherexample.com/feeds/otherId"},
	{"http://localhost/feeds", "otherId", "http://localhost/feeds/otherId"},
	{"http://localhost:8080/feeds", "otherId", "http://localhost:8080/feeds/otherId"},
	{"https://localhost:8080/feeds", "otherId", "https://localhost:8080/feeds/otherId"},
	{"https://localhost:8080/somePrefix/feeds", "otherId", "https://localhost:8080/feeds/otherId"},
}

func TestInputs(t *testing.T) {
	for _, tt := range tests {
		u, _ := url.Parse(tt.url)
		actual := getFeedPath(u, tt.id)
		if actual != tt.expected {
			t.Errorf("Got %v, expected %v for url: %v, id: %v",
			actual, tt.expected, tt.url, tt.id)
		}
	}
}
