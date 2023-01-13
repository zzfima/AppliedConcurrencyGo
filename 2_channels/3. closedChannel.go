package main

import (
	"log"
	"time"
)

func mainClosedChannel() {
	ch := make(chan string)
	go doIt2(&ch)
	ch <- "abra"
	time.Sleep(time.Second)
	close(ch)
	time.Sleep(2 * time.Second)
	//Panic: send to closed channel:
	//ch <- "kadabra"
}

func doIt2(ch *chan string) {
	s, ok := <-*ch
	if !ok {
		log.Println("1 not ok")
	} else {
		log.Println(s)
	}

	s, ok = <-*ch
	if !ok {
		log.Println("2 not ok")
	} else {
		log.Println(s)
	}
}
