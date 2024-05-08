package main

import "fmt"

// Kommentar für die Mensch

//go:generate gogen stack.gogen int
//go:generate gogen stack.gogen myType

type myType struct {
	a, b string
}

func main() {
	is := intStack{}
	is.Push(1)
	is.Push(2)
	is.Push(3)
	fmt.Println(is.Pop())
	fmt.Println(is.Pop())
	fmt.Println(is.Pop())
	fmt.Println(is.Pop())
	fmt.Println(is.Pop())
	fmt.Println(is.Pop())
	ms := myTypeStack{}
	ms.Push(myType{"a", "b"})
}
