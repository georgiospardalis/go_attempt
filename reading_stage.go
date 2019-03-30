package main

import (
	"bufio"
	"encoding/csv"
	"github.com/georgiospardalis/beat_assignment/ride"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFromFile(channelBuffer int, absoluteFilePath string) <-chan ride.Ride {
	out := make(chan ride.Ride, channelBuffer)

	go func() {
		defer close(out)

		csvFile, err := os.Open(absoluteFilePath)

		if err != nil {
			log.Fatal(err)
		}

		reader := csv.NewReader(bufio.NewReader(csvFile))
		line, err := reader.Read()

		var ridePositions []ride.Position

		previousId, lat, lng, ts, parseErr := parseLine(line)

		if len(parseErr) > 0 {
			log.Fatal(parseErr)
		}

		ridePositions = append(ridePositions, ride.NewPosition(lat, lng, ts))

		for {
			line, err := reader.Read()

			if err == io.EOF {
				if ridePositions != nil {
					cabRide := ride.NewRide(previousId, ridePositions)
					out <- cabRide
				}

				break
			} else if err != nil {
				log.Fatal(err)
			}

			currentId, lat, lng, ts, parseErr := parseLine(line)

			if len(parseErr) > 0 {
				log.Fatal(parseErr)
			}

			if currentId != previousId {
				cabRide := ride.NewRide(previousId, ridePositions)

				out <- cabRide

				previousId = currentId
				ridePositions = nil
			}

			ridePositions = append(ridePositions, ride.NewPosition(lat, lng, ts))
		}
	}()

	return out
}

func parseLine(line []string) (id int, lat, lng float64, ts int64, errors string){
	var errs []string

	id, err := strconv.Atoi(line[0])

	if err != nil {
		errs = append(errs, err.Error())
	}

	lat, err = strconv.ParseFloat(line[1], 64)

	if err != nil {
		errs = append(errs, err.Error())
	}

	lng, err = strconv.ParseFloat(line[2], 64)

	if err != nil {
		errs = append(errs, err.Error())
	}

	ts, err = strconv.ParseInt(line[3], 10, 64)
	if err != nil {
		errs = append(errs, err.Error())
	}

	return id, lat, lng, ts, strings.Join(errs, "\n")
}