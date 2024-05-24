package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

type server struct {
	router  *http.ServeMux
	server  *http.Server
	counter int
	mtx     *sync.Mutex
}

func main() {
	addr := ":8080"
	s := &server{
		router: http.NewServeMux(),
		server: &http.Server{
			Addr:           addr,
			ReadTimeout:    2 * time.Second,
			WriteTimeout:   2 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
		mtx: &sync.Mutex{},
	}
	// Definition der Handler
	s.router.HandleFunc("/", s.handleCount())
	s.server.Handler = s.router
	fmt.Printf("Server started at %s", s.server.Addr)
	log.Fatal(s.server.ListenAndServe())
}

func (s *server) handleCount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.mtx.Lock()
		s.counter++
		io.WriteString(w, fmt.Sprintf("Counter: %03d", s.counter))
		s.mtx.Unlock()
	}
}
