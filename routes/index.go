package routes

import (
	"github.com/FereshtehDehghani/golang-todo-REST-API/routes/handlers"
	"github.com/gin-gonic/gin"
)


func MuonteRoutes() *gin.Engine{

	 handler :=gin.Default()

	 handler.POST("/task",handlers.SaveTask)

	 return handler
}