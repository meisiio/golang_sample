package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func userHnadler(c *gin.Context) {
	var someone user
	if err := c.ShouldBindWith(&someone, binding.FormPost); err == nil {
		c.JSON(http.StatusOK, someone)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	log.Printf("some one is : %+v", someone)
}
