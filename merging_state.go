package main

import "sync"

func merge(channelBuffer int, channels []<-chan []string) <-chan []string {
	var waitGroup sync.WaitGroup

	mergedOut := make(chan []string, channelBuffer)

	output := func(channel <-chan []string) {
		for channelOutput := range channel {
			mergedOut <- channelOutput
		}

		waitGroup.Done()
	}

	waitGroup.Add(len(channels))

	for _, channel := range channels {
		go output(channel)
	}

	go func() {
		waitGroup.Wait()

		close(mergedOut)
	}()

	return mergedOut
}