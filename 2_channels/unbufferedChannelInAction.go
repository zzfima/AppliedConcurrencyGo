// Unbuffered receiver: accept send ONLY if there appropriate receiver.
// If no receiver on channel, it will be wait on sending
// An unbuffered channel is used to perform synchronous communication between goroutines,
// while a buffered channel is used for perform asynchronous communication.

package main

import (
	"log"
	"time"
)

func mainUnbufferedChannelInAction() {
	ch := make(chan string)
	go doIt(ch)
	log.Println("start get value....")
	log.Println("... finished get value: " + <-ch)
}

func doIt(ch chan string) {
	log.Println("start put value....")
	ch <- "hello world"
	time.Sleep(3 * time.Second)
	log.Println("... finished put value")
}
