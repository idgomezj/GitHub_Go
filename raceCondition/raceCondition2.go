package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
With this GOMAXPROCS we are requesting an application transition to multiple cores.
And the go keywords added before the function is executed,
can be executed separately on different cores, increasing application performance.
*/

func main() {
	// added paralelism. Uncomment to see  difference in timing

	// runtime.GOMAXPROCS(10)
	// runtime.GOMAXPROCS(2)
	runtime.GOMAXPROCS(5)
	start := time.Now()
	go func() {
		// func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}()

	go func() {
		for i := 4; i < 10; i++ {
			fmt.Println(i)
		}
	}()

	elapsedTime := time.Since(start)

	fmt.Println("Total Time For Execution: " + elapsedTime.String())
	// go fmt.Println("Total Time For Execution: " + elapsedTime.String())

	time.Sleep(time.Second)
}
