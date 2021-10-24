package exercise

import (
	"fmt"
)

func go5() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Printf("\n")
	for i := 0; i < 5; i = i + 2 {
		fmt.Println(i)
	}

	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	for i := 0; i < 5; {
		fmt.Println(i)
		i++
	}
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
	j := 0
	for {
		fmt.Println(j)
		j++
		if j == 5 {
			break
		}
	}

	//loop:  then break Loop to break the first loop
	fmt.Println("ivan")
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}

	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
		"New York":   19745289,
	}

	for k, v := range statePopulations {
		fmt.Println(k, v)
	}
}
