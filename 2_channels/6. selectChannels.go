package main

import (
	"log"
	"time"
)

func main() {
	hamburgerChannel := make(chan string, 4)
	pizzaChannel := make(chan string, 4)

	go doHamburger(&hamburgerChannel)
	go doPizza(&pizzaChannel)

	for {
		select {
		case hamburgerChannelString, ok := <-hamburgerChannel:
			if !ok {
				hamburgerChannel = nil
				return
			}
			log.Println(hamburgerChannelString)
		case pizzaChannelString, ok := <-pizzaChannel:
			if !ok {
				pizzaChannel = nil
				return
			}
			log.Println(pizzaChannelString)
		}
	}
}

func doHamburger(ch *chan string) {
	*ch <- "h1"
	*ch <- "h2"
	*ch <- "h3"

	time.Sleep(100 * time.Millisecond)
	close(*ch)
}

func doPizza(ch *chan string) {
	*ch <- "pz1"
	*ch <- "pz2"

	time.Sleep(100 * time.Millisecond)
	close(*ch)
}
