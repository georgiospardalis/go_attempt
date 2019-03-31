package main

import (
	"github.com/georgiospardalis/beat_assignment/pipeline"
	"log"
	"os"
	"runtime"
)

func main() {
	cpuCoresAvailable := runtime.NumCPU()
	channelBuffer := cpuCoresAvailable * 5

	if len(os.Args) < 3 {
		log.Fatal("Not enough arguments. Please, provide absolute file paths...\n" +
			"Usage: [application_executable] [input_file] [output_file]")
	}

	fileToRead := os.Args[1]
	fileToWrite := os.Args[2]

	readFileChannel := pipeline.ReadFromFile(channelBuffer, fileToRead)

	var processChannels []<-chan []string

	for i := 0; i < cpuCoresAvailable; i++ {
		processChannels = append(processChannels, pipeline.ProcessRide(channelBuffer, readFileChannel))
	}

	mergedProcessChannels := pipeline.Merge(channelBuffer*cpuCoresAvailable, processChannels)

	pipeline.WriteToFile(fileToWrite, mergedProcessChannels)
}