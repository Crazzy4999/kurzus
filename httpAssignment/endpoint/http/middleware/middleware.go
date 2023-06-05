package middleware

import "net/http"

type Func func(handler http.Handler) http.Handler
