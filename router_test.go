package Grouter

import (
	"fmt"
	"testing"
)

func TestRouter_Lookup(t *testing.T) {
	router := NewRouter()
	router.Get("/aaa/@name:([a-z]+)/@age:([0-9]+)", nil)
	router.Put("/bbb/@age:([0-9]+)", nil)
	router.Get("/ccc/@name:([a-z]+)/@age:([0-9]+)", nil)
	rule, found := router.Lookup("GET", "/ccc/liubang/26")
	if found {
		fmt.Println(rule)
	}
}
