package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("The test is true")
	}

	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
		"New York":   19745289,
	}

	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
	}

	// "=="
	// ">"
	// "<"
	// "!="
	// and -> &&
	// or -> ||
	// not operator -> !

	number := 50
	guess := 30
	if guess < 1 || returnTrue() || guess > 100 {
		fmt.Println("The guess mus be between 1 and 100!")
	} else if guess > 10000 {
		fmt.Println("The value is too high")
	} else {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You got it!")
		}
	}

	fmt.Println(!true)
	// switch i:= 2+3; i {}
	switch 2 {
	case 1, 5:
		fmt.Println("one or five")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("Not one, five or two")
	}

	ii := 10
	switch {
	case ii <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough
	case ii <= 20:
		fmt.Println("greater than twenty")
	default:
		fmt.Println("greater than twenty")

	}

	var j interface{} = 1
	switch j.(type) {
	case int:
		fmt.Println("j is an int")
	case float64:
		fmt.Println("j is a float64")
	default:
		fmt.Println("j is another type")
	}

}

func returnTrue() bool {
	fmt.Println("Return True")
	return true
}
