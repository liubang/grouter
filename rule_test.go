package grouter

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewRule(t *testing.T) {
	Convey("case 1:", t, func() {
		rule := NewRule("POST", "/liubang/@name:([a-z]+)/@age:([0-9])", nil)
		keys := rule.ParamKeys
		So(len(keys), ShouldEqual, 2)
		So(keys[0], ShouldEqual, "")
		So(keys[1], ShouldEqual, "age")
	})
}
