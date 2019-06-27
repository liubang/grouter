package Grouter

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type Router struct {
	Rules     []*Rule
	ChunkSize int
}

// 创建一个Router
func NewRouter(chunkSize int) *Router {
	return &Router{
		ChunkSize: chunkSize,
	}
}

func (router *Router) AddRule(rule *Rule) {
	router.Rules = append(router.Rules, rule)
}

func (router *Router) add(method string, path string, handler http.HandlerFunc) {
	rule := NewRule(method, path, handler)
	router.AddRule(rule)
}

func (router *Router) Get(path string, handler http.HandlerFunc) {
	router.add("GET", path, handler)
}

func (router *Router) Put(path string, handler http.HandlerFunc) {
	router.add("PUT", path, handler)
}

func (router *Router) Post(path string, handler http.HandlerFunc) {
	router.add("POST", path, handler)
}

func (router *Router) Delete(path string, handler http.HandlerFunc) {
	router.add("DELETE", path, handler)
}

func (router *Router) Option(path string, handler http.HandlerFunc) {
	router.add("OPTION", path, handler)
}

func (router *Router) Head(path string, handler http.HandlerFunc) {
	router.add("HEAD", path, handler)
}

func (router *Router) Lookup(method string, uri string) {
	method = strings.ToUpper(method)
	preg := "(?:"
	for idx, rule := range router.Rules {
		if rule.Method == method {
			preg += fmt.Sprintf("(?P<grouter_%d>%s)|", idx, rule.Reg)
		}
	}
	preg = strings.TrimRight(preg, "|")
	preg += ")"
	fmt.Println(preg)
	matcher := regexp.MustCompile(preg)
	matches := matcher.FindAllStringSubmatch(uri, -1)
	groups := matcher.SubexpNames()
	fmt.Println(groups)
	fmt.Println(matches)
	for idx, match := range matches {
		fmt.Println(groups[idx])
		fmt.Println(match)
	}
}
