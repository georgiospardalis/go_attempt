package main

import (
	"github.com/georgiospardalis/beat_assignment/ride"
	"strconv"
)

func processRide(channelBuffer int, input <-chan ride.Ride) <-chan []string {
	out := make(chan []string, channelBuffer)

	go func() {
		for cabRide := range input {
			cabRide.PrepareSegments()
			id := strconv.Itoa(cabRide.Id())
			cost := strconv.FormatFloat(cabRide.EstimateFare(), 'f', 2, 64)

			out <- []string{id, cost}
		}

		close(out)
	}()

	return out
}