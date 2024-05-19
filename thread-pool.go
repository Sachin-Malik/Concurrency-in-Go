package main

import (
	"fmt"
	"sync"
	"time"
)

type Job func()

type Pool struct {
	workQueue chan Job
	wg        sync.WaitGroup
}

func (p *Pool) addJob(job Job) {
	p.workQueue <- job
}

func (p *Pool) Wait() {
	close(p.workQueue)
	p.wg.Wait()
}

func createPool(workerCount int, bufferSize int) *Pool {
	pool := &Pool{workQueue: make(chan Job, bufferSize)}
	pool.wg.Add(workerCount)

	for i := 0; i < workerCount; i++ {
		go func() {
			defer pool.wg.Done()
			// this for range 
			for job := range pool.workQueue {
				fmt.Println("Started working on JOB")
				job()
			}
		}()
	}
	return pool
}

func dummyJob() {
	time.Sleep(time.Second * 1)
	fmt.Println("Job Completed")
}

func threadPoolDemo() {
	// create a thread pool with 2 workers and a buffer size of 10
	pool := createPool(2, 10)

	for i := 0; i < 30; i++ {
		fmt.Println("Job Added")
		pool.addJob(dummyJob)
	}

	pool.Wait()
}
