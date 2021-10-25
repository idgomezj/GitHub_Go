//if the program correctly prints the sorted slice after entering three distinct integers.
// **Points are awarded incrementally each time that an integer is added and it correctly prints the sorted slice. 

//the program correctly prints the sorted slice after entering four distinct integers.
// **Points are awarded if it correctly prints the sorted slice after adding the fourth integer.


package sortedSliceFile

import (
	"fmt"
)

func sortedSlice(){

	
	var backup int
	var level int
	var number int


	fmt.Printf("Can you tell us how many data your slice is going to have? ")
	fmt.Scan(&number)

	sli:= make([]int,number)

	for k:=0;k<number;k++{
		fmt.Printf("Please write the %v number \n",k)
		fmt.Scan(&sli[k])
	}


	for i:=0; i < len(sli);i++{
		for j:= 1; j<len(sli)-level;j++{
			if sli[j-1]>sli[j]{
				backup=sli[j]
				sli[j]=sli[j-1]
				sli[j-1]=backup
			}
		}
	level+=1
	}

	fmt.Printf("This is the Slice Sorted %v\n",sli)



}