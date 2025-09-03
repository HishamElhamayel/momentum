package handlers

import (
	"fmt"
	"net/http"

	"example.com/momentum/internal/models"
	"example.com/momentum/pkg"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler struct {
	DB *gorm.DB
}

func (h *AuthHandler) Register(c *gin.Context){
	var input struct{
		Name string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}	

	hashedPassword, err:= pkg.HashPassword(input.Password)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to hash password"})
		return
	}

	user:= models.User{
		Name: input.Name,
		Email: input.Email,
		Password: hashedPassword,
	}
	
	
	result := h.DB.Create(&user)
	if result.Error != nil {
		fmt.Printf("Failed to create user: %v\n", result.Error)
		fmt.Printf("Error details: %+v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}


func (h *AuthHandler) Login(c *gin.Context){
	var input struct{
		Email string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	var user models.User
	result := h.DB.Where("email = ?", input.Email).First(&user)
	if result.Error != nil{
		fmt.Printf("Failed to find user: %v\n", result.Error)
		fmt.Printf("Error details: %+v\n", result.Error)
		c.JSON(http.StatusNotFound, gin.H{"Error": "User not found"})
		return
	}

	if !pkg.ComparePassword(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid password"})
		return
	}

	token,err := pkg.GenerateToken(user.ID)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}