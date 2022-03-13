package messaging

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

type Event struct {
	Event string
	Key   string
	Team1 string
	Team2 string
}

func Send(c *gin.Context, topic string, data string) {
	// create a new context
	ctx := context.Background()
	go produce(ctx, topic, data)
}

func produce(ctx context.Context, topic string, data string) {

	l := log.New(os.Stdout, "kafka writer: ", 0)
	// intialize the writer with the broker addresses, and the topic

	brokerAddress := strings.Split(os.Getenv("BROKER_ADDRES"), ",")
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokerAddress,
		Topic:   topic,
		// assign the logger to the writer
		Logger: l,
	})

	err := w.WriteMessages(ctx, kafka.Message{
		Value: []byte(data),
	})

	if err != nil {
		panic("could not write message " + err.Error())
	}

	// after receiving the message, log its value
	fmt.Println("sent: ", data)
}
