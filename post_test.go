package feeds

import (
	"net/http"
	"testing"
	"strings"
	"net/http/httptest"
	"golang.org/x/net/context"
)

type mockDao struct {
	f *Feed
}

func newMockFetcher(f *Feed) *mockDao {
	return &mockDao{f}
}

func (m *mockDao) Get(ctx context.Context, id string) (*Feed, error) {
	return m.f, nil
}

func (m *mockDao) GetAll(ctx context.Context) ([]Feed, error) {
	return nil, nil
}

func (m *mockDao) Put(ctx context.Context, f *Feed) error {
	f.FeedId = "123"
	return nil
}

func newContext(*http.Request) context.Context {
	return context.Background()
}

func mockFetch(context.Context, string) (*Feed, error) {
	return &Feed{
		Url: "http://wiadomosci.wp.pl/kat,1329,ver,rss,rss.xml",
		Title: "WP",
	}, nil
}

func TestReturns303WhenFeedIsNotCreated(t *testing.T) {
	body := `{
		"Url": "http://wiadomosci.wp.pl/kat,1329,ver,rss,rss.xml"
	}`

	someFeed := &Feed{FeedId: "123"}
	dao := newMockFetcher(someFeed)
	req, err := http.NewRequest("POST", "http://localhost:8080/feeds", strings.NewReader(body))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := postHttpHandler{newContext, mockFetch, dao}

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status, got %v want %v, message: %v",
			status, http.StatusSeeOther, rr.Body.String())
	}

	expectedLocation := "http://localhost:8080/feeds/123"
	if location := rr.Header().Get("Location");
		location != expectedLocation {
		t.Errorf("handler returned wrong location header, got %v want %v",
			location, expectedLocation)
	}

}

func TestReturns303WhenFeedIsCreated(t *testing.T) {
	body := `{
		"Url": "http://wiadomosci.wp.pl/kat,1329,ver,rss,rss.xml"
	}`

	dao := newMockFetcher(nil)
	req, err := http.NewRequest("POST", "http://localhost:8080/feeds", strings.NewReader(body))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := postHttpHandler{newContext, mockFetch, dao}

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status, got %v want %v, message: %v",
			status, http.StatusSeeOther, rr.Body.String())
	}

	expectedLocation := "http://localhost:8080/feeds/123"
	if location := rr.Header().Get("Location");
		location != expectedLocation {
		t.Errorf("handler returned wrong location header, got %v want %v",
			location, expectedLocation)
	}
}

func TestReturns400WhenInvalidInput(t *testing.T) {
	body := `{
		invalid: "http://wiadomosci.wp.pl/kat,1329,ver,rss,rss.xml"
	}`

	someFeed := &Feed{FeedId: "123"}
	dao := newMockFetcher(someFeed)
	req, err := http.NewRequest("POST", "http://localhost:8080/feeds", strings.NewReader(body))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := postHttpHandler{nil, nil, dao}

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status, got %v want %v, message: %v",
			status, http.StatusSeeOther, rr.Body.String())
	}

}