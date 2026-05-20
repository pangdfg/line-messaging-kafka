package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"bytes"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"

	"kafka-go/internal/database"
	models "kafka-go/internal/schema"
)

const lineAPIURL = "https://api.line.me/v2/bot/message/push"

func consumer() {

	godotenv.Load()

	database.Connect()

	lineAccessToken := os.Getenv("LINE_ACCESS_TOKEN")

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{
			"localhost:9092",
			"localhost:9093",
		},
		Topic:   "message-topic",
		GroupID: "message-group",
	})

	log.Println("Kafka consumer started")

	for {

		message, err := reader.ReadMessage(context.Background())

		if err != nil {
			log.Println("consumer error:", err)
			continue
		}

		var messageData models.OrderMessage

		err = json.Unmarshal(message.Value, &messageData)

		if err != nil {
			log.Println("json parse error:", err)
			continue
		}

		log.Println("=== consumer message", messageData)

		body := models.LineBody{
			To: messageData.UserID,
			Messages: []models.LineMessage{
				{
					Type: "text",
					Text: "Buy product: " + messageData.ProductName + " successful!",
				},
			},
		}

		jsonBody, _ := json.Marshal(body)

		req, err := http.NewRequest(
			"POST",
			lineAPIURL,
			bytes.NewBuffer(jsonBody),
		)

		if err != nil {
			log.Println("request error:", err)
			continue
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+lineAccessToken)

		client := &http.Client{}

		resp, err := client.Do(req)

		if err != nil {
			log.Println("LINE API error:", err)
			continue
		}

		resp.Body.Close()

		log.Println("LINE message success")

		// update order status
		err = database.DB.Model(&models.Order{}).
			Where("id = ?", messageData.OrderID).
			Update("status", "success").Error

		if err != nil {
			log.Println("update order error:", err)
			continue
		}

		log.Println("Order updated success")
	}
}