package main

import (
	"log"
	"sync"
	"time"
)

func mainSignalingWorkFinished() {
	signalFinishGoroutine := make(chan struct{})
	signalFinishMain := make(chan struct{})
	go doWork(signalFinishGoroutine, signalFinishMain)
	time.Sleep(time.Second)
	signalFinishGoroutine <- struct{}{}
	time.Sleep(time.Second)
	<-signalFinishMain
	log.Println("Finish main")
}

func doWork(signalFinishGoroutine chan struct{}, signalFinishMain chan struct{}) {
	// Sync.Once make it happen once only
	var once sync.Once
	for {
		select {
		case <-signalFinishGoroutine:
			log.Println("Finish goroutine")
			time.Sleep(5 * time.Second)
			once.Do(func() {
				close(signalFinishGoroutine)
				signalFinishMain <- struct{}{}
			})
			return
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}
