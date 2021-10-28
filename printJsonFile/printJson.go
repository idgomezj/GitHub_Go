// the program correctly prints a JSON object with keys ("name", "address") 
// and they are associated with the name and address that was entered.

package printJsonFile

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	Name string `json:"name"`
	Address string `json:"address"`
}

func printJson(){

var p1 Person
/* p1 := Person{ 
	 Name:"ivan",
	Address:"123",
} */
 fmt.Println("Please give us the person's Name")
 fmt.Scan(&p1.Name)
 fmt.Println("Please give us the person's Address")
 fmt.Scan(&p1.Address)

 fmt.Println(p1)

 barr, err := json.MarshalIndent(p1,"", "    ")

 if err != nil{
	 fmt.Println("There was an error created Json")
 }

 fmt.Printf("The Json for person #1 is: %s \n", string(barr))





}