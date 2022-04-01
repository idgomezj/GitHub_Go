package course

import (
	"fmt"
	"time"
)


func Multiplexacion(){
	c1 := make(chan int)
	c2 := make(chan int)

	d1 := 4 * time.Second
	d2 := 2 * time.Second

	go doSomething3(d1, c1, 1)
	go doSomething3(d2, c2, 1)

	//fmt.Println(<-c1)
	//fmt.Println(<-c2)

	for i:= 0; i <2; i++{
		select{
		case channelMsg1 := <-c1:
			fmt.Println(channelMsg1)
		case channelMsg2 := <-c2:
			fmt.Println(channelMsg2)
		}
	}
}


func doSomething3(i time.Duration, c chan<- int, param int){
	time.Sleep(i)
	c <- param
}