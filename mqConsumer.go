package mq

import (
    "fmt"
    "log"
    "context"

    "modbusd/config"
    "github.com/streadway/amqp"
)

type RtuMessage struct{
    address string
}

func mqService(ctx context.Context, ch RtuMessage) {

    conn, err := amqp.Dial(config.RMQADDR)
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

    err = ch.QueueBind(
                "webq2",
                "webq2",
                "webex",
                false,
                nil,
            )
    failOnError(err, "Failed to bind exchange")

            
    for{
        select{
        case <-ctx.Done():
            log.Println("mqService got ctx.Done and exit.")
            return

        default:
            msgs, err := ch.Consume(
                "webq2",
                "",
                true, //Auto Ack
                false,
                false,
                false,
                nil,
            )

            if err != nil {
                log.Fatal(err)
            }

            for msg := range msgs {
                RtuMessage <- msg
                log.Printf("Consume a message: %s\n", msg.Body)
            }

        }
    }
            

    

    
}

func failOnError(err error, msg string) {
    if err != nil {
        fmt.Printf("%s: %s\n", msg, err)
    }
}