package main

import (
	"context"
	"net/http"

	"github.com/FereshtehDehghani/golang-todo-REST-API/config"
	"github.com/FereshtehDehghani/golang-todo-REST-API/db"
	"github.com/FereshtehDehghani/golang-todo-REST-API/routes"
)


func main(){

	 handler :=routes.MuonteRoutes()

	 config.Config.LoadConfig()
	
 db.InitDB()
	server :=&http.Server{
		Addr: config.Config.AppPort,
		Handler: handler,
	}
	
defer db.DB.Close(context.Background())

	server.ListenAndServe()
}

