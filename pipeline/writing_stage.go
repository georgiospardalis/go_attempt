package pipeline

import (
	"encoding/csv"
	"log"
	"os"
)

func WriteToFile(outputFile string, in <-chan []string) {
	file, err := os.Create(outputFile)

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for line := range in {
		err := writer.Write(line)

		if err != nil {
			log.Fatal(err)
		}
	}
}
