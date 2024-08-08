package controllers

import (
	"net/http"

	"task-manager/data"
	"task-manager/models"

	"github.com/gin-gonic/gin"
)

func RegistrationHandler(ctx *gin.Context){

	var user models.User
	err := ctx.BindJSON(&user)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message":"invalid request"})
		return
	}
	_, message := data.CreateUser(user)

	if message != "" {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message":message})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message":"User registered successfully!"})
}

func LoginHandler(ctx *gin.Context){
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message":"invalid request"})
	}
	response, message := data.Login(user)
	if message != "" {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"token":message})
		return
	}

	ctx.IndentedJSON(http.StatusOK, response)
}

func PromotoionHandler(ctx *gin.Context){
    id := ctx.Param("id")
    

    _, err := data.PromoteUser(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    ctx.IndentedJSON(http.StatusOK, gin.H{"message": "User promoted successfully"})
}