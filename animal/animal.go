package main

import (
	"fmt"
)

type Animal struct{
	food string
	locomotion string
	noise string	
}


func main(){
	for true {
		var animalValue string
		var actionValue string
		fmt.Println("> Please give us the ananimal according to the following options [cow, bird, snake]")
		fmt.Print("> ")
		_ ,err:=fmt.Scan(&animalValue)
		if err != nil{
			fmt.Println("> There was an error, try it again!")
			continue
		}

		fmt.Println("> Please give us the action to evaluated according to the following options [Eat, Move, Speak]")
		fmt.Print("> ")
		_ ,err2:=fmt.Scan(&actionValue)
		if err2 != nil{
			fmt.Println("> There was an error, try it again!")
			continue
		}

		funcs:= map[string] func(a Animal){
			"Eat": Eat,
			"Move": Move,
			"Speak": Speak,
		}

		switch animalValue{
		case "cow":
			funcs[actionValue](Animal{"grass","walk","moo"})
		case "bird":
			funcs[actionValue](Animal{"worms","fly","peep"})
		case "snake":
			funcs[actionValue](Animal{"mice","slither","hsss"})

		}


	}
}

func Eat(a Animal) {
	fmt.Println(a.food)
}

func Move(a Animal) {
	fmt.Println(a.locomotion)
}

func Speak(a Animal) {
	fmt.Println(a.noise)
}
