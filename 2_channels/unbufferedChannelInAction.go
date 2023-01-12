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
