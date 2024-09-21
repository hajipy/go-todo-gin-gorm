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

type NewTodo struct {
	Title string `json:"title" binding:"required"`
}

type UpdateTodo struct {
	IsCompleted *bool `json:"is_completed" binding:"required"`
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
		var newTodo NewTodo
		if err := context.ShouldBindJSON(&newTodo); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		todo := &Todo{Title: newTodo.Title}
		db.Create(todo)

		context.JSON(http.StatusCreated, todo)
	})

	server.PATCH("/todo/:id", func(context *gin.Context) {
		var updateTodo UpdateTodo
		if err := context.ShouldBindJSON(&updateTodo); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var todo Todo
		if err := db.First(&todo, context.Param("id")).Error; err != nil {
			context.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf("Todo with ID %s not found", context.Param("id")),
			})
			return
		}

		if *updateTodo.IsCompleted {
			now := time.Now()
			todo.CompletedAt = &now
		} else {
			todo.CompletedAt = nil
		}

		db.Save(&todo)

		context.JSON(http.StatusOK, todo)
	})

	server.DELETE("/todo/:id", func(context *gin.Context) {
		var todo Todo
		if err := db.First(&todo, context.Param("id")).Error; err != nil {
			context.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf("Todo with ID %s not found", context.Param("id")),
			})
			return
		}

		db.Delete(&todo)

		context.JSON(http.StatusNoContent, nil)
	})

	if err := server.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
