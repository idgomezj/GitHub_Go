//if the program correctly prints "Found!" when a string that contains
//the characters ‘i’, ‘a’, and ‘n’, such as "iaaaan" is entered.

//if the program correctly prints "Not Found!" when a string that
//does not contain the characters ‘i’, ‘a’, and ‘n’ is entered.

package containString

import (
	"fmt"
	"strings"
)

func main() {
	var stringValue string
	fmt.Printf("Please give as a string input ")
	_, err := fmt.Scan(&stringValue)

	for {
		if err != nil {
			fmt.Println("There was an error updating the string!, Trying again...!")
			continue
		} else {
			if strings.Contains(stringValue, "i") &&
				strings.Contains(stringValue, "a") &&
				strings.Contains(stringValue, "n") {
				fmt.Println("Found!")
			} else {
				fmt.Println("Not Found!")
			}

		}
		break
	}
}
