package Grouter

import (
	"net/http"
	"regexp"
)

type Rule struct {
	Method  string
	Path    string
	Reg     string
	Params  map[string]string
	Handler http.HandlerFunc
}

var (
	matcher = regexp.MustCompile("(?:@(.*?):)")
)

// create a new rule
func NewRule(method string, path string, handler http.HandlerFunc) *Rule {
	rule := &Rule{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
	params := make(map[string]string)
	matches := matcher.FindAllStringSubmatch(path, -1)
	if len(matches) != 0 {
		for _, p := range matches {
			params[p[1]] = ""
		}
		reg := matcher.ReplaceAllString(path, "")
		rule.Reg = reg
	} else {
		rule.Reg = path
	}
	rule.Params = params

	return rule
}
