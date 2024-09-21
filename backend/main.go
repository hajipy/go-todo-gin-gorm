package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DBHost     string `env:"TODO_DB_HOST" envDefault:"127.0.0.1"`
	DBPort     int    `env:"TODO_DB_PORT" envDefault:"33306"`
	DBUser     string `env:"TODO_DB_USER" envDefault:"todo"`
	DBPassword string `env:"TODO_DB_PASSWORD" envDefault:"todo"`
	DBName     string `env:"TODO_DB_NAME" envDefault:"todo"`
}

type Todo struct {
	gorm.Model
	Title       string
	CompletedAt *time.Time
}

func main() {
	config := &Config{}
	if err := env.Parse(config); err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}

	dbDsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	// connect mysql with gorm
	db, err := gorm.Open(mysql.Open(dbDsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(&Todo{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	server := gin.Default()

	server.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	server.POST("/todo", func(context *gin.Context) {
		todo := &Todo{Title: "todo"}
		db.Create(todo)

		context.JSON(http.StatusCreated, todo)
	})

	if err := server.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
