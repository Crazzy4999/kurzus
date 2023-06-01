package router

import (
	"fmt"
	"net/http"
)

type Route struct {
	method  string
	url     string
	handler http.Handler
}

type Router struct {
	port   string
	Params []parameter
	Routes []Route
}

func NewRouter(port string) *Router {
	router := new(Router)
	router.port = port
	return router
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// if serve http is implemented than we can serve pages based on a result
	http.NotFound(w, r)
}

type parameter string

func (router *Router) parseURL(url string) []parameter {
	router.Params = nil
	recording := false
	params := make([]parameter, 0)
	param := ""
	for index, bChar := range url {
		char := string(bChar)
		if char == ":" {
			recording = true
			continue
		}
		if recording && char != "/" {
			param += char
		}
		if (char == "/" || index == len(url)-1) && param != "" {
			recording = false
			params = append(params, parameter(param))
			param = ""
			continue
		}
	}
	router.Params = params
	return params
}

func (router *Router) GetParam(param parameter) parameter {
	for _, p := range router.Params {
		if p == param {
			return p
		}
	}
	return ""
}

func (router *Router) addRoute() {

}

func (router *Router) POST(url string, f func(w http.ResponseWriter, r *http.Request)) {

}

func (router *Router) GET(url string, f func(w http.ResponseWriter, r *http.Request)) {
	router.parseURL(url)

	/*
		When there is a dynamic parameter like :id
		then parse the url to 'suppliers/' and a map of parameters with any values like
		map[":id"] = 1
		than based on the new url that's going to be generated which is suppliers/1
		run the ServeHTTP function on that handler which is passed by 'f' of type func(...) {}
	*/

}

func (router *Router) PUT(url string, f func(w http.ResponseWriter, r *http.Request)) {

}

func (router *Router) DELETE(url string, f func(w http.ResponseWriter, r *http.Request)) {

}

func (router *Router) Start() {
	fmt.Printf("Server started on: localhost%s\n", router.port)
	http.ListenAndServe(router.port, router)
}
