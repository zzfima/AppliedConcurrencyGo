/*
example of fatal error: concurrent map writes
run code with flag: "go run -race .\raceConditions.go"
Analyzer will output:
==================
WARNING: DATA RACE
Write at 0x00c000027e00 by goroutine 7:
  runtime.mapassign_fast64()
      C:/Program Files/Go/src/runtime/map_fast64.go:93 +0x0
  main.updateMap()
      C:/Users/efzabar/source/repos/AppliedConcurrencyGo/1_goroutines/raceConditions.go:10 +0x4b
  main.main.func1()
      C:/Users/efzabar/source/repos/AppliedConcurrencyGo/1_goroutines/raceConditions.go:15 +0x47

Previous write at 0x00c000027e00 by goroutine 6:
  runtime.mapassign_fast64()
      C:/Program Files/Go/src/runtime/map_fast64.go:93 +0x0
  main.updateMap()
      C:/Users/efzabar/source/repos/AppliedConcurrencyGo/1_goroutines/raceConditions.go:10 +0x4b
  main.main.func1()
      C:/Users/efzabar/source/repos/AppliedConcurrencyGo/1_goroutines/raceConditions.go:15 +0x47

Goroutine 7 (running) created at:
  main.main()
      C:/Users/efzabar/source/repos/AppliedConcurrencyGo/1_goroutines/raceConditions.go:15 +0x95

Goroutine 6 (finished) created at:
  main.main()
      C:/Users/efzabar/source/repos/AppliedConcurrencyGo/1_goroutines/raceConditions.go:15 +0x95
==================
Found 1 data race(s)
exit status 66
*/

package main

import (
	"time"
)

var testMap = make(map[int]int)

func updateMap(key int, val int) {
	testMap[key] = val
}

func mainRaceConditions() {
	for i := 0; i < 100; i++ {
		go updateMap(i, i+100)
	}

	time.Sleep(time.Second)
}
