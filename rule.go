package Grouter

import (
	"regexp"
	"strings"
)

type Rule struct {
	Method    string
	Path      string
	Reg       string
	matcher   *regexp.Regexp
	Params    map[string]string
	ParamKeys []string
	Handler   HandlerFunc
}

var (
	// HTTPMETHOD list the supported http methods.
	HTTPMETHOD = map[string]bool{
		"GET":       true,
		"POST":      true,
		"PUT":       true,
		"DELETE":    true,
		"PATCH":     true,
		"OPTIONS":   true,
		"HEAD":      true,
		"TRACE":     true,
		"CONNECT":   true,
		"MKCOL":     true,
		"COPY":      true,
		"MOVE":      true,
		"PROPFIND":  true,
		"PROPPATCH": true,
		"LOCK":      true,
		"UNLOCK":    true,
	}

	matcher = regexp.MustCompile("(?:@(.*?):)")
)

// create a new rule
func NewRule(method string, path string, handler HandlerFunc) *Rule {
	method = strings.ToUpper(method)
	if !HTTPMETHOD[method] {
		panic("invalid request method " + method)
	}
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
			rule.ParamKeys = append(rule.ParamKeys, p[1])
		}
		reg := matcher.ReplaceAllString(path, "")
		rule.Reg = reg
		rule.matcher = regexp.MustCompile(reg)
	} else {
		rule.Reg = path
	}
	rule.Params = params
	return rule
}
