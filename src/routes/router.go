package routes

import (
	"github.com/gorilla/mux"
	"go_project/src/users/controllers"
	"net/http"
)

type Route struct {
	Router *mux.Router
	Serv   *http.Server
	UserAction controllers.App
}

type Router interface {
	createRoute()
   initializeRoutes()
}

func (a *Route) CreateRoute() {
	a.initializeRoutes()
}
func (a *Route) initializeRoutes() {
	a.Router.Use(controllers.JwtAuthentication)
	a.Router.HandleFunc("/api/users", a.UserAction.GetUsers).Methods("GET")
	a.Router.HandleFunc("/api/user/register",  a.UserAction.UserRegister).Methods("POST")
	//jwtMiddleware.Handler(c)
	a.Router.HandleFunc("/api/user/login", a.UserAction.LoginUser).Methods("POST")
	a.Router.HandleFunc("/api/user/{id:[0-9]+}", a.UserAction.GetUser).Methods("GET")
	a.Router.HandleFunc("/api/user/update/{id:[0-9]+}", a.UserAction.UpdateUser).Methods("PUT")
	a.Router.HandleFunc("/api/user/delete/{id:[0-9]+}", a.UserAction.DeleteUser).Methods("DELETE")
}
