package main

import (
	"log"

	"simple-user-web/config"
	"simple-user-web/user"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	db, err := gorm.Open(
		postgres.Open("postgres://postgres:root@localhost:5432/postgres?sslmode=disable"),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.New(db)

	app := gin.Default()
	router := app.Group("/")
	user.RegisterRouter(router, cfg)

	if err := app.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
