package main

import (
	"fmt"
)

func main(){
	var number int
	fmt.Printf("Can you tell us how many data your slice is going to have? ")
	fmt.Scan(&number)

	sli:= make([]int,number)

	for k:=0;k<number;k++{
		fmt.Printf("Please write the %v number \n",k)
		fmt.Scan(&sli[k])
	}
	BubbleSort(sli)
	fmt.Println(sli)
}

func BubbleSort(x []int){
	level:=len(x)
	for i:=0; i<len(x)-1; i++{
		for j:=0; j<level-1; j++{
			if x[j]>x[j+1]{
				Swap(x, j)
			}
		}
		level-=1
	}	
}

func Swap(x []int, i int){
	x[i],x[i+1]=x[i+1],x[i]	
}