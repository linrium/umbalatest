package main

import (
	"sync"
	"time"
)

var (
	done = make(chan bool)
)

func main() {
	var wg sync.WaitGroup

	manager := InitManager(TOTAL_ELEVATORS)
	ticker := time.NewTicker(time.Duration(1) * time.Second)
	wg.Add(1)

	go Start(manager, ticker, &wg, done)

	requests := InitRequest(TOTAL_PASSENGER_REQUESTS)
	go SendRequests(manager, requests)

	wg.Wait()
}

func InitManager(totalElevator int) *Manager {
	elevators := []*Elevator{}

	for i := 0; i < totalElevator; i++ {
		elevator := &Elevator{
			Id:              i,
			Current:         DEFAULT_FLOOR,
			Direction:       direction.UP,
			PickUpRequests:  [TOTAL_FLOORS][]Passenger{},
			DropOffRequests: [TOTAL_FLOORS][]Passenger{},
			IsOpen:          true,
		}

		elevators = append(elevators, elevator)
	}

	return &Manager{TotalFloors: TOTAL_FLOORS, Elevators: elevators}
}

func InitRequest(totalRequests int) chan Request {

	var passengerRequests = make(chan Request, totalRequests)

	for i := 0; i < totalRequests; i++ {
		destination := RandFloors()

		passengerRequests <- Passenger{
			Id:          i,
			Current:     0,
			Destination: destination,
			Direction:   direction.UP,
		}
	}
	close(passengerRequests)

	return passengerRequests
}
