package grouter

import (
	"fmt"
	"testing"
)

func TestRouter_Lookup(t *testing.T) {
	router := NewRouter()
	router.Get("/aaa/@name:([a-z]+)/@age:([0-9]+)", nil)
	router.Put("/bbb/@age:([0-9]+)", nil)
	router.Get("/ccc/@name:([a-z]+)/@age:([0-9]+)", nil)
	router.Delete("/this/is/test", nil)
	rule, found := router.Lookup("GET", "/ccc/liubang/26")
	if !found {
		t.Error("查找路由失败")
	} else {
		fmt.Println(rule)
	}

	rule, found = router.Lookup("DELETE", "/this/is/test")
	if !found {
		t.Error("查找路由失败")
	} else {
		fmt.Println(rule)
	}
}
