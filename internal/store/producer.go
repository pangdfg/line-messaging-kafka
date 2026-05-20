package store

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"

	"kafka-go/internal/database"
	models "kafka-go/internal/schema"
)


var writer *kafka.Writer

func CreateProduct(c *gin.Context) {

	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := database.DB.Create(&product).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something wrong",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func PlaceOrder(c *gin.Context) {

	type Request struct {
		ProductID uint   `json:"productId"`
		UserID    string `json:"userId"`
	}

	var req Request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var product models.Product

	err := database.DB.First(&product, req.ProductID).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "product not found",
		})
		return
	}

	if product.Amount <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "product out of stock",
		})
		return
	}

	product.Amount -= 1

	database.DB.Save(&product)

	order := models.Order{
		ProductID:   product.ID,
		UserLineUid: req.UserID,
		Status:      "pending",
	}

	database.DB.Create(&order)

	orderMessage := models.OrderMessage{
		ProductName: product.Name,
		UserID:      req.UserID,
		OrderID:     order.ID,
	}

	sendKafkaMessage(orderMessage)

	c.JSON(http.StatusOK, gin.H{
		"message": "buy product " + product.Name + " successful. waiting message for confirm.",
	})
}

func AddOrder(c *gin.Context) {

	type Request struct {
		ProductID uint   `json:"productId"`
		UserID    string `json:"userId"`
		Amount    int    `json:"amount"`
	}

	var req Request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var product models.Product

	err := database.DB.First(&product, req.ProductID).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "product not in data",
		})
		return
	}

	product.Amount += req.Amount

	database.DB.Save(&product)

	order := models.Order{
		ProductID:   product.ID,
		UserLineUid: req.UserID,
		Status:      "pending",
	}

	database.DB.Create(&order)

	orderMessage := models.OrderMessage{
		ProductName: product.Name,
		UserID:      req.UserID,
		OrderID:     order.ID,
	}

	sendKafkaMessage(orderMessage)

	c.JSON(http.StatusOK, gin.H{
		"message": "add product " + product.Name +
			" quantity " + strconv.Itoa(req.Amount) +
			" pieces successful. waiting message for confirm.",
	})
}

func sendKafkaMessage(data models.OrderMessage) {

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("marshal error:", err)
		return
	}

	err = writer.WriteMessages(
		context.Background(),
		kafka.Message{
			Value: jsonData,
		},
	)

	if err != nil {
		log.Println("kafka send error:", err)
		return
	}

	log.Println("message sent:", string(jsonData))
}