package main

import (
	"fmt"
	"math"
)

var f func(float64) float64

func main(){
var acceleration, velocity, displacement, time float64
acceleration=10
velocity=2
displacement=1
time=3
f=GenDisplaceFn(acceleration,velocity,displacement)
fmt.Println(f(time))
}

func GenDisplaceFn(a,v,s float64) func(float64) float64{
	return func(t float64) float64 {
		return v*t+s+((math.Pow(t,2)*a)/2)
	}
}

