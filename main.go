//i³+j³+k³ = ijk
//i.e. 166³ + 500³ + 333³ = 166500333
package main

import (
	"flag"
	"fmt"
	"runtime"
	"sync"
	"time"
)

var chunks int
var max int

func check(data ijk) {
	isOk, word := bingo(data)
	if isOk {
		fmt.Println(data, "=", word)
	}
}

func doIt(from int, to int, max int, wg *sync.WaitGroup) {
	// fmt.Println("chunk:", from, "-", to)
	defer wg.Done()
	for i := from; i < to; i++ {
		for j := 1; j < max; j++ {
			for k := 1; k < max; k++ {
				check(ijk{i, j, k})
			}
		}
	}
}

func init() {
	cores := runtime.GOMAXPROCS(0)
	fmt.Println("cores::", cores)
	flag.IntVar(&chunks, "ch", cores*2, "Chunks count")
	flag.IntVar(&max, "max", 999, "Max number")
	flag.Parse()
	fmt.Println("chunks::", chunks)
	fmt.Println("max::", max)
	fmt.Println("")
}

func main() {
	chunksize := max / chunks
	chunkrest := max % chunks

	var wg sync.WaitGroup
	start := time.Now()

	chunklasti := 0
	if chunksize > 0 {
		wg.Add(chunks)
		for i := 1; i <= chunks; i++ {
			from := chunklasti + 1
			chunklasti = i * chunksize
			to := chunklasti
			go doIt(from, to, max, &wg)
		}
	}
	if chunkrest > 0 {
		wg.Add(1)
		go doIt(chunklasti+1, chunklasti+chunkrest, max, &wg)
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("")
	fmt.Println("elapsed=", elapsed)
}
