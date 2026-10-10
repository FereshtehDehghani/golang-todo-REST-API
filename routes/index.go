package routes

import (
	"net/http"

	"github.com/FereshtehDehghani/golang-todo-REST-API/routes/handlers"
	"github.com/gin-gonic/gin"
)


func MuonteRoutes() *gin.Engine{

	 handler :=gin.Default()

	 handler.POST("/task",handlers.SaveTask)

	 handler.NoRoute(func(c *gin.Context){
        c.JSON(http.StatusNotFound,gin.H{"message":"Routr not found"})
	 })

	 return handler
}