package main

import (
	"log"
	"net/http"
	"sync/atomic"

	"github.com/iliubang/grouter"
)

var count uint64

func init() {
	atomic.StoreUint64(&count, 0)
}

func home(w http.ResponseWriter, r *http.Request, pathvariables map[string]string) {
	atomic.AddUint64(&count, 1)
	c := atomic.LoadUint64(&count)
	log.Printf("count: %d\n", c)
	if (c & 1) == 1 {
		panic("hahaha")
	} else {
		w.Write([]byte("hello world"))
	}
}

func main() {
	router := grouter.NewRouter()
	router.Get("/index.html", home)
	httpServer := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}
	httpServer.ListenAndServe()
}
