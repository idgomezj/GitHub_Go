package oop1

import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

func (c Car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax)
}

func (c Car) mmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax / kmh_multiple)
}

func Structur_oop2() {
	a_car2 := Car{gas_pedal: 22341, brake_pedal: 0, streering_wheel: 12561, top_speed_kmh: 225.0}
	fmt.Println(a_car2.kmh())
	fmt.Println(a_car2.mmh())

}
