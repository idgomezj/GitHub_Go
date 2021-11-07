/* Write a program to sort an array of integers.
The program should partition the array into 4 parts, each of which is sorted by a different goroutine.
Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted
subarrays into one large sorted array. */

package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Printf("Can you tell us how many data your slice is going to have? ")
	fmt.Scan(&number)
	array := make([]int, number)
	for k := 0; k < number; k++ {
		fmt.Printf("Please write the %v number \n", k)
		fmt.Scan(&array[k])
	}
	fmt.Println("This is the array to sorted")
	fmt.Println(array)
	c := make(chan int)
	n := len(array) / 4
	for i := 1; i < 4; i++ {
		go sorted(array[n*(i-1):n*i], c)
	}
	go sorted(array[n*3:len(array)], c)
	for i := 1; i < 5; i++ {
		<-c
	}
	nn := len(array) / 2
	go sorted(array[0:nn-1], c)
	go sorted(array[nn:len(array)], c)
	for i := 1; i < 2; i++ {
		<-c
	}

	go sorted(array, c)
	<-c
	fmt.Println("This is the array sorted")
	fmt.Println(array)

}

func sorted(array []int, c chan int) {
	for j := 0; j < len(array); j++ {
		for k := 1; k < len(array); k++ {
			if array[k] < array[k-1] {
				array[k], array[k-1] = array[k-1], array[k]
			}
		}
	}
	fmt.Printf("Data after Sorted %v\n", array)

	c <- 2
}
