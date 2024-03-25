package gocontext

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// func (s *SpyStore) assertWasCancelled() {
// 	s.t.Helper()
// 	if !s.cancelled {
// 		s.t.Error("store was not told to cancel")
// 	}
// }

// func (s *SpyStore) assertWasNotCancelled() {
// 	s.t.Helper()
// 	if s.cancelled {
// 		s.t.Error("store was told to cancel")
// 	}
// }

func TestServer(t *testing.T) {

	t.Run("First test case with httptest", func(t *testing.T) {
		data := "Hello, World!"
		svr := Server(&SpyStore{
			response: data,
		})

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)
		if response.written {
			t.Error("a response should not have written")
		}
	})

	t.Run("Tells store to cancel work if request is cancelled", func(t *testing.T) {

		data := "Hello, World!"
		store := &SpyStore{
			response: data,
			t:        t,
		}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		ctx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(ctx)

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should not have written")
		}

	})
}
