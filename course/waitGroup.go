package course

import (
	"fmt"
	"sync"
	"time"
)

func Wait() {

	var wg sync.WaitGroup

	for i := 0; i <10; i++{
		wg.Add(1)
		go doSomething2(i, &wg)
	}
	wg.Wait()
}

func doSomething2(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("start %d\n",i)
	time.Sleep(2 * time.Second)
	fmt.Printf("finish %d\n",i)
}