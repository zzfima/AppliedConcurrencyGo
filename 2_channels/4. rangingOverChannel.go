package main

import (
	"log"
	"strconv"
	"time"
)

func mainRangingOverChannel() {
	ch1 := make(chan string, 1)
	go doIt3(&ch1)

	for i := 0; i < 10; i++ {
		ch1 <- "i" + strconv.Itoa(i)
	}
	close(ch1)
	time.Sleep(3 * time.Second)
}

func doIt3(ch1 *chan string) {
	for {
		v, ok := <-*ch1

		if !ok {
			return
		}
		log.Println(v)
	}
}
