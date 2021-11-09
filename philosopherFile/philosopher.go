/* Write a program to sort an array of integers.
The program should partition the array into 4 parts, each of which is sorted by a different goroutine.
Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted
subarrays into one large sorted array. */

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	firstCS, secondCS *ChopS
}

func main() {
	CSticks := make([]*ChopS, 5)
	var wg sync.WaitGroup
	var x, y int

	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)

	wg.Add(5)
	for i := 0; i < 5; i++ {
		x = rand.Intn(5)
		for {
			y = rand.Intn(5)
			if x != y {
				break
			}
		}
		philos[i] = &Philo{CSticks[x], CSticks[y]}
		go philos[i].eat(i, x, y, &wg)
	}
	wg.Wait()
	fmt.Println("All philosophers already eaten")

}

func (p Philo) eat(number int, x int, y int, wg *sync.WaitGroup) {
	for i := 1; i < 4; i++ {
		p.firstCS.Lock()
		p.secondCS.Lock()

		fmt.Printf("starting to eat #%v and it is %v time \n", number+1, i)
		fmt.Printf("Using  %v and  %v CSticks \n", x, y)
		time.Sleep(1 * time.Second)
		fmt.Printf("finishing eating  #%v \n", number+1)
		p.firstCS.Unlock()
		p.secondCS.Unlock()

	}
	wg.Done()
}
