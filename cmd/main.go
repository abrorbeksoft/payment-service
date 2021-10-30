package main

import (
	"fmt"
	"log"
	"github.com/streadway/amqp"
)

func main()  {
	conn, err:=amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Error while connecting")
	channel, err:=conn.Channel()

	q,err:=channel.QueueDeclare("hello", false, false, false, false, nil)
	failOnError(err, "Error while creating queue")

	msg,err:=channel.Consume(q.Name, "", true, false, false, false, nil)
	failOnError(err, "Error while consuming")
	func() {
		for d := range msg {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	fmt.Println("Hello world")
}

func failOnError(err error, message string)  {
	if err!= nil {
		log.Fatal(message)
	}
}