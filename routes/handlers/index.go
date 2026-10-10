package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/FereshtehDehghani/golang-todo-REST-API/db"
	"github.com/gin-gonic/gin"
)


type PostTaskPayload struct{
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status string `json:"status" binding:"required"`
}

func SaveTask(c *gin.Context){
	var payload PostTaskPayload



	if err := c.ShouldBindJSON(&payload); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
         log.Printf(string(err.Error()))
		return
	}

	var id int

	query :=`Insert into tasks (title,description,status) VALUES ($1,$2,$3) RETURNING id;`    
	err := db.DB.QueryRow(context.Background(),query,payload.Title,payload.Description,payload.Status).Scan(&id)

		if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":true,"msg":err.Error()})
         log.Printf(string(err.Error()))
		return
	}

	// log.Printf(string(payload))
	c.JSON(http.StatusOK,gin.H{"error":false,"msg":id})
}