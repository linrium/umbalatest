package main

import "fmt"

type Manager struct {
	TotalFloors int
	Elevators   []*Elevator
}

func (this *Manager) Logger() {
	fmt.Println("-------LOGGER-------")
	for _, el := range this.Elevators {
		fmt.Printf("Elevator %v on floor %v\n", el.Id, el.Current)
	}
	fmt.Println("--------------------")
}

func (this *Manager) Operate() {
	for _, e := range this.Elevators {
		if e.IsOpen == false {
			e.pickUp()
			e.dropOff()
			e.move()
		}
	}
}

func (this *Manager) Schedule(req Request) {
	var suitable *Elevator = nil
	elevators := this.Elevators

	switch req.(type) {
	case Passenger:
		for _, el := range elevators {
			if suitable == nil {
				suitable = el
			}

			if el.computeNearest(req.(Passenger), this.TotalFloors) > suitable.computeNearest(req.(Passenger), this.TotalFloors) {
				suitable = el
			}
		}
		fmt.Printf("scheluded request %v, elevator %v, destination %v\n", req.(Passenger).Id, suitable.Id, req.(Passenger).Destination)
		suitable.add(req.(Passenger))
	}
}
