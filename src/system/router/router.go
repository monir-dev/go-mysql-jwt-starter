package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	WebRoutes(r)

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
}
