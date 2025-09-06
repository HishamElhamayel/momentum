package handlers

import (
	"fmt"
	"net/http"
	"time"

	"example.com/momentum/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HabitLogHandler struct {
DB *gorm.DB
}

func (h *HabitLogHandler) ToggleHabitLog (c *gin.Context){
	userID, _ := c.Get("userID")
	habitID := c.Param("id")

	fmt.Println(userID,habitID)

	var habit models.Habit
	result := h.DB.Where("id = ? AND user_id = ?", habitID, userID).First(&habit)
	if result.Error != nil{
		c.JSON(http.StatusNotFound,gin.H{"Error":"Couldn't find habit"})
		return
	}

	today := time.Now().Truncate(24 * time.Hour)

	var log models.HabitLog

	result = h.DB.Where("habit_id = ? AND date = ?", habit.ID, today).First(&log)
	
	fmt.Println(log)
	
	if result.Error == gorm.ErrRecordNotFound{
		log := models.HabitLog{
			HabitID: habit.ID,
			Date: today,
		}

		h.DB.Create(&log)
		c.JSON(http.StatusCreated, log)
		return
	}

	h.DB.Delete(&log)
	c.JSON(http.StatusOK, gin.H{"message": "Habit unchecked"})

	
}

func (h *HabitLogHandler) GetLogs(c *gin.Context) {
    userID, _ := c.Get("userID")
    habitID := c.Param("id")

    var habit models.Habit
    if err := h.DB.Where("id = ? AND user_id = ?", habitID, userID).First(&habit).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Habit not found"})
        return
    }

    var logs []models.HabitLog
    h.DB.Where("habit_id = ?", habit.ID).Order("date desc").Find(&logs)

    c.JSON(http.StatusOK, logs)
}

func (h *HabitLogHandler) GetStreak(c *gin.Context) {
    userID, _ := c.Get("userID")
    habitID := c.Param("id")

    var habit models.Habit
    result := h.DB.Where("id = ? AND user_id = ?", habitID, userID).First(&habit)
	
	if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Habit not found"})
        return
    }

    var logs []models.HabitLog
    h.DB.Where("habit_id = ?", habit.ID).Order("date desc").Find(&logs)

    streak := 0
    today := time.Now().Truncate(24 * time.Hour)

    for i := 0; i < len(logs); i++ {
        expectedDate := today.AddDate(0, 0, -i)
        if logs[i].Date.Equal(expectedDate) {
            streak++
        } else {
            break
        }
    }

    c.JSON(http.StatusOK, gin.H{"streak": streak})
}