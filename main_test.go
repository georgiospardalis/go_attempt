package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestPipelineEndToEnd(t *testing.T) {
	absInputFile, _ := filepath.Abs("./test_resources/dummyInput.csv")
	absOutputFile, _ := filepath.Abs("dummyOutput.csv")

	cliArgs := []string{"dummyApp", absInputFile, absOutputFile}
	os.Args = cliArgs

	main()

	csvFile, err := os.Open(absOutputFile)

	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))
	line, err := reader.Read()

	if line[0] != "1" || line[1] != "3.47" {
		t.Errorf("Excpected output: %d,%f Got: %s,%s", 1, 3.47, line[0], line[1])
	}

	csvFile.Close()
	os.Remove(absOutputFile)
}