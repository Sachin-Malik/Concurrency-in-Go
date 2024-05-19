package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

type Item struct {
	id    int
	value string
}

type Queue struct {
	queue []Item
	mu    sync.Mutex
}

func (q *Queue) Enqueue(item Item) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.queue = append(q.queue, item)

}

func (q *Queue) Dequeue() Item {
	q.mu.Lock()
	if len(q.queue) == 0 {
		panic("No items to dequeue")
	}

	defer q.mu.Unlock()
	firstItem := q.queue[0]
	q.queue = q.queue[1:]
	return firstItem
}

func (q *Queue) Size() int {
	return len(q.queue)
}

func threadSafeDemo() {
	itemsCount := 1000000
	q := Queue{}

	wgE := sync.WaitGroup{}

	// Inserting a million Values
	for i := 0; i < itemsCount; i++ {
		wgE.Add(1)
		dummy := Item{id: rand.Intn(100000), value: strconv.Itoa(rand.Intn(9999999))}
		go func() {
			q.Enqueue(dummy)
			wgE.Done()
		}()
	}

	wgE.Wait()
	fmt.Printf("Size of queue is %d \n", q.Size())

	wgD := sync.WaitGroup{}
	for i := 0; i < itemsCount; i++ {
		wgD.Add(1)
		go func() {
			q.Dequeue()
			wgD.Done()
		}()
	}

	wgD.Wait()

	fmt.Printf("Size of queue is %d \n", q.Size())

}
