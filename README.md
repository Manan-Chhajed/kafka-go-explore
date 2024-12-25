# Apache Kafka Producer-Consumer with Go

This project demonstrates a Kafka Producer-Consumer system built with Go. It includes a REST API for sending messages to a Kafka producer, which are consumed and processed by a Kafka consumer.

## Features
- **Kafka Integration**: Uses Apache Kafka for message streaming.
- **REST API**: Producer exposes a simple HTTP API using the Go Fiber framework.
- **Asynchronous Processing**: Consumer processes messages asynchronously.
- **Dockerized Setup**: Kafka and Zookeeper run via Docker Compose.

## Technologies
- **Go**: Programming language used for the producer and consumer.
- **Apache Kafka**: Event streaming platform.
- **Zookeeper**: Used for coordinating Kafka brokers.
- **Fiber**: Fast web framework for Go.
- **Sarama**: Go client for Kafka.

## Setup and Installation

1. **Clone the repository**:
    ```bash
    git clone https://github.com/Manan-Chhajed/kafka-go-explore.git
    ```
    ```bash
    cd kafka-go-explore
    ```
2. **Start Kafka and Zookeeper**:
    ```bash
    docker-compose up -d
    ```
3. **Install Go dependencies**:
    ```bash
    go mod tidy
    ```
4. **Run the Producer**:
    ```bash
    go run producer/producer.go
    ```
5. **Run the Consumer**:
    ```bash
    go run worker/worker.go
    ```

## Usage

1. **Send a message to the Kafka producer** using the following `curl` command:

    ```bash
    curl --location --request POST 'http://127.0.0.1:3000/api/v1/comments' --header 'Content-Type: application/json' --data-raw '{ "text": "message 1" }'
    ```

    ```bash
    curl --location --request POST 'http://127.0.0.1:3000/api/v1/comments' --header 'Content-Type: application/json' --data-raw '{ "text": "message 2" }'
    ```

2. **Test Consumer Resilience**:
   - Stop the consumer service.
   - Send another message using the producer.
   - Restart the consumer. The message should still be processed, demonstrating the Kafka consumer's ability to handle message persistence and recovery after restarts.
