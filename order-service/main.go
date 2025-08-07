package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/RealNameHidden/kafka-pizza-tracker/common"

	"github.com/segmentio/kafka-go"
)

func main() {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "pizza-orders",
	})

	pizzas := []string{"Margherita", "Pepperoni", "Hawaiian"}
	sizes := []string{"Small", "Medium", "Large"}

	for i := 0; i < 5; i++ {
		order := common.PizzaOrder{
			OrderID: time.Now().Format("150405") + "-" + string(i+'0'),
			User:    "User" + string(i+'A'),
			Pizza:   pizzas[rand.Intn(len(pizzas))],
			Size:    sizes[rand.Intn(len(sizes))],
		}

		msgBytes, _ := json.Marshal(order)
		err := writer.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(order.OrderID),
				Value: msgBytes,
			},
		)
		if err != nil {
			log.Println("Error writing message:", err)
		} else {
			log.Printf("Order sent: %+v\n", order)
		}

		time.Sleep(2 * time.Second)
	}

	writer.Close()
}
