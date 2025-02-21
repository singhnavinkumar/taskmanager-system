package routes

import (
    "github.com/gin-gonic/gin"
    "taskmanager/models"
)

var tasks []models.Task

func CreateTask(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(400, gin.H{"error": "Invalid input"})
        return
    }
    tasks = append(tasks, task)
    c.JSON(200, gin.H{"message": "Task created", "task": task})
}

func GetTasks(c *gin.Context) {
    c.JSON(200, gin.H{"tasks": tasks})
}