/*
example of fix fatal error: concurrent map writes
*/
package main

import (
	"log"
	"sync"
)

var concurrencyMap sync.Map
var wg sync.WaitGroup

func updateMap1(key int, val int) {
	defer wg.Done()
	concurrencyMap.Store(key, val)
}

func mainRaceConditionsFixedSyncMap() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go updateMap1(i, i+100)
	}

	wg.Wait()

	v, o := concurrencyMap.Load(88)
	log.Println(v, o)
}
