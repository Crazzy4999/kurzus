package router

import (
	"fmt"
	"hangryAPI/internal/endpoint/middleware"
	"net/http"
	"regexp"
)

type route struct {
	method  string
	url     *regexp.Regexp
	handler http.Handler
	m       middleware.Func
}

type Router struct {
	routes []route
}

func NewRouter() *Router {
	return &Router{}
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range router.routes {
		if route.url.MatchString(r.URL.String()) && route.method == r.Method {
			if route.m != nil {
				route.m(route.handler).ServeHTTP(w, r)
				return
			}
			route.handler.ServeHTTP(w, r)
			return
		} else if route.url.MatchString(r.URL.Path) && route.method != r.Method {
			http.Error(w, fmt.Sprintf("Only %s method is allowed for route %s", route.method, r.URL.Path), http.StatusBadRequest)
		}
	}
}

func (router *Router) addRoute(method string, url string, handler func(w http.ResponseWriter, r *http.Request), middleware middleware.Func) {
	router.routes = append(router.routes, route{
		method:  method,
		url:     regexp.MustCompile(fmt.Sprintf("^%s$", url)),
		handler: http.HandlerFunc(handler),
		m:       middleware,
	})
}

func (router *Router) GET(url string, handler func(w http.ResponseWriter, r *http.Request), middleware middleware.Func) {
	router.addRoute(http.MethodGet, url, handler, middleware)
}

func (router *Router) POST(url string, handler func(w http.ResponseWriter, r *http.Request), middleware middleware.Func) {
	router.addRoute(http.MethodPost, url, handler, middleware)
}

func (router *Router) PUT(url string, handler func(w http.ResponseWriter, r *http.Request), middleware middleware.Func) {
	router.addRoute(http.MethodPut, url, handler, middleware)
}

func (router *Router) DELETE(url string, handler func(w http.ResponseWriter, r *http.Request), middleware middleware.Func) {
	router.addRoute(http.MethodDelete, url, handler, middleware)
}
