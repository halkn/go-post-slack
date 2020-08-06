package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newMockHandler(status int) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		fmt.Fprintln(w, "called mock handler")
	}
}

func TestSlackClient_PostRequest(t *testing.T) {
	tests := map[string]struct {
		mockFunc func(http.ResponseWriter, *http.Request)
	}{
		"OK Test": {
			mockFunc: newMockHandler(200),
		},
	}
	sc := &SlackClient{}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(tt.mockFunc))
			defer ts.Close()
			err := sc.PostRequest(ts.URL, "")
			if err != nil {
				t.Errorf("err: \n%v", err)
			}
		})
	}

}
