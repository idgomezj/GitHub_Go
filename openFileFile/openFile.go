//1. Opens a named text file. 2. Prints all first name/ last name pairs.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"encoding/json"
)

type Person struct {
	Name string `json:"name"`
	LastName string `json:"lastName"`
}


func main(){
	var p1 Person
	file, err := ioutil.ReadFile("test.json")
	if err!= nil{
		log.Fatal("Failed to generate json", err)
	}
	err2 := json.Unmarshal(file, &p1)
	if err2!= nil{
		log.Fatal("Failed to generate json", err)
	}

	fmt.Printf("The Json for person #1 is: %s \n", p1)


}