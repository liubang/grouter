package grouter

import (
	"net/http"
	"strings"
)

type Router struct {
	Rules    []*Rule
	NotFound HandlerFunc
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request, pathvariables map[string]string)

var (
	NotFound = func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
		http.NotFound(w, r)
	}
)

// 创建一个Router
func NewRouter() *Router {
	return &Router{}
}

func (router *Router) SetNotFound(notFoundHandler HandlerFunc) {
	router.NotFound = notFoundHandler
}

func (router *Router) AddRule(rule *Rule) {
	router.Rules = append(router.Rules, rule)
}

func (router *Router) add(method string, path string, handler HandlerFunc) {
	rule := NewRule(method, path, handler)
	router.AddRule(rule)
}

func (router *Router) Get(path string, handler HandlerFunc) {
	router.add(http.MethodGet, path, handler)
}

func (router *Router) Put(path string, handler HandlerFunc) {
	router.add(http.MethodPut, path, handler)
}

func (router *Router) Post(path string, handler HandlerFunc) {
	router.add(http.MethodPost, path, handler)
}

func (router *Router) Delete(path string, handler HandlerFunc) {
	router.add(http.MethodDelete, path, handler)
}

func (router *Router) Option(path string, handler HandlerFunc) {
	router.add(http.MethodOptions, path, handler)
}

func (router *Router) Head(path string, handler HandlerFunc) {
	router.add(http.MethodHead, path, handler)
}

func (router *Router) Connect(path string, handler HandlerFunc) {
	router.add(http.MethodConnect, path, handler)
}

func (router *Router) Lookup(method string, uri string) (*Rule, bool) {
	method = strings.ToUpper(method)
	for _, rule := range router.Rules {
		if rule.Method == method {
			if rule.matcher == nil {
				// string match
				if rule.Reg == uri {
					return rule, true
				}
			} else {
				matches := rule.matcher.FindStringSubmatch(uri)
				if len(matches) == 0 {
					continue
				}
				for idx, item := range matches {
					if idx == 0 {
						continue
					}
					rule.Params[rule.ParamKeys[idx-1]] = item
				}
				return rule, true
			}
		}
	}

	return nil, false
}

// implements http.Handler interface.
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rule, f := router.Lookup(r.Method, r.URL.Path)
	if f {
		rule.Handler(w, r, rule.Params)
	} else {
		if router.NotFound != nil {
			router.NotFound(w, r, nil)
		} else {
			NotFound(w, r, nil)
		}
	}
}
