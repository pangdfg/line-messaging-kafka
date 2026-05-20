# Line Messaging Kafka - Go Edition

## Overview

A Go-based Line Messaging API service that leverages Apache Kafka for asynchronous message processing and MySQL for data persistence. This project demonstrates building scalable messaging systems with real-time Line Bot integration, message queueing, and a RESTful API using the Gin framework.

## Features

- **REST API** - Gin-based HTTP endpoints for managing products and orders
- **Kafka Message Queue** - Asynchronous message processing with Kafka topics
- **Consumer Service** - Kafka consumer that pushes messages to Line users
- **Database Integration** - MySQL with GORM ORM for persistent storage
- **Line Bot Integration** - Send messages to Line users via Line Messaging API
- **Docker Support** - Easy deployment with Docker Compose (Kafka, Zookeeper, MySQL, phpMyAdmin)

## Requirements

- Go 1.25.4 or higher
- Docker & Docker Compose
- Line Developer Account (for LINE_ACCESS_TOKEN)
- Git

## Setup & Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/pangdfg/line-messaging-kafka.git
   cd line-messaging-kafka
   ```

2. **Set up environment variables:**

   Create a `.env` file in the project root:
   ```bash
   LINE_ACCESS_TOKEN=your_line_bot_token_here
   PORT=8080
   ```

3. **Start Docker services (Kafka, MySQL, phpMyAdmin):**

   ```bash
   docker-compose up -d
   ```

   This will start:
   - **Kafka & Zookeeper** - Message broker on port 9092
   - **MySQL** - Database on port 3306 (credentials: root/root)
   - **phpMyAdmin** - Database UI on http://localhost:8080

4. **Run the REST API server and Kafka consumer:**

   ```bash
   go run cmd/.
   ```

   The API will be available at `http://localhost:8080`

   The consumer will listen for messages on the Kafka topic and push them to Line users.

## API Endpoints

- `GET /` - Health check
- `POST /api/create-product` - Create a new product
- `PUT /api/placeorder` - Place a new order
- `PUT /api/addorder` - Add item to order

## Services

### Producer (cmd/main.go)
- Provides REST API endpoints
- Accepts product and order requests
- Publishes messages to Kafka topic `message-topic`

### Consumer (cmd/consumer.go)
- Listens to Kafka messages
- Sends push notifications to Line users
- Integrates with Line Messaging API

## Database

- **Host:** localhost:3306
- **Database:** defaultdb
- **User:** root
- **Password:** root
- **phpMyAdmin:** http://localhost:8080

## Dependencies

Key dependencies:
- `github.com/gin-gonic/gin` - Web framework
- `github.com/segmentio/kafka-go` - Kafka client
- `gorm.io/gorm` - ORM for database
- `gorm.io/driver/mysql` - MySQL driver
- `github.com/joho/godotenv` - Environment variable loading

## Stopping Services

To stop all Docker services:

```bash
docker-compose down
```