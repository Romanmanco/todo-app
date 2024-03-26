package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type errorHandlers struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Print(message)
	c.AbortWithStatusJSON(statusCode, errorHandlers{message})
}
