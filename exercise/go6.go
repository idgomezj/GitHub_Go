package exercise

import (
	"fmt"
)

func go6() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("hello Go!"))
}

//INTERFACE

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
