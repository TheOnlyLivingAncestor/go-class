package wordcount

import (
	"runtime"
	"strings"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

func CountWords(texts []string) map[string]int {
	workers := runtime.NumCPU()

	jobs := make(chan string)
	results := make(chan map[string]int)

	//Workerek végzik a stringeken a számolást
	for i := 0; i < workers; i++ {
		go func() {
			local := make(map[string]int)
			for text := range jobs {
				words := strings.Fields(text)
				for _, w := range words {
					local[w]++
				}
			}
			results <- local
		}()
	}

	//Producer csinálja a feldolgozandó szövegeket a slice-ból
	go func() {
		for _, t := range texts {
			jobs <- t
		}
		close(jobs)
	}()

	final := make(map[string]int)
	for i := 0; i < workers; i++ {
		partial := <-results
		for w, c := range partial {
			final[w] += c
		}
	}

	return final
}
