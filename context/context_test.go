package context

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
	t         testing.T
}

func (s *SpyStore) assertWasCancelled() {
	s.t.Helper()

	if !s.cancelled {
		s.t.Errorf("store was not told to cancel")
	}
}

func (s *SpyStore) assertWasNotCancelled() {
	s.t.Helper()

	if s.cancelled {
		s.t.Errorf("store was told to cancel")
	}
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result strings.Builder

		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(5 * time.Millisecond)
				result.WriteString(string(c))
			}
		}

		data <- result.String()
	}()

	select {
	case d := <-data:
		return d, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func TestServer(t *testing.T) {
	data := "hello, world"
	t.Run("", func(t *testing.T) {

		svr := Server(&SpyStore{response: data})

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {

		store := &SpyStore{response: data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancelCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancelCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		store.assertWasNotCancelled()
	})

	t.Run("returns data from the store", func(t *testing.T) {
		stored := &SpyStore{response: data}
		svr := Server(stored)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}

		stored.assertWasCancelled()
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		stored := &SpyStore{response: data}
		svr := Server(stored)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should not have been written")
		}
	})
}
