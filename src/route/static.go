package route

import (
	"net/http"
)

type Static struct {
	staticPath string
}

func (r *Static) Path() string {
	return "/static/"
}

func (r *Static) Handler() http.Handler {
	fs := http.FileServer(http.Dir(r.staticPath))
	h := http.StripPrefix("/static/", fs)

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Cache-Control", "max-age=31536000")
		h.ServeHTTP(w, req)
	})
}

func NewStatic(staticPath string) Route {
	return &Static{
		staticPath: staticPath,
	}
}
