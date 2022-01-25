package oop1

import (
	"fmt"
)

type Car struct {
	gas_pedal       uint16 //min 0 max 65535
	brake_pedal     uint16
	streering_wheel int16 // -32k to +32k
	top_speed_kmh   float64
}

func Structur_oop1() {
	x := 15
	a := &x
	b := *a
	b += 2
	fmt.Println(*a)
	fmt.Println(b)
	fmt.Println(*a)

	a_car := Car{gas_pedal: 22341, brake_pedal: 0, streering_wheel: 12561, top_speed_kmh: 225.0}
	fmt.Println(a_car)
}
