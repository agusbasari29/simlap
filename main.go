package main

import (
	"os"

	"github.com/agusbasari29/simlap.git/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = database.SetupDBConnection()
)

func main() {

	g := gin.Default()
	g.Run(os.Getenv("SERVER_PORT"))
}
