package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"sync"

	"testing"
)

func TestServerRace(t *testing.T) {
	ts := &server{
		mtx: &sync.Mutex{},
	}
	handler := ts.handleCount()
	getCounter := func(out chan []byte) {
		w := httptest.NewRecorder()
		handler(w, &http.Request{})
		resp := w.Result()
		body, _ := io.ReadAll(resp.Body)
		out <- body
	}
	out := make(chan []byte)
	go getCounter(out)
	go getCounter(out)
	go getCounter(out)
	go getCounter(out)
	go getCounter(out)
	<-out
	<-out
	<-out
	<-out
	<-out
	want := 5
	if ts.counter != want {
		t.Errorf("Got: %d - Want: %d", ts.counter, want)
	}
}
