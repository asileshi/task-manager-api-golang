package controllers

import (
	"net/http"
	"time"

	"task-manager/data"
	"task-manager/models"

	"github.com/gin-gonic/gin"
)
func GetTasksHandler(ctx *gin.Context) {
    tasks, err := data.GetTasks()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.IndentedJSON(http.StatusOK, tasks)
}

func GetTaskByIDHandler(ctx *gin.Context) {
    id := ctx.Param("id")
    task, err := data.GetTaskByID(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    ctx.IndentedJSON(http.StatusOK, task)

}

func CreateTaskHandler(ctx *gin.Context) {
    var task models.Task
    err := ctx.BindJSON(&task)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	if task.DueDate.IsZero() {
		task.DueDate = time.Now().UTC()
	}
    task, err = data.CreateTask(task)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.IndentedJSON(http.StatusCreated, task)

}

func UpdateTaskHandler(ctx *gin.Context) {
    id := ctx.Param("id")
    var task models.Task
    err := ctx.BindJSON(&task)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    task, err = data.UpdateTask(id, task)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    ctx.IndentedJSON(http.StatusOK, task)
}

func DeleteTaskHandler(ctx *gin.Context) {
    id := ctx.Param("id")
    err := data.DeleteTask(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    ctx.IndentedJSON(http.StatusNoContent, gin.H{})

}