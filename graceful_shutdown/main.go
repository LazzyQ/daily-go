package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	name string
}

func NewWorker(name string) *Worker {
	return &Worker{name: name}
}

func (w *Worker) Name() string {
	return w.name
}

func (w *Worker) Report() {
	for counter := 0; counter < 5; counter++ {
		fmt.Printf("%s report number %v\n", w.Name(), counter)
		time.Sleep(3 * time.Second)
	}
}

func(w *Worker) Run(ctx context.Context) error {
	for {
		w.Report()
	}
}

func main() {
	ctx := context.Background()
	w := NewWorker("Alice")

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer  wg.Done()
		err := w.Run(ctx)
		fmt.Printf("%s stopped: error = %v", w.Name(), err)
	}()
	wg.Wait()
}
