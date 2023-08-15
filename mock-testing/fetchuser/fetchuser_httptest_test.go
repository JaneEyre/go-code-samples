package fetchuser

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProcessUser_HttpTest(t *testing.T) {
	user := User{ID: 1, Name: "Alice"}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userJSON, _ := json.Marshal(user)
		n, err := w.Write(userJSON)
		if err != nil {
			t.Errorf("test server: unexpected error after writing %d bytes: %v", n, err)
		}
	}))
	defer ts.Close()

	fetcher := &RealAPIFetcher{
		ApiURL: ts.URL, // mock URL provided by httptest
	}
	http.DefaultClient = ts.Client() // client provided by httptest

	result, err := ProcessUser(fetcher, 1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if result != user {
		t.Errorf("expected user: %v, got: %v", user, result)
	}
}
