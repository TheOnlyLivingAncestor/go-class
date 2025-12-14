package sleepSort

import (
	"runtime"
	"sync"
	"time"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

func reverseSleepSort(input []uint) []uint {
	//kellenek workerek akiknek szétosztom az inputon lévő számokat
	if len(input) == 0 {
		return nil
	}
	workers := runtime.NumCPU()

	//kell egy channel ahova bepotyogtatják a számokat
	jobs := make(chan uint)

	//végén össze kell szedni a számokat egy másik channelen
	results := make(chan uint)

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//megkapja a worker a számot és sleep-elni kell vele
			for num := range jobs {
				ms := 500 - int(num)*10
				time.Sleep(time.Duration(ms) * time.Millisecond)
				results <- num
			}
		}()
	}

	//Be kell potyogtatni a számokat a channelre
	go func() {
		for _, t := range input {
			jobs <- t
		}
		close(jobs)
	}()

	//Bevárjuk a workereket
	go func() {
		wg.Wait()
		close(results)
	}()

	//és erről a második channelről kell kialakítani a visszatérési értéket
	final := make([]uint, 0, len(input))
	for v := range results {
		final = append(final, v)
	}

	return final

}
