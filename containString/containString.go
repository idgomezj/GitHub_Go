//if the program correctly prints "Found!" when a string that contains
//the characters ‘i’, ‘a’, and ‘n’, such as "iaaaan" is entered.

//if the program correctly prints "Not Found!" when a string that
//does not contain the characters ‘i’, ‘a’, and ‘n’ is entered.

package containString

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func containString() {
	//var stringValue string
	//var stdin *bufio.Reader

	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Print("Please give as a string input.... \n")

	stringValue, err := consoleReader.ReadString('\n')
	for {
		if err != nil {
			fmt.Println("There was an error updating the string!, Trying again...!")
			continue
		} else {
			stringValue = strings.ToUpper(stringValue)
			//fmt.Print(string(stringValue[len(stringValue)-2]))
			if string(stringValue[0]) == "I" &&
				strings.Contains(stringValue, "A") &&
				string(stringValue[len(stringValue)-2]) == "N" {
				fmt.Println("Found!")
			} else {
				fmt.Println("Not Found!")
			}

		}
		break
	}
}
