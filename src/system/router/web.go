package router

import (
	controller "Structure/src/Controller"
	db "Structure/src/system/db"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var DB = db.Connect()

func WebRoutes(r *mux.Router) {
	// r.HandleFunc("/", middleware.AuthRequired(indexGetHandler)).Methods("GET")
	r.HandleFunc("/", HomeHandler)

	r.HandleFunc("/register", controller.Register).Methods("GET")
	r.HandleFunc("/login", controller.Login).Methods("POST")

	r.Handle("/check", IsAuthenticated(CheckHandler))

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// token, err := jwt.CreateJwtToken()
	// if err != nil {
	// 	w.Write([]byte("err"))
	// } else {
	// 	w.Write([]byte(token))
	// }
	w.Write([]byte("Home"))
}

func CheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(strconv.Itoa(USER_ID)))
}
