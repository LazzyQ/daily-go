// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"testing"

// 	"github.com/apache/pulsar-client-go/pulsar"
// )

// func TestConsumer(t *testing.T) {
// 	client, err := pulsar.NewClient(pulsar.ClientOptions{
// 		URL: "pulsar://localhost:6650",
// 	})
// 	if err != nil {
// 		log.Fatal("创建pulsar客户端err: ", err)
// 	}
// 	defer client.Close()

// 	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
// 		Topic:            "beats",
// 		SubscriptionName: "test-sub",
// 	})

// 	if err != nil {
// 		log.Fatal("创建pulsar客户端err: ", err)
// 	}

// 	defer consumer.Close()

// 	for {
// 		msg, err := consumer.Receive(context.Background())
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		fmt.Printf("Received message msgId: %#v -- content: '%s'\n",
// 			msg.ID(), string(msg.Payload()))

// 		consumer.Ack(msg)
// 	}

// }
