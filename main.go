package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	seed = flag.Int("seed", 0, "seed for random generator. unix(now) be default")
	start = flag.Int("start", 1, "start of interval for random generator")
	end = flag.Int("end", 6, "end of interval for random generator")
	n = flag.Int("n", 1, "num of output values")
	norepeat = flag.Bool("norepeat", false, "no repeat for values in output values")
)

// Фукнция должна вернуть число из интервала [l,r]
func randInterval(l, r int) int{
	return rand.Intn(r-l)+l
}


func main() {
	flag.Parse()
	if *seed == 0 {
		rand.Seed(time.Now().UnixNano())
	} else {
		rand.Seed(int64(*seed))
	}
	if *start > *end{
		fmt.Println("ERROR")
	} else {
		if *norepeat {
			var x [10000]int
			j := *start
			for i := 0; j <= *end; i++ {
				x[i] = j
				j++
			}
			rand.Shuffle(*n, func(i, j int) {
				x[i], x[j] = x[j], x[i]
			})
			for i := 0; i < *n; i++ {
				fmt.Print(x[i], " ")
			}
		} else{
			for i := 0; i < *n; i++ {
				fmt.Println(randInterval(*start+0, *end+1))
			}
		}
	}
}

