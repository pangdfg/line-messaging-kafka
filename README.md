# Line Messaging Kafka

## Overview

This project is a demonstration of building a Line Messaging API using Kafka as the messaging backbone. It provides a simple setup to send and receive Line messages using Kafka topics for communication.

## Features

- Send messages to Line users.
- Receive messages from Line users.
- Utilizes Kafka for message queueing and processing.

## Requirements

- Docker
- Node.js
- Line Developer Account
- Kafka

## Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/pangdfg/line-messaging-kafka
   ```

2. Navigate to the project directory:

   ```bash
   cd line-messaging-kafka
   ```

3. Set up your Line Messaging API credentials:
    - example in `.env.example` Add the following lines to the `.env` file
   ```bash
   LINE_ACCESS_TOKEN =
   PORT =
   ```

4. Install dependencies:
   ```bash
   npm install
   ```

5. Start the application:
    - database kafka phpmyadmin
   ```bash
   docker-compose up -d
   ```
    - run index.html
   ```bash
   npm start
   ```
    - run producer.js
   ```bash
   npx nodemon producer.js
   ```
