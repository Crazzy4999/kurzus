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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	for _, route := range router.routes {
		if route.url.MatchString(r.URL.String()) && route.method == r.Method {
			if route.m != nil {
				route.m(route.handler).ServeHTTP(w, r)
				return
			}
			route.handler.ServeHTTP(w, r)
			return
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
