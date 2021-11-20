package main

import (
	"context"
	"net/http"

	"github.com/iliubang/grouter"
)

func foo(_ context.Context, w http.ResponseWriter, _ *http.Request, _ map[string]string) {
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(`{"code": 0, "data": "hello world"}`))
}

func bar(_ context.Context, w http.ResponseWriter, _ *http.Request, params map[string]string) {
	w.Write([]byte(params["name"]))
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	r := grouter.NewRouter(ctx)
	r.Get("/index.html", foo)
	r.Get("/aaa/@name:([^/]+)", bar)
	http.ListenAndServe("127.0.0.1:8080", r)
}
