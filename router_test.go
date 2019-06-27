package Grouter

import "testing"

func TestRouter_Lookup(t *testing.T) {
	router := NewRouter(10)
	router.Get("/aaa/@name:([a-z]+)/@age:([0-9]+)", nil)
	router.Put("/bbb/@age:([0-9]+)", nil)
	router.Get("/ccc/@name:([a-z]+)/@age:([0-9]+)", nil)
	router.Lookup("GET", "/ccc/liubang/26")
}
