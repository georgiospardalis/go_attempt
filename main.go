package main

import (
	"encoding/csv"
	"log"
	"os"
	"runtime"
)

func main() {
	cpuCoresAvailable := runtime.NumCPU()
	channelBuffer := cpuCoresAvailable * 5

	if len(os.Args) < 3 {
		log.Fatal("Not enough command line arguments provided...")
		return
	}

	fileToRead := os.Args[1]
	fileToWrite := os.Args[2]

	readFileChannel := readFromFile(channelBuffer, fileToRead)

	var processChannels []<-chan []string

	for i := 0; i < cpuCoresAvailable; i++ {
		processChannels = append(processChannels, processRide(channelBuffer, readFileChannel))
	}

	mergedProcessChannels := merge(channelBuffer * cpuCoresAvailable, processChannels)

	file, err := os.Create(fileToWrite)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for line := range mergedProcessChannels {
		err := writer.Write(line)

		if err != nil {
			log.Fatal(err)
		}
	}
}