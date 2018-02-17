package main

import "fmt"

type Int int

func (i Int) String() string {
	return "I"
}

type String string

func (s String) String() string {
	return "S"
}

type Bool bool

func (b Bool) String() string {
	return "B"
}

type Stringer interface {
	String() string
}

func receive(stringer Stringer) {
	switch stringer.(type) {
	case Int:
		fmt.Printf("Type: %s, value: %v\n", "int", stringer)
	case String:
		fmt.Printf("Type: %s, value: %v\n", "string", stringer)
	case Bool:
		fmt.Printf("Type: %s, value: %v\n", "bool", stringer)
	}
}
func main() {
	receive(String(200))
	receive(Int(200))
	receive(Bool(false))
}
