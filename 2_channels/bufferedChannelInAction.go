// Buffered receiver: accept send not matter if there appropriate receiver.
// If no receiver on channel, it will be do sending and continue
// An unbuffered channel is used to perform synchronous communication between goroutines,
// while a buffered channel is used for perform asynchronous communication.

package main

import (
	"log"
	"time"
)

func mainBufferedChannelInAction() {
	ch := make(chan string, 2)
	go doIt1(ch)
	log.Println("start get value....")
	log.Println("... finished get value: " + <-ch)
	time.Sleep(3 * time.Second)
}

func doIt1(ch chan string) {
	log.Println("start put value....")
	ch <- "hello world"
	time.Sleep(3 * time.Second)
	ch <- "goodbye world"
	log.Println("... finished put value")
}
