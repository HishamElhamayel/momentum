package router

import (
	"example.com/momentum/internal/handlers"
	"example.com/momentum/internal/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) *gin.Engine{
	
	r := gin.Default()

	r.GET("/ping", handlers.PingHandler)

	authHandler := handlers.AuthHandler{DB: db}
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	habitHandler := handlers.HabitHandler{DB: db}
	habits := r.Group("/habits", middleware.AuthMiddleware())
	{
		habits.POST("/", habitHandler.CreateHabit)
		habits.GET("/", habitHandler.GetHabits)
		habits.PUT("/:id", habitHandler.UpdateHabit)
		habits.DELETE("/:id", habitHandler.DeleteHabit)
	}

	habitLogsHandler := handlers.HabitLogHandler{DB: db}
	r.POST("/:id/logs/toggle", middleware.AuthMiddleware(),habitLogsHandler.ToggleHabitLog)
	r.GET("/:id/logs", middleware.AuthMiddleware(),habitLogsHandler.GetLogs)

	return r
}