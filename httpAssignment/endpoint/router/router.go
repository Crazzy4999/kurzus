package router

import (
	"fmt"
	"httpA/endpoint/middleware"
	"net/http"
	"regexp"
)

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.Handler
	m       middleware.Func
}

type Router struct {
	port   string
	Routes []route
}

func NewRouter(port string) *Router {
	router := new(Router)
	router.port = port
	return router
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range router.Routes {
		if route.regex.MatchString(r.URL.Path) && route.method == r.Method {
			if route.m != nil {
				route.m(route.handler).ServeHTTP(w, r)
			} else {
				route.handler.ServeHTTP(w, r)
			}
			return
		}
	}
	http.NotFound(w, r)
}

func (router *Router) addRoute(method, regex string, f func(w http.ResponseWriter, r *http.Request), m middleware.Func) {
	router.Routes = append(router.Routes, route{
		method:  method,
		regex:   regexp.MustCompile(fmt.Sprint("^%s$", regex)),
		handler: http.HandlerFunc(f),
		m:       m,
	})
}

func (router *Router) POST(regexURL string, f func(w http.ResponseWriter, r *http.Request), m middleware.Func) {
	router.addRoute(http.MethodPost, regexURL, f, m)
}

func (router *Router) GET(regexURL string, f func(w http.ResponseWriter, r *http.Request), m middleware.Func) {
	router.addRoute(http.MethodGet, regexURL, f, m)
}

func (router *Router) PUT(regexURL string, f func(w http.ResponseWriter, r *http.Request), m middleware.Func) {
	router.addRoute(http.MethodPut, regexURL, f, m)
}

func (router *Router) DELETE(regexURL string, f func(w http.ResponseWriter, r *http.Request), m middleware.Func) {
	router.addRoute(http.MethodDelete, regexURL, f, m)
}

func (router *Router) Start() {
	fmt.Printf("Server started on: localhost%s\n", router.port)
	http.ListenAndServe(router.port, router)
}
