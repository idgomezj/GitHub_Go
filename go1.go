package main

import (
	"errors"
	"fmt"
	"math"
)

type person struct {
	name string
	age  int
}

func main() {
	x := 5
	var y int = 7
	var a [5]int
	a[2] = 7
	b := [5]int{1, 2, 3, 4, 5}
	sum := 0
	if x > 6 {
		sum = x + y
	} else {
		sum = 8
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}

	vertices := make(map[string]int)
	vertices["triangle"] = 3
	vertices["square"] = 4
	vertices["dodecagon"] = 12
	fmt.Println(sum)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(vertices)

	result := suma(3, 2)
	fmt.Println(result)

	result2, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result2)
	}
	ivan := 6
	inc(&ivan)
	fmt.Println(ivan)
}

func suma(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(x), nil

}

func inc(x *int) {
	*x++
}
