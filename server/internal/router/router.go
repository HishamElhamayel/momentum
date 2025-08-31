package router

import (
	"example.com/momentum/internal/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) *gin.Engine{
	
	r := gin.Default()

	r.GET("/ping", handlers.PingHandler)

	return r
}