package router

import "fmt"

type Router struct {
	Port   string
	Params []parameter
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

func (router *Router) POST(url string) {

}

func (router *Router) GET(url string) {
	router.parseURL(url)
	for i, p := range router.Params {
		fmt.Println(i, p)
	}
}

func (router *Router) PUT(url string) {

}

func (router *Router) DELETE(url string) {

}

func (router *Router) Start() {
	//router.Port
}
