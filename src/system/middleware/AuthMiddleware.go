package middleware

import (
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Person struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type MyClaims struct {
	User Person
	jwt.StandardClaims
}

// Standared Claim properties
// Audience  string
// ExpiresAt int64
// Id        string
// IssuedAt  int64
// Issuer    string
// NotBefore int64
// Subject   string

func CreateJwtToken(id int, name string, email string) (map[string]string, error) {

	jwtTokenSecret := []byte(getJwtTokenSecretKey())

	// Create the Claims
	claims := MyClaims{
		Person{
			ID:    id,
			Name:  name,
			Email: email,
		},
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Issuer:    name,
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := rawToken.SignedString(jwtTokenSecret)

	if err != nil {
		var userInfo map[string]string
		return userInfo, err
	}

	userInfo := map[string]string{
		"id":         strconv.Itoa(claims.User.ID),
		"name":       claims.User.Name,
		"email":      claims.User.Email,
		"token":      token,
		"expires_at": strconv.FormatInt(claims.ExpiresAt, 10), // 10 means decimal convertion
	}

	return userInfo, nil
}

func PurseToken(t string) (string, error) {
	tokenString := t

	jwtTokenSecret := []byte(getJwtTokenSecretKey())

	var response string
	var errr error

	if tokenString == "" {
		return response, errr
	}

	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtTokenSecret), nil
	})

	if claims, ok := token.Claims.(*MyClaims); ok {
		if time.Now().Unix() > claims.StandardClaims.ExpiresAt {
			response = "Your token is expired"
		} else if token.Valid {
			response = strconv.Itoa(claims.User.ID)
		}
	} else {
		errr = err
	}
	return response, errr
}

func getJwtTokenSecretKey() string {
	jwt_secret := os.Getenv("JWT_TOKEN_SECRET")

	if len(jwt_secret) > 0 {
		return jwt_secret
	}
	return ""
}
