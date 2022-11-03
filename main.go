package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type MathRequest struct {
	OperationType OperationType `json:"operation_type"`
	X             int64         `json:"x"`
	Y             int64         `json:"y"`
}

type MathResponse struct {
	SlackUsername string        `json:"slackUsername"`
	OperationType OperationType `json:"operation_type"`
	Result        int64         `json:"result"`
}

type OperationType string

const (
	multiplication OperationType = "multiplication"
	addition       OperationType = "addition"
	subtraction    OperationType = "subtraction"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.POST("/", Calculate())
	err := r.Run()
	if err != nil {
		log.Println("An error occurred while starting up server")
		return
	}
}

func Calculate() gin.HandlerFunc {
	return func(c *gin.Context) {
		problem := MathRequest{}
		c.ShouldBind(&problem)
		var result int64
		switch problem.OperationType {
		case addition:
			result = problem.X + problem.Y
		case multiplication:
			result = problem.X * problem.Y
		case subtraction:
			result = problem.X - problem.Y
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid arithmetic operation"})
			return
		}
		solution := MathResponse{
			SlackUsername: "Nade",
			OperationType: problem.OperationType,
			Result:        result,
		}
		c.JSON(http.StatusOK, solution)
		return
	}
}
