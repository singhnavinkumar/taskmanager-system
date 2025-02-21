package ai

import (
    "bytes"
    "encoding/json"
    "net/http"
    "github.com/gin-gonic/gin"
)

func GetTaskSuggestions(c *gin.Context) {
    var request struct {
        Prompt string `json:"prompt"`
    }
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(400, gin.H{"error": "Invalid input"})
        return
    }

    // Prepare the request body for OpenAI API
    requestBody, _ := json.Marshal(map[string]interface{}{
        "model":      "text-davinci-003", // Use the appropriate OpenAI model
        "prompt":     request.Prompt,
        "max_tokens": 50,
    })

    // Create a new HTTP request to OpenAI API
    req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(requestBody))
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to create OpenAI request"})
        return
    }

    // Set headers for OpenAI API
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer YOUR_OPENAI_API_KEY") // Replace with your OpenAI API key

    // Send the request to OpenAI API
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to call OpenAI API"})
        return
    }
    defer resp.Body.Close()

    // Parse the response from OpenAI API
    var result map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        c.JSON(500, gin.H{"error": "Failed to decode OpenAI response"})
        return
    }

    // Extract the suggestion from the response
    choices, ok := result["choices"].([]interface{})
    if !ok || len(choices) == 0 {
        c.JSON(500, gin.H{"error": "No suggestions found"})
        return
    }
    suggestion := choices[0].(map[string]interface{})["text"].(string)

    // Return the suggestion to the client
    c.JSON(200, gin.H{"suggestion": suggestion})
}