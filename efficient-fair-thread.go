package main

import (
  "fmt"
  "math"
  "strconv"
  "sync"
  "sync/atomic"
  "time"
)

var totalPrimeNumbers int32 = 0
var currentNumber int32 = 0
var CONCURRENCY int = 10
var LIMIT int = 10000000

func checkPrime(x int32) {
  if x&1 == 0 {
    return
  }
  for i := 3; i <= int(math.Sqrt(float64(x))); i++ {
    if x%int32(i) == 0 {
      return
    }
  }

  atomic.AddInt32(&totalPrimeNumbers, 1)
}

func doWork(threadId string, wg *sync.WaitGroup) {
  defer wg.Done()
  start := time.Now()
  for {
    x := atomic.AddInt32(&currentNumber, 1)
    if x > int32(LIMIT) {
      break
    }
    checkPrime(x)
  }
  fmt.Printf("Thread %s took %s\n", threadId, time.Since(start))
}



func fairThreadDemo () {
  wg := sync.WaitGroup{}
  for i := 0; i < CONCURRENCY; i++ {
    wg.Add(1)
    go doWork(strconv.Itoa(i), &wg)
  }

  wg.Wait()

  fmt.Println("Counted till ", LIMIT, " found ", totalPrimeNumbers, " prime Numbers")
}