package course

import (
	"fmt"
	"time"
)
func Object(){
	
	x:=8
	y := 7.0

	fmt.Println(x)
	fmt.Println(y)
	m := make(map[string] int)
	m["Key"]=7
	fmt.Println(m)

	s := []int{1,2,3}
	s = append(s, 5)
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}
	c := make(chan int)
	go doSomething(c)
	<-c

	g := 25

	fmt.Println(g)

	h := &g
	g=3
	fmt.Println(h)
	fmt.Println(*h)
}



func doSomething( c chan int){
	time.Sleep(3 * time.Second)
	fmt.Println("DOne")
	c <- 1
}