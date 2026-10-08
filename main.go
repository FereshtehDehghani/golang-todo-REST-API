package main

import (
	"context"
	"net/http"

	"github.com/FereshtehDehghani/golang-todo-REST-API/config"
	"github.com/FereshtehDehghani/golang-todo-REST-API/db"
	"github.com/gin-gonic/gin"
)


func main(){
 db.InitDB()
	 handler :=gin.Default()


	 config.Config.LoadConfig()

	

	server :=&http.Server{
		Addr: config.Config.AppPort,
		Handler: handler,
	}
	
defer db.DB.Close(context.Background())

	server.ListenAndServe()
}

