package main

import (
    "github.com/gin-gonic/gin"
    "taskmanager/auth"
    "taskmanager/routes"
    "taskmanager/websocket"
    "taskmanager/ai"
)

func main() {
    r := gin.Default()

    // Authentication routes
    r.POST("/signup", auth.Signup)
    r.POST("/login", auth.Login)

    // Protected routes
    protected := r.Group("/api")
    protected.Use(auth.AuthMiddleware())
    {
        protected.POST("/tasks", routes.CreateTask)
        protected.GET("/tasks", routes.GetTasks)
        protected.POST("/suggest", ai.GetTaskSuggestions) // Add AI suggestion endpoint
    }

    // WebSocket route
    r.GET("/ws", func(c *gin.Context) {
        websocket.HandleConnections(c.Writer, c.Request)
    })

    // Start the server
    r.Run(":8080")
}
