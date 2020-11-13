package main

import (
	"fmt"
)

func main() {
	fmt.Println(`===== CHANNEL ASYNC =====`)

	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hello golang !"

	select {
	case messages <- msg:
		fmt.Println(`sent message`, msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println(`received message`, msg)
	case sig := <-signals:
		fmt.Println(`recerived signal`, sig)
	default:
		fmt.Println(`nothing.`)
	}
}
