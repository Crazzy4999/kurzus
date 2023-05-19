package router

import "net/http"

type Router struct{}

func (r *Router) Handle(url string, f http.Handler) {
	//mux := http.NewServeMux()
}
