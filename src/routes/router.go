package routes

import (
	"github.com/gorilla/mux"
	"go_project/src/users/controllers"
	"net/http"
)

type Route struct {
	Router *mux.Router
	Serv   *http.Server
	Action controllers.AppController
}

type Router interface {
	createRoute()
   initializeRoutes()
}

func (a *Route) CreateRoute() {
	a.initializeRoutes()
}
func (a *Route) initializeRoutes() {
	//user routes
	func (u *controllers.UserController) {
		a.Router.Use(controllers.JwtAuthentication)
		a.Router.HandleFunc("/api/users", u.GetUsers).Methods("GET")
		a.Router.HandleFunc("/api/user/register", u.UserRegister).Methods("POST")
		//jwtMiddleware.Handler(c)
		a.Router.HandleFunc("/api/user/login", u.LoginUser).Methods("POST")
		a.Router.HandleFunc("/api/user/{id:[0-9]+}", u.GetUser).Methods("GET")
		a.Router.HandleFunc("/api/user/update/{id:[0-9]+}", u.UpdateUser).Methods("PUT")
		a.Router.HandleFunc("/api/user/delete/{id:[0-9]+}", u.DeleteUser).Methods("DELETE")
	}(&controllers.UserController{
		AppController: &a.Action,
		UserService: a.Action.Services.UserService,
	})
}
