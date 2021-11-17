package main

import (
	"net/http"

	"github.com/iliubang/grouter"
)

func index(w http.ResponseWriter, _ *http.Request, _ map[string]string) {
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(`{"code": 0, "data": "hello world"}`))
}

func aaa(w http.ResponseWriter, _ *http.Request, params map[string]string) {
	w.Write([]byte(params["name"]))
}

func main() {
	r := grouter.NewRouter()
	r.Get("/index.html", index)
	r.Get("/aaa/@name:([^/]+)", aaa)
	http.ListenAndServe("127.0.0.1:8080", r)
}
