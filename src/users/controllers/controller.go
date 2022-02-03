package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq"
	"go_project/src/common"
	"go_project/src/users/models"
	"net/http"
	"os"
	"strings"
	"time"
)

type AppController struct {
	Application *common.App
}

type Token struct {
	Username string
	jwt.MapClaims
}
var secretKey = os.Getenv("SECRET_KEY")

func createToken(p *models.User) {
	var mySigningKey = []byte(secretKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["admin"] = true
	claims["username"] = p.Username
	claims["email"] = p.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token.Claims = claims
	tokenString, _ := token.SignedString(mySigningKey)
	p.Token = tokenString
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		notAuth := []string{"/api/user/login", "/api/users", "/api/user/register"} //List of endpoints that doesn't require auth
		requestPath := r.URL.Path //current request path

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization") //Grab the token from the header

		if tokenHeader == "" {
			c := map[string]string{"message": "Missing auth token"}//Token is missing, returns with error code 403 Unauthorized
			respondWithJSON(w, http.StatusForbidden, c)
			return
		}

		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitted) != 2 {
			c := map[string]string{"message": "Invalid/Malformed auth token"}//Token is missing, returns with error code 403 Unauthorized
			respondWithJSON(w, http.StatusForbidden, c)
			return
		}

		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
		tk := &Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil { //Malformed token, returns with users code 403 as usual
			c := map[string]string{"message": "Malformed authentication token"}//Token is missing, returns with error code 403 Unauthorized
			respondWithJSON(w, http.StatusForbidden, c)
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			c := map[string]string{"message": "Token is not valid."}//Token is missing, returns with error code 403 Unauthorized
			respondWithJSON(w, http.StatusForbidden, c)
			return
		}

		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		fmt.Sprintf("User %", tk.Username) //Useful for monitoring
		ctx := context.WithValue(r.Context(), "user", tk.Username)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) //proceed in the middleware chain!
	})
}
