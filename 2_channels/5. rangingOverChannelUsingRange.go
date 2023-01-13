package main

import (
	"log"
	"strconv"
	"time"
)

func mainRangingOverChannelUsingRange() {
	ch1 := make(chan string, 1)
	go doIt4(&ch1)

	for i := 0; i < 10; i++ {
		ch1 <- "i" + strconv.Itoa(i)
	}
	close(ch1)
	time.Sleep(3 * time.Second)
}

func doIt4(ch1 *chan string) {
	for v := range *ch1 {
		log.Println(v)
	}
}
