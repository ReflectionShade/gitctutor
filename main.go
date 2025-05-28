package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"go.uber.org/zap"
)

// thirty edit from github.....

func main() {
	// loggers
	logger, err := zap.NewDevelopment()
	sugar := logger.Sugar()
	//

	pubSub := gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(false, false),
	)

	messages, err := pubSub.Subscribe(context.Background(), "example.topic")
	if err != nil {
		sugar.Info(err)
	}

	go process(messages)

	publishMessages(pubSub)
}

// my second edit from the web

func publishMessages(publisher message.Publisher) {
	for {
		msg := message.NewMessage(watermill.NewUUID(), []byte("Hello world"))

		if err := publisher.Publish("example.topic", msg); err != nil {
			log.Fatal(err)
		}

		time.Sleep(time.Second * 5)
	}
}

func process(messages <-chan *message.Message) {
	for msg := range messages {
		fmt.Printf("received message: %s, payload: %s\n", msg.UUID, string(msg.Payload))

		msg.Ack()
	}
}

func helloGit(onTheLevel string) {
	fmt.Println(onTheLevel)
}

// my small update from website 
