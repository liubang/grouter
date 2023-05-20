package grouter

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRouter_Lookup(t *testing.T) {
	Convey("case 1:", t, func() {
		router := NewRouter()
		router.Get("/aaa/@name:([a-z]+)/@age:([0-9]+)", nil)
		router.Put("/bbb/@age:([0-9]+)", nil)
		router.Get("/ccc/@name:([a-z]+)/@age:([0-9]+)", nil)
		router.Delete("/this/is/test", nil)
		rule, found := router.Lookup(http.MethodGet, "/ccc/liubang/26")

		So(found, ShouldBeTrue)
		So(rule.Params["name"], ShouldEqual, "liubang")
		So(rule.Params["age"], ShouldEqual, "26")

		_, found = router.Lookup(http.MethodDelete, "/this/is/test")
		So(found, ShouldBeTrue)
		_, found = router.Lookup(http.MethodPut, "/this/is/test")
		So(found, ShouldBeFalse)
	})
}
