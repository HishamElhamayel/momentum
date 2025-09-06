package handlers

import (
	"fmt"
	"net/http"

	"example.com/momentum/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


type HabitHandler struct {
	DB *gorm.DB
}

func (h *HabitHandler) CreateHabit(c *gin.Context){

	userID,_:= c.Get("userID")

	var input struct{
		Title string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	
	habit := models.Habit{
		Title: input.Title,
		Description: input.Description,
		UserID: userID.(uint),
	}


	result:= h.DB.Create(&habit)

	if result.Error != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to create habit"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Habit created successfully"})
}

func (h *HabitHandler) GetHabits(c *gin.Context){
	userID,_:= c.Get("userID")
	var habits []models.Habit

	if err := h.DB.Where("user_id = ?", userID).Find(&habits).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to get habits"})
		return
	}

	c.JSON(http.StatusOK, habits)

}

func (h *HabitHandler) UpdateHabit(c *gin.Context){
	userID,_ := c.Get("userID")
	id := c.Param("id")

	var habit models.Habit
	result := h.DB.Where("id=? AND user_id=?",id,userID).First(&habit); 
	if result.Error != nil{
		c.JSON(http.StatusNotFound, gin.H{"Error": "Couldn't find habit"})
		return
	}

	fmt.Println(result.Error)

	var input struct {
		Title string `json:"title"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	habit.Title = input.Title
	habit.Description = input.Description

	h.DB.Save(&habit)
	c.JSON(http.StatusOK, habit)
}

func (h *HabitHandler) DeleteHabit(c *gin.Context){
	userID, _ := c.Get("userID")
	id := c.Param("id")

	result := h.DB.Where("id=? AND user_id=?",id,userID).Delete(&models.Habit{}); 

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete habit"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Habit not found"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})

}