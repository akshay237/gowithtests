package goselect

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	fastUrl := "http://www.facebook.com"
	slowUrl := "http://www.quii.dev"

	want := fastUrl
	got := Racer(slowUrl, fastUrl)
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
	t.Log("Fast url is: ", got)

}

func makeHTTPServer(delay time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
	return server
}

func TestRacerWithMockServer(t *testing.T) {

	t.Run("returns ther url which responds faster in 10 secs", func(t *testing.T) {
		slowServer := makeHTTPServer(20 * time.Second)
		defer slowServer.Close()

		fastServer := makeHTTPServer(0 * time.Second)
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL
		t.Log("fast Url: ", fastURL)
		t.Log("slow url: ", slowURL)

		want := fastURL
		got, err := RacerWithSelect(slowURL, fastURL)

		if err != nil {
			t.Errorf("expected nilbut got error %q ", err)
		}

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("return an error when the server doesn't responds within 10 secs", func(t *testing.T) {
		slowServer := makeHTTPServer(11 * time.Second)
		defer slowServer.Close()

		fastServer := makeHTTPServer(12 * time.Second)
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL
		t.Log("fast Url: ", fastURL)
		t.Log("slow url: ", slowURL)

		want := fastURL
		got, err := RacerWithSelect(slowURL, fastURL)

		if err != nil {
			t.Fatalf("expected nil but got error %q ", err)
		}

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

}
