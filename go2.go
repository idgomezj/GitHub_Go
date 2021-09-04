package main

import (
	"fmt"
	"time"
)

func main() {
	go count("sheep")
	go count("fish")
	count("bear")
}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
