package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"

	"kafka-go/internal/database"
	"kafka-go/internal/store"
)

var writer *kafka.Writer

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found")
	}

	database.Connect()

	writer = &kafka.Writer{
		Addr: kafka.TCP(
			"localhost:9092",
			"localhost:9093",
		),
		Topic:    "message-topic",
		Balancer: &kafka.LeastBytes{},
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Kafka Go API running",
		})
	})

	r.POST("/api/create-product", store.CreateProduct)
	r.PUT("/api/placeorder", store.PlaceOrder)
	r.PUT("/api/addorder", store.AddOrder)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on :%s\n", port)

	r.Run(":" + port)
}