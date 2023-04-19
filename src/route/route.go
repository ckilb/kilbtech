package route

import "net/http"

type Route interface {
	Path() string
	Handler() http.Handler
}
