package main

import (
	"fmt"
	"math"
)

type Elevator struct {
	Id              int
	Current         int
	Direction       string
	PickUpRequests  [TOTAL_FLOORS][]Passenger
	DropOffRequests [TOTAL_FLOORS][]Passenger
	IsOpen          bool
}

func (this *Elevator) pickUp() {
	requests := this.PickUpRequests[this.Current]

	if len(requests) > 0 {
		fmt.Printf("picked up request = ")
		for _, req := range requests {
			this.DropOffRequests[req.Destination] = append(this.DropOffRequests[req.Destination], req)
			fmt.Printf("%v", req.Id)
		}
		fmt.Printf(", elevator = %v\n", this.Id)
	}

	this.PickUpRequests[this.Current] = []Passenger{}
}

func (this *Elevator) dropOff() {
	requests := this.DropOffRequests[this.Current]

	if !this.canOpen() {
		this.IsOpen = true
	}

	if len(requests) > 0 {
		fmt.Print("dropped off request = ")
		for _, req := range requests {
			fmt.Printf("%v", req.Id)
		}
		fmt.Printf(", floor = %v\n", this.Current)
	}

	this.DropOffRequests[this.Current] = []Passenger{}
}

func (this *Elevator) canOpen() bool {
	for _, r := range this.PickUpRequests {
		if len(r) > 0 {
			return true
		}
	}

	for _, r := range this.DropOffRequests {
		if len(r) > 0 {
			return true
		}
	}

	return false
}

func (this *Elevator) move() {
	if this.IsOpen {
		return
	}

	if this.Direction == direction.UP {
		this.Current++
		if this.Current >= this.max() {
			this.Direction = direction.DOWN
		}
		return
	}

	this.Current--
	if this.Current <= this.min() {
		this.Direction = direction.UP
	}
}

func (this *Elevator) min() int {
	pickUpMin := 0
	for i := 0; i < len(this.PickUpRequests); i++ {
		if len(this.PickUpRequests[i]) > 0 {
			pickUpMin = i
			break
		}
	}

	dropOffMin := 0
	for i := 0; i < len(this.DropOffRequests); i++ {
		if len(this.DropOffRequests[i]) > 0 {
			dropOffMin = i
			break
		}
	}

	if pickUpMin < dropOffMin {
		return pickUpMin
	}

	return dropOffMin
}

func (this *Elevator) max() int {
	pickUpMax := 0
	for i := len(this.PickUpRequests) - 1; i >= 0; i-- {
		if len(this.PickUpRequests[i]) > 0 {
			pickUpMax = i
			break
		}
	}

	dropOffMax := 0
	for i := len(this.DropOffRequests) - 1; i >= 0; i-- {
		if len(this.DropOffRequests[i]) > 0 {
			dropOffMax = i
			break
		}
	}

	if pickUpMax > dropOffMax {
		return pickUpMax
	}

	return dropOffMax
}

func (this *Elevator) computeNearest(passenger Passenger, totalFloors int) float64 {
	n := float64(totalFloors - 1)
	d := math.Abs(float64(this.Current - passenger.Current))

	fs := float64(1)

	t :=
		(passenger.Current >= this.Current && this.Direction == direction.UP) ||
			(passenger.Current <= this.Current && this.Direction == direction.DOWN)

	if t {
		if this.Direction == passenger.Direction {
			fs = n + 2 - d
		} else {
			fs = n + 1 - d
		}
	}

	return fs
}

func (this *Elevator) add(passenger Passenger) {
	this.IsOpen = false

	this.PickUpRequests[passenger.Current] = append(this.PickUpRequests[passenger.Current], passenger)
}