package main

import (
	"fmt"
	"log"
)

type AnimalInterface interface{
	Eat()
	Move()
	Speak()
}

type AnimalStruct struct {
	Food string
	Locomtion string
	Spoke string
}


func (a AnimalStruct) Eat() {
	fmt.Println(a.Food)
}
func (a AnimalStruct) Move() {
	fmt.Println(a.Locomtion)
}
func (a AnimalStruct) Speak() {
	fmt.Println(a.Spoke)
}




func main(){
	var activity string
	var name string
	var kindAnimal string
	var animalAction string
	saveInterface:= make(map[string] AnimalInterface)


	for true{
	fmt.Println("> Do you want to created ne animal (newanimal) or consult an animal (query)")
	fmt.Print("> ")
	_, err := fmt.Scan(&activity)
	if err!=nil{
		log.Fatalln("There was an error")
	}


	if activity == "newanimal"{
	fmt.Println("> What is the name?")
	fmt.Print("> ")
	_, err2 := fmt.Scan(&name)
	if err2!=nil{
		log.Fatalln("There was an error")
	}

	fmt.Println("> What kind of animal would youlike to creat? [Cow, Bird or Snake]")
	fmt.Print("> ")
	_, err3 := fmt.Scan(&kindAnimal)
	if err3!=nil{
		log.Fatalln("There was an error")
	}

	switch kindAnimal {
		case "Cow":
			a:= AnimalStruct{"grass","walk","moo"}
			var a1 AnimalInterface
			a1=a
			saveInterface[name]=a1
			fmt.Println("Created it!")
		case "Bird":
			a:= AnimalStruct{"worms","fly","peep"}
			var a1 AnimalInterface
			a1=a
			saveInterface[name]=a1
			fmt.Println("Created it!")
		case "Snake":
			a:= AnimalStruct{"mice","slither","hsss"}
			var a1 AnimalInterface
			a1=a
			saveInterface[name]=a1
			fmt.Println("Created it!")
	}
	


	}else if activity == "query"{

	fmt.Println("> Please give us the name")
	fmt.Print("> ")
	_, err4 := fmt.Scan(&name)
	if err4!=nil{
		log.Fatalln("There was an error")
	}
	nameInterface, ok := saveInterface[name]
	if !ok {
		log.Fatalln("This name does not existed")
	}


	fmt.Println("> Please tell us the action you want the animal does [Eat, Move, Speak]")
	fmt.Print("> ")
	_, err5 := fmt.Scan(&animalAction)
	if err5!=nil{
		log.Fatalln("There was an error")
	}

	
	switch animalAction{
	case "Eat":
		nameInterface.Eat()
	case "Move":
		nameInterface.Move()
	case "Speak":
		nameInterface.Speak()
	}



	}else{
		log.Fatalln("The process selected is not defined")
	}

	}

}