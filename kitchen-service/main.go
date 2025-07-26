
package main

import (
    "context"
    "encoding/json"
    "log"
    "time"

    "github.com/segmentio/kafka-go"
    "kafka-pizza-tracker/common"
)

func main() {
    reader := kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{"localhost:9092"},
        Topic:   "pizza-orders",
        GroupID: "kitchen-group",
    })

    writer := kafka.NewWriter(kafka.WriterConfig{
        Brokers: []string{"localhost:9092"},
        Topic:   "pizza-ready",
    })

    for {
        msg, err := reader.ReadMessage(context.Background())
        if err != nil {
            log.Println("Error reading message:", err)
            continue
        }

        var order common.PizzaOrder
        json.Unmarshal(msg.Value, &order)

        log.Printf("Kitchen received order: %+v\n", order)
        time.Sleep(3 * time.Second) // simulate cooking

        err = writer.WriteMessages(context.Background(),
            kafka.Message{
                Key:   []byte(order.OrderID),
                Value: msg.Value,
            },
        )
        if err != nil {
            log.Println("Error writing ready message:", err)
        } else {
            log.Printf("Pizza ready: %+v\n", order)
        }
    }
}
