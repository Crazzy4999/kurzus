package router

import "strconv"

type Router struct{}

type parameter string

func parseURL(url string) []parameter {
	recording := false
	params := make([]parameter, 0)
	param := ""
	for bChar := range url {
		char := strconv.Itoa(bChar)
		if char == ":" {
			recording = true
			continue
		}
		if char == "/" {
			recording = false
			params = append(params, parameter(param))
			param = ""
			continue
		}
		if recording {
			param += char
		}
	}
	return params
}

func (router *Router) POST(url string) {

}

func (router *Router) GET(url string) {

}

func (router *Router) PUT(url string) {

}

func (router *Router) DELETE(url string) {

}
