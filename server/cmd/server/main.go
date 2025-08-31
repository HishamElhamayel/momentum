package main

import (
	"log"

	"example.com/momentum/internal/config"
	"example.com/momentum/internal/database"
	"example.com/momentum/internal/router"
)


func main(){
	cnfg := config.Load()

	db,err := database.Connect(cnfg.DBUrl)

	if err != nil{
		log.Fatal("Failed to connect to database: ", err)
	}

	r := router.Setup(db)

	log.Println("🚀 Server running on port", cnfg.Port)
	r.Run(":" + cnfg.Port)
}