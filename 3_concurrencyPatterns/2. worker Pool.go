package main

import (
	"log"
	"time"
)

func mainWorkerPool() {
	numOfJobs := 5
	jobs := make(chan int, numOfJobs)
	results := make(chan int, numOfJobs)

	defer close(jobs)
	defer close(results)

	for i := 0; i < 3; i++ {
		go workerSingle(i*10, jobs, results)
	}

	time.Sleep(time.Second)

	for i := 0; i < 3; i++ {
		log.Println("Push job! i = ", i)
		jobs <- i
	}

	for i := 0; i < 3; i++ {
		log.Println("Got result! i = ", <-results)
	}
}

func workerSingle(id int, jobs chan int, results chan int) {
	log.Println(id, " worker single ready!")
	i := <-jobs
	log.Println(id, " worker single got job to do! Payload = ", i)
	time.Sleep(time.Second)
	results <- i + 100
}

func workerRange(id int, jobs chan int, results chan int) {
	log.Println(id, " worker range ready!")
	for i := range jobs {
		log.Println(id, " worker range got job to do! Payload = ", i)
		time.Sleep(time.Second)
		results <- i + 100
	}
}
