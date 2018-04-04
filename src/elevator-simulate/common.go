package main

import (
	"time"
	"math/rand"
)

func RandFloors() int {
	rand.Seed(time.Now().UnixNano())
	destination := rand.Intn(TOTAL_FLOORS)

	return destination
}