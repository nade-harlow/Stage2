package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
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
	r.POST("/", calculate())
	err := r.Run()
	if err != nil {
		log.Println("An error occurred while starting up server")
		return
	}
}

func calculate() gin.HandlerFunc {
	return func(c *gin.Context) {
		problem := MathRequest{}
		err := c.ShouldBind(&problem)
		log.Println("INPUTS: ", problem)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "check fields and try again"})
			return
		}
		if problem.OperationType == "" {
			c.JSON(http.StatusBadRequest, "The operation type field is required.")
		}
		solution := solveProblem(problem)
		if solution == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid arithmetic operation"})
			return
		}
		c.JSON(200, solution)
		return
	}
}

func solveProblem(problem MathRequest) *MathResponse {
	var result int64
	if strings.Contains(string(problem.OperationType), "add") ||
		strings.Contains(string(problem.OperationType), "addition") ||
		strings.Contains(string(problem.OperationType), "plus") {
		result = problem.X + problem.Y
		solution := &MathResponse{
			SlackUsername: "Nade",
			OperationType: addition,
			Result:        result,
		}
		return solution

	} else if strings.Contains(string(problem.OperationType), "sub") ||
		strings.Contains(string(problem.OperationType), "subtract") ||
		strings.Contains(string(problem.OperationType), "minus") ||
		strings.Contains(string(problem.OperationType), "subtraction") {
		result = problem.X - problem.Y
		solution := &MathResponse{
			SlackUsername: "Nade",
			OperationType: subtraction,
			Result:        result,
		}
		return solution

	} else if strings.Contains(string(problem.OperationType), "multiply") ||
		strings.Contains(string(problem.OperationType), "multiplied") ||
		strings.Contains(string(problem.OperationType), "multiplication") ||
		strings.Contains(string(problem.OperationType), "times") {
		result = problem.X * problem.Y
		solution := &MathResponse{
			SlackUsername: "Nade",
			OperationType: multiplication,
			Result:        result,
		}
		return solution

	} else {
		solution := &MathResponse{}
		return solution
	}
}
