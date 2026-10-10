package handlers

import (
	"log"
	"net/http"

	"github.com/FereshtehDehghani/golang-todo-REST-API/db"
	"github.com/gin-gonic/gin"
)




func SaveTask(c *gin.Context){
	var payload db.PostTaskPayload



	if err := c.ShouldBindJSON(&payload); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
         log.Printf(string(err.Error()))
		return
	}

id,err := db.TaskRepository.SaveTaskQuery(payload)
		if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":true,"msg":err.Error()})
         log.Printf(string(err.Error()))
		return
	}

	// log.Printf(string(payload))
	c.JSON(http.StatusOK,gin.H{"error":false,"msg":id})
}