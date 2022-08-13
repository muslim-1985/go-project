package main

import (
	"fmt"
	"go_project/internal/users/user/store"
	"go_project/pkg/common/controllers"
	"go_project/pkg/common/di"

	//"fmt"
	"log"
	//"net/http"
	"go_project/internal/config"
	"go_project/internal/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env.example file")
	}

	a := config.App{}

	a.InitializeDb(
		"goland",  //os.Getenv("APP_DB_USERNAME"),
		"goland",  //os.Getenv("APP_DB_PASSWORD"),
		"goland",  //os.Getenv("APP_DB_NAME"),
		"disable", //os.Getenv("SSL_MODE"),
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
	fmt.Println("fdsfddfdf")
	//quoteChan := parser.Grab()
	//for i := 0; i < 5; i++ { //получаем 5 цитат и закругляемся
	//	fmt.Println(<-quoteChan, "\n")
	//}
	ServerRun(":8080", r)
}
