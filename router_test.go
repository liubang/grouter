package grouter

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouter_Lookup(t *testing.T) {
	router := NewRouter()
	router.Get("/aaa/@name:([a-z]+)/@age:([0-9]+)", nil)
	router.Put("/bbb/@age:([0-9]+)", nil)
	router.Get("/ccc/@name:([a-z]+)/@age:([0-9]+)", nil)
	router.Delete("/this/is/test", nil)
	rule, found := router.Lookup(http.MethodGet, "/ccc/liubang/26")
	assert.Equal(t, true, found)
	assert.Equal(t, "liubang", rule.Params["name"])
	assert.Equal(t, "26", rule.Params["age"])
	_, found = router.Lookup(http.MethodDelete, "/this/is/test")
	assert.Equal(t, true, found)
	_, found = router.Lookup(http.MethodPut, "/this/is/test")
	assert.Equal(t, false, found)
}
