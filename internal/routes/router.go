package routes

import (
	"github.com/gorilla/mux"
	controllers3 "go_project/internal/users/user/controllers"
	controllers2 "go_project/pkg/common/controllers"
	"net/http"
)

type Route struct {
	Router *mux.Router
	Serv   *http.Server
	Action controllers2.AppController
}

type Router interface {
	createRoute()
   initializeRoutes()
}

func (a *Route) CreateRoute() {
	a.initializeRoutes()
}
func (a *Route) initializeRoutes() {
	a.Router.Use(a.Action.JwtAuthentication)
	//user routes
	func (u *controllers3.UserController) {
		a.Router.HandleFunc("/", u.Home).Methods("GET")
		a.Router.HandleFunc("/api/users", u.GetUsers).Methods("GET")
		////jwtMiddleware.Handler(c)
		a.Router.HandleFunc("/api/user/{id:[0-9]+}", u.GetUser).Methods("GET")
		a.Router.HandleFunc("/api/user/update/{id:[0-9]+}", u.UpdateUser).Methods("PUT")
		a.Router.HandleFunc("/api/user/delete/{id:[0-9]+}", u.DeleteUser).Methods("DELETE")
	}(&controllers3.UserController{
		AppController: a.Action,
		UserService: a.Action.Services.UserService,
	})
	//auth routes
	func (u *controllers3.AuthController) {
		a.Router.HandleFunc("/api/user/register", u.UserRegister).Methods("POST")
		a.Router.HandleFunc("/api/user/login", u.LoginUser).Methods("POST")
	}(&controllers3.AuthController{
		AppController: a.Action,
		UserService: a.Action.Services.AuthService,
	})
}
