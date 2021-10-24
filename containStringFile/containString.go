package intToFloating

//if the program correctly prints the truncated integer when a floating point number is entered.
import (
	"fmt"
)

func intToFloating() {
	var value float64
	var integerValue int

	fmt.Printf("Please, give a float number ")
	_, err := fmt.Scanf("%f", &value)
	//fmt.Println(err)
	if err != nil {
		fmt.Println("The value given was wrong!")
	} else {
		integerValue = int(value)
		fmt.Printf("Integer value is %v\n", integerValue)
	}

}
