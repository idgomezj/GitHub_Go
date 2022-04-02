package course

import "fmt"

type Person struct {
	name string
	age int
}

type Employee struct {
	id int
	Person
}

type PrintInfo interface {
	getMessage() string
}



func (employee Employee) getMessage() string{
	return "Este es un Employee"
}


func (e *Employee) SetId(id int){
	e.id = id
}

func (e *Employee) SetName(name string){
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func Object2(){
	e:= Employee{}
	fmt.Println(e)
	e.SetId(3)
	e.SetName("Dario")
	e.SetName("Lore")
	fmt.Println(e)
}