
package main

import (
    "context"
    "encoding/json"
    "log"

    "github.com/segmentio/kafka-go"
    "kafka-pizza-tracker/common"
)

func main() {
    reader := kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{"localhost:9092"},
        Topic:   "pizza-ready",
        GroupID: "delivery-group",
    })

    for {
        msg, err := reader.ReadMessage(context.Background())
        if err != nil {
            log.Println("Error reading message:", err)
            continue
        }

        var order common.PizzaOrder
        json.Unmarshal(msg.Value, &order)

        log.Printf("Delivered to %s: %s (%s)\n", order.User, order.Pizza, order.Size)
    }
}
