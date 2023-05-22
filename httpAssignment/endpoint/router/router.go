package router

type Router struct {
	Port string
}

type parameter string

func parseURL(url string) []parameter {
	recording := false
	params := make([]parameter, 0)
	param := ""
	for index, bChar := range url {
		char := string(bChar)
		if char == ":" {
			recording = true
			continue
		}
		if recording {
			param += char
		}
		if char == "/" || index == len(url)-1 {
			recording = false
			params = append(params, parameter(param))
			param = ""
			continue
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

func (router *Router) Start() {
	//router.Port
}
