package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

func main() {
	topic := "comments"
	worker, err := ConnectConsumer([]string{"localhost:29092"})

	if err != nil {
		log.Println("Error connect consumer: ", err)
		panic(err)
	}

	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest) 

	if(err != nil){
		log.Println("Error consume partition: ", err)
		panic(err)
	}

	fmt.Println("Worker is running")
	
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	msgCnt := 0

	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				log.Println("Consumer error: ", err)
			case msg := <-consumer.Messages():
				msgCnt++
				fmt.Println("Received message count: ", msgCnt, " topic is: ", msg.Topic, " message is: ", string(msg.Value))
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("Processed ", msgCnt, " messages")
	fmt.Println("Worker is stopped")

	if err := worker.Close(); err != nil {
		log.Println("Error close consumer: ", err)
		panic(err)
	}
}

func ConnectConsumer(brokersUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	conn, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		log.Println("Error consumer: ", err)
		return nil, err
	}

	return conn, nil
}