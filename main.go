package main

import (
	"go_project/src/common/di"
	"go_project/src/users/store"
	//"fmt"
	"log"
	//"net/http"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go_project/src/config"
	"go_project/src/routes"
	"go_project/src/users/controllers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	a := config.App{}

	a.InitializeDb(
		"goland",//os.Getenv("APP_DB_USERNAME"),
		"goland",//os.Getenv("APP_DB_PASSWORD"),
		"goland",//os.Getenv("APP_DB_NAME"),
		"disable",//os.Getenv("SSL_MODE"),
	)

	services := di.AppService{}
	initServices := services.InitService()

	store.Init(a)
	//_ = config.InitWorkers
	controller := controllers.AppController{}
	controller.Services = initServices

	route := mux.NewRouter()
	r := routes.Route{Action: controller, Router: route}
	r.CreateRoute()
	//quoteChan := parser.Grab()
	//for i := 0; i < 5; i++ { //получаем 5 цитат и закругляемся
	//	fmt.Println(<-quoteChan, "\n")
	//}
	serverRun("localhost:8080", r)
}
