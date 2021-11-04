package main

import (
	"fmt"
	"time"
)

var ValueA int = 0

func printValueA() {
	for {
		var val int = ValueA
		fmt.Println(val)
		time.Sleep(10 * time.Second)
	}
}

func writeValueA() {
	for {
		ValueA = ValueA + 1
	}
}

func main() {
	go writeValueA()
	go printValueA()
	time.Sleep(10 * time.Second)
}
