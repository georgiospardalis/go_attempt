package main

import (
	"github.com/georgiospardalis/beat_assignment/ride"
	"bufio"
	"encoding/csv"
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
			return
		}

		reader := csv.NewReader(bufio.NewReader(csvFile))
		line, err := reader.Read()

		var ridePositions []ride.Position

		previousId, lat, lng, ts, parseErr := parseLine(line)

		if len(parseErr) > 0 {
			log.Fatal(parseErr)

			return
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
				return
			}

			currentId, lat, lng, ts, parseErr := parseLine(line)

			if len(parseErr) > 0 {
				log.Fatal(parseErr)

				continue
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