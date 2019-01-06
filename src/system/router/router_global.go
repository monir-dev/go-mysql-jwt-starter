package router

import (
	jwt "Structure/src/system/middleware"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var USER_ID int

// Check if current request is authenticated
func IsAuthenticated(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		bearer := r.Header.Get("Authorization")
		if bearer != "" {
			token := strings.Replace(bearer, "Bearer ", "", 1)
			response, err := jwt.PurseToken(token)

			checkErr(err)

			// set user id
			user_id, err := strconv.Atoi(response)
			checkErr(err)
			USER_ID = user_id

			// process request
			endpoint(w, r)

		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

// Error check
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
