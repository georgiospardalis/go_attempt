package main

import "sync"

func merge(channelBuffer int, channels []<-chan []string) <-chan []string {
	var wg sync.WaitGroup

	mergedOut := make(chan []string, channelBuffer)

	output := func(channel <-chan []string) {
		for channelOutput := range channel {
			mergedOut <- channelOutput
		}

		wg.Done()
	}

	wg.Add(len(channels))

	for _, channel := range channels {
		go output(channel)
	}

	go func() {
		wg.Wait()

		close(mergedOut)
	}()

	return mergedOut
}