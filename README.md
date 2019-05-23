# GO Attempt
## About
Cab Ride Fare Estimation app

## Set Up
Before testing, building or installing the project,
place the unzipped folder under your workspace. It should look
like this:
`$GOPATH/src/github.com/georgiospardalis/beat_assignment`

## Testing the app
Just run `go test github.com/georgiospardalis/beat_assignment`

## Get an executable
In order to get the executable under your `$GOPATH/bin` directory
execute `go install github.com/georgiospardalis/beat_assignment`

## Running the app
Once you have generated the executable, use it as follows: 
`$beat_assignment [inputFile] [outputFile]`, where `inputFile`
and `outputFile` should be the absolute path to the input and
output files accordingly.

## About the app structure
The app is divided in 3 main  sections:
- reading from the input file
- processing data (estimating fares)
- write output data to specified file

This can be also seen through the pipeline where: 
- a channel is provided, with data being read line by line
(`reading_stage`)
- goroutines that "listen" to this channel process data,
  estimate fares and forward the end results to the output channel
  of their own (`processing_stage`)
- another goroutine merges the output of the aforementioned channels
  into a single one, which is in turn fed to (`merging_stage`)
- the last goroutine that outputs data to the
  requested file. (`writing_stage`)
  
The entry code snippet is `main.go` which bootstraps all the stages
and executes the writing one.
