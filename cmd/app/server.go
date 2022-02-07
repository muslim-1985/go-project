package main

import (
	"go_project/internal/routes"
	"log"
	"net/http"
	"time"
)

func serverRun(address string, route routes.Route)  {
	server := &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      route.Router,
	}
	log.Fatal(server.ListenAndServe())
}