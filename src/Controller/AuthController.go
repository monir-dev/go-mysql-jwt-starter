package controller

import (
	model "Structure/src/Model"
	"Structure/src/system/db"
	jwt "Structure/src/system/middleware"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type Cred struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var DB = db.Connect()

func Login(w http.ResponseWriter, r *http.Request) {

	var user model.User

	// var cred Cred

	var res map[string]interface{}
	json.NewDecoder(r.Body).Decode(&res)

	// For Form data
	// var email = r.FormValue("email")
	// var password = r.FormValue("password")

	// DB.First(&user, 1)
	DB.Where("email = ? AND password = ?", res["email"], res["password"]).First(&user)

	if user.Email == "" {
		w.Write([]byte("Invalid Username and password"))
	} else {
		userInfo, err := jwt.CreateJwtToken(user.ID, user.Name, user.Email)
		if err != nil {
			w.Write([]byte("Token creation error"))
		} else {
			json.NewEncoder(w).Encode(userInfo)
		}
	}
}

func Register(w http.ResponseWriter, r *http.Request) {

	var user model.User
	// user.ID = 5
	user.CreatedAt = strconv.FormatInt(time.Now().Unix(), 10)
	user.UpdatedAt = strconv.FormatInt(time.Now().Unix(), 10)
	json.NewDecoder(r.Body).Decode(&user)
	DB.Create(&user)

	json.NewEncoder(w).Encode(user)
}
