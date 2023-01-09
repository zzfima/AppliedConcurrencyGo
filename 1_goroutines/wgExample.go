package main

import (
	"log"
	"sync"
)

func mainWgExample() {
	var wg sync.WaitGroup
	wg.Add(1)
	log.Printf("wg address: %p", &wg)
	go printHello(&wg)
	wg.Wait()
	printGoodBye()
}

func printHello(wg *sync.WaitGroup) {
	log.Printf("wg address: %p", wg)
	defer wg.Done()
	log.Println("Hello!")
}

func printGoodBye() {
	log.Println("Good Bye!")
}
