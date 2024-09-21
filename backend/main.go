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
	DBHost     string `env:"DB_HOST" envDefault:"127.0.0.1"`
	DBPort     int    `env:"DB_PORT" envDefault:"33306"`
	DBUser     string `env:"DB_USER" envDefault:"todo"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"todo"`
	DBName     string `env:"DB_NAME" envDefault:"todo"`
}

type Todo struct {
	gorm.Model
	Title       string
	CompletedAt *time.Time
}

func connectDB(dbDsn string, maxRetries int) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	waitTime := 1 * time.Second

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(mysql.Open(dbDsn), &gorm.Config{})
		if err == nil {
			return db, nil
		}

		log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)

		time.Sleep(waitTime)
		waitTime *= 2
	}

	return nil, err
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

	maxRetries := 3
	db, err := connectDB(dbDsn, maxRetries)
	if err != nil {
		log.Fatalf("Failed to connect to database after %d attempts: %v", maxRetries, err)
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

	server.GET("/todo", func(context *gin.Context) {
		var todos []Todo
		db.Order("created_at asc").Find(&todos)

		var response []gin.H
		for _, todo := range todos {
			response = append(response, gin.H{
				"id":           todo.ID,
				"title":        todo.Title,
				"completed_at": todo.CompletedAt,
			})
		}

		context.JSON(http.StatusOK, response)
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
