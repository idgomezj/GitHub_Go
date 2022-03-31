package course

import "fmt"

func Chanels() {
	c := make(chan int, 3)

	c <- 1
	c <- 2
	c <- 3


	fmt.Println( <- c)

}