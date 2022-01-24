package main

import (
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
	//_ = config.InitWorkers
	a.Initialize(
		"goland",//os.Getenv("APP_DB_USERNAME"),
		"goland",//os.Getenv("APP_DB_PASSWORD"),
		"goland",//os.Getenv("APP_DB_NAME"),
		"disable",//os.Getenv("SSL_MODE"),
		)
	userModule := &controllers.App{DB: a.DB}
	route := mux.NewRouter()
	r := routes.Route{UserAction: *userModule, Router: route}
	r.CreateRoute()
	//quoteChan := parser.Grab()
	//for i := 0; i < 5; i++ { //получаем 5 цитат и закругляемся
	//	fmt.Println(<-quoteChan, "\n")
	//}
	r.Run("localhost:8080")
}