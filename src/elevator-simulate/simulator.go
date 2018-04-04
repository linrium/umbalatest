package main

import (
	"time"
	"sync"
	"os"
	"fmt"
)

type Run interface {
	Step()
	Logger()
}

type Scheduler interface {
	Schedule(req Request)
}

type Request interface {
}

func Start(run Run, ticker *time.Ticker, wg *sync.WaitGroup, done chan bool) {
	defer wg.Done()

	for {
		select {
		case <-ticker.C:
			run.Logger()
			run.Step()
		case <-done:
			fmt.Printf("Helloooooo")
			ticker.Stop()
			os.Exit(-1)
			return
		}
	}
}

func SendRequests(scheduler Scheduler, requests <- chan Request) {
	for req := range requests {
		scheduler.Schedule(req)
		time.Sleep(time.Duration(3) * time.Second)
	}
}
