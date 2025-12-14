package concurrentprimes

import "sort"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// Számokat generálunk 2-től n-ig, és a számokat a channelre kiküldjük
func generate(n int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 2; i <= n; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

// Bevesszük az előző channelről a számot, és ha az eddig összegyűlt prímekkel maradék nélkül osztható, akkor nem prím
// a prímeket visszeküldjük egy másik channelre
func filter(in <-chan int, prime int) <-chan int {
	out := make(chan int)
	go func() {
		for x := range in {
			if x%prime != 0 {
				out <- x
			}
		}
		close(out)
	}()
	return out
}

func GeneratePrimes(n int) []int {
	ch := generate(n)
	primes := []int{}

	for {
		p, ok := <-ch
		if !ok {
			break
		}
		primes = append(primes, p)
		ch = filter(ch, p)
	}

	sort.Ints(primes)

	return primes

}
