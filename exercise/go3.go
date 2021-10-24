package exercise

import (
	"fmt"
	"reflect"
)

//Array and Slice
func go3() {
	a := []int{}
	// a := make([]int,10,100) len 10 capacity 100
	//ARRAY and SLICES.
	bb := [3]int{1, 2, 3} //ARRAY
	// b:= [...]int
	fmt.Println(bb)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	a = append(a, 1)
	a = append(a, []int{2, 3, 4, 5, 6, 7}...)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	// MAPS and STRUCTS
	// MAP
	// statePopulations := make(map[string]int)
	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
		"New York":   19745289,
	}
	statePopulations["ivan"] = 20
	delete(statePopulations, "New York")
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Texas"])

	pop, ok := statePopulations["New York"]
	fmt.Println(pop, ok)
	fmt.Println(len(statePopulations))

	//STRUCTS
	type Doctor struct {
		number     int
		actorName  string
		companions []string
	}

	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
		},
	}

	// anonymous aDoctor := struct{name string}{name: "John Pertwee"}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[1])

	// No inheritance, but can use Composition via embedding
	type Animal struct {
		Name   string `required max:"100"` //This is a Tag.
		Origin string
	}

	type Bird struct {
		Animal
		SpeedKPH float32
		CanFly   bool
	}

	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)

	//TAG

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
