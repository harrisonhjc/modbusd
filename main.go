package main

import (
    "fmt"
    "log"

    "modbusd/config"
    "github.com/streadway/amqp"
)

func main() {

    conn, err := amqp.Dial(config.RMQADDR)
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    forever := make(chan struct{})

    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

    for routine := 0; routine < config.CONSUMERCNT; routine++ {
        go func(routineNum int) {
            err = ch.QueueBind(
                "webq2",
                "webq2",
                "webex",
                false,
                nil,
            )
            failOnError(err, "Failed to bind exchange")

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
                log.Printf("In %d consume a message: %s\n", routineNum, msg.Body)
            }

        }(routine)
    }

    <-forever
}

func failOnError(err error, msg string) {
    if err != nil {
        fmt.Printf("%s: %s\n", msg, err)
    }
}